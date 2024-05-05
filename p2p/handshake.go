package p2p

type HandshakeFunc func(peer Peer) error

func NOPHandshakeFunc(Peer Peer) error {
	return nil
}
