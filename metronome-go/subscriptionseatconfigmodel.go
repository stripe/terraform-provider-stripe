//
// File generated from our OpenAPI spec
//

package metronome

type SubscriptionSeatConfigModel struct {
	APIResource
	// The property name, sent on usage events, that identifies the seat ID associated with the usage event.  For example, the property name might be seat_id or user_id. The property must be set as a group key on billable metrics and a presentation/pricing group key on contract products.  This allows linked recurring credits with an allocation per seat to be consumed by only one seat's usage.
	SeatGroupKey string `json:"seat_group_key"`
}
