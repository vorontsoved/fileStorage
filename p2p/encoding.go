package p2p

import (
	"encoding/gob"
	"fmt"
	"io"
)

type Decoder interface {
	Decode(reader io.Reader, msg *Message) error
}

type GOPDecoder struct {
}

func (dec GOPDecoder) Decode(r io.Reader, v any) error {
	return gob.NewDecoder(r).Decode(v)
}

type DefaultDecoder struct {
}

func (dec DefaultDecoder) Decode(r io.Reader, msg *Message) error {
	buf := make([]byte, 1028)

	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	msg.Payload = buf[:n]
	fmt.Println(string(buf[:n]))
	return nil
}
