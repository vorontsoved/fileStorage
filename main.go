package main

import (
	"github.com/vorontsoved/fileStorage/p2p"
	"log"
)

func main() {

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:      ":45000",
		HandleShakeFunc: p2p.NOPHandshakeFunc,
		Decoder:         p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
