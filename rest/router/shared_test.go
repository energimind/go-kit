package router

import "github.com/gin-gonic/gin"

type testRoute struct{}

// Ensure that the testRoute implements the route interface.
var _ route = &testRoute{}

// RegisterRoutes registers the routes to the router.
func (r *testRoute) RegisterRoutes(root gin.IRouter) {
	root.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "test-message"})
	})
}

func testMiddleware(c *gin.Context) {
	c.Header("X-Test-Middleware", "test-header")
	c.Next()
}
