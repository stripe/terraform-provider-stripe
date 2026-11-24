//
// File generated from our OpenAPI spec
//

package metronome

// Indicates how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
type AppliedCommitOrCreditModelAccessType string

// List of values that AppliedCommitOrCreditModelAccessType can take
const (
	AppliedCommitOrCreditModelAccessTypeSPEND    AppliedCommitOrCreditModelAccessType = "SPEND"
	AppliedCommitOrCreditModelAccessTypeQUANTITY AppliedCommitOrCreditModelAccessType = "QUANTITY"
)

type AppliedCommitOrCreditModelType string

// List of values that AppliedCommitOrCreditModelType can take
const (
	AppliedCommitOrCreditModelTypePREPAID  AppliedCommitOrCreditModelType = "PREPAID"
	AppliedCommitOrCreditModelTypePOSTPAID AppliedCommitOrCreditModelType = "POSTPAID"
	AppliedCommitOrCreditModelTypeCREDIT   AppliedCommitOrCreditModelType = "CREDIT"
)

type AppliedCommitOrCreditModel struct {
	APIResource
	// Indicates how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType AppliedCommitOrCreditModelAccessType `json:"access_type,omitempty"`
	ID         string                               `json:"id"`
	Type       AppliedCommitOrCreditModelType       `json:"type"`
}
