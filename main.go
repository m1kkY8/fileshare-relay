package main

import (
	"fileshare-relay/src/server"
	"log"
	"net"
)

func main() {
	server := server.NewServer()
	listener, err := net.Listen("tcp", "0.0.0.0:9001")
	if err != nil {
		log.Printf("err starting server %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("err accepting connection %v", err)
			continue
		}

		var fileSize int64
		go server.HandleConnection(conn, &fileSize)
	}
}
