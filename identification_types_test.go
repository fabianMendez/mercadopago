package mercadopago_test

import (
	"github.com/fabianMendez/mercadopago"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_GetIdentificationTypes(t *testing.T) {
	tests := []struct {
		name             string
		expectedResponse mercadopago.IdentificationTypes
		expectedErr      error
		respStatus       int
		respBody         string
	}{
		{
			name:       "successful response",
			respStatus: http.StatusOK,
			respBody:   `[{"id":"CC","name":"C.C.","type":"number","min_length":5,"max_length":20},{"id":"CE","name":"C.E.","type":"number","min_length":5,"max_length":20},{"id":"NIT","name":"NIT","type":"number","min_length":5,"max_length":20},{"id":"Otro","name":"Otro","type":"number","min_length":5,"max_length":20}]`,
			expectedResponse: mercadopago.IdentificationTypes{
				{ID: "CC", Name: "C.C.", Type: "number", MinLength: 5, MaxLength: 20},
				{ID: "CE", Name: "C.E.", Type: "number", MinLength: 5, MaxLength: 20},
				{ID: "NIT", Name: "NIT", Type: "number", MinLength: 5, MaxLength: 20},
				{ID: "Otro", Name: "Otro", Type: "number", MinLength: 5, MaxLength: 20},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			publicKey := "public-key"
			accessToken := "access-token"

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "/identification_types", r.URL.Path)

				//assert.Equal(t, publicKey, r.URL.Query().Get("public_key"))
				assert.Equal(t, "Bearer "+accessToken, r.Header.Get("Authorization"))

				w.WriteHeader(tt.respStatus)
				_, _ = w.Write([]byte(tt.respBody))
			}))

			c := mercadopago.NewClient(server.URL, publicKey, accessToken)

			got, err := c.GetIdentificationTypes()
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expectedResponse, got)
		})
	}
}
