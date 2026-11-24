//
// File generated from our OpenAPI spec
//

package metronome

type SubscriptionSeatUpdateInputModel struct {
	APIResource
	// Adds seat IDs to the subscription.  If there are unassigned seats, the new seat IDs will fill these unassigned seats and not increase the total subscription quantity. Otherwise, if there are more new seat IDs than unassigned seats, the total subscription quantity will increase.
	AddSeatIDs []*SubscriptionSeatUpdateAssignedSeatModel `json:"add_seat_ids,omitempty"`
	// Adds unassigned seats to the subscription. This will increase the total subscription quantity.
	AddUnassignedSeats []*SubscriptionSeatUpdateUnassignedSeatModel `json:"add_unassigned_seats,omitempty"`
	// Removes seat IDs from the subscription, if possible.  If a seat ID is removed, the total subscription quantity will decrease. Otherwise, if the seat ID is not found on the subscription, this is a no-op.
	RemoveSeatIDs []*SubscriptionSeatUpdateAssignedSeatModel `json:"remove_seat_ids,omitempty"`
	// Removes unassigned seats from the subscription. This will decrease the total subscription quantity if there are are unassigned seats.
	RemoveUnassignedSeats []*SubscriptionSeatUpdateUnassignedSeatModel `json:"remove_unassigned_seats,omitempty"`
}
