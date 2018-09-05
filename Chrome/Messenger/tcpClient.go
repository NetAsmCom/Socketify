package main

import (
	"io"
	"net"
	"os"
)

var tcpClientSocket *net.TCPConn

func tcpClient(addrStr string) {
	address, error := net.ResolveTCPAddr("tcp", addrStr)
	if error != nil {
		write(message{
			Event: "open",
			Error: "cannot resolve tcp address",
			Debug: error.Error(),
		})
		os.Exit(1)
	}

	tcpClientSocket, error = net.DialTCP("tcp", nil, address)
	if error != nil {
		write(message{
			Event: "open",
			Error: "cannot connect tcp socket",
			Debug: error.Error(),
		})
		os.Exit(1)
	}

	defer tcpClientSocket.Close()
	go tcpClientReceive()
	write(message{
		Event:   "open",
		Address: tcpClientSocket.LocalAddr().String(),
	})

	for {
		msg := read()
		switch msg.Event {
		case "error":
			write(message{
				Event: "close",
				Error: msg.Error,
				Debug: msg.Debug,
			})
			os.Exit(1)
			break
		case "send":
			tcpClientSend(msg)
			break
		case "close":
			write(message{
				Event: "close",
			})
			tcpClientSocket.Close()
			os.Exit(0)
			break
		}
	}
}

func tcpClientSend(msg message) {
	_, error := tcpClientSocket.Write([]byte(msg.Payload))
	if error != nil {
		write(message{
			Event: "close",
			Error: "cannot write to tcp socket",
			Debug: error.Error(),
		})
		os.Exit(1)
	}
}

func tcpClientReceive() {
	for {
		buffer := make([]byte, 1500)
		length, error := tcpClientSocket.Read(buffer)
		if error != nil {
			if error == io.EOF {
				write(message{
					Event: "close",
					Error: "connection closed",
					Debug: error.Error(),
				})
				os.Exit(0)
			} else {
				write(message{
					Event: "close",
					Error: "cannot read from tcp socket",
					Debug: error.Error(),
				})
				os.Exit(1)
			}
		}

		write(message{
			Event:   "receive",
			Payload: string(buffer[:length]),
		})
	}
}
