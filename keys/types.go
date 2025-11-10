package keys

type Algorithm string

const (
    AlgoShimSig Algorithm = "shim-sig-v1" // placeholder; swap to SPHINCS+ later
)

type PublicKey []byte
type PrivateKey []byte
type Signature []byte
type Address string

type Keypair struct {
    Algo Algorithm
    Pub  PublicKey
    Pri  PrivateKey
}

type Signer interface {
    Algorithm() Algorithm
    Public() PublicKey
    Address() Address
    Sign(msg []byte) (Signature, error)
}

type Verifier interface {
    Algorithm() Algorithm
    Verify(pub PublicKey, msg []byte, sig Signature) bool
    AddressFrom(pub PublicKey) Address
}
