//
// File generated from our OpenAPI spec
//

package metronome

type WebhookDeliveryAttemptModel struct {
	APIResource
	AttemptedAt        string `json:"attempted_at"`
	ResponseStatusCode Number `json:"response_status_code"`
}
