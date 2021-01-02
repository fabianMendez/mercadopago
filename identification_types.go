package mercadopago

import (
	"fmt"
	"net/http"
)

type IdentificationType struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	MinLength int    `json:"min_length"`
	MaxLength int    `json:"max_length"`
}

type IdentificationTypes []IdentificationType

// GetIdentificationTypes returns all available identification types
func (c *Client) GetIdentificationTypes() (IdentificationTypes, error) {
	u := fmt.Sprintf("%s/identification_types", c.baseURL)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))

	var response IdentificationTypes
	err = c.requestAndDecode(req, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
