package route

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestErrors(t *testing.T) {
	t.Parallel()

	tester := func(err, exp error) {
		t.Helper()

		require.IsType(t, exp, err)
		require.Equal(t, "test:42", err.Error())
	}

	tester(NewBadJSONError("test:%d", 42), BadJSONError{})
}

func TestErrorIs(t *testing.T) {
	t.Parallel()

	tester := func(f func(err error) bool, exp error) {
		require.True(t, f(exp))
		require.False(t, f(errors.New("other-error")))
	}

	tester(IsBadJSONError, BadJSONError{})
}
