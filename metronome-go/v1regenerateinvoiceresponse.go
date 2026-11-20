//
// File generated from our OpenAPI spec
//

package metronome

type V1RegenerateInvoiceResponseData struct {
	// The new invoice id
	ID string `json:"id"`
}
type V1RegenerateInvoiceResponse struct {
	APIResource
	Data *V1RegenerateInvoiceResponseData `json:"data,omitempty"`
}
