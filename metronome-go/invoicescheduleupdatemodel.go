//
// File generated from our OpenAPI spec
//

package metronome

type InvoiceScheduleUpdateModelAddScheduleItem struct {
	Amount    Number `json:"amount,omitempty"`
	Quantity  Number `json:"quantity,omitempty"`
	Timestamp string `json:"timestamp"`
	UnitPrice Number `json:"unit_price,omitempty"`
}
type InvoiceScheduleUpdateModelUpdateScheduleItem struct {
	Amount    Number `json:"amount,omitempty"`
	ID        string `json:"id"`
	Quantity  Number `json:"quantity,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	UnitPrice Number `json:"unit_price,omitempty"`
}
type InvoiceScheduleUpdateModelRemoveScheduleItem struct {
	ID string `json:"id"`
}
type InvoiceScheduleUpdateModel struct {
	APIResource
	AddScheduleItems    []*InvoiceScheduleUpdateModelAddScheduleItem    `json:"add_schedule_items,omitempty"`
	RemoveScheduleItems []*InvoiceScheduleUpdateModelRemoveScheduleItem `json:"remove_schedule_items,omitempty"`
	UpdateScheduleItems []*InvoiceScheduleUpdateModelUpdateScheduleItem `json:"update_schedule_items,omitempty"`
}
