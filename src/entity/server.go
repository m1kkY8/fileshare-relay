package entity

import (
	"io"
	"log"
	"net"
)

type Server struct {
	Connections map[string][]net.Conn // Store connections based on keyword
	Receiver    net.Conn
	Sender      net.Conn
}

type ConnectionInfo struct {
	Conn      net.Conn
	Handshake Handshake
}

func NewServer() Server {
	return Server{
		Connections: make(map[string][]net.Conn),
	}
}

func (server *Server) HandleConnection(conn net.Conn) {
	handshakeBytes, err := GetHandshake(conn)
	if err != nil {
		log.Printf("error reading handshake %v", err)
		return
	}

	decodedHandshake, err := DecodeHandshake(handshakeBytes)
	if err != nil {
		log.Printf("error decoding handshake %v", err)
	}

	if decodedHandshake.Intent == "s" {
		log.Println("sender connecte")
		server.Sender = conn
	}

	if decodedHandshake.Intent == "r" {
		log.Println("receiver connecte")
		server.Receiver = conn
	}

	if server.Sender != nil && server.Receiver != nil {
		go io.Copy(server.Receiver, server.Sender)

		// server.Sender.Close()
		// server.Receiver.Close()

		server.Sender = nil
		server.Receiver = nil
	}
}
