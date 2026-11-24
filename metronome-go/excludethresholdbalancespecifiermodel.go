//
// File generated from our OpenAPI spec
//

package metronome

type ExcludeThresholdBalanceSpecifierModel struct {
	APIResource
	// If provided, balances with all the custom fields will not be considered when evaluating threshold billing
	CustomFieldFilters []*BalanceCustomFieldFilterTypeModel `json:"custom_field_filters"`
}
