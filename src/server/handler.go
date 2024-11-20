package server

import (
	"encoding/binary"
	"fileshare-relay/src/ack"
	"fileshare-relay/src/handshake"
	"log"
	"net"
)

func (server *Server) handleConnection(conn net.Conn) {
	handshake, err := handshake.ReadHandshake(conn)
	if err != nil {
		log.Printf("error getting handshake %v", err)
	}

	if handshake.Intent == "s" {
		server.Senders[handshake] = conn
		return
	}

	if handshake.Intent == "r" {
		server.Receivers[handshake] = conn

		senderHandshake, senderConn, found := server.getSenderByKeyword(handshake.Keyword)
		if !found {
			log.Println("sender not found")
			return
		}

		err := binary.Write(conn, binary.LittleEndian, senderHandshake.FileSize)
		if err != nil {
			log.Println("error sending filesize")
			return
		}

		err = ack.SendAck(senderConn, "ready")
		if err != nil {
			log.Println("error sending ack")
			return
		}

		server.transfer(conn, senderConn, senderHandshake.FileSize)

		delete(server.Receivers, handshake)
		delete(server.Receivers, senderHandshake)

		defer conn.Close()
		defer senderConn.Close()
	}
}
