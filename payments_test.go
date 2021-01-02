package mercadopago_test

import (
	"github.com/fabianMendez/mercadopago/v1"
	"github.com/stretchr/testify/assert"
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
			name:       "successful response",
			params:     mercadopago.PaymentParams{},
			respStatus: http.StatusOK,
			respBody:   `{"id":1232478005,"date_created":"2021-01-01T22:42:18.000-04:00","date_approved":null,"date_last_updated":"2021-01-01T22:42:18.000-04:00","date_of_expiration":null,"money_release_date":null,"operation_type":"regular_payment","issuer_id":"202","payment_method_id":"amex","payment_type_id":"credit_card","status":"rejected","status_detail":"cc_rejected_other_reason","currency_id":"COP","description":"Pago de prueba","live_mode":false,"sponsor_id":null,"authorization_code":null,"money_release_schema":null,"taxes_amount":0,"counter_currency":null,"brand_id":null,"shipping_amount":0,"pos_id":null,"store_id":null,"integrator_id":null,"platform_id":null,"corporation_id":null,"collector_id":695864235,"payer":{"first_name":"Test","last_name":"Test","email":"test_user_80507629@testuser.com","identification":{"number":"32659430","type":"DNI"},"phone":{"area_code":"01","number":"1111-1111","extension":""},"type":"registered","entity_type":null,"id":"695869370"},"marketplace_owner":null,"metadata":{},"additional_info":{"available_balance":null,"nsu_processadora":null},"order":{},"external_reference":null,"transaction_amount":2000,"net_amount":1680.67,"taxes":[{"value":319.33,"type":"IVA"}],"transaction_amount_refunded":0,"coupon_amount":0,"differential_pricing_id":null,"deduction_schema":null,"installments":1,"transaction_details":{"payment_method_reference_id":null,"net_received_amount":0,"total_paid_amount":2000,"overpaid_amount":0,"external_resource_url":null,"installment_amount":2000,"financial_institution":null,"payable_deferral_period":null,"acquirer_reference":null},"fee_details":[],"charges_details":[],"captured":true,"binary_mode":false,"call_for_authorize_id":null,"statement_descriptor":"MERCADOPAGO","card":{"id":null,"first_six_digits":"374378","last_four_digits":"5283","expiration_month":11,"expiration_year":2025,"date_created":"2021-01-01T22:42:18.000-04:00","date_last_updated":"2021-01-01T22:42:18.000-04:00","cardholder":{"name":"OTHE","identification":{"number":"23090923","type":"CC"}}},"notification_url":null,"refunds":[],"processing_mode":"aggregator","merchant_account_id":null,"merchant_number":null,"acquirer_reconciliation":[]}`,
			expectedResponse: &mercadopago.Payment{ID: 1232478005, DateCreated: "2021-01-01T22:42:18.000-04:00", DateApproved: "", DateLastUpdated: "2021-01-01T22:42:18.000-04:00", DateOfExpiration: interface{}(nil), MoneyReleaseDate: "", OperationType: "regular_payment", IssuerID: "202", PaymentMethodID: "amex", PaymentTypeID: "credit_card", Status: "rejected", StatusDetail: "cc_rejected_other_reason", CurrencyID: "COP", Description: "Pago de prueba", LiveMode: false, SponsorID: interface{}(nil), AuthorizationCode: interface{}(nil), MoneyReleaseSchema: interface{}(nil), TaxesAmount: 0, CounterCurrency: interface{}(nil), ShippingAmount: 0, PosID: interface{}(nil), StoreID: interface{}(nil), CollectorID: 695864235, Payer: struct {
				FirstName      string "json:\"first_name\""
				LastName       string "json:\"last_name\""
				Email          string "json:\"email\""
				Identification struct {
					Number string "json:\"number\""
					Type   string "json:\"type\""
				} "json:\"identification\""
				Phone struct {
					AreaCode  string "json:\"area_code\""
					Number    string "json:\"number\""
					Extension string "json:\"extension\""
				} "json:\"phone\""
				Type       string      "json:\"type\""
				EntityType interface{} "json:\"entity_type\""
				ID         interface{} "json:\"id\""
			}{FirstName: "Test", LastName: "Test", Email: "test_user_80507629@testuser.com", Identification: struct {
				Number string "json:\"number\""
				Type   string "json:\"type\""
			}{Number: "32659430", Type: "DNI"}, Phone: struct {
				AreaCode  string "json:\"area_code\""
				Number    string "json:\"number\""
				Extension string "json:\"extension\""
			}{AreaCode: "01", Number: "1111-1111", Extension: ""}, Type: "registered", EntityType: interface{}(nil), ID: "695869370"}, Metadata: struct{}{}, AdditionalInfo: struct {
				Items []struct {
					ID          string  "json:\"id\""
					Title       string  "json:\"title\""
					Description string  "json:\"description\""
					PictureURL  string  "json:\"picture_url\""
					CategoryID  string  "json:\"category_id\""
					Quantity    int     "json:\"quantity\""
					UnitPrice   float64 "json:\"unit_price\""
				} "json:\"items\""
				Payer struct {
					Phone struct {
						AreaCode string "json:\"area_code\""
						Number   string "json:\"number\""
					} "json:\"phone\""
					Address struct {
						ZipCode      string "json:\"zip_code\""
						StreetName   string "json:\"street_name\""
						StreetNumber string "json:\"street_number\""
					} "json:\"address\""
					FirstName        string "json:\"first_name\""
					LastName         string "json:\"last_name\""
					RegistrationDate string "json:\"registration_date\""
				} "json:\"payer\""
				Shipments struct {
					ReceiverAddress struct {
						StreetName   string "json:\"street_name\""
						StreetNumber int    "json:\"street_number\""
						ZipCode      string "json:\"zip_code\""
						CityName     string "json:\"city_name\""
						StateName    string "json:\"state_name\""
					} "json:\"receiver_address\""
				} "json:\"shipments\""
			}{Items: []struct {
				ID          string  "json:\"id\""
				Title       string  "json:\"title\""
				Description string  "json:\"description\""
				PictureURL  string  "json:\"picture_url\""
				CategoryID  string  "json:\"category_id\""
				Quantity    int     "json:\"quantity\""
				UnitPrice   float64 "json:\"unit_price\""
			}(nil), Payer: struct {
				Phone struct {
					AreaCode string "json:\"area_code\""
					Number   string "json:\"number\""
				} "json:\"phone\""
				Address struct {
					ZipCode      string "json:\"zip_code\""
					StreetName   string "json:\"street_name\""
					StreetNumber string "json:\"street_number\""
				} "json:\"address\""
				FirstName        string "json:\"first_name\""
				LastName         string "json:\"last_name\""
				RegistrationDate string "json:\"registration_date\""
			}{Phone: struct {
				AreaCode string "json:\"area_code\""
				Number   string "json:\"number\""
			}{AreaCode: "", Number: ""}, Address: struct {
				ZipCode      string "json:\"zip_code\""
				StreetName   string "json:\"street_name\""
				StreetNumber string "json:\"street_number\""
			}{ZipCode: "", StreetName: "", StreetNumber: ""}, FirstName: "", LastName: "", RegistrationDate: ""}, Shipments: struct {
				ReceiverAddress struct {
					StreetName   string "json:\"street_name\""
					StreetNumber int    "json:\"street_number\""
					ZipCode      string "json:\"zip_code\""
					CityName     string "json:\"city_name\""
					StateName    string "json:\"state_name\""
				} "json:\"receiver_address\""
			}{ReceiverAddress: struct {
				StreetName   string "json:\"street_name\""
				StreetNumber int    "json:\"street_number\""
				ZipCode      string "json:\"zip_code\""
				CityName     string "json:\"city_name\""
				StateName    string "json:\"state_name\""
			}{StreetName: "", StreetNumber: 0, ZipCode: "", CityName: "", StateName: ""}}}, Order: struct{}{}, ExternalReference: "", TransactionAmount: 2000, TransactionAmountRefunded: 0, CouponAmount: 0, DifferentialPricingID: interface{}(nil), DeductionSchema: interface{}(nil), TransactionDetails: struct {
				PaymentMethodReferenceID interface{} "json:\"payment_method_reference_id\""
				NetReceivedAmount        float64     "json:\"net_received_amount\""
				TotalPaidAmount          float64     "json:\"total_paid_amount\""
				OverpaidAmount           int         "json:\"overpaid_amount\""
				ExternalResourceURL      interface{} "json:\"external_resource_url\""
				InstallmentAmount        float64     "json:\"installment_amount\""
				FinancialInstitution     interface{} "json:\"financial_institution\""
				PayableDeferralPeriod    interface{} "json:\"payable_deferral_period\""
				AcquirerReference        interface{} "json:\"acquirer_reference\""
			}{PaymentMethodReferenceID: interface{}(nil), NetReceivedAmount: 0, TotalPaidAmount: 2000, OverpaidAmount: 0, ExternalResourceURL: interface{}(nil), InstallmentAmount: 2000, FinancialInstitution: interface{}(nil), PayableDeferralPeriod: interface{}(nil), AcquirerReference: interface{}(nil)}, FeeDetails: []struct {
				Type     string  "json:\"type\""
				Amount   float64 "json:\"amount\""
				FeePayer string  "json:\"fee_payer\""
			}{}, Captured: true, BinaryMode: false, CallForAuthorizeID: interface{}(nil), StatementDescriptor: "MERCADOPAGO", Installments: 1, Card: struct {
				ID              interface{} "json:\"id\""
				FirstSixDigits  string      "json:\"first_six_digits\""
				LastFourDigits  string      "json:\"last_four_digits\""
				ExpirationMonth int         "json:\"expiration_month\""
				ExpirationYear  int         "json:\"expiration_year\""
				DateCreated     string      "json:\"date_created\""
				DateLastUpdated string      "json:\"date_last_updated\""
				Cardholder      struct {
					Name           string "json:\"name\""
					Identification struct {
						Number string "json:\"number\""
						Type   string "json:\"type\""
					} "json:\"identification\""
				} "json:\"cardholder\""
			}{ID: interface{}(nil), FirstSixDigits: "374378", LastFourDigits: "5283", ExpirationMonth: 11, ExpirationYear: 2025, DateCreated: "2021-01-01T22:42:18.000-04:00", DateLastUpdated: "2021-01-01T22:42:18.000-04:00", Cardholder: struct {
				Name           string "json:\"name\""
				Identification struct {
					Number string "json:\"number\""
					Type   string "json:\"type\""
				} "json:\"identification\""
			}{Name: "OTHE", Identification: struct {
				Number string "json:\"number\""
				Type   string "json:\"type\""
			}{Number: "23090923", Type: "CC"}}}, NotificationURL: "", Refunds: []interface{}{}, ProcessingMode: "aggregator", MerchantAccountID: interface{}(nil), Acquirer: interface{}(nil), MerchantNumber: interface{}(nil), AcquirerReconciliation: []interface{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accessToken := "access-token"

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, "/payments", r.URL.Path)

				assert.Equal(t, "Bearer "+accessToken, r.Header.Get("Authorization"))

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
