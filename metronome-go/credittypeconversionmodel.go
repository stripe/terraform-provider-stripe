//
// File generated from our OpenAPI spec
//

package metronome

type CreditTypeConversionModel struct {
	APIResource
	CustomCreditType    *CreditTypeModel `json:"custom_credit_type"`
	FiatPerCustomCredit string           `json:"fiat_per_custom_credit"`
}
