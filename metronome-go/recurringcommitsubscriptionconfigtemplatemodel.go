//
// File generated from our OpenAPI spec
//

package metronome

type RecurringCommitSubscriptionConfigTemplateModelAllocation string

// List of values that RecurringCommitSubscriptionConfigTemplateModelAllocation can take
const (
	RecurringCommitSubscriptionConfigTemplateModelAllocationINDIVIDUAL RecurringCommitSubscriptionConfigTemplateModelAllocation = "INDIVIDUAL"
	RecurringCommitSubscriptionConfigTemplateModelAllocationPOOLED     RecurringCommitSubscriptionConfigTemplateModelAllocation = "POOLED"
)

// Controls when a customer can access recurring commit or credit child balances when the subscription is payment-gated. `BILLING_PERIOD_PAID` releases each balance after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases all balances after the first payment.
type RecurringCommitSubscriptionConfigTemplateModelAccessPolicy string

// List of values that RecurringCommitSubscriptionConfigTemplateModelAccessPolicy can take
const (
	RecurringCommitSubscriptionConfigTemplateModelAccessPolicyBILLINGPERIODPAID            RecurringCommitSubscriptionConfigTemplateModelAccessPolicy = "BILLING_PERIOD_PAID"
	RecurringCommitSubscriptionConfigTemplateModelAccessPolicyINITIALBILLINGPERIODPAIDONLY RecurringCommitSubscriptionConfigTemplateModelAccessPolicy = "INITIAL_BILLING_PERIOD_PAID_ONLY"
)

type RecurringCommitSubscriptionConfigTemplateModel struct {
	APIResource
	// Controls when a customer can access recurring commit or credit child balances when the subscription is payment-gated. `BILLING_PERIOD_PAID` releases each balance after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases all balances after the first payment.
	AccessPolicy            RecurringCommitSubscriptionConfigTemplateModelAccessPolicy `json:"access_policy,omitempty"`
	Allocation              RecurringCommitSubscriptionConfigTemplateModelAllocation   `json:"allocation"`
	ApplySeatIncreaseConfig *ApplySeatIncreaseConfigForRecurringCommitModel            `json:"apply_seat_increase_config"`
	SubscriptionTemplateID  string                                                     `json:"subscription_template_id"`
}
