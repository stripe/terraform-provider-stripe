//
// File generated from our OpenAPI spec
//

package metronome

// Use group_key and group_filters instead. Use a single group key to group by. Compound group keys are not supported.
type V1UsageGetPagedUsageGroupByParams struct {
	// The name of the group_by key to use
	Key *string `form:"key" json:"key"`
	// Values of the group_by key to return in the query. Accepts at most 200 values. Omit this if you'd like all values for the key returned.
	Values []*string `form:"values" json:"values,omitempty"`
}

// Retrieve granular usage data for a specific customer and billable metric, with the ability to break down usage by custom grouping dimensions. This endpoint enables deep usage analytics by segmenting data across attributes like region, user, model type, or any custom dimension defined in your billable metrics.
//
// ### Use this endpoint to:
// - Analyze usage patterns broken down by specific attributes (region, user, department, etc.)
// - Build detailed usage dashboards with dimensional filtering
// - Identify high-usage segments for optimization opportunities
//
// ### Request parameters
// Use [`group_key`](https://docs.stripe.com/api#body-group-key) and [`group_filters`](https://docs.stripe.com/api#body-group-filters) to group by one or more dimensions. This is required for compound group keys and recommended for all new integrations:
// ```json
//
//	{
//	  "group_key": ["region", "team"],
//	  "group_filters": {
//	    "region": ["US-East", "US-West"]
//	  }
//	}
//
// ```
// Important: For compound group keys, you must pass the complete set of keys that make up the compound key. Partial key combinations are not supported. For example, if your billable metric has a compound group key [region, team, environment], you must pass all three keys in group_key—you cannot pass just `[region]` or `[region, team]`.
//
// ### Key response fields:
// An array of `PagedUsageAggregate` objects containing:
// - `starting_on` and `ending_before`: Time window boundaries
// - `group`: Object mapping group keys to their values
//   - For simple groups, this will be a map with a single key-value pair (e.g., `{"region": "US-East"}`)
//   - For compound groups, this will be a map with multiple key-value pairs (e.g., `{"region": "US-East", "team": "engineering"}`)
//
// - `value`: Aggregated usage for this group and time window
// - `next_page`: Pagination cursor for large datasets
//
// ### Usage guidelines:
// - Required parameters: Must specify `customer_id`, `billable_metric_id`, and `window_size`
// - Time windows: Set `window_size` to hour, day, or none for different granularities
// - Group filtering: Use `group_key` and `group_filters` to specify groups and group filters. Across all arrays in `group_filters`, include at most 200 filter values total. Requests with more than 200 filter values are rejected when this limit is enforced
// - Response limit: When using compound group keys (2+ keys in `group_key`), the default and maximum page size is 100
// - Pagination: Use limit and `next_page` for large result sets
// - Null handling: Group values may be null for events missing the group key property
type V1UsageGetPagedUsageParams struct {
	Params           `form:"*"`
	BillableMetricID *string `form:"billable_metric_id" json:"billable_metric_id"`
	CustomerID       *string `form:"customer_id" json:"customer_id"`
	// Must be aligned to UTC midnight and at least one day after `starting_on`.
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// Use group_key and group_filters instead. Use a single group key to group by. Compound group keys are not supported.
	GroupBy      *V1UsageGetPagedUsageGroupByParams `form:"group_by" json:"group_by,omitempty"`
	GroupFilters map[string][]*string               `form:"group_filters" json:"group_filters,omitempty"`
	GroupKey     []*string                          `form:"group_key" json:"group_key,omitempty"`
	// Max number of results that should be returned
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
	// Must be aligned to UTC midnight, e.g. `2024-01-01T00:00:00Z`.
	StartingOn *string `form:"starting_on" json:"starting_on,omitempty"`
	// A window_size of "day" or "hour" will return the usage for the specified period segmented into daily or hourly aggregates. A window_size of "none" will return a single usage aggregate for the entirety of the specified period.
	WindowSize *string `form:"window_size" json:"window_size"`
}

// Returns the number of seats purchased for the current period.
type V1UsageGetPurchasedSeatsParams struct {
	Params `form:"*"`
}

// Fetch seats usage data for the specified customer, broken into intervals of the specified length.
type V1UsageGetSeatsUsageParams struct {
	Params     `form:"*"`
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// Must be aligned to UTC midnight and at least one day after `starting_on`.
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// Max number of results that should be returned
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor that indicates where the next page of results should start.
	NextPage     *string `form:"next_page" json:"next_page,omitempty"`
	SeatMetricID *string `form:"seat_metric_id" json:"seat_metric_id"`
	// Must be aligned to UTC midnight, e.g. `2024-01-01T00:00:00Z`.
	StartingOn *string `form:"starting_on" json:"starting_on,omitempty"`
	// The type of usage to return, max or latest value.
	UsageType *string `form:"usage_type" json:"usage_type"`
	// A window_size of "day" or "hour" will return the usage for the specified period segmented into daily or hourly aggregates. A window_size of "none" will return a single usage aggregate for the entirety of the specified period.
	WindowSize *string `form:"window_size" json:"window_size"`
}
type V1UsageGetUsageBatchBillableMetricGroupByParams struct {
	// The name of the group_by key to use
	Key *string `form:"key" json:"key"`
	// Values of the group_by key to return in the query. If this field is omitted, all available values will be returned, up to a maximum of 200.
	Values []*string `form:"values" json:"values,omitempty"`
}

// A list of billable metrics to fetch usage for. If absent, all billable metrics will be returned.
type V1UsageGetUsageBatchBillableMetricParams struct {
	GroupBy *V1UsageGetUsageBatchBillableMetricGroupByParams `form:"group_by" json:"group_by,omitempty"`
	ID      *string                                          `form:"id" json:"id"`
}

// Retrieve aggregated usage data across multiple customers and billable metrics in a single query. This batch endpoint enables you to fetch usage patterns at scale, broken down by time windows, making it ideal for building analytics dashboards, generating reports, and monitoring platform-wide usage trends.
//
// ### Use this endpoint to:
// - Generate platform-wide usage reports for internal teams
// - Monitor aggregate usage trends across your entire customer base
// - Create comparative usage analyses between customers or time periods
// - Support capacity planning with historical usage patterns
//
// ### Key response fields:
// An array of `UsageBatchAggregate` objects containing:
//
// - `customer_id`: The customer this usage belongs to
// - `billable_metric_id` and `billable_metric_name`: What was measured
// - `start_timestamp` and `end_timestamp`: Time window for this data point
// - `value`: Aggregated usage amount for the period
// - `groups` (optional): Usage broken down by group keys with values
// - `next_page`: Pagination cursor for large result sets
//
// ### Usage guidelines:
// - Time windows: Set `window_size` to `hour`, `day`, or `none` (entire period)
// - Required parameters: Must specify `starting_on`, `ending_before`, and `window_size`
// - Filtering options:
//   - `customer_ids`: Limit to specific customers (omit for all customers)
//   - `billable_metrics`: Limit to specific metrics (omit for all metrics)
//
// - Pagination: Use `next_page` cursor to retrieve large datasets
// - Null values: Group values may be null when no usage matches that group
type V1UsageGetUsageBatchParams struct {
	Params `form:"*"`
	// A list of billable metrics to fetch usage for. If absent, all billable metrics will be returned.
	BillableMetrics []*V1UsageGetUsageBatchBillableMetricParams `form:"billable_metrics" json:"billable_metrics,omitempty"`
	// A list of Metronome customer IDs to fetch usage for. If absent, usage for all customers will be returned.
	CustomerIDs []*string `form:"customer_ids" json:"customer_ids,omitempty"`
	// Must be aligned to UTC midnight and at least one day after `starting_on`.
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
	// Must be aligned to UTC midnight, e.g. `2024-01-01T00:00:00Z`.
	StartingOn *string `form:"starting_on" json:"starting_on"`
	// A window_size of "day" or "hour" will return the usage for the specified period segmented into daily or hourly aggregates. A window_size of "none" will return a single usage aggregate for the entirety of the specified period.
	WindowSize *string `form:"window_size" json:"window_size"`
}
type V1UsageIngestBodyParams struct {
	CustomerID *string         `form:"customer_id" json:"customer_id"`
	EventType  *string         `form:"event_type" json:"event_type"`
	Properties map[string]JSON `form:"properties" json:"properties,omitempty"`
	// RFC 3339 formatted
	Timestamp     *string `form:"timestamp" json:"timestamp"`
	TransactionID *string `form:"transaction_id" json:"transaction_id"`
}

// The ingest endpoint is the primary method for sending usage events to Metronome, serving as the foundation for all billing calculations in your usage-based pricing model. This high-throughput endpoint is designed for real-time streaming ingestion, supports backdating 34 days, and is built to handle mission-critical usage data with enterprise-grade reliability. Metronome supports 100,000 events per second without requiring pre-aggregation or rollups and can scale up from there. See the [Send usage events](https://docs.stripe.com/guides/events/send-usage-events) guide to learn more about usage events.
//
// ### Use this endpoint to:
// Create a customer usage pipeline into Metronome that drives billable metrics, credit drawdown, and invoicing. Track customer behavior, resource consumption, and feature usage
//
// ### What happens when you send events:
// - Events are validated and processed in real-time
// - Events are matched to customers using customer IDs or customer ingest aliases
// - Events are matched to billable metrics and are immediately available for usage and spend calculations
//
// ### Usage guidelines:
// - Historical events can be backdated up to 34 days and will immediately impact live customer spend
// - Duplicate events are automatically detected and ignored (34-day deduplication window)
//
// #### Event structure:
// Usage events are simple JSON objects designed for flexibility and ease of integration:
// ```json
//
//	{
//	  "transaction_id": "2021-01-01T00:00:00Z_cluster42",
//	  "customer_id": "team@example.com",
//	  "event_type": "api_request",
//	  "timestamp": "2021-01-01T00:00:00Z",
//	  "properties": {
//	    "endpoint": "/v1/users",
//	    "method": "POST",
//	    "response_time_ms": 45,
//	    "region": "us-west-2"
//	  }
//	}
//
// ```
//
// Learn more about [usage event structure definitions](https://docs.stripe.com/guides/events/design-usage-events).
//
// #### Transaction ID
//
//	The transaction_id serves as your idempotency key, ensuring events are processed exactly once. Metronome maintains a 34-day deduplication window - significantly longer than typical 12-hour windows - enabling robust backfill scenarios without duplicate billing.
//	- Best Practices:
//	  - Use UUIDs for one-time events: uuid4()
//	  - For heartbeat events, use deterministic IDs
//	  - Include enough context to avoid collisions across different event sources
//
// #### Customer ID
// Identifies which customer should be billed for this usage. Supports two identification methods:
//   - Metronome Customer ID: The UUID returned when creating a customer
//   - Ingest Alias: Your system's identifier (email, account number, etc.)
//
// Ingest aliases enable seamless integration without requiring ID mapping, and customers can have multiple aliases for flexibility.
//
// #### Event Type:
// Categorizes the event type for billable metric matching. Choose descriptive names that aligns with the product surface area.
//
// #### Properties:
// Flexible metadata also used to match billable metrics or to be used to serve as group keys to create multiple pricing dimensions or breakdown costs by novel properties for end customers or internal finance teams measuring underlying COGs.
type V1UsageIngestParams struct {
	Params `form:"*"`
	Body   []*V1UsageIngestBodyParams `form:"body" json:"body"`
}

// This endpoint retrieves events by transaction ID for events that occurred within the last 34 days. It is specifically designed for sampling-based testing workflows to detect revenue leakage. The Event Search API provides a critical observability tool that validates the integrity of your usage pipeline by allowing you to sample raw events and verify their matching against active billable metrics.
//
// Why event observability matters for revenue leakage:
// Silent revenue loss occurs when events are dropped, delayed, or fail to match billable metrics due to:
// - Upstream system failures
// - Event format changes
// - Misconfigured billable metrics
//
// ### Use this endpoint to:
// - Sample raw events and validate they match the expected billable metrics
// - Build custom leakage detection alerts to prevent silent revenue loss
// - Verify event processing accuracy during system changes or metric updates
// - Debug event matching issues in real-time
//
// ### Key response fields:
// - Complete event details including transaction ID, customer ID, and properties
// - Matched Metronome customer (if any)
// - Matched billable metric information (if any)
// - Processing status and duplicate detection flags
//
// ### Usage guidelines:
// ⚠️ Important: This endpoint is heavily rate limited and designed for sampling workflows only. Do not use this endpoint to check every event in your system. Instead, implement a sampling strategy to randomly validate a subset of events for observability purposes.
type V1UsageSearchEventsParams struct {
	Params `form:"*"`
	// The transaction IDs of the events to retrieve
	TransactionIDs []*string `form:"transactionIds" json:"transactionIds"`
}
