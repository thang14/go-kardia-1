package store

import (
	"github.com/kardiachain/go-kardia/kai/kaidb"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

const (
	baseKeyPending   = byte(0x01)
	baseKeyCompleted = byte(0x02)
)

type Store struct {
	db kaidb.Database
}

func (s *Store) GetPendingDeposit(hash []byte) (*dproto.Deposit, error) {
	buf, err := s.db.Get(keyPending(&dproto.Deposit{Hash: hash}))
	if err != nil {
		return nil, err
	}
	deposit := new(dproto.Deposit)

	if err := deposit.Unmarshal(buf); err != nil {
		return nil, err
	}
	return deposit, nil
}

func (s *Store) isCompleted(deposit *dproto.Deposit) bool {
	ok, err := s.db.Has(keyCompleted(deposit))
	if err != nil {
		panic(err)
	}
	return ok
}

func (s *Store) isPending(deposit *dproto.Deposit) bool {
	ok, err := s.db.Has(keyPending(deposit))
	if err != nil {
		panic(err)
	}
	return ok
}

func (s *Store) SetDeposit(deposit *dproto.Deposit) error {
	if s.isCompleted(deposit) {
		return nil
	}

	buf, err := deposit.Marshal()
	if err != nil {
		return err
	}
	return s.db.Put(keyPending(deposit), buf)
}

func (s *Store) PendingDeposit() ([]*dproto.Deposit, error) {
	return s.listDeposit(baseKeyPending)
}

func (s *Store) listDeposit(keyPrefix byte) ([]*dproto.Deposit, error) {
	iter := s.db.NewIterator([]byte{keyPrefix}, nil)
	deposits := make([]*dproto.Deposit, 0)
	for iter.Next() {
		var dpv *dproto.Deposit
		if err := dpv.Unmarshal(iter.Value()); err != nil {
			return nil, err
		}
		deposits = append(deposits, dpv)
	}
	return deposits, nil
}

func (s *Store) MarkDepositCompleted(deposit *dproto.Deposit) error {
	if err := s.db.Delete(keyPending(deposit)); err != nil {
		return err
	}
	return s.db.Put(keyCompleted(deposit), []byte{0x01})
}

func keyPending(deposit *dproto.Deposit) []byte {
	return append([]byte{baseKeyPending}, deposit.Hash...)
}

func keyCompleted(deposit *dproto.Deposit) []byte {
	return append([]byte{baseKeyCompleted}, deposit.Hash...)
}
