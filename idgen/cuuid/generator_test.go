package cuuid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerator_Generate(t *testing.T) {
	g := NewGenerator()

	generated := make(map[string]struct{})

	for range 200 {
		id := g.GenerateID()

		require.Len(t, id, 26)
		require.NotContains(t, id, "-")

		_, found := generated[id]
		if found {
			require.Fail(t, "duplicated ID", id)
		}

		generated[id] = struct{}{}
	}
}
