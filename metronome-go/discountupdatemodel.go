//
// File generated from our OpenAPI spec
//

package metronome

type DiscountUpdateModel struct {
	APIResource
	CustomFields         map[string]string                `json:"custom_fields,omitempty"`
	ID                   string                           `json:"id"`
	Name                 string                           `json:"name,omitempty"`
	NetsuiteSalesOrderID string                           `json:"netsuite_sales_order_id,omitempty"`
	Schedule             *SchedulePointInTimeInputV2Model `json:"schedule,omitempty"`
}
