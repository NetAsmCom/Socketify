package main

import (
	"encoding/json"
	"flag"
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
	versionPtr := flag.Bool("version", false, "print version")
	installPtr := flag.Bool("install", false, "install host app for browsers")
	chromeExtIDPtr := flag.String("chromeExtId", "aaaaaaaaaaabbbbbbbbbbccccccccccc", "chrome extension id to allow")
	firefoxExtIDPtr := flag.String("firefoxExtId", "extension@socketify.net", "firefox extension id to allow")
	uninstallPtr := flag.Bool("uninstall", false, "uninstall host app from browsers")

	flag.Parse()

	if *versionPtr {
		os.Stdout.Write([]byte("version: 1.0\n"))
		os.Exit(0)
	}

	if *installPtr {
		if install(*chromeExtIDPtr, *firefoxExtIDPtr) {
			os.Stdout.Write([]byte("install: succeeded\n"))
			os.Exit(0)
		} else {
			os.Stdout.Write([]byte("install: failed\n"))
			os.Exit(1)
		}
	}

	if *uninstallPtr {
		if uninstall() {
			os.Stdout.Write([]byte("uninstall: succeeded\n"))
			os.Exit(0)
		} else {
			os.Stdout.Write([]byte("uninstall: failed\n"))
			os.Exit(1)
		}
	}

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
