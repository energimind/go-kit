package route

import "github.com/gin-gonic/gin"

// handleOutFn is a signature for a function that handles a request using
// the output type.
type handleOutFn[Out any] func(ctx *Context) (Out, error)

// HandleOut creates a new handler for a request.
// It uses the output type.
// Note: HandleOut will not set an error status code or send any response in case of an error.
// This is to allow for middleware to catch this error and respond accordingly.
func HandleOut[Out any](fn handleOutFn[Out]) gin.HandlerFunc {
	return handleOutOp[Out]{fn: fn}.execute
}

// handleOutOp is used to handle a request using an output type.
type handleOutOp[Out any] struct {
	fn handleOutFn[Out]
}

// execute executes the handler.
func (h handleOutOp[Out]) execute(ctx *gin.Context) {
	rcx := newContext(ctx)

	out, err := h.fn(rcx)
	if err != nil {
		abortWithError(ctx, err)

		return
	}

	ctx.Header(contentType, applicationJSON)

	abortIfError(ctx, encodeJSON(ctx.Writer, out))
}
