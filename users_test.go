package mercadopago_test

import (
	"github.com/fabianMendez/mercadopago"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_NewTestUser(t *testing.T) {
	tests := []struct {
		name             string
		params           mercadopago.TestUserParams
		expectedBody     string
		expectedResponse *mercadopago.TestUser
		expectedErr      error
		respStatus       int
		respBody         string
	}{
		{
			name: "successful response",
			params: mercadopago.TestUserParams{
				SiteID: "MCO",
			},
			expectedBody: `{"site_id":"MCO"}`,
			respStatus:   http.StatusOK,
			respBody:     `{"id": 123456, "nickname": "TT123456", "password": "qatest123456", "site_status": "active", "email": "test_user_123456@testuser.com"}`,
			expectedResponse: &mercadopago.TestUser{
				ID:         123456,
				Nickname:   "TT123456",
				Password:   "qatest123456",
				SiteStatus: "active",
				Email:      "test_user_123456@testuser.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accessToken := "access-token"

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, "/users/test_user", r.URL.Path)

				assert.Equal(t, "Bearer "+accessToken, r.Header.Get("Authorization"))

				body, err := ioutil.ReadAll(r.Body)
				require.NoError(t, err)

				assert.Equal(t, tt.expectedBody, string(body))

				w.WriteHeader(tt.respStatus)
				_, _ = w.Write([]byte(tt.respBody))
			}))

			c := mercadopago.NewClient(server.URL, "", accessToken)

			got, err := c.NewTestUser(tt.params)

			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expectedResponse, got)
		})
	}
}
