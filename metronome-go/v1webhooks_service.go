//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v1WebhooksService is used to invoke v1webhooks related APIs.
type v1WebhooksService struct {
	B   Backend
	Key string
}

// Retrieves a triggered notification and its delivery attempts for a webhook. Webhook delivery history is retained for 30 days.
func (c v1WebhooksService) GetWebhookDelivery(ctx context.Context, id string, params *V1WebhooksGetWebhookDeliveryParams) (*V1GetWebhookDeliveryResponse, error) {
	if params == nil {
		params = &V1WebhooksGetWebhookDeliveryParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/webhooks/%s/triggered_notifications/%s", StringValue(
			params.WebhookID), id)
	v1getwebhookdeliveryresponse := &V1GetWebhookDeliveryResponse{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, v1getwebhookdeliveryresponse)
	return v1getwebhookdeliveryresponse, err
}

// Lists triggered notifications delivered to a webhook, ordered by the time of the first delivery attempt.
//
// Results are limited to the most recent 30 days. The requested date range cannot exceed 30 days.
func (c v1WebhooksService) ListWebhookDeliveries(ctx context.Context, id string, params *V1WebhooksListWebhookDeliveriesParams) (*V1ListWebhookDeliveriesResponse, error) {
	if params == nil {
		params = &V1WebhooksListWebhookDeliveriesParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/webhooks/%s/triggered_notifications", id)
	v1listwebhookdeliveriesresponse := &V1ListWebhookDeliveriesResponse{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, v1listwebhookdeliveriesresponse)
	return v1listwebhookdeliveriesresponse, err
}
