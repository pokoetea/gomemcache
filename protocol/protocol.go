package protocol

import (
	"reflect"
	"unsafe"
)

const (
	HeaderSize = 24
)

const (
	MagicReq byte = 0x80
	MagicRes byte = 0x81
)

const (
	StatusNoError                   uint16 = 0x0000
	StatusKeyNotFound               uint16 = 0x0001
	StatusKeyExists                 uint16 = 0x0002
	StatusValueTooLarge             uint16 = 0x0003
	StatusInvalidArguments          uint16 = 0x0004
	StatusItemNotStored             uint16 = 0x0005
	StatusIncrDecrOnNonNumericValue uint16 = 0x0006
)

const (
	CmdGet        byte = 0x00
	CmdSet        byte = 0x01
	CmdAdd        byte = 0x02
	CmdReplace    byte = 0x03
	CmdDelete     byte = 0x04
	CmdIncrement  byte = 0x05
	CmdDecrement  byte = 0x06
	CmdQuit       byte = 0x07
	CmdFlush      byte = 0x08
	CmdGetQ       byte = 0x09
	CmdNoOp       byte = 0x0A
	CmdVersion    byte = 0x0B
	CmdGetK       byte = 0x0C
	CmdGetKQ      byte = 0x0D
	CmdAppend     byte = 0x0E
	CmdPrepend    byte = 0x0F
	CmdStat       byte = 0x10
	CmdSetQ       byte = 0x11
	CmdAddQ       byte = 0x12
	CmdReplaceQ   byte = 0x13
	CmdDeleteQ    byte = 0x14
	CmdIncrementQ byte = 0x15
	CmdDecrementQ byte = 0x16
	CmdQuitQ      byte = 0x17
	CmdFlushQ     byte = 0x18
	CmdAppendQ    byte = 0x19
	CmdPrependQ   byte = 0x1A
)

const (
	FirstCmd = CmdGet
	LastCmd  = CmdSet
)

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{sh.Data, sh.Len, 0}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
