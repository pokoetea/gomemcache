package cli

import (
	"github.com/pokoetea/gomemcache/server"
)

type ServerCommand struct {
}

func NewServerCommand() *ServerCommand {
	return &ServerCommand{}
}

func (s *ServerCommand) Run() error {
	server := server.New()
	return server.Run()
}

var _ Command = (*ServerCommand)(nil)
