package route

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_decodeJSON(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		var in testIn

		err := decodeJSON(strings.NewReader(validInJSON), &in)

		require.NoError(t, err)
		require.Equal(t, validIn, in)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		var in testIn

		err := decodeJSON(strings.NewReader(invalidJSON), &in)

		require.Error(t, err)
		require.True(t, IsBadJSONError(err))
	})
}

func Test_encodeJSON(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		buf := strings.Builder{}

		err := encodeJSON(&buf, validOut)

		require.NoError(t, err)
		require.JSONEq(t, validOutJSON, buf.String())
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		buf := strings.Builder{}

		err := encodeJSON(&buf, func() {})

		require.Error(t, err)
		require.True(t, IsBadJSONError(err))
	})
}
