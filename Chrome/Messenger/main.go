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

	os.Stdout.Write(append(lenBytes, msgBytes...))
}

func read() message {
	buffer := make([]byte, 2048)
	length, error := os.Stdin.Read(buffer)
	if error != nil {
		return message{
			Event: "error",
			Error: "cannot read input",
			Debug: error.Error(),
		}
	}

	msg := message{}
	error = json.Unmarshal(buffer[4:length], &msg)
	if error != nil {
		return message{
			Event: "error",
			Error: "cannot deserialize input",
			Debug: error.Error(),
		}
	}

	return msg
}

func main() {
	init := read()
	if init.Event == "error" {
		write(message{
			Event: "open",
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
		// TODO
		break
	case "open-tcpClient":
		// TODO
		break
	default:
		write(message{
			Event: "open",
			Error: "unknown socket type",
		})
		os.Exit(1)
		break
	}
}
