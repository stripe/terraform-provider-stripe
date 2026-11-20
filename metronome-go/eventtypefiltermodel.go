//
// File generated from our OpenAPI spec
//

package metronome

// An optional filtering rule to match the 'event_type' property of an event.
type EventTypeFilterModel struct {
	APIResource
	// A list of event types that are explicitly included in the billable metric. If specified, only events of these types will match the billable metric. Must be non-empty if present.
	InValues []string `json:"in_values,omitempty"`
	// A list of event types that are explicitly excluded from the billable metric. If specified, events of these types will not match the billable metric. Must be non-empty if present.
	NotInValues []string `json:"not_in_values,omitempty"`
}
