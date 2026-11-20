//
// File generated from our OpenAPI spec
//

package metronome

type PropertyFilterModel struct {
	APIResource
	// Determines whether the property must exist in the event. If true, only events with this property will pass the filter. If false, only events without this property will pass the filter. If null or omitted, the existence of the property is optional.
	Exists bool `json:"exists,omitempty"`
	// Specifies the allowed values for the property to match an event. An event will pass the filter only if its property value is included in this list. If undefined, all property values will pass the filter. Must be non-empty if present.
	InValues []string `json:"in_values,omitempty"`
	// The name of the event property.
	Name string `json:"name"`
	// Specifies the values that prevent an event from matching the filter. An event will not pass the filter if its property value is included in this list. If null or empty, all property values will pass the filter. Must be non-empty if present.
	NotInValues []string `json:"not_in_values,omitempty"`
}
