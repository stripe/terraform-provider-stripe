//
// File generated from our OpenAPI spec
//

package metronome

type OverrideSpecifierModelBillingFrequency string

// List of values that OverrideSpecifierModelBillingFrequency can take
const (
	OverrideSpecifierModelBillingFrequencyMONTHLY   OverrideSpecifierModelBillingFrequency = "MONTHLY"
	OverrideSpecifierModelBillingFrequencyQUARTERLY OverrideSpecifierModelBillingFrequency = "QUARTERLY"
	OverrideSpecifierModelBillingFrequencyANNUAL    OverrideSpecifierModelBillingFrequency = "ANNUAL"
	OverrideSpecifierModelBillingFrequencyWEEKLY    OverrideSpecifierModelBillingFrequency = "WEEKLY"
	OverrideSpecifierModelBillingFrequencyMonthly   OverrideSpecifierModelBillingFrequency = "monthly"
	OverrideSpecifierModelBillingFrequencyQuarterly OverrideSpecifierModelBillingFrequency = "quarterly"
	OverrideSpecifierModelBillingFrequencyAnnual    OverrideSpecifierModelBillingFrequency = "annual"
	OverrideSpecifierModelBillingFrequencyWeekly    OverrideSpecifierModelBillingFrequency = "weekly"
)

type OverrideSpecifierModel struct {
	APIResource
	AnyCommitOrCreditIDs []string                               `json:"any_commit_or_credit_ids,omitempty"`
	BillingFrequency     OverrideSpecifierModelBillingFrequency `json:"billing_frequency,omitempty"`
	CommitIDs            []string                               `json:"commit_ids,omitempty"`
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude                 []*ExcludeSpecifierModel `json:"exclude,omitempty"`
	PresentationGroupValues map[string]*string       `json:"presentation_group_values,omitempty"`
	PricingGroupValues      map[string]string        `json:"pricing_group_values,omitempty"`
	ProductID               string                   `json:"product_id,omitempty"`
	ProductTags             []string                 `json:"product_tags,omitempty"`
	RecurringCommitIDs      []string                 `json:"recurring_commit_ids,omitempty"`
}
