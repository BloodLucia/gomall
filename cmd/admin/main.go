package main

import (
	"log"

	"github.com/kalougata/gomall/cmd/wire"
	httppkg "github.com/kalougata/gomall/pkg/http"
)

func main() {
	server, cleanup, err := wire.NewApp()

	if err != nil {
		log.Fatalln(err)
	}

	httppkg.Run(server.AdminServerHTTP, ":8001")

	defer cleanup()
}
