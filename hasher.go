package hasher

import (
	"fmt"
	"hash"
)

// Hasher is a hash interface implemening a hash function
type Hasher interface {
	Clone() Hasher
	Name() string
	New() hash.Hash
	Size() int
}

// Default returns the default hasher
func Default() Hasher {
	return &SHA256Hasher{}
}

// New returns a new hasher based on the given algorithm
func New(algo string) (Hasher, error) {
	switch algo {
	case "sha256":
		return Default(), nil

	}

	return nil, fmt.Errorf("hash algorithm not supported: %s", algo)
}
