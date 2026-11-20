//
// File generated from our OpenAPI spec
//

package metronome

type V1VoidInvoiceResponseData struct {
	ID string `json:"id"`
}
type V1VoidInvoiceResponse struct {
	APIResource
	Data *V1VoidInvoiceResponseData `json:"data,omitempty"`
}
