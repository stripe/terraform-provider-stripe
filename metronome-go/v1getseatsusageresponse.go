//
// File generated from our OpenAPI spec
//

package metronome

type V1GetSeatsUsageResponse struct {
	APIResource
	Data     []*SeatsUsageAggregateModel `json:"data"`
	NextPage string                      `json:"next_page"`
}
