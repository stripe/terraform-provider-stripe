//
// File generated from our OpenAPI spec
//

package metronome

// Use this endpoint to retire billable metrics that are no longer used. After a billable metric is archived, that billable metric can no longer be used in any new Products to define how that product should be metered. If you archive a billable metric that is already associated with a Product, the Product will continue to function as usual, metering based on the definition of the archived billable metric.
//
// Archived billable metrics will be returned on the `getBillableMetric` and `listBillableMetrics` endpoints with a populated `archived_at` field.
type V1BillableMetricsArchiveBillableMetricParams struct {
	Params `form:"*"`
	ID     *string `form:"id" json:"id"`
}
type V1BillableMetricsCreateBillableMetricV1EventTypeFilterParams struct {
	// A list of event types that are explicitly included in the billable metric. If specified, only events of these types will match the billable metric. Must be non-empty if present.
	InValues []*string `form:"in_values" json:"in_values,omitempty"`
	// A list of event types that are explicitly excluded from the billable metric. If specified, events of these types will not match the billable metric. Must be non-empty if present.
	NotInValues []*string `form:"not_in_values" json:"not_in_values,omitempty"`
}
type V1BillableMetricsCreateBillableMetricV1PropertyFilterParams struct {
	// Determines whether the property must exist in the event. If true, only events with this property will pass the filter. If false, only events without this property will pass the filter. If null or omitted, the existence of the property is optional.
	Exists *bool `form:"exists" json:"exists,omitempty"`
	// Specifies the allowed values for the property to match an event. An event will pass the filter only if its property value is included in this list. If undefined, all property values will pass the filter. Must be non-empty if present.
	InValues []*string `form:"in_values" json:"in_values,omitempty"`
	// The name of the event property.
	Name *string `form:"name" json:"name"`
	// Specifies the values that prevent an event from matching the filter. An event will not pass the filter if its property value is included in this list. If null or empty, all property values will pass the filter. Must be non-empty if present.
	NotInValues []*string `form:"not_in_values" json:"not_in_values,omitempty"`
}

// Create billable metrics programmatically with this endpoint—an essential step in configuring your pricing and packaging in Metronome.
//
// A billable metric is a customizable query that filters and aggregates events from your event stream. These metrics are continuously tracked as usage data enters Metronome through the ingestion pipeline. The ingestion process transforms raw usage data into actionable pricing metrics, enabling accurate metering and billing for your products.
//
// ### Use this endpoint to:
// - Create individual or multiple billable metrics as part of a setup workflow.
// - Automate the entire pricing configuration process, from metric creation to customer contract setup.
// - Define metrics using either standard filtering/aggregation or a custom SQL query.
//
// ### Key response fields:
// - The ID of the billable metric that was created
// - The created billable metric will be available to be used in Products, usage endpoints, and alerts.
//
// ### Usage guidelines:
// - Metrics defined using standard filtering and aggregation are Streaming billable metrics, which have been optimized for ultra low latency and high throughput workflows.
// - Use SQL billable metrics if you require more flexible aggregation options.
type V1BillableMetricsCreateBillableMetricV1Params struct {
	Params `form:"*"`
	// Specifies the type of aggregation performed on matching events. Required if `sql` is not provided.
	AggregationKey  *string `form:"aggregation_key" json:"aggregation_key,omitempty"`
	AggregationType *string `form:"aggregation_type" json:"aggregation_type,omitempty"`
	// Custom fields to attach to the billable metric.
	CustomFields    map[string]string                                             `form:"custom_fields" json:"custom_fields,omitempty"`
	EventTypeFilter *V1BillableMetricsCreateBillableMetricV1EventTypeFilterParams `form:"event_type_filter" json:"event_type_filter,omitempty"`
	GroupKeys       []*[]*string                                                  `form:"group_keys" json:"group_keys,omitempty"`
	// The display name of the billable metric.
	Name            *string                                                        `form:"name" json:"name"`
	PropertyFilters []*V1BillableMetricsCreateBillableMetricV1PropertyFilterParams `form:"property_filters" json:"property_filters,omitempty"`
	// The SQL query associated with the billable metric. This field is mutually exclusive with aggregation_type, event_type_filter, property_filters, aggregation_key, and group_keys. If provided, these other fields must be omitted.
	SQL *string `form:"sql" json:"sql,omitempty"`
}

// Retrieves the complete configuration for a specific billable metric by its ID. Use this to review billable metric setup before associating it with products. Returns the metric's `name`, `event_type_filter`, `property_filters`, `aggregation_type`, `aggregation_key`, `group_keys`, `custom fields`, and `SQL query` (if it's a SQL billable metric).
//
// Important:
// - Archived billable metrics will include an `archived_at` timestamp; they no longer process new usage events but remain accessible for historical reference.
type V1BillableMetricsGetBillableMetricParams struct {
	Params `form:"*"`
}

// Retrieves all billable metrics with their complete configurations. Use this for programmatic discovery and management of billable metrics, such as associating metrics to products and auditing for orphaned or archived metrics.
// Important: Archived metrics are excluded by default; use `include_archived`=`true` parameter to include them.
type V1BillableMetricsListAllBillableMetricsParams struct {
	Params `form:"*"`
	// If true, the list of returned metrics will include archived metrics
	IncludeArchived *bool `form:"include_archived" json:"include_archived,omitempty"`
	// Max number of results that should be returned
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
}

// Get all billable metrics available for a specific customer. Supports pagination and filtering by current plan status or archived metrics. Use this endpoint to see which metrics are being tracked for billing calculations for a given customer.
type V1BillableMetricsListBillableMetricsParams struct {
	Params `form:"*"`
	// If true, the list of returned metrics will include archived metrics
	IncludeArchived *bool `form:"include_archived" json:"include_archived,omitempty"`
	// Max number of results that should be returned
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
	// If true, the list of metrics will be filtered to just ones that are on the customer's current plan
	OnCurrentPlan *bool `form:"on_current_plan" json:"on_current_plan,omitempty"`
}

// Updates only the display name of an existing billable metric. Use this to correct mistakes or apply standardized naming conventions across all billable metrics. Returns the billable metric ID to confirm the update.
//
// Important: Only the name can be modified via this endpoint; configurations cannot be changed after creation.
//
// #### Example workflow:
// If you need to make changes to a streaming billable metric, for example, Metronome supports easily rolling out these changes using a simple workflow:
// 1. Duplicate the billable metric
// 2. Make required changes
// 3. Save the metric
// 4. Navigate to the product you have associated with the incorrect metric
// 5. Schedule the product to reference the newly created metric on the appropriate date
type V1BillableMetricsUpdateBillableMetricParams struct {
	Params `form:"*"`
	// The new name of the metric
	Name *string `form:"name" json:"name"`
}
