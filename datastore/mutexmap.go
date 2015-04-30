package datastore

import "github.com/pokoetea/gomemcache/protocol"

type MutexMap struct {
	// *sync.RWMutex
	data map[string]*Item
}

func NewMutexMap() *MutexMap {
	return &MutexMap{
		// RWMutex: new(sync.RWMutex),
		data: make(map[string]*Item, 1024),
	}
}

func (m *MutexMap) Get(key []byte) (value []byte, flags []byte, err error) {
	k := protocol.BytesToString(key)
	// m.RLock()
	// defer m.RUnlock()
	item, ok := m.data[k]
	if !ok {
		err = ErrKeyNotFound
		return
	}
	value = item.Value
	flags = item.Flags
	return
}

func (m *MutexMap) Set(key []byte, value []byte, flags []byte, expiry uint32) (cas uint64, err error) {
	k := protocol.BytesToString(key)
	// m.Lock()
	// defer m.Unlock()
	item, ok := m.data[k]
	if !ok {
		item = &Item{}
		m.data[k] = item
	}
	item.Value = value
	item.Flags = flags
	item.Expiry = expiry
	item.CAS++
	cas = item.CAS
	return
}

var _ DataStore = (*MutexMap)(nil)
