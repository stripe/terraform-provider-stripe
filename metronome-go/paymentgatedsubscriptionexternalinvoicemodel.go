//
// File generated from our OpenAPI spec
//

package metronome

type PaymentGatedSubscriptionExternalInvoiceModelExternalStatus string

// List of values that PaymentGatedSubscriptionExternalInvoiceModelExternalStatus can take
const (
	PaymentGatedSubscriptionExternalInvoiceModelExternalStatusDRAFT               PaymentGatedSubscriptionExternalInvoiceModelExternalStatus = "DRAFT"
	PaymentGatedSubscriptionExternalInvoiceModelExternalStatusFINALIZED           PaymentGatedSubscriptionExternalInvoiceModelExternalStatus = "FINALIZED"
	PaymentGatedSubscriptionExternalInvoiceModelExternalStatusPAID                PaymentGatedSubscriptionExternalInvoiceModelExternalStatus = "PAID"
	PaymentGatedSubscriptionExternalInvoiceModelExternalStatusPARTIALLYPAID       PaymentGatedSubscriptionExternalInvoiceModelExternalStatus = "PARTIALLY_PAID"
	PaymentGatedSubscriptionExternalInvoiceModelExternalStatusUNCOLLECTIBLE       PaymentGatedSubscriptionExternalInvoiceModelExternalStatus = "UNCOLLECTIBLE"
	PaymentGatedSubscriptionExternalInvoiceModelExternalStatusVOID                PaymentGatedSubscriptionExternalInvoiceModelExternalStatus = "VOID"
	PaymentGatedSubscriptionExternalInvoiceModelExternalStatusDELETED             PaymentGatedSubscriptionExternalInvoiceModelExternalStatus = "DELETED"
	PaymentGatedSubscriptionExternalInvoiceModelExternalStatusPAYMENTFAILED       PaymentGatedSubscriptionExternalInvoiceModelExternalStatus = "PAYMENT_FAILED"
	PaymentGatedSubscriptionExternalInvoiceModelExternalStatusINVALIDREQUESTERROR PaymentGatedSubscriptionExternalInvoiceModelExternalStatus = "INVALID_REQUEST_ERROR"
	PaymentGatedSubscriptionExternalInvoiceModelExternalStatusSKIPPED             PaymentGatedSubscriptionExternalInvoiceModelExternalStatus = "SKIPPED"
	PaymentGatedSubscriptionExternalInvoiceModelExternalStatusSENT                PaymentGatedSubscriptionExternalInvoiceModelExternalStatus = "SENT"
	PaymentGatedSubscriptionExternalInvoiceModelExternalStatusQUEUED              PaymentGatedSubscriptionExternalInvoiceModelExternalStatus = "QUEUED"
)

type PaymentGatedSubscriptionExternalInvoiceModel struct {
	APIResource
	BillingProviderError string                                                     `json:"billing_provider_error,omitempty"`
	ExternalStatus       PaymentGatedSubscriptionExternalInvoiceModelExternalStatus `json:"external_status"`
	InvoiceID            string                                                     `json:"invoice_id,omitempty"`
}
