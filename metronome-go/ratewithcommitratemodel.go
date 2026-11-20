//
// File generated from our OpenAPI spec
//

package metronome

type RateWithCommitRateModelRateType string

// List of values that RateWithCommitRateModelRateType can take
const (
	RateWithCommitRateModelRateTypeFLAT             RateWithCommitRateModelRateType = "FLAT"
	RateWithCommitRateModelRateTypeFlat             RateWithCommitRateModelRateType = "flat"
	RateWithCommitRateModelRateTypePERCENTAGE       RateWithCommitRateModelRateType = "PERCENTAGE"
	RateWithCommitRateModelRateTypePercentage       RateWithCommitRateModelRateType = "percentage"
	RateWithCommitRateModelRateTypeSUBSCRIPTION     RateWithCommitRateModelRateType = "SUBSCRIPTION"
	RateWithCommitRateModelRateTypeSubscription     RateWithCommitRateModelRateType = "subscription"
	RateWithCommitRateModelRateTypeCUSTOM           RateWithCommitRateModelRateType = "CUSTOM"
	RateWithCommitRateModelRateTypeCustom           RateWithCommitRateModelRateType = "custom"
	RateWithCommitRateModelRateTypeTIERED           RateWithCommitRateModelRateType = "TIERED"
	RateWithCommitRateModelRateTypeTiered           RateWithCommitRateModelRateType = "tiered"
	RateWithCommitRateModelRateTypeTIEREDPERCENTAGE RateWithCommitRateModelRateType = "TIERED_PERCENTAGE"
	RateWithCommitRateModelRateTypeTieredPercentage RateWithCommitRateModelRateType = "tiered_percentage"
)

type RateWithCommitRateModel struct {
	APIResource
	CommitRate *CommitRateModel `json:"commit_rate,omitempty"`
	CreditType *CreditTypeModel `json:"credit_type,omitempty"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate processors.
	CustomRate map[string]JSON `json:"custom_rate,omitempty"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be set to true.
	IsProrated bool `json:"is_prorated,omitempty"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price Number `json:"price,omitempty"`
	// if pricing groups are used, this will contain the values used to calculate the price
	PricingGroupValues map[string]string `json:"pricing_group_values,omitempty"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity Number                          `json:"quantity,omitempty"`
	RateType RateWithCommitRateModelRateType `json:"rate_type"`
	// Only set for TIERED rate_type.
	Tiers []*TierModel `json:"tiers,omitempty"`
}
