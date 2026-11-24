//
// File generated from our OpenAPI spec
//

package metronome

// The operation to perform on the quantity
type QuantityConversionModelOperation string

// List of values that QuantityConversionModelOperation can take
const (
	QuantityConversionModelOperationMultiply QuantityConversionModelOperation = "multiply"
	QuantityConversionModelOperationDivide   QuantityConversionModelOperation = "divide"
	QuantityConversionModelOperationMULTIPLY QuantityConversionModelOperation = "MULTIPLY"
	QuantityConversionModelOperationDIVIDE   QuantityConversionModelOperation = "DIVIDE"
)

// Optional. Only valid for USAGE products. If provided, the quantity will be converted using the provided conversion factor and operation. For example, if the operation is "multiply" and the conversion factor is 100, then the quantity will be multiplied by 100. This can be used in cases where data is sent in one unit and priced in another.  For example, data could be sent in MB and priced in GB. In this case, the conversion factor would be 1024 and the operation would be "divide".
type QuantityConversionModel struct {
	APIResource
	// The factor to multiply or divide the quantity by.
	ConversionFactor Number `json:"conversion_factor"`
	// Optional name for this conversion.
	Name string `json:"name,omitempty"`
	// The operation to perform on the quantity
	Operation QuantityConversionModelOperation `json:"operation"`
}
