//
// File generated from our OpenAPI spec
//

package metronome

type SubscriptionSeatUpdateUnassignedSeatModel struct {
	APIResource
	// The number of unassigned seats on the subscription will increase/decrease by this delta. Must be greater than 0.
	Quantity Number `json:"quantity"`
	// Unassigned seats will be updated starting at this date.
	StartingAt string `json:"starting_at"`
}
