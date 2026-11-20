//
// File generated from our OpenAPI spec
//

package metronome

type SubscriptionBillingCycleConfigTemplateModelInvoicePlacement string

// List of values that SubscriptionBillingCycleConfigTemplateModelInvoicePlacement can take
const (
	SubscriptionBillingCycleConfigTemplateModelInvoicePlacementONSCHEDULEDINVOICE SubscriptionBillingCycleConfigTemplateModelInvoicePlacement = "ON_SCHEDULED_INVOICE"
	SubscriptionBillingCycleConfigTemplateModelInvoicePlacementONUSAGEINVOICE     SubscriptionBillingCycleConfigTemplateModelInvoicePlacement = "ON_USAGE_INVOICE"
)

type SubscriptionBillingCycleConfigTemplateModel struct {
	APIResource
	InvoicePlacement SubscriptionBillingCycleConfigTemplateModelInvoicePlacement `json:"invoice_placement,omitempty"`
}
