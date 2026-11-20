//
// File generated from our OpenAPI spec
//

package metronome

type V1ListCustomerCreditsResponse struct {
	APIResource
	Data     []*CreditModel `json:"data"`
	NextPage string         `json:"next_page"`
}
