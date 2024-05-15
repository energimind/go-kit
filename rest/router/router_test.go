package router

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestRouter(t *testing.T) {
	rtr := New(testMiddleware)

	rtr.RegisterRoutes(&testRoute{})
	rtr.RegisterRoutesFunc(func(rtr gin.IRouter) {
		rtr.GET("/demo", func(c *gin.Context) {})
	})

	routes := rtr.GetRoutes()

	require.Len(t, routes, 2)

	require.Equal(t, "/test", routes[0].Path)
	require.Equal(t, "/demo", routes[1].Path)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)

	rtr.ServeHTTP(rec, req)

	require.Equal(t, 200, rec.Code)
	require.JSONEq(t, `{"message": "test-message"}`, rec.Body.String())
	require.Equal(t, "test-header", rec.Header().Get("X-Test-Middleware"))
}
