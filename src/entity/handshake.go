package entity

import (
	"net"

	"github.com/vmihailenco/msgpack/v5"
)

type Handshake struct {
	Intent   string `msgpack:"intent"`   // r for receive, s for send
	Keyword  string `msgpack:"keyword"`  // keyword used for pairing clients
	FileName string `msgpack:"filename"` // r for receive, s for send
	FileSize int64  `msgpack:"filesize"` // r for receive, s for send
}

func GetHandshake(conn net.Conn) ([]byte, error) {
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func DecodeHandshake(buf []byte) (Handshake, error) {
	var decodedHandshake Handshake

	err := msgpack.Unmarshal(buf, &decodedHandshake)
	if err != nil {
		return Handshake{}, err
	}

	return decodedHandshake, nil
}
