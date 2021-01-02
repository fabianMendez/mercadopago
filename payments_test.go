package mercadopago_test

import (
	"github.com/fabianMendez/mercadopago"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_NewPayment(t *testing.T) {
	tests := []struct {
		name             string
		params           mercadopago.PaymentParams
		expectedBody     string
		expectedResponse *mercadopago.Payment
		expectedErr      error
		respStatus       int
		respBody         string
	}{
		{
			name: "successful response",
			params: mercadopago.PaymentParams{
				TransactionAmount: 1234.56,
				PaymentMethodID:   "visa",
				Payer: mercadopago.Payer{
					Email: "testuser@test.com",
					Identification: mercadopago.Identification{
						Type:   "CC",
						Number: "1234567890",
					},
				},
				Token:               "card-token",
				Description:         "payment test",
				Installments:        1,
				NotificationURL:     "http://localhost:1234/callback",
				BinaryMode:          false,
				ExternalReference:   "MP001",
				StatementDescriptor: "MercadoPago",
				AdditionalInfo: map[string]interface{}{
					"items": []map[string]interface{}{
						{"id": "123", "title": "Test Product 1", "quantity": 1, "unit_price": 1234.5},
					},
				},
			},
			expectedBody: `{"transaction_amount":1234.56,"payment_method_id":"visa","payer":{"email":"testuser@test.com","identification":{"type":"CC","number":"1234567890"},"phone":{"area_code":"","number":"","extension":""}},"token":"card-token","description":"payment test","installments":1,"notification_url":"http://localhost:1234/callback","sponsor_id":null,"binary_mode":false,"external_reference":"MP001","statement_descriptor":"MercadoPago","additional_info":{"items":[{"id":"123","quantity":1,"title":"Test Product 1","unit_price":1234.5}]}}`,
			respStatus:   http.StatusOK,
			respBody:     `{"id":1232478005,"date_created":"2021-01-01T22:42:18.000-04:00","date_approved":null,"date_last_updated":"2021-01-01T22:42:18.000-04:00","date_of_expiration":null,"money_release_date":null,"operation_type":"regular_payment","issuer_id":"202","payment_method_id":"amex","payment_type_id":"credit_card","status":"rejected","status_detail":"cc_rejected_other_reason","currency_id":"COP","description":"Pago de prueba","live_mode":false,"sponsor_id":null,"authorization_code":null,"money_release_schema":null,"taxes_amount":0,"counter_currency":null,"brand_id":null,"shipping_amount":0,"pos_id":null,"store_id":null,"integrator_id":null,"platform_id":null,"corporation_id":null,"collector_id":695864235,"payer":{"first_name":"Test","last_name":"Test","email":"test_user_80507629@testuser.com","identification":{"number":"32659430","type":"DNI"},"phone":{"area_code":"01","number":"1111-1111","extension":""},"type":"registered","entity_type":null,"id":"695869370"},"marketplace_owner":null,"metadata":{},"additional_info":{"available_balance":null,"nsu_processadora":null},"order":{},"external_reference":null,"transaction_amount":2000,"net_amount":1680.67,"taxes":[{"value":319.33,"type":"IVA"}],"transaction_amount_refunded":0,"coupon_amount":0,"differential_pricing_id":null,"deduction_schema":null,"installments":1,"transaction_details":{"payment_method_reference_id":null,"net_received_amount":0,"total_paid_amount":2000,"overpaid_amount":0,"external_resource_url":null,"installment_amount":2000,"financial_institution":null,"payable_deferral_period":null,"acquirer_reference":null},"fee_details":[],"charges_details":[],"captured":true,"binary_mode":false,"call_for_authorize_id":null,"statement_descriptor":"MERCADOPAGO","card":{"id":null,"first_six_digits":"374378","last_four_digits":"5283","expiration_month":11,"expiration_year":2025,"date_created":"2021-01-01T22:42:18.000-04:00","date_last_updated":"2021-01-01T22:42:18.000-04:00","cardholder":{"name":"OTHE","identification":{"number":"23090923","type":"CC"}}},"notification_url":null,"refunds":[],"processing_mode":"aggregator","merchant_account_id":null,"merchant_number":null,"acquirer_reconciliation":[]}`,
			expectedResponse: &mercadopago.Payment{
				ID:               1232478005,
				DateCreated:      "2021-01-01T22:42:18.000-04:00",
				DateApproved:     "",
				DateLastUpdated:  "2021-01-01T22:42:18.000-04:00",
				MoneyReleaseDate: "",
				OperationType:    "regular_payment", IssuerID: "202",
				PaymentMethodID: "amex", PaymentTypeID: "credit_card",
				Status: "rejected", StatusDetail: "cc_rejected_other_reason",
				CurrencyID: "COP", Description: "Pago de prueba",
				LiveMode: false, TaxesAmount: 0, ShippingAmount: 0,
				CollectorID: 695864235,
				Payer: mercadopago.Payer{
					FirstName: "Test", LastName: "Test", Email: "test_user_80507629@testuser.com",
					Identification: mercadopago.Identification{Number: "32659430", Type: "DNI"},
					Phone:          mercadopago.Phone{AreaCode: "01", Number: "1111-1111", Extension: ""},
					Type:           "registered",
					ID:             "695869370",
				},
				Metadata: struct{}{},
				AdditionalInfo: map[string]interface{}{
					"available_balance": nil, "nsu_processadora": nil,
				}, Order: struct{}{}, ExternalReference: "", TransactionAmount: 2000, TransactionAmountRefunded: 0, CouponAmount: 0,
				TransactionDetails: mercadopago.TransactionDetails{
					NetReceivedAmount: 0, TotalPaidAmount: 2000,
					OverpaidAmount: 0, InstallmentAmount: 2000,
				},
				FeeDetails: mercadopago.FeeDetails{},
				Captured:   true, BinaryMode: false,
				StatementDescriptor: "MERCADOPAGO", Installments: 1,
				Card: mercadopago.Card{
					FirstSixDigits: "374378", LastFourDigits: "5283",
					ExpirationMonth: 11, ExpirationYear: 2025,
					DateCreated:     "2021-01-01T22:42:18.000-04:00",
					DateLastUpdated: "2021-01-01T22:42:18.000-04:00",
					Cardholder: mercadopago.Cardholder{
						Name:           "OTHE",
						Identification: mercadopago.Identification{Number: "23090923", Type: "CC"},
					},
				},
				NotificationURL:        "",
				Refunds:                []interface{}{},
				ProcessingMode:         "aggregator",
				AcquirerReconciliation: []interface{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accessToken := "access-token"

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, "/payments", r.URL.Path)

				assert.Equal(t, "Bearer "+accessToken, r.Header.Get("Authorization"))

				body, err := ioutil.ReadAll(r.Body)
				require.NoError(t, err)

				assert.Equal(t, tt.expectedBody, string(body))

				w.WriteHeader(tt.respStatus)
				_, _ = w.Write([]byte(tt.respBody))
			}))

			c := mercadopago.NewClient(server.URL, "", accessToken)

			got, err := c.NewPayment(tt.params)

			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expectedResponse, got)
		})
	}
}
