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
	Phone          Phone          `json:"phone,omitempty"`
	Type           string         `json:"type,omitempty"`
	EntityType     interface{}    `json:"entity_type,omitempty"`
	ID             interface{}    `json:"id,omitempty"`
	// ,"first_name":"","last_name":"","phone":{"area_code":"","number":"","extension":""},"type":"","entity_type":null,"id":null}
	// ,"phone":{"area_code":"","number":"","extension":""}}
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
