package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var (
	data       []byte
	dataJSON   []byte
	outerCount = int(10e6)
	innerCount = 8
)

func init() {
	data = generateNestedSlice(outerCount, innerCount)
	dataJSON = generateNestedJSONSlice(outerCount, innerCount)
}

func Test_DecodeJSON(t *testing.T) {
	fmt.Printf("json: data size is %d\n", len(dataJSON))
	var parsed [][]uint64
	beforeDecode := time.Now()
	json.Unmarshal(dataJSON, &parsed)
	fmt.Printf("json: spent %s decoding\n", time.Since(beforeDecode))

	beforeSum := time.Now()
	sum := uint64(0)
	for _, inner := range parsed {
		for _, elem := range inner {
			sum += elem
		}
	}
	fmt.Printf("json: spent %s summing\n", time.Since(beforeSum))
}

func Test_DecodeFirst(t *testing.T) {
	fmt.Printf("binary decode: data size is %d\n", len(data))
	beforeDecode := time.Now()
	parsed := make([][]uint64, outerCount)
	offset := 0
	i := 0
	for offset < len(data) {
		length := binary.LittleEndian.Uint16(data[offset : offset+2])
		offset += 2

		parsed[i] = make([]uint64, length)
		for j := uint16(0); j < length; j++ {
			parsed[i][j] = binary.LittleEndian.Uint64(data[offset : offset+8])
			offset += 8
		}
	}
	fmt.Printf("binary decode: spent %s decoding\n", time.Since(beforeDecode))

	beforeSum := time.Now()
	sum := uint64(0)
	for _, inner := range parsed {
		for _, elem := range inner {
			sum += elem
		}
	}
	fmt.Printf("binary decode: spent %s summing\n", time.Since(beforeSum))
}

func Test_DecodeWhenUsed(t *testing.T) {
	before := time.Now()
	offset := 0
	sum := uint64(0)
	for offset < len(data) {
		length := binary.LittleEndian.Uint16(data[offset : offset+2])
		offset += 2

		for j := uint16(0); j < length; j++ {
			sum += binary.LittleEndian.Uint64(data[offset : offset+8])
			offset += 8
		}
	}
	fmt.Printf("delayed decode: spent %s decoding+summing\n", time.Since(before))
}

func generateNestedSlice(outer, inner int) []byte {
	out := make([]byte, outer*2+outer*inner*8)

	offset := 0

	for i := 0; i < outer; i++ {
		binary.LittleEndian.PutUint16(out[offset:offset+2], uint16(inner))
		offset += 2

		for j := 0; j < inner; j++ {
			binary.LittleEndian.PutUint64(out[offset:offset+8], uint64(rand.Int()))
			offset += 8
		}

	}

	return out
}

func generateNestedJSONSlice(outer, inner int) []byte {
	nested := make([][]uint64, outer)
	for i := range nested {
		nested[i] = make([]uint64, inner)
		for j := range nested[i] {
			nested[i][j] = uint64(rand.Int())
		}
	}

	d, _ := json.Marshal(nested)
	return d
}
