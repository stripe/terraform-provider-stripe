//
// File generated from our OpenAPI spec
//

package metronome

type V1ListCustomerBalancesResponse struct {
	APIResource
	Data     []JSON `json:"data"`
	NextPage string `json:"next_page"`
}
