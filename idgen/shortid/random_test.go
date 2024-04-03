package shortid

import (
	"crypto/rand"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_random(t *testing.T) {
	id := random(rand.Read)

	require.NotEqual(t, 0, id)
}

func Test_random_panics_on_error(t *testing.T) {
	require.Panics(t, func() {
		random(func(b []byte) (n int, err error) {
			return 0, errors.New("forced-error")
		})
	})
}
