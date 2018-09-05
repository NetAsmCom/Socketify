package main

import (
	"io"
	"net"
	"os"
)

var tcpServerClosed = false
var tcpServerSocket *net.TCPListener
var tcpServerConnections = make(map[string]*net.TCPConn)

func tcpServer(addrStr string) {
	address, error := net.ResolveTCPAddr("tcp", addrStr)
	if error != nil {
		write(message{
			Event: "close",
			Error: "cannot resolve tcp address",
			Debug: error.Error(),
		})
		os.Exit(1)
	}

	tcpServerSocket, error = net.ListenTCP("tcp", address)
	if error != nil {
		write(message{
			Event: "close",
			Error: "cannot listen tcp socket",
			Debug: error.Error(),
		})
		os.Exit(1)
	}

	defer tcpServerSocket.Close()
	go tcpServerAccept()
	write(message{
		Event:   "open",
		Address: tcpServerSocket.Addr().String(),
	})

	for {
		msg := read()
		switch msg.Event {
		case "error":
			if !tcpServerClosed {
				write(message{
					Event: "close",
					Error: msg.Error,
					Debug: msg.Debug,
				})
				tcpServerClosed = true
			}
			os.Exit(1)
			break
		case "send":
			tcpServerSend(msg)
			break
		case "drop":
			tcpServerDrop(msg.Address, "", "")
			break
		case "close":
			if !tcpServerClosed {
				write(message{
					Event: "close",
				})
				tcpServerClosed = true
			}
			tcpServerSocket.Close()
			os.Exit(0)
			break
		}
	}
}

func tcpServerAccept() {
	for {
		connection, error := tcpServerSocket.AcceptTCP()
		if error != nil {
			if !tcpServerClosed {
				write(message{
					Event: "close",
					Error: "cannot accept tcp connection",
					Debug: error.Error(),
				})
				tcpServerClosed = true
			}
			os.Exit(1)
		}

		addrStr := connection.RemoteAddr().String()
		tcpServerConnections[addrStr] = connection
		go tcpServerReceive(connection)
		write(message{
			Event:   "connect",
			Address: addrStr,
		})
	}
}

func tcpServerReceive(conn *net.TCPConn) {
	addrStr := conn.RemoteAddr().String()
	for {
		buffer := make([]byte, 1500)
		length, error := conn.Read(buffer)
		if error != nil {
			if error == io.EOF {
				tcpServerDrop(addrStr, "connection closed", error.Error())
			} else {
				tcpServerDrop(addrStr, "cannot read from tcp socket", error.Error())
			}
			return
		}

		write(message{
			Event:   "receive",
			Address: addrStr,
			Payload: string(buffer[:length]),
		})
	}
}

func tcpServerSend(msg message) {
	conn, ok := tcpServerConnections[msg.Address]
	if !ok {
		return
	}

	_, error := conn.Write([]byte(msg.Payload))
	if error != nil {
		tcpServerDrop(conn.RemoteAddr().String(), "cannot write to tcp socket", error.Error())
	}
}

func tcpServerDrop(addrStr string, error string, debug string) {
	conn, ok := tcpServerConnections[addrStr]
	if !ok {
		return
	}

	write(message{
		Event:   "disconnect",
		Address: addrStr,
		Error:   error,
		Debug:   debug,
	})

	delete(tcpServerConnections, addrStr)
	conn.Close()
}
