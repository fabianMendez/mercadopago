package mercadopago

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PaymentParams struct {
	TransactionAmount   float64     `json:"transaction_amount"`
	PaymentMethodID     string      `json:"payment_method_id"`
	Payer               Payer       `json:"payer"`
	Token               string      `json:"token"`
	Description         string      `json:"description"`
	Installments        int         `json:"installments"`
	NotificationURL     string      `json:"notification_url"`
	SponsorID           interface{} `json:"sponsor_id"`
	BinaryMode          bool        `json:"binary_mode"`
	ExternalReference   string      `json:"external_reference"`
	StatementDescriptor string      `json:"statement_descriptor"`
	AdditionalInfo      interface{} `json:"additional_info"`
}

type Payment struct {
	ID                 int         `json:"id"`
	DateCreated        string      `json:"date_created"`
	DateApproved       string      `json:"date_approved"`
	DateLastUpdated    string      `json:"date_last_updated"`
	DateOfExpiration   interface{} `json:"date_of_expiration"`
	MoneyReleaseDate   string      `json:"money_release_date"`
	OperationType      string      `json:"operation_type"`
	IssuerID           string      `json:"issuer_id"`
	PaymentMethodID    string      `json:"payment_method_id"`
	PaymentTypeID      string      `json:"payment_type_id"`
	Status             string      `json:"status"`
	StatusDetail       string      `json:"status_detail"`
	CurrencyID         string      `json:"currency_id"`
	Description        string      `json:"description"`
	LiveMode           bool        `json:"live_mode"`
	SponsorID          interface{} `json:"sponsor_id"`
	AuthorizationCode  interface{} `json:"authorization_code"`
	MoneyReleaseSchema interface{} `json:"money_release_schema"`
	TaxesAmount        int         `json:"taxes_amount"`
	CounterCurrency    interface{} `json:"counter_currency"`
	ShippingAmount     int         `json:"shipping_amount"`
	PosID              interface{} `json:"pos_id"`
	StoreID            interface{} `json:"store_id"`
	CollectorID        int         `json:"collector_id"`
	Payer              struct {
		FirstName      string `json:"first_name"`
		LastName       string `json:"last_name"`
		Email          string `json:"email"`
		Identification struct {
			Number string `json:"number"`
			Type   string `json:"type"`
		} `json:"identification"`
		Phone struct {
			AreaCode  string `json:"area_code"`
			Number    string `json:"number"`
			Extension string `json:"extension"`
		} `json:"phone"`
		Type       string      `json:"type"`
		EntityType interface{} `json:"entity_type"`
		ID         interface{} `json:"id"`
	} `json:"payer"`
	Metadata struct {
	} `json:"metadata"`
	AdditionalInfo struct {
		Items []struct {
			ID          string  `json:"id"`
			Title       string  `json:"title"`
			Description string  `json:"description"`
			PictureURL  string  `json:"picture_url"`
			CategoryID  string  `json:"category_id"`
			Quantity    int     `json:"quantity"`
			UnitPrice   float64 `json:"unit_price"`
		} `json:"items"`
		Payer struct {
			Phone struct {
				AreaCode string `json:"area_code"`
				Number   string `json:"number"`
			} `json:"phone"`
			Address struct {
				ZipCode      string `json:"zip_code"`
				StreetName   string `json:"street_name"`
				StreetNumber string `json:"street_number"`
			} `json:"address"`
			FirstName        string `json:"first_name"`
			LastName         string `json:"last_name"`
			RegistrationDate string `json:"registration_date"`
		} `json:"payer"`
		Shipments struct {
			ReceiverAddress struct {
				StreetName   string `json:"street_name"`
				StreetNumber int    `json:"street_number"`
				ZipCode      string `json:"zip_code"`
				CityName     string `json:"city_name"`
				StateName    string `json:"state_name"`
			} `json:"receiver_address"`
		} `json:"shipments"`
	} `json:"additional_info"`
	Order struct {
	} `json:"order"`
	ExternalReference         string      `json:"external_reference"`
	TransactionAmount         float64     `json:"transaction_amount"`
	TransactionAmountRefunded int         `json:"transaction_amount_refunded"`
	CouponAmount              int         `json:"coupon_amount"`
	DifferentialPricingID     interface{} `json:"differential_pricing_id"`
	DeductionSchema           interface{} `json:"deduction_schema"`
	TransactionDetails        struct {
		PaymentMethodReferenceID interface{} `json:"payment_method_reference_id"`
		NetReceivedAmount        float64     `json:"net_received_amount"`
		TotalPaidAmount          float64     `json:"total_paid_amount"`
		OverpaidAmount           int         `json:"overpaid_amount"`
		ExternalResourceURL      interface{} `json:"external_resource_url"`
		InstallmentAmount        float64     `json:"installment_amount"`
		FinancialInstitution     interface{} `json:"financial_institution"`
		PayableDeferralPeriod    interface{} `json:"payable_deferral_period"`
		AcquirerReference        interface{} `json:"acquirer_reference"`
	} `json:"transaction_details"`
	FeeDetails []struct {
		Type     string  `json:"type"`
		Amount   float64 `json:"amount"`
		FeePayer string  `json:"fee_payer"`
	} `json:"fee_details"`
	Captured            bool        `json:"captured"`
	BinaryMode          bool        `json:"binary_mode"`
	CallForAuthorizeID  interface{} `json:"call_for_authorize_id"`
	StatementDescriptor string      `json:"statement_descriptor"`
	Installments        int         `json:"installments"`
	Card                struct {
		ID              interface{} `json:"id"`
		FirstSixDigits  string      `json:"first_six_digits"`
		LastFourDigits  string      `json:"last_four_digits"`
		ExpirationMonth int         `json:"expiration_month"`
		ExpirationYear  int         `json:"expiration_year"`
		DateCreated     string      `json:"date_created"`
		DateLastUpdated string      `json:"date_last_updated"`
		Cardholder      struct {
			Name           string `json:"name"`
			Identification struct {
				Number string `json:"number"`
				Type   string `json:"type"`
			} `json:"identification"`
		} `json:"cardholder"`
	} `json:"card"`
	NotificationURL        string        `json:"notification_url"`
	Refunds                []interface{} `json:"refunds"`
	ProcessingMode         string        `json:"processing_mode"`
	MerchantAccountID      interface{}   `json:"merchant_account_id"`
	Acquirer               interface{}   `json:"acquirer"`
	MerchantNumber         interface{}   `json:"merchant_number"`
	AcquirerReconciliation []interface{} `json:"acquirer_reconciliation"`
}

// NewPayment creates a new Payment
func (c *client) NewPayment(params PaymentParams) (*Payment, error) {
	u := fmt.Sprintf("%s/payments", c.baseURL)
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))

	var response Payment
	err = c.requestAndDecode(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
