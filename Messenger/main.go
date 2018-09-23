package main

import (
	"encoding/json"
	"os"
)

type message struct {
	Event   string `json:"event"`
	Address string `json:"address,omitempty"`
	Payload string `json:"payload,omitempty"`
	Error   string `json:"error,omitempty"`
	Debug   string `json:"debug,omitempty"`
}

func write(msg message) {
	msgBytes, error := json.Marshal(msg)
	if error != nil {
		return
	}

	lenBytes := make([]byte, 4)
	nativeEndian.PutUint32(lenBytes, uint32(len(msgBytes)))
	os.Stdout.Write(lenBytes)

	os.Stdout.Write(msgBytes)
}

func read() message {
	buffer := make([]byte, 4)
	length, error := os.Stdin.Read(buffer)
	if error != nil || length != 4 {
		return message{
			Event: "error",
			Error: "cannot read input size",
			Debug: error.Error() + ">" + string(buffer),
		}
	}

	size := nativeEndian.Uint32(buffer)
	buffer = make([]byte, size)
	length, error = os.Stdin.Read(buffer)
	if error != nil && length != int(size) {
		return message{
			Event: "error",
			Error: "cannot read input message",
			Debug: error.Error() + ">" + string(buffer),
		}
	}

	msg := message{}
	error = json.Unmarshal(buffer, &msg)
	if error != nil {
		return message{
			Event: "error",
			Error: "cannot deserialize input message",
			Debug: error.Error() + ">" + string(buffer),
		}
	}

	return msg
}

func main() {
	init := read()
	if init.Event == "error" {
		write(message{
			Event: "close",
			Error: init.Error,
			Debug: init.Debug,
		})
		os.Exit(1)
	}

	switch init.Event {
	case "open-udpPeer":
		udpPeer(init.Address)
		break
	case "open-tcpServer":
		tcpServer(init.Address)
		break
	case "open-tcpClient":
		tcpClient(init.Address)
		break
	default:
		write(message{
			Event: "close",
			Error: "unknown socket type",
		})
		os.Exit(1)
		break
	}
}
