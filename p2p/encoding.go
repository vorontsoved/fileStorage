package p2p

import "io"

type Decoder interface {
	Decode(reader io.Reader, any2 any) error
}
