//
// File generated from our OpenAPI spec
//

package metronome

type OverrideSpecifierV2ModelBillingFrequency string

// List of values that OverrideSpecifierV2ModelBillingFrequency can take
const (
	OverrideSpecifierV2ModelBillingFrequencyMONTHLY   OverrideSpecifierV2ModelBillingFrequency = "MONTHLY"
	OverrideSpecifierV2ModelBillingFrequencyQUARTERLY OverrideSpecifierV2ModelBillingFrequency = "QUARTERLY"
	OverrideSpecifierV2ModelBillingFrequencyANNUAL    OverrideSpecifierV2ModelBillingFrequency = "ANNUAL"
	OverrideSpecifierV2ModelBillingFrequencyWEEKLY    OverrideSpecifierV2ModelBillingFrequency = "WEEKLY"
)

type OverrideSpecifierV2Model struct {
	APIResource
	AnyCommitOrCreditIDs []string                                 `json:"any_commit_or_credit_ids,omitempty"`
	BillingFrequency     OverrideSpecifierV2ModelBillingFrequency `json:"billing_frequency,omitempty"`
	CommitIDs            []string                                 `json:"commit_ids,omitempty"`
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude                 []*ExcludeSpecifierModel `json:"exclude,omitempty"`
	PresentationGroupValues map[string]*string       `json:"presentation_group_values,omitempty"`
	PricingGroupValues      map[string]string        `json:"pricing_group_values,omitempty"`
	ProductID               string                   `json:"product_id,omitempty"`
	ProductTags             []string                 `json:"product_tags,omitempty"`
	RecurringCommitIDs      []string                 `json:"recurring_commit_ids,omitempty"`
}
