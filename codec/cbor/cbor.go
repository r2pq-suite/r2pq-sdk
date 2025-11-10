package cbor

import (
	sdkerr "github.com/r2pq-suite/r2pq-sdk/internal/errors"
	"github.com/r2pq-suite/r2pq-sdk/codec/registry"
)

const Name = "cbor"

// Register installs the (placeholder) CBOR codec.
func Register() { registry.Register(Codec{}) }

type Codec struct{}

func (Codec) Name() string                     { return Name }
func (Codec) Encode(v any) ([]byte, error)     { return nil, sdkerr.ErrUnimplemented }
func (Codec) Decode(_ []byte, _ any) error     { return sdkerr.ErrUnimplemented }
