package pipe

import (
	"encoding/binary"
	"unsafe"
)

var nativeEndian = map[bool]binary.ByteOrder{
	true:  binary.BigEndian,
	false: binary.LittleEndian,
}[func() bool {
	i := 0x1
	m := (*[int(unsafe.Sizeof(0))]byte)(unsafe.Pointer(&i))
	return m[0] == 0
}()]
