package server

import (
	"fmt"
	"os"
	"sync/atomic"
	"testing"

	"github.com/dustin/gomemcached/client"
)

var c *memcached.Client

func TestMain(m *testing.M) {
	client, err := memcached.Connect("tcp", ":11211")
	if err != nil {
		os.Exit(1)
	}
	c = client
	defer c.Close()
	os.Exit(m.Run())
}

func main() {
	c, err := memcached.Connect("tcp", ":11211")
	if err != nil {
		fmt.Println(err)
		return
	}
	key := "test"
	value := []byte("test")
	res, err := c.Set(0, key, 0, 0, value)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("res:", res)
	res, err = c.Get(0, key)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("res:", string(res.Body))
}

func BenchmarkSingleKeySetGet(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := "singlekey"
			value := []byte("singlevalue")
			res, err := c.Set(0, key, 0, 0, value)
			if err != nil {
				b.Fatalf("set err: %v", err)
			}
			res, err = c.Get(0, key)
			if err != nil {
				b.Fatalf("get err: %v", err)
			}
			if string(res.Body) != string(value) {
				b.Fatalf("wrong: real=%v, expect=%v", string(res.Body), value)
			}
		}
	})
}

func BenchmarkIncrementalKeySetGet(b *testing.B) {
	var counter uint64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := atomic.AddUint64(&counter, 1)
			key := "key_" + string(i)
			value := []byte("value_" + string(i))
			res, err := c.Set(0, key, 0, 0, value)
			if err != nil {
				b.Fatalf("set err: %v", err)
			}
			res, err = c.Get(0, key)
			if err != nil {
				b.Fatalf("get err: %v", err)
			}
			if string(res.Body) != string(value) {
				b.Fatalf("wrong: real=%v, expect=%v", string(res.Body), value)
			}
		}
	})
}
