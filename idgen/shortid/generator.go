// Package shortid provides a generator for short IDs.
// Generated IDs are based on short IDs (github.com/teris-io/shortid).
package shortid

import (
	"crypto/rand"

	"github.com/teris-io/shortid"
)

// Generator implements an ID generator based on short IDs.
type Generator struct {
	rnd func(b []byte) (n int, err error)
	gen *shortid.Shortid
}

// NewGenerator returns a new Generator instance.
func NewGenerator() *Generator {
	g := Generator{
		rnd: rand.Read,
	}

	g.gen = shortid.MustNew(1, shortid.DefaultABC, random(g.rnd))

	return &g
}

// GenerateID generates a new ID.
func (g *Generator) GenerateID() string {
	return g.gen.MustGenerate()
}
