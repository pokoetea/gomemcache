package datastore

import (
	"os"
	"sync/atomic"
	"testing"
)

var ds DataStore

func TestMain(m *testing.M) {
	ds = NewMutexMap()
	os.Exit(m.Run())
}

func BenchmarkSingleKeySetGet(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := []byte("singlekey")
			value := []byte("singlevalue")
			_, err := ds.Set(key, value, nil, 0)
			if err != nil {
				b.Fatalf("set err: %v", err)
			}
			res, _, err := ds.Get(key)
			if err != nil {
				b.Fatalf("get err: %v", err)
			}
			if string(res) != string(value) {
				b.Fatalf("wrong: real=%v, expect=%v", string(res), value)
			}
		}
	})
}

func BenchmarkIncrementalKeySetGet(b *testing.B) {
	var counter uint64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := atomic.AddUint64(&counter, 1)
			key := []byte("key_" + string(i))
			value := []byte("value_" + string(i))
			_, err := ds.Set(key, value, nil, 0)
			if err != nil {
				b.Fatalf("set err: %v", err)
			}
			res, _, err := ds.Get(key)
			if err != nil {
				b.Fatalf("get err: %v", err)
			}
			if string(res) != string(value) {
				b.Fatalf("wrong: real=%v, expect=%v", string(res), value)
			}
		}
	})
}
