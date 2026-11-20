//
// File generated from our OpenAPI spec
//

package metronome

type RateCardModel struct {
	APIResource
	Aliases               []*RateCardAliasModel        `json:"aliases,omitempty"`
	CreatedAt             string                       `json:"created_at"`
	CreatedBy             string                       `json:"created_by"`
	CreditTypeConversions []*CreditTypeConversionModel `json:"credit_type_conversions,omitempty"`
	CustomFields          map[string]string            `json:"custom_fields,omitempty"`
	Description           string                       `json:"description,omitempty"`
	FiatCreditType        *CreditTypeModel             `json:"fiat_credit_type,omitempty"`
	ID                    string                       `json:"id"`
	Name                  string                       `json:"name"`
}
