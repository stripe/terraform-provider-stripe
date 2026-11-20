//
// File generated from our OpenAPI spec
//

package metronome

type V1GetContractRateScheduleResponse struct {
	APIResource
	Data     []*ContractRateScheduleModel `json:"data"`
	NextPage string                       `json:"next_page,omitempty"`
}
