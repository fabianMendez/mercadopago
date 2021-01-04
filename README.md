# Mercadopago
Idiomatic go client for MercadoPago


## :eyes: Example

Create a new payment using a test user

```golang
package main

import (
	"fmt"
	"github.com/fabianMendez/mercadopago"
)

func main() {
	client := mercadopago.NewClient("https://api.mercadopago.com/v1", "TEST-PUBLIC-KEY", "TEST-ACCESS-TOKEN")

	buyer, err := client.NewTestUser(mercadopago.TestUserParams{SiteID: "MCO"})
	if err != nil {
		panic(err)
	}

	identification := mercadopago.Identification{Type: "CC", Number: "19119119100"}

	cardToken, err := client.NewCardToken(mercadopago.CardTokenParams{
		ExpirationMonth: 11,
		ExpirationYear:  2025,
		Cardholder:      mercadopago.Cardholder{Name: "APRO", Identification: identification},
		SecurityCode:    "123",
		CardNumber:      "4013540682746260",
	})
	if err != nil {
		panic(err)
	}

	payment, err := client.NewPayment(mercadopago.PaymentParams{
		PaymentMethodID:   "visa",
		TransactionAmount: 1234.5,
		Payer: mercadopago.Payer{
			Email:          buyer.Email,
			Identification: identification,
		},
		Token:        cardToken.ID,
		Description:  "Test Payment",
		Installments: 1,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(payment.Status)
}

```
