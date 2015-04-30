package protocol

type Header []byte

func (h Header) Magic() byte {
	return h[0]
}

func (h Header) SetMagic(magic byte) {
	h[0] = magic
}

func (h Header) OpCode() byte {
	return h[1]
}

func (h Header) SetOpCode(opCode byte) {
	h[1] = opCode
}

func (h Header) KeyLength() uint16 {
	return uint16(h[3]) | uint16(h[2])<<8
}

func (h Header) SetKeyLength(keyLength uint16) {
	h[2] = byte(keyLength >> 8)
	h[3] = byte(keyLength)
}

func (h Header) ExtrasLength() uint8 {
	return uint8(h[4])
}

func (h Header) SetExtrasLength(extrasLength uint8) {
	h[4] = byte(extrasLength)
}

func (h Header) DataType() uint8 {
	return uint8(h[5])
}

func (h Header) SetDataType(dataType uint8) {
	h[5] = byte(dataType)
}

func (h Header) Reserved() []byte {
	return h[6:8]
}

func (h Header) SetReserved(reserved []byte) {
	h[6] = reserved[0]
	h[7] = reserved[1]
}

func (h Header) Status() uint16 {
	return uint16(h[7]) | uint16(h[6])<<8
}

func (h Header) SetStatus(status uint16) {
	h[6] = byte(status >> 8)
	h[7] = byte(status)
}

func (h Header) TotalBodyLength() uint32 {
	return uint32(h[11]) | uint32(h[10])<<8 | uint32(h[9])<<16 | uint32(h[8])<<24
}

func (h Header) SetTotalBodyLength(totalBodyLength uint32) {
	h[8] = byte(totalBodyLength >> 24)
	h[9] = byte(totalBodyLength >> 16)
	h[10] = byte(totalBodyLength >> 8)
	h[11] = byte(totalBodyLength)
}

func (h Header) Opaque() []byte {
	return h[12:16]
}

func (h Header) SetOpaque(opaque []byte) {
	h[12] = opaque[0]
	h[13] = opaque[1]
	h[14] = opaque[2]
	h[15] = opaque[3]
}

func (h Header) CAS() uint64 {
	return uint64(h[23]) | uint64(h[22])<<8 | uint64(h[21])<<16 | uint64(h[20])<<24 |
		uint64(h[19])<<32 | uint64(h[18])<<40 | uint64(h[17])<<48 | uint64(h[16])<<56
}

func (h Header) SetCAS(cas uint64) {
	h[16] = byte(cas >> 56)
	h[17] = byte(cas >> 48)
	h[18] = byte(cas >> 40)
	h[19] = byte(cas >> 32)
	h[20] = byte(cas >> 24)
	h[21] = byte(cas >> 16)
	h[22] = byte(cas >> 8)
	h[23] = byte(cas)
}
