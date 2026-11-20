//
// File generated from our OpenAPI spec
//

package metronome

// Controls whether this subscription consolidates onto usage invoices or gets its own scheduled invoice.
type SubscriptionBillingCycleConfigModelInvoicePlacement string

// List of values that SubscriptionBillingCycleConfigModelInvoicePlacement can take
const (
	SubscriptionBillingCycleConfigModelInvoicePlacementONSCHEDULEDINVOICE SubscriptionBillingCycleConfigModelInvoicePlacement = "ON_SCHEDULED_INVOICE"
	SubscriptionBillingCycleConfigModelInvoicePlacementONUSAGEINVOICE     SubscriptionBillingCycleConfigModelInvoicePlacement = "ON_USAGE_INVOICE"
)

type SubscriptionBillingCycleConfigModel struct {
	APIResource
	// The date this subscription's billing cycle is anchored to.
	AnchorDate string `json:"anchor_date"`
	// Controls whether this subscription consolidates onto usage invoices or gets its own scheduled invoice.
	InvoicePlacement SubscriptionBillingCycleConfigModelInvoicePlacement `json:"invoice_placement"`
}
