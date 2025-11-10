package wire

import (
	"github.com/r2pq-suite/r2pq-sdk/codec/registry"
	"github.com/r2pq-suite/r2pq-sdk/protocol/message"
)

// Frame encodes a Message with the named codec.
func Frame(codecName string, m message.Message) ([]byte, error) {
	c, err := registry.Get(codecName)
	if err != nil {
		return nil, err
	}
	return c.Encode(m)
}

// Unframe decodes bytes into a Message with the named codec.
func Unframe(codecName string, b []byte) (message.Message, error) {
	c, err := registry.Get(codecName)
	if err != nil {
		return message.Message{}, err
	}
	var m message.Message
	return m, c.Decode(b, &m)
}
