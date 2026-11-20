//
// File generated from our OpenAPI spec
//

package metronome

type OverrideTemplateModelType string

// List of values that OverrideTemplateModelType can take
const (
	OverrideTemplateModelTypeOVERWRITE  OverrideTemplateModelType = "OVERWRITE"
	OverrideTemplateModelTypeMULTIPLIER OverrideTemplateModelType = "MULTIPLIER"
	OverrideTemplateModelTypeTIERED     OverrideTemplateModelType = "TIERED"
)

type OverrideTemplateModelTarget string

// List of values that OverrideTemplateModelTarget can take
const (
	OverrideTemplateModelTargetCOMMITRATE OverrideTemplateModelTarget = "COMMIT_RATE"
	OverrideTemplateModelTargetLISTRATE   OverrideTemplateModelTarget = "LIST_RATE"
)

type OverrideTemplateModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type OverrideTemplateModel struct {
	APIResource
	ApplicableProductTags []string                          `json:"applicable_product_tags,omitempty"`
	Duration              *RelativeDateModel                `json:"duration,omitempty"`
	Entitled              bool                              `json:"entitled,omitempty"`
	ID                    string                            `json:"id"`
	IsCommitSpecific      bool                              `json:"is_commit_specific,omitempty"`
	Multiplier            Number                            `json:"multiplier,omitempty"`
	OverrideSpecifiers    []*OverrideSpecifierTemplateModel `json:"override_specifiers"`
	OverrideTiers         []*OverrideTierModel              `json:"override_tiers,omitempty"`
	OverwriteRate         *OverwriteRateModel               `json:"overwrite_rate,omitempty"`
	Priority              Number                            `json:"priority,omitempty"`
	Product               *OverrideTemplateModelProduct     `json:"product,omitempty"`
	StartingAtOffset      *RelativeDateModel                `json:"starting_at_offset"`
	Target                OverrideTemplateModelTarget       `json:"target,omitempty"`
	Type                  OverrideTemplateModelType         `json:"type,omitempty"`
}
