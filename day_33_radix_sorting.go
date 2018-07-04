package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	list := []int32{91, 28, 73, 46, 55, 64, 37, 82, 19}
	fmt.Println("unsorted:", list)
	radixsort(list)
	fmt.Println("sorted:  ", list)
}

func radixsort(list []int32) {
	const wordLen = 4
	const highBit = -1 << 31

	buf := bytes.NewBuffer(nil)
	ds := make([][]byte, len(list))
	for i, x := range list {
		binary.Write(buf, binary.LittleEndian, x^highBit)
		b := make([]byte, wordLen)
		buf.Read(b)
		ds[i] = b
	}
	bins := make([][][]byte, 256)
	for i := 0; i < wordLen; i++ {
		for _, b := range ds {
			bins[b[i]] = append(bins[b[i]], b)
		}
		j := 0
		for k, bs := range bins {
			copy(ds[j:], bs)
			j += len(bs)
			bins[k] = bs[:0]
		}
	}
	var w int32
	for i, b := range ds {
		buf.Write(b)
		binary.Read(buf, binary.LittleEndian, &w)
		list[i] = w ^ highBit
	}
}
