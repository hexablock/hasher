package hasher

import (
	"crypto/sha256"
	"hash"
)

type SHA256Hasher struct{}

func (h *SHA256Hasher) Clone() Hasher {
	return &SHA256Hasher{}
}

func (h *SHA256Hasher) Name() string {
	return "sha256"
}
func (h *SHA256Hasher) New() hash.Hash {
	return sha256.New()
}
func (h *SHA256Hasher) Size() int {
	return sha256.Size
}
