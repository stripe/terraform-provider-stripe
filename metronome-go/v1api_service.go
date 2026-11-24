//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v1APIService is used to invoke v1api related APIs.
type v1APIService struct {
	B   Backend
	Key string
}

// Create a new billable metric
func (c v1APIService) CreateBillableMetric(ctx context.Context, params *V1APICreateBillableMetricParams) (*V1CreateBillableMetricResponse, error) {
	if params == nil {
		params = &V1APICreateBillableMetricParams{}
	}
	params.Context = ctx
	v1createbillablemetricresponse := &V1CreateBillableMetricResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/billable-metrics", c.Key, params, v1createbillablemetricresponse)
	return v1createbillablemetricresponse, err
}

// Lists all packages with details including name, aliases, duration, and terms. To view contracts on a specific package, use the `listContractsOnPackage` endpoint.
func (c v1APIService) ListPackages(ctx context.Context, params *V1APIListPackagesParams) (*V1ListPackagesResponse, error) {
	if params == nil {
		params = &V1APIListPackagesParams{}
	}
	params.Context = ctx
	v1listpackagesresponse := &V1ListPackagesResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/packages/list", c.Key, params, v1listpackagesresponse)
	return v1listpackagesresponse, err
}
