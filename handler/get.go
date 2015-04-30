package handler

import (
	"github.com/pokoetea/gomemcache/datastore"
	"github.com/pokoetea/gomemcache/protocol"
)

type GetHandler struct {
	ds datastore.DataStore
}

func NewGetHandler(ds datastore.DataStore) Handler {
	return &GetHandler{
		ds: ds,
	}
}

func (h *GetHandler) Handle(req *protocol.Packet, resChan chan<- *protocol.Packet) {
	defer response(req, resChan)
	value, flags, _ := h.ds.Get(req.Key)
	req.Extras = flags
	req.Value = value
	req.Header.SetExtrasLength(uint8(len(flags)))
	req.Header.SetTotalBodyLength(uint32(len(req.Key) + len(flags) + len(value)))
}
