//
// File generated from our OpenAPI spec
//

package metronome

type SpendBreakdownInvoiceModel struct {
	APIResource
	BreakdownEndTimestamp   string `json:"breakdown_end_timestamp"`
	BreakdownStartTimestamp string `json:"breakdown_start_timestamp"`
	ContractID              string `json:"contract_id,omitempty"`
	// When the invoice was created (UTC). This field is present for correction invoices only.
	CreatedAt  string           `json:"created_at,omitempty"`
	CreditType *CreditTypeModel `json:"credit_type"`
	CustomerID string           `json:"customer_id"`
	// End of the usage period this invoice covers (UTC)
	EndTimestamp string `json:"end_timestamp,omitempty"`
	ID           string `json:"id"`
	// When the invoice was issued (UTC)
	IssuedAt  string                  `json:"issued_at,omitempty"`
	LineItems []*InvoiceLineItemModel `json:"line_items"`
	// Beginning of the usage period this invoice covers (UTC)
	StartTimestamp string `json:"start_timestamp,omitempty"`
	Status         string `json:"status"`
	Type           string `json:"type"`
}
