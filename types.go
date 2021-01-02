package mercadopago

type Identification struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Payer struct {
	Email          string         `json:"email"`
	Identification Identification `json:"identification"`
}

type CardHolder struct {
	Name           string         `json:"name"`
	Identification Identification `json:"identification"`
}
