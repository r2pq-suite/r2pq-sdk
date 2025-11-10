package message

import "github.com/r2pq-suite/r2pq-sdk/version"

type Header struct {
	Version string
	Type    string
}

// Message is the generic wire message.
type Message struct {
	Header  Header
	Payload []byte
}

func New(t string, payload []byte) Message {
	return Message{
		Header:  Header{Version: version.Version, Type: t},
		Payload: payload,
	}
}
