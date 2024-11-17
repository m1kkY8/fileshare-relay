package ack

import (
	"net"

	"github.com/vmihailenco/msgpack/v5"
)

type Acknowledge struct {
	Ready   bool   `msgpack:"ready"` // send message to sender about status
	Message string `msgpack:"msg"`   // test message
}

func SendAck(conn net.Conn, message string) error {
	ack := Acknowledge{
		Ready:   true,
		Message: message,
	}

	ackBytes, err := msgpack.Marshal(ack)
	if err != nil {
		return err
	}

	_, err = conn.Write(ackBytes)
	if err != nil {
		return err
	}

	return nil
}
