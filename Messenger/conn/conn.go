package conn

import (
	"crypto/tls"
	"net"
)

var connected = false

var encrypted = false
var keyAES string

var connTLS *tls.Conn
var connTCP *net.TCPConn
var connUDP *net.UDPConn

var chanConn chan error
var chanTCP chan []byte
var chanUDP chan []byte

// Connect TODO
func Connect(host string, secure bool) chan error {
	chanConn = make(chan error)
	chanTCP = make(chan []byte)
	chanUDP = make(chan []byte)

	go connect(host, secure)

	return chanConn
}

func connect(host string, secure bool) {
	encrypted = secure

	if encrypted {
		cfg := tls.Config{InsecureSkipVerify: true}
		conn, err := tls.Dial("tcp", host, &cfg)
		if err != nil {
			chanConn <- err
			return
		}

		connTLS = conn
	} else {
		addr, err := net.ResolveTCPAddr("tcp", host)
		if err != nil {
			chanConn <- err
			return
		}

		conn, err := net.DialTCP("tcp", nil, addr)
		if err != nil {
			chanConn <- err
			return
		}

		connTCP = conn
	}
}
