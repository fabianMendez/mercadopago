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
		{
			name:       "invalid parameter",
			respStatus: http.StatusBadRequest,
			respBody:   `{"message":"Invalid Value for Field: cardholder.name","status":400,"error":"bad_request","cause":[{"description":"Invalid parameter 'cardholder.name'","code":"316"}]}`,
			expectedErr: Error{
				Message:   "Invalid Value for Field: cardholder.name",
				Status:    400,
				ErrorCode: "bad_request",
				Cause: []ErrorCause{
					{Code: "316", Description: `Invalid parameter 'cardholder.name'`},
				},
			},
		},
		{
			name:       "notificaction_url attribute must be url valid",
			respStatus: http.StatusBadRequest,
			respBody:   `{"message":"notificaction_url attribute must be url valid","error":"bad_request","status":400,"cause":[{"code":4020,"description":"notificaction_url attribute must be url valid","data":null}]}`,
			expectedErr: Error{
				Message:   "notificaction_url attribute must be url valid",
				Status:    400,
				ErrorCode: "bad_request",
				Cause: []ErrorCause{
					{Code: 4020, Description: `notificaction_url attribute must be url valid`},
				},
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

			c := NewClient(server.URL, "access-token", "public-key").(*client)

			req, err := http.NewRequest(http.MethodGet, server.URL+"/", nil)
			require.NoError(t, err)

			err = c.requestAndDecode(req, tt.response)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
