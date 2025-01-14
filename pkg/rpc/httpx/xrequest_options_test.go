package httpx

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleRequestOptions(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "GET", r.Method)
		require.Equal(t, "/testpath", r.URL.Path)
		require.Equal(t, "b", r.URL.Query().Get("a"))
		require.Equal(t, "d", r.Header.Get("c"))
		bodyRaw, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		require.Equal(t, "", string(bodyRaw))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	}))

	client := NewXClient()

	ctx := context.Background()
	resp, err := client.XGet(ctx, s.URL+"/testpath",
		WithQuery("a", "b"),
		WithHeader("c", "d"),
	)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
