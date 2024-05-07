package route

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestHandleVoid(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		h := HandleVoid(func(ctx *Context) error {
			ctx.Created()

			return nil
		})

		ctx, _ := testRequest("")

		h(ctx)

		require.Equal(t, http.StatusCreated, ctx.Writer.Status())
		require.Empty(t, ctx.Writer.Header().Get(contentType))
	})

	t.Run("handler-error", func(t *testing.T) {
		h := HandleVoid(func(ctx *Context) error {
			return errors.New("test-error")
		})

		ctx, _ := testRequest("")

		h(ctx)

		require.Error(t, ctx.Errors.Last())

		requireNoResponseAndStatusSent(t, ctx)
	})
}
