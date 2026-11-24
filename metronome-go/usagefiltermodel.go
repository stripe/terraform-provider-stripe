//
// File generated from our OpenAPI spec
//

package metronome

type UsageFilterModel struct {
	APIResource
	Current *BaseUsageFilterModel                         `json:"current"`
	Initial *BaseUsageFilterModel                         `json:"initial"`
	Updates []*BaseUsageFilterWithRequiredStartingAtModel `json:"updates"`
}
