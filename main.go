package main

import (
	"fileshare-relay/src/server"
)

/*
starts tcp server on port 9001
server accepts connections from senders and receivers
and just acts as middle-man between two clients
*/

func main() {
	server := server.NewServer()
	server.Start()
}
