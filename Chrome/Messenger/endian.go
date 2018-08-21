package main

import (
	"encoding/binary"
	"unsafe"
)

func isBigEndian() bool {
	i := 0x1
	m := (*[int(unsafe.Sizeof(0))]byte)(unsafe.Pointer(&i))
	return m[0] == 0
}

func isLittleEndian() bool {
	return !isBigEndian()
}

func getNativeEndian() binary.ByteOrder {
	if isBigEndian() {
		return binary.BigEndian
	}
	return binary.LittleEndian
}
