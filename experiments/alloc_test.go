package main

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestBinaryRead(t *testing.T) {
	count := int(1e9)
	buf := make([]byte, 8*count)
	r := bytes.NewReader(buf)

	for i := 0; i < count; i++ {
		var num uint64
		binary.Read(r, binary.LittleEndian, &num)
	}
}

func TestBinaryAllocFree(t *testing.T) {
	count := int(1e9)
	buf := make([]byte, 8*count)

	for i := 0; i < count; i++ {
		num := binary.LittleEndian.Uint64(buf[i*8 : i*8+8])
		_ = num
	}
}
