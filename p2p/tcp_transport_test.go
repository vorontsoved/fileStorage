package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPTransport(listenAddr)
	assert.Equal(t, tr.listenAddress, listenAddr)

	//Server
	assert.Nil(t, tr.ListenAndAccept())

	select {}
}
