//
// File generated from our OpenAPI spec
//

package metronome

type AlertSpecifierModel struct {
	APIResource
	// A list of custom field filters for notification types that support advanced filtering
	CustomFieldFilters []*AlertSpecifierInclusiveCustomFieldFilterModel `json:"custom_field_filters,omitempty"`
	// If provided, the specifier will not apply to balances that matches the inclusion criteria and any of the excluding values.
	Exclude []*ExcludeAlertSpecifierModel `json:"exclude,omitempty"`
}
