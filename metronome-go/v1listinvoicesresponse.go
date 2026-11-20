//
// File generated from our OpenAPI spec
//

package metronome

type V1ListInvoicesResponse struct {
	APIResource
	Data     []*InvoiceModel `json:"data"`
	NextPage string          `json:"next_page"`
}
