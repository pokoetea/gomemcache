package server

import (
	"github.com/pokoetea/gomemcache/datastore"
	"github.com/pokoetea/gomemcache/handler"
	"github.com/pokoetea/gomemcache/protocol"
)

type Dispatcher interface {
	Dispatch(packet *protocol.Packet, resChan chan<- *protocol.Packet)
}

type SliceDispatcher struct {
	handlers []handler.Handler
}

func NewSliceDispatcher(ds datastore.DataStore) Dispatcher {
	handlers := make([]handler.Handler, int(protocol.LastCmd)+1)
	handlers[protocol.CmdGet] = handler.NewGetHandler(ds)
	handlers[protocol.CmdSet] = handler.NewSetHandler(ds)
	return &SliceDispatcher{
		handlers: handlers,
	}
}

func (s *SliceDispatcher) Dispatch(packet *protocol.Packet, resChan chan<- *protocol.Packet) {
	opCode := packet.Header.OpCode()
	if opCode > protocol.LastCmd {
		return
	}
	s.handlers[opCode].Handle(packet, resChan)
}
