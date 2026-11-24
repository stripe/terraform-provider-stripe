//
// File generated from our OpenAPI spec
//

package metronome

type ScheduledChargeAddModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type ScheduledChargeAddModel struct {
	APIResource
	ID string `json:"id"`
	// displayed on invoices
	Name string `json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string                          `json:"netsuite_sales_order_id,omitempty"`
	Product              *ScheduledChargeAddModelProduct `json:"product"`
	Schedule             *SchedulePointInTimeModel       `json:"schedule"`
}
