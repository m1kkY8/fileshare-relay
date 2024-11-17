package server

import (
	"fileshare-relay/src/ack"
	"fileshare-relay/src/handshake"
	"log"
	"net"
	"time"
)

type Server struct {
	Receivers map[string]net.Conn
	Senders   map[string]net.Conn
}

func NewServer() Server {
	return Server{
		Receivers: make(map[string]net.Conn),
		Senders:   make(map[string]net.Conn),
	}
}

func (server *Server) HandleConnection(conn net.Conn) {
	handshakeBytes, err := handshake.GetHandshake(conn)
	if err != nil {
		log.Printf("error reading handshake %v", err)
		return
	}

	handshake, err := handshake.DecodeHandshake(handshakeBytes)
	if err != nil {
		log.Printf("error decoding handshake %v", err)
	}

	if handshake.Intent == "s" {
		server.Senders[handshake.Keyword] = conn
		// if nema receiver cekaj
	}

	if handshake.Intent == "r" {
		server.Receivers[handshake.Keyword] = conn

		// ako nema sender break i posalji da nema sta da primi
		// if ima sender sa keyword
		// posalji ack da reciever postoji
		// pocni transfer
		v, ok := server.Senders[handshake.Keyword]
		if ok {
			err := ack.SendAck(v, "ready")
			if err != nil {
				log.Println("error sending ack")
				return
			}
		} else {
			log.Println("sender not found")

			conn.Close()
			return
		}

		// nakon 10 sekundi treba da se prekinu
		server.Transfer(server.Senders[handshake.Keyword], server.Receivers[handshake.Keyword], handshake)

		log.Println("receiver connecte")
	}
}

// mock funkcija da simulira neki transfer i prekid konekcije
func (server *Server) Transfer(conn1 net.Conn, conn2 net.Conn, handshake handshake.Handshake) {
	time.Sleep(5 * time.Second)

	defer conn1.Close()
	defer conn2.Close()

	conn1 = nil
	conn2 = nil
}
