package shortid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerator_GenerateID(t *testing.T) {
	g := NewGenerator()

	generated := make(map[string]struct{})

	for range 200 {
		id := g.GenerateID()

		_, found := generated[id]
		if found {
			require.Fail(t, "duplicated ID", id)
		}

		generated[id] = struct{}{}
	}
}
