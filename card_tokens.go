package mercadopago

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// CardTokenParams represent the params used to create or update a CardToken
type CardTokenParams struct {
	ExpirationMonth int        `json:"expiration_month"`
	ExpirationYear  int        `json:"expiration_year"`
	Cardholder      CardHolder `json:"cardholder"`
	SecurityCode    string     `json:"security_code"`
	CardNumber      string     `json:"card_number"`
}

// CardToken contains the information for a token which represents a card
type CardToken struct {
	ID              string     `json:"id"`
	PublicKey       string     `json:"public_key"`
	Cardholder      CardHolder `json:"cardholder"`
	Status          string     `json:"status"`
	DateCreated     string     `json:"date_created"`
	DateLastUpdated string     `json:"date_last_updated"`
	DateDue         string     `json:"date_due"`
	LuhnValidation  bool       `json:"luhn_validation"`
	LiveMode        bool       `json:"live_mode"`
	RequireEsc      bool       `json:"require_esc"`
}

// NewCardToken creates a new CardToken
func (c *client) NewCardToken(params CardTokenParams) (*CardToken, error) {
	u := fmt.Sprintf("%s/card_tokens?public_key=%s&js_version=%s", c.baseURL, c.publicKey, c.version)

	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var response CardToken
	err = c.requestAndDecode(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
