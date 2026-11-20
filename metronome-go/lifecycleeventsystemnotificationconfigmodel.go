//
// File generated from our OpenAPI spec
//

package metronome

type LifecycleEventSystemNotificationConfigModel struct {
	APIResource
	// Whether or not webhook publishing for this lifecycle event is enabled
	IsEnabled bool                             `json:"is_enabled,omitempty"`
	Policy    *LifecycleEventSystemPolicyModel `json:"policy"`
	// Indicates this is a system lifecycle event notification
	Type string `json:"type"`
}
