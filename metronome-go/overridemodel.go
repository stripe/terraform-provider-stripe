//
// File generated from our OpenAPI spec
//

package metronome

type OverrideModelType string

// List of values that OverrideModelType can take
const (
	OverrideModelTypeOVERWRITE  OverrideModelType = "OVERWRITE"
	OverrideModelTypeMULTIPLIER OverrideModelType = "MULTIPLIER"
	OverrideModelTypeTIERED     OverrideModelType = "TIERED"
)

type OverrideModelTarget string

// List of values that OverrideModelTarget can take
const (
	OverrideModelTargetCOMMITRATE OverrideModelTarget = "COMMIT_RATE"
	OverrideModelTargetLISTRATE   OverrideModelTarget = "LIST_RATE"
)

type OverrideModelRateType string

// List of values that OverrideModelRateType can take
const (
	OverrideModelRateTypeFLAT             OverrideModelRateType = "FLAT"
	OverrideModelRateTypeFlat             OverrideModelRateType = "flat"
	OverrideModelRateTypePERCENTAGE       OverrideModelRateType = "PERCENTAGE"
	OverrideModelRateTypePercentage       OverrideModelRateType = "percentage"
	OverrideModelRateTypeSUBSCRIPTION     OverrideModelRateType = "SUBSCRIPTION"
	OverrideModelRateTypeSubscription     OverrideModelRateType = "subscription"
	OverrideModelRateTypeTIERED           OverrideModelRateType = "TIERED"
	OverrideModelRateTypeTiered           OverrideModelRateType = "tiered"
	OverrideModelRateTypeTIEREDPERCENTAGE OverrideModelRateType = "TIERED_PERCENTAGE"
	OverrideModelRateTypeTieredPercentage OverrideModelRateType = "tiered_percentage"
	OverrideModelRateTypeCUSTOM           OverrideModelRateType = "CUSTOM"
	OverrideModelRateTypeCustom           OverrideModelRateType = "custom"
)

type OverrideModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type OverrideModel struct {
	APIResource
	ApplicableProductTags []string         `json:"applicable_product_tags,omitempty"`
	CreatedAt             string           `json:"created_at"`
	CreditType            *CreditTypeModel `json:"credit_type,omitempty"`
	EndingBefore          string           `json:"ending_before,omitempty"`
	Entitled              bool             `json:"entitled,omitempty"`
	ID                    string           `json:"id"`
	IsCommitSpecific      bool             `json:"is_commit_specific,omitempty"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be set to true.
	IsProrated         bool                      `json:"is_prorated,omitempty"`
	Multiplier         Number                    `json:"multiplier,omitempty"`
	OverrideSpecifiers []*OverrideSpecifierModel `json:"override_specifiers,omitempty"`
	OverrideTiers      []*OverrideTierModel      `json:"override_tiers,omitempty"`
	OverwriteRate      *OverwriteRateModel       `json:"overwrite_rate,omitempty"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    Number                `json:"price,omitempty"`
	Priority Number                `json:"priority,omitempty"`
	Product  *OverrideModelProduct `json:"product,omitempty"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity   Number                `json:"quantity,omitempty"`
	RateType   OverrideModelRateType `json:"rate_type,omitempty"`
	StartingAt string                `json:"starting_at"`
	Target     OverrideModelTarget   `json:"target,omitempty"`
	// Only set for TIERED rate_type.
	Tiers []*TierModel      `json:"tiers,omitempty"`
	Type  OverrideModelType `json:"type,omitempty"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate processors.
	Value map[string]JSON `json:"value,omitempty"`
}
