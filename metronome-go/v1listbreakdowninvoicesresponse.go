//
// File generated from our OpenAPI spec
//

package metronome

type V1ListBreakdownInvoicesResponse struct {
	APIResource
	Data     []*BreakdownInvoiceModel `json:"data"`
	NextPage string                   `json:"next_page"`
}
