//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v1ThresholdBillingService is used to invoke v1thresholdbilling related APIs.
type v1ThresholdBillingService struct {
	B   Backend
	Key string
}

// A unified API for updating prepaid balance threshold configs for clients who are migrating customers from Metronome 1.0 and 2.0. Customer can only have one active config across all Plans or Contracts.
func (c v1ThresholdBillingService) UpdateActiveRechargeConfig(ctx context.Context, params *V1ThresholdBillingUpdateActiveRechargeConfigParams) (*V1UpdateActiveRechargeConfigResponse, error) {
	if params == nil {
		params = &V1ThresholdBillingUpdateActiveRechargeConfigParams{}
	}
	params.Context = ctx
	v1updateactiverechargeconfigresponse := &V1UpdateActiveRechargeConfigResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/threshold-billing/update-active-recharge-config", c.Key, params, v1updateactiverechargeconfigresponse)
	return v1updateactiverechargeconfigresponse, err
}
