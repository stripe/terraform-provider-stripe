//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v2NotificationsService is used to invoke v2notifications related APIs.
type v2NotificationsService struct {
	B   Backend
	Key string
}

// Archive an offset lifecycle event notification configuration. Archived notifications are not processed.
func (c v2NotificationsService) ArchiveNotificationConfig(ctx context.Context, params *V2NotificationsArchiveNotificationConfigParams) (*V2ArchiveNotificationConfigResponse, error) {
	if params == nil {
		params = &V2NotificationsArchiveNotificationConfigParams{}
	}
	params.Context = ctx
	v2archivenotificationconfigresponse := &V2ArchiveNotificationConfigResponse{}
	err := c.B.Call(
		http.MethodPost, "/v2/notifications/archive", c.Key, params, v2archivenotificationconfigresponse)
	return v2archivenotificationconfigresponse, err
}

// Create an offset lifecycle event notification configuration. The lifecycle event type is inferred from the policy.type field.
func (c v2NotificationsService) CreateNotificationConfig(ctx context.Context, params *V2NotificationsCreateNotificationConfigParams) (*V2CreateNotificationConfigResponse, error) {
	if params == nil {
		params = &V2NotificationsCreateNotificationConfigParams{}
	}
	params.Context = ctx
	v2createnotificationconfigresponse := &V2CreateNotificationConfigResponse{}
	err := c.B.Call(
		http.MethodPost, "/v2/notifications/create", c.Key, params, v2createnotificationconfigresponse)
	return v2createnotificationconfigresponse, err
}

// Edit an existing offset notification, or enable/disable a system notification
func (c v2NotificationsService) EditNotificationConfig(ctx context.Context, params *V2NotificationsEditNotificationConfigParams) (*V2EditNotificationConfigResponse, error) {
	if params == nil {
		params = &V2NotificationsEditNotificationConfigParams{}
	}
	params.Context = ctx
	v2editnotificationconfigresponse := &V2EditNotificationConfigResponse{}
	err := c.B.Call(
		http.MethodPost, "/v2/notifications/edit", c.Key, params, v2editnotificationconfigresponse)
	return v2editnotificationconfigresponse, err
}

// Retrieve a specific offset lifecycle event notification configuration by ID.
func (c v2NotificationsService) GetNotificationConfig(ctx context.Context, params *V2NotificationsGetNotificationConfigParams) (*V2GetNotificationConfigResponse, error) {
	if params == nil {
		params = &V2NotificationsGetNotificationConfigParams{}
	}
	params.Context = ctx
	v2getnotificationconfigresponse := &V2GetNotificationConfigResponse{}
	err := c.B.Call(
		http.MethodPost, "/v2/notifications/get", c.Key, params, v2getnotificationconfigresponse)
	return v2getnotificationconfigresponse, err
}

// List offset lifecycle event notification configurations. These are user-created notifications that trigger at a specified time offset relative to lifecycle events. Returns a maximum of 400 results per request.
func (c v2NotificationsService) ListOffsetNotificationConfigs(ctx context.Context, params *V2NotificationsListOffsetNotificationConfigsParams) (*V2ListOffsetNotificationConfigsResponse, error) {
	if params == nil {
		params = &V2NotificationsListOffsetNotificationConfigsParams{}
	}
	params.Context = ctx
	v2listoffsetnotificationconfigsresponse := &V2ListOffsetNotificationConfigsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v2/notifications/offset/list", c.Key, params, v2listoffsetnotificationconfigsresponse)
	return v2listoffsetnotificationconfigsresponse, err
}

// List available system notification types. You can enable these notifications directly or use supported types to create offset notifications.
func (c v2NotificationsService) ListSystemNotificationConfigs(ctx context.Context, params *V2NotificationsListSystemNotificationConfigsParams) (*V2ListSystemNotificationConfigsResponse, error) {
	if params == nil {
		params = &V2NotificationsListSystemNotificationConfigsParams{}
	}
	params.Context = ctx
	v2listsystemnotificationconfigsresponse := &V2ListSystemNotificationConfigsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v2/notifications/system/list", c.Key, params, v2listsystemnotificationconfigsresponse)
	return v2listsystemnotificationconfigsresponse, err
}
