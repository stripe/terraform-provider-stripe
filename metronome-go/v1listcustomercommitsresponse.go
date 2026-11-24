//
// File generated from our OpenAPI spec
//

package metronome

type V1ListCustomerCommitsResponse struct {
	APIResource
	Data     []*CommitModel `json:"data"`
	NextPage string         `json:"next_page"`
}
