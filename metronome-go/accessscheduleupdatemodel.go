//
// File generated from our OpenAPI spec
//

package metronome

type AccessScheduleUpdateModelAddScheduleItem struct {
	Amount Number `json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore string `json:"ending_before"`
	// RFC 3339 timestamp (inclusive)
	StartingAt string `json:"starting_at"`
}
type AccessScheduleUpdateModelUpdateScheduleItem struct {
	Amount Number `json:"amount,omitempty"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore string `json:"ending_before,omitempty"`
	ID           string `json:"id"`
	// RFC 3339 timestamp (inclusive)
	StartingAt string `json:"starting_at,omitempty"`
}
type AccessScheduleUpdateModelRemoveScheduleItem struct {
	ID string `json:"id"`
}
type AccessScheduleUpdateModel struct {
	APIResource
	AddScheduleItems    []*AccessScheduleUpdateModelAddScheduleItem    `json:"add_schedule_items,omitempty"`
	RemoveScheduleItems []*AccessScheduleUpdateModelRemoveScheduleItem `json:"remove_schedule_items,omitempty"`
	UpdateScheduleItems []*AccessScheduleUpdateModelUpdateScheduleItem `json:"update_schedule_items,omitempty"`
}
