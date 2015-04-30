package handler

import (
	"github.com/pokoetea/gomemcache/datastore"
	"github.com/pokoetea/gomemcache/protocol"
)

type SetHandler struct {
	ds datastore.DataStore
}

func NewSetHandler(ds datastore.DataStore) Handler {
	return &SetHandler{
		ds: ds,
	}
}

func (h *SetHandler) Handle(req *protocol.Packet, resChan chan<- *protocol.Packet) {
	defer response(req, resChan)
	if len(req.Extras) != 8 {
		return
	}
	flags := req.Extras[0:4]
	e := req.Extras[4:8]
	expiry := uint32(e[3]) | uint32(e[2])<<8 | uint32(e[1])<<16 | uint32(e[0])<<24
	cas, _ := h.ds.Set(req.Key, req.Value, flags, expiry)
	req.Header.SetCAS(cas)
}
