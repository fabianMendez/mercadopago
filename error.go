package mercadopago

// Error represents an error returned by the api
type Error struct {
	Message   string       `json:"message"`
	ErrorCode string       `json:"error"`
	Status    int          `json:"status"`
	Cause     []ErrorCause `json:"cause"`
}

type ErrorCause struct {
	Code        int         `json:"code"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

// Error returns the string representation of the error
func (e Error) Error() string {
	if e.ErrorCode != "" {
		return e.ErrorCode + " - " + e.Message
	}
	return e.Message
}
