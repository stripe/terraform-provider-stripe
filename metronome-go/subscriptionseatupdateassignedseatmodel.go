//
// File generated from our OpenAPI spec
//

package metronome

type SubscriptionSeatUpdateAssignedSeatModel struct {
	APIResource
	SeatIDs []string `json:"seat_ids"`
	// Assigned seats will be added/removed starting at this date.
	StartingAt string `json:"starting_at"`
}
