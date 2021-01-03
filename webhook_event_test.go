package mercadopago

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJsonUnmarshal(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    *WebhookEvent
		expectedErr error
	}{
		{
			name:  "valid json",
			input: `{"action":"payment.created","api_version":"v1","data":{"id":"1232485150"},"date_created":"2021-01-02T22:37:13Z","id":6872856233,"live_mode":false,"type":"payment","user_id":"695864235"}`,
			expected: &WebhookEvent{
				Action:     "payment.created",
				APIVersion: "v1",
				Data: map[string]interface{}{
					"id": "1232485150",
				},
				DateCreated: time.Date(2021, 1, 2, 22, 37, 13, 0, time.UTC),
				ID:          6872856233,
				LiveMode:    false,
				Type:        "payment",
				UserID:      "695864235",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual WebhookEvent
			err := json.Unmarshal([]byte(tt.input), &actual)

			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expected, &actual)
		})
	}
}
