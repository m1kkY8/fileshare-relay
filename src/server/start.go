package server

import (
	"log"
	"net"
)

func (server *Server) Start() {
	listener, err := net.Listen("tcp", "0.0.0.0:9001")
	if err != nil {
		log.Fatalf("err starting server %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("err accepting connection %v", err)
			continue
		}

		go server.handleConnection(conn)
	}
}
