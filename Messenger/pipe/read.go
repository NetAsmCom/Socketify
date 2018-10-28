package pipe

import (
	"os"
)

// Read TODO
func Read() (MsgType, []byte) {
	buf := make([]byte, 4)
	_, err := os.Stdin.Read(buf)
	if err != nil {
		return InMsgError, nil
	}

	len := nativeEndian.Uint32(buf)
	buf = make([]byte, len)
	_, err = os.Stdin.Read(buf)
	if err != nil {
		return InMsgError, nil
	}

	msgType := MsgType(buf[0])
	if msgType != InMsgConfig &&
		msgType != InMsgSendTCP &&
		msgType != InMsgSendUDP {
		return InMsgError, nil
	}

	return msgType, buf[1:]
}
