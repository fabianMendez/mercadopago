package mercadopago

import (
	"fmt"
	"net/http"
)

type PaymentMethods []PaymentMethod

type PaymentMethod struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	PaymentTypeID   string `json:"payment_type_id"`
	Status          string `json:"status"`
	SecureThumbnail string `json:"secure_thumbnail"`
	Thumbnail       string `json:"thumbnail"`
	DeferredCapture string `json:"deferred_capture"`
	Settings        []struct {
		CardNumber struct {
			Validation string `json:"validation"`
			Length     int    `json:"length"`
		} `json:"card_number"`
		Bin struct {
			Pattern             string `json:"pattern"`
			InstallmentsPattern string `json:"installments_pattern"`
			ExclusionPattern    string `json:"exclusion_pattern"`
		} `json:"bin"`
		SecurityCode struct {
			Length       int    `json:"length"`
			CardLocation string `json:"card_location"`
			Mode         string `json:"mode"`
		} `json:"security_code"`
	} `json:"settings"`
	AdditionalInfoNeeded  []string `json:"additional_info_needed"`
	MinAllowedAmount      int      `json:"min_allowed_amount"`
	MaxAllowedAmount      int      `json:"max_allowed_amount"`
	AccreditationTime     int      `json:"accreditation_time"`
	FinancialInstitutions []struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"financial_institutions"`
	ProcessingModes []string `json:"processing_modes"`
	// from the search response only
	PayerCosts []struct {
		InstallmentRate       int           `json:"installment_rate"`
		DiscountRate          int           `json:"discount_rate"`
		MinAllowedAmount      float64       `json:"min_allowed_amount"`
		Labels                []interface{} `json:"labels"`
		Installments          int           `json:"installments"`
		ReimbursementRate     interface{}   `json:"reimbursement_rate"`
		MaxAllowedAmount      float64       `json:"max_allowed_amount"`
		PaymentMethodOptionID string        `json:"payment_method_option_id"`
	} `json:"payer_costs"`
	Issuer struct {
		Default bool   `json:"default"`
		Name    string `json:"name"`
		ID      int    `json:"id"`
	} `json:"issuer"`
	TotalFinancialCost   interface{}   `json:"total_financial_cost"`
	MinAccreditationDays int           `json:"min_accreditation_days"`
	MaxAccreditationDays int           `json:"max_accreditation_days"`
	MerchantAccountID    interface{}   `json:"merchant_account_id"`
	Bins                 []interface{} `json:"bins"`
	Marketplace          string        `json:"marketplace"`
	Agreements           []interface{} `json:"agreements"`
	Labels               []string      `json:"labels"`
	FinancingDeals       struct {
		Legals         interface{} `json:"legals"`
		Installments   interface{} `json:"installments"`
		ExpirationDate interface{} `json:"expiration_date"`
		StartDate      interface{} `json:"start_date"`
		Status         string      `json:"status"`
	} `json:"financing_deals"`
	SiteID         string `json:"site_id"`
	ProcessingMode string `json:"processing_mode"`
}

// GetPaymentMethods returns all available payment methods
func (c *Client) GetPaymentMethods() (PaymentMethods, error) {
	u := fmt.Sprintf("%s/payment_methods", c.baseURL)
	if c.publicKey != "" {
		u += fmt.Sprintf("?public_key=%s", c.publicKey)
	}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	if c.accessToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	}

	var methods PaymentMethods
	err = c.requestAndDecode(req, &methods)
	if err != nil {
		return nil, err
	}

	return methods, nil
}

type paymentMethodsSearchResponse struct {
	Paging struct {
		Total  int `json:"total"`
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	} `json:"paging"`
	Results PaymentMethods `json:"results"`
}

// GetPaymentMethodsForBin returns all available payment methods for the given bin
func (c *Client) GetPaymentMethodsForBin(bin string) (PaymentMethods, error) {
	u := fmt.Sprintf("%s/payment_methods/search?public_key=%s&marketplace=NONE&status=active&js_version=%s",
		c.baseURL, c.publicKey, c.version)

	if bin != "" {
		u += fmt.Sprintf("&bins=%s", bin)
	}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var results paymentMethodsSearchResponse
	err = c.requestAndDecode(req, &results)
	if err != nil {
		return nil, err
	}

	return results.Results, nil
}

type Installments []Installment

type Installment struct {
	PaymentMethodID string `json:"payment_method_id"`
	PaymentTypeID   string `json:"payment_type_id"`
	Issuer          struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		SecureThumbnail string `json:"secure_thumbnail"`
		Thumbnail       string `json:"thumbnail"`
	} `json:"issuer"`
	ProcessingMode    string      `json:"processing_mode"`
	MerchantAccountID interface{} `json:"merchant_account_id"`
	PayerCosts        []struct {
		Installments             int           `json:"installments"`
		InstallmentRate          int           `json:"installment_rate"`
		DiscountRate             int           `json:"discount_rate"`
		ReimbursementRate        interface{}   `json:"reimbursement_rate"`
		Labels                   []interface{} `json:"labels"`
		InstallmentRateCollector []string      `json:"installment_rate_collector"`
		MinAllowedAmount         int           `json:"min_allowed_amount"`
		MaxAllowedAmount         int           `json:"max_allowed_amount"`
		RecommendedMessage       string        `json:"recommended_message"`
		InstallmentAmount        float64       `json:"installment_amount"`
		TotalAmount              int           `json:"total_amount"`
		PaymentMethodOptionID    string        `json:"payment_method_option_id"`
	} `json:"payer_costs"`
	Agreements interface{} `json:"agreements"`
}

type GetInstallmentsParams struct {
	PaymentMethodID string  `json:"payment_method_id"`
	Amount          float64 `json:"amount"`
	IssuerID        string  `json:"issuer_id"`
}

// GetInstallments returns all available installments for the given
// payment method, amount and issuer
func (c *Client) GetInstallments(params GetInstallmentsParams) (Installments, error) {
	u := fmt.Sprintf("%s/payment_methods/installments?public_key=%s&js_version=%s",
		c.baseURL, c.publicKey, c.version)

	if params.PaymentMethodID != "" {
		u += fmt.Sprintf("&payment_method_id=%s", params.PaymentMethodID)
	}

	if params.Amount != 0 {
		u += fmt.Sprintf("&amount=%f", params.Amount)
	}

	if params.IssuerID != "" {
		u += fmt.Sprintf("&issuer.id=%s", params.IssuerID)
	}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var installments Installments
	err = c.requestAndDecode(req, &installments)
	if err != nil {
		return nil, err
	}

	return installments, nil
}

type Issuers []Issuer

type Issuer struct {
	ID                string      `json:"id"`
	Name              string      `json:"name"`
	SecureThumbnail   string      `json:"secure_thumbnail"`
	Thumbnail         string      `json:"thumbnail"`
	ProcessingMode    string      `json:"processing_mode"`
	MerchantAccountID interface{} `json:"merchant_account_id"`
}

// GetIssuers returns all available issuers for the given payment method
func (c *Client) GetIssuers(paymentMethodID string) (Issuers, error) {
	u := fmt.Sprintf("%s/payment_methods/card_issuers?public_key=%s&js_version=%s&payment_method_id=%s",
		c.baseURL, c.publicKey, c.version, paymentMethodID)

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var issuers Issuers
	err = c.requestAndDecode(req, &issuers)
	if err != nil {
		return nil, err
	}

	return issuers, nil
}
