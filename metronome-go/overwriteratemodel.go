//
// File generated from our OpenAPI spec
//

package metronome

type OverwriteRateModelRateType string

// List of values that OverwriteRateModelRateType can take
const (
	OverwriteRateModelRateTypeFLAT             OverwriteRateModelRateType = "FLAT"
	OverwriteRateModelRateTypeFlat             OverwriteRateModelRateType = "flat"
	OverwriteRateModelRateTypePERCENTAGE       OverwriteRateModelRateType = "PERCENTAGE"
	OverwriteRateModelRateTypePercentage       OverwriteRateModelRateType = "percentage"
	OverwriteRateModelRateTypeSUBSCRIPTION     OverwriteRateModelRateType = "SUBSCRIPTION"
	OverwriteRateModelRateTypeSubscription     OverwriteRateModelRateType = "subscription"
	OverwriteRateModelRateTypeTIERED           OverwriteRateModelRateType = "TIERED"
	OverwriteRateModelRateTypeTiered           OverwriteRateModelRateType = "tiered"
	OverwriteRateModelRateTypeTIEREDPERCENTAGE OverwriteRateModelRateType = "TIERED_PERCENTAGE"
	OverwriteRateModelRateTypeTieredPercentage OverwriteRateModelRateType = "tiered_percentage"
	OverwriteRateModelRateTypeCUSTOM           OverwriteRateModelRateType = "CUSTOM"
	OverwriteRateModelRateTypeCustom           OverwriteRateModelRateType = "custom"
)

type OverwriteRateModel struct {
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
	Quantity Number                     `json:"quantity,omitempty"`
	RateType OverwriteRateModelRateType `json:"rate_type"`
	// Only set for TIERED rate_type.
	Tiers []*TierModel `json:"tiers,omitempty"`
}
