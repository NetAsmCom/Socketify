package main

import (
	"encoding/json"
	"os"

	"./pipe"
)

// Config TODO
type Config struct {
	Host   string
	Secure bool
}

func main() {
	msgType, payload := pipe.Read()
	if msgType != pipe.InMsgConfig {
		pipe.Write(pipe.OutMsgDisconnect, []byte("cannot read config"))
		os.Exit(1)
		return
	}

	cfg := Config{}
	err := json.Unmarshal(payload, &cfg)
	if err != nil {
		pipe.Write(pipe.OutMsgDisconnect, []byte("cannot parse config"))
		os.Exit(1)
		return
	}

	for {
		msgType, payload = pipe.Read()
		switch msgType {
		case pipe.InMsgSendTCP:
		case pipe.InMsgSendUDP:
		default:
			pipe.Write(pipe.OutMsgDisconnect, []byte("unexpected message type"))
			os.Exit(1)
			return
		}
	}
}
