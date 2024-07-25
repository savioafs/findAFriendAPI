package main

import (
	"github.com/savioafs/findAFriendAPI/server"
)

func main() {
	server := server.NewServer()
	server.Run()
}
