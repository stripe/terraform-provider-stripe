//
// File generated from our OpenAPI spec
//

package metronome

type LifecycleEventSystemPolicyModel struct {
	APIResource
	// The type of lifecycle event (e.g., "contract.create", "contract.start")
	Type string `json:"type"`
}
