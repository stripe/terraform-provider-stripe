//
// File generated from our OpenAPI spec
//

package metronome

// Create a new billable metric
type V1APICreateBillableMetricParams struct {
	Params       `form:"*"`
	Aggregate    *string `form:"aggregate" json:"aggregate"`
	AggregateKey *string `form:"aggregate_key" json:"aggregate_key,omitempty"`
	// JSON Schema filter to apply to the metric
	Filter      map[string]JSON `form:"filter" json:"filter"`
	GroupKeys   []*string       `form:"group_keys" json:"group_keys,omitempty"`
	GroupValues []*string       `form:"group_values" json:"group_values,omitempty"`
	Name        *string         `form:"name" json:"name"`
}
type V1APIIntegrationCloudwatchLoggerDimensionParams struct {
	Name  *string `form:"name" json:"name"`
	Value *string `form:"value" json:"value"`
}

// Log events from integration services to Cloudwatch
type V1APIIntegrationCloudwatchLoggerParams struct {
	Params      `form:"*"`
	Dimensions  []*V1APIIntegrationCloudwatchLoggerDimensionParams `form:"dimensions" json:"dimensions"`
	MetricName  *string                                            `form:"metric_name" json:"metric_name"`
	ServiceName *string                                            `form:"service_name" json:"service_name"`
}

// Lists all packages with details including name, aliases, duration, and terms. To view contracts on a specific package, use the `listContractsOnPackage` endpoint.
type V1APIListPackagesParams struct {
	Params `form:"*"`
	// Filter packages by archived status. Defaults to NOT_ARCHIVED.
	ArchiveFilter *string `form:"archive_filter" json:"archive_filter,omitempty"`
	// Max number of results that should be returned
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
}
