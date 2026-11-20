//
// File generated from our OpenAPI spec
//

package metronome

// Indicates how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
type ScheduleDurationModelAccessType string

// List of values that ScheduleDurationModelAccessType can take
const (
	ScheduleDurationModelAccessTypeSPEND    ScheduleDurationModelAccessType = "SPEND"
	ScheduleDurationModelAccessTypeQUANTITY ScheduleDurationModelAccessType = "QUANTITY"
)

type ScheduleDurationModelScheduleItem struct {
	Amount       Number `json:"amount"`
	EndingBefore string `json:"ending_before"`
	ID           string `json:"id"`
	StartingAt   string `json:"starting_at"`
}
type ScheduleDurationModel struct {
	APIResource
	// Indicates how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType    ScheduleDurationModelAccessType      `json:"access_type,omitempty"`
	CreditType    *CreditTypeModel                     `json:"credit_type,omitempty"`
	ScheduleItems []*ScheduleDurationModelScheduleItem `json:"schedule_items"`
}
