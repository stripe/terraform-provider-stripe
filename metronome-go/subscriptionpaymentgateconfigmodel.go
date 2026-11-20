//
// File generated from our OpenAPI spec
//

package metronome

// Successful Stripe payment gates access to the subscription.
type SubscriptionPaymentGateConfigModelPaymentGateType string

// List of values that SubscriptionPaymentGateConfigModelPaymentGateType can take
const (
	SubscriptionPaymentGateConfigModelPaymentGateTypeSTRIPE SubscriptionPaymentGateConfigModelPaymentGateType = "STRIPE"
)

// Determines whether Stripe calculates tax for the payment-gated subscription.
type SubscriptionPaymentGateConfigModelTaxType string

// List of values that SubscriptionPaymentGateConfigModelTaxType can take
const (
	SubscriptionPaymentGateConfigModelTaxTypeNONE   SubscriptionPaymentGateConfigModelTaxType = "NONE"
	SubscriptionPaymentGateConfigModelTaxTypeSTRIPE SubscriptionPaymentGateConfigModelTaxType = "STRIPE"
)

type SubscriptionPaymentGateConfigModel struct {
	APIResource
	// Successful Stripe payment gates access to the subscription.
	PaymentGateType SubscriptionPaymentGateConfigModelPaymentGateType `json:"payment_gate_type"`
	// Determines whether Stripe calculates tax for the payment-gated subscription.
	TaxType SubscriptionPaymentGateConfigModelTaxType `json:"tax_type,omitempty"`
}
