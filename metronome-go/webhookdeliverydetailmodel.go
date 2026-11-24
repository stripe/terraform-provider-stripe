//
// File generated from our OpenAPI spec
//

package metronome

type WebhookDeliveryDetailModel struct {
	APIResource
	// Delivery attempts ordered from oldest to newest.
	Attempts   []*WebhookDeliveryAttemptModel `json:"attempts"`
	CustomerID string                         `json:"customer_id"`
	// ID of the triggered notification.
	ID string `json:"id"`
	// ID of the notification configuration, when applicable.
	NotificationID string `json:"notification_id"`
	// The exact payload sent to the webhook destination.
	NotificationPayload string `json:"notification_payload"`
	NotificationType    string `json:"notification_type"`
	// ID of the webhook destination.
	WebhookID string `json:"webhook_id"`
}
