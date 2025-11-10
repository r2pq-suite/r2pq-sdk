package hash

// Hasher is a minimal interface so we can swap implementations (Keccak, SHAKE, etc.).
type Hasher interface {
	Write(p []byte) (n int, err error)
	Sum(b []byte) []byte
	Reset()
}

// NewKeccak256 returns a placeholder hasher; real impl wired in later.
func NewKeccak256() Hasher { return nop{} }

type nop struct{}
func (nop) Write(p []byte) (int, error) { return len(p), nil }
func (nop) Sum(b []byte) []byte         { return nil }
func (nop) Reset()                      {}
