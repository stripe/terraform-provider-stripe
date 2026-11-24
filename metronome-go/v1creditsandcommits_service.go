//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v1CreditsAndCommitsService is used to invoke v1creditsandcommits related APIs.
type v1CreditsAndCommitsService struct {
	B   Backend
	Key string
}

// Manually adjust the available balance on a commit or credit. This entry is appended to the commit ledger as a new event. Optionally include a description that provides the reasoning for the entry.
//
// ### Use this endpoint to:
// - Address incorrect usage burn-down caused by malformed usage or invalid config
// - Decrease available balance to account for outages where usage may have not been tracked or sent to Metronome
// - Issue credits to customers in the form of increased balance on existing commit or credit
//
// ### Usage guidelines:
// Manual ledger entries can be extremely useful for resolving discrepancies in Metronome. However, most corrections to inaccurate billings can be modified upstream of the commit, whether that is via contract editing, rate editing, or other actions that cause an invoice to be recalculated.
func (c v1CreditsAndCommitsService) AddManualBalanceLedgerEntry(ctx context.Context, params *V1CreditsAndCommitsAddManualBalanceLedgerEntryParams) (*V1AddManualBalanceLedgerEntryResponse, error) {
	if params == nil {
		params = &V1CreditsAndCommitsAddManualBalanceLedgerEntryParams{}
	}
	params.Context = ctx
	v1addmanualbalanceledgerentryresponse := &V1AddManualBalanceLedgerEntryResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/addManualBalanceLedgerEntry", c.Key, params, v1addmanualbalanceledgerentryresponse)
	return v1addmanualbalanceledgerentryresponse, err
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
func (c v1CreditsAndCommitsService) CreateCustomerCommit(ctx context.Context, params *V1CreditsAndCommitsCreateCustomerCommitParams) (*V1CreateCustomerCommitResponse, error) {
	if params == nil {
		params = &V1CreditsAndCommitsCreateCustomerCommitParams{}
	}
	params.Context = ctx
	v1createcustomercommitresponse := &V1CreateCustomerCommitResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/customerCommits/create", c.Key, params, v1createcustomercommitresponse)
	return v1createcustomercommitresponse, err
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
func (c v1CreditsAndCommitsService) CreateCustomerCredit(ctx context.Context, params *V1CreditsAndCommitsCreateCustomerCreditParams) (*V1CreateCustomerCreditResponse, error) {
	if params == nil {
		params = &V1CreditsAndCommitsCreateCustomerCreditParams{}
	}
	params.Context = ctx
	v1createcustomercreditresponse := &V1CreateCustomerCreditResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/customerCredits/create", c.Key, params, v1createcustomercreditresponse)
	return v1createcustomercreditresponse, err
}

// Disable the true-up invoice for a postpaid commit. If used, the true-up invoice will not be generated.
//
// For postpaid commits, usage during the access period is paid for in arrears. If the total amount paid during the access period is less than the committed amount, there's a final true-up invoice on the invoice_date.
func (c v1CreditsAndCommitsService) DisableCommitTrueup(ctx context.Context, params *V1CreditsAndCommitsDisableCommitTrueupParams) (*V1DisableCommitTrueupResponse, error) {
	if params == nil {
		params = &V1CreditsAndCommitsDisableCommitTrueupParams{}
	}
	params.Context = ctx
	v1disablecommittrueupresponse := &V1DisableCommitTrueupResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/commits/disableTrueup", c.Key, params, v1disablecommittrueupresponse)
	return v1disablecommittrueupresponse, err
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
func (c v1CreditsAndCommitsService) GetNetBalance(ctx context.Context, params *V1CreditsAndCommitsGetNetBalanceParams) (*V1GetNetBalanceResponse, error) {
	if params == nil {
		params = &V1CreditsAndCommitsGetNetBalanceParams{}
	}
	params.Context = ctx
	v1getnetbalanceresponse := &V1GetNetBalanceResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/customerBalances/getNetBalance", c.Key, params, v1getnetbalanceresponse)
	return v1getnetbalanceresponse, err
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
func (c v1CreditsAndCommitsService) ListCustomerBalances(ctx context.Context, params *V1CreditsAndCommitsListCustomerBalancesParams) (*V1ListCustomerBalancesResponse, error) {
	if params == nil {
		params = &V1CreditsAndCommitsListCustomerBalancesParams{}
	}
	params.Context = ctx
	v1listcustomerbalancesresponse := &V1ListCustomerBalancesResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/customerBalances/list", c.Key, params, v1listcustomerbalancesresponse)
	return v1listcustomerbalancesresponse, err
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
func (c v1CreditsAndCommitsService) ListCustomerCommits(ctx context.Context, params *V1CreditsAndCommitsListCustomerCommitsParams) (*V1ListCustomerCommitsResponse, error) {
	if params == nil {
		params = &V1CreditsAndCommitsListCustomerCommitsParams{}
	}
	params.Context = ctx
	v1listcustomercommitsresponse := &V1ListCustomerCommitsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/customerCommits/list", c.Key, params, v1listcustomercommitsresponse)
	return v1listcustomercommitsresponse, err
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
func (c v1CreditsAndCommitsService) ListCustomerCredits(ctx context.Context, params *V1CreditsAndCommitsListCustomerCreditsParams) (*V1ListCustomerCreditsResponse, error) {
	if params == nil {
		params = &V1CreditsAndCommitsListCustomerCreditsParams{}
	}
	params.Context = ctx
	v1listcustomercreditsresponse := &V1ListCustomerCreditsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/customerCredits/list", c.Key, params, v1listcustomercreditsresponse)
	return v1listcustomercreditsresponse, err
}

// Retrieve historical commit and credit balances for a customer. Historical balances include their schedules and current balance. Set `include_ledgers` to include transaction history. Ledgers and balances may change when invoices are voided and regenerated.
//
// Use the date filters to select commits by either their access schedule or the time they became historical.
func (c v1CreditsAndCommitsService) ListHistoricalBalances(ctx context.Context, params *V1CreditsAndCommitsListHistoricalBalancesParams) (*V1ListHistoricalBalancesResponse, error) {
	if params == nil {
		params = &V1CreditsAndCommitsListHistoricalBalancesParams{}
	}
	params.Context = ctx
	v1listhistoricalbalancesresponse := &V1ListHistoricalBalancesResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/historicalBalances/list", c.Key, params, v1listhistoricalbalancesresponse)
	return v1listhistoricalbalancesresponse, err
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
func (c v1CreditsAndCommitsService) ListSeatBalances(ctx context.Context, params *V1CreditsAndCommitsListSeatBalancesParams) (*V1ListSeatBalancesResponse, error) {
	if params == nil {
		params = &V1CreditsAndCommitsListSeatBalancesParams{}
	}
	params.Context = ctx
	v1listseatbalancesresponse := &V1ListSeatBalancesResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/seatBalances/list", c.Key, params, v1listseatbalancesresponse)
	return v1listseatbalancesresponse, err
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
func (c v1CreditsAndCommitsService) ReleaseExternalPaymentGateThresholdCommit(ctx context.Context, params *V1CreditsAndCommitsReleaseExternalPaymentGateThresholdCommitParams) (*V1ReleaseExternalPaymentGateThresholdCommitResponse, error) {
	if params == nil {
		params = &V1CreditsAndCommitsReleaseExternalPaymentGateThresholdCommitParams{}
	}
	params.Context = ctx
	v1releaseexternalpaymentgatethresholdcommitresponse := &V1ReleaseExternalPaymentGateThresholdCommitResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/commits/threshold-billing/release", c.Key, params, v1releaseexternalpaymentgatethresholdcommitresponse)
	return v1releaseexternalpaymentgatethresholdcommitresponse, err
}

// Shortens the end date of a prepaid commit to terminate it earlier than originally scheduled. Use this endpoint when you need to cancel or reduce the duration of an existing prepaid commit. Only works with prepaid commit types and can only move the end date forward (earlier), not extend it.
//
// ### Usage guidelines:
// To extend commit end dates or make other comprehensive edits, use the 'edit commit' endpoint instead.
func (c v1CreditsAndCommitsService) UpdateCommitEndDate(ctx context.Context, params *V1CreditsAndCommitsUpdateCommitEndDateParams) (*V1UpdateCommitEndDateResponse, error) {
	if params == nil {
		params = &V1CreditsAndCommitsUpdateCommitEndDateParams{}
	}
	params.Context = ctx
	v1updatecommitenddateresponse := &V1UpdateCommitEndDateResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/customerCommits/updateEndDate", c.Key, params, v1updatecommitenddateresponse)
	return v1updatecommitenddateresponse, err
}

// Shortens the end date of an existing customer credit to terminate it earlier than originally scheduled. Only allows moving end dates forward (earlier), not extending them.
//
// Note: To extend credit end dates or make comprehensive edits, use the 'edit credit' endpoint instead.
func (c v1CreditsAndCommitsService) UpdateCreditEndDate(ctx context.Context, params *V1CreditsAndCommitsUpdateCreditEndDateParams) (*V1UpdateCreditEndDateResponse, error) {
	if params == nil {
		params = &V1CreditsAndCommitsUpdateCreditEndDateParams{}
	}
	params.Context = ctx
	v1updatecreditenddateresponse := &V1UpdateCreditEndDateResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contracts/customerCredits/updateEndDate", c.Key, params, v1updatecreditenddateresponse)
	return v1updatecreditenddateresponse, err
}
