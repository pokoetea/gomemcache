package protocol

type Packet struct {
	Header Header
	Extras []byte
	Key    []byte
	Value  []byte
}

func NewPacket(header Header) *Packet {
	return &Packet{
		Header: header,
	}
}

func (p *Packet) Len() int {
	return len(p.Header) + len(p.Extras) + len(p.Key) + len(p.Value)
}

func (p *Packet) ParseBytes() []byte {
	b := make([]byte, 0, p.Len())
	b = append(b, p.Header...)
	b = append(b, p.Extras...)
	b = append(b, p.Key...)
	b = append(b, p.Value...)
	return b
}
