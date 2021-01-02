package mercadopago

import (
	"encoding/json"
	"net/http"
)

// Client is the api client
type Client struct {
	version     string
	httpClient  *http.Client
	baseURL     string
	accessToken string
	publicKey   string
}

// NewClient creates a new Client
func NewClient(baseURL, publicKey, accessToken string) *Client {
	return &Client{
		httpClient:  http.DefaultClient,
		baseURL:     baseURL,
		accessToken: accessToken,
		publicKey:   publicKey,
		version:     "1.3.1",
	}
}

func (c *Client) requestAndDecode(req *http.Request, response interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		err = json.NewDecoder(resp.Body).Decode(response)

		return err
	}

	var errResp Error
	err = json.NewDecoder(resp.Body).Decode(&errResp)
	if err != nil {
		return err
	}

	return errResp
}
