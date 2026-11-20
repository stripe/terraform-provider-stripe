//
// File generated from our OpenAPI spec
//

package metronome

type CommitAddInvoiceScheduleModelScheduleItem struct {
	Amount    Number `json:"amount,omitempty"`
	ID        string `json:"id"`
	InvoiceID string `json:"invoice_id,omitempty"`
	Quantity  Number `json:"quantity,omitempty"`
	Timestamp string `json:"timestamp"`
	UnitPrice Number `json:"unit_price,omitempty"`
}
type CommitAddInvoiceScheduleModel struct {
	APIResource
	CreditType *CreditTypeModel `json:"credit_type,omitempty"`
	// If true, this schedule will not generate an invoice.
	DoNotInvoice  bool                                         `json:"do_not_invoice,omitempty"`
	ScheduleItems []*CommitAddInvoiceScheduleModelScheduleItem `json:"schedule_items,omitempty"`
}
