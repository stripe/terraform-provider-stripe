//
// File generated from our OpenAPI spec
//

package metronome

type OverrideV2ModelType string

// List of values that OverrideV2ModelType can take
const (
	OverrideV2ModelTypeOVERWRITE  OverrideV2ModelType = "OVERWRITE"
	OverrideV2ModelTypeMULTIPLIER OverrideV2ModelType = "MULTIPLIER"
	OverrideV2ModelTypeTIERED     OverrideV2ModelType = "TIERED"
)

type OverrideV2ModelTarget string

// List of values that OverrideV2ModelTarget can take
const (
	OverrideV2ModelTargetCOMMITRATE OverrideV2ModelTarget = "COMMIT_RATE"
	OverrideV2ModelTargetLISTRATE   OverrideV2ModelTarget = "LIST_RATE"
)

type OverrideV2ModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type OverrideV2Model struct {
	APIResource
	ApplicableProductTags []string                    `json:"applicable_product_tags,omitempty"`
	CreatedAt             string                      `json:"created_at"`
	EndingBefore          string                      `json:"ending_before,omitempty"`
	Entitled              bool                        `json:"entitled,omitempty"`
	ID                    string                      `json:"id"`
	IsCommitSpecific      bool                        `json:"is_commit_specific,omitempty"`
	Multiplier            Number                      `json:"multiplier,omitempty"`
	OverrideSpecifiers    []*OverrideSpecifierV2Model `json:"override_specifiers,omitempty"`
	OverrideTiers         []*OverrideTierModel        `json:"override_tiers,omitempty"`
	OverwriteRate         *OverwriteRateV2Model       `json:"overwrite_rate,omitempty"`
	Priority              Number                      `json:"priority,omitempty"`
	Product               *OverrideV2ModelProduct     `json:"product,omitempty"`
	StartingAt            string                      `json:"starting_at"`
	Target                OverrideV2ModelTarget       `json:"target,omitempty"`
	Type                  OverrideV2ModelType         `json:"type,omitempty"`
}
