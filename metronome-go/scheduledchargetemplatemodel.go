//
// File generated from our OpenAPI spec
//

package metronome

type ScheduledChargeTemplateModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type ScheduledChargeTemplateModel struct {
	APIResource
	CustomFields map[string]string                    `json:"custom_fields,omitempty"`
	Description  string                               `json:"description,omitempty"`
	ID           string                               `json:"id"`
	Name         string                               `json:"name,omitempty"`
	Product      *ScheduledChargeTemplateModelProduct `json:"product"`
	Schedule     *RelativeSchedulePointInTimeModel    `json:"schedule"`
}
