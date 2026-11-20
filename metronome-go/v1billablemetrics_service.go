//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v1BillableMetricsService is used to invoke v1billablemetrics related APIs.
type v1BillableMetricsService struct {
	B   Backend
	Key string
}

// Use this endpoint to retire billable metrics that are no longer used. After a billable metric is archived, that billable metric can no longer be used in any new Products to define how that product should be metered. If you archive a billable metric that is already associated with a Product, the Product will continue to function as usual, metering based on the definition of the archived billable metric.
//
// Archived billable metrics will be returned on the `getBillableMetric` and `listBillableMetrics` endpoints with a populated `archived_at` field.
func (c v1BillableMetricsService) ArchiveBillableMetric(ctx context.Context, params *V1BillableMetricsArchiveBillableMetricParams) (*V1ArchiveBillableMetricResponse, error) {
	if params == nil {
		params = &V1BillableMetricsArchiveBillableMetricParams{}
	}
	params.Context = ctx
	v1archivebillablemetricresponse := &V1ArchiveBillableMetricResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/billable-metrics/archive", c.Key, params, v1archivebillablemetricresponse)
	return v1archivebillablemetricresponse, err
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
func (c v1BillableMetricsService) CreateBillableMetricV1(ctx context.Context, params *V1BillableMetricsCreateBillableMetricV1Params) (*V1CreateBillableMetricV1Response, error) {
	if params == nil {
		params = &V1BillableMetricsCreateBillableMetricV1Params{}
	}
	params.Context = ctx
	v1createbillablemetricv1response := &V1CreateBillableMetricV1Response{}
	err := c.B.Call(
		http.MethodPost, "/v1/billable-metrics/create", c.Key, params, v1createbillablemetricv1response)
	return v1createbillablemetricv1response, err
}

// Retrieves the complete configuration for a specific billable metric by its ID. Use this to review billable metric setup before associating it with products. Returns the metric's `name`, `event_type_filter`, `property_filters`, `aggregation_type`, `aggregation_key`, `group_keys`, `custom fields`, and `SQL query` (if it's a SQL billable metric).
//
// Important:
// - Archived billable metrics will include an `archived_at` timestamp; they no longer process new usage events but remain accessible for historical reference.
func (c v1BillableMetricsService) GetBillableMetric(ctx context.Context, id string, params *V1BillableMetricsGetBillableMetricParams) (*V1GetBillableMetricResponse, error) {
	if params == nil {
		params = &V1BillableMetricsGetBillableMetricParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billable-metrics/%s", id)
	v1getbillablemetricresponse := &V1GetBillableMetricResponse{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, v1getbillablemetricresponse)
	return v1getbillablemetricresponse, err
}

// Retrieves all billable metrics with their complete configurations. Use this for programmatic discovery and management of billable metrics, such as associating metrics to products and auditing for orphaned or archived metrics.
// Important: Archived metrics are excluded by default; use `include_archived`=`true` parameter to include them.
func (c v1BillableMetricsService) ListAllBillableMetrics(ctx context.Context, params *V1BillableMetricsListAllBillableMetricsParams) (*V1ListAllBillableMetricsResponse, error) {
	if params == nil {
		params = &V1BillableMetricsListAllBillableMetricsParams{}
	}
	params.Context = ctx
	v1listallbillablemetricsresponse := &V1ListAllBillableMetricsResponse{}
	err := c.B.Call(
		http.MethodGet, "/v1/billable-metrics", c.Key, params, v1listallbillablemetricsresponse)
	return v1listallbillablemetricsresponse, err
}

// Get all billable metrics available for a specific customer. Supports pagination and filtering by current plan status or archived metrics. Use this endpoint to see which metrics are being tracked for billing calculations for a given customer.
func (c v1BillableMetricsService) ListBillableMetrics(ctx context.Context, id string, params *V1BillableMetricsListBillableMetricsParams) (*V1ListBillableMetricsResponse, error) {
	if params == nil {
		params = &V1BillableMetricsListBillableMetricsParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/customers/%s/billable-metrics", id)
	v1listbillablemetricsresponse := &V1ListBillableMetricsResponse{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, v1listbillablemetricsresponse)
	return v1listbillablemetricsresponse, err
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
func (c v1BillableMetricsService) UpdateBillableMetric(ctx context.Context, id string, params *V1BillableMetricsUpdateBillableMetricParams) (*V1UpdateBillableMetricResponse, error) {
	if params == nil {
		params = &V1BillableMetricsUpdateBillableMetricParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billable-metrics/%s", id)
	v1updatebillablemetricresponse := &V1UpdateBillableMetricResponse{}
	err := c.B.Call(
		http.MethodPut, path, c.Key, params, v1updatebillablemetricresponse)
	return v1updatebillablemetricresponse, err
}
