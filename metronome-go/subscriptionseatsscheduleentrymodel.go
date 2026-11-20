//
// File generated from our OpenAPI spec
//

package metronome

type SubscriptionSeatsScheduleEntryModel struct {
	APIResource
	// Array of seat IDs that are assigned in this period
	AssignedSeatIDs []string `json:"assigned_seat_ids"`
	// The end time of this seat schedule period (null if ongoing)
	EndingBefore string `json:"ending_before"`
	// The start time of this seat schedule period
	StartingAt string `json:"starting_at"`
	// Total number of assigned and unassigned seats in this period
	TotalQuantity int64 `json:"total_quantity"`
}
