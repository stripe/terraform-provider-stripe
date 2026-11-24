//
// File generated from our OpenAPI spec
//

package metronome

type V1ListCreditTypesResponseData struct {
	ID         string `json:"id,omitempty"`
	IsCurrency bool   `json:"is_currency,omitempty"`
	Name       string `json:"name,omitempty"`
}
type V1ListCreditTypesResponse struct {
	APIResource
	Data     []*V1ListCreditTypesResponseData `json:"data"`
	NextPage string                           `json:"next_page"`
}
