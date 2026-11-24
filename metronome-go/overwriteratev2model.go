//
// File generated from our OpenAPI spec
//

package metronome

type OverwriteRateV2ModelRateType string

// List of values that OverwriteRateV2ModelRateType can take
const (
	OverwriteRateV2ModelRateTypeFLAT             OverwriteRateV2ModelRateType = "FLAT"
	OverwriteRateV2ModelRateTypeFlat             OverwriteRateV2ModelRateType = "flat"
	OverwriteRateV2ModelRateTypePERCENTAGE       OverwriteRateV2ModelRateType = "PERCENTAGE"
	OverwriteRateV2ModelRateTypePercentage       OverwriteRateV2ModelRateType = "percentage"
	OverwriteRateV2ModelRateTypeSUBSCRIPTION     OverwriteRateV2ModelRateType = "SUBSCRIPTION"
	OverwriteRateV2ModelRateTypeSubscription     OverwriteRateV2ModelRateType = "subscription"
	OverwriteRateV2ModelRateTypeTIERED           OverwriteRateV2ModelRateType = "TIERED"
	OverwriteRateV2ModelRateTypeTiered           OverwriteRateV2ModelRateType = "tiered"
	OverwriteRateV2ModelRateTypeTIEREDPERCENTAGE OverwriteRateV2ModelRateType = "TIERED_PERCENTAGE"
	OverwriteRateV2ModelRateTypeTieredPercentage OverwriteRateV2ModelRateType = "tiered_percentage"
	OverwriteRateV2ModelRateTypeCUSTOM           OverwriteRateV2ModelRateType = "CUSTOM"
	OverwriteRateV2ModelRateTypeCustom           OverwriteRateV2ModelRateType = "custom"
)

type OverwriteRateV2Model struct {
	APIResource
	CreditType *CreditTypeModel `json:"credit_type,omitempty"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate processors.
	CustomRate map[string]JSON `json:"custom_rate,omitempty"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be set to true.
	IsProrated    bool                `json:"is_prorated,omitempty"`
	MinimumConfig *MinimumConfigModel `json:"minimum_config,omitempty"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price Number `json:"price,omitempty"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity Number                       `json:"quantity,omitempty"`
	RateType OverwriteRateV2ModelRateType `json:"rate_type"`
	// Only set for TIERED rate_type.
	Tiers []*TierModel `json:"tiers,omitempty"`
}
