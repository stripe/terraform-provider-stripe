//
// File generated from our OpenAPI spec
//

package metronome

type BillableMetricV1Model struct {
	APIResource
	AggregationKey  string `json:"aggregation_key,omitempty"`
	AggregationType string `json:"aggregation_type,omitempty"`
	// RFC 3339 timestamp indicating when the billable metric was archived. If not provided, the billable metric is not archived.
	ArchivedAt      string                `json:"archived_at,omitempty"`
	CustomFields    map[string]string     `json:"custom_fields,omitempty"`
	EventTypeFilter *EventTypeFilterModel `json:"event_type_filter,omitempty"`
	GroupKeys       [][]string            `json:"group_keys,omitempty"`
	// ID of the billable metric
	ID string `json:"id"`
	// The display name of the billable metric.
	Name            string                 `json:"name"`
	PropertyFilters []*PropertyFilterModel `json:"property_filters,omitempty"`
	// The SQL query associated with the billable metric
	SQL string `json:"sql,omitempty"`
}
