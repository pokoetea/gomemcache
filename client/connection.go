package client

import (
	"fmt"
	"io"
	"net"
	"sync/atomic"

	"github.com/pokoetea/gomemcache/protocol"
)

type Conn struct {
	conn   net.Conn
	opaque uint32
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{
		conn: conn,
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

func (c *Conn) Request(packet *protocol.Packet) (*protocol.Packet, error) {
	id := atomic.AddUint64(&c.opaque, 1)
	b := make([]byte, 0, 4)
	b = append(b, byte(id>>24))
	b = append(b, byte(id>>16))
	b = append(b, byte(id>>8))
	b = append(b, byte(id))
	packet.Header.SetOpaque(b)
	return nil, nil
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
		fmt.Println("read:", packet.ParseBytes())
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
			b := packet.ParseBytes()
			fmt.Println("write:", b)
			if _, err := c.conn.Write(b); err != nil {
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
