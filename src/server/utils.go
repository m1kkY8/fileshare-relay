package server

import (
	"fileshare-relay/src/handshake"
	"net"
)

func (server *Server) getSenderByKeyword(keyword string) (handshake.Handshake, net.Conn, bool) {
	for h, c := range server.Senders {
		if h.Keyword == keyword {
			return h, c, true
		}
	}
	return handshake.Handshake{}, nil, false
}
