//
// File generated from our OpenAPI spec
//

package metronome

type UpdatePrepaidBalanceThresholdCommitModelDurationUnit string

// List of values that UpdatePrepaidBalanceThresholdCommitModelDurationUnit can take
const (
	UpdatePrepaidBalanceThresholdCommitModelDurationUnitDAYS   UpdatePrepaidBalanceThresholdCommitModelDurationUnit = "DAYS"
	UpdatePrepaidBalanceThresholdCommitModelDurationUnitWEEKS  UpdatePrepaidBalanceThresholdCommitModelDurationUnit = "WEEKS"
	UpdatePrepaidBalanceThresholdCommitModelDurationUnitMONTHS UpdatePrepaidBalanceThresholdCommitModelDurationUnit = "MONTHS"
	UpdatePrepaidBalanceThresholdCommitModelDurationUnitYEARS  UpdatePrepaidBalanceThresholdCommitModelDurationUnit = "YEARS"
)

// Whether the created commits will be charged at commit rate or list rate. Set to null to clear a previously configured rate type.
type UpdatePrepaidBalanceThresholdCommitModelRateType string

// List of values that UpdatePrepaidBalanceThresholdCommitModelRateType can take
const (
	UpdatePrepaidBalanceThresholdCommitModelRateTypeCOMMITRATE UpdatePrepaidBalanceThresholdCommitModelRateType = "COMMIT_RATE"
	UpdatePrepaidBalanceThresholdCommitModelRateTypeCommitRate UpdatePrepaidBalanceThresholdCommitModelRateType = "commit_rate"
	UpdatePrepaidBalanceThresholdCommitModelRateTypeLISTRATE   UpdatePrepaidBalanceThresholdCommitModelRateType = "LIST_RATE"
	UpdatePrepaidBalanceThresholdCommitModelRateTypeListRate   UpdatePrepaidBalanceThresholdCommitModelRateType = "list_rate"
)

// The length of time the created commit will be valid, starting from the end of the invoice's service period. Set to null to clear a previously configured duration.
type UpdatePrepaidBalanceThresholdCommitModelDuration struct {
	Unit  UpdatePrepaidBalanceThresholdCommitModelDurationUnit `json:"unit"`
	Value int64                                                `json:"value"`
}
type UpdatePrepaidBalanceThresholdCommitModel struct {
	APIResource
	// Which products the threshold commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitempty"`
	// Which tags the threshold commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags []string `json:"applicable_product_tags,omitempty"`
	Description           string   `json:"description,omitempty"`
	// The length of time the created commit will be valid, starting from the end of the invoice's service period. Set to null to clear a previously configured duration.
	Duration *UpdatePrepaidBalanceThresholdCommitModelDuration `json:"duration,omitempty"`
	// Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.
	Name string `json:"name,omitempty"`
	// The priority of the commit, used to determine drawdown order. Lower priority commits are consumed first. Defaults to 100 if not specified. On updates, set to null to clear a previously configured priority.
	Priority Number `json:"priority,omitempty"`
	// The commit product that will be used to generate the line item for commit payment.
	ProductID string `json:"product_id,omitempty"`
	// Whether the created commits will be charged at commit rate or list rate. Set to null to clear a previously configured rate type.
	RateType UpdatePrepaidBalanceThresholdCommitModelRateType `json:"rate_type,omitempty"`
	// Fraction of the created commit's unused balance that will roll over. Must be between 0 and 1. Set to null to clear a previously configured rollover fraction.
	RolloverFraction Number `json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
	Specifiers []*CommitSpecifierInputModel `json:"specifiers,omitempty"`
}
