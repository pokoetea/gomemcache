package client

import (
	"net"

	"github.com/pokoetea/gomemcache/protocol"
)

type Client struct {
	conn *Conn
}

func New() *Client {
	return &Client{}
}

func (c *Client) Run() error {
	c, err := net.Dial("tcp", ":11211")
	if err != nil {
		return err
	}
	c.conn = NewConn(c)
	go c.conn.Run()
	return nil
}

func (c *Client) Get(key []byte) (value []byte, flags []byte, err error) {
	header := Header(make([]byte, 0, protocol.HeaderSize))
	header.SetMagic(protocol.MagicReq)
	header.SetOpCode(protocol.CmdGet)
	header.SetKeyLength(uint16(len(key)))
	header.SetTotalBodyLength(uint32(len(key)))
	packet := protocol.NewPacket(header)
	packet.Key = key
	res, err := conn.Request(packet)
	if err != nil {
		return nil, nil, err
	}
	return res.Value, res.Extras, nil
}

func (c *Client) Set(key []byte, value []byte, flags []byte, expiry uint32) (cas uint64, err error) {
	if len(flags) != 4 {
		return 0, ErrInvalidArgument
	}
	header := Header(make([]byte, 0, protocol.HeaderSize))
	header.SetMagic(protocol.MagicReq)
	header.SetOpCode(protocol.CmdSet)
	header.SetExtrasLength(uint8(8))
	header.SetKeyLength(uint16(len(key)))
	header.SetTotalBodyLength(int32(len(key) + len(value) + 8))
	packet := protocol.NewPacket(header)
	extras := make([]byte, 8)
	extras = append(extras, flags...)
	extras[4] = byte(expiry >> 24)
	extras[5] = byte(expiry >> 16)
	extras[6] = byte(expiry >> 8)
	extras[7] = byte(expiry)
	packet.Extras = extras
	packet.Key = key
	packet.Value = value
	res, err := conn.Request(packet)
	if err != nil {
		return 0, err
	}
	return res.Header.CAS(), nil
}
