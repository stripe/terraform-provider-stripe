//
// File generated from our OpenAPI spec
//

package metronome

type ProrationRoundingConfigV2ModelRoundingMethod string

// List of values that ProrationRoundingConfigV2ModelRoundingMethod can take
const (
	ProrationRoundingConfigV2ModelRoundingMethodHALFUP  ProrationRoundingConfigV2ModelRoundingMethod = "HALF_UP"
	ProrationRoundingConfigV2ModelRoundingMethodFLOOR   ProrationRoundingConfigV2ModelRoundingMethod = "FLOOR"
	ProrationRoundingConfigV2ModelRoundingMethodCEILING ProrationRoundingConfigV2ModelRoundingMethod = "CEILING"
	ProrationRoundingConfigV2ModelRoundingMethodHalfUp  ProrationRoundingConfigV2ModelRoundingMethod = "half_up"
	ProrationRoundingConfigV2ModelRoundingMethodFloor   ProrationRoundingConfigV2ModelRoundingMethod = "floor"
	ProrationRoundingConfigV2ModelRoundingMethodCeiling ProrationRoundingConfigV2ModelRoundingMethod = "ceiling"
)

type ProrationRoundingConfigV2Model struct {
	APIResource
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit).
	DecimalPlaces  Number                                       `json:"decimal_places"`
	RoundingMethod ProrationRoundingConfigV2ModelRoundingMethod `json:"rounding_method"`
}
