package server

import (
	"io"
	"log"
	"net"
)

func (server *Server) transfer(conn1 net.Conn, conn2 net.Conn, fileSize int64) {
	_, err := io.CopyN(conn1, conn2, fileSize)
	if err != nil {
		if err == io.EOF {
			log.Println("reached EOF")
			return
		}
		log.Printf("error writing to conn %v", err)
		return
	}
}
