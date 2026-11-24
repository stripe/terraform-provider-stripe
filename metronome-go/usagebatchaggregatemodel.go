//
// File generated from our OpenAPI spec
//

package metronome

type UsageBatchAggregateModel struct {
	APIResource
	BillableMetricID   string `json:"billable_metric_id"`
	BillableMetricName string `json:"billable_metric_name"`
	CustomerID         string `json:"customer_id"`
	EndTimestamp       string `json:"end_timestamp"`
	// Values will be either a number or null. Null indicates that there were no matches for the group_by value.
	Groups         map[string]*Number `json:"groups,omitempty"`
	StartTimestamp string             `json:"start_timestamp"`
	Value          Number             `json:"value"`
}
