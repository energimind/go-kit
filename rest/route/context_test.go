package route

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContext_Status(t *testing.T) {
	t.Parallel()

	tester := func(expCode int, fn func(ctx *Context)) {
		t.Run(strconv.FormatInt(int64(expCode), 10), func(t *testing.T) {
			t.Helper()

			ctx, rec := testRequest("")

			fn(newContext(ctx))

			ctx.Writer.Flush()

			require.Equal(t, expCode, rec.Code)
		})
	}

	tester(http.StatusOK, func(ctx *Context) { ctx.OK() })
	tester(http.StatusCreated, func(ctx *Context) { ctx.Created() })
	tester(http.StatusAccepted, func(ctx *Context) { ctx.Accepted() })
	tester(http.StatusNoContent, func(ctx *Context) { ctx.NoContent() })
}

func TestContext_Header(t *testing.T) {
	t.Parallel()

	ctx, _ := testRequest("", withHeader("header1", "value1"))
	c := newContext(ctx)

	fmt.Println("!", ctx.Request.Header)

	require.Equal(t, "value1", c.Header("header1"))
}

func TestContext_SetHeader(t *testing.T) {
	t.Parallel()

	ctx, rec := testRequest("")
	c := newContext(ctx)

	c.SetHeader("header1", "value1")

	require.Equal(t, "value1", rec.Header().Get("header1"))
}

func TestContext_Query(t *testing.T) {
	t.Parallel()

	ctx, _ := testRequest("", withQuery("query1", "value1"))
	c := newContext(ctx)

	require.Equal(t, "value1", c.Query("query1"))
}

func TestContext_Deadline_Done(t *testing.T) {
	t.Parallel()

	ctx, _ := testRequest("")
	c := newContext(ctx)

	dl, exp := c.Deadline()

	require.Zero(t, dl)
	require.False(t, exp)

	require.Nil(t, c.Done())
}

func TestContext_Err(t *testing.T) {
	t.Parallel()

	ctx, _ := testRequest("")
	c := newContext(ctx)

	require.Nil(t, c.Err())
}

func TestContext_Value(t *testing.T) {
	t.Parallel()

	ctx, _ := testRequest("")
	c := newContext(ctx)

	require.Nil(t, c.Value("key"))

	c.gc.Set("key", "value")

	require.Equal(t, "value", c.Value("key"))
}
