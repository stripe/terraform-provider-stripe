//
// File generated from our OpenAPI spec
//

package metronome

type V1InvoicesChargeSeatsSeatChargeParams struct {
	ChargeID         *string `form:"charge_id" json:"charge_id"`
	CurrentSeatCount *int64  `form:"current_seat_count" json:"current_seat_count,omitempty"`
	SeatCount        *int64  `form:"seat_count" json:"seat_count"`
}

// Creates an prorated invoice for a seat addition. As an alternative to this endpoint, you can elect to use automatic seat invoicing feature. Metronome will check for new seat usage every hour and automatically invoice for any new seats. For newly created active customer plans, there will be up to 4 hour delay before the first automatic seat invoice is generated.
type V1InvoicesChargeSeatsParams struct {
	Params      `form:"*"`
	SeatCharges []*V1InvoicesChargeSeatsSeatChargeParams `form:"seat_charges" json:"seat_charges"`
}

// Retrieve detailed information for a specific invoice by its unique identifier. This endpoint returns comprehensive invoice data including line items, applied credits, totals, and billing period details for both finalized and draft invoices.
//
// ### Use this endpoint to:
// - Display historical invoice details in customer-facing dashboards or billing portals.
// - Retrieve current month draft invoices to show customers their month-to-date spend.
// - Access finalized invoices for historical billing records and payment reconciliation.
// - Validate customer pricing and credit applications for customer support queries.
//
// ### Key response fields:
// Invoice status (DRAFT, FINALIZED, VOID)
// Billing period start and end dates
// Total amount and amount due after credits
// Detailed line items broken down by:
// - Customer and contract information
// - Invoice line item type
// - Product/service name and ID
// - Quantity consumed
// - Unit and total price
// - Time period for usage-based charges
// - Applied credits or prepaid commitments
//
// ### Usage guidelines:
// - Draft invoices update in real-time as usage is reported and may change before finalization
// - The response includes both usage-based line items (e.g., API calls, data processed) and scheduled charges (e.g., monthly subscriptions, commitment fees)
// - Credit and commitment applications are shown as separate line items with negative amounts
// - For voided invoices, the response will indicate VOID status but retain all original line item details
type V1InvoicesGetInvoiceParams struct {
	Params     `form:"*"`
	CustomerID *string `form:"-" json:"-"` // Included in URL
	// If set, all zero quantity line items will be filtered out of the response
	SkipZeroQtyLineItems *bool `form:"skip_zero_qty_line_items" json:"skip_zero_qty_line_items,omitempty"`
}

// Retrieve a PDF version of a specific invoice by its unique identifier. This endpoint generates a professionally formatted invoice document suitable for sharing with customers, accounting teams, or for record-keeping purposes.
//
// ### Use this endpoint to:
// - Provide customers with downloadable or emailable copies of their invoices
// - Support accounting and finance teams with official billing documents
// - Maintain accurate records of billing transactions for audits and compliance
//
// ### Key response details:
// - The response is a binary PDF file representing the full invoice
// - The PDF includes all standard invoice information such as line items, totals, billing period, and customer details
// - The document is formatted for clarity and professionalism, suitable for official use
//
// ### Usage guidelines:
// - Ensure the `invoice_id` corresponds to an existing invoice for the specified `customer_id`
// - The PDF is generated on-demand; frequent requests for the same invoice may impact performance
// - Use appropriate headers to handle the binary response in your application (e.g., setting `Content-Type: application/pdf`)
type V1InvoicesGetInvoicePDFParams struct {
	Params     `form:"*"`
	CustomerID *string `form:"-" json:"-"` // Included in URL
}

// Retrieve granular time-series breakdowns of invoice data at hourly or daily intervals. This endpoint transforms standard invoices into detailed timelines, enabling you to track usage patterns, identify consumption spikes, and provide customers with transparency into their billing details throughout the billing period.
//
// ### Use this endpoint to:
// - Build usage analytics dashboards showing daily or hourly consumption trends
// - Identify peak usage periods for capacity planning and cost optimization
// - Generate detailed billing reports for finance teams and customer success
// - Troubleshoot billing disputes by examining usage patterns at specific times
// - Power real-time cost monitoring and alerting systems
//
// ### Key response fields:
// An array of BreakdownInvoice objects, each containing:
// - All standard invoice fields (ID, customer, commit, line items, totals, status)
// - Line items with quantities and costs for that specific period
// - `breakdown_start_timestamp`: Start of the specific time window
// - `breakdown_end_timestamp`: End of the specific time window
// - `next_page`: Pagination cursor for large result sets
//
// ### Usage guidelines:
// - Time granularity: Set `window_size` to hour or day based on your analysis needs
// - Response limits: Daily breakdowns return up to 35 days; hourly breakdowns return up to 24 hours per request
// - Date filtering: Use `starting_on` and `ending_before` to focus on specific periods
// - Performance: For large date ranges, use pagination to retrieve all data efficiently
// - Backdated usage: If usage events arrive after invoice finalization, breakdowns will reflect the updated usage
// - Zero quantity filtering: Use `skip_zero_qty_line_items=true` to exclude periods with no usage
type V1InvoicesListBreakdownInvoicesParams struct {
	Params `form:"*"`
	// Only return invoices for the specified credit type
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// RFC 3339 timestamp. Breakdowns will only be returned for time windows that end on or before this time.
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	// Max number of results that should be returned. For daily breakdowns, the response can return up to 35 days worth of breakdowns. For hourly breakdowns, the response can return up to 24 hours. If there are more results, a cursor to the next page is returned.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
	// If set, all zero quantity line items will be filtered out of the response
	SkipZeroQtyLineItems *bool `form:"skip_zero_qty_line_items" json:"skip_zero_qty_line_items,omitempty"`
	// Invoice sort order by issued_at, e.g. date_asc or date_desc.  Defaults to date_asc.
	Sort *string `form:"sort" json:"sort,omitempty"`
	// RFC 3339 timestamp. Breakdowns will only be returned for time windows that start on or after this time.
	StartingOn *string `form:"starting_on" json:"starting_on"`
	// Invoice status, e.g. DRAFT or FINALIZED
	Status *string `form:"status" json:"status,omitempty"`
	// The granularity of the breakdowns to return. Defaults to day.
	WindowSize *string `form:"window_size" json:"window_size,omitempty"`
}

// Retrieves a paginated list of invoices for a specific customer, with flexible filtering options to narrow results by status, date range, credit type, and more. This endpoint provides a comprehensive view of a customer's billing history and current charges, supporting both real-time billing dashboards and historical reporting needs.
//
// ### Use this endpoint to:
// - Display historical invoice details in customer-facing dashboards or billing portals.
// - Retrieve current month draft invoices to show customers their month-to-date spend.
// - Access finalized invoices for historical billing records and payment reconciliation.
// - Validate customer pricing and credit applications for customer support queries.
// - Generate financial reports by filtering invoices within specific date ranges
//
// ### Key response fields:
// Array of invoice objects containing:
// - Invoice ID and status (DRAFT, FINALIZED, VOID)
// - Invoice type (USAGE, SCHEDULED)
// - Billing period start and end dates
// - Issue date and due date
// - Total amount, subtotal, and amount due
// - Applied credits summary
// - Contract ID reference
// - External billing provider status (if integrated with Stripe, etc.)
// - Pagination metadata `next_page` cursor
//
// ### Usage guidelines:
// - The endpoint returns invoice summaries; use the Get Invoice endpoint for detailed line items
// - Draft invoices are continuously updated as new usage is reported and will show real-time spend
// - Results are ordered by creation date descending by default (newest first)
// - When filtering by date range, the filter applies to the billing period, not the issue date
// - For customers with many invoices, implement pagination to ensure all results are retrieved
// External billing provider statuses (like Stripe payment status) are included when applicable
// - Voided invoices are included in results by default unless filtered out by status
type V1InvoicesListInvoicesParams struct {
	Params `form:"*"`
	// Only return invoices for the specified contract
	ContractID *string `form:"contract_id" json:"contract_id,omitempty"`
	// Only return invoices for the specified credit type
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// RFC 3339 timestamp (exclusive). Invoices will only be returned for billing periods that end before this time.
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// When true, includes retired commit invoices alongside active invoices. Defaults to false.
	IncludeRetiredCommitInvoices *bool `form:"include_retired_commit_invoices" json:"include_retired_commit_invoices,omitempty"`
	// Max number of results that should be returned
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
	// If set, all zero quantity line items will be filtered out of the response
	SkipZeroQtyLineItems *bool `form:"skip_zero_qty_line_items" json:"skip_zero_qty_line_items,omitempty"`
	// Invoice sort order by issued_at, e.g. date_asc or date_desc.  Defaults to date_asc.
	Sort *string `form:"sort" json:"sort,omitempty"`
	// RFC 3339 timestamp (inclusive). Invoices will only be returned for billing periods that start at or after this time.
	StartingOn *string `form:"starting_on" json:"starting_on,omitempty"`
	// Invoice status, e.g. DRAFT, FINALIZED, or VOID
	Status *string `form:"status" json:"status,omitempty"`
	// Filter invoices by type. Defaults to returning all invoice types.
	Type *string `form:"type" json:"type,omitempty"`
	// Indicates that this API request was triggered by a webhook notification with the provided ID.
	WebhookNotificationID *string `form:"webhook_notification_id" json:"webhook_notification_id,omitempty"`
}

// Granularly analyze customer spend patterns by dynamically slicing and dicing costs across any dimension. This endpoint empowers you to break down spending by granular properties like user, organization, model, region, or any custom event property—even if these aren't the default groupings on your invoices. Unlike standard invoice breakdowns, this endpoint focuses purely on spend analysis, making helpful for building powerful cost analytics dashboards that show spend before credit/commit application.
//
// ### Use this endpoint to:
// - Identify cost drivers: Pinpoint which users, teams, or resources are driving the most spend
// - Build usage analytics dashboards: Let customers explore their costs by any event property (user_id, org_id, model_id, region, etc.)
// - Enable showback/chargeback: Allocate costs to specific departments, projects, or cost centers
// - Detect anomalies: Find unexpected spending patterns by analyzing costs across different dimensions
// - Optimize resource usage: Help customers identify underutilized or over-provisioned resources
// - Support multi-tenancy: Show spending breakdowns for specific organizations within a single account
// - Create custom reports: Generate executive dashboards with spending by any business-relevant dimension
//
// ### Key response fields:
// Spend-focused invoice data with:
// - Pure spend information: No commits or credits—just raw spending data for cleaner analysis
// - Dynamic grouping: Line items grouped by your specified group_keys (overriding default presentation groups)
// - Filtered results: Only line items matching your group_filters criteria
// - Flexible time windows: Daily, hourly, or full-period (none) breakdowns
// - Complete line item details: Including quantities, unit prices, and custom presentation group values
//
// ### Usage guidelines:
// - Group key setup: All keys used in group_keys, group_filters, and pricing groups must exist in the same compound group key on the billable metric
// - Supported window sizes: hour, day, or none (for full period analysis)
// - Filtering power: Use group_filters to focus on specific values (e.g., only show data for specific user_ids)
// - Override flexibility: Change how costs are grouped without affecting actual invoicing
//
// Limitations:
// - Cannot override group keys when using:
//   - MAX aggregation billable metrics
//   - Tiered pricing
//   - Quantity rounding
//   - Commit-specific overrides
//   - Overrides on presentation group values
type V1InvoicesListSpendBreakdownInvoicesParams struct {
	Params `form:"*"`
	// If provided, only invoices with the specified credit type id will be included in the response.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// RFC 3339 timestamp. Breakdowns will only be returned for time windows that end on or before this time.
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	// An object where the keys are the group keys and the values are arrays of group values. If the values is an empty array, the returned usage data will be aggregated across all values for that key. The value will be ignored and the default keys will be used if: billable metric is MAX type; product uses tiered pricing model; product has quantity rounding enabled; there are contract overrides based on presentation group keys
	//
	GroupFilters map[string][]*string `form:"group_filters" json:"group_filters,omitempty"`
	// A list of keys that can be used to additionally segment the values of the billable metric when making usage queries. Must be valid keys that are present in the billable metrics or an empty list. The value will be ignored and the default keys will be used if: billable metric is MAX type; product uses tiered pricing model; product has quantity rounding enabled; there are contract overrides based on presentation group keys
	//
	GroupKeys []*string `form:"group_keys" json:"group_keys,omitempty"`
	// If set, list prices will be returned for each contract usage and subscription line item.
	IncludeListPrices *bool `form:"include_list_prices" json:"include_list_prices,omitempty"`
	// Max number of results that should be returned. For daily and "none" breakdowns, the response can return up to 35 days worth of breakdowns. If there are more results, a cursor to the next page is returned.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
	// If set, all zero quantity line items will be filtered out of the response.
	SkipZeroQtyLineItems *bool `form:"skip_zero_qty_line_items" json:"skip_zero_qty_line_items,omitempty"`
	// Invoice sort order by issued_at, e.g. date_asc or date_desc.  Defaults to date_asc.
	Sort *string `form:"sort" json:"sort,omitempty"`
	// RFC 3339 timestamp. Breakdowns will only be returned for time windows that start on or after this time.
	StartingOn *string `form:"starting_on" json:"starting_on"`
	// The granularity of the breakdowns to return. Defaults to "day".
	WindowSize *string `form:"window_size" json:"window_size,omitempty"`
}

// Array of usage events to include in the preview calculation. Must contain at least one event in `merge` mode.
type V1InvoicesPreviewCustomerEventsEventParams struct {
	EventType  *string         `form:"event_type" json:"event_type"`
	Properties map[string]JSON `form:"properties" json:"properties,omitempty"`
	// RFC 3339 formatted. If not provided, the current time will be used.
	Timestamp *string `form:"timestamp" json:"timestamp,omitempty"`
	// Optional unique identifier for event deduplication. When provided, preview events are automatically deduplicated against historical events from the past 34 days.  Duplicate transaction IDs within the same request will return an error.
	//
	TransactionID *string `form:"transaction_id" json:"transaction_id,omitempty"`
}

// Preview how a set of events will affect a customer's invoices. Generates draft invoices for a customer using their current contract configuration and the provided events.  This is useful for testing how new events will affect the customer's invoices before they are actually processed. Customers on contracts with SQL billable metrics are not supported.
type V1InvoicesPreviewCustomerEventsParams struct {
	Params `form:"*"`
	// Array of usage events to include in the preview calculation. Must contain at least one event in `merge` mode.
	//
	Events []*V1InvoicesPreviewCustomerEventsEventParams `form:"events" json:"events"`
	// Controls how the provided events are combined with existing usage data. Use `replace` to calculate the preview as if these are the only events for the customer, ignoring all historical usage.  Use `merge` to combine these events with the customer's existing usage.  Defaults to `replace`.
	//
	Mode *string `form:"mode" json:"mode,omitempty"`
	// When `true`, line items with zero quantity are excluded from the response.
	SkipZeroQtyLineItems *bool `form:"skip_zero_qty_line_items" json:"skip_zero_qty_line_items,omitempty"`
}

// This endpoint regenerates a voided invoice and recalculates the invoice based on up-to-date rates, available balances, and other fees regardless of the billing period.
//
// ### Use this endpoint to:
// Recalculate an invoice with updated rate terms, available balance, and fees to correct billing disputes or discrepancies
//
// ### Key response fields:
// The regenerated invoice id, which is distinct from the previously voided invoice.
//
// ### Usage guidelines:
// If an invoice is attached to a contract with a billing provider on it, the regenerated invoice will be distributed based on the configuration.
type V1InvoicesRegenerateInvoiceParams struct {
	Params `form:"*"`
	// The invoice id to regenerate
	ID *string `form:"id" json:"id"`
}

// Permanently cancels an invoice by setting its status to voided, preventing collection and removing it from customer billing. Use this to correct billing errors, cancel incorrect charges, or handle disputed invoices that should not be collected. Returns the voided invoice ID with the status change applied immediately to stop any payment processing.
type V1InvoicesVoidInvoiceParams struct {
	Params `form:"*"`
	// The invoice id to void
	ID *string `form:"id" json:"id"`
}
