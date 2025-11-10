package pq

type Signer interface {
	PublicKey() []byte
	Sign(msg []byte) ([]byte, error)
}

type Verifier interface {
	Verify(msg, sig []byte) bool
}
