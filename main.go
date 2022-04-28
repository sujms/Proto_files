package main

import (
	"proto_files/client"
	"proto_files/server"
)

func main() {

	client.Client()
	client.Clienttest()
	server.Server()
	server.Servertest()

}
