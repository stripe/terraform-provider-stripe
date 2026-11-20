//
// File generated from our OpenAPI spec
//

package metronome

type ScheduledChargeUpdateModel struct {
	APIResource
	ID                   string                      `json:"id"`
	InvoiceSchedule      *InvoiceScheduleUpdateModel `json:"invoice_schedule,omitempty"`
	Name                 string                      `json:"name,omitempty"`
	NetsuiteSalesOrderID string                      `json:"netsuite_sales_order_id,omitempty"`
}
