//
// File generated from our OpenAPI spec
//

package metronome

type ScheduledChargeModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type ScheduledChargeModel struct {
	APIResource
	ArchivedAt   string            `json:"archived_at,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
	ID           string            `json:"id"`
	// displayed on invoices
	Name string `json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string                       `json:"netsuite_sales_order_id,omitempty"`
	Product              *ScheduledChargeModelProduct `json:"product"`
	Schedule             *SchedulePointInTimeModel    `json:"schedule"`
}
