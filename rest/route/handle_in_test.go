package route

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestHandleIn(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		h := HandleIn[testIn](func(ctx *Context, in testIn) error {
			ctx.Created()

			return nil
		})

		ctx, _ := testRequest(validInJSON)

		h(ctx)

		require.Equal(t, http.StatusCreated, ctx.Writer.Status())
		require.Empty(t, ctx.Writer.Header().Get(contentType))
	})

	t.Run("invalidBody-error", func(t *testing.T) {
		h := HandleIn[testIn](func(ctx *Context, in testIn) error {
			return nil
		})

		ctx, _ := testRequest(invalidJSON)

		h(ctx)

		require.ErrorAs(t, ctx.Errors.Last(), &BadJSONError{})

		requireNoResponseAndStatusSent(t, ctx)
	})

	t.Run("handler-error", func(t *testing.T) {
		h := HandleIn[testIn](func(ctx *Context, in testIn) error {
			return errors.New("test-error")
		})

		ctx, _ := testRequest(validInJSON)

		h(ctx)

		require.Error(t, ctx.Errors.Last())

		requireNoResponseAndStatusSent(t, ctx)
	})
}
