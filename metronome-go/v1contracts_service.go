//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v1ContractsService is used to invoke v1contracts related APIs.
type v1ContractsService struct {
	B   Backend
	Key string
}

// Amendments will be replaced by Contract editing. New clients should implement using the `editContract` endpoint. Read more about the migration to contract editing [here](https://docs.stripe.com/guides/implement-metronome/migrate-amendments-to-edits/) and contact us via the [Metronome support portal](https://support.metronome.com/) for more details. Once contract editing is enabled, access to this endpoint will be removed.
func (c v1ContractsService) AmendContract(ctx context.Context, params *V1ContractsAmendContractParams) (*V1AmendContractResponse, error) {
	if params == nil {
		params = &V1ContractsAmendContractParams{}
	}
	params.Context = ctx
	v1amendcontractresponse := &V1AmendContractResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/amend", c.Key, params, v1amendcontractresponse)
	return v1amendcontractresponse, err
}

// Permanently end and archive a contract along with all its terms. Any draft invoices will be canceled, and all upcoming scheduled invoices will be voided–also all finalized invoices can optionally be voided. Use this in the event a contract was incorrectly created and needed to be removed from a customer.
//
// #### Impact on commits and credits:
// When archiving a contract, all associated commits and credits are also archived. For prepaid commits with active segments, Metronome automatically generates expiration ledger entries to close out any remaining balances, ensuring accurate accounting of unused prepaid amounts. These ledger entries will appear in the commit's transaction history with type `PREPAID_COMMIT_EXPIRATION`.
//
// #### Archived contract visibility:
// Archived contracts remain accessible for historical reporting and audit purposes. They can be retrieved using the `ListContracts` endpoint by setting the `include_archived` parameter to `true` or in the Metronome UI when the "Show archived" option is enabled.
func (c v1ContractsService) ArchiveContract(ctx context.Context, params *V1ContractsArchiveContractParams) (*V1ArchiveContractResponse, error) {
	if params == nil {
		params = &V1ContractsArchiveContractParams{}
	}
	params.Context = ctx
	v1archivecontractresponse := &V1ArchiveContractResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/archive", c.Key, params, v1archivecontractresponse)
	return v1archivecontractresponse, err
}

// Contracts define a customer's products, pricing, discounts, access duration, and billing configuration. Contracts serve as the central billing agreement for both PLG and Enterprise customers. You can automatically grant customers access to your products and services directly from your product or CRM.
//
// ### Use this endpoint to:
// - PLG onboarding: Automatically provision new self-serve customers with contracts when they sign up.
// - Enterprise sales: Push negotiated contracts from Salesforce with custom pricing and commitments
// - Promotional pricing: Implement time-limited discounts and free trials through overrides
//
// ### Key components:
// #### Contract Term and Billing Schedule
// - Set contract duration using `starting_at` and `ending_before` fields. PLG contracts typically use perpetual agreements (no end date), while Enterprise contracts have fixed end dates which can be edited over time in the case of co-term upsells.
//
// #### Rate Card
// If you are offering usage based pricing, you can set a rate card for the contract to reference through `rate_card_id` or `rate_card_alias`. The rate card is a store of all of your usage based products and their centralized pricing. Any new products or price changes on the rate card can be set to automatically propagate to all associated contracts - this ensures consistent pricing and product launches flow to contracts without manual updates and migrations. The `usage_statement_schedule` determines the cadence on which Metronome will finalize a usage invoice for the customer. This defaults to monthly on the 1st, with options for custom dates, quarterly, or annual cadences. Note: Most usage based billing companies align usage statements to be evaluated aligned to the first of the month.
// Read more about [Rate Cards](https://docs.metronome.com/pricing-packaging/create-manage-rate-cards/).
//
// #### Overrides and discounts
// Customize pricing on the contract through time-bounded overrides that can target specific products, product families, or complex usage scenarios. Overrides enable two key capabilities:
// - Discounts: Apply percentage discounts, fixed rate reductions, or quantity-based pricing tiers
// - Entitlements: Provide special pricing or access to specific products for negotiated deals
//
// Read more about [Contract Overrides](https://docs.metronome.com/manage-product-access/add-contract-override/).
//
// #### Commits and Credits
// Using commits, configure prepaid or postpaid spending commitments where customers promise to spend a certain amount over the contract period paid in advance or in arrears. Use credits to provide free spending allowances. Under the hood these are the same mechanisms, however, credits are typically offered for free (SLA or promotional) or as a part of an allotment associated with a Subscription.
//
// In Metronome, you can set commits and credits to only be applicable for a subset of usage. Use `applicable_product_ids` or `applicable_product_tags` to create product or product-family specific commits or credits, or you can build complex boolean logic specifiers to target usage based on pricing  and presentation group values using `override_specifiers`.
//
// These objects can also also be configured to have a recurrence schedule to easily model customer packaging which includes recurring monthly or quarterly allotments.
//
// Commits support rollover settings (`rollover_fraction`) to transfer unused balances between contract periods, either entirely or as a percentage.
//
// Read more about [Credits and Commits](https://docs.metronome.com/pricing-packaging/apply-credits-commits/).
//
// #### Subscriptions
// You can add a fixed recurring charge to a contract, like monthly licenses or seat-based fees, using the subscription charge. Subscription charges are defined on your rate card and you can select which subscription is applicable to add to each contract.         When you add a subscription to a contract you need to:
// - Define whether the subscription is paid for in-advance or in-arrears (`collection_schedule`)
// - Define the proration behavior (`proration`)
// - Specify an initial quantity (`initial_quantity`)
// - Define which subscription rate on the rate card should be used (`subscription_rate`)
//
// Read more about [Subscriptions](https://docs.metronome.com/manage-product-access/create-subscription/).
//
// #### Scheduled Charges
// Set up one-time, recurring, or entirely custom charges that occur on specific dates, separate from usage-based billing or commitments. These can be used to model non-recurring platform charges or professional services.
//
// #### Threshold Billing
// Metronome allows you to configure automatic billing triggers when customers reach spending thresholds to prevent fraud and manage risk. You can use `spend_threshold_configuration` to trigger an invoice to cover current charges whenever the threshold is reached or you can ensure the customer maintains a minimum prepaid balance using the `prepaid_balance_configuration`.
//
// Read more about [Spend Threshold](https://docs.metronome.com/manage-product-access/spend-thresholds/) and [Prepaid Balance Thresholds](https://docs.metronome.com/manage-product-access/prepaid-balance-thresholds/).
//
// ### Usage guidelines:
// - You can always [Edit Contracts](https://docs.metronome.com/manage-product-access/edit-contract/) after it has been created, using the `editContract` endpoint. Metronome keeps track of all edits, both in the audit log and over the `getEditHistory` endpoint.
// - Customers in Metronome can have multiple concurrent contracts at one time. Use `usage_filters` to route the correct usage to each contract. [Read more about usage filters](https://docs.metronome.com/manage-product-access/provision-customer/#create-a-usage-filter).
func (c v1ContractsService) CreateContract(ctx context.Context, params *V1ContractsCreateContractParams) (*V1CreateContractResponse, error) {
	if params == nil {
		params = &V1ContractsCreateContractParams{}
	}
	params.Context = ctx
	v1createcontractresponse := &V1CreateContractResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/create", c.Key, params, v1createcontractresponse)
	return v1createcontractresponse, err
}

// Create a new customer and provision a contract. This endpoint's availability is dependent on your client's configuration.
func (c v1ContractsService) CreateCustomerWithContract(ctx context.Context, params *V1ContractsCreateCustomerWithContractParams) (*V1CreateCustomerWithContractResponse, error) {
	if params == nil {
		params = &V1ContractsCreateCustomerWithContractParams{}
	}
	params.Context = ctx
	v1createcustomerwithcontractresponse := &V1CreateCustomerWithContractResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/composite/createCustomerWithContract", c.Key, params, v1createcustomerwithcontractresponse)
	return v1createcustomerwithcontractresponse, err
}

// Create historical usage invoices for past billing periods on specific contracts. Use this endpoint to generate retroactive invoices with custom usage line items, quantities, and date ranges. Supports preview mode to validate invoice data before creation. Ideal for billing migrations or correcting past billing periods.
func (c v1ContractsService) CreateHistoricalContractUsageInvoices(ctx context.Context, params *V1ContractsCreateHistoricalContractUsageInvoicesParams) (*V1CreateHistoricalContractUsageInvoicesResponse, error) {
	if params == nil {
		params = &V1ContractsCreateHistoricalContractUsageInvoicesParams{}
	}
	params.Context = ctx
	v1createhistoricalcontractusageinvoicesresponse := &V1CreateHistoricalContractUsageInvoicesResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/createHistoricalInvoices", c.Key, params, v1createhistoricalcontractusageinvoicesresponse)
	return v1createhistoricalcontractusageinvoicesresponse, err
}

// Create a package that defines a set of reusable, time-relative contract terms that can be used across cohorts of customers. Packages provide an abstraction layer on top of rate cards to provide an easy way to provision customers with standard pricing.
//
// ### **Use this endpoint to:**
// - Model standard pay-as-you-go pricing packages that can be easily reused across customers
// - Define standardized contract terms and discounting for sales-led motions
// - Set aliases for the package to facilitate easy package transition. Aliases are human-readable names that you can use in the place of the id of the package when provisioning a customer's contract. By using an alias, you can easily create a contract and provision a customer by choosing the “Starter Plan” package, without storing the package ID in your internal systems. This is helpful when launching terms for a package, as you can create a new package with the “Starter Plan” alias scheduled to be assigned without updating your provisioning code.
//
// ### Key input fields:
// - `starting_at_offset`: Starting date relative to contract start. Generates the `starting_at` date when a contract is provisioned using a package.
// - `duration`: Duration starting from `starting_at_offset`. Generates the `ending_before` date when a contract is provisioned using a package.
// - `date_offset`: Date relative to contract start. Used for point-in-time dates without a duration.
// - `aliases`: Human-readable name to use when provisioning contracts with a package.
//
// ### Usage guidelines:
// - Use packages for standard self-serve use cases where customers have consistent terms. For customers with negotiated custom contract terms, use the `createContract` endpoint for maximum flexibility.
// - Billing provider configuration can be set when creating a package by using `billing_provider` and `delivery_method`. To provision a customer successfully with a package, the customer must have one and only one billing provider configuration that matches the billing provider configuration set on the package.
// - A package alias can only be used by one package at a time. If you create a new package with an alias that is already in use by another package, the original package's alias schedule will be updated. The alias will reference the package to which it was most recently assigned.
// - Terms can only be specified using times relative to the contract start date. Supported granularities are: `days`, `weeks`, `months`, `years`
// - Packages cannot be edited once created. Use the rate card to easily add new rates across all of your customers or make direct edits to a contract after provisioning with a package. Edited contracts will still be associated with the package used during provisioning.
func (c v1ContractsService) CreatePackage(ctx context.Context, params *V1ContractsCreatePackageParams) (*V1CreatePackageResponse, error) {
	if params == nil {
		params = &V1ContractsCreatePackageParams{}
	}
	params.Context = ctx
	v1createpackageresponse := &V1CreatePackageResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/packages/create", c.Key, params, v1createpackageresponse)
	return v1createpackageresponse, err
}

// This is the v1 endpoint to get a contract. New clients should implement using the v2 endpoint.
func (c v1ContractsService) GetContract(ctx context.Context, params *V1ContractsGetContractParams) (*V1GetContractResponse, error) {
	if params == nil {
		params = &V1ContractsGetContractParams{}
	}
	params.Context = ctx
	v1getcontractresponse := &V1GetContractResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/get", c.Key, params, v1getcontractresponse)
	return v1getcontractresponse, err
}

// For a specific customer and contract, get the rates at a specific point in time. This endpoint takes the contract's rate card into consideration, including scheduled changes. It also takes into account overrides on the contract.
//
// For example, if you want to show your customer a summary of the prices they are paying, inclusive of any negotiated discounts or promotions, use this endpoint. This endpoint only returns rates that are entitled.
func (c v1ContractsService) GetContractRateSchedule(ctx context.Context, params *V1ContractsGetContractRateScheduleParams) (*V1GetContractRateScheduleResponse, error) {
	if params == nil {
		params = &V1ContractsGetContractRateScheduleParams{}
	}
	params.Context = ctx
	v1getcontractratescheduleresponse := &V1GetContractRateScheduleResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/getContractRateSchedule", c.Key, params, v1getcontractratescheduleresponse)
	return v1getcontractratescheduleresponse, err
}

// Get the history of subscription quantities and prices over time for a given `subscription_id`. This endpoint can be used to power an in-product experience where you show a customer their historical changes to seat count. Future changes are not included in this endpoint - use the `getContract` endpoint to view the future scheduled changes to a subscription's quantity.
//
// Subscriptions are used to model fixed recurring fees as well as seat-based recurring fees. To model changes to the number of seats in Metronome, you can increment or decrement the quantity on a subscription at any point in the past or future.
func (c v1ContractsService) GetSubscriptionQuantityHistory(ctx context.Context, params *V1ContractsGetSubscriptionQuantityHistoryParams) (*V1GetSubscriptionQuantityHistoryResponse, error) {
	if params == nil {
		params = &V1ContractsGetSubscriptionQuantityHistoryParams{}
	}
	params.Context = ctx
	v1getsubscriptionquantityhistoryresponse := &V1GetSubscriptionQuantityHistoryResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/getSubscriptionQuantityHistory", c.Key, params, v1getsubscriptionquantityhistoryresponse)
	return v1getsubscriptionquantityhistoryresponse, err
}

// Get the history of subscription seats schedule over time for a given `subscription_id`. This endpoint provides information about seat assignments and total quantities for different time periods, allowing you to track how seat assignments have changed over time.
//
// ### Use this endpoint to:
// - Track changes to seat assignments over time
// - Get seat schedule for a specific date using the `covering_date` parameter
// - Get seat schedule history with optional date range filtering using `starting_at` and `ending_before`
//
// ### Key response fields:
// - data: array of seat schedule entries with time periods, quantity, and assignments
// - next_page: cursor for pagination to retrieve additional results
//
// ### Usage guidelines:
// - Use `covering_date` to get the active seats for a specific point in time. `covering_date` cannot be used with `starting_at` or `ending_before`.
// - Use `starting_at` and `ending_before` to filter results by time range. `starting_at` and `ending_before` cannot be used with `covering_date`.
// - Maximum limit is 10 seat schedule entries per request
// - Results are ordered by `starting_at` timestamp
func (c v1ContractsService) GetSubscriptionSeatsHistory(ctx context.Context, params *V1ContractsGetSubscriptionSeatsHistoryParams) (*V1GetSubscriptionSeatsHistoryResponse, error) {
	if params == nil {
		params = &V1ContractsGetSubscriptionSeatsHistoryParams{}
	}
	params.Context = ctx
	v1getsubscriptionseatshistoryresponse := &V1GetSubscriptionSeatsHistoryResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/getSubscriptionSeatsHistory", c.Key, params, v1getsubscriptionseatshistoryresponse)
	return v1getsubscriptionseatshistoryresponse, err
}

// Get the current payment and invoice status for a payment-gated subscription.
func (c v1ContractsService) GetSubscriptionStatus(ctx context.Context, params *V1ContractsGetSubscriptionStatusParams) (*V1GetSubscriptionStatusResponse, error) {
	if params == nil {
		params = &V1ContractsGetSubscriptionStatusParams{}
	}
	params.Context = ctx
	v1getsubscriptionstatusresponse := &V1GetSubscriptionStatusResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/subscriptions/getStatus", c.Key, params, v1getsubscriptionstatusresponse)
	return v1getsubscriptionstatusresponse, err
}

// Retrieves a page of contracts for a specific customer, including pricing, terms, credits, and commitments. Use this to view a customer's contract history and current agreements for billing management. Returns contract details with optional ledgers and balance information.
//
// ### Usage guidelines:
// - Pagination: Results are limited to 20 contracts per page; use 'cursor' for more
//
// ⚠️ Note: This is the legacy v1 endpoint - new integrations should use the v2 endpoint for enhanced features.
func (c v1ContractsService) ListContracts(ctx context.Context, params *V1ContractsListContractsParams) (*V1ListContractsResponse, error) {
	if params == nil {
		params = &V1ContractsListContractsParams{}
	}
	params.Context = ctx
	v1listcontractsresponse := &V1ListContractsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/list", c.Key, params, v1listcontractsresponse)
	return v1listcontractsresponse, err
}

// For a given package, returns all contract IDs and customer IDs associated with the package over a specific time period.
//
// ### Use this endpoint to:
// - Understand which customers are provisioned on a package at any given time for internal cohort management
// - Manage customer migrations to a new package. For example, to migrate all active customers to a new package, call this endpoint, end contracts, and provision customers on a new package.
//
// ### **Usage guidelines:**
// Use the **`starting_at`**, **`covering_date`**, and **`include_archived`** parameters to filter the list of returned contracts. For example, to list only currently active contracts, pass **`covering_date`** equal to the current time.
func (c v1ContractsService) ListContractsOnPackage(ctx context.Context, params *V1ContractsListContractsOnPackageParams) (*V1ListContractsOnPackageResponse, error) {
	if params == nil {
		params = &V1ContractsListContractsOnPackageParams{}
	}
	params.Context = ctx
	v1listcontractsonpackageresponse := &V1ListContractsOnPackageResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/packages/listContractsOnPackage", c.Key, params, v1listcontractsonpackageresponse)
	return v1listcontractsonpackageresponse, err
}

// Create a new scheduled invoice for Professional Services terms on a contract. This endpoint's availability is dependent on your client's configuration.
func (c v1ContractsService) ScheduleProServicesInvoice(ctx context.Context, params *V1ContractsScheduleProServicesInvoiceParams) (*V1ScheduleProServicesInvoiceResponse, error) {
	if params == nil {
		params = &V1ContractsScheduleProServicesInvoiceParams{}
	}
	params.Context = ctx
	v1scheduleproservicesinvoiceresponse := &V1ScheduleProServicesInvoiceResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/scheduleProServicesInvoice", c.Key, params, v1scheduleproservicesinvoiceresponse)
	return v1scheduleproservicesinvoiceresponse, err
}

// If a customer has multiple contracts with overlapping rates, the usage filter routes usage to the appropriate contract based on a predefined group key.
//
// As an example, imagine you have a customer associated with two projects. Each project is associated with its own contract. You can create a usage filter with group key `project_id`
// on each contract, and route usage for `project_1` to the first contract and `project_2` to the second contract.
//
// ### Use this endpoint to:
// - Support enterprise contracting scenarios where multiple contracts are associated to the same customer with the same rates.
// - Update the usage filter associated with the contract over time.
//
// ### Usage guidelines:
// To use usage filters, the `group_key` must be defined on the billable metrics underlying the rate card on the contracts.
func (c v1ContractsService) SetUsageFilter(ctx context.Context, params *V1ContractsSetUsageFilterParams) (*V1SetUsageFilterResponse, error) {
	if params == nil {
		params = &V1ContractsSetUsageFilterParams{}
	}
	params.Context = ctx
	v1setusagefilterresponse := &V1SetUsageFilterResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/setUsageFilter", c.Key, params, v1setusagefilterresponse)
	return v1setusagefilterresponse, err
}

// Update or add an end date to a contract. Ending a contract early will impact draft usage statements, truncate any terms, and remove upcoming scheduled invoices. Moving the date into the future will only extend the contract length. Terms and scheduled invoices are not extended. In-advance subscriptions will not be extended. Use this if a contract's end date has changed or if a perpetual contract ends.
func (c v1ContractsService) UpdateContractEndDate(ctx context.Context, params *V1ContractsUpdateContractEndDateParams) (*V1UpdateContractEndDateResponse, error) {
	if params == nil {
		params = &V1ContractsUpdateContractEndDateParams{}
	}
	params.Context = ctx
	v1updatecontractenddateresponse := &V1UpdateContractEndDateResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/updateEndDate", c.Key, params, v1updatecontractenddateresponse)
	return v1updatecontractenddateresponse, err
}

// Updates the issue date of a specific DRAFT invoice within a contract. Use this endpoint to reschedule when an invoice should be issued without affecting future billing cycles or the underlying contract terms. Only works with invoices still in DRAFT status, and the new issue date cannot be later than the contract's end date.
//
// ### Usage guidelines:
// This only changes the individual invoice's issue date - it does not modify the recurring invoice schedule of associated charges or commits. To update both the issue date and future billing schedule, use the 'edit contract' or 'edit commit' endpoints instead.
func (c v1ContractsService) UpdateInvoiceIssueDate(ctx context.Context, params *V1ContractsUpdateInvoiceIssueDateParams) (*V1UpdateInvoiceIssueDateResponse, error) {
	if params == nil {
		params = &V1ContractsUpdateInvoiceIssueDateParams{}
	}
	params.Context = ctx
	v1updateinvoiceissuedateresponse := &V1UpdateInvoiceIssueDateResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/updateInvoiceIssueDate", c.Key, params, v1updateinvoiceissuedateresponse)
	return v1updateinvoiceissuedateresponse, err
}
