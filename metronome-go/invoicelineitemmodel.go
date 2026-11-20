//
// File generated from our OpenAPI spec
//

package metronome

type InvoiceLineItemModelResellerType string

// List of values that InvoiceLineItemModelResellerType can take
const (
	InvoiceLineItemModelResellerTypeAWS           InvoiceLineItemModelResellerType = "AWS"
	InvoiceLineItemModelResellerTypeAWSPROSERVICE InvoiceLineItemModelResellerType = "AWS_PRO_SERVICE"
	InvoiceLineItemModelResellerTypeGCP           InvoiceLineItemModelResellerType = "GCP"
	InvoiceLineItemModelResellerTypeGCPPROSERVICE InvoiceLineItemModelResellerType = "GCP_PRO_SERVICE"
)

// Present on line items from invoices with type USAGE_CONSOLIDATED. Indicates the original customer, contract, invoice and line item from which this line item was copied.
type InvoiceLineItemModelOrigin struct {
	ContractID string `json:"contract_id"`
	CustomerID string `json:"customer_id"`
	InvoiceID  string `json:"invoice_id"`
	LineItemID string `json:"line_item_id"`
}
type InvoiceLineItemModel struct {
	APIResource
	// Details about the credit or commit that was applied to this line item. Only present on line items with product of `USAGE`, `SUBSCRIPTION`, `COMPOSITE`, or `CPU_CONVERSION` types.
	AppliedCommitOrCredit *AppliedCommitOrCreditModel `json:"applied_commit_or_credit,omitempty"`
	CommitCustomFields    map[string]string           `json:"commit_custom_fields,omitempty"`
	// For line items with product of `USAGE`, `SUBSCRIPTION`, `COMPOSITE`, or `CPU_CONVERSION` types, the ID of the credit or commit that was applied to this line item. For line items with product type of `FIXED`, the ID of the prepaid or postpaid commit that is being paid for.
	CommitID                   string `json:"commit_id,omitempty"`
	CommitNetsuiteItemID       string `json:"commit_netsuite_item_id,omitempty"`
	CommitNetsuiteSalesOrderID string `json:"commit_netsuite_sales_order_id,omitempty"`
	CommitSegmentID            string `json:"commit_segment_id,omitempty"`
	// `PrepaidCommit` (for commit types `PREPAID` and `CREDIT`) or `PostpaidCommit` (for commit type `POSTPAID`).
	CommitType           string            `json:"commit_type,omitempty"`
	CreditType           *CreditTypeModel  `json:"credit_type"`
	CustomFields         map[string]string `json:"custom_fields,omitempty"`
	DiscountCustomFields map[string]string `json:"discount_custom_fields,omitempty"`
	// ID of the discount applied to this line item.
	DiscountID string `json:"discount_id,omitempty"`
	// The line item's end date (exclusive).
	EndingBefore string `json:"ending_before,omitempty"`
	// Indicates whether the line item is prorated for `SUBSCRIPTION` type product.
	IsProrated bool `json:"is_prorated,omitempty"`
	// Only present for contract invoices and when the `include_list_prices` query parameter is set to true. This will include the list rate for the charge if applicable.  Only present for usage and subscription line items.
	//
	ListPrice *RateModel `json:"list_price,omitempty"`
	Metadata  string     `json:"metadata,omitempty"`
	Name      string     `json:"name"`
	// The end date for the billing period on the invoice.
	NetsuiteInvoiceBillingEnd string `json:"netsuite_invoice_billing_end,omitempty"`
	// The start date for the billing period on the invoice.
	NetsuiteInvoiceBillingStart string `json:"netsuite_invoice_billing_start,omitempty"`
	NetsuiteItemID              string `json:"netsuite_item_id,omitempty"`
	// Present on line items from invoices with type USAGE_CONSOLIDATED. Indicates the original customer, contract, invoice and line item from which this line item was copied.
	Origin *InvoiceLineItemModelOrigin `json:"origin,omitempty"`
	// Only present for line items paying for a postpaid commit true-up.
	PostpaidCommit *PostpaidCommitModel `json:"postpaid_commit,omitempty"`
	// Includes the presentation group values associated with this line item if presentation group keys are used.
	PresentationGroupValues map[string]*string `json:"presentation_group_values,omitempty"`
	// Includes the pricing group values associated with this line item if dimensional pricing is used.
	PricingGroupValues  map[string]string `json:"pricing_group_values,omitempty"`
	ProductCustomFields map[string]string `json:"product_custom_fields,omitempty"`
	// ID of the product associated with the line item.
	ProductID string `json:"product_id,omitempty"`
	// The current product tags associated with the line item's `product_id`.
	ProductTags []string `json:"product_tags,omitempty"`
	// The type of the line item's product. Possible values are `FixedProductListItem` (for `FIXED` type products), `UsageProductListItem` (for `USAGE` type products), `SubscriptionProductListItem` (for `SUBSCRIPTION` type products) or `CompositeProductListItem` (for `COMPOSITE` type products).  For scheduled charges, commit and credit payments, the value is `FixedProductListItem`.
	//
	ProductType                     string            `json:"product_type,omitempty"`
	ProfessionalServiceCustomFields map[string]string `json:"professional_service_custom_fields,omitempty"`
	ProfessionalServiceID           string            `json:"professional_service_id,omitempty"`
	// The quantity associated with the line item.
	Quantity Number `json:"quantity,omitempty"`
	// Present on applied commit line items for quantity-based commits. Represents the unit quantity deducted the commit.
	QuantityConsumed            Number                           `json:"quantity_consumed,omitempty"`
	ResellerType                InvoiceLineItemModelResellerType `json:"reseller_type,omitempty"`
	ScheduledChargeCustomFields map[string]string                `json:"scheduled_charge_custom_fields,omitempty"`
	// ID of scheduled charge.
	ScheduledChargeID string `json:"scheduled_charge_id,omitempty"`
	// The line item's start date (inclusive).
	StartingAt               string            `json:"starting_at,omitempty"`
	SubscriptionCustomFields map[string]string `json:"subscription_custom_fields,omitempty"`
	// ID of the subscription that this line item is associated with. Only present on line items with product of `SUBSCRIPTION` type.
	SubscriptionID string `json:"subscription_id,omitempty"`
	// Populated if the line item has a tiered price.
	Tier  *TierMetadataModel `json:"tier,omitempty"`
	Total Number             `json:"total"`
	// The type of line item.
	//
	// - `scheduled`: Line item is associated with a scheduled charge. View the scheduled_charge_id on the line item.
	// - `commit_purchase`: Line item is associated with a payment for a prepaid commit. View the commit_id on the line item.
	// - `usage`: Line item is associated with a usage product or composite product. View the product_id on the line item to determine which product.
	// - `subscription`: Line item is associated with a subscription. e.g. monthly recurring payment for an in-advance subscription.
	// - `applied_commit_or_credit`: On metronome invoices, applied commits and credits are associated with their own line items. These line items have negative totals. Use the applied_commit_or_credit object on the line item to understand the id of the applied commit or credit, and its type. Note that the application of a postpaid commit is associated with a line item, but the total on the line item is not included in the invoice's total as postpaid commits are paid in-arrears.
	// - `cpu_conversion`: Line item converting between a custom pricing unit and fiat currency, using the conversion rate set on the rate card. This line item will appear when there are products priced in custom pricing units, and there is insufficient prepaid commit/credit in that custom pricing unit to fully cover the spend. Then, the outstanding spend in custom pricing units will be converted to fiat currency using a cpu_conversion line item.
	//
	Type string `json:"type"`
	// The unit price associated with the line item.
	UnitPrice Number `json:"unit_price,omitempty"`
}
