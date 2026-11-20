//
// File generated from our OpenAPI spec
//

package metronome

type V1ListSpendBreakdownInvoicesResponse struct {
	APIResource
	Data     []*SpendBreakdownInvoiceModel `json:"data"`
	NextPage string                        `json:"next_page"`
}
