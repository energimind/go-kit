package route

import (
	"io"

	"github.com/gin-gonic/gin"
)

// handleInFn is a signature for a function that handles a request using
// the input type.
type handleInFn[In any] func(ctx *Context, in In) error

// HandleIn creates a new handler for a request.
// It uses the input type.
// Note: HandleIn will not set an error status code or send any response in case of an error.
// This is to allow for middleware to catch this error and respond accordingly.
func HandleIn[In any](fn handleInFn[In]) gin.HandlerFunc {
	return handleInOp[In]{fn: fn}.execute
}

// handleInOp is used to handle a request using an input type.
type handleInOp[In any] struct {
	fn handleInFn[In]
}

// execute executes the handler.
func (h handleInOp[In]) execute(ctx *gin.Context) {
	var in In

	defer io.NopCloser(ctx.Request.Body)

	if err := decodeJSON(ctx.Request.Body, &in); err != nil {
		abortWithError(ctx, err)

		return
	}

	rcx := newContext(ctx)

	err := h.fn(rcx, in)
	if err != nil {
		abortWithError(ctx, err)

		return
	}
}
