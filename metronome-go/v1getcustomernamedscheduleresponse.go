//
// File generated from our OpenAPI spec
//

package metronome

type V1GetCustomerNamedScheduleResponseData struct {
	EndingBefore string `json:"ending_before,omitempty"`
	StartingAt   string `json:"starting_at"`
	Value        JSON   `json:"value"`
}
type V1GetCustomerNamedScheduleResponse struct {
	APIResource
	Data []*V1GetCustomerNamedScheduleResponseData `json:"data"`
}
