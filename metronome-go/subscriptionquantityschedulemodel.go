//
// File generated from our OpenAPI spec
//

package metronome

type SubscriptionQuantityScheduleModel struct {
	APIResource
	EndingBefore string `json:"ending_before,omitempty"`
	Quantity     Number `json:"quantity"`
	StartingAt   string `json:"starting_at"`
}
