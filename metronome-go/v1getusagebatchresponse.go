//
// File generated from our OpenAPI spec
//

package metronome

type V1GetUsageBatchResponse struct {
	APIResource
	Data     []*UsageBatchAggregateModel `json:"data"`
	NextPage string                      `json:"next_page"`
}
