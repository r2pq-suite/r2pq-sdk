package verifier

import (
	"sync"

	sdkerr "github.com/r2pq-suite/r2pq-sdk/internal/errors"
)

type Verifier interface {
	// Verify returns nil on success, or an error on failure.
	Verify(payload []byte) error
	Name() string
}

var (
	mu   sync.RWMutex
	reg  = map[string]Verifier{}
)

func Register(v Verifier) {
	mu.Lock(); defer mu.Unlock()
	reg[v.Name()] = v
}

func Get(name string) (Verifier, error) {
	mu.RLock(); defer mu.RUnlock()
	v, ok := reg[name]
	if !ok { return nil, sdkerr.ErrUnimplemented }
	return v, nil
}
