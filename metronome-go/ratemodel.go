//
// File generated from our OpenAPI spec
//

package metronome

type RateModelRateType string

// List of values that RateModelRateType can take
const (
	RateModelRateTypeFLAT             RateModelRateType = "FLAT"
	RateModelRateTypeFlat             RateModelRateType = "flat"
	RateModelRateTypePERCENTAGE       RateModelRateType = "PERCENTAGE"
	RateModelRateTypePercentage       RateModelRateType = "percentage"
	RateModelRateTypeSUBSCRIPTION     RateModelRateType = "SUBSCRIPTION"
	RateModelRateTypeSubscription     RateModelRateType = "subscription"
	RateModelRateTypeCUSTOM           RateModelRateType = "CUSTOM"
	RateModelRateTypeCustom           RateModelRateType = "custom"
	RateModelRateTypeTIERED           RateModelRateType = "TIERED"
	RateModelRateTypeTiered           RateModelRateType = "tiered"
	RateModelRateTypeTIEREDPERCENTAGE RateModelRateType = "TIERED_PERCENTAGE"
	RateModelRateTypeTieredPercentage RateModelRateType = "tiered_percentage"
)

type RateModel struct {
	APIResource
	CreditType *CreditTypeModel `json:"credit_type,omitempty"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate processors.
	CustomRate map[string]JSON `json:"custom_rate,omitempty"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be set to true.
	IsProrated    bool                `json:"is_prorated,omitempty"`
	MinimumConfig *MinimumConfigModel `json:"minimum_config,omitempty"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price Number `json:"price,omitempty"`
	// if pricing groups are used, this will contain the values used to calculate the price
	PricingGroupValues map[string]string `json:"pricing_group_values,omitempty"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity Number            `json:"quantity,omitempty"`
	RateType RateModelRateType `json:"rate_type"`
	// Only set for TIERED rate_type.
	Tiers []*TierModel `json:"tiers,omitempty"`
}
