//
// File generated from our OpenAPI spec
//

package metronome

type CustomerContractNamedScheduleModelSegment struct {
	EndingBefore string `json:"ending_before,omitempty"`
	StartingAt   string `json:"starting_at"`
	Value        JSON   `json:"value"`
}
type CustomerContractNamedScheduleModel struct {
	APIResource
	// ID of the contract this named schedule belongs to.
	ContractID string `json:"contract_id"`
	// Properties on this named schedule.
	Properties map[string]string `json:"properties,omitempty"`
	// Name of the schedule.
	ScheduleName string                                       `json:"schedule_name"`
	Segments     []*CustomerContractNamedScheduleModelSegment `json:"segments"`
}
