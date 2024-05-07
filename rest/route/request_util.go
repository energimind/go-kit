package route

import (
	"github.com/gin-gonic/gin"
)

const (
	contentType     = "Content-Type"
	applicationJSON = "application/json; charset=utf-8"
)

// abortWithError aborts the request with an error.
// This is a helper function to avoid repeating the same code.
func abortWithError(ctx *gin.Context, err error) {
	_ = ctx.Error(err)

	ctx.Abort()
}

// abortIfError aborts the request if there is an error.
// This is a helper function to avoid repeating the same code.
func abortIfError(ctx *gin.Context, err error) {
	if err == nil {
		return
	}

	abortWithError(ctx, err)
}
