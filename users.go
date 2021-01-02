package mercadopago

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TestUser struct {
	ID         int    `json:"id"`
	Nickname   string `json:"nickname"`
	Password   string `json:"password"`
	SiteStatus string `json:"site_status"`
	Email      string `json:"email"`
}

type TestUserParams struct {
	SiteID string `json:"site_id"`
}

// NewTestUser creates a new test user
func (c *Client) NewTestUser(params TestUserParams) (*TestUser, error) {
	u := fmt.Sprintf("%s/users/test_user", c.baseURL)

	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))

	var testUser TestUser
	err = c.requestAndDecode(req, &testUser)
	if err != nil {
		return nil, err
	}

	return &testUser, nil
}
