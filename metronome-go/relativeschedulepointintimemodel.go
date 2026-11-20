//
// File generated from our OpenAPI spec
//

package metronome

type RelativeSchedulePointInTimeModelScheduleItem struct {
	DateOffset *RelativeDateModel `json:"date_offset"`
	ID         string             `json:"id"`
	Quantity   Number             `json:"quantity"`
	UnitPrice  Number             `json:"unit_price"`
}
type RelativeSchedulePointInTimeModel struct {
	APIResource
	CreditType    *CreditTypeModel                                `json:"credit_type"`
	ScheduleItems []*RelativeSchedulePointInTimeModelScheduleItem `json:"schedule_items"`
}
