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

package consensus

import (
	"github.com/kardiachain/go-kardia/kai/state/cstate"
	"github.com/kardiachain/go-kardia/lib/common"
	stypes "github.com/kardiachain/go-kardia/mainchain/staking/types"
	"github.com/kardiachain/go-kardia/types"
)

type BaseBlockOperations interface {
	Base() uint64
	Height() uint64
	LoadBlock(height uint64) *types.Block
	LoadBlockCommit(height uint64) *types.Commit
	LoadSeenCommit(height uint64) *types.Commit
	CreateProposalBlock(height uint64, state cstate.LatestBlockState, proposerAddr common.Address, commit *types.Commit) (*types.Block, *types.PartSet)
	CommitAndValidateBlockTxs(block *types.Block, lastCommit stypes.LastCommitInfo, byzVals []stypes.Evidence) ([]*types.Validator, common.Hash, error)
	CommitBlockTxsIfNotFound(block *types.Block, lastCommit stypes.LastCommitInfo, byzVals []stypes.Evidence) ([]*types.Validator, common.Hash, error)
	SaveBlock(block *types.Block, partSet *types.PartSet, seenCommit *types.Commit)
	LoadBlockPart(height uint64, index int) *types.Part
	LoadBlockMeta(height uint64) *types.BlockMeta
}
