package main

import (
	"log"

	"github.com/IveGotNorto/jam/server"
)

func main() {
	log.SetFlags(log.Lshortfile)
	server, err := server.NewServer()

	if err != nil {
		log.Fatal(err)
	}

	server.Start()
}
