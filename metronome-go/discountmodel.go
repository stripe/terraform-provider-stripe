//
// File generated from our OpenAPI spec
//

package metronome

type DiscountModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type DiscountModel struct {
	APIResource
	CustomFields map[string]string `json:"custom_fields,omitempty"`
	ID           string            `json:"id"`
	Name         string            `json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string                    `json:"netsuite_sales_order_id,omitempty"`
	Product              *DiscountModelProduct     `json:"product"`
	Schedule             *SchedulePointInTimeModel `json:"schedule"`
}
