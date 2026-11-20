//
// File generated from our OpenAPI spec
//

package metronome

type V1ListCustomerAlertsResponse struct {
	APIResource
	Data     []*CustomerAlertModel `json:"data"`
	NextPage string                `json:"next_page"`
}
