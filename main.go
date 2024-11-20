package main

import (
	"fileshare-relay/src/server"
)

func main() {
	server := server.NewServer()
	server.Start()
}
