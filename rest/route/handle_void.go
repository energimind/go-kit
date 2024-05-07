package route

import "github.com/gin-gonic/gin"

// handleVoidFn is a signature for a function that handles a request without any input
// or output types.
type handleVoidFn func(ctx *Context) error

// HandleVoid creates a new handler for a request.
// It does not use any input or output types.
// Note: HandleVoid will not set an error status code or send any response in case of an error.
// This is to allow for middleware to catch this error and respond accordingly.
func HandleVoid(fn handleVoidFn) gin.HandlerFunc {
	return handleVoidOp{fn: fn}.execute
}

// handleVoidOp is used to handle a request without any input or output types.
type handleVoidOp struct {
	fn handleVoidFn
}

// execute executes the handler.
func (h handleVoidOp) execute(ctx *gin.Context) {
	rcx := newContext(ctx)

	err := h.fn(rcx)
	if err != nil {
		abortWithError(ctx, err)

		return
	}
}
