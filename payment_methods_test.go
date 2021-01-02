package mercadopago_test

import (
	"github.com/fabianMendez/mercadopago"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestClient_GetPaymentMethods(t *testing.T) {
	tests := []struct {
		name             string
		publicKey        string
		accessToken      string
		expectedResponse mercadopago.PaymentMethods
		expectedErr      error
		respStatus       int
		respBody         string
	}{
		{
			name:        "successful response using access token",
			accessToken: "access-token",
			respStatus:  http.StatusOK,
			respBody:    `[{"id":"diners","name":"Diners","payment_type_id":"credit_card","status":"active","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/diners.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/diners.gif","deferred_capture":"unsupported","settings":[{"card_number":{"validation":"standard","length":14},"bin":{"pattern":"^((30)|(36)|(38))","installments_pattern":"^((30)|(36)|(38))","exclusion_pattern":null},"security_code":{"length":3,"card_location":"back","mode":"mandatory"}}],"additional_info_needed":["cardholder_name","cardholder_identification_number","cardholder_identification_type"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"accreditation_time":2880,"financial_institutions":[],"processing_modes":["aggregator"]},{"id":"codensa","name":"Crédito Fácil Codensa","payment_type_id":"credit_card","status":"active","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/codensa.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/codensa.gif","deferred_capture":"unsupported","settings":[{"card_number":{"validation":"standard","length":16},"bin":{"pattern":"^(590712|529448)","installments_pattern":"^(590712|529448)","exclusion_pattern":null},"security_code":{"length":3,"card_location":"back","mode":"mandatory"}}],"additional_info_needed":["cardholder_name","cardholder_identification_number","cardholder_identification_type"],"min_allowed_amount":1000,"max_allowed_amount":5000000,"accreditation_time":null,"financial_institutions":[],"processing_modes":["aggregator"]},{"id":"debvisa","name":"Visa Débito","payment_type_id":"debit_card","status":"active","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/debvisa.gif","thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/debvisa.gif","deferred_capture":"unsupported","settings":[{"card_number":{"validation":"standard","length":16},"bin":{"pattern":"^(488233|462896|484192|441509|486367|431385|454106|489635|455982|498476|492468|491268|459317|418253|404279|423949|457605|402739|450942|457604|457603|491511)","installments_pattern":null,"exclusion_pattern":null},"security_code":{"length":3,"card_location":"back","mode":"mandatory"}}],"additional_info_needed":["cardholder_name","cardholder_identification_type","cardholder_identification_number"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"accreditation_time":0,"financial_institutions":[],"processing_modes":["aggregator"]},{"id":"debmaster","name":"Mastercard Débito","payment_type_id":"debit_card","status":"active","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/debmaster.gif","thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/debmaster.gif","deferred_capture":"unsupported","settings":[{"card_number":{"validation":"standard","length":16},"bin":{"pattern":"^(533254|534778|557555|525315|537980|523719|559749|530715|530716|530717|530710|530721|514332|529768|536782|552558|530729|530712|530724|530727|530695|510342|530719|530713|530720|530725|530726|530691|530711|530728|530714|530723|528633|524627|530718|530722|536126|518503|548185|525358|545409|535803|547692|520922|517404|516451|518092|517393|544039|515516|526557)","installments_pattern":"","exclusion_pattern":null},"security_code":{"length":3,"card_location":"back","mode":"mandatory"}}],"additional_info_needed":["cardholder_name","cardholder_identification_type","cardholder_identification_number"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"accreditation_time":0,"financial_institutions":[],"processing_modes":["aggregator"]},{"id":"visa","name":"Visa","payment_type_id":"credit_card","status":"active","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/visa.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/visa.gif","deferred_capture":"unsupported","settings":[{"card_number":{"validation":"standard","length":16},"bin":{"pattern":"^(4)","installments_pattern":"^(4)","exclusion_pattern":"^(488233|462896|484192|441509|486367|431385|454106|489635|455982|498476|492468|491268|459317|418253|404279|423949|457605|402739|450942|457604|457603|491511)"},"security_code":{"length":3,"card_location":"back","mode":"mandatory"}}],"additional_info_needed":["cardholder_name","cardholder_identification_number","cardholder_identification_type"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"accreditation_time":2880,"financial_institutions":[],"processing_modes":["aggregator"]},{"id":"amex","name":"American Express","payment_type_id":"credit_card","status":"active","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/amex.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/amex.gif","deferred_capture":"unsupported","settings":[{"card_number":{"validation":"standard","length":15},"bin":{"pattern":"^((34)|(37))","installments_pattern":"^((34)|(37))","exclusion_pattern":null},"security_code":{"length":4,"card_location":"front","mode":"mandatory"}}],"additional_info_needed":["cardholder_name","cardholder_identification_number","cardholder_identification_type"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"accreditation_time":2880,"financial_institutions":[],"processing_modes":["aggregator"]},{"id":"master","name":"Mastercard","payment_type_id":"credit_card","status":"active","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/master.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/master.gif","deferred_capture":"unsupported","settings":[{"card_number":{"validation":"standard","length":16},"bin":{"pattern":"^(5|(2(221|222|223|224|225|226|227|228|229|23|24|25|26|27|28|29|3|4|5|6|70|71|720)))","installments_pattern":"^(5|(2(221|222|223|224|225|226|227|228|229|23|24|25|26|27|28|29|3|4|5|6|70|71|720)))","exclusion_pattern":"^(590712|529448|533254|534778|557555|525315|537980|523719|559749|530715|530716|530717|530710|530721|514332|529768|536782|552558|530729|530712|530724|530727|530695|510342|530719|530713|530720|530725|530726|530691|530711|530728|530714|530723|528633|524627|530718|530722|536126|518503|548185|525358|545409|535803|547692|520922|517404|516451|518092|517393|544039|515516|526557)"},"security_code":{"length":3,"card_location":"back","mode":"mandatory"}}],"additional_info_needed":["cardholder_name","cardholder_identification_number","cardholder_identification_type"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"accreditation_time":2880,"financial_institutions":[],"processing_modes":["aggregator"]},{"id":"efecty","name":"Efecty","payment_type_id":"ticket","status":"active","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/efecty.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/efecty.gif","deferred_capture":"does_not_apply","settings":[],"additional_info_needed":[],"min_allowed_amount":5000,"max_allowed_amount":4000000,"accreditation_time":0,"financial_institutions":[],"processing_modes":["aggregator"]},{"id":"pse","name":"PSE","payment_type_id":"bank_transfer","status":"active","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/pse.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/pse.gif","deferred_capture":"does_not_apply","settings":[],"additional_info_needed":["entity_type"],"min_allowed_amount":1600,"max_allowed_amount":30000000,"accreditation_time":30,"financial_institutions":[{"id":"1040","description":"Banco Agrario"},{"id":"1507","description":"NEQUI"},{"id":"1052","description":"Banco AV Villas"},{"id":"1032","description":"Banco Caja Social"},{"id":"1019","description":"SCOTIABANK COLPATRIA"},{"id":"1066","description":"Banco Cooperativo Coopcentral"},{"id":"1051","description":"Banco Davivienda"},{"id":"1001","description":"Banco De Bogota"},{"id":"1023","description":"Banco De Occidente"},{"id":"1062","description":"Banco Falabella"},{"id":"1012","description":"Banco GNB Sudameris"},{"id":"1060","description":"Banco Pichincha S.A."},{"id":"1002","description":"Banco Popular"},{"id":"1058","description":"Banco Procredit"},{"id":"1007","description":"Bancolombia"},{"id":"1061","description":"Bancoomeva S.A."},{"id":"1013","description":"Banco BBVA Colombia S.A."},{"id":"1009","description":"Citibank"},{"id":"1006","description":"Banco Itaú"},{"id":"1292","description":"Confiar Cooperativa Financiera"},{"id":"1551","description":"DaviPlata"},{"id":"1069","description":"BANCO SERFINANZA"},{"id":"1065","description":"BANCO SANTANDER COLOMBIA"},{"id":"1283","description":"CFA COOPERATIVA FINANCIERA"},{"id":"1059","description":"BANCO DE LAS MICROFINANZAS BANCAMIA"},{"id":"1151","description":"RAPPIPAY"},{"id":"1289","description":"Cotrafa"},{"id":"1370","description":"COLTEFINANCIERA"}],"processing_modes":["aggregator"]},{"id":"baloto","name":"Baloto","payment_type_id":"ticket","status":"active","secure_thumbnail":"https://http2.mlstatic.com/storage/logos-api-admin/bb3d20a0-99d5-11e9-a52b-295618a86fe2-xs.svg","thumbnail":"https://http2.mlstatic.com/storage/logos-api-admin/bb3d20a0-99d5-11e9-a52b-295618a86fe2-xs.svg","deferred_capture":"does_not_apply","settings":[],"additional_info_needed":[],"min_allowed_amount":1500,"max_allowed_amount":1000000,"accreditation_time":0,"financial_institutions":[],"processing_modes":["aggregator"]}]`,
			expectedResponse: mercadopago.PaymentMethods{
				{ID: "diners", Name: "Diners", PaymentTypeID: "credit_card", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/diners.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/diners.gif", DeferredCapture: "unsupported", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{{CardNumber: struct {
					Validation string "json:\"validation\""
					Length     int    "json:\"length\""
				}{Validation: "standard", Length: 14}, Bin: struct {
					Pattern             string "json:\"pattern\""
					InstallmentsPattern string "json:\"installments_pattern\""
					ExclusionPattern    string "json:\"exclusion_pattern\""
				}{Pattern: "^((30)|(36)|(38))", InstallmentsPattern: "^((30)|(36)|(38))", ExclusionPattern: ""}, SecurityCode: struct {
					Length       int    "json:\"length\""
					CardLocation string "json:\"card_location\""
					Mode         string "json:\"mode\""
				}{Length: 3, CardLocation: "back", Mode: "mandatory"}}}, AdditionalInfoNeeded: []string{"cardholder_name", "cardholder_identification_number", "cardholder_identification_type"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, AccreditationTime: 2880, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
				{ID: "codensa", Name: "Crédito Fácil Codensa", PaymentTypeID: "credit_card", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/codensa.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/codensa.gif", DeferredCapture: "unsupported", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{{CardNumber: struct {
					Validation string "json:\"validation\""
					Length     int    "json:\"length\""
				}{Validation: "standard", Length: 16}, Bin: struct {
					Pattern             string "json:\"pattern\""
					InstallmentsPattern string "json:\"installments_pattern\""
					ExclusionPattern    string "json:\"exclusion_pattern\""
				}{Pattern: "^(590712|529448)", InstallmentsPattern: "^(590712|529448)", ExclusionPattern: ""}, SecurityCode: struct {
					Length       int    "json:\"length\""
					CardLocation string "json:\"card_location\""
					Mode         string "json:\"mode\""
				}{Length: 3, CardLocation: "back", Mode: "mandatory"}}}, AdditionalInfoNeeded: []string{"cardholder_name", "cardholder_identification_number", "cardholder_identification_type"}, MinAllowedAmount: 1000, MaxAllowedAmount: 5000000, AccreditationTime: 0, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
				{ID: "debvisa", Name: "Visa Débito", PaymentTypeID: "debit_card", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/debvisa.gif", Thumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/debvisa.gif", DeferredCapture: "unsupported", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{{CardNumber: struct {
					Validation string "json:\"validation\""
					Length     int    "json:\"length\""
				}{Validation: "standard", Length: 16}, Bin: struct {
					Pattern             string "json:\"pattern\""
					InstallmentsPattern string "json:\"installments_pattern\""
					ExclusionPattern    string "json:\"exclusion_pattern\""
				}{Pattern: "^(488233|462896|484192|441509|486367|431385|454106|489635|455982|498476|492468|491268|459317|418253|404279|423949|457605|402739|450942|457604|457603|491511)", InstallmentsPattern: "", ExclusionPattern: ""}, SecurityCode: struct {
					Length       int    "json:\"length\""
					CardLocation string "json:\"card_location\""
					Mode         string "json:\"mode\""
				}{Length: 3, CardLocation: "back", Mode: "mandatory"}}}, AdditionalInfoNeeded: []string{"cardholder_name", "cardholder_identification_type", "cardholder_identification_number"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, AccreditationTime: 0, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
				{ID: "debmaster", Name: "Mastercard Débito", PaymentTypeID: "debit_card", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/debmaster.gif", Thumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/debmaster.gif", DeferredCapture: "unsupported", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{{CardNumber: struct {
					Validation string "json:\"validation\""
					Length     int    "json:\"length\""
				}{Validation: "standard", Length: 16}, Bin: struct {
					Pattern             string "json:\"pattern\""
					InstallmentsPattern string "json:\"installments_pattern\""
					ExclusionPattern    string "json:\"exclusion_pattern\""
				}{Pattern: "^(533254|534778|557555|525315|537980|523719|559749|530715|530716|530717|530710|530721|514332|529768|536782|552558|530729|530712|530724|530727|530695|510342|530719|530713|530720|530725|530726|530691|530711|530728|530714|530723|528633|524627|530718|530722|536126|518503|548185|525358|545409|535803|547692|520922|517404|516451|518092|517393|544039|515516|526557)", InstallmentsPattern: "", ExclusionPattern: ""}, SecurityCode: struct {
					Length       int    "json:\"length\""
					CardLocation string "json:\"card_location\""
					Mode         string "json:\"mode\""
				}{Length: 3, CardLocation: "back", Mode: "mandatory"}}}, AdditionalInfoNeeded: []string{"cardholder_name", "cardholder_identification_type", "cardholder_identification_number"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, AccreditationTime: 0, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
				{ID: "visa", Name: "Visa", PaymentTypeID: "credit_card", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/visa.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/visa.gif", DeferredCapture: "unsupported", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{{CardNumber: struct {
					Validation string "json:\"validation\""
					Length     int    "json:\"length\""
				}{Validation: "standard", Length: 16}, Bin: struct {
					Pattern             string "json:\"pattern\""
					InstallmentsPattern string "json:\"installments_pattern\""
					ExclusionPattern    string "json:\"exclusion_pattern\""
				}{Pattern: "^(4)", InstallmentsPattern: "^(4)", ExclusionPattern: "^(488233|462896|484192|441509|486367|431385|454106|489635|455982|498476|492468|491268|459317|418253|404279|423949|457605|402739|450942|457604|457603|491511)"}, SecurityCode: struct {
					Length       int    "json:\"length\""
					CardLocation string "json:\"card_location\""
					Mode         string "json:\"mode\""
				}{Length: 3, CardLocation: "back", Mode: "mandatory"}}}, AdditionalInfoNeeded: []string{"cardholder_name", "cardholder_identification_number", "cardholder_identification_type"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, AccreditationTime: 2880, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
				{ID: "amex", Name: "American Express", PaymentTypeID: "credit_card", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/amex.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/amex.gif", DeferredCapture: "unsupported", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{{CardNumber: struct {
					Validation string "json:\"validation\""
					Length     int    "json:\"length\""
				}{Validation: "standard", Length: 15}, Bin: struct {
					Pattern             string "json:\"pattern\""
					InstallmentsPattern string "json:\"installments_pattern\""
					ExclusionPattern    string "json:\"exclusion_pattern\""
				}{Pattern: "^((34)|(37))", InstallmentsPattern: "^((34)|(37))", ExclusionPattern: ""}, SecurityCode: struct {
					Length       int    "json:\"length\""
					CardLocation string "json:\"card_location\""
					Mode         string "json:\"mode\""
				}{Length: 4, CardLocation: "front", Mode: "mandatory"}}}, AdditionalInfoNeeded: []string{"cardholder_name", "cardholder_identification_number", "cardholder_identification_type"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, AccreditationTime: 2880, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
				{ID: "master", Name: "Mastercard", PaymentTypeID: "credit_card", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/master.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/master.gif", DeferredCapture: "unsupported", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{{CardNumber: struct {
					Validation string "json:\"validation\""
					Length     int    "json:\"length\""
				}{Validation: "standard", Length: 16}, Bin: struct {
					Pattern             string "json:\"pattern\""
					InstallmentsPattern string "json:\"installments_pattern\""
					ExclusionPattern    string "json:\"exclusion_pattern\""
				}{Pattern: "^(5|(2(221|222|223|224|225|226|227|228|229|23|24|25|26|27|28|29|3|4|5|6|70|71|720)))", InstallmentsPattern: "^(5|(2(221|222|223|224|225|226|227|228|229|23|24|25|26|27|28|29|3|4|5|6|70|71|720)))", ExclusionPattern: "^(590712|529448|533254|534778|557555|525315|537980|523719|559749|530715|530716|530717|530710|530721|514332|529768|536782|552558|530729|530712|530724|530727|530695|510342|530719|530713|530720|530725|530726|530691|530711|530728|530714|530723|528633|524627|530718|530722|536126|518503|548185|525358|545409|535803|547692|520922|517404|516451|518092|517393|544039|515516|526557)"}, SecurityCode: struct {
					Length       int    "json:\"length\""
					CardLocation string "json:\"card_location\""
					Mode         string "json:\"mode\""
				}{Length: 3, CardLocation: "back", Mode: "mandatory"}}}, AdditionalInfoNeeded: []string{"cardholder_name", "cardholder_identification_number", "cardholder_identification_type"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, AccreditationTime: 2880, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
				{ID: "efecty", Name: "Efecty", PaymentTypeID: "ticket", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/efecty.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/efecty.gif", DeferredCapture: "does_not_apply", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{}, AdditionalInfoNeeded: []string{}, MinAllowedAmount: 5000, MaxAllowedAmount: 4000000, AccreditationTime: 0, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
				{ID: "pse", Name: "PSE", PaymentTypeID: "bank_transfer", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/pse.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/pse.gif", DeferredCapture: "does_not_apply", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{}, AdditionalInfoNeeded: []string{"entity_type"}, MinAllowedAmount: 1600, MaxAllowedAmount: 30000000, AccreditationTime: 30, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{{ID: "1040", Description: "Banco Agrario"}, {ID: "1507", Description: "NEQUI"}, {ID: "1052", Description: "Banco AV Villas"}, {ID: "1032", Description: "Banco Caja Social"}, {ID: "1019", Description: "SCOTIABANK COLPATRIA"}, {ID: "1066", Description: "Banco Cooperativo Coopcentral"}, {ID: "1051", Description: "Banco Davivienda"}, {ID: "1001", Description: "Banco De Bogota"}, {ID: "1023", Description: "Banco De Occidente"}, {ID: "1062", Description: "Banco Falabella"}, {ID: "1012", Description: "Banco GNB Sudameris"}, {ID: "1060", Description: "Banco Pichincha S.A."}, {ID: "1002", Description: "Banco Popular"}, {ID: "1058", Description: "Banco Procredit"}, {ID: "1007", Description: "Bancolombia"}, {ID: "1061", Description: "Bancoomeva S.A."}, {ID: "1013", Description: "Banco BBVA Colombia S.A."}, {ID: "1009", Description: "Citibank"}, {ID: "1006", Description: "Banco Itaú"}, {ID: "1292", Description: "Confiar Cooperativa Financiera"}, {ID: "1551", Description: "DaviPlata"}, {ID: "1069", Description: "BANCO SERFINANZA"}, {ID: "1065", Description: "BANCO SANTANDER COLOMBIA"}, {ID: "1283", Description: "CFA COOPERATIVA FINANCIERA"}, {ID: "1059", Description: "BANCO DE LAS MICROFINANZAS BANCAMIA"}, {ID: "1151", Description: "RAPPIPAY"}, {ID: "1289", Description: "Cotrafa"}, {ID: "1370", Description: "COLTEFINANCIERA"}}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
				{ID: "baloto", Name: "Baloto", PaymentTypeID: "ticket", Status: "active", SecureThumbnail: "https://http2.mlstatic.com/storage/logos-api-admin/bb3d20a0-99d5-11e9-a52b-295618a86fe2-xs.svg", Thumbnail: "https://http2.mlstatic.com/storage/logos-api-admin/bb3d20a0-99d5-11e9-a52b-295618a86fe2-xs.svg", DeferredCapture: "does_not_apply", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{}, AdditionalInfoNeeded: []string{}, MinAllowedAmount: 1500, MaxAllowedAmount: 1000000, AccreditationTime: 0, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
			},
		},
		{
			name:       "successful response using public key",
			publicKey:  "public-key",
			respStatus: http.StatusOK,
			respBody:   `[{"id":"diners","name":"Diners","payment_type_id":"credit_card","status":"active","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/diners.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/diners.gif","deferred_capture":"unsupported","settings":[{"card_number":{"validation":"standard","length":14},"bin":{"pattern":"^((30)|(36)|(38))","installments_pattern":"^((30)|(36)|(38))","exclusion_pattern":null},"security_code":{"length":3,"card_location":"back","mode":"mandatory"}}],"additional_info_needed":["cardholder_name","cardholder_identification_number","cardholder_identification_type"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"accreditation_time":2880,"financial_institutions":[],"processing_modes":["aggregator"]},{"id":"codensa","name":"Crédito Fácil Codensa","payment_type_id":"credit_card","status":"active","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/codensa.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/codensa.gif","deferred_capture":"unsupported","settings":[{"card_number":{"validation":"standard","length":16},"bin":{"pattern":"^(590712|529448)","installments_pattern":"^(590712|529448)","exclusion_pattern":null},"security_code":{"length":3,"card_location":"back","mode":"mandatory"}}],"additional_info_needed":["cardholder_name","cardholder_identification_number","cardholder_identification_type"],"min_allowed_amount":1000,"max_allowed_amount":5000000,"accreditation_time":null,"financial_institutions":[],"processing_modes":["aggregator"]},{"id":"debvisa","name":"Visa Débito","payment_type_id":"debit_card","status":"active","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/debvisa.gif","thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/debvisa.gif","deferred_capture":"unsupported","settings":[{"card_number":{"validation":"standard","length":16},"bin":{"pattern":"^(488233|462896|484192|441509|486367|431385|454106|489635|455982|498476|492468|491268|459317|418253|404279|423949|457605|402739|450942|457604|457603|491511)","installments_pattern":null,"exclusion_pattern":null},"security_code":{"length":3,"card_location":"back","mode":"mandatory"}}],"additional_info_needed":["cardholder_name","cardholder_identification_type","cardholder_identification_number"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"accreditation_time":0,"financial_institutions":[],"processing_modes":["aggregator"]},{"id":"debmaster","name":"Mastercard Débito","payment_type_id":"debit_card","status":"active","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/debmaster.gif","thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/debmaster.gif","deferred_capture":"unsupported","settings":[{"card_number":{"validation":"standard","length":16},"bin":{"pattern":"^(533254|534778|557555|525315|537980|523719|559749|530715|530716|530717|530710|530721|514332|529768|536782|552558|530729|530712|530724|530727|530695|510342|530719|530713|530720|530725|530726|530691|530711|530728|530714|530723|528633|524627|530718|530722|536126|518503|548185|525358|545409|535803|547692|520922|517404|516451|518092|517393|544039|515516|526557)","installments_pattern":"","exclusion_pattern":null},"security_code":{"length":3,"card_location":"back","mode":"mandatory"}}],"additional_info_needed":["cardholder_name","cardholder_identification_type","cardholder_identification_number"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"accreditation_time":0,"financial_institutions":[],"processing_modes":["aggregator"]},{"id":"visa","name":"Visa","payment_type_id":"credit_card","status":"active","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/visa.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/visa.gif","deferred_capture":"unsupported","settings":[{"card_number":{"validation":"standard","length":16},"bin":{"pattern":"^(4)","installments_pattern":"^(4)","exclusion_pattern":"^(488233|462896|484192|441509|486367|431385|454106|489635|455982|498476|492468|491268|459317|418253|404279|423949|457605|402739|450942|457604|457603|491511)"},"security_code":{"length":3,"card_location":"back","mode":"mandatory"}}],"additional_info_needed":["cardholder_name","cardholder_identification_number","cardholder_identification_type"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"accreditation_time":2880,"financial_institutions":[],"processing_modes":["aggregator"]},{"id":"amex","name":"American Express","payment_type_id":"credit_card","status":"active","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/amex.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/amex.gif","deferred_capture":"unsupported","settings":[{"card_number":{"validation":"standard","length":15},"bin":{"pattern":"^((34)|(37))","installments_pattern":"^((34)|(37))","exclusion_pattern":null},"security_code":{"length":4,"card_location":"front","mode":"mandatory"}}],"additional_info_needed":["cardholder_name","cardholder_identification_number","cardholder_identification_type"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"accreditation_time":2880,"financial_institutions":[],"processing_modes":["aggregator"]},{"id":"master","name":"Mastercard","payment_type_id":"credit_card","status":"active","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/master.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/master.gif","deferred_capture":"unsupported","settings":[{"card_number":{"validation":"standard","length":16},"bin":{"pattern":"^(5|(2(221|222|223|224|225|226|227|228|229|23|24|25|26|27|28|29|3|4|5|6|70|71|720)))","installments_pattern":"^(5|(2(221|222|223|224|225|226|227|228|229|23|24|25|26|27|28|29|3|4|5|6|70|71|720)))","exclusion_pattern":"^(590712|529448|533254|534778|557555|525315|537980|523719|559749|530715|530716|530717|530710|530721|514332|529768|536782|552558|530729|530712|530724|530727|530695|510342|530719|530713|530720|530725|530726|530691|530711|530728|530714|530723|528633|524627|530718|530722|536126|518503|548185|525358|545409|535803|547692|520922|517404|516451|518092|517393|544039|515516|526557)"},"security_code":{"length":3,"card_location":"back","mode":"mandatory"}}],"additional_info_needed":["cardholder_name","cardholder_identification_number","cardholder_identification_type"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"accreditation_time":2880,"financial_institutions":[],"processing_modes":["aggregator"]},{"id":"efecty","name":"Efecty","payment_type_id":"ticket","status":"active","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/efecty.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/efecty.gif","deferred_capture":"does_not_apply","settings":[],"additional_info_needed":[],"min_allowed_amount":5000,"max_allowed_amount":4000000,"accreditation_time":0,"financial_institutions":[],"processing_modes":["aggregator"]},{"id":"pse","name":"PSE","payment_type_id":"bank_transfer","status":"active","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/pse.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/pse.gif","deferred_capture":"does_not_apply","settings":[],"additional_info_needed":["entity_type"],"min_allowed_amount":1600,"max_allowed_amount":30000000,"accreditation_time":30,"financial_institutions":[{"id":"1040","description":"Banco Agrario"},{"id":"1507","description":"NEQUI"},{"id":"1052","description":"Banco AV Villas"},{"id":"1032","description":"Banco Caja Social"},{"id":"1019","description":"SCOTIABANK COLPATRIA"},{"id":"1066","description":"Banco Cooperativo Coopcentral"},{"id":"1051","description":"Banco Davivienda"},{"id":"1001","description":"Banco De Bogota"},{"id":"1023","description":"Banco De Occidente"},{"id":"1062","description":"Banco Falabella"},{"id":"1012","description":"Banco GNB Sudameris"},{"id":"1060","description":"Banco Pichincha S.A."},{"id":"1002","description":"Banco Popular"},{"id":"1058","description":"Banco Procredit"},{"id":"1007","description":"Bancolombia"},{"id":"1061","description":"Bancoomeva S.A."},{"id":"1013","description":"Banco BBVA Colombia S.A."},{"id":"1009","description":"Citibank"},{"id":"1006","description":"Banco Itaú"},{"id":"1292","description":"Confiar Cooperativa Financiera"},{"id":"1551","description":"DaviPlata"},{"id":"1069","description":"BANCO SERFINANZA"},{"id":"1065","description":"BANCO SANTANDER COLOMBIA"},{"id":"1283","description":"CFA COOPERATIVA FINANCIERA"},{"id":"1059","description":"BANCO DE LAS MICROFINANZAS BANCAMIA"},{"id":"1151","description":"RAPPIPAY"},{"id":"1289","description":"Cotrafa"},{"id":"1370","description":"COLTEFINANCIERA"}],"processing_modes":["aggregator"]},{"id":"baloto","name":"Baloto","payment_type_id":"ticket","status":"active","secure_thumbnail":"https://http2.mlstatic.com/storage/logos-api-admin/bb3d20a0-99d5-11e9-a52b-295618a86fe2-xs.svg","thumbnail":"https://http2.mlstatic.com/storage/logos-api-admin/bb3d20a0-99d5-11e9-a52b-295618a86fe2-xs.svg","deferred_capture":"does_not_apply","settings":[],"additional_info_needed":[],"min_allowed_amount":1500,"max_allowed_amount":1000000,"accreditation_time":0,"financial_institutions":[],"processing_modes":["aggregator"]}]`,
			expectedResponse: mercadopago.PaymentMethods{
				{ID: "diners", Name: "Diners", PaymentTypeID: "credit_card", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/diners.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/diners.gif", DeferredCapture: "unsupported", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{{CardNumber: struct {
					Validation string "json:\"validation\""
					Length     int    "json:\"length\""
				}{Validation: "standard", Length: 14}, Bin: struct {
					Pattern             string "json:\"pattern\""
					InstallmentsPattern string "json:\"installments_pattern\""
					ExclusionPattern    string "json:\"exclusion_pattern\""
				}{Pattern: "^((30)|(36)|(38))", InstallmentsPattern: "^((30)|(36)|(38))", ExclusionPattern: ""}, SecurityCode: struct {
					Length       int    "json:\"length\""
					CardLocation string "json:\"card_location\""
					Mode         string "json:\"mode\""
				}{Length: 3, CardLocation: "back", Mode: "mandatory"}}}, AdditionalInfoNeeded: []string{"cardholder_name", "cardholder_identification_number", "cardholder_identification_type"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, AccreditationTime: 2880, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
				{ID: "codensa", Name: "Crédito Fácil Codensa", PaymentTypeID: "credit_card", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/codensa.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/codensa.gif", DeferredCapture: "unsupported", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{{CardNumber: struct {
					Validation string "json:\"validation\""
					Length     int    "json:\"length\""
				}{Validation: "standard", Length: 16}, Bin: struct {
					Pattern             string "json:\"pattern\""
					InstallmentsPattern string "json:\"installments_pattern\""
					ExclusionPattern    string "json:\"exclusion_pattern\""
				}{Pattern: "^(590712|529448)", InstallmentsPattern: "^(590712|529448)", ExclusionPattern: ""}, SecurityCode: struct {
					Length       int    "json:\"length\""
					CardLocation string "json:\"card_location\""
					Mode         string "json:\"mode\""
				}{Length: 3, CardLocation: "back", Mode: "mandatory"}}}, AdditionalInfoNeeded: []string{"cardholder_name", "cardholder_identification_number", "cardholder_identification_type"}, MinAllowedAmount: 1000, MaxAllowedAmount: 5000000, AccreditationTime: 0, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
				{ID: "debvisa", Name: "Visa Débito", PaymentTypeID: "debit_card", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/debvisa.gif", Thumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/debvisa.gif", DeferredCapture: "unsupported", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{{CardNumber: struct {
					Validation string "json:\"validation\""
					Length     int    "json:\"length\""
				}{Validation: "standard", Length: 16}, Bin: struct {
					Pattern             string "json:\"pattern\""
					InstallmentsPattern string "json:\"installments_pattern\""
					ExclusionPattern    string "json:\"exclusion_pattern\""
				}{Pattern: "^(488233|462896|484192|441509|486367|431385|454106|489635|455982|498476|492468|491268|459317|418253|404279|423949|457605|402739|450942|457604|457603|491511)", InstallmentsPattern: "", ExclusionPattern: ""}, SecurityCode: struct {
					Length       int    "json:\"length\""
					CardLocation string "json:\"card_location\""
					Mode         string "json:\"mode\""
				}{Length: 3, CardLocation: "back", Mode: "mandatory"}}}, AdditionalInfoNeeded: []string{"cardholder_name", "cardholder_identification_type", "cardholder_identification_number"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, AccreditationTime: 0, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
				{ID: "debmaster", Name: "Mastercard Débito", PaymentTypeID: "debit_card", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/debmaster.gif", Thumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/debmaster.gif", DeferredCapture: "unsupported", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{{CardNumber: struct {
					Validation string "json:\"validation\""
					Length     int    "json:\"length\""
				}{Validation: "standard", Length: 16}, Bin: struct {
					Pattern             string "json:\"pattern\""
					InstallmentsPattern string "json:\"installments_pattern\""
					ExclusionPattern    string "json:\"exclusion_pattern\""
				}{Pattern: "^(533254|534778|557555|525315|537980|523719|559749|530715|530716|530717|530710|530721|514332|529768|536782|552558|530729|530712|530724|530727|530695|510342|530719|530713|530720|530725|530726|530691|530711|530728|530714|530723|528633|524627|530718|530722|536126|518503|548185|525358|545409|535803|547692|520922|517404|516451|518092|517393|544039|515516|526557)", InstallmentsPattern: "", ExclusionPattern: ""}, SecurityCode: struct {
					Length       int    "json:\"length\""
					CardLocation string "json:\"card_location\""
					Mode         string "json:\"mode\""
				}{Length: 3, CardLocation: "back", Mode: "mandatory"}}}, AdditionalInfoNeeded: []string{"cardholder_name", "cardholder_identification_type", "cardholder_identification_number"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, AccreditationTime: 0, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
				{ID: "visa", Name: "Visa", PaymentTypeID: "credit_card", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/visa.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/visa.gif", DeferredCapture: "unsupported", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{{CardNumber: struct {
					Validation string "json:\"validation\""
					Length     int    "json:\"length\""
				}{Validation: "standard", Length: 16}, Bin: struct {
					Pattern             string "json:\"pattern\""
					InstallmentsPattern string "json:\"installments_pattern\""
					ExclusionPattern    string "json:\"exclusion_pattern\""
				}{Pattern: "^(4)", InstallmentsPattern: "^(4)", ExclusionPattern: "^(488233|462896|484192|441509|486367|431385|454106|489635|455982|498476|492468|491268|459317|418253|404279|423949|457605|402739|450942|457604|457603|491511)"}, SecurityCode: struct {
					Length       int    "json:\"length\""
					CardLocation string "json:\"card_location\""
					Mode         string "json:\"mode\""
				}{Length: 3, CardLocation: "back", Mode: "mandatory"}}}, AdditionalInfoNeeded: []string{"cardholder_name", "cardholder_identification_number", "cardholder_identification_type"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, AccreditationTime: 2880, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
				{ID: "amex", Name: "American Express", PaymentTypeID: "credit_card", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/amex.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/amex.gif", DeferredCapture: "unsupported", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{{CardNumber: struct {
					Validation string "json:\"validation\""
					Length     int    "json:\"length\""
				}{Validation: "standard", Length: 15}, Bin: struct {
					Pattern             string "json:\"pattern\""
					InstallmentsPattern string "json:\"installments_pattern\""
					ExclusionPattern    string "json:\"exclusion_pattern\""
				}{Pattern: "^((34)|(37))", InstallmentsPattern: "^((34)|(37))", ExclusionPattern: ""}, SecurityCode: struct {
					Length       int    "json:\"length\""
					CardLocation string "json:\"card_location\""
					Mode         string "json:\"mode\""
				}{Length: 4, CardLocation: "front", Mode: "mandatory"}}}, AdditionalInfoNeeded: []string{"cardholder_name", "cardholder_identification_number", "cardholder_identification_type"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, AccreditationTime: 2880, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
				{ID: "master", Name: "Mastercard", PaymentTypeID: "credit_card", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/master.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/master.gif", DeferredCapture: "unsupported", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{{CardNumber: struct {
					Validation string "json:\"validation\""
					Length     int    "json:\"length\""
				}{Validation: "standard", Length: 16}, Bin: struct {
					Pattern             string "json:\"pattern\""
					InstallmentsPattern string "json:\"installments_pattern\""
					ExclusionPattern    string "json:\"exclusion_pattern\""
				}{Pattern: "^(5|(2(221|222|223|224|225|226|227|228|229|23|24|25|26|27|28|29|3|4|5|6|70|71|720)))", InstallmentsPattern: "^(5|(2(221|222|223|224|225|226|227|228|229|23|24|25|26|27|28|29|3|4|5|6|70|71|720)))", ExclusionPattern: "^(590712|529448|533254|534778|557555|525315|537980|523719|559749|530715|530716|530717|530710|530721|514332|529768|536782|552558|530729|530712|530724|530727|530695|510342|530719|530713|530720|530725|530726|530691|530711|530728|530714|530723|528633|524627|530718|530722|536126|518503|548185|525358|545409|535803|547692|520922|517404|516451|518092|517393|544039|515516|526557)"}, SecurityCode: struct {
					Length       int    "json:\"length\""
					CardLocation string "json:\"card_location\""
					Mode         string "json:\"mode\""
				}{Length: 3, CardLocation: "back", Mode: "mandatory"}}}, AdditionalInfoNeeded: []string{"cardholder_name", "cardholder_identification_number", "cardholder_identification_type"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, AccreditationTime: 2880, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
				{ID: "efecty", Name: "Efecty", PaymentTypeID: "ticket", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/efecty.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/efecty.gif", DeferredCapture: "does_not_apply", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{}, AdditionalInfoNeeded: []string{}, MinAllowedAmount: 5000, MaxAllowedAmount: 4000000, AccreditationTime: 0, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
				{ID: "pse", Name: "PSE", PaymentTypeID: "bank_transfer", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/pse.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/pse.gif", DeferredCapture: "does_not_apply", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{}, AdditionalInfoNeeded: []string{"entity_type"}, MinAllowedAmount: 1600, MaxAllowedAmount: 30000000, AccreditationTime: 30, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{{ID: "1040", Description: "Banco Agrario"}, {ID: "1507", Description: "NEQUI"}, {ID: "1052", Description: "Banco AV Villas"}, {ID: "1032", Description: "Banco Caja Social"}, {ID: "1019", Description: "SCOTIABANK COLPATRIA"}, {ID: "1066", Description: "Banco Cooperativo Coopcentral"}, {ID: "1051", Description: "Banco Davivienda"}, {ID: "1001", Description: "Banco De Bogota"}, {ID: "1023", Description: "Banco De Occidente"}, {ID: "1062", Description: "Banco Falabella"}, {ID: "1012", Description: "Banco GNB Sudameris"}, {ID: "1060", Description: "Banco Pichincha S.A."}, {ID: "1002", Description: "Banco Popular"}, {ID: "1058", Description: "Banco Procredit"}, {ID: "1007", Description: "Bancolombia"}, {ID: "1061", Description: "Bancoomeva S.A."}, {ID: "1013", Description: "Banco BBVA Colombia S.A."}, {ID: "1009", Description: "Citibank"}, {ID: "1006", Description: "Banco Itaú"}, {ID: "1292", Description: "Confiar Cooperativa Financiera"}, {ID: "1551", Description: "DaviPlata"}, {ID: "1069", Description: "BANCO SERFINANZA"}, {ID: "1065", Description: "BANCO SANTANDER COLOMBIA"}, {ID: "1283", Description: "CFA COOPERATIVA FINANCIERA"}, {ID: "1059", Description: "BANCO DE LAS MICROFINANZAS BANCAMIA"}, {ID: "1151", Description: "RAPPIPAY"}, {ID: "1289", Description: "Cotrafa"}, {ID: "1370", Description: "COLTEFINANCIERA"}}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
				{ID: "baloto", Name: "Baloto", PaymentTypeID: "ticket", Status: "active", SecureThumbnail: "https://http2.mlstatic.com/storage/logos-api-admin/bb3d20a0-99d5-11e9-a52b-295618a86fe2-xs.svg", Thumbnail: "https://http2.mlstatic.com/storage/logos-api-admin/bb3d20a0-99d5-11e9-a52b-295618a86fe2-xs.svg", DeferredCapture: "does_not_apply", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{}, AdditionalInfoNeeded: []string{}, MinAllowedAmount: 1500, MaxAllowedAmount: 1000000, AccreditationTime: 0, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string{"aggregator"}, PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}(nil), Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: false, Name: "", ID: 0}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 0, MerchantAccountID: interface{}(nil), Bins: []interface{}(nil), Marketplace: "", Agreements: []interface{}(nil), Labels: []string(nil), FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: ""}, SiteID: "", ProcessingMode: ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "/payment_methods", r.URL.Path)

				assert.Equal(t, tt.publicKey, r.URL.Query().Get("public_key"))
				if tt.accessToken != "" {
					assert.Equal(t, "Bearer "+tt.accessToken, r.Header.Get("Authorization"))
				}

				w.WriteHeader(tt.respStatus)
				_, _ = w.Write([]byte(tt.respBody))
			}))

			c := mercadopago.NewClient(server.URL, tt.publicKey, tt.accessToken)

			got, err := c.GetPaymentMethods()

			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expectedResponse, got)
		})
	}
}

func TestClient_GetPaymentMethodsForBin(t *testing.T) {
	tests := []struct {
		name             string
		bin              string
		expectedResponse mercadopago.PaymentMethods
		expectedErr      error
		respStatus       int
		respBody         string
	}{
		{
			name:       "successful response for mastercard bin",
			bin:        "525413",
			respStatus: http.StatusOK,
			respBody:   `{"paging":{"total":1,"limit":30,"offset":0},"results":[{"financial_institutions":[],"secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/master.gif","payer_costs":[{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":[],"installments":1,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":[],"installments":2,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":[],"installments":3,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":[],"installments":4,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":[],"installments":5,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":[],"installments":6,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":[],"installments":12,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":[],"installments":18,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":[],"installments":24,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":["recommended_installment"],"installments":36,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}],"issuer":{"default":true,"name":"Mastercard","id":204},"total_financial_cost":null,"min_accreditation_days":0,"max_accreditation_days":2,"merchant_account_id":null,"id":"master","payment_type_id":"credit_card","accreditation_time":2880,"thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/master.gif","bins":[],"marketplace":"NONE","deferred_capture":"unsupported","agreements":[],"labels":["recommended_method"],"financing_deals":{"legals":null,"installments":null,"expiration_date":null,"start_date":null,"status":"deactive"},"name":"Mastercard","site_id":"MCO","processing_mode":"aggregator","additional_info_needed":["cardholder_name","cardholder_identification_number","cardholder_identification_type"],"status":"active","settings":[{"security_code":{"mode":"mandatory","card_location":"back","length":3},"card_number":{"length":16,"validation":"standard"},"bin":{"pattern":"^(5|(2(221|222|223|224|225|226|227|228|229|23|24|25|26|27|28|29|3|4|5|6|70|71|720)))","installments_pattern":"^(5|(2(221|222|223|224|225|226|227|228|229|23|24|25|26|27|28|29|3|4|5|6|70|71|720)))","exclusion_pattern":"^(590712|529448|533254|534778|557555|525315|537980|523719|559749|530715|530716|530717|530710|530721|514332|529768|536782|552558|530729|530712|530724|530727|530695|510342|530719|530713|530720|530725|530726|530691|530711|530728|530714|530723|528633|524627|530718|530722|536126|518503|548185|525358|545409|535803|547692|520922|517404|516451|518092|517393|544039|515516|526557)"}}]}]}`,
			expectedResponse: mercadopago.PaymentMethods{mercadopago.PaymentMethod{ID: "master", Name: "Mastercard", PaymentTypeID: "credit_card", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/master.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/master.gif", DeferredCapture: "unsupported", Settings: []struct {
				CardNumber struct {
					Validation string "json:\"validation\""
					Length     int    "json:\"length\""
				} "json:\"card_number\""
				Bin struct {
					Pattern             string "json:\"pattern\""
					InstallmentsPattern string "json:\"installments_pattern\""
					ExclusionPattern    string "json:\"exclusion_pattern\""
				} "json:\"bin\""
				SecurityCode struct {
					Length       int    "json:\"length\""
					CardLocation string "json:\"card_location\""
					Mode         string "json:\"mode\""
				} "json:\"security_code\""
			}{{CardNumber: struct {
				Validation string "json:\"validation\""
				Length     int    "json:\"length\""
			}{Validation: "standard", Length: 16}, Bin: struct {
				Pattern             string "json:\"pattern\""
				InstallmentsPattern string "json:\"installments_pattern\""
				ExclusionPattern    string "json:\"exclusion_pattern\""
			}{Pattern: "^(5|(2(221|222|223|224|225|226|227|228|229|23|24|25|26|27|28|29|3|4|5|6|70|71|720)))", InstallmentsPattern: "^(5|(2(221|222|223|224|225|226|227|228|229|23|24|25|26|27|28|29|3|4|5|6|70|71|720)))", ExclusionPattern: "^(590712|529448|533254|534778|557555|525315|537980|523719|559749|530715|530716|530717|530710|530721|514332|529768|536782|552558|530729|530712|530724|530727|530695|510342|530719|530713|530720|530725|530726|530691|530711|530728|530714|530723|528633|524627|530718|530722|536126|518503|548185|525358|545409|535803|547692|520922|517404|516451|518092|517393|544039|515516|526557)"}, SecurityCode: struct {
				Length       int    "json:\"length\""
				CardLocation string "json:\"card_location\""
				Mode         string "json:\"mode\""
			}{Length: 3, CardLocation: "back", Mode: "mandatory"}}}, AdditionalInfoNeeded: []string{"cardholder_name", "cardholder_identification_number", "cardholder_identification_type"}, MinAllowedAmount: 0, MaxAllowedAmount: 0, AccreditationTime: 2880, FinancialInstitutions: []struct {
				ID          string "json:\"id\""
				Description string "json:\"description\""
			}{}, ProcessingModes: []string(nil), PayerCosts: []struct {
				InstallmentRate       int           "json:\"installment_rate\""
				DiscountRate          int           "json:\"discount_rate\""
				MinAllowedAmount      float64       "json:\"min_allowed_amount\""
				Labels                []interface{} "json:\"labels\""
				Installments          int           "json:\"installments\""
				ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
				MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
				PaymentMethodOptionID string        "json:\"payment_method_option_id\""
			}{{InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{}, Installments: 1, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{}, Installments: 2, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{}, Installments: 3, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{}, Installments: 4, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{}, Installments: 5, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{}, Installments: 6, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{}, Installments: 12, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{}, Installments: 18, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{}, Installments: 24, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{"recommended_installment"}, Installments: 36, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}}, Issuer: struct {
				Default bool   "json:\"default\""
				Name    string "json:\"name\""
				ID      int    "json:\"id\""
			}{Default: true, Name: "Mastercard", ID: 204}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 2, MerchantAccountID: interface{}(nil), Bins: []interface{}{}, Marketplace: "NONE", Agreements: []interface{}{}, Labels: []string{"recommended_method"}, FinancingDeals: struct {
				Legals         interface{} "json:\"legals\""
				Installments   interface{} "json:\"installments\""
				ExpirationDate interface{} "json:\"expiration_date\""
				StartDate      interface{} "json:\"start_date\""
				Status         string      "json:\"status\""
			}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: "deactive"}, SiteID: "MCO", ProcessingMode: "aggregator"}},
		},
		{
			name:       "successful response for visa bin",
			bin:        "401354",
			respStatus: http.StatusOK,
			respBody:   `{"paging":{"total":1,"limit":30,"offset":0},"results":[{"financial_institutions":[],"secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/visa.gif","payer_costs":[{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":[],"installments":1,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":[],"installments":2,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":[],"installments":3,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":[],"installments":4,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":[],"installments":5,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":[],"installments":6,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":[],"installments":12,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":[],"installments":18,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":[],"installments":24,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installment_rate":0,"discount_rate":0,"min_allowed_amount":1000,"labels":["recommended_installment"],"installments":36,"reimbursement_rate":null,"max_allowed_amount":5.0E7,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}],"issuer":{"default":true,"name":"Visa","id":205},"total_financial_cost":null,"min_accreditation_days":0,"max_accreditation_days":2,"merchant_account_id":null,"id":"visa","payment_type_id":"credit_card","accreditation_time":2880,"thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/visa.gif","bins":[],"marketplace":"NONE","deferred_capture":"unsupported","agreements":[],"labels":["recommended_method"],"financing_deals":{"legals":null,"installments":null,"expiration_date":null,"start_date":null,"status":"deactive"},"name":"Visa","site_id":"MCO","processing_mode":"aggregator","additional_info_needed":["cardholder_name","cardholder_identification_number","cardholder_identification_type"],"status":"active","settings":[{"security_code":{"mode":"mandatory","card_location":"back","length":3},"card_number":{"length":16,"validation":"standard"},"bin":{"pattern":"^(4)","installments_pattern":"^(4)","exclusion_pattern":"^(488233|462896|484192|441509|486367|431385|454106|489635|455982|498476|492468|491268|459317|418253|404279|423949|457605|402739|450942|457604|457603|491511)"}}]}]}`,
			expectedResponse: mercadopago.PaymentMethods{
				mercadopago.PaymentMethod{ID: "visa", Name: "Visa", PaymentTypeID: "credit_card", Status: "active", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/visa.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/visa.gif", DeferredCapture: "unsupported", Settings: []struct {
					CardNumber struct {
						Validation string "json:\"validation\""
						Length     int    "json:\"length\""
					} "json:\"card_number\""
					Bin struct {
						Pattern             string "json:\"pattern\""
						InstallmentsPattern string "json:\"installments_pattern\""
						ExclusionPattern    string "json:\"exclusion_pattern\""
					} "json:\"bin\""
					SecurityCode struct {
						Length       int    "json:\"length\""
						CardLocation string "json:\"card_location\""
						Mode         string "json:\"mode\""
					} "json:\"security_code\""
				}{{CardNumber: struct {
					Validation string "json:\"validation\""
					Length     int    "json:\"length\""
				}{Validation: "standard", Length: 16}, Bin: struct {
					Pattern             string "json:\"pattern\""
					InstallmentsPattern string "json:\"installments_pattern\""
					ExclusionPattern    string "json:\"exclusion_pattern\""
				}{Pattern: "^(4)", InstallmentsPattern: "^(4)", ExclusionPattern: "^(488233|462896|484192|441509|486367|431385|454106|489635|455982|498476|492468|491268|459317|418253|404279|423949|457605|402739|450942|457604|457603|491511)"}, SecurityCode: struct {
					Length       int    "json:\"length\""
					CardLocation string "json:\"card_location\""
					Mode         string "json:\"mode\""
				}{Length: 3, CardLocation: "back", Mode: "mandatory"}}}, AdditionalInfoNeeded: []string{"cardholder_name", "cardholder_identification_number", "cardholder_identification_type"}, MinAllowedAmount: 0, MaxAllowedAmount: 0, AccreditationTime: 2880, FinancialInstitutions: []struct {
					ID          string "json:\"id\""
					Description string "json:\"description\""
				}{}, ProcessingModes: []string(nil), PayerCosts: []struct {
					InstallmentRate       int           "json:\"installment_rate\""
					DiscountRate          int           "json:\"discount_rate\""
					MinAllowedAmount      float64       "json:\"min_allowed_amount\""
					Labels                []interface{} "json:\"labels\""
					Installments          int           "json:\"installments\""
					ReimbursementRate     interface{}   "json:\"reimbursement_rate\""
					MaxAllowedAmount      float64       "json:\"max_allowed_amount\""
					PaymentMethodOptionID string        "json:\"payment_method_option_id\""
				}{{InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{}, Installments: 1, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{}, Installments: 2, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{}, Installments: 3, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{}, Installments: 4, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{}, Installments: 5, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{}, Installments: 6, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{}, Installments: 12, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{}, Installments: 18, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{}, Installments: 24, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {InstallmentRate: 0, DiscountRate: 0, MinAllowedAmount: 1000, Labels: []interface{}{"recommended_installment"}, Installments: 36, ReimbursementRate: interface{}(nil), MaxAllowedAmount: 5e+07, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}}, Issuer: struct {
					Default bool   "json:\"default\""
					Name    string "json:\"name\""
					ID      int    "json:\"id\""
				}{Default: true, Name: "Visa", ID: 205}, TotalFinancialCost: interface{}(nil), MinAccreditationDays: 0, MaxAccreditationDays: 2, MerchantAccountID: interface{}(nil), Bins: []interface{}{}, Marketplace: "NONE", Agreements: []interface{}{}, Labels: []string{"recommended_method"}, FinancingDeals: struct {
					Legals         interface{} "json:\"legals\""
					Installments   interface{} "json:\"installments\""
					ExpirationDate interface{} "json:\"expiration_date\""
					StartDate      interface{} "json:\"start_date\""
					Status         string      "json:\"status\""
				}{Legals: interface{}(nil), Installments: interface{}(nil), ExpirationDate: interface{}(nil), StartDate: interface{}(nil), Status: "deactive"}, SiteID: "MCO", ProcessingMode: "aggregator"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			publicKey := "public-key"

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "/payment_methods/search", r.URL.Path)

				assert.Equal(t, publicKey, r.URL.Query().Get("public_key"))
				assert.Equal(t, "NONE", r.URL.Query().Get("marketplace"))
				assert.Equal(t, "active", r.URL.Query().Get("status"))

				w.WriteHeader(tt.respStatus)
				_, _ = w.Write([]byte(tt.respBody))
			}))

			c := mercadopago.NewClient(server.URL, publicKey, "")

			got, err := c.GetPaymentMethodsForBin(tt.bin)

			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expectedResponse, got)
		})
	}
}

func TestClient_GetInstallments(t *testing.T) {
	tests := []struct {
		name             string
		params           mercadopago.GetInstallmentsParams
		expectedResponse mercadopago.Installments
		expectedErr      error
		respStatus       int
		respBody         string
	}{
		{
			name: "successful response",
			params: mercadopago.GetInstallmentsParams{
				PaymentMethodID: "amex",
				Amount:          1234.56,
				IssuerID:        "987",
			},
			respStatus: http.StatusOK,
			respBody:   `[{"payment_method_id":"amex","payment_type_id":"credit_card","issuer":{"id":"551","name":"Bancolombia","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/amex.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/amex.gif"},"processing_mode":"aggregator","merchant_account_id":null,"payer_costs":[{"installments":1,"installment_rate":0,"discount_rate":0,"reimbursement_rate":null,"labels":[],"installment_rate_collector":["MERCADOPAGO"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"recommended_message":"1 cuota de $ 2.000 ($ 2.000)","installment_amount":2000,"total_amount":2000,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installments":2,"installment_rate":0,"discount_rate":0,"reimbursement_rate":null,"labels":[],"installment_rate_collector":["THIRD_PARTY"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"recommended_message":"2 cuotas de $ 1.000 ($ 2.000)","installment_amount":1000,"total_amount":2000,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installments":3,"installment_rate":0,"discount_rate":0,"reimbursement_rate":null,"labels":[],"installment_rate_collector":["THIRD_PARTY"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"recommended_message":"3 cuotas de $ 667 ($ 2.000)","installment_amount":666.67,"total_amount":2000,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installments":4,"installment_rate":0,"discount_rate":0,"reimbursement_rate":null,"labels":[],"installment_rate_collector":["THIRD_PARTY"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"recommended_message":"4 cuotas de $ 500 ($ 2.000)","installment_amount":500,"total_amount":2000,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installments":5,"installment_rate":0,"discount_rate":0,"reimbursement_rate":null,"labels":[],"installment_rate_collector":["THIRD_PARTY"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"recommended_message":"5 cuotas de $ 400 ($ 2.000)","installment_amount":400,"total_amount":2000,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installments":6,"installment_rate":0,"discount_rate":0,"reimbursement_rate":null,"labels":[],"installment_rate_collector":["THIRD_PARTY"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"recommended_message":"6 cuotas de $ 333 ($ 2.000)","installment_amount":333.33,"total_amount":2000,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installments":12,"installment_rate":0,"discount_rate":0,"reimbursement_rate":null,"labels":[],"installment_rate_collector":["THIRD_PARTY"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"recommended_message":"12 cuotas de $ 167 ($ 2.000)","installment_amount":166.67,"total_amount":2000,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installments":18,"installment_rate":0,"discount_rate":0,"reimbursement_rate":null,"labels":[],"installment_rate_collector":["THIRD_PARTY"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"recommended_message":"18 cuotas de $ 111 ($ 2.000)","installment_amount":111.11,"total_amount":2000,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installments":24,"installment_rate":0,"discount_rate":0,"reimbursement_rate":null,"labels":[],"installment_rate_collector":["THIRD_PARTY"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"recommended_message":"24 cuotas de $ 83 ($ 2.000)","installment_amount":83.33,"total_amount":2000,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"},{"installments":36,"installment_rate":0,"discount_rate":0,"reimbursement_rate":null,"labels":["recommended_installment"],"installment_rate_collector":["THIRD_PARTY"],"min_allowed_amount":1000,"max_allowed_amount":50000000,"recommended_message":"36 cuotas de $ 56 ($ 2.000)","installment_amount":55.56,"total_amount":2000,"payment_method_option_id":"1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}],"agreements":null}]`,
			expectedResponse: mercadopago.Installments{mercadopago.Installment{PaymentMethodID: "amex", PaymentTypeID: "credit_card", Issuer: struct {
				ID              string "json:\"id\""
				Name            string "json:\"name\""
				SecureThumbnail string "json:\"secure_thumbnail\""
				Thumbnail       string "json:\"thumbnail\""
			}{ID: "551", Name: "Bancolombia", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/amex.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/amex.gif"}, ProcessingMode: "aggregator", MerchantAccountID: interface{}(nil), PayerCosts: []struct {
				Installments             int           "json:\"installments\""
				InstallmentRate          int           "json:\"installment_rate\""
				DiscountRate             int           "json:\"discount_rate\""
				ReimbursementRate        interface{}   "json:\"reimbursement_rate\""
				Labels                   []interface{} "json:\"labels\""
				InstallmentRateCollector []string      "json:\"installment_rate_collector\""
				MinAllowedAmount         int           "json:\"min_allowed_amount\""
				MaxAllowedAmount         int           "json:\"max_allowed_amount\""
				RecommendedMessage       string        "json:\"recommended_message\""
				InstallmentAmount        float64       "json:\"installment_amount\""
				TotalAmount              int           "json:\"total_amount\""
				PaymentMethodOptionID    string        "json:\"payment_method_option_id\""
			}{{Installments: 1, InstallmentRate: 0, DiscountRate: 0, ReimbursementRate: interface{}(nil), Labels: []interface{}{}, InstallmentRateCollector: []string{"MERCADOPAGO"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, RecommendedMessage: "1 cuota de $ 2.000 ($ 2.000)", InstallmentAmount: 2000, TotalAmount: 2000, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {Installments: 2, InstallmentRate: 0, DiscountRate: 0, ReimbursementRate: interface{}(nil), Labels: []interface{}{}, InstallmentRateCollector: []string{"THIRD_PARTY"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, RecommendedMessage: "2 cuotas de $ 1.000 ($ 2.000)", InstallmentAmount: 1000, TotalAmount: 2000, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {Installments: 3, InstallmentRate: 0, DiscountRate: 0, ReimbursementRate: interface{}(nil), Labels: []interface{}{}, InstallmentRateCollector: []string{"THIRD_PARTY"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, RecommendedMessage: "3 cuotas de $ 667 ($ 2.000)", InstallmentAmount: 666.67, TotalAmount: 2000, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {Installments: 4, InstallmentRate: 0, DiscountRate: 0, ReimbursementRate: interface{}(nil), Labels: []interface{}{}, InstallmentRateCollector: []string{"THIRD_PARTY"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, RecommendedMessage: "4 cuotas de $ 500 ($ 2.000)", InstallmentAmount: 500, TotalAmount: 2000, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {Installments: 5, InstallmentRate: 0, DiscountRate: 0, ReimbursementRate: interface{}(nil), Labels: []interface{}{}, InstallmentRateCollector: []string{"THIRD_PARTY"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, RecommendedMessage: "5 cuotas de $ 400 ($ 2.000)", InstallmentAmount: 400, TotalAmount: 2000, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {Installments: 6, InstallmentRate: 0, DiscountRate: 0, ReimbursementRate: interface{}(nil), Labels: []interface{}{}, InstallmentRateCollector: []string{"THIRD_PARTY"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, RecommendedMessage: "6 cuotas de $ 333 ($ 2.000)", InstallmentAmount: 333.33, TotalAmount: 2000, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {Installments: 12, InstallmentRate: 0, DiscountRate: 0, ReimbursementRate: interface{}(nil), Labels: []interface{}{}, InstallmentRateCollector: []string{"THIRD_PARTY"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, RecommendedMessage: "12 cuotas de $ 167 ($ 2.000)", InstallmentAmount: 166.67, TotalAmount: 2000, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {Installments: 18, InstallmentRate: 0, DiscountRate: 0, ReimbursementRate: interface{}(nil), Labels: []interface{}{}, InstallmentRateCollector: []string{"THIRD_PARTY"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, RecommendedMessage: "18 cuotas de $ 111 ($ 2.000)", InstallmentAmount: 111.11, TotalAmount: 2000, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {Installments: 24, InstallmentRate: 0, DiscountRate: 0, ReimbursementRate: interface{}(nil), Labels: []interface{}{}, InstallmentRateCollector: []string{"THIRD_PARTY"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, RecommendedMessage: "24 cuotas de $ 83 ($ 2.000)", InstallmentAmount: 83.33, TotalAmount: 2000, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}, {Installments: 36, InstallmentRate: 0, DiscountRate: 0, ReimbursementRate: interface{}(nil), Labels: []interface{}{"recommended_installment"}, InstallmentRateCollector: []string{"THIRD_PARTY"}, MinAllowedAmount: 1000, MaxAllowedAmount: 50000000, RecommendedMessage: "36 cuotas de $ 56 ($ 2.000)", InstallmentAmount: 55.56, TotalAmount: 2000, PaymentMethodOptionID: "1.AQokODllZjQyNjktYjAzMy00OWU1LWJhMWQtNDE0NjQyNTM3MzY4EJaFuevHLg"}}, Agreements: interface{}(nil)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			publicKey := "public-key"

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "/payment_methods/installments", r.URL.Path)

				assert.Equal(t, publicKey, r.URL.Query().Get("public_key"))

				if tt.params.PaymentMethodID != "" {
					assert.Equal(t, tt.params.PaymentMethodID, r.URL.Query().Get("payment_method_id"))
				}

				if tt.params.Amount != 0 {
					assert.Equal(t, strconv.FormatFloat(tt.params.Amount, 'f', 6, 64), r.URL.Query().Get("amount"))
				}

				if tt.params.IssuerID != "" {
					assert.Equal(t, tt.params.IssuerID, r.URL.Query().Get("issuer.id"))
				}

				w.WriteHeader(tt.respStatus)
				_, _ = w.Write([]byte(tt.respBody))
			}))

			c := mercadopago.NewClient(server.URL, publicKey, "")

			got, err := c.GetInstallments(tt.params)

			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expectedResponse, got)
		})
	}
}

func TestClient_GetCardIssuers(t *testing.T) {
	tests := []struct {
		name             string
		paymentMethodID  string
		expectedResponse mercadopago.Issuers
		expectedErr      error
		respStatus       int
		respBody         string
	}{
		{
			name:             "successful response",
			paymentMethodID:  "amex",
			respStatus:       http.StatusOK,
			respBody:         `[{"id":"551","name":"Bancolombia","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/amex.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/amex.gif","processing_mode":"aggregator","merchant_account_id":null},{"id":"12467","name":"Davivienda","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/amex.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/amex.gif","processing_mode":"aggregator","merchant_account_id":null},{"id":"12475","name":"Scotiabank Colpatria","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/amex.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/amex.gif","processing_mode":"aggregator","merchant_account_id":null},{"id":"202","name":"Otro","secure_thumbnail":"https://www.mercadopago.com/org-img/MP3/API/logos/amex.gif","thumbnail":"http://img.mlstatic.com/org-img/MP3/API/logos/amex.gif","processing_mode":"aggregator","merchant_account_id":null}]`,
			expectedResponse: mercadopago.Issuers{mercadopago.Issuer{ID: "551", Name: "Bancolombia", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/amex.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/amex.gif", ProcessingMode: "aggregator", MerchantAccountID: interface{}(nil)}, mercadopago.Issuer{ID: "12467", Name: "Davivienda", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/amex.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/amex.gif", ProcessingMode: "aggregator", MerchantAccountID: interface{}(nil)}, mercadopago.Issuer{ID: "12475", Name: "Scotiabank Colpatria", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/amex.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/amex.gif", ProcessingMode: "aggregator", MerchantAccountID: interface{}(nil)}, mercadopago.Issuer{ID: "202", Name: "Otro", SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/amex.gif", Thumbnail: "http://img.mlstatic.com/org-img/MP3/API/logos/amex.gif", ProcessingMode: "aggregator", MerchantAccountID: interface{}(nil)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			publicKey := "public-key"

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "/payment_methods/card_issuers", r.URL.Path)

				assert.Equal(t, publicKey, r.URL.Query().Get("public_key"))
				assert.Equal(t, tt.paymentMethodID, r.URL.Query().Get("payment_method_id"))

				w.WriteHeader(tt.respStatus)
				_, _ = w.Write([]byte(tt.respBody))
			}))

			c := mercadopago.NewClient(server.URL, publicKey, "")

			got, err := c.GetCardIssuers(tt.paymentMethodID)

			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expectedResponse, got)
		})
	}
}
