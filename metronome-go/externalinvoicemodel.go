//
// File generated from our OpenAPI spec
//

package metronome

type ExternalInvoiceModelBillingProviderType string

// List of values that ExternalInvoiceModelBillingProviderType can take
const (
	ExternalInvoiceModelBillingProviderTypeAwsMarketplace   ExternalInvoiceModelBillingProviderType = "aws_marketplace"
	ExternalInvoiceModelBillingProviderTypeStripe           ExternalInvoiceModelBillingProviderType = "stripe"
	ExternalInvoiceModelBillingProviderTypeNetsuite         ExternalInvoiceModelBillingProviderType = "netsuite"
	ExternalInvoiceModelBillingProviderTypeCustom           ExternalInvoiceModelBillingProviderType = "custom"
	ExternalInvoiceModelBillingProviderTypeAzureMarketplace ExternalInvoiceModelBillingProviderType = "azure_marketplace"
	ExternalInvoiceModelBillingProviderTypeQuickbooksOnline ExternalInvoiceModelBillingProviderType = "quickbooks_online"
	ExternalInvoiceModelBillingProviderTypeWorkday          ExternalInvoiceModelBillingProviderType = "workday"
	ExternalInvoiceModelBillingProviderTypeGcpMarketplace   ExternalInvoiceModelBillingProviderType = "gcp_marketplace"
	ExternalInvoiceModelBillingProviderTypeMetronome        ExternalInvoiceModelBillingProviderType = "metronome"
)

type ExternalInvoiceModelExternalStatus string

// List of values that ExternalInvoiceModelExternalStatus can take
const (
	ExternalInvoiceModelExternalStatusDRAFT               ExternalInvoiceModelExternalStatus = "DRAFT"
	ExternalInvoiceModelExternalStatusFINALIZED           ExternalInvoiceModelExternalStatus = "FINALIZED"
	ExternalInvoiceModelExternalStatusPAID                ExternalInvoiceModelExternalStatus = "PAID"
	ExternalInvoiceModelExternalStatusPARTIALLYPAID       ExternalInvoiceModelExternalStatus = "PARTIALLY_PAID"
	ExternalInvoiceModelExternalStatusUNCOLLECTIBLE       ExternalInvoiceModelExternalStatus = "UNCOLLECTIBLE"
	ExternalInvoiceModelExternalStatusVOID                ExternalInvoiceModelExternalStatus = "VOID"
	ExternalInvoiceModelExternalStatusDELETED             ExternalInvoiceModelExternalStatus = "DELETED"
	ExternalInvoiceModelExternalStatusPAYMENTFAILED       ExternalInvoiceModelExternalStatus = "PAYMENT_FAILED"
	ExternalInvoiceModelExternalStatusINVALIDREQUESTERROR ExternalInvoiceModelExternalStatus = "INVALID_REQUEST_ERROR"
	ExternalInvoiceModelExternalStatusSKIPPED             ExternalInvoiceModelExternalStatus = "SKIPPED"
	ExternalInvoiceModelExternalStatusSENT                ExternalInvoiceModelExternalStatus = "SENT"
	ExternalInvoiceModelExternalStatusQUEUED              ExternalInvoiceModelExternalStatus = "QUEUED"
)

// Tax details for the invoice, if available from the billing provider.
type ExternalInvoiceModelTax struct {
	// The total taxable amount of the invoice.
	TotalTaxableAmount Number `json:"total_taxable_amount,omitempty"`
	// The total tax amount applied to the invoice.
	TotalTaxAmount Number `json:"total_tax_amount,omitempty"`
	// The transaction ID associated with the tax calculation.
	TransactionID string `json:"transaction_id,omitempty"`
}
type ExternalInvoiceModel struct {
	APIResource
	// Error message from the billing provider, if available.
	BillingProviderError string                                  `json:"billing_provider_error,omitempty"`
	BillingProviderType  ExternalInvoiceModelBillingProviderType `json:"billing_provider_type"`
	// The ID of the payment in the external system, if available.
	ExternalPaymentID string                             `json:"external_payment_id,omitempty"`
	ExternalStatus    ExternalInvoiceModelExternalStatus `json:"external_status,omitempty"`
	// The subtotal amount invoiced, if available from the billing provider.
	InvoicedSubTotal Number `json:"invoiced_sub_total,omitempty"`
	// The total amount invoiced, if available from the billing provider.
	InvoicedTotal     Number `json:"invoiced_total,omitempty"`
	InvoiceID         string `json:"invoice_id,omitempty"`
	IssuedAtTimestamp string `json:"issued_at_timestamp,omitempty"`
	// A URL to the PDF of the invoice, if available from the billing provider.
	PDFURL string `json:"pdf_url,omitempty"`
	// Tax details for the invoice, if available from the billing provider.
	Tax *ExternalInvoiceModelTax `json:"tax,omitempty"`
}
