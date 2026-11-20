//
// File generated from our OpenAPI spec
//

package metronome

type BillableMetricWithDeprecatedFieldsModel struct {
	APIResource
	// (DEPRECATED) use aggregation_type instead
	Aggregate string `json:"aggregate,omitempty"`
	// (DEPRECATED) use aggregation_key instead
	AggregateKeys   []string `json:"aggregate_keys,omitempty"`
	AggregationKey  string   `json:"aggregation_key,omitempty"`
	AggregationType string   `json:"aggregation_type,omitempty"`
	// RFC 3339 timestamp indicating when the billable metric was archived. If not provided, the billable metric is not archived.
	ArchivedAt      string                `json:"archived_at,omitempty"`
	CustomFields    map[string]string     `json:"custom_fields,omitempty"`
	EventTypeFilter *EventTypeFilterModel `json:"event_type_filter,omitempty"`
	// (DEPRECATED) use property_filters & event_type_filter instead
	Filter map[string]JSON `json:"filter,omitempty"`
	// (DEPRECATED) use group_keys instead
	GroupBy         []string               `json:"group_by,omitempty"`
	GroupKeys       [][]string             `json:"group_keys,omitempty"`
	ID              string                 `json:"id"`
	Name            string                 `json:"name"`
	PropertyFilters []*PropertyFilterModel `json:"property_filters,omitempty"`
	// The SQL query associated with the billable metric
	SQL string `json:"sql,omitempty"`
}
