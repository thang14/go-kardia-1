/*
 *  Copyright 2018 KardiaChain
 *  This file is part of the go-kardia library.
 *
 *  The go-kardia library is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU Lesser General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  The go-kardia library is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 *  GNU Lesser General Public License for more details.
 *
 *  You should have received a copy of the GNU Lesser General Public License
 *  along with the go-kardia library. If not, see <http://www.gnu.org/licenses/>.
 */

package blockchain

import (
	"sync/atomic"

	lru "github.com/hashicorp/golang-lru"
	"github.com/kardiachain/go-kardia/configs"
	"github.com/kardiachain/go-kardia/kai/kaidb"
	"github.com/kardiachain/go-kardia/kai/storage/kvstore"
	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/lib/log"
	"github.com/kardiachain/go-kardia/types"
)

const (
	headerCacheLimit = 512
	heightCacheLimit = 2048
)

//TODO(huny@): Add detailed description
type HeaderChain struct {
	config *configs.ChainConfig

	kaiDb types.StoreDB

	genesisHeader *types.Header

	currentHeader     atomic.Value // Current head of the header chain (may be above the block chain!)
	currentHeaderHash common.Hash  // Hash of the current head of the header chain (prevent recomputing all the time)

	headerCache *lru.Cache // Cache for the most recent block headers
	heightCache *lru.Cache // Cache for the most recent block height
}

// CurrentHeader retrieves the current head header of the canonical chain. The
// header is retrieved from the HeaderChain's internal cache.
func (hc *HeaderChain) CurrentHeader() *types.Header {
	return hc.currentHeader.Load().(*types.Header)
}

// NewHeaderChain creates a new HeaderChain structure.
//  getValidator should return the parent's validator
//  procInterrupt points to the parent's interrupt semaphore
//  wg points to the parent's shutdown wait group
func NewHeaderChain(kaiDb types.StoreDB, config *configs.ChainConfig) (*HeaderChain, error) {
	headerCache, _ := lru.New(headerCacheLimit)
	heightCache, _ := lru.New(heightCacheLimit)

	hc := &HeaderChain{
		config:      config,
		kaiDb:       kaiDb,
		headerCache: headerCache,
		heightCache: heightCache,
	}

	hc.genesisHeader = hc.GetHeaderByHeight(0)
	if hc.genesisHeader == nil {
		return nil, ErrNoGenesis
	}

	hc.currentHeader.Store(hc.genesisHeader)
	if head := kaiDb.ReadHeadBlockHash(); head != (common.Hash{}) {
		if chead := hc.GetHeaderByHash(head); chead != nil {
			hc.currentHeader.Store(chead)
		}
	}
	hc.currentHeaderHash = hc.CurrentHeader().Hash()

	return hc, nil
}

// GetHeaderByHeight retrieves a block header from the database by height,
// caching it (associated with its hash) if found.
func (hc *HeaderChain) GetHeaderByHeight(height uint64) *types.Header {
	hash := hc.kaiDb.ReadCanonicalHash(height)
	if hash == (common.Hash{}) {
		return nil
	}
	return hc.GetHeader(hash, height)
}

// GetHeader retrieves a block header from the database by hash and height,
// caching it if found.
func (hc *HeaderChain) GetHeader(hash common.Hash, height uint64) *types.Header {
	// Short circuit if the header's already in the cache, retrieve otherwise
	if header, ok := hc.headerCache.Get(hash); ok {
		return header.(*types.Header)
	}
	header := hc.kaiDb.ReadHeader(height)
	if header == nil {
		return nil
	}
	// Cache the found header for next time and return
	hc.headerCache.Add(hash, header)
	return header
}

// GetHeaderByHash retrieves a block header from the database by hash, caching it if
// found.
func (hc *HeaderChain) GetHeaderByHash(hash common.Hash) *types.Header {
	height := hc.GetBlockHeight(hash)
	if height == nil {
		return nil
	}
	return hc.GetHeader(hash, *height)
}

// GetBlockHeight retrieves the block height belonging to the given hash
// from the cache or database
func (hc *HeaderChain) GetBlockHeight(hash common.Hash) *uint64 {
	if cached, ok := hc.heightCache.Get(hash); ok {
		height := cached.(uint64)
		return &height
	}
	height := hc.kaiDb.ReadHeaderHeight(hash)
	if height != nil {
		hc.heightCache.Add(hash, *height)
	}
	return height
}

// SetCurrentHeader sets the current head header of the canonical chain.
func (hc *HeaderChain) SetCurrentHeader(head *types.Header) {
	hc.currentHeader.Store(head)
	hc.currentHeaderHash = head.Hash()
}

// SetGenesis sets a new genesis block header for the chain
func (hc *HeaderChain) SetGenesis(head *types.Header) {
	hc.genesisHeader = head
}

type (
	// UpdateHeadBlocksCallback is a callback function that is called by SetHead
	// before head header is updated. The method will return the actual block it
	// updated the head to (missing state) and a flag if setHead should continue
	// rewinding till that forcefully (exceeded ancient limits)
	UpdateHeadBlocksCallback func(kaidb.KeyValueWriter, *types.Header) (uint64, bool)

	// DeleteCallback is a callback function that is called by SetHead before
	// each header is deleted.
	DeleteCallback func(kaidb.KeyValueWriter, kaidb.KeyValueReader, common.Hash, uint64)
)

// SetHead rewinds the local chain to a new head. Everything above the new head
// will be deleted and the new one set.
func (hc *HeaderChain) SetHead(head uint64, delFn DeleteCallback) {
	var (
		chainDb    = hc.kaiDb.DB()
		parentHash common.Hash
		batch      = chainDb.NewBatch()
		origin     = true
	)
	for hdr := hc.CurrentHeader(); hdr != nil && hdr.Height > head; hdr = hc.CurrentHeader() {
		num := hdr.Height

		// Rewind block chain to new head.
		parent := hc.GetHeader(hdr.LastBlockID.Hash, num-1)
		if parent == nil {
			parent = hc.genesisHeader
		}
		parentHash = parent.Hash()

		// Update head first(head fast block, head full block) before deleting the data.
		markerBatch := chainDb.NewBatch()

		// Update head blcok then.
		hc.kaiDb.WriteHeadBlockHash(parentHash)
		if err := markerBatch.Write(); err != nil {
			log.Crit("Failed to update chain markers", "error", err)
		}
		hc.currentHeader.Store(parent)
		hc.currentHeaderHash = parentHash

		// If this is the first iteration, wipe any leftover data upwards too so
		// we don't end up with dangling daps in the database
		var nums []uint64
		if origin {
			for n := num + 1; len(kvstore.ReadAllHashes(chainDb, n)) > 0; n++ {
				nums = append([]uint64{n}, nums...) // suboptimal, but we don't really expect this path
			}
			origin = false
		}
		nums = append(nums, num)

		// Remove the related data from the database on all sidechains
		for _, num := range nums {
			// Gather all the side fork hashes
			hashes := kvstore.ReadAllHashes(chainDb, num)
			if len(hashes) == 0 {
				// No hashes in the database whatsoever, probably frozen already
				hashes = append(hashes, hdr.Hash())
			}
			for _, hash := range hashes {
				if delFn != nil {
					delFn(batch, hc.kaiDb.DB(), hash, num)
				}
				kvstore.DeleteHeader(batch, hash, num)
			}
			kvstore.DeleteCanonicalHash(batch, num)
		}
	}
	// Flush all accumulated deletions.
	if err := batch.Write(); err != nil {
		log.Crit("Failed to rewind block", "error", err)
	}

	// Clear out any stale content from the caches
	hc.headerCache.Purge()
	hc.heightCache.Purge()
}
