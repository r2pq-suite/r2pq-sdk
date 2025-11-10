package pq

type algoID int

const (
	AlgoShimSig algoID = iota + 1
)

type (
	signerCtor   func() (Signer, error)
	verifierCtor func(pub []byte) (Verifier, error)
)

var reg = map[algoID]struct {
	newSigner   signerCtor
	newVerifier verifierCtor
}{}

func Register(id algoID, s signerCtor, v verifierCtor) {
	reg[id] = struct {
		newSigner   signerCtor
		newVerifier verifierCtor
	}{s, v}
}

// Helpers you can use later:
func NewSigner(id algoID) (Signer, bool, error) {
	entry, ok := reg[id]
	if !ok {
		return nil, false, nil
	}
	s, err := entry.newSigner()
	return s, true, err
}

func NewVerifier(id algoID, pub []byte) (Verifier, bool) {
	entry, ok := reg[id]
	if !ok {
		return nil, false
	}
	return entry.newVerifier(pub), true
}
