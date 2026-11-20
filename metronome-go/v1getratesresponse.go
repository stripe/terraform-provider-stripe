//
// File generated from our OpenAPI spec
//

package metronome

type V1GetRatesResponse struct {
	APIResource
	Data     []*RateScheduleModel `json:"data"`
	NextPage string               `json:"next_page,omitempty"`
}
