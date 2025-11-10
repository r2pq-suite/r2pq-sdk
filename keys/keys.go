package keys

import sdkerr "github.com/r2pq-suite/r2pq-sdk/internal/errors"

type PublicKey []byte
type PrivateKey []byte

type KeyPair struct {
	Public  PublicKey
	Private PrivateKey
}

func (k KeyPair) IsZero() bool { return len(k.Public) == 0 && len(k.Private) == 0 }

// Generate is a placeholder for a PQ-safe key generator plugged in later.
func Generate() (KeyPair, error) { return KeyPair{}, sdkerr.ErrUnimplemented }
