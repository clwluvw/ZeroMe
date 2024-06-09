package promquerier

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundTrip(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		headers map[string]string
	}{
		{
			name:    "nil headers",
			headers: nil,
		},
		{
			name:    "empty headers",
			headers: map[string]string{},
		},
		{
			name: "non-empty headers",
			headers: map[string]string{
				"Authorization": "Bearer token",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			transport := &headerTransport{
				headers: tc.headers,
			}

			req, err := http.NewRequest("GET", "http://example.com", nil)
			require.NoError(t, err)

			_, err = transport.RoundTrip(req)
			require.NoError(t, err)

			for k, v := range tc.headers {
				require.Equal(t, v, req.Header.Get(k))
			}
		})
	}
}
