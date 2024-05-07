package router

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRouter(t *testing.T) {
	rtr := New(testMiddleware)

	rtr.RegisterRoutes(&testRoute{})

	routes := rtr.GetRoutes()

	require.Len(t, routes, 1)

	require.Equal(t, "/test", routes[0].Path)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test", nil)

	rtr.ServeHTTP(rec, req)

	require.Equal(t, 200, rec.Code)
	require.JSONEq(t, `{"message": "test-message"}`, rec.Body.String())
	require.Equal(t, "test-header", rec.Header().Get("X-Test-Middleware"))
}
