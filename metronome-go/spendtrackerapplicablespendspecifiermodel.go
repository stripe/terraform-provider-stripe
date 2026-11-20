//
// File generated from our OpenAPI spec
//

package metronome

type SpendTrackerApplicableSpendSpecifierModelSpendType string

// List of values that SpendTrackerApplicableSpendSpecifierModelSpendType can take
const (
	SpendTrackerApplicableSpendSpecifierModelSpendTypeCOMMITPURCHASE SpendTrackerApplicableSpendSpecifierModelSpendType = "COMMIT_PURCHASE"
)

type SpendTrackerApplicableSpendSpecifierModelSource string

// List of values that SpendTrackerApplicableSpendSpecifierModelSource can take
const (
	SpendTrackerApplicableSpendSpecifierModelSourceTHRESHOLDRECHARGE SpendTrackerApplicableSpendSpecifierModelSource = "THRESHOLD_RECHARGE"
	SpendTrackerApplicableSpendSpecifierModelSourceMANUAL            SpendTrackerApplicableSpendSpecifierModelSource = "MANUAL"
)

type SpendTrackerApplicableSpendSpecifierModelDiscounted string

// List of values that SpendTrackerApplicableSpendSpecifierModelDiscounted can take
const (
	SpendTrackerApplicableSpendSpecifierModelDiscountedANY              SpendTrackerApplicableSpendSpecifierModelDiscounted = "ANY"
	SpendTrackerApplicableSpendSpecifierModelDiscountedDISCOUNTEDONLY   SpendTrackerApplicableSpendSpecifierModelDiscounted = "DISCOUNTED_ONLY"
	SpendTrackerApplicableSpendSpecifierModelDiscountedUNDISCOUNTEDONLY SpendTrackerApplicableSpendSpecifierModelDiscounted = "UNDISCOUNTED_ONLY"
)

type SpendTrackerApplicableSpendSpecifierModel struct {
	APIResource
	Discounted SpendTrackerApplicableSpendSpecifierModelDiscounted `json:"discounted,omitempty"`
	Sources    []SpendTrackerApplicableSpendSpecifierModelSource   `json:"sources"`
	SpendType  SpendTrackerApplicableSpendSpecifierModelSpendType  `json:"spend_type"`
}
