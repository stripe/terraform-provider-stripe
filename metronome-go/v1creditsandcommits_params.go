//
// File generated from our OpenAPI spec
//

package metronome

// Manually adjust the available balance on a commit or credit. This entry is appended to the commit ledger as a new event. Optionally include a description that provides the reasoning for the entry.
//
// ### Use this endpoint to:
// - Address incorrect usage burn-down caused by malformed usage or invalid config
// - Decrease available balance to account for outages where usage may have not been tracked or sent to Metronome
// - Issue credits to customers in the form of increased balance on existing commit or credit
//
// ### Usage guidelines:
// Manual ledger entries can be extremely useful for resolving discrepancies in Metronome. However, most corrections to inaccurate billings can be modified upstream of the commit, whether that is via contract editing, rate editing, or other actions that cause an invoice to be recalculated.
type V1CreditsAndCommitsAddManualBalanceLedgerEntryParams struct {
	Params `form:"*"`
	// Amount to add to the segment. A negative number will draw down from the balance.
	Amount *Number `form:"amount" json:"amount"`
	// ID of the contract to update. Leave blank to update a customer level balance.
	ContractID *string `form:"contract_id" json:"contract_id,omitempty"`
	// ID of the customer whose balance is to be updated.
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// ID of the balance (commit or credit) to update.
	ID *string `form:"id" json:"id"`
	// If using individually configured commits/credits attached to seat managed subscriptions, the amount to add for each seat. Must sum to total amount.
	PerGroupAmounts map[string]Number `form:"per_group_amounts" json:"per_group_amounts,omitempty"`
	// Reason for the manual adjustment. This will be displayed in the ledger.
	Reason *string `form:"reason" json:"reason"`
	// ID of the segment to update.
	SegmentID *string `form:"segment_id" json:"segment_id"`
	// RFC 3339 timestamp indicating when the manual adjustment takes place. If not provided, it will default to the start of the segment.
	Timestamp     *string `form:"timestamp" json:"timestamp,omitempty"`
	UniquenessKey *string `form:"uniqueness_key" json:"uniqueness_key,omitempty"`
}
type V1CreditsAndCommitsCreateCustomerCommitAccessScheduleScheduleItemParams struct {
	Amount *Number `form:"amount" json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	// RFC 3339 timestamp (inclusive)
	StartingAt *string `form:"starting_at" json:"starting_at"`
}

// Schedule for distributing the commit to the customer. For "POSTPAID" commits only one schedule item is allowed and amount must match invoice_schedule total.
type V1CreditsAndCommitsCreateCustomerCommitAccessScheduleParams struct {
	// Determines how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used. Defaults to `SPEND` if omitted.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID  *string                                                                    `form:"credit_type_id" json:"credit_type_id,omitempty"`
	ScheduleItems []*V1CreditsAndCommitsCreateCustomerCommitAccessScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items"`
}

// Either provide amount or provide both unit_price and quantity.
type V1CreditsAndCommitsCreateCustomerCommitInvoiceScheduleScheduleItemParams struct {
	// Amount for the charge. Can be provided instead of unit_price and quantity. If amount is sent, the unit_price is assumed to be the amount and quantity is inferred to be 1.
	Amount *Number `form:"amount" json:"amount,omitempty"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the amount and must be specified with unit_price. If specified amount cannot be provided.
	Quantity *Number `form:"quantity" json:"quantity,omitempty"`
	// timestamp of the scheduled event
	Timestamp *string `form:"timestamp" json:"timestamp"`
	// Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified with quantity. If specified amount cannot be provided.
	UnitPrice *Number `form:"unit_price" json:"unit_price,omitempty"`
}

// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
type V1CreditsAndCommitsCreateCustomerCommitInvoiceScheduleRecurringScheduleParams struct {
	// Amount for the charge. Can be provided instead of unit_price and quantity. If amount is sent, the unit_price is assumed to be the amount and quantity is inferred to be 1.
	Amount             *Number `form:"amount" json:"amount,omitempty"`
	AmountDistribution *string `form:"amount_distribution" json:"amount_distribution"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	Frequency    *string `form:"frequency" json:"frequency"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the amount and must be specified with unit_price. If specified amount cannot be provided.
	Quantity *Number `form:"quantity" json:"quantity,omitempty"`
	// RFC 3339 timestamp (inclusive).
	StartingAt *string `form:"starting_at" json:"starting_at"`
	// Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified with quantity. If specified amount cannot be provided.
	UnitPrice *Number `form:"unit_price" json:"unit_price,omitempty"`
}

// Required for "POSTPAID" commits: the true up invoice will be generated at this time and only one schedule item is allowed; the total must match accesss_schedule amount. Optional for "PREPAID" commits: if not provided, this will be a "complimentary" commit with no invoice.
type V1CreditsAndCommitsCreateCustomerCommitInvoiceScheduleParams struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is only applicable to commit invoice schedules. If true, this schedule will not generate an invoice.
	DoNotInvoice *bool `form:"do_not_invoice" json:"do_not_invoice,omitempty"`
	// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
	RecurringSchedule *V1CreditsAndCommitsCreateCustomerCommitInvoiceScheduleRecurringScheduleParams `form:"recurring_schedule" json:"recurring_schedule,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []*V1CreditsAndCommitsCreateCustomerCommitInvoiceScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items,omitempty"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1CreditsAndCommitsCreateCustomerCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1CreditsAndCommitsCreateCustomerCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1CreditsAndCommitsCreateCustomerCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// ⚠️ For most contract amendments, use `contracts/edit` directly. Use this endpoint only for cross-contract or enterprise-wide commits.
//
// Creates customer-level commits that establish spending commitments for customers across their Metronome usage. Commits represent contracted spending obligations that can be either prepaid (paid upfront) or postpaid (billed later).
//
// Note: In most cases, you should add commitments directly to customer contracts using the contract/create or contract/edit APIs.
//
// ### Use this endpoint to:
// Use this endpoint when you need to establish customer-level spending commitments that can be applied across multiple contracts or scoped to specific contracts. Customer-level commits are ideal for:
// - Enterprise-wide minimum spending agreements that span multiple contracts
// - Multi-contract volume commitments with shared spending pools
// - Cross-contract discount tiers based on aggregate usage
//
// #### Commit type Requirements:
// - You must specify either "prepaid" or "postpaid" as the commit type:
// - Prepaid commits: Customer pays upfront; invoice_schedule is optional (if omitted, creates a commit without an invoice)
// - Postpaid commits: Customer pays when the commitment expires (the end of the access_schedule); invoice_schedule is required and must match access_schedule totals.
//
// #### Billing configuration:
// - invoice_contract_id is required for postpaid commits and for prepaid commits with billing (only optional for free prepaid commits) unless do_not_invoice is set to true
// - For postpaid commits: access_schedule and invoice_schedule must have matching amounts
// - For postpaid commits: only one schedule item is allowed in both schedules.
//
// #### Scoping flexibility:
// Customer-level commits can be configured in a few ways:
// - Contract-specific: Use the `applicable_contract_ids` field to limit the commit to specific contracts
// - Cross-contract: Leave `applicable_contract_ids` empty to allow the commit to be used across all of the customer's contracts
//
// #### Product targeting:
// Commits can be scoped to specific products using applicable_product_ids, applicable_product_tags, or specifiers, or left unrestricted to apply to all products.
//
// #### Priority considerations:
// When multiple commits are applicable, the one with the lower priority value will be consumed first. If there is a tie, contract level commits and credits will be applied before customer level commits and credits. Plan your priority scheme carefully to ensure commits are applied in the desired order.
//
// ### Usage guidelines:
// ⚠️ Preferred Alternative: In most cases, you should add commits directly to contracts using the create contract or edit contract APIs instead of creating customer-level commits. Contract-level commits provide better organization and are the recommended approach for standard use cases.
type V1CreditsAndCommitsCreateCustomerCommitParams struct {
	Params `form:"*"`
	// Schedule for distributing the commit to the customer. For "POSTPAID" commits only one schedule item is allowed and amount must match invoice_schedule total.
	AccessSchedule *V1CreditsAndCommitsCreateCustomerCommitAccessScheduleParams `form:"access_schedule" json:"access_schedule"`
	// Which contract the commit applies to. If not provided, the commit applies to all contracts.
	ApplicableContractIDs []*string `form:"applicable_contract_ids" json:"applicable_contract_ids,omitempty"`
	// Which products the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductTags []*string         `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	CustomerID            *string           `form:"customer_id" json:"customer_id"`
	CustomFields          map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// Used only in UI/API. It is not exposed to end customers.
	Description *string `form:"description" json:"description,omitempty"`
	// The contract that this commit will be billed on. This is required for "POSTPAID" commits and for "PREPAID" commits unless there is no invoice schedule above (i.e., the commit is 'free'), or if do_not_invoice is set to true.
	InvoiceContractID *string `form:"invoice_contract_id" json:"invoice_contract_id,omitempty"`
	// Required for "POSTPAID" commits: the true up invoice will be generated at this time and only one schedule item is allowed; the total must match accesss_schedule amount. Optional for "PREPAID" commits: if not provided, this will be a "complimentary" commit with no invoice.
	InvoiceSchedule *V1CreditsAndCommitsCreateCustomerCommitInvoiceScheduleParams `form:"invoice_schedule" json:"invoice_schedule,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// If multiple credits or commits are applicable, the one with the lower priority will apply first.
	Priority *Number `form:"priority" json:"priority"`
	// ID of the fixed product associated with the commit. This is required because products are used to invoice the commit amount.
	ProductID *string `form:"product_id" json:"product_id"`
	RateType  *string `form:"rate_type" json:"rate_type,omitempty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID *string `form:"salesforce_opportunity_id" json:"salesforce_opportunity_id,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers    []*V1CreditsAndCommitsCreateCustomerCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
	Type          *string                                                   `form:"type" json:"type"`
	UniquenessKey *string                                                   `form:"uniqueness_key" json:"uniqueness_key,omitempty"`
}
type V1CreditsAndCommitsCreateCustomerCreditAccessScheduleScheduleItemParams struct {
	Amount *Number `form:"amount" json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	// RFC 3339 timestamp (inclusive)
	StartingAt *string `form:"starting_at" json:"starting_at"`
}

// Schedule for distributing the credit to the customer.
type V1CreditsAndCommitsCreateCustomerCreditAccessScheduleParams struct {
	// Determines how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used. Defaults to `SPEND` if omitted.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID  *string                                                                    `form:"credit_type_id" json:"credit_type_id,omitempty"`
	ScheduleItems []*V1CreditsAndCommitsCreateCustomerCreditAccessScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1CreditsAndCommitsCreateCustomerCreditSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1CreditsAndCommitsCreateCustomerCreditSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1CreditsAndCommitsCreateCustomerCreditSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// ⚠️ For most contract amendments, use `contracts/edit` directly. Use this endpoint only for cross-contract or enterprise-wide commits.
//
// Creates customer-level credits that provide spending allowances or free credit balances for customers across their Metronome usage. Note: In most cases, you should add credits directly to customer contracts using the contract/create or contract/edit APIs.
//
// ### Use this endpoint to:
// Use this endpoint when you need to provision credits directly at the customer level that can be applied across multiple contracts or scoped to specific contracts. Customer-level credits are ideal for:
// - Customer onboarding incentives that apply globally
// - Flexible spending allowances that aren't tied to a single contract
// - Migration scenarios where you need to preserve existing customer balances
//
// #### Scoping flexibility:
// Customer-level credits can be configured in two ways:
// - Contract-specific: Use the applicable_contract_ids field to limit the credit to specific contracts
// - Cross-contract: Leave applicable_contract_ids empty to allow the credit to be used across all of the customer's contracts
//
// #### Product Targeting:
// Credits can be scoped to specific products using `applicable_product_ids` or `applicable_product_tags`, or left unrestricted to apply to all products.
//
// #### Priority considerations:
// When multiple credits are applicable, the one with the lower priority value will be consumed first. If there is a tie, contract level commits and credits will be applied before customer level commits and credits. Plan your priority scheme carefully to ensure credits are applied in the desired order.
//
// #### Access Schedule Required:
// You must provide an `access_schedule` that defines when and how much credit becomes available to the customer over time. This usually is aligned to the contract schedule or starts immediately and is set to expire in the future.
//
// ### Usage Guidelines:
// ⚠️ Preferred Alternative: In most cases, you should add credits directly to contracts using the contract/create or contract/edit APIs instead of creating customer-level credits. Contract-level credits provide better organization, and are easier for finance teams to recognize revenue, and are the recommended approach for most use cases.
type V1CreditsAndCommitsCreateCustomerCreditParams struct {
	Params `form:"*"`
	// Schedule for distributing the credit to the customer.
	AccessSchedule *V1CreditsAndCommitsCreateCustomerCreditAccessScheduleParams `form:"access_schedule" json:"access_schedule"`
	// Which contract the credit applies to. If not provided, the credit applies to all contracts.
	ApplicableContractIDs []*string `form:"applicable_contract_ids" json:"applicable_contract_ids,omitempty"`
	// Which products the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductTags []*string         `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	CustomerID            *string           `form:"customer_id" json:"customer_id"`
	CustomFields          map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// Used only in UI/API. It is not exposed to end customers.
	Description *string `form:"description" json:"description,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// If multiple credits or commits are applicable, the one with the lower priority will apply first.
	Priority  *Number `form:"priority" json:"priority"`
	ProductID *string `form:"product_id" json:"product_id"`
	RateType  *string `form:"rate_type" json:"rate_type,omitempty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID *string `form:"salesforce_opportunity_id" json:"salesforce_opportunity_id,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers    []*V1CreditsAndCommitsCreateCustomerCreditSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
	UniquenessKey *string                                                   `form:"uniqueness_key" json:"uniqueness_key,omitempty"`
}

// Disable the true-up invoice for a postpaid commit. If used, the true-up invoice will not be generated.
//
// For postpaid commits, usage during the access period is paid for in arrears. If the total amount paid during the access period is less than the committed amount, there's a final true-up invoice on the invoice_date.
type V1CreditsAndCommitsDisableCommitTrueupParams struct {
	Params `form:"*"`
	// If applicable, the amendment ID that the commit is on
	AmendmentID *string `form:"amendment_id" json:"amendment_id,omitempty"`
	// ID of the commit to update
	CommitID *string `form:"commit_id" json:"commit_id"`
	// ID of the contract that the commit is on
	ContractID *string `form:"contract_id" json:"contract_id"`
	// ID of the customer whose commit is to be updated
	CustomerID *string `form:"customer_id" json:"customer_id"`
}

// Balance filters are OR'd together, so if a given commit or credit matches any of the filters, it will be included in the net balance.
type V1CreditsAndCommitsGetNetBalanceFilterParams struct {
	// The balance type to filter by.
	BalanceTypes []*string `form:"balance_types" json:"balance_types,omitempty"`
	// Custom fields to compute balance across. Must match all custom fields
	CustomFields map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// Specific IDs to compute balance across.
	IDs []*string `form:"ids" json:"ids,omitempty"`
}

// Retrieve the combined current balance across any grouping of credits and commits for a customer in a single API call.
// - Display real-time available balance to customers in billing dashboards
// - Build finance dashboards showing credit utilization across customer segments
// - Validate expected vs. actual balance during billing reconciliation
//
// ### Key response fields:
// - `balance`: The combined net balance available to use at this moment across all matching commits and credits
// - `credit_type_id`: The credit type (fiat or custom pricing unit) the balance is denominated in
//
// ### Filtering options:
// Balance filters allow you to scope the calculation to specific subsets of commits and credits. When using multiple filter objects, they are OR'd together — if a commit or credit matches any filter, it's included in the net balance. Within a single filter object, all specified conditions are AND'd together.
// - **Balance types**: Include any combination of `PREPAID_COMMIT`, `POSTPAID_COMMIT`, and `CREDIT` (e.g., `["PREPAID_COMMIT", "CREDIT"]` to exclude postpaid commits). If not specified, all balance types are included.
// - **Specific IDs**: Target exact commit or credit IDs for precise balance queries
// - **Custom fields**: Filter by custom field key-value pairs; when multiple pairs are provided, commits must match all of them
//
// **Example**: To get the balance of all free-trial credits OR all signup-promotion commits, you'd pass two filter objects — one filtering for CREDIT with custom field campaign: free-trial, and another filtering for PREPAID_COMMIT with custom field campaign: signup-promotion.
//
// ### Usage guidelines:
// - **Balance ledger details**: Use the [listBalances](https://docs.metronome.com/api-reference/credits-and-commits/list-balances) endpoint instead to understand detailed ledger drawdowns for each individual balance
// - **Draft invoice handling**: Use `invoice_inclusion_mode` to control whether pending draft invoice deductions are included (`FINALIZED_AND_DRAFT`, the default) or excluded (`FINALIZED`) from the balance calculation
// - **Account hierarchies**: When querying a child customer, shared commits from parent contracts are not included — query the parent customer directly to see shared commit balances
// - **Negative balances**: Manual ledger entries can cause negative segment balances; these are treated as zero when calculating the net balance
// - **Credit types**: If `credit_type_id` is not specified, the balance defaults to USD (cents)
type V1CreditsAndCommitsGetNetBalanceParams struct {
	Params `form:"*"`
	// Filters balances by how they are drawn down. Defaults to `SPEND`. If set to `QUANTITY`, `credit_type_id` must not be provided.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// The ID of the credit type (can be fiat or a custom pricing unit) to get the balance for. Defaults to USD (cents) if not specified.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// The ID of the customer.
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// Balance filters are OR'd together, so if a given commit or credit matches any of the filters, it will be included in the net balance.
	Filters []*V1CreditsAndCommitsGetNetBalanceFilterParams `form:"filters" json:"filters,omitempty"`
	// Controls which invoices are considered when calculating the remaining balance. `FINALIZED` considers only deductions from finalized invoices. `FINALIZED_AND_DRAFT` also includes deductions from pending draft invoices.
	InvoiceInclusionMode *string `form:"invoice_inclusion_mode" json:"invoice_inclusion_mode,omitempty"`
}

// Retrieve a comprehensive view of all available balances (commits and credits) for a customer. This endpoint provides real-time visibility into prepaid funds, postpaid commitments, promotional credits, and other balance types that can offset usage charges, helping you build transparent billing experiences.
//
// ### Use this endpoint to:
// - Display current available balances in customer dashboards
// - Verify available funds before approving high-usage operations
// - Generate balance reports for finance teams
// - Filter balances by contract or date ranges
//
// ### Key response fields:
// An array of balance objects (all credits and commits) containing:
//
// - Balance details: Current available amount for each commit or credit
// - Metadata: Product associations, priorities, applicable date ranges
// - Optional ledger entries: Detailed transaction history (if `include_ledgers=true`)
// - Balance calculations: Including pending transactions and future-dated entries
// - Custom fields: Any additional metadata attached to balances
//
// ### Usage guidelines:
// - Use the [getNetBalance](https://docs.metronome.com/api-reference/credits-and-commits/get-the-net-balance-of-a-customer) endpoint to retrieve a single combined current balance
// - Date filtering: Use `effective_before` to include only balances with access before a specific date (exclusive)
// - Set `include_balance=true` for calculated balance amounts on each commit or credit
// - Set `include_ledgers=true` for full transaction history
// - Set `include_contract_balances = true` to see contract level balances
// - Balance logic: Reflects currently accessible amounts, excluding expired/future segments
// - Manual adjustments: Includes all manual ledger entries, even future-dated ones
type V1CreditsAndCommitsListCustomerBalancesParams struct {
	Params `form:"*"`
	// Filters balances by how they are drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Return only balances that have access schedules that "cover" the provided date
	CoveringDate *string `form:"covering_date" json:"covering_date,omitempty"`
	CustomerID   *string `form:"customer_id" json:"customer_id"`
	// Include only balances that have any access before the provided date (exclusive)
	EffectiveBefore *string `form:"effective_before" json:"effective_before,omitempty"`
	// Exclude balances with zero amounts from the response.
	ExcludeZeroBalances *bool   `form:"exclude_zero_balances" json:"exclude_zero_balances,omitempty"`
	ID                  *string `form:"id" json:"id,omitempty"`
	// Include archived credits and credits from archived contracts.
	IncludeArchived *bool `form:"include_archived" json:"include_archived,omitempty"`
	// Include the balance of credits and commits in the response. Setting this flag may cause the query to be slower.
	IncludeBalance *bool `form:"include_balance" json:"include_balance,omitempty"`
	// Include balances on the contract level.
	IncludeContractBalances *bool `form:"include_contract_balances" json:"include_contract_balances,omitempty"`
	// Include ledgers in the response. Setting this flag may cause the query to be slower.
	IncludeLedgers *bool `form:"include_ledgers" json:"include_ledgers,omitempty"`
	// The maximum number of commits to return. Defaults to 25.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// The next page token from a previous response.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
	// Include only balances that have any access on or after the provided date
	StartingAt *string `form:"starting_at" json:"starting_at,omitempty"`
	// Indicates that this API request was triggered by a webhook notification with the provided ID.
	WebhookNotificationID *string `form:"webhook_notification_id" json:"webhook_notification_id,omitempty"`
}

// Retrieve all commit agreements for a customer, including both prepaid and postpaid commitments. This endpoint provides comprehensive visibility into contractual spending obligations, enabling you to track commitment utilization and manage customer contracts effectively.
//
// ### Use this endpoint to:
// - Display commitment balances and utilization in customer dashboards
// - Track prepaid commitment drawdown and remaining balances
// - Monitor postpaid commitment progress toward minimum thresholds
// - Build commitment tracking and forecasting tools
// - Show commitment history with optional ledger details
// - Manage rollover balances between contract periods
//
// ### Key response fields:
// An array of Commit objects containing:
// - Commit type: PREPAID (pay upfront) or POSTPAID (pay at true-up)
// - Rate type: COMMIT_RATE (discounted) or LIST_RATE (standard pricing)
// - Access schedule: When commitment funds become available
// - Invoice schedule: When the customer is billed
// - Product targeting: Which product(s) usage is eligible to draw from this commit
// - Optional ledger entries: Transaction history (if `include_ledgers=true`)
// - Balance information: Current available amount (if `include_balance=true`)
// - Rollover settings: Fraction of unused amount that carries forward
//
// ### Usage guidelines:
// - Pagination: Results limited to 25 commits per page; use 'next_page' for more
// - Date filtering options:
//   - `covering_date`: Commits active on a specific date
//   - `starting_at`: Commits with access on/after a date
//   - `effective_before`: Commits with access before a date (exclusive)
//
// - Scope options:
//   - `include_contract_commits`: Include contract-level commits (not just customer-level)
//   - `include_archived`: Include archived commits and commits from archived contracts
//
// - Performance considerations:
//   - include_ledgers: Adds detailed transaction history (slower)
//   - include_balance: Adds current balance calculation (slower)
//
// - Optional filtering: Use commit_id to retrieve a specific commit
type V1CreditsAndCommitsListCustomerCommitsParams struct {
	Params `form:"*"`
	// Filters commits by how their balances are drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	CommitID   *string `form:"commit_id" json:"commit_id,omitempty"`
	// Include only commits that have access schedules that "cover" the provided date
	CoveringDate *string `form:"covering_date" json:"covering_date,omitempty"`
	CustomerID   *string `form:"customer_id" json:"customer_id"`
	// Include only commits that have any access before the provided date (exclusive)
	EffectiveBefore *string `form:"effective_before" json:"effective_before,omitempty"`
	// Include archived commits and commits from archived contracts.
	IncludeArchived *bool `form:"include_archived" json:"include_archived,omitempty"`
	// Include the balance in the response. Setting this flag may cause the query to be slower.
	IncludeBalance *bool `form:"include_balance" json:"include_balance,omitempty"`
	// Include commits on the contract level.
	IncludeContractCommits *bool `form:"include_contract_commits" json:"include_contract_commits,omitempty"`
	// Include commit ledgers in the response. Setting this flag may cause the query to be slower.
	IncludeLedgers *bool `form:"include_ledgers" json:"include_ledgers,omitempty"`
	// The maximum number of commits to return. Defaults to 25.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// The next page token from a previous response.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
	// Include only commits that have any access on or after the provided date
	StartingAt *string `form:"starting_at" json:"starting_at,omitempty"`
}

// Retrieve a detailed list of all credits available to a customer, including promotional credits and contract-specific credits. This endpoint provides comprehensive visibility into credit balances, access schedules, and usage rules, enabling you to build credit management interfaces and track available funding.
//
// ### Use this endpoint to:
// - Display all available credits in customer billing dashboards
// - Show credit balances and expiration dates
// - Track credit usage history with optional ledger details
// - Build credit management and reporting tools
// - Monitor promotional credit utilization
// • Support customer inquiries about available credits
//
// ### Key response fields:
// An array of Credit objects containing:
// - Credit details: Name, priority, and which applicable products/tags it applies to
// - Product ID: The `product_id` of the credit. This is for external mapping into your quote-to-cash stack, not the product it applies to.
// - Access schedule: When credits become available and expire
// - Optional ledger entries: Transaction history (if `include_ledgers=true`)
// - Balance information: Current available amount (if `include_balance=true`)
// - Metadata: Custom fields and usage specifiers
//
// ### Usage guidelines:
// - Pagination: Results limited to 25 commits per page; use next_page for more
// - Date filtering options:
//   - `covering_date`: Credits active on a specific date
//   - `starting_at`: Credits with access on/after a date
//   - `effective_before`: Credits with access before a date (exclusive)
//
// - Scope options:
//   - `include_contract_credits`: Include contract-level credits (not just customer-level)
//   - `include_archived`: Include archived credits and credits from archived contracts
//
// - Performance considerations:
//   - `include_ledgers`: Adds detailed transaction history (slower)
//   - `include_balance`: Adds current balance calculation (slower)
//
// - Optional filtering: Use credit_id to retrieve a specific commit
type V1CreditsAndCommitsListCustomerCreditsParams struct {
	Params `form:"*"`
	// Return only credits that have access schedules that "cover" the provided date
	CoveringDate *string `form:"covering_date" json:"covering_date,omitempty"`
	CreditID     *string `form:"credit_id" json:"credit_id,omitempty"`
	CustomerID   *string `form:"customer_id" json:"customer_id"`
	// Include only credits that have any access before the provided date (exclusive)
	EffectiveBefore *string `form:"effective_before" json:"effective_before,omitempty"`
	// Include archived credits and credits from archived contracts.
	IncludeArchived *bool `form:"include_archived" json:"include_archived,omitempty"`
	// Include the balance in the response. Setting this flag may cause the query to be slower.
	IncludeBalance *bool `form:"include_balance" json:"include_balance,omitempty"`
	// Include credits on the contract level.
	IncludeContractCredits *bool `form:"include_contract_credits" json:"include_contract_credits,omitempty"`
	// Include credit ledgers in the response. Setting this flag may cause the query to be slower.
	IncludeLedgers *bool `form:"include_ledgers" json:"include_ledgers,omitempty"`
	// The maximum number of commits to return. Defaults to 25.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// The next page token from a previous response.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
	// Include only credits that have any access on or after the provided date
	StartingAt *string `form:"starting_at" json:"starting_at,omitempty"`
}

// Retrieve historical commit and credit balances for a customer. Historical balances include their schedules and current balance. Set `include_ledgers` to include transaction history. Ledgers and balances may change when invoices are voided and regenerated.
//
// Use the date filters to select commits by either their access schedule or the time they became historical.
type V1CreditsAndCommitsListHistoricalBalancesParams struct {
	Params `form:"*"`
	// Return only historical balances from this contract.
	ContractID *string `form:"contract_id" json:"contract_id,omitempty"`
	// Return only historical balances with access schedules that cover this date.
	CoveringDate *string `form:"covering_date" json:"covering_date,omitempty"`
	CustomerID   *string `form:"customer_id" json:"customer_id"`
	// Return only historical balances with any access before this date (exclusive).
	EffectiveBefore *string `form:"effective_before" json:"effective_before,omitempty"`
	// Return only the historical balance with this ID.
	ID *string `form:"id" json:"id,omitempty"`
	// Include ledgers in the response. Setting this flag may cause the query to be slower.
	IncludeLedgers *bool `form:"include_ledgers" json:"include_ledgers,omitempty"`
	// The maximum number of historical balances to return. Defaults to 25.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// The next page token from a previous response.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
	// Return only balances that became historical at or after this date.
	RetiredAfter *string `form:"retired_after" json:"retired_after,omitempty"`
	// Return only balances that became historical before this date (exclusive).
	RetiredBefore *string `form:"retired_before" json:"retired_before,omitempty"`
	// Return only historical balances with any access on or after this date.
	StartingAt *string `form:"starting_at" json:"starting_at,omitempty"`
}

// Retrieve detailed balance for seat-based credits and commits from the contract's subscriptions, broken down by individual seats.
//
// ### Use this endpoint to:
// - Display per-seat balance information in customer dashboards
// - Filter balance data by subscription or specific seats
//
// ### Key response fields:
// An array of seat balance objects containing:
// - Seat id
// - Balance: current total balance across all commits and credits
//
// ### Usage guidelines:
// - Date filtering: use `covering_date` OR `starting_at`/`ending_before` to filter balance data by time range
// - Set `include_credits_and_commits=true` for detailed commits and credits breakdown per seat
// - Set `include_ledgers=true` for detailed transaction history per commit/credit per seat
type V1CreditsAndCommitsListSeatBalancesParams struct {
	Params `form:"*"`
	// The contract ID to retrieve seat balances for
	ContractID *string `form:"contract_id" json:"contract_id"`
	// Include only commits or credits with access that cover this specific date (cannot be used with starting_at or ending_before).
	CoveringDate *string `form:"covering_date" json:"covering_date,omitempty"`
	// Page token from a previous response to retrieve the next page
	Cursor *string `form:"cursor" json:"cursor,omitempty"`
	// The customer ID to retrieve seat balances for
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// Include only commits or credits with access effective on or before this date (cannot be used with covering_date).
	EffectiveBefore *string `form:"effective_before" json:"effective_before,omitempty"`
	// Include credits and commits in the response
	IncludeCreditsAndCommits *bool `form:"include_credits_and_commits" json:"include_credits_and_commits,omitempty"`
	// Include ledger entries for each commit and commit. `include_credits_and_commits` must be set to `true` for `include_ledgers=true` to apply.
	IncludeLedgers *bool `form:"include_ledgers" json:"include_ledgers,omitempty"`
	// Maximum number of seats to return. Range: 1-100. Default: 25.
	// When `include_credits_and_commits = true`, if the total commits/credits across all seats exceeds 100, a limit of 100 applies to the total credits and commits. Seats are included greedily to maximize the number of seats returned.
	// Example: if seat 1 has 98 commits and seat 2 has 10 commits, both seats will be returned (total: 108 commits). Each returned seat includes all of its associated credits and commits.
	//
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Optional filter to only include specific seats.
	SeatIDs []*string `form:"seat_ids" json:"seat_ids,omitempty"`
	// When true, any seat_ids not found in contract subscriptions will be silently omitted from the response instead of returning a 400 error.
	SkipMissingSeatIDs *bool `form:"skip_missing_seat_ids" json:"skip_missing_seat_ids,omitempty"`
	// Include only commits or credits with access effective on or after this date (cannot be used with covering_date).
	StartingAt *string `form:"starting_at" json:"starting_at,omitempty"`
	// Optional filter to only include seats from specific subscriptions. If subscriptions ids are not mapped to SEAT_BASED subscriptions, error will be returned.
	SubscriptionIDs []*string `form:"subscription_ids" json:"subscription_ids,omitempty"`
}

// If using threshold billing with an external payment gateway, Metronome does not facilitate the payment gating process on behalf of the client. As a result, clients must facilitate the transaction themselves. This end-point is used to either release or cancel the commit pending on the outcome of the external payment attempt.
//
// To release the commit, you must pass the `workflow_id` provided in the `payment_gate.external_initiate` webhook.
//
// ### Use this endpoint to:
// Facilitate payment gating workflows for threshold billing if using a payment gateway Metronome does not support today.
//
// ### Usage guidelines:
// Ensure that you are set up to consume the `payment_gate.external_initiate` webhook and save the `workflow_id`.
type V1CreditsAndCommitsReleaseExternalPaymentGateThresholdCommitParams struct {
	Params `form:"*"`
	// The outcome of the external payment gate
	Outcome *string `form:"outcome" json:"outcome"`
	// ID of the workflow to continue
	WorkflowID *string `form:"workflow_id" json:"workflow_id"`
}

// Shortens the end date of a prepaid commit to terminate it earlier than originally scheduled. Use this endpoint when you need to cancel or reduce the duration of an existing prepaid commit. Only works with prepaid commit types and can only move the end date forward (earlier), not extend it.
//
// ### Usage guidelines:
// To extend commit end dates or make other comprehensive edits, use the 'edit commit' endpoint instead.
type V1CreditsAndCommitsUpdateCommitEndDateParams struct {
	Params `form:"*"`
	// RFC 3339 timestamp indicating when access to the commit will end and it will no longer be possible to draw it down (exclusive). If not provided, the access will not be updated.
	AccessEndingBefore *string `form:"access_ending_before" json:"access_ending_before,omitempty"`
	// ID of the commit to update. Only supports "PREPAID" commits.
	CommitID *string `form:"commit_id" json:"commit_id"`
	// ID of the customer whose commit is to be updated
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// RFC 3339 timestamp indicating when the commit will stop being invoiced (exclusive). If not provided, the invoice schedule will not be updated.
	InvoicesEndingBefore *string `form:"invoices_ending_before" json:"invoices_ending_before,omitempty"`
}

// Shortens the end date of an existing customer credit to terminate it earlier than originally scheduled. Only allows moving end dates forward (earlier), not extending them.
//
// Note: To extend credit end dates or make comprehensive edits, use the 'edit credit' endpoint instead.
type V1CreditsAndCommitsUpdateCreditEndDateParams struct {
	Params `form:"*"`
	// RFC 3339 timestamp indicating when access to the credit will end and it will no longer be possible to draw it down (exclusive).
	AccessEndingBefore *string `form:"access_ending_before" json:"access_ending_before"`
	// ID of the commit to update
	CreditID *string `form:"credit_id" json:"credit_id"`
	// ID of the customer whose credit is to be updated
	CustomerID *string `form:"customer_id" json:"customer_id"`
}
