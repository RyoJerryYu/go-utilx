package httpx

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type ErrorTransport struct{}

var ErrTransport = errors.New("error from transport")

func (e *ErrorTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, ErrTransport
}

func TestUnwrapTransportError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	client := NewXClientFromHttp(&http.Client{
		Transport: &ErrorTransport{},
	}, WithUnwrapTransportError())

	resp, err := client.Get(context.Background(), server.URL)
	require.Error(t, err)
	require.Nil(t, resp)

	require.Equal(t, ErrTransport, err)
}

func TestReturnErrorIfNot2xx(t *testing.T) {
	testCases := []struct {
		name          string
		ReturnStatus  int
		ExpectedError bool
	}{
		{
			name:          "200",
			ReturnStatus:  http.StatusOK,
			ExpectedError: false,
		},
		{
			name:          "400",
			ReturnStatus:  http.StatusBadRequest,
			ExpectedError: true,
		},
		{
			name:          "500",
			ReturnStatus:  http.StatusInternalServerError,
			ExpectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.ReturnStatus)
			}))

			client := NewXClient(
				WithoutDefaultOption(),
				WithReturnErrorIfNot2xx(),
			)

			resp, err := client.Get(context.Background(), server.URL)
			if tc.ExpectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
			}
		})
	}
}
