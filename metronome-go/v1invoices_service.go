//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v1InvoicesService is used to invoke v1invoices related APIs.
type v1InvoicesService struct {
	B   Backend
	Key string
}

// Creates an prorated invoice for a seat addition. As an alternative to this endpoint, you can elect to use automatic seat invoicing feature. Metronome will check for new seat usage every hour and automatically invoice for any new seats. For newly created active customer plans, there will be up to 4 hour delay before the first automatic seat invoice is generated.
func (c v1InvoicesService) ChargeSeats(ctx context.Context, id string, params *V1InvoicesChargeSeatsParams) (*V1ChargeSeatsResponse, error) {
	if params == nil {
		params = &V1InvoicesChargeSeatsParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/customers/%s/invoices/invoice_seats", id)
	v1chargeseatsresponse := &V1ChargeSeatsResponse{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, v1chargeseatsresponse)
	return v1chargeseatsresponse, err
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
func (c v1InvoicesService) GetInvoice(ctx context.Context, id string, params *V1InvoicesGetInvoiceParams) (*V1GetInvoiceResponse, error) {
	if params == nil {
		params = &V1InvoicesGetInvoiceParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/invoices/%s", StringValue(params.CustomerID), id)
	v1getinvoiceresponse := &V1GetInvoiceResponse{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, v1getinvoiceresponse)
	return v1getinvoiceresponse, err
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
func (c v1InvoicesService) GetInvoicePDF(ctx context.Context, id string, params *V1InvoicesGetInvoicePDFParams) (*V1GetInvoicePDFResponse, error) {
	if params == nil {
		params = &V1InvoicesGetInvoicePDFParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/invoices/%s/pdf", StringValue(params.CustomerID), id)
	v1getinvoicepdfresponse := &V1GetInvoicePDFResponse{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, v1getinvoicepdfresponse)
	return v1getinvoicepdfresponse, err
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
func (c v1InvoicesService) ListBreakdownInvoices(ctx context.Context, id string, params *V1InvoicesListBreakdownInvoicesParams) (*V1ListBreakdownInvoicesResponse, error) {
	if params == nil {
		params = &V1InvoicesListBreakdownInvoicesParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/customers/%s/invoices/breakdowns", id)
	v1listbreakdowninvoicesresponse := &V1ListBreakdownInvoicesResponse{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, v1listbreakdowninvoicesresponse)
	return v1listbreakdowninvoicesresponse, err
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
func (c v1InvoicesService) ListInvoices(ctx context.Context, id string, params *V1InvoicesListInvoicesParams) (*V1ListInvoicesResponse, error) {
	if params == nil {
		params = &V1InvoicesListInvoicesParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/customers/%s/invoices", id)
	v1listinvoicesresponse := &V1ListInvoicesResponse{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, v1listinvoicesresponse)
	return v1listinvoicesresponse, err
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
func (c v1InvoicesService) ListSpendBreakdownInvoices(ctx context.Context, id string, params *V1InvoicesListSpendBreakdownInvoicesParams) (*V1ListSpendBreakdownInvoicesResponse, error) {
	if params == nil {
		params = &V1InvoicesListSpendBreakdownInvoicesParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/customers/%s/invoices/spend-breakdowns", id)
	v1listspendbreakdowninvoicesresponse := &V1ListSpendBreakdownInvoicesResponse{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, v1listspendbreakdowninvoicesresponse)
	return v1listspendbreakdowninvoicesresponse, err
}

// Preview how a set of events will affect a customer's invoices. Generates draft invoices for a customer using their current contract configuration and the provided events.  This is useful for testing how new events will affect the customer's invoices before they are actually processed. Customers on contracts with SQL billable metrics are not supported.
func (c v1InvoicesService) PreviewCustomerEvents(ctx context.Context, id string, params *V1InvoicesPreviewCustomerEventsParams) (*V1PreviewCustomerEventsResponse, error) {
	if params == nil {
		params = &V1InvoicesPreviewCustomerEventsParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/customers/%s/previewEvents", id)
	v1previewcustomereventsresponse := &V1PreviewCustomerEventsResponse{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, v1previewcustomereventsresponse)
	return v1previewcustomereventsresponse, err
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
func (c v1InvoicesService) RegenerateInvoice(ctx context.Context, params *V1InvoicesRegenerateInvoiceParams) (*V1RegenerateInvoiceResponse, error) {
	if params == nil {
		params = &V1InvoicesRegenerateInvoiceParams{}
	}
	params.Context = ctx
	v1regenerateinvoiceresponse := &V1RegenerateInvoiceResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/invoices/regenerate", c.Key, params, v1regenerateinvoiceresponse)
	return v1regenerateinvoiceresponse, err
}

// Permanently cancels an invoice by setting its status to voided, preventing collection and removing it from customer billing. Use this to correct billing errors, cancel incorrect charges, or handle disputed invoices that should not be collected. Returns the voided invoice ID with the status change applied immediately to stop any payment processing.
func (c v1InvoicesService) VoidInvoice(ctx context.Context, params *V1InvoicesVoidInvoiceParams) (*V1VoidInvoiceResponse, error) {
	if params == nil {
		params = &V1InvoicesVoidInvoiceParams{}
	}
	params.Context = ctx
	v1voidinvoiceresponse := &V1VoidInvoiceResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/invoices/void", c.Key, params, v1voidinvoiceresponse)
	return v1voidinvoiceresponse, err
}
