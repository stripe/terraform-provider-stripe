//
// File generated from our OpenAPI spec
//

package metronome

type ProrationRoundingConfigModelRoundingMethod string

// List of values that ProrationRoundingConfigModelRoundingMethod can take
const (
	ProrationRoundingConfigModelRoundingMethodHALFUP  ProrationRoundingConfigModelRoundingMethod = "HALF_UP"
	ProrationRoundingConfigModelRoundingMethodFLOOR   ProrationRoundingConfigModelRoundingMethod = "FLOOR"
	ProrationRoundingConfigModelRoundingMethodCEILING ProrationRoundingConfigModelRoundingMethod = "CEILING"
	ProrationRoundingConfigModelRoundingMethodHalfUp  ProrationRoundingConfigModelRoundingMethod = "half_up"
	ProrationRoundingConfigModelRoundingMethodFloor   ProrationRoundingConfigModelRoundingMethod = "floor"
	ProrationRoundingConfigModelRoundingMethodCeiling ProrationRoundingConfigModelRoundingMethod = "ceiling"
)

type ProrationRoundingConfigModel struct {
	APIResource
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).
	DecimalPlaces  Number                                     `json:"decimal_places"`
	RoundingMethod ProrationRoundingConfigModelRoundingMethod `json:"rounding_method"`
}
