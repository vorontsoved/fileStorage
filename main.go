package main

import (
	"github.com/vorontsoved/fileStorage/p2p"
	"log"
)

func main() {
	tr := p2p.NewTCPTransport(":45000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
