//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v1NamedSchedulesService is used to invoke v1namedschedules related APIs.
type v1NamedSchedulesService struct {
	B   Backend
	Key string
}

// Get a named schedule for the given contract. This endpoint's availability is dependent on your client's configuration.
func (c v1NamedSchedulesService) GetContractNamedSchedule(ctx context.Context, params *V1NamedSchedulesGetContractNamedScheduleParams) (*V1GetContractNamedScheduleResponse, error) {
	if params == nil {
		params = &V1NamedSchedulesGetContractNamedScheduleParams{}
	}
	params.Context = ctx
	v1getcontractnamedscheduleresponse := &V1GetContractNamedScheduleResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/getNamedSchedule", c.Key, params, v1getcontractnamedscheduleresponse)
	return v1getcontractnamedscheduleresponse, err
}

// Get a named schedule for the given customer. This endpoint's availability is dependent on your client's configuration.
func (c v1NamedSchedulesService) GetCustomerNamedSchedule(ctx context.Context, params *V1NamedSchedulesGetCustomerNamedScheduleParams) (*V1GetCustomerNamedScheduleResponse, error) {
	if params == nil {
		params = &V1NamedSchedulesGetCustomerNamedScheduleParams{}
	}
	params.Context = ctx
	v1getcustomernamedscheduleresponse := &V1GetCustomerNamedScheduleResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/customers/getNamedSchedule", c.Key, params, v1getcustomernamedscheduleresponse)
	return v1getcustomernamedscheduleresponse, err
}

// Get a named schedule for the given rate card. This endpoint's availability is dependent on your client's configuration.
func (c v1NamedSchedulesService) GetRateCardNamedSchedule(ctx context.Context, params *V1NamedSchedulesGetRateCardNamedScheduleParams) (*V1GetRateCardNamedScheduleResponse, error) {
	if params == nil {
		params = &V1NamedSchedulesGetRateCardNamedScheduleParams{}
	}
	params.Context = ctx
	v1getratecardnamedscheduleresponse := &V1GetRateCardNamedScheduleResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contract-pricing/rate-cards/getNamedSchedule", c.Key, params, v1getratecardnamedscheduleresponse)
	return v1getratecardnamedscheduleresponse, err
}

// List contract-level named schedules for a customer, optionally scoped to a single contract.
func (c v1NamedSchedulesService) ListContractsNamedSchedules(ctx context.Context, params *V1NamedSchedulesListContractsNamedSchedulesParams) (*V1ListContractsNamedSchedulesResponse, error) {
	if params == nil {
		params = &V1NamedSchedulesListContractsNamedSchedulesParams{}
	}
	params.Context = ctx
	v1listcontractsnamedschedulesresponse := &V1ListContractsNamedSchedulesResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/listNamedSchedules", c.Key, params, v1listcontractsnamedschedulesresponse)
	return v1listcontractsnamedschedulesresponse, err
}

// Update a named schedule for the given contract. This endpoint's availability is dependent on your client's configuration.
func (c v1NamedSchedulesService) UpdateContractNamedSchedule(ctx context.Context, params *V1NamedSchedulesUpdateContractNamedScheduleParams) (*V1UpdateContractNamedScheduleResponse, error) {
	if params == nil {
		params = &V1NamedSchedulesUpdateContractNamedScheduleParams{}
	}
	params.Context = ctx
	v1updatecontractnamedscheduleresponse := &V1UpdateContractNamedScheduleResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/updateNamedSchedule", c.Key, params, v1updatecontractnamedscheduleresponse)
	return v1updatecontractnamedscheduleresponse, err
}

// Update a named schedule for the given customer. This endpoint's availability is dependent on your client's configuration.
func (c v1NamedSchedulesService) UpdateCustomerNamedSchedule(ctx context.Context, params *V1NamedSchedulesUpdateCustomerNamedScheduleParams) (*V1UpdateCustomerNamedScheduleResponse, error) {
	if params == nil {
		params = &V1NamedSchedulesUpdateCustomerNamedScheduleParams{}
	}
	params.Context = ctx
	v1updatecustomernamedscheduleresponse := &V1UpdateCustomerNamedScheduleResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/customers/updateNamedSchedule", c.Key, params, v1updatecustomernamedscheduleresponse)
	return v1updatecustomernamedscheduleresponse, err
}

// Update a named schedule for the given rate card. This endpoint's availability is dependent on your client's configuration.
func (c v1NamedSchedulesService) UpdateRateCardNamedSchedule(ctx context.Context, params *V1NamedSchedulesUpdateRateCardNamedScheduleParams) (*V1UpdateRateCardNamedScheduleResponse, error) {
	if params == nil {
		params = &V1NamedSchedulesUpdateRateCardNamedScheduleParams{}
	}
	params.Context = ctx
	v1updateratecardnamedscheduleresponse := &V1UpdateRateCardNamedScheduleResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contract-pricing/rate-cards/updateNamedSchedule", c.Key, params, v1updateratecardnamedscheduleresponse)
	return v1updateratecardnamedscheduleresponse, err
}
