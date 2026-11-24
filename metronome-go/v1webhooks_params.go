//
// File generated from our OpenAPI spec
//

package metronome

// Retrieves a triggered notification and its delivery attempts for a webhook. Webhook delivery history is retained for 30 days.
type V1WebhooksGetWebhookDeliveryParams struct {
	Params    `form:"*"`
	WebhookID *string `form:"-" json:"-"` // Included in URL
}
type V1WebhooksListWebhookDeliveriesFiltersParams struct {
	CustomerIDs []*string `form:"customer_ids" json:"customer_ids,omitempty"`
	// IDs of notification configurations.
	NotificationIDs   []*string `form:"notification_ids" json:"notification_ids,omitempty"`
	NotificationTypes []*string `form:"notification_types" json:"notification_types,omitempty"`
}

// Lists triggered notifications delivered to a webhook, ordered by the time of the first delivery attempt.
//
// Results are limited to the most recent 30 days. The requested date range cannot exceed 30 days.
type V1WebhooksListWebhookDeliveriesParams struct {
	Params `form:"*"`
	// Exclusive end timestamp. Defaults to the current time. The requested range cannot exceed 30 days.
	EndDate *string                                       `form:"end_date" json:"end_date,omitempty"`
	Filters *V1WebhooksListWebhookDeliveriesFiltersParams `form:"filters" json:"filters,omitempty"`
	// Number of deliveries to return per page.
	Limit *Number `form:"limit" json:"limit,omitempty"`
	// Cursor for the next page of results.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
	// Inclusive start timestamp. Must be within the last 30 days.
	StartDate *string `form:"start_date" json:"start_date"`
}
