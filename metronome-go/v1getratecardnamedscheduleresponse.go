//
// File generated from our OpenAPI spec
//

package metronome

type V1GetRateCardNamedScheduleResponseData struct {
	EndingBefore string `json:"ending_before,omitempty"`
	StartingAt   string `json:"starting_at"`
	Value        JSON   `json:"value"`
}
type V1GetRateCardNamedScheduleResponse struct {
	APIResource
	Data []*V1GetRateCardNamedScheduleResponseData `json:"data"`
}
