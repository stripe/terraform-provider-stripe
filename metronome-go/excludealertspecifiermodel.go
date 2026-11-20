//
// File generated from our OpenAPI spec
//

package metronome

type ExcludeAlertSpecifierModel struct {
	APIResource
	// A list of custom field filters for notification types that support advanced filtering
	CustomFieldFilters []*CustomFieldFilterTypeModel `json:"custom_field_filters,omitempty"`
}
