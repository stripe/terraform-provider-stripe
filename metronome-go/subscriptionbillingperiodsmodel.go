//
// File generated from our OpenAPI spec
//

package metronome

// Previous, current, and next billing periods for the subscription.
type SubscriptionBillingPeriodsModel struct {
	APIResource
	Current  *SubscriptionBillingPeriodModel `json:"current,omitempty"`
	Next     *SubscriptionBillingPeriodModel `json:"next,omitempty"`
	Previous *SubscriptionBillingPeriodModel `json:"previous,omitempty"`
}
