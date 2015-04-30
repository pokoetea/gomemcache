package server

import (
	"net"

	"github.com/pokoetea/gomemcache/datastore"
)

type Server struct {
	dispatcher Dispatcher
}

func New() *Server {
	ds := datastore.NewMutexMap()
	return &Server{
		dispatcher: NewSliceDispatcher(ds),
	}
}

func (s *Server) Run() error {
	l, err := net.Listen("tcp", ":11211")
	if err != nil {
		return err
	}
	for {
		c, err := l.Accept()
		if err != nil {
			// TODO: handle error
			continue
		}
		conn := NewConn(c, s.dispatcher)
		go conn.Run()
	}
}
