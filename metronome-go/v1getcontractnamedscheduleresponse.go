//
// File generated from our OpenAPI spec
//

package metronome

type V1GetContractNamedScheduleResponseData struct {
	EndingBefore string `json:"ending_before,omitempty"`
	StartingAt   string `json:"starting_at"`
	Value        JSON   `json:"value"`
}
type V1GetContractNamedScheduleResponse struct {
	APIResource
	Data []*V1GetContractNamedScheduleResponseData `json:"data"`
}
