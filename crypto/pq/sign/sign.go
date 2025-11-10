package sign

import sdkerr "github.com/r2pq-suite/r2pq-sdk/internal/errors"

type Scheme string

const (
	SchemeR2PQ Scheme = "r2pq-sign"
)

type PrivateKey []byte
type PublicKey  []byte
type Signature  []byte

type Signer interface {
	Sign(priv PrivateKey, msg []byte) (Signature, error)
	Verify(pub PublicKey, msg []byte, sig Signature) bool
	Scheme() Scheme
}

// New returns a placeholder signer; real scheme injected later.
func New() Signer { return nop{} }

type nop struct{}
func (nop) Sign(_ PrivateKey, _ []byte) (Signature, error) { return nil, sdkerr.ErrUnimplemented }
func (nop) Verify(_ PublicKey, _ []byte, _ Signature) bool { return false }
func (nop) Scheme() Scheme                                 { return SchemeR2PQ }
