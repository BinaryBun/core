package main

import (
	"Core/config"
	"Core/internal/server"
	"github.com/go-resty/resty/v2"
	"os"
)

func main() {
	var url string
	args := os.Args
	if len(args) <= 1 {
		url = ""
	} else {
		url = args[1]
	}

	// init conf
	configs := config.LoadConfig()

	// init resty
	app := resty.New()

	newServer := server.NewServer(url, configs, app)
	newServer.Run()
}
