// Package uuid provides a generator for IDs based on UUIDs.
// Generated IDs are based on UUIDs (github.com/google/uuid).
package uuid

import (
	"github.com/google/uuid"
)

// Generator implements an ID generator based on UUIDs.
type Generator struct{}

// NewGenerator returns a new Generator instance.
func NewGenerator() *Generator {
	return &Generator{}
}

// GenerateID generates a new ID.
func (g *Generator) GenerateID() string {
	return uuid.New().String()
}
