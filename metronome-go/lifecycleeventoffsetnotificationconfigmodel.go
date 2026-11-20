//
// File generated from our OpenAPI spec
//

package metronome

type LifecycleEventOffsetNotificationConfigModel struct {
	APIResource
	// When this notification configuration was archived
	ArchivedAt string `json:"archived_at"`
	// RFC 3339 timestamp when this notification configuration was created.
	//
	CreatedAt string `json:"created_at"`
	// Who created this notification configuration
	CreatedBy string `json:"created_by"`
	// The environment type where this notification configuration was created.
	//
	EnvironmentType string `json:"environment_type"`
	// ID for this offset notification configuration
	ID string `json:"id"`
	// The name for this offset notification configuration.
	//
	Name   string                           `json:"name"`
	Policy *LifecycleEventOffsetPolicyModel `json:"policy"`
	// Indicates this is an offset lifecycle event notification
	Type string `json:"type"`
}
