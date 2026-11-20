//
// File generated from our OpenAPI spec
//

package metronome

// Whether the created commits will be charged at commit rate or list rate.
type PrepaidBalanceThresholdCommitV2ModelRateType string

// List of values that PrepaidBalanceThresholdCommitV2ModelRateType can take
const (
	PrepaidBalanceThresholdCommitV2ModelRateTypeCOMMITRATE PrepaidBalanceThresholdCommitV2ModelRateType = "COMMIT_RATE"
	PrepaidBalanceThresholdCommitV2ModelRateTypeCommitRate PrepaidBalanceThresholdCommitV2ModelRateType = "commit_rate"
	PrepaidBalanceThresholdCommitV2ModelRateTypeLISTRATE   PrepaidBalanceThresholdCommitV2ModelRateType = "LIST_RATE"
	PrepaidBalanceThresholdCommitV2ModelRateTypeListRate   PrepaidBalanceThresholdCommitV2ModelRateType = "list_rate"
)

type PrepaidBalanceThresholdCommitV2Model struct {
	APIResource
	// Which products the threshold commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitempty"`
	// Which tags the threshold commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductTags []string `json:"applicable_product_tags,omitempty"`
	Description           string   `json:"description,omitempty"`
	// The length of time the created commit will be valid, starting from the end of the invoice's service period. If not provided, defaults to one year.
	Duration *RelativeDateModel `json:"duration,omitempty"`
	// Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.
	Name string `json:"name,omitempty"`
	// The priority of the commit, used to determine drawdown order. Lower priority commits are consumed first. Defaults to 100 if not specified.
	Priority Number `json:"priority,omitempty"`
	// The commit product that will be used to generate the line item for commit payment.
	ProductID string `json:"product_id"`
	// Whether the created commits will be charged at commit rate or list rate.
	RateType PrepaidBalanceThresholdCommitV2ModelRateType `json:"rate_type,omitempty"`
	// Fraction of the created commit's unused balance that will roll over. Must be between 0 and 1.
	RolloverFraction Number `json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
	Specifiers []*CommitSpecifierInputModel `json:"specifiers,omitempty"`
}
