package route

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestHandleOut(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		h := HandleOut[testOut](func(ctx *Context) (testOut, error) {
			ctx.Created()

			return validOut, nil
		})

		ctx, rec := testRequest(validInJSON)

		h(ctx)

		require.Equal(t, http.StatusCreated, ctx.Writer.Status())
		require.Equal(t, applicationJSON, ctx.Writer.Header().Get(contentType))
		require.JSONEq(t, validOutJSON, rec.Body.String())
	})

	t.Run("handler-error", func(t *testing.T) {
		h := HandleOut[testOut](func(ctx *Context) (testOut, error) {
			return testOut{}, errors.New("test-error")
		})

		ctx, _ := testRequest(validInJSON)

		h(ctx)

		require.Error(t, ctx.Errors.Last())

		requireNoResponseAndStatusSent(t, ctx)
	})
}
