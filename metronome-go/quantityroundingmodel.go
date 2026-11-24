//
// File generated from our OpenAPI spec
//

package metronome

type QuantityRoundingModelRoundingMethod string

// List of values that QuantityRoundingModelRoundingMethod can take
const (
	QuantityRoundingModelRoundingMethodRoundUp     QuantityRoundingModelRoundingMethod = "round_up"
	QuantityRoundingModelRoundingMethodRoundDown   QuantityRoundingModelRoundingMethod = "round_down"
	QuantityRoundingModelRoundingMethodRoundHalfUp QuantityRoundingModelRoundingMethod = "round_half_up"
	QuantityRoundingModelRoundingMethodROUNDUP     QuantityRoundingModelRoundingMethod = "ROUND_UP"
	QuantityRoundingModelRoundingMethodROUNDDOWN   QuantityRoundingModelRoundingMethod = "ROUND_DOWN"
	QuantityRoundingModelRoundingMethodROUNDHALFUP QuantityRoundingModelRoundingMethod = "ROUND_HALF_UP"
)

// Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.
type QuantityRoundingModel struct {
	APIResource
	DecimalPlaces  Number                              `json:"decimal_places"`
	RoundingMethod QuantityRoundingModelRoundingMethod `json:"rounding_method"`
}
