package registry

import (
	"sync"

	sdkerr "github.com/r2pq-suite/r2pq-sdk/internal/errors"
)

type Codec interface {
	Name() string
	Encode(v any) ([]byte, error)
	Decode(b []byte, v any) error
}

var (
	mu     sync.RWMutex
	codecs = map[string]Codec{}
)

func Register(c Codec) {
	mu.Lock()
	defer mu.Unlock()
	codecs[c.Name()] = c
}

func Get(name string) (Codec, error) {
	mu.RLock()
	defer mu.RUnlock()
	c, ok := codecs[name]
	if !ok {
		return nil, sdkerr.ErrCodecNotFound
	}
	return c, nil
}
