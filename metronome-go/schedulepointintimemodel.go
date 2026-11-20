//
// File generated from our OpenAPI spec
//

package metronome

type SchedulePointInTimeModelScheduleItem struct {
	Amount    Number `json:"amount"`
	ID        string `json:"id"`
	InvoiceID string `json:"invoice_id,omitempty"`
	Quantity  Number `json:"quantity"`
	Timestamp string `json:"timestamp"`
	UnitPrice Number `json:"unit_price"`
}
type SchedulePointInTimeModel struct {
	APIResource
	CreditType *CreditTypeModel `json:"credit_type,omitempty"`
	// This field is only applicable to commit invoice schedules. If true, this schedule will not generate an invoice.
	DoNotInvoice  bool                                    `json:"do_not_invoice,omitempty"`
	ScheduleItems []*SchedulePointInTimeModelScheduleItem `json:"schedule_items,omitempty"`
}
