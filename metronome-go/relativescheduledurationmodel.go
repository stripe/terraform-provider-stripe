//
// File generated from our OpenAPI spec
//

package metronome

// Indicates how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
type RelativeScheduleDurationModelAccessType string

// List of values that RelativeScheduleDurationModelAccessType can take
const (
	RelativeScheduleDurationModelAccessTypeSPEND    RelativeScheduleDurationModelAccessType = "SPEND"
	RelativeScheduleDurationModelAccessTypeQUANTITY RelativeScheduleDurationModelAccessType = "QUANTITY"
)

type RelativeScheduleDurationModelScheduleItem struct {
	Amount           Number             `json:"amount"`
	Duration         *RelativeDateModel `json:"duration"`
	ID               string             `json:"id"`
	StartingAtOffset *RelativeDateModel `json:"starting_at_offset"`
}
type RelativeScheduleDurationModel struct {
	APIResource
	// Indicates how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType    RelativeScheduleDurationModelAccessType      `json:"access_type,omitempty"`
	CreditType    *CreditTypeModel                             `json:"credit_type"`
	ScheduleItems []*RelativeScheduleDurationModelScheduleItem `json:"schedule_items"`
}
