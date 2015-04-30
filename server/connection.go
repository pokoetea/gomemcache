package server

import (
	"io"
	"net"
	"sync"

	"github.com/pokoetea/gomemcache/protocol"
)

type Conn struct {
	*sync.RWMutex
	conn       net.Conn
	dispatcher Dispatcher
	writeChan  chan *protocol.Packet
	stopChan   chan struct{}
}

func NewConn(conn net.Conn, dispatcher Dispatcher) *Conn {
	return &Conn{
		RWMutex:    new(sync.RWMutex),
		conn:       conn,
		dispatcher: dispatcher,
		writeChan:  make(chan *protocol.Packet, 1000),
		stopChan:   make(chan struct{}),
	}
}

func (c *Conn) Run() {
	syncChan := make(chan struct{})
	go func() {
		c.writeLoop()
		close(syncChan)
	}()
	c.readLoop()
	<-syncChan
}

func (c *Conn) Stop() {
	c.Lock()
	defer c.Unlock()
	if c.running() {
		close(c.stopChan)
	}
}

func (c *Conn) readLoop() error {
	defer c.Stop()
	for {
		if !c.running() {
			return nil
		}
		headerBuffer := make([]byte, protocol.HeaderSize)
		if _, err := io.ReadFull(c.conn, headerBuffer); err != nil {
			return err
		}
		header := protocol.Header(headerBuffer)
		packet := protocol.NewPacket(header)
		totalLength := int(header.TotalBodyLength())
		if totalLength > 0 {
			total := make([]byte, totalLength)
			if _, err := io.ReadFull(c.conn, total); err != nil {
				return err
			}
			index := 0
			extrasLength := int(header.ExtrasLength())
			if extrasLength > 0 {
				packet.Extras = total[index : index+extrasLength]
				index += extrasLength
			}
			keyLength := int(header.KeyLength())
			if keyLength > 0 {
				packet.Key = total[index : index+keyLength]
				index += keyLength
			}
			valueLength := totalLength - extrasLength - keyLength
			if valueLength > 0 {
				packet.Value = total[index : index+valueLength]
				index += valueLength
			}
		}
		c.dispatcher.Dispatch(packet, c.writeChan)
	}
}

func (c *Conn) writeLoop() error {
	defer c.Stop()
	for {
		if !c.running() {
			return nil
		}
		select {
		case packet := <-c.writeChan:
			if _, err := c.conn.Write(packet.ParseBytes()); err != nil {
				return err
			}
		case <-c.stopChan:
			return nil
		}
	}
}

func (c *Conn) running() bool {
	select {
	case <-c.stopChan:
		return false
	default:
		return true
	}
}
