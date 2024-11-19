package server

import (
	"encoding/binary"
	"fileshare-relay/src/ack"
	"fileshare-relay/src/handshake"
	"io"
	"log"
	"net"
)

type Mock struct {
	Receivers map[handshake.Handshake]net.Conn
	Senders   map[handshake.Handshake]net.Conn
}
type Server struct {
	Receivers map[handshake.Handshake]net.Conn
	Senders   map[handshake.Handshake]net.Conn
}

func NewServer() Server {
	return Server{
		Receivers: make(map[handshake.Handshake]net.Conn),
		Senders:   make(map[handshake.Handshake]net.Conn),
	}
}

func (server *Server) HandleConnection(conn net.Conn, fileSize *int64) {
	var amogus int64
	handshake, err := handshake.ReadHandshake(conn)
	if err != nil {
		log.Printf("error getting handshake %v", err)
	}

	if handshake.Intent == "s" {
		server.Senders[handshake] = conn
	}

	if handshake.Intent == "r" {
		server.Receivers[handshake] = conn

		// ako nema sender break i posalji da nema sta da primi
		// if ima sender sa keyword
		// posalji ack da reciever postoji
		// pocni transfer

		for hs, c := range server.Senders {
			// match connections
			if hs.Keyword == handshake.Keyword {

				// posalji filesize
				log.Println(hs.FileSize)
				binary.Write(conn, binary.LittleEndian, hs.FileSize)
				amogus = hs.FileSize

				// posalji ack
				err := ack.SendAck(c, "ready")
				if err != nil {
					log.Println("error sending ack")
					return
				}

				server.Transfer(conn, c, amogus)

				delete(server.Receivers, handshake)
				delete(server.Receivers, hs)
			}
		}

		// nakon 10 sekundi treba da se prekinu
	}
}

// mock funkcija da simulira neki transfer i prekid konekcije
func (server *Server) Transfer(conn1 net.Conn, conn2 net.Conn, fileSize int64) {
	_, err := io.CopyN(conn1, conn2, fileSize)
	if err != nil {
		log.Printf("error writing to conn %v", err)
		return
	}

	defer conn1.Close()
	defer conn2.Close()
}
