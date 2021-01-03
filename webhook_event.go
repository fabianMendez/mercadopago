package mercadopago

import "time"

type WebhookEvent struct {
	Action      string      `json:"action"`
	APIVersion  string      `json:"api_version"`
	Data        interface{} `json:"data"`
	DateCreated time.Time   `json:"date_created"`
	ID          int64       `json:"id"`
	LiveMode    bool        `json:"live_mode"`
	Type        string      `json:"type"`
	UserID      string      `json:"user_id"`
}
