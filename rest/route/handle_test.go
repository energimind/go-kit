package route

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		h := Handle[testIn, testOut](func(ctx *Context, in testIn) (testOut, error) {
			ctx.Created()

			return validOut, nil
		})

		ctx, rec := testRequest(validInJSON)

		h(ctx)

		require.Equal(t, http.StatusCreated, ctx.Writer.Status())
		require.Equal(t, applicationJSON, ctx.Writer.Header().Get(contentType))
		require.JSONEq(t, validOutJSON, rec.Body.String())
	})

	t.Run("invalidBody-error", func(t *testing.T) {
		h := Handle[testIn, testOut](func(ctx *Context, in testIn) (testOut, error) {
			return testOut{}, nil
		})

		ctx, _ := testRequest(invalidJSON)

		h(ctx)

		require.ErrorAs(t, ctx.Errors.Last(), &BadJSONError{})

		requireNoResponseAndStatusSent(t, ctx)
	})

	t.Run("handler-error", func(t *testing.T) {
		h := Handle[testIn, testOut](func(ctx *Context, in testIn) (testOut, error) {
			return testOut{}, errors.New("test-error")
		})

		ctx, _ := testRequest(validInJSON)

		h(ctx)

		require.Error(t, ctx.Errors.Last())

		requireNoResponseAndStatusSent(t, ctx)
	})
}
