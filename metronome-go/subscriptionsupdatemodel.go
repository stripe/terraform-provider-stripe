//
// File generated from our OpenAPI spec
//

package metronome

type SubscriptionsUpdateModelQuantityUpdate struct {
	Quantity      Number `json:"quantity,omitempty"`
	QuantityDelta Number `json:"quantity_delta,omitempty"`
	StartingAt    string `json:"starting_at"`
}
type SubscriptionsUpdateModel struct {
	APIResource
	EndingBefore    string                                    `json:"ending_before,omitempty"`
	ID              string                                    `json:"id"`
	Name            string                                    `json:"name,omitempty"`
	QuantityUpdates []*SubscriptionsUpdateModelQuantityUpdate `json:"quantity_updates,omitempty"`
	// Manage subscription seats for subscriptions in SEAT_BASED mode.
	SeatUpdates *SubscriptionSeatUpdateInputModel `json:"seat_updates,omitempty"`
}
