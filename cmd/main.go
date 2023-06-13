package main

import (
	"Core/internal/server"
)

func main() {
	newServer := server.NewServer()
	newServer.Run()

	//configs := config.LoadConfig()
}
