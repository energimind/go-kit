package route

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Context is the context of a request.
// It wraps the gin.Context.
type Context struct {
	gc *gin.Context
}

// Ensure Context implements the context.Context interface.
var _ context.Context = (*Context)(nil)

// newContext creates a new Context.
func newContext(gc *gin.Context) *Context {
	return &Context{gc: gc}
}

// Status sets the status of the response.
func (c *Context) Status(status int) {
	c.gc.Status(status)
}

// OK sets the status of the response to OK.
func (c *Context) OK() {
	c.Status(http.StatusOK)
}

// Created sets the status of the response to Created.
func (c *Context) Created() {
	c.Status(http.StatusCreated)
}

// Accepted sets the status of the response to Accepted.
func (c *Context) Accepted() {
	c.Status(http.StatusAccepted)
}

// NoContent sets the status of the response to No Content.
func (c *Context) NoContent() {
	c.Status(http.StatusNoContent)
}

// Header returns the value of the header identified by key.
func (c *Context) Header(key string) string {
	return c.gc.GetHeader(key)
}

// SetHeader sets the value of the header identified by key.
func (c *Context) SetHeader(key, value string) {
	c.gc.Header(key, value)
}

// Param returns the value of the parameter identified by key.
func (c *Context) Param(key string) string {
	return c.gc.Param(key)
}

// Query returns the value of the query identified by key.
func (c *Context) Query(key string) string {
	return c.gc.Query(key)
}

// Get returns the value for the given key.
func (c *Context) Get(key string) (any, bool) {
	return c.gc.Get(key)
}

// Deadline implements the context.Context interface.
func (c *Context) Deadline() (time.Time, bool) {
	return c.gc.Deadline()
}

// Done implements the context.Context interface.
func (c *Context) Done() <-chan struct{} {
	return c.gc.Done()
}

// Err implements the context.Context interface.
func (c *Context) Err() error {
	return c.gc.Err() //nolint:wrapcheck // leave as is
}

// Value implements the context.Context interface.
func (c *Context) Value(key any) any {
	return c.gc.Value(key)
}
