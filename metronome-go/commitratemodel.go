//
// File generated from our OpenAPI spec
//

package metronome

type CommitRateModelRateType string

// List of values that CommitRateModelRateType can take
const (
	CommitRateModelRateTypeFLAT             CommitRateModelRateType = "FLAT"
	CommitRateModelRateTypeFlat             CommitRateModelRateType = "flat"
	CommitRateModelRateTypePERCENTAGE       CommitRateModelRateType = "PERCENTAGE"
	CommitRateModelRateTypePercentage       CommitRateModelRateType = "percentage"
	CommitRateModelRateTypeSUBSCRIPTION     CommitRateModelRateType = "SUBSCRIPTION"
	CommitRateModelRateTypeSubscription     CommitRateModelRateType = "subscription"
	CommitRateModelRateTypeTIERED           CommitRateModelRateType = "TIERED"
	CommitRateModelRateTypeTiered           CommitRateModelRateType = "tiered"
	CommitRateModelRateTypeTIEREDPERCENTAGE CommitRateModelRateType = "TIERED_PERCENTAGE"
	CommitRateModelRateTypeTieredPercentage CommitRateModelRateType = "tiered_percentage"
	CommitRateModelRateTypeCUSTOM           CommitRateModelRateType = "CUSTOM"
	CommitRateModelRateTypeCustom           CommitRateModelRateType = "custom"
)

// A distinct rate on the rate card. You can choose to use this rate rather than list rate when consuming a credit or commit.
type CommitRateModel struct {
	APIResource
	MinimumConfig *MinimumConfigModel `json:"minimum_config,omitempty"`
	// Commit rate price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    Number                  `json:"price,omitempty"`
	RateType CommitRateModelRateType `json:"rate_type"`
	// Only set for TIERED rate_type.
	Tiers []*TierModel `json:"tiers,omitempty"`
}
