//
// File generated from our OpenAPI spec
//

package metronome

type RecurringCommitSubscriptionConfigModelAllocation string

// List of values that RecurringCommitSubscriptionConfigModelAllocation can take
const (
	RecurringCommitSubscriptionConfigModelAllocationINDIVIDUAL RecurringCommitSubscriptionConfigModelAllocation = "INDIVIDUAL"
	RecurringCommitSubscriptionConfigModelAllocationPOOLED     RecurringCommitSubscriptionConfigModelAllocation = "POOLED"
)

// Controls when a customer can access recurring commit or credit child balances when the attached subscription is payment-gated. `BILLING_PERIOD_PAID` releases each balance after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases all balances after the first payment.
type RecurringCommitSubscriptionConfigModelAccessPolicy string

// List of values that RecurringCommitSubscriptionConfigModelAccessPolicy can take
const (
	RecurringCommitSubscriptionConfigModelAccessPolicyBILLINGPERIODPAID            RecurringCommitSubscriptionConfigModelAccessPolicy = "BILLING_PERIOD_PAID"
	RecurringCommitSubscriptionConfigModelAccessPolicyINITIALBILLINGPERIODPAIDONLY RecurringCommitSubscriptionConfigModelAccessPolicy = "INITIAL_BILLING_PERIOD_PAID_ONLY"
)

type RecurringCommitSubscriptionConfigModel struct {
	APIResource
	// Controls when a customer can access recurring commit or credit child balances when the attached subscription is payment-gated. `BILLING_PERIOD_PAID` releases each balance after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases all balances after the first payment.
	AccessPolicy            RecurringCommitSubscriptionConfigModelAccessPolicy `json:"access_policy,omitempty"`
	Allocation              RecurringCommitSubscriptionConfigModelAllocation   `json:"allocation"`
	ApplySeatIncreaseConfig *ApplySeatIncreaseConfigForRecurringCommitModel    `json:"apply_seat_increase_config"`
	SubscriptionID          string                                             `json:"subscription_id"`
}
