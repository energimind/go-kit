package route

import (
	"io"

	"github.com/gin-gonic/gin"
)

// handleFn is a signature for a function that handles a request using
// the input and output types.
type handleFn[In, Out any] func(ctx *Context, in In) (Out, error)

// Handle creates a new handler for a request.
// It uses the input and output types.
// Note: Handle will not set an error status code or send any response in case of an error.
// This is to allow for middleware to catch this error and respond accordingly.
func Handle[In, Out any](fn handleFn[In, Out]) gin.HandlerFunc {
	return handleOp[In, Out]{fn: fn}.execute
}

// handleOp is used to handle a request using input and output types.
type handleOp[In, Out any] struct {
	fn handleFn[In, Out]
}

// execute executes the handler.
func (h handleOp[In, Out]) execute(ctx *gin.Context) {
	var in In

	defer io.NopCloser(ctx.Request.Body)

	if err := decodeJSON(ctx.Request.Body, &in); err != nil {
		abortWithError(ctx, err)

		return
	}

	rcx := newContext(ctx)

	out, err := h.fn(rcx, in)
	if err != nil {
		abortWithError(ctx, err)

		return
	}

	ctx.Header(contentType, applicationJSON)

	abortIfError(ctx, encodeJSON(ctx.Writer, out))
}
