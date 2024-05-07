package route

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func Test_abortWithError(t *testing.T) {
	t.Parallel()

	ctx := &gin.Context{}

	abortWithError(ctx, errors.New("test error"))

	require.True(t, ctx.IsAborted())
	require.Error(t, ctx.Errors.Last())
}

func Test_abortIfError(t *testing.T) {
	t.Parallel()

	t.Run("no-error", func(t *testing.T) {
		ctx := &gin.Context{}

		abortIfError(ctx, nil)

		require.False(t, ctx.IsAborted())
	})

	t.Run("error", func(t *testing.T) {
		ctx := &gin.Context{}

		abortIfError(ctx, errors.New("test error"))

		require.True(t, ctx.IsAborted())
	})
}
