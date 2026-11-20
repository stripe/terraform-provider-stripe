//
// File generated from our OpenAPI spec
//

package metronome

type V1ListWebhookDeliveriesResponse struct {
	APIResource
	Data []*WebhookDeliveryModel `json:"data"`
	// Cursor for the next page of results, when one exists.
	NextPage string `json:"next_page,omitempty"`
}
