package datastore

import "errors"

var (
	ErrKeyNotFound = errors.New("key not found")
)

type DataStore interface {
	Get(key []byte) (value []byte, flags []byte, err error)
	Set(key []byte, value []byte, flags []byte, expiry uint32) (cas uint64, err error)
}

type Item struct {
	Value  []byte
	Flags  []byte
	Expiry uint32
	CAS    uint64
}
