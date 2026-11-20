//
// File generated from our OpenAPI spec
//

package metronome

type V1ChargeSeatsResponseResellerRoyaltyResellerType string

// List of values that V1ChargeSeatsResponseResellerRoyaltyResellerType can take
const (
	V1ChargeSeatsResponseResellerRoyaltyResellerTypeAWS           V1ChargeSeatsResponseResellerRoyaltyResellerType = "AWS"
	V1ChargeSeatsResponseResellerRoyaltyResellerTypeAWSPROSERVICE V1ChargeSeatsResponseResellerRoyaltyResellerType = "AWS_PRO_SERVICE"
	V1ChargeSeatsResponseResellerRoyaltyResellerTypeGCP           V1ChargeSeatsResponseResellerRoyaltyResellerType = "GCP"
	V1ChargeSeatsResponseResellerRoyaltyResellerTypeGCPPROSERVICE V1ChargeSeatsResponseResellerRoyaltyResellerType = "GCP_PRO_SERVICE"
)

// Indicates if the invoice has been or will be sent to the configured customer billing provider. Defaults to `billable`.
type V1ChargeSeatsResponseBillableStatus string

// List of values that V1ChargeSeatsResponseBillableStatus can take
const (
	V1ChargeSeatsResponseBillableStatusBillable   V1ChargeSeatsResponseBillableStatus = "billable"
	V1ChargeSeatsResponseBillableStatusUnbillable V1ChargeSeatsResponseBillableStatus = "unbillable"
)

type V1ChargeSeatsResponseCorrectionRecord struct {
	CorrectedExternalInvoice *ExternalInvoiceModel `json:"corrected_external_invoice,omitempty"`
	CorrectedInvoiceID       string                `json:"corrected_invoice_id"`
	Memo                     string                `json:"memo"`
	Reason                   string                `json:"reason"`
}
type V1ChargeSeatsResponseResellerRoyaltyAwsOptions struct {
	AwsAccountNumber    string `json:"aws_account_number,omitempty"`
	AwsOfferID          string `json:"aws_offer_id,omitempty"`
	AwsPayerReferenceID string `json:"aws_payer_reference_id,omitempty"`
}
type V1ChargeSeatsResponseResellerRoyaltyGcpOptions struct {
	GcpAccountID string `json:"gcp_account_id,omitempty"`
	GcpOfferID   string `json:"gcp_offer_id,omitempty"`
}

// Only present for contract invoices with reseller royalties.
type V1ChargeSeatsResponseResellerRoyalty struct {
	AwsOptions         *V1ChargeSeatsResponseResellerRoyaltyAwsOptions  `json:"aws_options,omitempty"`
	Fraction           string                                           `json:"fraction"`
	GcpOptions         *V1ChargeSeatsResponseResellerRoyaltyGcpOptions  `json:"gcp_options,omitempty"`
	NetsuiteResellerID string                                           `json:"netsuite_reseller_id"`
	ResellerType       V1ChargeSeatsResponseResellerRoyaltyResellerType `json:"reseller_type"`
}

// Required on invoices with type USAGE_CONSOLIDATED. List of constituent invoices that were consolidated to create this invoice.
type V1ChargeSeatsResponseConstituentInvoice struct {
	ContractID string `json:"contract_id"`
	CustomerID string `json:"customer_id"`
	InvoiceID  string `json:"invoice_id"`
}

// Required for account hierarchy usage invoices. An object containing the contract and customer UUIDs that pay for this invoice.
type V1ChargeSeatsResponsePayer struct {
	ContractID string `json:"contract_id"`
	CustomerID string `json:"customer_id"`
}
type V1ChargeSeatsResponse struct {
	APIResource
	AmendmentID string `json:"amendment_id,omitempty"`
	// Indicates if the invoice has been or will be sent to the configured customer billing provider. Defaults to `billable`.
	BillableStatus V1ChargeSeatsResponseBillableStatus `json:"billable_status,omitempty"`
	// Required on invoices with type USAGE_CONSOLIDATED. List of constituent invoices that were consolidated to create this invoice.
	ConstituentInvoices  []*V1ChargeSeatsResponseConstituentInvoice `json:"constituent_invoices,omitempty"`
	ContractCustomFields map[string]string                          `json:"contract_custom_fields,omitempty"`
	ContractID           string                                     `json:"contract_id,omitempty"`
	CorrectionRecord     *V1ChargeSeatsResponseCorrectionRecord     `json:"correction_record,omitempty"`
	// When the invoice was created (UTC). This field is present for correction invoices only.
	CreatedAt            string            `json:"created_at,omitempty"`
	CreditType           *CreditTypeModel  `json:"credit_type"`
	CustomerCustomFields map[string]string `json:"customer_custom_fields,omitempty"`
	CustomerID           string            `json:"customer_id"`
	CustomFields         map[string]JSON   `json:"custom_fields,omitempty"`
	// End of the usage period this invoice covers (UTC)
	EndTimestamp    string                `json:"end_timestamp,omitempty"`
	ExternalInvoice *ExternalInvoiceModel `json:"external_invoice,omitempty"`
	ID              string                `json:"id"`
	// When the invoice was issued (UTC)
	IssuedAt            string                  `json:"issued_at,omitempty"`
	LineItems           []*InvoiceLineItemModel `json:"line_items"`
	NetPaymentTermsDays Number                  `json:"net_payment_terms_days,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id,omitempty"`
	// Required for account hierarchy usage invoices. An object containing the contract and customer UUIDs that pay for this invoice.
	Payer                    *V1ChargeSeatsResponsePayer `json:"payer,omitempty"`
	RegeneratedFromInvoiceID string                      `json:"regenerated_from_invoice_id,omitempty"`
	// Only present for contract invoices with reseller royalties.
	ResellerRoyalty       *V1ChargeSeatsResponseResellerRoyalty `json:"reseller_royalty,omitempty"`
	RevenueSystemInvoices []*RevenueSystemInvoiceModel          `json:"revenue_system_invoices,omitempty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id,omitempty"`
	// Beginning of the usage period this invoice covers (UTC)
	StartTimestamp string `json:"start_timestamp,omitempty"`
	Status         string `json:"status"`
	Total          Number `json:"total"`
	Type           string `json:"type"`
}
