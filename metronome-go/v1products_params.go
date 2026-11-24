//
// File generated from our OpenAPI spec
//

package metronome

// Archive a product. Any current rate cards associated with this product will continue to function as normal. However, it will no longer be available as an option for newly created rates. Once you archive a product, you can still retrieve it in the UI and API, but you cannot unarchive it.
type V1ProductsArchiveProductListItemParams struct {
	Params `form:"*"`
	// ID of the product to be archived
	ProductID *string `form:"product_id" json:"product_id"`
}
type V1ProductsCreateProductQuantityConversionParams struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor *Number `form:"conversion_factor" json:"conversion_factor"`
	// Optional name for this conversion.
	Name *string `form:"name" json:"name,omitempty"`
	// The operation to perform on the quantity
	Operation *string `form:"operation" json:"operation"`
}
type V1ProductsCreateProductQuantityRoundingParams struct {
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}

// Create a new product object. Products in Metronome represent your company's individual product or service offerings. A Product can be thought of as the basic unit of a line item on the invoice. This is analogous to SKUs or items in an ERP system. Give the product a meaningful name as they will appear on customer invoices.
type V1ProductsCreateProductParams struct {
	Params `form:"*"`
	// Required for USAGE products
	BillableMetricID *string `form:"billable_metric_id" json:"billable_metric_id,omitempty"`
	// Required for COMPOSITE products
	CompositeProductIDs []*string `form:"composite_product_ids" json:"composite_product_ids,omitempty"`
	// Determines what spend contributes to calculating this charge. When composite scope is contract, calculates composite charge based on applicable spend per contract. When composite scope is customer, calculates composite charge based on applicable spend across all customer contracts. Defaults to contract.
	CompositeScope *string `form:"composite_scope" json:"composite_scope,omitempty"`
	// Required for COMPOSITE products
	CompositeTags []*string         `form:"composite_tags" json:"composite_tags,omitempty"`
	CustomFields  map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// Beta feature only available for composite products. If true, products with $0 will not be included when computing composite usage. Defaults to false
	ExcludeFreeUsage *bool `form:"exclude_free_usage" json:"exclude_free_usage,omitempty"`
	// Only for composite products. If true, allows a composite to incorporate spend from other composite products. Defaults to false
	IncludeCompositeSpend *bool `form:"include_composite_spend" json:"include_composite_spend,omitempty"`
	// This field's availability is dependent on your client's configuration. Defaults to true.
	IsRefundable *bool `form:"is_refundable" json:"is_refundable,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteInternalItemID *string `form:"netsuite_internal_item_id" json:"netsuite_internal_item_id,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteOverageItemID *string                                          `form:"netsuite_overage_item_id" json:"netsuite_overage_item_id,omitempty"`
	PresentationGroupKey  []*string                                        `form:"presentation_group_key" json:"presentation_group_key,omitempty"`
	PricingGroupKey       []*string                                        `form:"pricing_group_key" json:"pricing_group_key,omitempty"`
	QuantityConversion    *V1ProductsCreateProductQuantityConversionParams `form:"quantity_conversion" json:"quantity_conversion,omitempty"`
	QuantityRounding      *V1ProductsCreateProductQuantityRoundingParams   `form:"quantity_rounding" json:"quantity_rounding,omitempty"`
	// Defines the breakdown behavior when calculating usage from SQL Billable Metrics. If set to 'service_period' (default), the usage will be evaluated once for all events the invoice service period and the usage will be applied at the last instant of the invoice. If set to 'hour', it will be broken down and evaluated for each hour. For most use cases, 'hour' is recommended. The setting has no effect for Streaming Billable Metrics.
	SQLBreakdownGranularity *string   `form:"sql_breakdown_granularity" json:"sql_breakdown_granularity,omitempty"`
	Tags                    []*string `form:"tags" json:"tags,omitempty"`
	Type                    *string   `form:"type" json:"type"`
}

// Retrieve a product by its ID, including all metadata and historical changes.
type V1ProductsGetProductParams struct {
	Params `form:"*"`
	ID     *string `form:"id" json:"id"`
}

// Get a paginated list of all products in your organization with their complete configuration, version history, and metadata. By default excludes archived products unless explicitly requested via the `archive_filter` parameter.
type V1ProductsListProductsParams struct {
	Params `form:"*"`
	// Filter options for the product list. If not provided, defaults to not archived.
	ArchiveFilter *string `form:"archive_filter" json:"archive_filter,omitempty"`
	// Max number of results that should be returned
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
}
type V1ProductsUpdateProductQuantityConversionParams struct {
	// The factor to multiply or divide the quantity by.
	ConversionFactor *Number `form:"conversion_factor" json:"conversion_factor"`
	// Optional name for this conversion.
	Name *string `form:"name" json:"name,omitempty"`
	// The operation to perform on the quantity
	Operation *string `form:"operation" json:"operation"`
}
type V1ProductsUpdateProductQuantityRoundingParams struct {
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}

// Updates a product's configuration while maintaining billing continuity for active customers. Use this endpoint to modify product names, metrics, pricing rules, and composite settings without disrupting ongoing billing cycles. Changes are scheduled using the starting_at timestamp, which must be on an hour boundary—set future dates to schedule updates ahead of time, or past dates for retroactive changes. Returns the updated product ID upon success.
//
// ### Usage guidance:
// - Product type cannot be changed after creation. For incorrect product types, create a new product and archive the original instead.
type V1ProductsUpdateProductParams struct {
	Params `form:"*"`
	// Available for USAGE products only. If not provided, defaults to product's current billable metric.
	BillableMetricID *string `form:"billable_metric_id" json:"billable_metric_id,omitempty"`
	// Available for COMPOSITE products only. If not provided, defaults to product's current composite_product_ids.
	CompositeProductIDs []*string `form:"composite_product_ids" json:"composite_product_ids,omitempty"`
	// Available for COMPOSITE products only. If not provided, defaults to product's current composite_tags.
	CompositeTags []*string `form:"composite_tags" json:"composite_tags,omitempty"`
	// Beta feature only available for composite products. If true, products with $0 will not be included when computing composite usage. Defaults to false
	ExcludeFreeUsage *bool `form:"exclude_free_usage" json:"exclude_free_usage,omitempty"`
	// Only for composite products. If true, allows a composite to incorporate spend from other composite products. Defaults to false
	IncludeCompositeSpend *bool `form:"include_composite_spend" json:"include_composite_spend,omitempty"`
	// Defaults to product's current refundability status. This field's availability is dependent on your client's configuration.
	IsRefundable *bool `form:"is_refundable" json:"is_refundable,omitempty"`
	// displayed on invoices. If not provided, defaults to product's current name.
	Name *string `form:"name" json:"name,omitempty"`
	// If not provided, defaults to product's current netsuite_internal_item_id. This field's availability is dependent on your client's configuration.
	NetsuiteInternalItemID *string `form:"netsuite_internal_item_id" json:"netsuite_internal_item_id,omitempty"`
	// Available for USAGE and COMPOSITE products only. If not provided, defaults to product's current netsuite_overage_item_id. This field's availability is dependent on your client's configuration.
	NetsuiteOverageItemID *string   `form:"netsuite_overage_item_id" json:"netsuite_overage_item_id,omitempty"`
	PresentationGroupKey  []*string `form:"presentation_group_key" json:"presentation_group_key,omitempty"`
	PricingGroupKey       []*string `form:"pricing_group_key" json:"pricing_group_key,omitempty"`
	// ID of the product to update
	ProductID          *string                                          `form:"product_id" json:"product_id"`
	QuantityConversion *V1ProductsUpdateProductQuantityConversionParams `form:"quantity_conversion" json:"quantity_conversion,omitempty"`
	QuantityRounding   *V1ProductsUpdateProductQuantityRoundingParams   `form:"quantity_rounding" json:"quantity_rounding,omitempty"`
	// Defines the breakdown behavior when calculating usage from SQL Billable Metrics. If set to 'service_period' (default), the usage will be evaluated once for all events the invoice service period and the usage will be applied at the last instant of the invoice. If set to 'hour', it will be broken down and evaluated for each hour. For most use cases, 'hour' is recommended. The setting has no effect for Streaming Billable Metrics.
	SQLBreakdownGranularity *string `form:"sql_breakdown_granularity" json:"sql_breakdown_granularity,omitempty"`
	// Timestamp representing when the update should go into effect. It must be on an hour boundary (e.g. 1:00, not 1:30).
	StartingAt *string `form:"starting_at" json:"starting_at"`
	// If not provided, defaults to product's current tags
	Tags []*string `form:"tags" json:"tags,omitempty"`
}
