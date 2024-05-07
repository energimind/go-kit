package route

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

const (
	validInJSON  = `{"id":"1","name":"test-in"}`
	validOutJSON = `{"id":"1","name":"test-out"}`
	invalidJSON  = `{"id":`
)

var (
	validIn = testIn{
		ID:   "1",
		Name: "test-in",
	}
	validOut = testOut{
		ID:   "1",
		Name: "test-out",
	}
)

type testIn struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type testOut struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func init() {
	gin.SetMode(gin.TestMode)
}

func testRequest(body string, opts ...testOption) (*gin.Context, *httptest.ResponseRecorder) {
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	uri, _ := url.Parse("http://localhost")

	ctx.Request = &http.Request{
		URL:  uri,
		Body: io.NopCloser(strings.NewReader(body)),
	}

	for _, opt := range opts {
		opt(ctx.Request)
	}

	return ctx, rec
}

type testOption func(*http.Request)

func withHeader(key, value string) testOption {
	return func(req *http.Request) {
		if req.Header == nil {
			req.Header = http.Header{}
		}

		req.Header.Add(key, value)
	}
}

func withQuery(key, value string) testOption {
	return func(req *http.Request) {
		q := req.URL.Query()

		q.Add(key, value)

		req.URL.RawQuery = q.Encode()
	}
}

func requireNoResponseAndStatusSent(t *testing.T, ctx *gin.Context) {
	t.Helper()

	require.Equal(t, http.StatusOK, ctx.Writer.Status())
	require.False(t, ctx.Writer.Written())
}
