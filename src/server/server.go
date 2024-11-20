package server

import (
	"fileshare-relay/src/handshake"
	"net"
	"sync"
)

type Server struct {
	Receivers map[handshake.Handshake]net.Conn
	Senders   map[handshake.Handshake]net.Conn
	mu        sync.Mutex
}

func NewServer() Server {
	return Server{
		Receivers: make(map[handshake.Handshake]net.Conn),
		Senders:   make(map[handshake.Handshake]net.Conn),
	}
}
