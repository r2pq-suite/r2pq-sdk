package kem

import sdkerr "github.com/r2pq-suite/r2pq-sdk/internal/errors"

type Scheme string

const SchemeR2PQ Scheme = "r2pq-kem"

type PublicKey []byte
type SecretKey []byte
type Ciphertext []byte
type SharedSecret []byte

type KEM interface {
	Encaps(pk PublicKey) (ct Ciphertext, ss SharedSecret, err error)
	Decaps(sk SecretKey, ct Ciphertext) (SharedSecret, error)
	Scheme() Scheme
}

func New() KEM { return nop{} }

type nop struct{}
func (nop) Encaps(_ PublicKey) (Ciphertext, SharedSecret, error) { return nil, nil, sdkerr.ErrUnimplemented }
func (nop) Decaps(_ SecretKey, _ Ciphertext) (SharedSecret, error) { return nil, sdkerr.ErrUnimplemented }
func (nop) Scheme() Scheme { return SchemeR2PQ }
