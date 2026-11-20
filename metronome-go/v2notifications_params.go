//
// File generated from our OpenAPI spec
//

package metronome

// Archive an offset lifecycle event notification configuration. Archived notifications are not processed.
type V2NotificationsArchiveNotificationConfigParams struct {
	Params `form:"*"`
	// The ID of the offset lifecycle event notification configuration to archive.
	//
	ID *string `form:"id" json:"id"`
}

// The offset lifecycle event policy that defines when and how this notification should be triggered. The lifecycle event type is inferred from the policy.type field.
type V2NotificationsCreateNotificationConfigPolicyParams struct {
	// ISO-8601 duration string indicating how much time before or after the base event this notification should be sent. Positive values indicate notifications after the event, negative values indicate notifications before the event. Examples: "P1D" (1 day after), "-PT2H" (2 hours before)
	//
	Offset *string `form:"offset" json:"offset"`
	// The type of lifecycle event that this offset is based on.
	//
	Type *string `form:"type" json:"type"`
}

// Create an offset lifecycle event notification configuration. The lifecycle event type is inferred from the policy.type field.
type V2NotificationsCreateNotificationConfigParams struct {
	Params `form:"*"`
	// The name for this offset notification configuration.
	//
	Name *string `form:"name" json:"name"`
	// The offset lifecycle event policy that defines when and how this notification should be triggered. The lifecycle event type is inferred from the policy.type field.
	//
	Policy *V2NotificationsCreateNotificationConfigPolicyParams `form:"policy" json:"policy"`
	// Optional uniqueness key to prevent duplicate notification configurations.
	//
	UniquenessKey *string `form:"uniqueness_key" json:"uniqueness_key,omitempty"`
}

// Edit an existing offset notification, or enable/disable a system notification
type V2NotificationsEditNotificationConfigParams struct {
	Params `form:"*"`
	// The ID of the notification configuration to edit. Not provided when updating the configuration for system events
	//
	ID *string `form:"id" json:"id,omitempty"`
	// Set to true to enable webhook messages for the notification indicated in the policy, false to disable.
	// Only supported by system lifecycle events.
	//
	IsEnabled *bool `form:"is_enabled" json:"is_enabled,omitempty"`
	// Updated policy configuration. The policy.type must match the existing lifecycle event type.
	Policy *JSON `form:"policy" json:"policy"`
}

// Retrieve a specific offset lifecycle event notification configuration by ID.
type V2NotificationsGetNotificationConfigParams struct {
	Params `form:"*"`
	// The ID of the notification configuration to retrieve
	ID *string `form:"id" json:"id"`
}

// List offset lifecycle event notification configurations. These are user-created notifications that trigger at a specified time offset relative to lifecycle events. Returns a maximum of 400 results per request.
type V2NotificationsListOffsetNotificationConfigsParams struct {
	Params `form:"*"`
	// Filter options for the notification configurations. If not provided, defaults to NOT_ARCHIVED.
	ArchiveFilter *string `form:"archive_filter" json:"archive_filter,omitempty"`
	Cursor        *string `form:"cursor" json:"cursor,omitempty"`
	Limit         *Number `form:"limit" json:"limit,omitempty"`
}

// List available system notification types. You can enable these notifications directly or use supported types to create offset notifications.
type V2NotificationsListSystemNotificationConfigsParams struct {
	Params `form:"*"`
}
