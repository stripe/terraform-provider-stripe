//
// File generated from our OpenAPI spec
//

package metronome

type RelativeDateModelUnit string

// List of values that RelativeDateModelUnit can take
const (
	RelativeDateModelUnitDAYS   RelativeDateModelUnit = "DAYS"
	RelativeDateModelUnitWEEKS  RelativeDateModelUnit = "WEEKS"
	RelativeDateModelUnitMONTHS RelativeDateModelUnit = "MONTHS"
	RelativeDateModelUnitYEARS  RelativeDateModelUnit = "YEARS"
)

type RelativeDateModel struct {
	APIResource
	Unit  RelativeDateModelUnit `json:"unit"`
	Value int64                 `json:"value"`
}
