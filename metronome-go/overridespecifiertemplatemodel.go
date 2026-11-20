//
// File generated from our OpenAPI spec
//

package metronome

type OverrideSpecifierTemplateModelBillingFrequency string

// List of values that OverrideSpecifierTemplateModelBillingFrequency can take
const (
	OverrideSpecifierTemplateModelBillingFrequencyMONTHLY   OverrideSpecifierTemplateModelBillingFrequency = "MONTHLY"
	OverrideSpecifierTemplateModelBillingFrequencyQUARTERLY OverrideSpecifierTemplateModelBillingFrequency = "QUARTERLY"
	OverrideSpecifierTemplateModelBillingFrequencyANNUAL    OverrideSpecifierTemplateModelBillingFrequency = "ANNUAL"
	OverrideSpecifierTemplateModelBillingFrequencyWEEKLY    OverrideSpecifierTemplateModelBillingFrequency = "WEEKLY"
)

type OverrideSpecifierTemplateModel struct {
	APIResource
	AnyCommitOrCreditTemplateIDs []string                                       `json:"any_commit_or_credit_template_ids,omitempty"`
	BillingFrequency             OverrideSpecifierTemplateModelBillingFrequency `json:"billing_frequency,omitempty"`
	CommitTemplateIDs            []string                                       `json:"commit_template_ids,omitempty"`
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude                    []*ExcludeSpecifierModel `json:"exclude,omitempty"`
	PresentationGroupValues    map[string]*string       `json:"presentation_group_values,omitempty"`
	PricingGroupValues         map[string]string        `json:"pricing_group_values,omitempty"`
	ProductID                  string                   `json:"product_id,omitempty"`
	ProductTags                []string                 `json:"product_tags,omitempty"`
	RecurringCommitTemplateIDs []string                 `json:"recurring_commit_template_ids,omitempty"`
}
