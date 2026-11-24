//
// File generated from our OpenAPI spec
//

package metronome

type V1ListCustomersResponse struct {
	APIResource
	Data     []*CustomerDetailModel `json:"data"`
	NextPage string                 `json:"next_page"`
}
