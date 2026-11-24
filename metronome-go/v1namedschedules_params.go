//
// File generated from our OpenAPI spec
//

package metronome

// Get a named schedule for the given contract. This endpoint's availability is dependent on your client's configuration.
type V1NamedSchedulesGetContractNamedScheduleParams struct {
	Params `form:"*"`
	// ID of the contract whose named schedule is to be retrieved
	ContractID *string `form:"contract_id" json:"contract_id"`
	// If provided, at most one schedule segment will be returned (the one that covers this date). If not provided, all segments will be returned.
	CoveringDate *string `form:"covering_date" json:"covering_date,omitempty"`
	// ID of the customer whose named schedule is to be retrieved
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// A set of key-value pairs that qualifies which schedule should be applied when looking up a schedule by name.
	Properties map[string]string `form:"properties" json:"properties,omitempty"`
	// The identifier for the schedule to be retrieved
	ScheduleName *string `form:"schedule_name" json:"schedule_name"`
}

// Get a named schedule for the given customer. This endpoint's availability is dependent on your client's configuration.
type V1NamedSchedulesGetCustomerNamedScheduleParams struct {
	Params `form:"*"`
	// If provided, at most one schedule segment will be returned (the one that covers this date). If not provided, all segments will be returned.
	CoveringDate *string `form:"covering_date" json:"covering_date,omitempty"`
	// ID of the customer whose named schedule is to be retrieved
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// The identifier for the schedule to be retrieved
	ScheduleName *string `form:"schedule_name" json:"schedule_name"`
}

// Get a named schedule for the given rate card. This endpoint's availability is dependent on your client's configuration.
type V1NamedSchedulesGetRateCardNamedScheduleParams struct {
	Params `form:"*"`
	// If provided, at most one schedule segment will be returned (the one that covers this date). If not provided, all segments will be returned.
	CoveringDate *string `form:"covering_date" json:"covering_date,omitempty"`
	// ID of the rate card whose named schedule is to be retrieved
	RateCardID *string `form:"rate_card_id" json:"rate_card_id"`
	// The identifier for the schedule to be retrieved
	ScheduleName *string `form:"schedule_name" json:"schedule_name"`
}

// List contract-level named schedules for a customer, optionally scoped to a single contract.
type V1NamedSchedulesListContractsNamedSchedulesParams struct {
	Params `form:"*"`
	// Optional ID of a contract to scope the results to a single contract
	ContractID *string `form:"contract_id" json:"contract_id,omitempty"`
	// If provided, result contains at most one segment covering this date for each schedule.
	CoveringDate *string `form:"covering_date" json:"covering_date,omitempty"`
	// ID of the customer to list contract named schedules for
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// Maximum number of named schedules to return.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor for pagination. Use the value from a previous response's next_page.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
	// Optional key-value filters. A schedule matches when it contains at least this subset of properties.
	Properties map[string]string `form:"properties" json:"properties,omitempty"`
	// Optional filter to scope the results to a single schedule name
	ScheduleName *string `form:"schedule_name" json:"schedule_name,omitempty"`
}

// Update a named schedule for the given contract. This endpoint's availability is dependent on your client's configuration.
type V1NamedSchedulesUpdateContractNamedScheduleParams struct {
	Params `form:"*"`
	// ID of the contract whose named schedule is to be updated
	ContractID *string `form:"contract_id" json:"contract_id"`
	// ID of the customer whose named schedule is to be updated
	CustomerID   *string `form:"customer_id" json:"customer_id"`
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// A set of key-value pairs that qualifies which schedule should be applied when looking up a schedule by name.
	Properties map[string]string `form:"properties" json:"properties,omitempty"`
	// The identifier for the schedule to be updated
	ScheduleName *string `form:"schedule_name" json:"schedule_name"`
	StartingAt   *string `form:"starting_at" json:"starting_at"`
	// The value to set for the named schedule. The structure of this object is specific to the named schedule.
	Value *JSON `form:"value" json:"value"`
}

// Update a named schedule for the given customer. This endpoint's availability is dependent on your client's configuration.
type V1NamedSchedulesUpdateCustomerNamedScheduleParams struct {
	Params `form:"*"`
	// ID of the customer whose named schedule is to be updated
	CustomerID   *string `form:"customer_id" json:"customer_id"`
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// The identifier for the schedule to be updated
	ScheduleName *string `form:"schedule_name" json:"schedule_name"`
	StartingAt   *string `form:"starting_at" json:"starting_at"`
	// The value to set for the named schedule. The structure of this object is specific to the named schedule.
	Value *JSON `form:"value" json:"value"`
}

// Update a named schedule for the given rate card. This endpoint's availability is dependent on your client's configuration.
type V1NamedSchedulesUpdateRateCardNamedScheduleParams struct {
	Params       `form:"*"`
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// ID of the rate card whose named schedule is to be updated
	RateCardID *string `form:"rate_card_id" json:"rate_card_id"`
	// The identifier for the schedule to be updated
	ScheduleName *string `form:"schedule_name" json:"schedule_name"`
	StartingAt   *string `form:"starting_at" json:"starting_at"`
	// The value to set for the named schedule. The structure of this object is specific to the named schedule.
	Value *JSON `form:"value" json:"value"`
}
