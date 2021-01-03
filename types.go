package mercadopago

type Identification struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Payer struct {
	Email          string         `json:"email"`
	Identification Identification `json:"identification"`
	FirstName      string         `json:"first_name,omitempty"`
	LastName       string         `json:"last_name,omitempty"`
	Phone          *Phone         `json:"phone,omitempty"`
	Type           string         `json:"type,omitempty"`
	EntityType     interface{}    `json:"entity_type,omitempty"`
	ID             interface{}    `json:"id,omitempty"`
}

type Cardholder struct {
	Name           string         `json:"name"`
	Identification Identification `json:"identification"`
}

type Phone struct {
	AreaCode  string `json:"area_code"`
	Number    string `json:"number"`
	Extension string `json:"extension"`
}

type Card struct {
	ID              interface{} `json:"id"`
	FirstSixDigits  string      `json:"first_six_digits"`
	LastFourDigits  string      `json:"last_four_digits"`
	ExpirationMonth int         `json:"expiration_month"`
	ExpirationYear  int         `json:"expiration_year"`
	DateCreated     string      `json:"date_created"`
	DateLastUpdated string      `json:"date_last_updated"`
	Cardholder      Cardholder  `json:"cardholder"`
}

type PayerCosts []PayerCost

type PayerCost struct {
	InstallmentRate          int           `json:"installment_rate"`
	DiscountRate             int           `json:"discount_rate"`
	MinAllowedAmount         float64       `json:"min_allowed_amount"`
	Labels                   []interface{} `json:"labels"`
	Installments             int           `json:"installments"`
	ReimbursementRate        interface{}   `json:"reimbursement_rate"`
	MaxAllowedAmount         float64       `json:"max_allowed_amount"`
	PaymentMethodOptionID    string        `json:"payment_method_option_id"`
	InstallmentRateCollector []string      `json:"installment_rate_collector"`
	RecommendedMessage       string        `json:"recommended_message"`
	InstallmentAmount        float64       `json:"installment_amount"`
	TotalAmount              float64       `json:"total_amount"`
}
