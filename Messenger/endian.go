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

var nativeEndian = map[bool]binary.ByteOrder{
	true:  binary.BigEndian,
	false: binary.LittleEndian,
}[isBigEndian()]
