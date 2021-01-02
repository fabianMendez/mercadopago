package mercadopago

import (
	"encoding/json"
	"net/http"
)

// Client is the api client
type Client interface {
	// NewCardToken creates a new CardToken
	NewCardToken(params CardTokenParams) (*CardToken, error)

	// GetIdentificationTypes returns all available identification types
	GetIdentificationTypes() (IdentificationTypes, error)

	// GetPaymentMethods returns all available payment methods
	GetPaymentMethods() (PaymentMethods, error)

	// GetPaymentMethodsForBin returns all available payment methods for the given bin
	GetPaymentMethodsForBin(bin string) (PaymentMethods, error)

	// GetInstallments returns all available installments for the given
	// payment method, amount and issuer
	GetInstallments(params GetInstallmentsParams) (Installments, error)

	// GetCardIssuers returns all available issuers for the given payment method
	GetCardIssuers(paymentMethodID string) (Issuers, error)

	// NewPayment creates a new Payment
	NewPayment(params PaymentParams) (*Payment, error)

	// NewTestUser creates a new test user
	NewTestUser(params TestUserParams) (*TestUser, error)
}

// client is the api client implementation
type client struct {
	version     string
	httpClient  *http.Client
	baseURL     string
	accessToken string
	publicKey   string
}

// NewClient creates a new Client
func NewClient(baseURL, publicKey, accessToken string) Client {
	return &client{
		httpClient:  http.DefaultClient,
		baseURL:     baseURL,
		accessToken: accessToken,
		publicKey:   publicKey,
		version:     "1.3.1",
	}
}

func (c *client) requestAndDecode(req *http.Request, response interface{}) error {
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
