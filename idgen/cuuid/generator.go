// Package cuuid provides a generator for condensed UUIDs.
// These are the base32 encoded version of the UUIDs.
package cuuid

import (
	"encoding/base32"

	"github.com/google/uuid"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz234567"

// Generator implements an ID generator based on condensed UUIDs.
type Generator struct {
	enc *base32.Encoding
}

// NewGenerator returns a new Generator instance.
func NewGenerator() *Generator {
	return &Generator{
		enc: base32.NewEncoding(alphabet).WithPadding(base32.NoPadding),
	}
}

// GenerateID generates a new ID.
func (g *Generator) GenerateID() string {
	u := uuid.New()

	return g.enc.EncodeToString(u[:])
}
