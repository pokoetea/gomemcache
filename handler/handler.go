package handler

import "github.com/pokoetea/gomemcache/protocol"

type Handler interface {
	Handle(req *protocol.Packet, resChan chan<- *protocol.Packet)
}

func response(p *protocol.Packet, resChan chan<- *protocol.Packet) {
	p.Header.SetMagic(protocol.MagicRes)
	resChan <- p
}
