package merkle

import "github.com/r2pq-suite/r2pq-sdk/internal/errors"

type Node []byte

type Tree struct {
	Leaves [][]byte
	Root   Node
}

// Build is a placeholder; final uses PQ-friendly hash.
func Build(leaves [][]byte) (*Tree, error) {
	if len(leaves) == 0 {
		return &Tree{}, nil
	}
	return nil, errors.ErrUnimplemented
}
