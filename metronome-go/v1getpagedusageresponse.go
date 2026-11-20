//
// File generated from our OpenAPI spec
//

package metronome

type V1GetPagedUsageResponse struct {
	APIResource
	Data     []*PagedUsageAggregateModel `json:"data"`
	NextPage string                      `json:"next_page"`
}
