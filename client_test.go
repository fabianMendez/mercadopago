package mercadopago

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_requestAndDecode(t *testing.T) {
	tests := []struct {
		name             string
		response         interface{}
		expectedResponse interface{}
		expectedErr      error
		respStatus       int
		respBody         string
	}{
		{
			name:       "internal server error",
			respStatus: http.StatusInternalServerError,
			respBody:   `{"message":"internal_error","error":null,"status":500,"cause":[]}`,
			expectedErr: Error{
				Message: "internal_error",
				Status:  500,
				Cause:   []ErrorCause{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/" {
					w.WriteHeader(http.StatusNotFound)
					return
				}

				w.WriteHeader(tt.respStatus)
				_, _ = w.Write([]byte(tt.respBody))
			}))

			c := NewClient(server.URL, "access-token", "public-key")

			req, err := http.NewRequest(http.MethodGet, server.URL+"/", nil)
			require.NoError(t, err)

			err = c.requestAndDecode(req, tt.response)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
