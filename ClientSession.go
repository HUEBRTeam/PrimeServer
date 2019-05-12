package PrimeServer

import (
	"net"
	"sync"
)

type ClientSession struct {
	sync.Mutex
	Conn    net.Conn
	Running bool
}

func MakeClientSession(conn net.Conn) *ClientSession {
	return &ClientSession{
		Conn:    conn,
		Running: false,
	}
}
