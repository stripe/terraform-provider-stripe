//
// File generated from our OpenAPI spec
//

package metronome

type SpendThresholdCommitModel struct {
	APIResource
	Description string `json:"description,omitempty"`
	// Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.
	Name string `json:"name,omitempty"`
	// The priority of the commit, used to determine drawdown order. Lower priority commits are consumed first. Defaults to 100 if not specified.
	Priority Number `json:"priority,omitempty"`
	// The commit product that will be used to generate the line item for commit payment.
	ProductID string `json:"product_id"`
}
