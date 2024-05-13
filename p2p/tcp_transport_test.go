package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTCPTransport(t *testing.T) {

	tr := NewTCPTransport(TCPTransportOpts{
		ListenAddr:      ":4500",
		HandleShakeFunc: nil,
		Decoder:         nil,
	})
	assert.Equal(t, tr.TCPTransportOpts.ListenAddr, tr)

	//Server
	assert.Nil(t, tr.ListenAndAccept())

	select {}
}
