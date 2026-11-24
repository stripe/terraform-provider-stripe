//
// File generated from our OpenAPI spec
//

package metronome

type WebhookDeliveryModel struct {
	APIResource
	CustomerID       string `json:"customer_id"`
	FirstAttemptedAt string `json:"first_attempted_at"`
	// ID of the triggered notification.
	ID                     string `json:"id"`
	LastAttemptedAt        string `json:"last_attempted_at"`
	LastResponseStatusCode Number `json:"last_response_status_code"`
	// ID of the notification configuration, when applicable.
	NotificationID   string `json:"notification_id"`
	NotificationType string `json:"notification_type"`
	NumAttempts      Number `json:"num_attempts"`
	// ID of the webhook destination.
	WebhookID string `json:"webhook_id"`
}
