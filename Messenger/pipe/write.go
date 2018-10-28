package pipe

import (
	"os"
)

// Write TODO
func Write(msgType MsgType, payload []byte) error {
	msg := append([]byte{uint8(msgType)}, payload...)

	len := uint32(len(msg))
	lenBytes := make([]byte, 4)
	nativeEndian.PutUint32(lenBytes, len)
	_, err := os.Stdout.Write(lenBytes)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(msg)
	return err
}
