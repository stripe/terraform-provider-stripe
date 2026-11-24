//
// File generated from our OpenAPI spec
//

package metronome

type V1ContractsAmendContractCommitAccessScheduleScheduleItemParams struct {
	Amount *Number `form:"amount" json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	// RFC 3339 timestamp (inclusive)
	StartingAt *string `form:"starting_at" json:"starting_at"`
}

// Required: Schedule for distributing the commit to the customer. For "POSTPAID" commits only one schedule item is allowed and amount must match invoice_schedule total.
type V1ContractsAmendContractCommitAccessScheduleParams struct {
	// Determines how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used. Defaults to `SPEND` if omitted.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID  *string                                                           `form:"credit_type_id" json:"credit_type_id,omitempty"`
	ScheduleItems []*V1ContractsAmendContractCommitAccessScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items"`
}

// Either provide amount or provide both unit_price and quantity.
type V1ContractsAmendContractCommitInvoiceScheduleScheduleItemParams struct {
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
type V1ContractsAmendContractCommitInvoiceScheduleRecurringScheduleParams struct {
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

// Required for "POSTPAID" commits: the true up invoice will be generated at this time and only one schedule item is allowed; the total must match access_schedule amount. Optional for "PREPAID" commits: if not provided, this will be a "complimentary" commit with no invoice.
type V1ContractsAmendContractCommitInvoiceScheduleParams struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is only applicable to commit invoice schedules. If true, this schedule will not generate an invoice.
	DoNotInvoice *bool `form:"do_not_invoice" json:"do_not_invoice,omitempty"`
	// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
	RecurringSchedule *V1ContractsAmendContractCommitInvoiceScheduleRecurringScheduleParams `form:"recurring_schedule" json:"recurring_schedule,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []*V1ContractsAmendContractCommitInvoiceScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items,omitempty"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsAmendContractCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsAmendContractCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsAmendContractCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Optional configuration for commit hierarchy access control
type V1ContractsAmendContractCommitHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}

// Optional attributes for spend tracker integration. Immutable after creation.
type V1ContractsAmendContractCommitSpendTrackerAttributesParams struct {
	// If true, this commit will be included in spend trackers with discounted set to DISCOUNTED_ONLY
	CountsAsDiscounted *bool `form:"counts_as_discounted" json:"counts_as_discounted"`
}
type V1ContractsAmendContractCommitParams struct {
	// Required: Schedule for distributing the commit to the customer. For "POSTPAID" commits only one schedule item is allowed and amount must match invoice_schedule total.
	AccessSchedule *V1ContractsAmendContractCommitAccessScheduleParams `form:"access_schedule" json:"access_schedule,omitempty"`
	// (DEPRECATED) Use access_schedule and invoice_schedule instead.
	Amount *Number `form:"amount" json:"amount,omitempty"`
	// Which products the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductTags []*string         `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	CustomFields          map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// Used only in UI/API. It is not exposed to end customers.
	Description *string `form:"description" json:"description,omitempty"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration *V1ContractsAmendContractCommitHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// Required for "POSTPAID" commits: the true up invoice will be generated at this time and only one schedule item is allowed; the total must match access_schedule amount. Optional for "PREPAID" commits: if not provided, this will be a "complimentary" commit with no invoice.
	InvoiceSchedule *V1ContractsAmendContractCommitInvoiceScheduleParams `form:"invoice_schedule" json:"invoice_schedule,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// If multiple commits are applicable, the one with the lower priority will apply first.
	Priority  *Number `form:"priority" json:"priority,omitempty"`
	ProductID *string `form:"product_id" json:"product_id"`
	RateType  *string `form:"rate_type" json:"rate_type,omitempty"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsAmendContractCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
	// Optional attributes for spend tracker integration. Immutable after creation.
	SpendTrackerAttributes *V1ContractsAmendContractCommitSpendTrackerAttributesParams `form:"spend_tracker_attributes" json:"spend_tracker_attributes,omitempty"`
	// A temporary ID for the commit that can be used to reference the commit for commit specific overrides.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
	Type        *string `form:"type" json:"type"`
}
type V1ContractsAmendContractCreditAccessScheduleScheduleItemParams struct {
	Amount *Number `form:"amount" json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	// RFC 3339 timestamp (inclusive)
	StartingAt *string `form:"starting_at" json:"starting_at"`
}

// Schedule for distributing the credit to the customer.
type V1ContractsAmendContractCreditAccessScheduleParams struct {
	// Determines how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used. Defaults to `SPEND` if omitted.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID  *string                                                           `form:"credit_type_id" json:"credit_type_id,omitempty"`
	ScheduleItems []*V1ContractsAmendContractCreditAccessScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsAmendContractCreditSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsAmendContractCreditSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsAmendContractCreditSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Optional configuration for credit hierarchy access control
type V1ContractsAmendContractCreditHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}
type V1ContractsAmendContractCreditParams struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule *V1ContractsAmendContractCreditAccessScheduleParams `form:"access_schedule" json:"access_schedule"`
	// Which products the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductTags []*string         `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	CustomFields          map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// Used only in UI/API. It is not exposed to end customers.
	Description *string `form:"description" json:"description,omitempty"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration *V1ContractsAmendContractCreditHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// If multiple credits are applicable, the one with the lower priority will apply first.
	Priority  *Number `form:"priority" json:"priority,omitempty"`
	ProductID *string `form:"product_id" json:"product_id"`
	RateType  *string `form:"rate_type" json:"rate_type,omitempty"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsAmendContractCreditSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
}

// Only set for TIERED rate_type.
type V1ContractsAmendContractOverrideOverwriteRateTierParams struct {
	Price *Number `form:"price" json:"price"`
	Size  *Number `form:"size" json:"size,omitempty"`
}
type V1ContractsAmendContractOverrideOverwriteRateMinimumConfigParams struct {
	Minimum *Number `form:"minimum" json:"minimum"`
}

// Required for OVERWRITE type.
type V1ContractsAmendContractOverrideOverwriteRateParams struct {
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate processors.
	CustomRate map[string]JSON `form:"custom_rate" json:"custom_rate,omitempty"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be set to true.
	IsProrated    *bool                                                             `form:"is_prorated" json:"is_prorated,omitempty"`
	MinimumConfig *V1ContractsAmendContractOverrideOverwriteRateMinimumConfigParams `form:"minimum_config" json:"minimum_config,omitempty"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price *Number `form:"price" json:"price,omitempty"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity *Number `form:"quantity" json:"quantity,omitempty"`
	RateType *string `form:"rate_type" json:"rate_type"`
	// Only set for TIERED rate_type.
	Tiers []*V1ContractsAmendContractOverrideOverwriteRateTierParams `form:"tiers" json:"tiers,omitempty"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsAmendContractOverrideOverrideSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Cannot be used in conjunction with product_id or applicable_product_tags. If provided, the override will apply to all products with the specified specifiers.
type V1ContractsAmendContractOverrideOverrideSpecifierParams struct {
	// Can only be used for commit specific overrides. Must be used in conjunction with one of `product_id`, `product_tags`, `pricing_group_values`, or `presentation_group_values`. Must be used instead of both `commit_ids` and `recurring_commit_ids` If provided, the override will apply to any specified commit, credit, recurring commit or recurring credit IDs.
	AnyCommitOrCreditIDs []*string `form:"any_commit_or_credit_ids" json:"any_commit_or_credit_ids,omitempty"`
	BillingFrequency     *string   `form:"billing_frequency" json:"billing_frequency,omitempty"`
	// Can only be used for commit specific overrides. Must be used in conjunction with one of `product_id`, `product_tags`, `pricing_group_values`, or `presentation_group_values`. If provided, the override will only apply to the specified commits. If not provided, the override will apply to all commits.
	CommitIDs []*string `form:"commit_ids" json:"commit_ids,omitempty"`
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsAmendContractOverrideOverrideSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// A map of group names to values. The override will only apply to line items with the specified presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// A map of pricing group names to values. The override will only apply to products with the specified pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the override will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the override will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
	// Can only be used for commit specific overrides. Must be used in conjunction with one of `product_id`, `product_tags`, `pricing_group_values`, or `presentation_group_values`. If provided, the override will only apply to commits created by the specified recurring commit ids.
	RecurringCommitIDs []*string `form:"recurring_commit_ids" json:"recurring_commit_ids,omitempty"`
}

// Required for TIERED type. Must have at least one tier.
type V1ContractsAmendContractOverrideTierParams struct {
	Multiplier *Number `form:"multiplier" json:"multiplier"`
	Size       *Number `form:"size" json:"size,omitempty"`
}
type V1ContractsAmendContractOverrideParams struct {
	// tags identifying products whose rates are being overridden. Cannot be used in conjunction with override_specifiers.
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	Entitled     *bool   `form:"entitled" json:"entitled,omitempty"`
	// Indicates whether the override should only apply to commits. Defaults to `false`. If `true` you can specify relevant commits in `override_specifiers` by passing `commit_ids`, `recurring_commit_ids`, or `any_commit_or_credit_ids`.  If you do not specify any of these fields, the override will apply when consuming any prepaid commit, postpaid commit, or credit
	IsCommitSpecific *bool `form:"is_commit_specific" json:"is_commit_specific,omitempty"`
	// Required for MULTIPLIER type. Must be >=0.
	Multiplier *Number `form:"multiplier" json:"multiplier,omitempty"`
	// Cannot be used in conjunction with product_id or applicable_product_tags. If provided, the override will apply to all products with the specified specifiers.
	OverrideSpecifiers []*V1ContractsAmendContractOverrideOverrideSpecifierParams `form:"override_specifiers" json:"override_specifiers,omitempty"`
	// Required for OVERWRITE type.
	OverwriteRate *V1ContractsAmendContractOverrideOverwriteRateParams `form:"overwrite_rate" json:"overwrite_rate,omitempty"`
	// Required for EXPLICIT multiplier prioritization scheme and all TIERED overrides. Under EXPLICIT prioritization, overwrites are prioritized first, and then tiered and multiplier overrides are prioritized by their priority value (lowest first). Must be > 0.
	Priority *Number `form:"priority" json:"priority,omitempty"`
	// ID of the product whose rate is being overridden. Cannot be used in conjunction with override_specifiers.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// RFC 3339 timestamp indicating when the override will start applying (inclusive)
	StartingAt *string `form:"starting_at" json:"starting_at"`
	// Indicates whether the override applies to commit rates or list rates. Can only be used for overrides that have `is_commit_specific` set to `true`. Defaults to `"LIST_RATE"`.
	Target *string `form:"target" json:"target,omitempty"`
	// Required for TIERED type. Must have at least one tier.
	Tiers []*V1ContractsAmendContractOverrideTierParams `form:"tiers" json:"tiers,omitempty"`
	// Overwrites are prioritized over multipliers and tiered overrides.
	Type *string `form:"type" json:"type,omitempty"`
}

// Either provide amount or provide both unit_price and quantity.
type V1ContractsAmendContractDiscountScheduleScheduleItemParams struct {
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
type V1ContractsAmendContractDiscountScheduleRecurringScheduleParams struct {
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
type V1ContractsAmendContractDiscountScheduleParams struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is only applicable to commit invoice schedules. If true, this schedule will not generate an invoice.
	DoNotInvoice *bool `form:"do_not_invoice" json:"do_not_invoice,omitempty"`
	// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
	RecurringSchedule *V1ContractsAmendContractDiscountScheduleRecurringScheduleParams `form:"recurring_schedule" json:"recurring_schedule,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []*V1ContractsAmendContractDiscountScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items,omitempty"`
}

// This field's availability is dependent on your client's configuration.
type V1ContractsAmendContractDiscountParams struct {
	CustomFields map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string                                         `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	ProductID            *string                                         `form:"product_id" json:"product_id"`
	Schedule             *V1ContractsAmendContractDiscountScheduleParams `form:"schedule" json:"schedule"`
}

// This field's availability is dependent on your client's configuration.
type V1ContractsAmendContractProfessionalServiceParams struct {
	CustomFields map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	Description  *string           `form:"description" json:"description,omitempty"`
	// Maximum amount for the term.
	MaxAmount *Number `form:"max_amount" json:"max_amount"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	ProductID            *string `form:"product_id" json:"product_id"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the amount.
	Quantity *Number `form:"quantity" json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified.
	UnitPrice *Number `form:"unit_price" json:"unit_price"`
}
type V1ContractsAmendContractResellerRoyaltyAwsOptionsParams struct {
	AwsAccountNumber    *string `form:"aws_account_number" json:"aws_account_number,omitempty"`
	AwsOfferID          *string `form:"aws_offer_id" json:"aws_offer_id,omitempty"`
	AwsPayerReferenceID *string `form:"aws_payer_reference_id" json:"aws_payer_reference_id,omitempty"`
}
type V1ContractsAmendContractResellerRoyaltyGcpOptionsParams struct {
	GcpAccountID *string `form:"gcp_account_id" json:"gcp_account_id,omitempty"`
	GcpOfferID   *string `form:"gcp_offer_id" json:"gcp_offer_id,omitempty"`
}

// This field's availability is dependent on your client's configuration.
type V1ContractsAmendContractResellerRoyaltyParams struct {
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductTags []*string                                                `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	AwsOptions            *V1ContractsAmendContractResellerRoyaltyAwsOptionsParams `form:"aws_options" json:"aws_options,omitempty"`
	// Use null to indicate that the existing end timestamp should be removed.
	EndingBefore          *string                                                  `form:"ending_before" json:"ending_before,omitempty"`
	Fraction              *Number                                                  `form:"fraction" json:"fraction,omitempty"`
	GcpOptions            *V1ContractsAmendContractResellerRoyaltyGcpOptionsParams `form:"gcp_options" json:"gcp_options,omitempty"`
	NetsuiteResellerID    *string                                                  `form:"netsuite_reseller_id" json:"netsuite_reseller_id,omitempty"`
	ResellerContractValue *Number                                                  `form:"reseller_contract_value" json:"reseller_contract_value,omitempty"`
	ResellerType          *string                                                  `form:"reseller_type" json:"reseller_type"`
	StartingAt            *string                                                  `form:"starting_at" json:"starting_at,omitempty"`
}

// Either provide amount or provide both unit_price and quantity.
type V1ContractsAmendContractScheduledChargeScheduleScheduleItemParams struct {
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
type V1ContractsAmendContractScheduledChargeScheduleRecurringScheduleParams struct {
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
type V1ContractsAmendContractScheduledChargeScheduleParams struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is only applicable to commit invoice schedules. If true, this schedule will not generate an invoice.
	DoNotInvoice *bool `form:"do_not_invoice" json:"do_not_invoice,omitempty"`
	// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
	RecurringSchedule *V1ContractsAmendContractScheduledChargeScheduleRecurringScheduleParams `form:"recurring_schedule" json:"recurring_schedule,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []*V1ContractsAmendContractScheduledChargeScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items,omitempty"`
}
type V1ContractsAmendContractScheduledChargeParams struct {
	CustomFields map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string                                                `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	ProductID            *string                                                `form:"product_id" json:"product_id"`
	Schedule             *V1ContractsAmendContractScheduledChargeScheduleParams `form:"schedule" json:"schedule"`
}

// Amendments will be replaced by Contract editing. New clients should implement using the `editContract` endpoint. Read more about the migration to contract editing [here](https://docs.stripe.com/guides/implement-metronome/migrate-amendments-to-edits/) and contact us via the [Metronome support portal](https://support.metronome.com/) for more details. Once contract editing is enabled, access to this endpoint will be removed.
type V1ContractsAmendContractParams struct {
	Params  `form:"*"`
	Commits []*V1ContractsAmendContractCommitParams `form:"commits" json:"commits,omitempty"`
	// ID of the contract to amend
	ContractID *string                                 `form:"contract_id" json:"contract_id"`
	Credits    []*V1ContractsAmendContractCreditParams `form:"credits" json:"credits,omitempty"`
	// ID of the customer whose contract is to be amended
	CustomerID   *string           `form:"customer_id" json:"customer_id"`
	CustomFields map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// This field's availability is dependent on your client's configuration.
	Discounts []*V1ContractsAmendContractDiscountParams `form:"discounts" json:"discounts,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string                                   `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	Overrides            []*V1ContractsAmendContractOverrideParams `form:"overrides" json:"overrides,omitempty"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []*V1ContractsAmendContractProfessionalServiceParams `form:"professional_services" json:"professional_services,omitempty"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties []*V1ContractsAmendContractResellerRoyaltyParams `form:"reseller_royalties" json:"reseller_royalties,omitempty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID *string                                          `form:"salesforce_opportunity_id" json:"salesforce_opportunity_id,omitempty"`
	ScheduledCharges        []*V1ContractsAmendContractScheduledChargeParams `form:"scheduled_charges" json:"scheduled_charges,omitempty"`
	// inclusive start time for the amendment
	StartingAt *string `form:"starting_at" json:"starting_at"`
	// This field's availability is dependent on your client's configuration.
	TotalContractValue *Number `form:"total_contract_value" json:"total_contract_value,omitempty"`
}

// Permanently end and archive a contract along with all its terms. Any draft invoices will be canceled, and all upcoming scheduled invoices will be voided–also all finalized invoices can optionally be voided. Use this in the event a contract was incorrectly created and needed to be removed from a customer.
//
// #### Impact on commits and credits:
// When archiving a contract, all associated commits and credits are also archived. For prepaid commits with active segments, Metronome automatically generates expiration ledger entries to close out any remaining balances, ensuring accurate accounting of unused prepaid amounts. These ledger entries will appear in the commit's transaction history with type `PREPAID_COMMIT_EXPIRATION`.
//
// #### Archived contract visibility:
// Archived contracts remain accessible for historical reporting and audit purposes. They can be retrieved using the `ListContracts` endpoint by setting the `include_archived` parameter to `true` or in the Metronome UI when the "Show archived" option is enabled.
type V1ContractsArchiveContractParams struct {
	Params `form:"*"`
	// ID of the contract to archive
	ContractID *string `form:"contract_id" json:"contract_id"`
	// ID of the customer whose contract is to be archived
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// If false, the existing finalized invoices will remain after the contract is archived.
	VoidInvoices *bool `form:"void_invoices" json:"void_invoices"`
}

// Per-subscription overrides for the subscriptions defined on the package template. Each entry targets one template subscription and sets its starting quantity or seats.
type V1ContractsCreateContractPackageCustomizationsModifySubscriptionParams struct {
	// The initial quantity for the subscription. Must be a non-negative value. Provide for a QUANTITY_ONLY subscription.
	InitialQuantity *Number `form:"initial_quantity" json:"initial_quantity,omitempty"`
	// Seat identifiers to assign at the start of a SEAT_BASED subscription.
	InitialSeatIDs []*string `form:"initial_seat_ids" json:"initial_seat_ids,omitempty"`
	// Number of unassigned seats to provision at the start of a SEAT_BASED subscription.
	InitialUnassignedSeats *Number `form:"initial_unassigned_seats" json:"initial_unassigned_seats,omitempty"`
	// Human-readable alias of the template subscription on the package to customize, as an alternative to template_subscription_id. Provide exactly one of the two.
	SubscriptionAlias *string `form:"subscription_alias" json:"subscription_alias,omitempty"`
	// ID of the template subscription on the package to customize. Provide exactly one of template_subscription_id or subscription_alias.
	TemplateSubscriptionID *string `form:"template_subscription_id" json:"template_subscription_id,omitempty"`
}

// Only set for TIERED rate_type.
type V1ContractsCreateContractPackageCustomizationsAddOverrideOverwriteRateTierParams struct {
	Price *Number `form:"price" json:"price"`
	Size  *Number `form:"size" json:"size,omitempty"`
}
type V1ContractsCreateContractPackageCustomizationsAddOverrideOverwriteRateMinimumConfigParams struct {
	Minimum *Number `form:"minimum" json:"minimum"`
}

// Required for OVERWRITE type.
type V1ContractsCreateContractPackageCustomizationsAddOverrideOverwriteRateParams struct {
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate processors.
	CustomRate map[string]JSON `form:"custom_rate" json:"custom_rate,omitempty"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be set to true.
	IsProrated    *bool                                                                                      `form:"is_prorated" json:"is_prorated,omitempty"`
	MinimumConfig *V1ContractsCreateContractPackageCustomizationsAddOverrideOverwriteRateMinimumConfigParams `form:"minimum_config" json:"minimum_config,omitempty"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price *Number `form:"price" json:"price,omitempty"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity *Number `form:"quantity" json:"quantity,omitempty"`
	RateType *string `form:"rate_type" json:"rate_type"`
	// Only set for TIERED rate_type.
	Tiers []*V1ContractsCreateContractPackageCustomizationsAddOverrideOverwriteRateTierParams `form:"tiers" json:"tiers,omitempty"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreateContractPackageCustomizationsAddOverrideOverrideSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Cannot be used in conjunction with product_id or applicable_product_tags. If provided, the override will apply to all products with the specified specifiers.
type V1ContractsCreateContractPackageCustomizationsAddOverrideOverrideSpecifierParams struct {
	// Can only be used for commit specific overrides. Must be used in conjunction with one of `product_id`, `product_tags`, `pricing_group_values`, or `presentation_group_values`. Must be used instead of both `commit_ids` and `recurring_commit_ids` If provided, the override will apply to any specified commit, credit, recurring commit or recurring credit IDs.
	AnyCommitOrCreditIDs []*string `form:"any_commit_or_credit_ids" json:"any_commit_or_credit_ids,omitempty"`
	BillingFrequency     *string   `form:"billing_frequency" json:"billing_frequency,omitempty"`
	// Can only be used for commit specific overrides. Must be used in conjunction with one of `product_id`, `product_tags`, `pricing_group_values`, or `presentation_group_values`. If provided, the override will only apply to the specified commits. If not provided, the override will apply to all commits.
	CommitIDs []*string `form:"commit_ids" json:"commit_ids,omitempty"`
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreateContractPackageCustomizationsAddOverrideOverrideSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// A map of group names to values. The override will only apply to line items with the specified presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// A map of pricing group names to values. The override will only apply to products with the specified pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the override will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the override will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
	// Can only be used for commit specific overrides. Must be used in conjunction with one of `product_id`, `product_tags`, `pricing_group_values`, or `presentation_group_values`. If provided, the override will only apply to commits created by the specified recurring commit ids.
	RecurringCommitIDs []*string `form:"recurring_commit_ids" json:"recurring_commit_ids,omitempty"`
}

// Required for TIERED type. Must have at least one tier.
type V1ContractsCreateContractPackageCustomizationsAddOverrideTierParams struct {
	Multiplier *Number `form:"multiplier" json:"multiplier"`
	Size       *Number `form:"size" json:"size,omitempty"`
}

// Rate overrides layered onto the contract on top of the package's baseline configuration. Same shape as create_contract's overrides.
type V1ContractsCreateContractPackageCustomizationsAddOverrideParams struct {
	// tags identifying products whose rates are being overridden. Cannot be used in conjunction with override_specifiers.
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	Entitled     *bool   `form:"entitled" json:"entitled,omitempty"`
	// Indicates whether the override should only apply to commits. Defaults to `false`. If `true` you can specify relevant commits in `override_specifiers` by passing `commit_ids`, `recurring_commit_ids`, or `any_commit_or_credit_ids`.  If you do not specify any of these fields, the override will apply when consuming any prepaid commit, postpaid commit, or credit
	IsCommitSpecific *bool `form:"is_commit_specific" json:"is_commit_specific,omitempty"`
	// Required for MULTIPLIER type. Must be >=0.
	Multiplier *Number `form:"multiplier" json:"multiplier,omitempty"`
	// Cannot be used in conjunction with product_id or applicable_product_tags. If provided, the override will apply to all products with the specified specifiers.
	OverrideSpecifiers []*V1ContractsCreateContractPackageCustomizationsAddOverrideOverrideSpecifierParams `form:"override_specifiers" json:"override_specifiers,omitempty"`
	// Required for OVERWRITE type.
	OverwriteRate *V1ContractsCreateContractPackageCustomizationsAddOverrideOverwriteRateParams `form:"overwrite_rate" json:"overwrite_rate,omitempty"`
	// Required for EXPLICIT multiplier prioritization scheme and all TIERED overrides. Under EXPLICIT prioritization, overwrites are prioritized first, and then tiered and multiplier overrides are prioritized by their priority value (lowest first). Must be > 0.
	Priority *Number `form:"priority" json:"priority,omitempty"`
	// ID of the product whose rate is being overridden. Cannot be used in conjunction with override_specifiers.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// RFC 3339 timestamp indicating when the override will start applying (inclusive)
	StartingAt *string `form:"starting_at" json:"starting_at"`
	// Indicates whether the override applies to commit rates or list rates. Can only be used for overrides that have `is_commit_specific` set to `true`. Defaults to `"LIST_RATE"`.
	Target *string `form:"target" json:"target,omitempty"`
	// Required for TIERED type. Must have at least one tier.
	Tiers []*V1ContractsCreateContractPackageCustomizationsAddOverrideTierParams `form:"tiers" json:"tiers,omitempty"`
	// Overwrites are prioritized over multipliers and tiered overrides.
	Type *string `form:"type" json:"type,omitempty"`
}
type V1ContractsCreateContractPackageCustomizationsAddCommitAccessScheduleScheduleItemParams struct {
	Amount *Number `form:"amount" json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	// RFC 3339 timestamp (inclusive)
	StartingAt *string `form:"starting_at" json:"starting_at"`
}

// Required: Schedule for distributing the commit to the customer. For "POSTPAID" commits only one schedule item is allowed and amount must match invoice_schedule total.
type V1ContractsCreateContractPackageCustomizationsAddCommitAccessScheduleParams struct {
	// Determines how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used. Defaults to `SPEND` if omitted.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID  *string                                                                                    `form:"credit_type_id" json:"credit_type_id,omitempty"`
	ScheduleItems []*V1ContractsCreateContractPackageCustomizationsAddCommitAccessScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items"`
}

// Either provide amount or provide both unit_price and quantity.
type V1ContractsCreateContractPackageCustomizationsAddCommitInvoiceScheduleScheduleItemParams struct {
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
type V1ContractsCreateContractPackageCustomizationsAddCommitInvoiceScheduleRecurringScheduleParams struct {
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

// Required for "POSTPAID" commits: the true up invoice will be generated at this time and only one schedule item is allowed; the total must match access_schedule amount. Optional for "PREPAID" commits: if not provided, this will be a "complimentary" commit with no invoice.
type V1ContractsCreateContractPackageCustomizationsAddCommitInvoiceScheduleParams struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is only applicable to commit invoice schedules. If true, this schedule will not generate an invoice.
	DoNotInvoice *bool `form:"do_not_invoice" json:"do_not_invoice,omitempty"`
	// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
	RecurringSchedule *V1ContractsCreateContractPackageCustomizationsAddCommitInvoiceScheduleRecurringScheduleParams `form:"recurring_schedule" json:"recurring_schedule,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []*V1ContractsCreateContractPackageCustomizationsAddCommitInvoiceScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items,omitempty"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreateContractPackageCustomizationsAddCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreateContractPackageCustomizationsAddCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreateContractPackageCustomizationsAddCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Optional configuration for commit hierarchy access control
type V1ContractsCreateContractPackageCustomizationsAddCommitHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}

// Optional attributes for spend tracker integration. Immutable after creation.
type V1ContractsCreateContractPackageCustomizationsAddCommitSpendTrackerAttributesParams struct {
	// If true, this commit will be included in spend trackers with discounted set to DISCOUNTED_ONLY
	CountsAsDiscounted *bool `form:"counts_as_discounted" json:"counts_as_discounted"`
}

// Commits added to the contract on top of the package's baseline configuration. Same shape as create_contract's commits.
type V1ContractsCreateContractPackageCustomizationsAddCommitParams struct {
	// Required: Schedule for distributing the commit to the customer. For "POSTPAID" commits only one schedule item is allowed and amount must match invoice_schedule total.
	AccessSchedule *V1ContractsCreateContractPackageCustomizationsAddCommitAccessScheduleParams `form:"access_schedule" json:"access_schedule,omitempty"`
	// (DEPRECATED) Use access_schedule and invoice_schedule instead.
	Amount *Number `form:"amount" json:"amount,omitempty"`
	// Which products the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductTags []*string         `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	CustomFields          map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// Used only in UI/API. It is not exposed to end customers.
	Description *string `form:"description" json:"description,omitempty"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration *V1ContractsCreateContractPackageCustomizationsAddCommitHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// Required for "POSTPAID" commits: the true up invoice will be generated at this time and only one schedule item is allowed; the total must match access_schedule amount. Optional for "PREPAID" commits: if not provided, this will be a "complimentary" commit with no invoice.
	InvoiceSchedule *V1ContractsCreateContractPackageCustomizationsAddCommitInvoiceScheduleParams `form:"invoice_schedule" json:"invoice_schedule,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// If multiple commits are applicable, the one with the lower priority will apply first.
	Priority  *Number `form:"priority" json:"priority,omitempty"`
	ProductID *string `form:"product_id" json:"product_id"`
	RateType  *string `form:"rate_type" json:"rate_type,omitempty"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreateContractPackageCustomizationsAddCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
	// Optional attributes for spend tracker integration. Immutable after creation.
	SpendTrackerAttributes *V1ContractsCreateContractPackageCustomizationsAddCommitSpendTrackerAttributesParams `form:"spend_tracker_attributes" json:"spend_tracker_attributes,omitempty"`
	// A temporary ID for the commit that can be used to reference the commit for commit specific overrides.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
	Type        *string `form:"type" json:"type"`
}
type V1ContractsCreateContractPackageCustomizationsAddCreditAccessScheduleScheduleItemParams struct {
	Amount *Number `form:"amount" json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	// RFC 3339 timestamp (inclusive)
	StartingAt *string `form:"starting_at" json:"starting_at"`
}

// Schedule for distributing the credit to the customer.
type V1ContractsCreateContractPackageCustomizationsAddCreditAccessScheduleParams struct {
	// Determines how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used. Defaults to `SPEND` if omitted.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID  *string                                                                                    `form:"credit_type_id" json:"credit_type_id,omitempty"`
	ScheduleItems []*V1ContractsCreateContractPackageCustomizationsAddCreditAccessScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreateContractPackageCustomizationsAddCreditSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreateContractPackageCustomizationsAddCreditSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreateContractPackageCustomizationsAddCreditSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Optional configuration for credit hierarchy access control
type V1ContractsCreateContractPackageCustomizationsAddCreditHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}

// Credits added to the contract on top of the package's baseline configuration. Same shape as create_contract's credits.
type V1ContractsCreateContractPackageCustomizationsAddCreditParams struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule *V1ContractsCreateContractPackageCustomizationsAddCreditAccessScheduleParams `form:"access_schedule" json:"access_schedule"`
	// Which products the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductTags []*string         `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	CustomFields          map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// Used only in UI/API. It is not exposed to end customers.
	Description *string `form:"description" json:"description,omitempty"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration *V1ContractsCreateContractPackageCustomizationsAddCreditHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// If multiple credits are applicable, the one with the lower priority will apply first.
	Priority  *Number `form:"priority" json:"priority,omitempty"`
	ProductID *string `form:"product_id" json:"product_id"`
	RateType  *string `form:"rate_type" json:"rate_type,omitempty"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreateContractPackageCustomizationsAddCreditSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
}

// The amount of commit to grant.
type V1ContractsCreateContractPackageCustomizationsAddRecurringCommitAccessAmountParams struct {
	// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is required unless a subscription is attached via `subscription_config`.
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	UnitPrice *Number `form:"unit_price" json:"unit_price"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreateContractPackageCustomizationsAddRecurringCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreateContractPackageCustomizationsAddRecurringCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreateContractPackageCustomizationsAddRecurringCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.
type V1ContractsCreateContractPackageCustomizationsAddRecurringCommitCommitDurationParams struct {
	Unit  *string `form:"unit" json:"unit,omitempty"`
	Value *Number `form:"value" json:"value"`
}
type V1ContractsCreateContractPackageCustomizationsAddRecurringCommitSubscriptionConfigApplySeatIncreaseConfigParams struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated *bool `form:"is_prorated" json:"is_prorated"`
}

// Attach a subscription to the recurring commit/credit.
type V1ContractsCreateContractPackageCustomizationsAddRecurringCommitSubscriptionConfigParams struct {
	// Controls when a customer can access recurring commit or credit child balances when the attached subscription is payment-gated. `BILLING_PERIOD_PAID` releases each balance after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases all balances after the first payment.
	AccessPolicy            *string                                                                                                          `form:"access_policy" json:"access_policy,omitempty"`
	Allocation              *string                                                                                                          `form:"allocation" json:"allocation,omitempty"`
	ApplySeatIncreaseConfig *V1ContractsCreateContractPackageCustomizationsAddRecurringCommitSubscriptionConfigApplySeatIncreaseConfigParams `form:"apply_seat_increase_config" json:"apply_seat_increase_config"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID *string `form:"subscription_id" json:"subscription_id"`
}

// Optional configuration for recurring commit/credit hierarchy access control
type V1ContractsCreateContractPackageCustomizationsAddRecurringCommitHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}

// The amount the customer should be billed for the commit. Not required.
type V1ContractsCreateContractPackageCustomizationsAddRecurringCommitInvoiceAmountParams struct {
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id"`
	Quantity     *Number `form:"quantity" json:"quantity"`
	UnitPrice    *Number `form:"unit_price" json:"unit_price"`
}
type V1ContractsCreateContractPackageCustomizationsAddRecurringCommitProrationRoundingAccessParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}
type V1ContractsCreateContractPackageCustomizationsAddRecurringCommitProrationRoundingInvoiceParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}

// Optional rounding configuration for prorated recurring commit amounts.
type V1ContractsCreateContractPackageCustomizationsAddRecurringCommitProrationRoundingParams struct {
	Access  *V1ContractsCreateContractPackageCustomizationsAddRecurringCommitProrationRoundingAccessParams  `form:"access" json:"access,omitempty"`
	Invoice *V1ContractsCreateContractPackageCustomizationsAddRecurringCommitProrationRoundingInvoiceParams `form:"invoice" json:"invoice,omitempty"`
}

// Recurring commits added to the contract on top of the package's baseline configuration. Same shape as create_contract's recurring_commits.
type V1ContractsCreateContractPackageCustomizationsAddRecurringCommitParams struct {
	// The amount of commit to grant.
	AccessAmount *V1ContractsCreateContractPackageCustomizationsAddRecurringCommitAccessAmountParams `form:"access_amount" json:"access_amount"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Will be passed down to the individual commits
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	// Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.
	CommitDuration *V1ContractsCreateContractPackageCustomizationsAddRecurringCommitCommitDurationParams `form:"commit_duration" json:"commit_duration"`
	// Will be passed down to the individual commits
	Description *string `form:"description" json:"description,omitempty"`
	// Determines when the contract will stop creating recurring commits. optional
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// Optional configuration for recurring commit/credit hierarchy access control
	HierarchyConfiguration *V1ContractsCreateContractPackageCustomizationsAddRecurringCommitHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount *V1ContractsCreateContractPackageCustomizationsAddRecurringCommitInvoiceAmountParams `form:"invoice_amount" json:"invoice_amount,omitempty"`
	// displayed on invoices. will be passed through to the individual commits
	Name *string `form:"name" json:"name,omitempty"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// Will be passed down to the individual commits
	Priority  *Number `form:"priority" json:"priority"`
	ProductID *string `form:"product_id" json:"product_id"`
	// Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration *string `form:"proration" json:"proration,omitempty"`
	// Optional rounding configuration for prorated recurring commit amounts.
	ProrationRounding *V1ContractsCreateContractPackageCustomizationsAddRecurringCommitProrationRoundingParams `form:"proration_rounding" json:"proration_rounding,omitempty"`
	// Whether the created commits will use the commit rate or list rate
	RateType *string `form:"rate_type" json:"rate_type,omitempty"`
	// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
	RecurrenceFrequency *string `form:"recurrence_frequency" json:"recurrence_frequency,omitempty"`
	// Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreateContractPackageCustomizationsAddRecurringCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
	// determines the start time for the first commit
	StartingAt *string `form:"starting_at" json:"starting_at"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig *V1ContractsCreateContractPackageCustomizationsAddRecurringCommitSubscriptionConfigParams `form:"subscription_config" json:"subscription_config,omitempty"`
	// A temporary ID that can be used to reference the recurring commit for commit specific overrides.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
}

// The amount of commit to grant.
type V1ContractsCreateContractPackageCustomizationsAddRecurringCreditAccessAmountParams struct {
	// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is required unless a subscription is attached via `subscription_config`.
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	UnitPrice *Number `form:"unit_price" json:"unit_price"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreateContractPackageCustomizationsAddRecurringCreditSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreateContractPackageCustomizationsAddRecurringCreditSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreateContractPackageCustomizationsAddRecurringCreditSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.
type V1ContractsCreateContractPackageCustomizationsAddRecurringCreditCommitDurationParams struct {
	Unit  *string `form:"unit" json:"unit,omitempty"`
	Value *Number `form:"value" json:"value"`
}
type V1ContractsCreateContractPackageCustomizationsAddRecurringCreditSubscriptionConfigApplySeatIncreaseConfigParams struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated *bool `form:"is_prorated" json:"is_prorated"`
}

// Attach a subscription to the recurring commit/credit.
type V1ContractsCreateContractPackageCustomizationsAddRecurringCreditSubscriptionConfigParams struct {
	// Controls when a customer can access recurring commit or credit child balances when the attached subscription is payment-gated. `BILLING_PERIOD_PAID` releases each balance after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases all balances after the first payment.
	AccessPolicy            *string                                                                                                          `form:"access_policy" json:"access_policy,omitempty"`
	Allocation              *string                                                                                                          `form:"allocation" json:"allocation,omitempty"`
	ApplySeatIncreaseConfig *V1ContractsCreateContractPackageCustomizationsAddRecurringCreditSubscriptionConfigApplySeatIncreaseConfigParams `form:"apply_seat_increase_config" json:"apply_seat_increase_config"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID *string `form:"subscription_id" json:"subscription_id"`
}

// Optional configuration for recurring commit/credit hierarchy access control
type V1ContractsCreateContractPackageCustomizationsAddRecurringCreditHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}
type V1ContractsCreateContractPackageCustomizationsAddRecurringCreditProrationRoundingAccessParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}

// Optional rounding configuration for prorated recurring credit amounts.
type V1ContractsCreateContractPackageCustomizationsAddRecurringCreditProrationRoundingParams struct {
	Access *V1ContractsCreateContractPackageCustomizationsAddRecurringCreditProrationRoundingAccessParams `form:"access" json:"access,omitempty"`
}

// Recurring credits added to the contract on top of the package's baseline configuration. Same shape as create_contract's recurring_credits.
type V1ContractsCreateContractPackageCustomizationsAddRecurringCreditParams struct {
	// The amount of commit to grant.
	AccessAmount *V1ContractsCreateContractPackageCustomizationsAddRecurringCreditAccessAmountParams `form:"access_amount" json:"access_amount"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Will be passed down to the individual commits
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	// Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.
	CommitDuration *V1ContractsCreateContractPackageCustomizationsAddRecurringCreditCommitDurationParams `form:"commit_duration" json:"commit_duration"`
	// Will be passed down to the individual commits
	Description *string `form:"description" json:"description,omitempty"`
	// Determines when the contract will stop creating recurring commits. optional
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// Optional configuration for recurring commit/credit hierarchy access control
	HierarchyConfiguration *V1ContractsCreateContractPackageCustomizationsAddRecurringCreditHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// displayed on invoices. will be passed through to the individual commits
	Name *string `form:"name" json:"name,omitempty"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// Will be passed down to the individual commits
	Priority  *Number `form:"priority" json:"priority"`
	ProductID *string `form:"product_id" json:"product_id"`
	// Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration *string `form:"proration" json:"proration,omitempty"`
	// Optional rounding configuration for prorated recurring credit amounts.
	ProrationRounding *V1ContractsCreateContractPackageCustomizationsAddRecurringCreditProrationRoundingParams `form:"proration_rounding" json:"proration_rounding,omitempty"`
	// Whether the created commits will use the commit rate or list rate
	RateType *string `form:"rate_type" json:"rate_type,omitempty"`
	// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
	RecurrenceFrequency *string `form:"recurrence_frequency" json:"recurrence_frequency,omitempty"`
	// Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreateContractPackageCustomizationsAddRecurringCreditSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
	// determines the start time for the first commit
	StartingAt *string `form:"starting_at" json:"starting_at"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig *V1ContractsCreateContractPackageCustomizationsAddRecurringCreditSubscriptionConfigParams `form:"subscription_config" json:"subscription_config,omitempty"`
	// A temporary ID that can be used to reference the recurring commit for commit specific overrides.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
}

// Either provide amount or provide both unit_price and quantity.
type V1ContractsCreateContractPackageCustomizationsAddScheduledChargeScheduleScheduleItemParams struct {
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
type V1ContractsCreateContractPackageCustomizationsAddScheduledChargeScheduleRecurringScheduleParams struct {
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
type V1ContractsCreateContractPackageCustomizationsAddScheduledChargeScheduleParams struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is only applicable to commit invoice schedules. If true, this schedule will not generate an invoice.
	DoNotInvoice *bool `form:"do_not_invoice" json:"do_not_invoice,omitempty"`
	// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
	RecurringSchedule *V1ContractsCreateContractPackageCustomizationsAddScheduledChargeScheduleRecurringScheduleParams `form:"recurring_schedule" json:"recurring_schedule,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []*V1ContractsCreateContractPackageCustomizationsAddScheduledChargeScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items,omitempty"`
}

// Scheduled charges added to the contract on top of the package's baseline configuration. Same shape as create_contract's scheduled_charges.
type V1ContractsCreateContractPackageCustomizationsAddScheduledChargeParams struct {
	CustomFields map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string                                                                         `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	ProductID            *string                                                                         `form:"product_id" json:"product_id"`
	Schedule             *V1ContractsCreateContractPackageCustomizationsAddScheduledChargeScheduleParams `form:"schedule" json:"schedule"`
}
type V1ContractsCreateContractPackageCustomizationsAddSpendTrackerApplicableSpendSpecifierParams struct {
	// Filter by whether the spend was discounted. Defaults to ANY if omitted.
	Discounted *string   `form:"discounted" json:"discounted,omitempty"`
	Sources    []*string `form:"sources" json:"sources"`
	SpendType  *string   `form:"spend_type" json:"spend_type"`
}

// Spend trackers added to the contract on top of the package's baseline configuration. Same shape as create_contract's spend_trackers; aliases must be unique within a contract.
type V1ContractsCreateContractPackageCustomizationsAddSpendTrackerParams struct {
	// Human-readable identifier, unique per contract.
	Alias                     *string                                                                                        `form:"alias" json:"alias"`
	ApplicableSpendSpecifiers []*V1ContractsCreateContractPackageCustomizationsAddSpendTrackerApplicableSpendSpecifierParams `form:"applicable_spend_specifiers" json:"applicable_spend_specifiers"`
	CreditTypeID              *string                                                                                        `form:"credit_type_id" json:"credit_type_id"`
	ResetFrequency            *string                                                                                        `form:"reset_frequency" json:"reset_frequency"`
}
type V1ContractsCreateContractPackageCustomizationsAddSubscriptionSubscriptionRateParams struct {
	// Frequency to bill subscription with. Together with product_id, must match existing rate on the rate card.
	BillingFrequency *string `form:"billing_frequency" json:"billing_frequency"`
	// Must be subscription type product
	ProductID *string `form:"product_id" json:"product_id"`
}
type V1ContractsCreateContractPackageCustomizationsAddSubscriptionProrationRoundingParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}
type V1ContractsCreateContractPackageCustomizationsAddSubscriptionProrationParams struct {
	// Indicates how mid-period quantity adjustments are invoiced.
	// **BILL_IMMEDIATELY**: Only available when collection schedule is `ADVANCE`. The quantity increase will be billed immediately on the scheduled date.
	// **BILL_ON_NEXT_COLLECTION_DATE**: The quantity increase will be billed for in-arrears at the end of the period.
	//
	InvoiceBehavior *string `form:"invoice_behavior" json:"invoice_behavior,omitempty"`
	// Indicates if the partial period will be prorated or charged a full amount.
	IsProrated *bool                                                                                 `form:"is_prorated" json:"is_prorated,omitempty"`
	Rounding   *V1ContractsCreateContractPackageCustomizationsAddSubscriptionProrationRoundingParams `form:"rounding" json:"rounding,omitempty"`
}
type V1ContractsCreateContractPackageCustomizationsAddSubscriptionSeatConfigParams struct {
	// The initial assigned seats on this subscription.
	InitialSeatIDs []*string `form:"initial_seat_ids" json:"initial_seat_ids"`
	// The initial amount of unassigned seats on this subscription.
	InitialUnassignedSeats *Number `form:"initial_unassigned_seats" json:"initial_unassigned_seats,omitempty"`
	// The property name, sent on usage events, that identifies the seat ID associated with the usage event.  For example, the property name might be seat_id or user_id. The property must be set as a group key on billable metrics and a presentation/pricing group key on contract products.  This allows linked recurring credits with an allocation per seat to be consumed by only one seat's usage.
	SeatGroupKey *string `form:"seat_group_key" json:"seat_group_key"`
}
type V1ContractsCreateContractPackageCustomizationsAddSubscriptionBillingCycleConfigParams struct {
	// The date to anchor the billing cycle to. If omitted, defaults to the contract's usage invoice billing cycle anchor date.
	AnchorDate *string `form:"anchor_date" json:"anchor_date,omitempty"`
	// Controls whether this subscription consolidates onto usage invoices or gets its own scheduled invoice. Defaults to ON_USAGE_INVOICE if omitted.
	InvoicePlacement *string `form:"invoice_placement" json:"invoice_placement,omitempty"`
}
type V1ContractsCreateContractPackageCustomizationsAddSubscriptionPaymentGateConfigParams struct {
	// Gate access to the subscription based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe.
	PaymentGateType *string `form:"payment_gate_type" json:"payment_gate_type"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType *string `form:"tax_type" json:"tax_type,omitempty"`
}

// New subscriptions added to the contract on top of the package's baseline configuration. Distinct from modify_subscriptions, which adjusts subscriptions defined on the package template. Same shape as create_contract's subscriptions.
type V1ContractsCreateContractPackageCustomizationsAddSubscriptionParams struct {
	BillingCycleConfig *V1ContractsCreateContractPackageCustomizationsAddSubscriptionBillingCycleConfigParams `form:"billing_cycle_config" json:"billing_cycle_config,omitempty"`
	CollectionSchedule *string                                                                                `form:"collection_schedule" json:"collection_schedule"`
	CustomFields       map[string]string                                                                      `form:"custom_fields" json:"custom_fields,omitempty"`
	Description        *string                                                                                `form:"description" json:"description,omitempty"`
	// Exclusive end time for the subscription. If not provided, subscription inherits contract end date.
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// The initial quantity for the subscription. It must be non-negative value. Required if quantity_management_mode is QUANTITY_ONLY.
	InitialQuantity   *Number                                                                               `form:"initial_quantity" json:"initial_quantity,omitempty"`
	Name              *string                                                                               `form:"name" json:"name,omitempty"`
	PaymentGateConfig *V1ContractsCreateContractPackageCustomizationsAddSubscriptionPaymentGateConfigParams `form:"payment_gate_config" json:"payment_gate_config,omitempty"`
	Proration         *V1ContractsCreateContractPackageCustomizationsAddSubscriptionProrationParams         `form:"proration" json:"proration"`
	// Determines how the subscription's quantity is controlled. Defaults to QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified directly on the subscription. `initial_quantity` must be provided with this option. Compatible with recurring commits/credits that use POOLED allocation. **SEAT_BASED**: Use when you want to pass specific seat identifiers (e.g. add user_123) to increment and decrement a subscription quantity, rather than directly providing the quantity. You must use a **SEAT_BASED** subscription to use a linked recurring credit with an allocation per seat. `seat_config` must be provided with this option.
	QuantityManagementMode *string                                                                        `form:"quantity_management_mode" json:"quantity_management_mode,omitempty"`
	SeatConfig             *V1ContractsCreateContractPackageCustomizationsAddSubscriptionSeatConfigParams `form:"seat_config" json:"seat_config,omitempty"`
	// Inclusive start time for the subscription. If not provided, defaults to contract start date
	StartingAt       *string                                                                              `form:"starting_at" json:"starting_at,omitempty"`
	SubscriptionRate *V1ContractsCreateContractPackageCustomizationsAddSubscriptionSubscriptionRateParams `form:"subscription_rate" json:"subscription_rate"`
	// A temporary ID used to reference the subscription in recurring commit/credit subscription configs created within the same payload.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
}
type V1ContractsCreateContractPackageCustomizationsSetSpendThresholdConfigurationCommitParams struct {
	Description *string `form:"description" json:"description,omitempty"`
	// Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.
	Name *string `form:"name" json:"name,omitempty"`
	// The priority of the commit, used to determine drawdown order. Lower priority commits are consumed first. Defaults to 100 if not specified.
	Priority *Number `form:"priority" json:"priority,omitempty"`
	// The commit product that will be used to generate the line item for commit payment.
	ProductID *string `form:"product_id" json:"product_id"`
}

// Only applicable if using STRIPE as your payment gate type.
type V1ContractsCreateContractPackageCustomizationsSetSpendThresholdConfigurationPaymentGateConfigStripeConfigParams struct {
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as your payment type.
	InvoiceMetadata map[string]string `form:"invoice_metadata" json:"invoice_metadata,omitempty"`
	// If left blank, will default to INVOICE
	PaymentType *string `form:"payment_type" json:"payment_type"`
}

// Only applicable if using PRECALCULATED as your tax type.
type V1ContractsCreateContractPackageCustomizationsSetSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams struct {
	// Amount of tax to be applied. This should be in the same currency and denomination  as the commit's invoice schedule
	TaxAmount *Number `form:"tax_amount" json:"tax_amount"`
	// Name of the tax to be applied. This may be used in an invoice line item description.
	TaxName *string `form:"tax_name" json:"tax_name,omitempty"`
}
type V1ContractsCreateContractPackageCustomizationsSetSpendThresholdConfigurationPaymentGateConfigParams struct {
	// Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.
	PaymentGateType *string `form:"payment_gate_type" json:"payment_gate_type"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig *V1ContractsCreateContractPackageCustomizationsSetSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams `form:"precalculated_tax_config" json:"precalculated_tax_config,omitempty"`
	// Only applicable if using STRIPE as your payment gate type.
	StripeConfig *V1ContractsCreateContractPackageCustomizationsSetSpendThresholdConfigurationPaymentGateConfigStripeConfigParams `form:"stripe_config" json:"stripe_config,omitempty"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType *string `form:"tax_type" json:"tax_type,omitempty"`
}

// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
type V1ContractsCreateContractPackageCustomizationsSetSpendThresholdConfigurationDiscountConfigurationCapParams struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount *Number `form:"amount" json:"amount"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias *string `form:"spend_tracker_alias" json:"spend_tracker_alias"`
}
type V1ContractsCreateContractPackageCustomizationsSetSpendThresholdConfigurationDiscountConfigurationParams struct {
	// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
	Cap *V1ContractsCreateContractPackageCustomizationsSetSpendThresholdConfigurationDiscountConfigurationCapParams `form:"cap" json:"cap,omitempty"`
	// The fraction of the original amount that the customer pays after applying the discount. For example, 0.85 means the customer pays 85% of the original amount (a 15% discount).
	PaymentFraction *Number `form:"payment_fraction" json:"payment_fraction"`
}

// Spend threshold configuration for the contract, using set (create-or-replace) semantics: supply the complete configuration and it replaces the package's baseline in full. Same shape as create_contract's spend_threshold_configuration.
type V1ContractsCreateContractPackageCustomizationsSetSpendThresholdConfigurationParams struct {
	Commit                *V1ContractsCreateContractPackageCustomizationsSetSpendThresholdConfigurationCommitParams                `form:"commit" json:"commit"`
	DiscountConfiguration *V1ContractsCreateContractPackageCustomizationsSetSpendThresholdConfigurationDiscountConfigurationParams `form:"discount_configuration" json:"discount_configuration,omitempty"`
	// When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.
	IsEnabled         *bool                                                                                                `form:"is_enabled" json:"is_enabled"`
	PaymentGateConfig *V1ContractsCreateContractPackageCustomizationsSetSpendThresholdConfigurationPaymentGateConfigParams `form:"payment_gate_config" json:"payment_gate_config"`
	// Specify the threshold amount for the contract. Each time the contract's usage hits this amount, a threshold charge will be initiated.
	ThresholdAmount *Number `form:"threshold_amount" json:"threshold_amount"`
}

// The length of time the created commit will be valid, starting from the end of the invoice's service period. If not provided, defaults to one year.
type V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationCommitDurationParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}
type V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationCommitParams struct {
	// Which products the threshold commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the threshold commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	Description           *string   `form:"description" json:"description,omitempty"`
	// The length of time the created commit will be valid, starting from the end of the invoice's service period. If not provided, defaults to one year.
	Duration *V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationCommitDurationParams `form:"duration" json:"duration,omitempty"`
	// Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.
	Name *string `form:"name" json:"name,omitempty"`
	// The priority of the commit, used to determine drawdown order. Lower priority commits are consumed first. Defaults to 100 if not specified.
	Priority *Number `form:"priority" json:"priority,omitempty"`
	// The commit product that will be used to generate the line item for commit payment.
	ProductID *string `form:"product_id" json:"product_id"`
	// Whether the created commits will be charged at commit rate or list rate.
	RateType *string `form:"rate_type" json:"rate_type,omitempty"`
	// Fraction of the created commit's unused balance that will roll over. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
}

// Only applicable if using STRIPE as your payment gate type.
type V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigParams struct {
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as your payment type.
	InvoiceMetadata map[string]string `form:"invoice_metadata" json:"invoice_metadata,omitempty"`
	// If left blank, will default to INVOICE
	PaymentType *string `form:"payment_type" json:"payment_type"`
}

// Only applicable if using PRECALCULATED as your tax type.
type V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams struct {
	// Amount of tax to be applied. This should be in the same currency and denomination  as the commit's invoice schedule
	TaxAmount *Number `form:"tax_amount" json:"tax_amount"`
	// Name of the tax to be applied. This may be used in an invoice line item description.
	TaxName *string `form:"tax_name" json:"tax_name,omitempty"`
}
type V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationPaymentGateConfigParams struct {
	// Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.
	PaymentGateType *string `form:"payment_gate_type" json:"payment_gate_type"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig *V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams `form:"precalculated_tax_config" json:"precalculated_tax_config,omitempty"`
	// Only applicable if using STRIPE as your payment gate type.
	StripeConfig *V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigParams `form:"stripe_config" json:"stripe_config,omitempty"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType *string `form:"tax_type" json:"tax_type,omitempty"`
}

// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
type V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationDiscountConfigurationCapParams struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount *Number `form:"amount" json:"amount"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias *string `form:"spend_tracker_alias" json:"spend_tracker_alias"`
}
type V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationDiscountConfigurationParams struct {
	// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
	Cap *V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationDiscountConfigurationCapParams `form:"cap" json:"cap,omitempty"`
	// The fraction of the original amount that the customer pays after applying the discount. For example, 0.85 means the customer pays 85% of the original amount (a 15% discount).
	PaymentFraction *Number `form:"payment_fraction" json:"payment_fraction"`
}

// If provided, balances with all the custom fields will not be considered when evaluating threshold billing
type V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilterParams struct {
	Entity *string `form:"entity" json:"entity"`
	Key    *string `form:"key" json:"key"`
	Value  *string `form:"value" json:"value"`
}

// If any of the exclude specifier is met, the balance is not considered when evaluating threshold billing
type V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeParams struct {
	// If provided, balances with all the custom fields will not be considered when evaluating threshold billing
	CustomFieldFilters []*V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilterParams `form:"custom_field_filters" json:"custom_field_filters"`
}

// Determines which balances are excluded from remaining balance calculation for threshold billing.
type V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierParams struct {
	// If any of the exclude specifier is met, the balance is not considered when evaluating threshold billing
	Exclude []*V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeParams `form:"exclude" json:"exclude"`
}

// Prepaid balance (credit balance recharge) threshold configuration for the contract, using set (create-or-replace) semantics: supply the complete configuration and it replaces the package's baseline in full. Same shape as create_contract's prepaid_balance_threshold_configuration.
type V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationParams struct {
	Commit *V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationCommitParams `form:"commit" json:"commit"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID    *string                                                                                                           `form:"custom_credit_type_id" json:"custom_credit_type_id,omitempty"`
	DiscountConfiguration *V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationDiscountConfigurationParams `form:"discount_configuration" json:"discount_configuration,omitempty"`
	// When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.
	IsEnabled         *bool                                                                                                         `form:"is_enabled" json:"is_enabled"`
	PaymentGateConfig *V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationPaymentGateConfigParams `form:"payment_gate_config" json:"payment_gate_config"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount *Number `form:"recharge_to_amount" json:"recharge_to_amount"`
	// Specify the threshold amount for the contract. Each time the contract's prepaid balance lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount *Number `form:"threshold_amount" json:"threshold_amount"`
	// Determines which balances are excluded from remaining balance calculation for threshold billing.
	ThresholdBalanceSpecifiers []*V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierParams `form:"threshold_balance_specifiers" json:"threshold_balance_specifiers,omitempty"`
}
type V1ContractsCreateContractPackageCustomizationsParams struct {
	// Commits added to the contract on top of the package's baseline configuration. Same shape as create_contract's commits.
	AddCommits []*V1ContractsCreateContractPackageCustomizationsAddCommitParams `form:"add_commits" json:"add_commits,omitempty"`
	// Credits added to the contract on top of the package's baseline configuration. Same shape as create_contract's credits.
	AddCredits []*V1ContractsCreateContractPackageCustomizationsAddCreditParams `form:"add_credits" json:"add_credits,omitempty"`
	// Rate overrides layered onto the contract on top of the package's baseline configuration. Same shape as create_contract's overrides.
	AddOverrides []*V1ContractsCreateContractPackageCustomizationsAddOverrideParams `form:"add_overrides" json:"add_overrides,omitempty"`
	// Recurring commits added to the contract on top of the package's baseline configuration. Same shape as create_contract's recurring_commits.
	AddRecurringCommits []*V1ContractsCreateContractPackageCustomizationsAddRecurringCommitParams `form:"add_recurring_commits" json:"add_recurring_commits,omitempty"`
	// Recurring credits added to the contract on top of the package's baseline configuration. Same shape as create_contract's recurring_credits.
	AddRecurringCredits []*V1ContractsCreateContractPackageCustomizationsAddRecurringCreditParams `form:"add_recurring_credits" json:"add_recurring_credits,omitempty"`
	// Scheduled charges added to the contract on top of the package's baseline configuration. Same shape as create_contract's scheduled_charges.
	AddScheduledCharges []*V1ContractsCreateContractPackageCustomizationsAddScheduledChargeParams `form:"add_scheduled_charges" json:"add_scheduled_charges,omitempty"`
	// Spend trackers added to the contract on top of the package's baseline configuration. Same shape as create_contract's spend_trackers; aliases must be unique within a contract.
	AddSpendTrackers []*V1ContractsCreateContractPackageCustomizationsAddSpendTrackerParams `form:"add_spend_trackers" json:"add_spend_trackers,omitempty"`
	// New subscriptions added to the contract on top of the package's baseline configuration. Distinct from modify_subscriptions, which adjusts subscriptions defined on the package template. Same shape as create_contract's subscriptions.
	AddSubscriptions []*V1ContractsCreateContractPackageCustomizationsAddSubscriptionParams `form:"add_subscriptions" json:"add_subscriptions,omitempty"`
	// Per-subscription overrides for the subscriptions defined on the package template. Each entry targets one template subscription and sets its starting quantity or seats.
	ModifySubscriptions []*V1ContractsCreateContractPackageCustomizationsModifySubscriptionParams `form:"modify_subscriptions" json:"modify_subscriptions,omitempty"`
	// Prepaid balance (credit balance recharge) threshold configuration for the contract, using set (create-or-replace) semantics: supply the complete configuration and it replaces the package's baseline in full. Same shape as create_contract's prepaid_balance_threshold_configuration.
	SetPrepaidBalanceThresholdConfiguration *V1ContractsCreateContractPackageCustomizationsSetPrepaidBalanceThresholdConfigurationParams `form:"set_prepaid_balance_threshold_configuration" json:"set_prepaid_balance_threshold_configuration,omitempty"`
	// Spend threshold configuration for the contract, using set (create-or-replace) semantics: supply the complete configuration and it replaces the package's baseline in full. Same shape as create_contract's spend_threshold_configuration.
	SetSpendThresholdConfiguration *V1ContractsCreateContractPackageCustomizationsSetSpendThresholdConfigurationParams `form:"set_spend_threshold_configuration" json:"set_spend_threshold_configuration,omitempty"`
}
type V1ContractsCreateContractCommitAccessScheduleScheduleItemParams struct {
	Amount *Number `form:"amount" json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	// RFC 3339 timestamp (inclusive)
	StartingAt *string `form:"starting_at" json:"starting_at"`
}

// Required: Schedule for distributing the commit to the customer. For "POSTPAID" commits only one schedule item is allowed and amount must match invoice_schedule total.
type V1ContractsCreateContractCommitAccessScheduleParams struct {
	// Determines how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used. Defaults to `SPEND` if omitted.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID  *string                                                            `form:"credit_type_id" json:"credit_type_id,omitempty"`
	ScheduleItems []*V1ContractsCreateContractCommitAccessScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items"`
}

// Either provide amount or provide both unit_price and quantity.
type V1ContractsCreateContractCommitInvoiceScheduleScheduleItemParams struct {
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
type V1ContractsCreateContractCommitInvoiceScheduleRecurringScheduleParams struct {
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

// Required for "POSTPAID" commits: the true up invoice will be generated at this time and only one schedule item is allowed; the total must match access_schedule amount. Optional for "PREPAID" commits: if not provided, this will be a "complimentary" commit with no invoice.
type V1ContractsCreateContractCommitInvoiceScheduleParams struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is only applicable to commit invoice schedules. If true, this schedule will not generate an invoice.
	DoNotInvoice *bool `form:"do_not_invoice" json:"do_not_invoice,omitempty"`
	// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
	RecurringSchedule *V1ContractsCreateContractCommitInvoiceScheduleRecurringScheduleParams `form:"recurring_schedule" json:"recurring_schedule,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []*V1ContractsCreateContractCommitInvoiceScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items,omitempty"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreateContractCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreateContractCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreateContractCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Optional configuration for commit hierarchy access control
type V1ContractsCreateContractCommitHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}

// Optional attributes for spend tracker integration. Immutable after creation.
type V1ContractsCreateContractCommitSpendTrackerAttributesParams struct {
	// If true, this commit will be included in spend trackers with discounted set to DISCOUNTED_ONLY
	CountsAsDiscounted *bool `form:"counts_as_discounted" json:"counts_as_discounted"`
}
type V1ContractsCreateContractCommitParams struct {
	// Required: Schedule for distributing the commit to the customer. For "POSTPAID" commits only one schedule item is allowed and amount must match invoice_schedule total.
	AccessSchedule *V1ContractsCreateContractCommitAccessScheduleParams `form:"access_schedule" json:"access_schedule,omitempty"`
	// (DEPRECATED) Use access_schedule and invoice_schedule instead.
	Amount *Number `form:"amount" json:"amount,omitempty"`
	// Which products the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductTags []*string         `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	CustomFields          map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// Used only in UI/API. It is not exposed to end customers.
	Description *string `form:"description" json:"description,omitempty"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration *V1ContractsCreateContractCommitHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// Required for "POSTPAID" commits: the true up invoice will be generated at this time and only one schedule item is allowed; the total must match access_schedule amount. Optional for "PREPAID" commits: if not provided, this will be a "complimentary" commit with no invoice.
	InvoiceSchedule *V1ContractsCreateContractCommitInvoiceScheduleParams `form:"invoice_schedule" json:"invoice_schedule,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// If multiple commits are applicable, the one with the lower priority will apply first.
	Priority  *Number `form:"priority" json:"priority,omitempty"`
	ProductID *string `form:"product_id" json:"product_id"`
	RateType  *string `form:"rate_type" json:"rate_type,omitempty"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreateContractCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
	// Optional attributes for spend tracker integration. Immutable after creation.
	SpendTrackerAttributes *V1ContractsCreateContractCommitSpendTrackerAttributesParams `form:"spend_tracker_attributes" json:"spend_tracker_attributes,omitempty"`
	// A temporary ID for the commit that can be used to reference the commit for commit specific overrides.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
	Type        *string `form:"type" json:"type"`
}
type V1ContractsCreateContractCreditAccessScheduleScheduleItemParams struct {
	Amount *Number `form:"amount" json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	// RFC 3339 timestamp (inclusive)
	StartingAt *string `form:"starting_at" json:"starting_at"`
}

// Schedule for distributing the credit to the customer.
type V1ContractsCreateContractCreditAccessScheduleParams struct {
	// Determines how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used. Defaults to `SPEND` if omitted.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID  *string                                                            `form:"credit_type_id" json:"credit_type_id,omitempty"`
	ScheduleItems []*V1ContractsCreateContractCreditAccessScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreateContractCreditSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreateContractCreditSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreateContractCreditSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Optional configuration for credit hierarchy access control
type V1ContractsCreateContractCreditHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}
type V1ContractsCreateContractCreditParams struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule *V1ContractsCreateContractCreditAccessScheduleParams `form:"access_schedule" json:"access_schedule"`
	// Which products the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductTags []*string         `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	CustomFields          map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// Used only in UI/API. It is not exposed to end customers.
	Description *string `form:"description" json:"description,omitempty"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration *V1ContractsCreateContractCreditHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// If multiple credits are applicable, the one with the lower priority will apply first.
	Priority  *Number `form:"priority" json:"priority,omitempty"`
	ProductID *string `form:"product_id" json:"product_id"`
	RateType  *string `form:"rate_type" json:"rate_type,omitempty"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreateContractCreditSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
}

// The amount of commit to grant.
type V1ContractsCreateContractRecurringCommitAccessAmountParams struct {
	// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is required unless a subscription is attached via `subscription_config`.
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	UnitPrice *Number `form:"unit_price" json:"unit_price"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreateContractRecurringCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreateContractRecurringCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreateContractRecurringCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.
type V1ContractsCreateContractRecurringCommitCommitDurationParams struct {
	Unit  *string `form:"unit" json:"unit,omitempty"`
	Value *Number `form:"value" json:"value"`
}
type V1ContractsCreateContractRecurringCommitSubscriptionConfigApplySeatIncreaseConfigParams struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated *bool `form:"is_prorated" json:"is_prorated"`
}

// Attach a subscription to the recurring commit/credit.
type V1ContractsCreateContractRecurringCommitSubscriptionConfigParams struct {
	// Controls when a customer can access recurring commit or credit child balances when the attached subscription is payment-gated. `BILLING_PERIOD_PAID` releases each balance after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases all balances after the first payment.
	AccessPolicy            *string                                                                                  `form:"access_policy" json:"access_policy,omitempty"`
	Allocation              *string                                                                                  `form:"allocation" json:"allocation,omitempty"`
	ApplySeatIncreaseConfig *V1ContractsCreateContractRecurringCommitSubscriptionConfigApplySeatIncreaseConfigParams `form:"apply_seat_increase_config" json:"apply_seat_increase_config"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID *string `form:"subscription_id" json:"subscription_id"`
}

// Optional configuration for recurring commit/credit hierarchy access control
type V1ContractsCreateContractRecurringCommitHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}

// The amount the customer should be billed for the commit. Not required.
type V1ContractsCreateContractRecurringCommitInvoiceAmountParams struct {
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id"`
	Quantity     *Number `form:"quantity" json:"quantity"`
	UnitPrice    *Number `form:"unit_price" json:"unit_price"`
}
type V1ContractsCreateContractRecurringCommitProrationRoundingAccessParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}
type V1ContractsCreateContractRecurringCommitProrationRoundingInvoiceParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}

// Optional rounding configuration for prorated recurring commit amounts.
type V1ContractsCreateContractRecurringCommitProrationRoundingParams struct {
	Access  *V1ContractsCreateContractRecurringCommitProrationRoundingAccessParams  `form:"access" json:"access,omitempty"`
	Invoice *V1ContractsCreateContractRecurringCommitProrationRoundingInvoiceParams `form:"invoice" json:"invoice,omitempty"`
}
type V1ContractsCreateContractRecurringCommitParams struct {
	// The amount of commit to grant.
	AccessAmount *V1ContractsCreateContractRecurringCommitAccessAmountParams `form:"access_amount" json:"access_amount"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Will be passed down to the individual commits
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	// Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.
	CommitDuration *V1ContractsCreateContractRecurringCommitCommitDurationParams `form:"commit_duration" json:"commit_duration"`
	// Will be passed down to the individual commits
	Description *string `form:"description" json:"description,omitempty"`
	// Determines when the contract will stop creating recurring commits. optional
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// Optional configuration for recurring commit/credit hierarchy access control
	HierarchyConfiguration *V1ContractsCreateContractRecurringCommitHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount *V1ContractsCreateContractRecurringCommitInvoiceAmountParams `form:"invoice_amount" json:"invoice_amount,omitempty"`
	// displayed on invoices. will be passed through to the individual commits
	Name *string `form:"name" json:"name,omitempty"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// Will be passed down to the individual commits
	Priority  *Number `form:"priority" json:"priority"`
	ProductID *string `form:"product_id" json:"product_id"`
	// Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration *string `form:"proration" json:"proration,omitempty"`
	// Optional rounding configuration for prorated recurring commit amounts.
	ProrationRounding *V1ContractsCreateContractRecurringCommitProrationRoundingParams `form:"proration_rounding" json:"proration_rounding,omitempty"`
	// Whether the created commits will use the commit rate or list rate
	RateType *string `form:"rate_type" json:"rate_type,omitempty"`
	// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
	RecurrenceFrequency *string `form:"recurrence_frequency" json:"recurrence_frequency,omitempty"`
	// Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreateContractRecurringCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
	// determines the start time for the first commit
	StartingAt *string `form:"starting_at" json:"starting_at"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig *V1ContractsCreateContractRecurringCommitSubscriptionConfigParams `form:"subscription_config" json:"subscription_config,omitempty"`
	// A temporary ID that can be used to reference the recurring commit for commit specific overrides.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
}

// The amount of commit to grant.
type V1ContractsCreateContractRecurringCreditAccessAmountParams struct {
	// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is required unless a subscription is attached via `subscription_config`.
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	UnitPrice *Number `form:"unit_price" json:"unit_price"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreateContractRecurringCreditSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreateContractRecurringCreditSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreateContractRecurringCreditSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.
type V1ContractsCreateContractRecurringCreditCommitDurationParams struct {
	Unit  *string `form:"unit" json:"unit,omitempty"`
	Value *Number `form:"value" json:"value"`
}
type V1ContractsCreateContractRecurringCreditSubscriptionConfigApplySeatIncreaseConfigParams struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated *bool `form:"is_prorated" json:"is_prorated"`
}

// Attach a subscription to the recurring commit/credit.
type V1ContractsCreateContractRecurringCreditSubscriptionConfigParams struct {
	// Controls when a customer can access recurring commit or credit child balances when the attached subscription is payment-gated. `BILLING_PERIOD_PAID` releases each balance after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases all balances after the first payment.
	AccessPolicy            *string                                                                                  `form:"access_policy" json:"access_policy,omitempty"`
	Allocation              *string                                                                                  `form:"allocation" json:"allocation,omitempty"`
	ApplySeatIncreaseConfig *V1ContractsCreateContractRecurringCreditSubscriptionConfigApplySeatIncreaseConfigParams `form:"apply_seat_increase_config" json:"apply_seat_increase_config"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID *string `form:"subscription_id" json:"subscription_id"`
}

// Optional configuration for recurring commit/credit hierarchy access control
type V1ContractsCreateContractRecurringCreditHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}
type V1ContractsCreateContractRecurringCreditProrationRoundingAccessParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}

// Optional rounding configuration for prorated recurring credit amounts.
type V1ContractsCreateContractRecurringCreditProrationRoundingParams struct {
	Access *V1ContractsCreateContractRecurringCreditProrationRoundingAccessParams `form:"access" json:"access,omitempty"`
}
type V1ContractsCreateContractRecurringCreditParams struct {
	// The amount of commit to grant.
	AccessAmount *V1ContractsCreateContractRecurringCreditAccessAmountParams `form:"access_amount" json:"access_amount"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Will be passed down to the individual commits
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	// Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.
	CommitDuration *V1ContractsCreateContractRecurringCreditCommitDurationParams `form:"commit_duration" json:"commit_duration"`
	// Will be passed down to the individual commits
	Description *string `form:"description" json:"description,omitempty"`
	// Determines when the contract will stop creating recurring commits. optional
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// Optional configuration for recurring commit/credit hierarchy access control
	HierarchyConfiguration *V1ContractsCreateContractRecurringCreditHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// displayed on invoices. will be passed through to the individual commits
	Name *string `form:"name" json:"name,omitempty"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// Will be passed down to the individual commits
	Priority  *Number `form:"priority" json:"priority"`
	ProductID *string `form:"product_id" json:"product_id"`
	// Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration *string `form:"proration" json:"proration,omitempty"`
	// Optional rounding configuration for prorated recurring credit amounts.
	ProrationRounding *V1ContractsCreateContractRecurringCreditProrationRoundingParams `form:"proration_rounding" json:"proration_rounding,omitempty"`
	// Whether the created commits will use the commit rate or list rate
	RateType *string `form:"rate_type" json:"rate_type,omitempty"`
	// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
	RecurrenceFrequency *string `form:"recurrence_frequency" json:"recurrence_frequency,omitempty"`
	// Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreateContractRecurringCreditSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
	// determines the start time for the first commit
	StartingAt *string `form:"starting_at" json:"starting_at"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig *V1ContractsCreateContractRecurringCreditSubscriptionConfigParams `form:"subscription_config" json:"subscription_config,omitempty"`
	// A temporary ID that can be used to reference the recurring commit for commit specific overrides.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
}

// Only set for TIERED rate_type.
type V1ContractsCreateContractOverrideOverwriteRateTierParams struct {
	Price *Number `form:"price" json:"price"`
	Size  *Number `form:"size" json:"size,omitempty"`
}
type V1ContractsCreateContractOverrideOverwriteRateMinimumConfigParams struct {
	Minimum *Number `form:"minimum" json:"minimum"`
}

// Required for OVERWRITE type.
type V1ContractsCreateContractOverrideOverwriteRateParams struct {
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate processors.
	CustomRate map[string]JSON `form:"custom_rate" json:"custom_rate,omitempty"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be set to true.
	IsProrated    *bool                                                              `form:"is_prorated" json:"is_prorated,omitempty"`
	MinimumConfig *V1ContractsCreateContractOverrideOverwriteRateMinimumConfigParams `form:"minimum_config" json:"minimum_config,omitempty"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price *Number `form:"price" json:"price,omitempty"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity *Number `form:"quantity" json:"quantity,omitempty"`
	RateType *string `form:"rate_type" json:"rate_type"`
	// Only set for TIERED rate_type.
	Tiers []*V1ContractsCreateContractOverrideOverwriteRateTierParams `form:"tiers" json:"tiers,omitempty"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreateContractOverrideOverrideSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Cannot be used in conjunction with product_id or applicable_product_tags. If provided, the override will apply to all products with the specified specifiers.
type V1ContractsCreateContractOverrideOverrideSpecifierParams struct {
	// Can only be used for commit specific overrides. Must be used in conjunction with one of `product_id`, `product_tags`, `pricing_group_values`, or `presentation_group_values`. Must be used instead of both `commit_ids` and `recurring_commit_ids` If provided, the override will apply to any specified commit, credit, recurring commit or recurring credit IDs.
	AnyCommitOrCreditIDs []*string `form:"any_commit_or_credit_ids" json:"any_commit_or_credit_ids,omitempty"`
	BillingFrequency     *string   `form:"billing_frequency" json:"billing_frequency,omitempty"`
	// Can only be used for commit specific overrides. Must be used in conjunction with one of `product_id`, `product_tags`, `pricing_group_values`, or `presentation_group_values`. If provided, the override will only apply to the specified commits. If not provided, the override will apply to all commits.
	CommitIDs []*string `form:"commit_ids" json:"commit_ids,omitempty"`
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreateContractOverrideOverrideSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// A map of group names to values. The override will only apply to line items with the specified presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// A map of pricing group names to values. The override will only apply to products with the specified pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the override will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the override will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
	// Can only be used for commit specific overrides. Must be used in conjunction with one of `product_id`, `product_tags`, `pricing_group_values`, or `presentation_group_values`. If provided, the override will only apply to commits created by the specified recurring commit ids.
	RecurringCommitIDs []*string `form:"recurring_commit_ids" json:"recurring_commit_ids,omitempty"`
}

// Required for TIERED type. Must have at least one tier.
type V1ContractsCreateContractOverrideTierParams struct {
	Multiplier *Number `form:"multiplier" json:"multiplier"`
	Size       *Number `form:"size" json:"size,omitempty"`
}
type V1ContractsCreateContractOverrideParams struct {
	// tags identifying products whose rates are being overridden. Cannot be used in conjunction with override_specifiers.
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	Entitled     *bool   `form:"entitled" json:"entitled,omitempty"`
	// Indicates whether the override should only apply to commits. Defaults to `false`. If `true` you can specify relevant commits in `override_specifiers` by passing `commit_ids`, `recurring_commit_ids`, or `any_commit_or_credit_ids`.  If you do not specify any of these fields, the override will apply when consuming any prepaid commit, postpaid commit, or credit
	IsCommitSpecific *bool `form:"is_commit_specific" json:"is_commit_specific,omitempty"`
	// Required for MULTIPLIER type. Must be >=0.
	Multiplier *Number `form:"multiplier" json:"multiplier,omitempty"`
	// Cannot be used in conjunction with product_id or applicable_product_tags. If provided, the override will apply to all products with the specified specifiers.
	OverrideSpecifiers []*V1ContractsCreateContractOverrideOverrideSpecifierParams `form:"override_specifiers" json:"override_specifiers,omitempty"`
	// Required for OVERWRITE type.
	OverwriteRate *V1ContractsCreateContractOverrideOverwriteRateParams `form:"overwrite_rate" json:"overwrite_rate,omitempty"`
	// Required for EXPLICIT multiplier prioritization scheme and all TIERED overrides. Under EXPLICIT prioritization, overwrites are prioritized first, and then tiered and multiplier overrides are prioritized by their priority value (lowest first). Must be > 0.
	Priority *Number `form:"priority" json:"priority,omitempty"`
	// ID of the product whose rate is being overridden. Cannot be used in conjunction with override_specifiers.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// RFC 3339 timestamp indicating when the override will start applying (inclusive)
	StartingAt *string `form:"starting_at" json:"starting_at"`
	// Indicates whether the override applies to commit rates or list rates. Can only be used for overrides that have `is_commit_specific` set to `true`. Defaults to `"LIST_RATE"`.
	Target *string `form:"target" json:"target,omitempty"`
	// Required for TIERED type. Must have at least one tier.
	Tiers []*V1ContractsCreateContractOverrideTierParams `form:"tiers" json:"tiers,omitempty"`
	// Overwrites are prioritized over multipliers and tiered overrides.
	Type *string `form:"type" json:"type,omitempty"`
}

// Either provide amount or provide both unit_price and quantity.
type V1ContractsCreateContractDiscountScheduleScheduleItemParams struct {
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
type V1ContractsCreateContractDiscountScheduleRecurringScheduleParams struct {
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
type V1ContractsCreateContractDiscountScheduleParams struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is only applicable to commit invoice schedules. If true, this schedule will not generate an invoice.
	DoNotInvoice *bool `form:"do_not_invoice" json:"do_not_invoice,omitempty"`
	// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
	RecurringSchedule *V1ContractsCreateContractDiscountScheduleRecurringScheduleParams `form:"recurring_schedule" json:"recurring_schedule,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []*V1ContractsCreateContractDiscountScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items,omitempty"`
}

// This field's availability is dependent on your client's configuration.
type V1ContractsCreateContractDiscountParams struct {
	CustomFields map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string                                          `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	ProductID            *string                                          `form:"product_id" json:"product_id"`
	Schedule             *V1ContractsCreateContractDiscountScheduleParams `form:"schedule" json:"schedule"`
}

// This field's availability is dependent on your client's configuration.
type V1ContractsCreateContractProfessionalServiceParams struct {
	CustomFields map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	Description  *string           `form:"description" json:"description,omitempty"`
	// Maximum amount for the term.
	MaxAmount *Number `form:"max_amount" json:"max_amount"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	ProductID            *string `form:"product_id" json:"product_id"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the amount.
	Quantity *Number `form:"quantity" json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified.
	UnitPrice *Number `form:"unit_price" json:"unit_price"`
}
type V1ContractsCreateContractResellerRoyaltyAwsOptionsParams struct {
	AwsAccountNumber    *string `form:"aws_account_number" json:"aws_account_number,omitempty"`
	AwsOfferID          *string `form:"aws_offer_id" json:"aws_offer_id,omitempty"`
	AwsPayerReferenceID *string `form:"aws_payer_reference_id" json:"aws_payer_reference_id,omitempty"`
}
type V1ContractsCreateContractResellerRoyaltyGcpOptionsParams struct {
	GcpAccountID *string `form:"gcp_account_id" json:"gcp_account_id,omitempty"`
	GcpOfferID   *string `form:"gcp_offer_id" json:"gcp_offer_id,omitempty"`
}

// This field's availability is dependent on your client's configuration.
type V1ContractsCreateContractResellerRoyaltyParams struct {
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductTags []*string                                                 `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	AwsOptions            *V1ContractsCreateContractResellerRoyaltyAwsOptionsParams `form:"aws_options" json:"aws_options,omitempty"`
	EndingBefore          *string                                                   `form:"ending_before" json:"ending_before,omitempty"`
	Fraction              *Number                                                   `form:"fraction" json:"fraction"`
	GcpOptions            *V1ContractsCreateContractResellerRoyaltyGcpOptionsParams `form:"gcp_options" json:"gcp_options,omitempty"`
	NetsuiteResellerID    *string                                                   `form:"netsuite_reseller_id" json:"netsuite_reseller_id"`
	ResellerContractValue *Number                                                   `form:"reseller_contract_value" json:"reseller_contract_value,omitempty"`
	ResellerType          *string                                                   `form:"reseller_type" json:"reseller_type"`
	StartingAt            *string                                                   `form:"starting_at" json:"starting_at"`
}

// Either provide amount or provide both unit_price and quantity.
type V1ContractsCreateContractScheduledChargeScheduleScheduleItemParams struct {
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
type V1ContractsCreateContractScheduledChargeScheduleRecurringScheduleParams struct {
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
type V1ContractsCreateContractScheduledChargeScheduleParams struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is only applicable to commit invoice schedules. If true, this schedule will not generate an invoice.
	DoNotInvoice *bool `form:"do_not_invoice" json:"do_not_invoice,omitempty"`
	// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
	RecurringSchedule *V1ContractsCreateContractScheduledChargeScheduleRecurringScheduleParams `form:"recurring_schedule" json:"recurring_schedule,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []*V1ContractsCreateContractScheduledChargeScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items,omitempty"`
}
type V1ContractsCreateContractScheduledChargeParams struct {
	CustomFields map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string                                                 `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	ProductID            *string                                                 `form:"product_id" json:"product_id"`
	Schedule             *V1ContractsCreateContractScheduledChargeScheduleParams `form:"schedule" json:"schedule"`
}
type V1ContractsCreateContractTransitionFutureInvoiceBehaviorParams struct {
	// Controls whether future trueup invoices are billed or removed. Default behavior is AS_IS if not specified.
	Trueup *string `form:"trueup" json:"trueup,omitempty"`
}
type V1ContractsCreateContractTransitionParams struct {
	FromContractID        *string                                                         `form:"from_contract_id" json:"from_contract_id"`
	FutureInvoiceBehavior *V1ContractsCreateContractTransitionFutureInvoiceBehaviorParams `form:"future_invoice_behavior" json:"future_invoice_behavior,omitempty"`
	// This field's available values may vary based on your client's configuration.
	Type *string `form:"type" json:"type"`
}
type V1ContractsCreateContractUsageFilterParams struct {
	GroupKey    *string   `form:"group_key" json:"group_key"`
	GroupValues []*string `form:"group_values" json:"group_values"`
	StartingAt  *string   `form:"starting_at" json:"starting_at,omitempty"`
}
type V1ContractsCreateContractUsageStatementScheduleParams struct {
	// Required when using CUSTOM_DATE. This option lets you set a historical billing anchor date, aligning future billing cycles with a chosen cadence. For example, if a contract starts on 2024-09-15 and you set the anchor date to 2024-09-10 with a MONTHLY frequency, the first usage statement will cover 09-15 to 10-10. Subsequent statements will follow the 10th of each month.
	BillingAnchorDate *string `form:"billing_anchor_date" json:"billing_anchor_date,omitempty"`
	// If not provided, defaults to the first day of the month.
	Day       *string `form:"day" json:"day,omitempty"`
	Frequency *string `form:"frequency" json:"frequency"`
	// The date Metronome should start generating usage invoices. If unspecified, contract start date will be used. This is useful to set if you want to import historical invoices via our 'Create Historical Invoices' API rather than having Metronome automatically generate them.
	InvoiceGenerationStartingAt *string `form:"invoice_generation_starting_at" json:"invoice_generation_starting_at,omitempty"`
}
type V1ContractsCreateContractBillingProviderConfigurationParams struct {
	// Do not specify if using billing_provider_configuration_id.
	BillingProvider *string `form:"billing_provider" json:"billing_provider,omitempty"`
	// The Metronome ID of the billing provider configuration. Use when a customer has multiple configurations with the same billing provider and delivery method. Otherwise, specify the billing_provider and delivery_method.
	BillingProviderConfigurationID *string `form:"billing_provider_configuration_id" json:"billing_provider_configuration_id,omitempty"`
	// Do not specify if using billing_provider_configuration_id.
	DeliveryMethod *string `form:"delivery_method" json:"delivery_method,omitempty"`
}
type V1ContractsCreateContractRevenueSystemConfigurationParams struct {
	// How revenue recognition records should be delivered to the revenue system. Do not specify if using revenue_system_configuration_id.
	DeliveryMethod *string `form:"delivery_method" json:"delivery_method,omitempty"`
	// The system that is providing services for revenue recognition. Do not specify if using revenue_system_configuration_id.
	Provider *string `form:"provider" json:"provider,omitempty"`
	// The Metronome ID of the revenue system configuration. Use when a customer has multiple configurations with the same provider and delivery method. Otherwise, specify the provider and delivery_method.
	RevenueSystemConfigurationID *string `form:"revenue_system_configuration_id" json:"revenue_system_configuration_id,omitempty"`
}
type V1ContractsCreateContractSpendThresholdConfigurationCommitParams struct {
	Description *string `form:"description" json:"description,omitempty"`
	// Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.
	Name *string `form:"name" json:"name,omitempty"`
	// The priority of the commit, used to determine drawdown order. Lower priority commits are consumed first. Defaults to 100 if not specified.
	Priority *Number `form:"priority" json:"priority,omitempty"`
	// The commit product that will be used to generate the line item for commit payment.
	ProductID *string `form:"product_id" json:"product_id"`
}

// Only applicable if using STRIPE as your payment gate type.
type V1ContractsCreateContractSpendThresholdConfigurationPaymentGateConfigStripeConfigParams struct {
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as your payment type.
	InvoiceMetadata map[string]string `form:"invoice_metadata" json:"invoice_metadata,omitempty"`
	// If left blank, will default to INVOICE
	PaymentType *string `form:"payment_type" json:"payment_type"`
}

// Only applicable if using PRECALCULATED as your tax type.
type V1ContractsCreateContractSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams struct {
	// Amount of tax to be applied. This should be in the same currency and denomination  as the commit's invoice schedule
	TaxAmount *Number `form:"tax_amount" json:"tax_amount"`
	// Name of the tax to be applied. This may be used in an invoice line item description.
	TaxName *string `form:"tax_name" json:"tax_name,omitempty"`
}
type V1ContractsCreateContractSpendThresholdConfigurationPaymentGateConfigParams struct {
	// Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.
	PaymentGateType *string `form:"payment_gate_type" json:"payment_gate_type"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig *V1ContractsCreateContractSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams `form:"precalculated_tax_config" json:"precalculated_tax_config,omitempty"`
	// Only applicable if using STRIPE as your payment gate type.
	StripeConfig *V1ContractsCreateContractSpendThresholdConfigurationPaymentGateConfigStripeConfigParams `form:"stripe_config" json:"stripe_config,omitempty"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType *string `form:"tax_type" json:"tax_type,omitempty"`
}

// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
type V1ContractsCreateContractSpendThresholdConfigurationDiscountConfigurationCapParams struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount *Number `form:"amount" json:"amount"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias *string `form:"spend_tracker_alias" json:"spend_tracker_alias"`
}
type V1ContractsCreateContractSpendThresholdConfigurationDiscountConfigurationParams struct {
	// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
	Cap *V1ContractsCreateContractSpendThresholdConfigurationDiscountConfigurationCapParams `form:"cap" json:"cap,omitempty"`
	// The fraction of the original amount that the customer pays after applying the discount. For example, 0.85 means the customer pays 85% of the original amount (a 15% discount).
	PaymentFraction *Number `form:"payment_fraction" json:"payment_fraction"`
}
type V1ContractsCreateContractSpendThresholdConfigurationParams struct {
	Commit                *V1ContractsCreateContractSpendThresholdConfigurationCommitParams                `form:"commit" json:"commit"`
	DiscountConfiguration *V1ContractsCreateContractSpendThresholdConfigurationDiscountConfigurationParams `form:"discount_configuration" json:"discount_configuration,omitempty"`
	// When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.
	IsEnabled         *bool                                                                        `form:"is_enabled" json:"is_enabled"`
	PaymentGateConfig *V1ContractsCreateContractSpendThresholdConfigurationPaymentGateConfigParams `form:"payment_gate_config" json:"payment_gate_config"`
	// Specify the threshold amount for the contract. Each time the contract's usage hits this amount, a threshold charge will be initiated.
	ThresholdAmount *Number `form:"threshold_amount" json:"threshold_amount"`
}

// The length of time the created commit will be valid, starting from the end of the invoice's service period. If not provided, defaults to one year.
type V1ContractsCreateContractPrepaidBalanceThresholdConfigurationCommitDurationParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreateContractPrepaidBalanceThresholdConfigurationCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreateContractPrepaidBalanceThresholdConfigurationCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreateContractPrepaidBalanceThresholdConfigurationCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}
type V1ContractsCreateContractPrepaidBalanceThresholdConfigurationCommitParams struct {
	// Which products the threshold commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the threshold commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	Description           *string   `form:"description" json:"description,omitempty"`
	// The length of time the created commit will be valid, starting from the end of the invoice's service period. If not provided, defaults to one year.
	Duration *V1ContractsCreateContractPrepaidBalanceThresholdConfigurationCommitDurationParams `form:"duration" json:"duration,omitempty"`
	// Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.
	Name *string `form:"name" json:"name,omitempty"`
	// The priority of the commit, used to determine drawdown order. Lower priority commits are consumed first. Defaults to 100 if not specified.
	Priority *Number `form:"priority" json:"priority,omitempty"`
	// The commit product that will be used to generate the line item for commit payment.
	ProductID *string `form:"product_id" json:"product_id"`
	// Whether the created commits will be charged at commit rate or list rate.
	RateType *string `form:"rate_type" json:"rate_type,omitempty"`
	// Fraction of the created commit's unused balance that will roll over. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreateContractPrepaidBalanceThresholdConfigurationCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
}

// Only applicable if using STRIPE as your payment gate type.
type V1ContractsCreateContractPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigParams struct {
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as your payment type.
	InvoiceMetadata map[string]string `form:"invoice_metadata" json:"invoice_metadata,omitempty"`
	// If left blank, will default to INVOICE
	PaymentType *string `form:"payment_type" json:"payment_type"`
}

// Only applicable if using PRECALCULATED as your tax type.
type V1ContractsCreateContractPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams struct {
	// Amount of tax to be applied. This should be in the same currency and denomination  as the commit's invoice schedule
	TaxAmount *Number `form:"tax_amount" json:"tax_amount"`
	// Name of the tax to be applied. This may be used in an invoice line item description.
	TaxName *string `form:"tax_name" json:"tax_name,omitempty"`
}
type V1ContractsCreateContractPrepaidBalanceThresholdConfigurationPaymentGateConfigParams struct {
	// Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.
	PaymentGateType *string `form:"payment_gate_type" json:"payment_gate_type"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig *V1ContractsCreateContractPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams `form:"precalculated_tax_config" json:"precalculated_tax_config,omitempty"`
	// Only applicable if using STRIPE as your payment gate type.
	StripeConfig *V1ContractsCreateContractPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigParams `form:"stripe_config" json:"stripe_config,omitempty"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType *string `form:"tax_type" json:"tax_type,omitempty"`
}

// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
type V1ContractsCreateContractPrepaidBalanceThresholdConfigurationDiscountConfigurationCapParams struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount *Number `form:"amount" json:"amount"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias *string `form:"spend_tracker_alias" json:"spend_tracker_alias"`
}
type V1ContractsCreateContractPrepaidBalanceThresholdConfigurationDiscountConfigurationParams struct {
	// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
	Cap *V1ContractsCreateContractPrepaidBalanceThresholdConfigurationDiscountConfigurationCapParams `form:"cap" json:"cap,omitempty"`
	// The fraction of the original amount that the customer pays after applying the discount. For example, 0.85 means the customer pays 85% of the original amount (a 15% discount).
	PaymentFraction *Number `form:"payment_fraction" json:"payment_fraction"`
}

// If provided, balances with all the custom fields will not be considered when evaluating threshold billing
type V1ContractsCreateContractPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilterParams struct {
	Entity *string `form:"entity" json:"entity"`
	Key    *string `form:"key" json:"key"`
	Value  *string `form:"value" json:"value"`
}

// If any of the exclude specifier is met, the balance is not considered when evaluating threshold billing
type V1ContractsCreateContractPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeParams struct {
	// If provided, balances with all the custom fields will not be considered when evaluating threshold billing
	CustomFieldFilters []*V1ContractsCreateContractPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilterParams `form:"custom_field_filters" json:"custom_field_filters"`
}

// Determines which balances are excluded from remaining balance calculation for threshold billing.
type V1ContractsCreateContractPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierParams struct {
	// If any of the exclude specifier is met, the balance is not considered when evaluating threshold billing
	Exclude []*V1ContractsCreateContractPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeParams `form:"exclude" json:"exclude"`
}
type V1ContractsCreateContractPrepaidBalanceThresholdConfigurationParams struct {
	Commit *V1ContractsCreateContractPrepaidBalanceThresholdConfigurationCommitParams `form:"commit" json:"commit"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID    *string                                                                                   `form:"custom_credit_type_id" json:"custom_credit_type_id,omitempty"`
	DiscountConfiguration *V1ContractsCreateContractPrepaidBalanceThresholdConfigurationDiscountConfigurationParams `form:"discount_configuration" json:"discount_configuration,omitempty"`
	// When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.
	IsEnabled         *bool                                                                                 `form:"is_enabled" json:"is_enabled"`
	PaymentGateConfig *V1ContractsCreateContractPrepaidBalanceThresholdConfigurationPaymentGateConfigParams `form:"payment_gate_config" json:"payment_gate_config"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount *Number `form:"recharge_to_amount" json:"recharge_to_amount"`
	// Specify the threshold amount for the contract. Each time the contract's prepaid balance lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount *Number `form:"threshold_amount" json:"threshold_amount"`
	// Determines which balances are excluded from remaining balance calculation for threshold billing.
	ThresholdBalanceSpecifiers []*V1ContractsCreateContractPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierParams `form:"threshold_balance_specifiers" json:"threshold_balance_specifiers,omitempty"`
}
type V1ContractsCreateContractSpendTrackerApplicableSpendSpecifierParams struct {
	// Filter by whether the spend was discounted. Defaults to ANY if omitted.
	Discounted *string   `form:"discounted" json:"discounted,omitempty"`
	Sources    []*string `form:"sources" json:"sources"`
	SpendType  *string   `form:"spend_type" json:"spend_type"`
}

// Spend trackers to attach to this contract. Aliases must be unique within a contract.
type V1ContractsCreateContractSpendTrackerParams struct {
	// Human-readable identifier, unique per contract.
	Alias                     *string                                                                `form:"alias" json:"alias"`
	ApplicableSpendSpecifiers []*V1ContractsCreateContractSpendTrackerApplicableSpendSpecifierParams `form:"applicable_spend_specifiers" json:"applicable_spend_specifiers"`
	CreditTypeID              *string                                                                `form:"credit_type_id" json:"credit_type_id"`
	ResetFrequency            *string                                                                `form:"reset_frequency" json:"reset_frequency"`
}
type V1ContractsCreateContractSubscriptionSubscriptionRateParams struct {
	// Frequency to bill subscription with. Together with product_id, must match existing rate on the rate card.
	BillingFrequency *string `form:"billing_frequency" json:"billing_frequency"`
	// Must be subscription type product
	ProductID *string `form:"product_id" json:"product_id"`
}
type V1ContractsCreateContractSubscriptionProrationRoundingParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}
type V1ContractsCreateContractSubscriptionProrationParams struct {
	// Indicates how mid-period quantity adjustments are invoiced.
	// **BILL_IMMEDIATELY**: Only available when collection schedule is `ADVANCE`. The quantity increase will be billed immediately on the scheduled date.
	// **BILL_ON_NEXT_COLLECTION_DATE**: The quantity increase will be billed for in-arrears at the end of the period.
	//
	InvoiceBehavior *string `form:"invoice_behavior" json:"invoice_behavior,omitempty"`
	// Indicates if the partial period will be prorated or charged a full amount.
	IsProrated *bool                                                         `form:"is_prorated" json:"is_prorated,omitempty"`
	Rounding   *V1ContractsCreateContractSubscriptionProrationRoundingParams `form:"rounding" json:"rounding,omitempty"`
}
type V1ContractsCreateContractSubscriptionSeatConfigParams struct {
	// The initial assigned seats on this subscription.
	InitialSeatIDs []*string `form:"initial_seat_ids" json:"initial_seat_ids"`
	// The initial amount of unassigned seats on this subscription.
	InitialUnassignedSeats *Number `form:"initial_unassigned_seats" json:"initial_unassigned_seats,omitempty"`
	// The property name, sent on usage events, that identifies the seat ID associated with the usage event.  For example, the property name might be seat_id or user_id. The property must be set as a group key on billable metrics and a presentation/pricing group key on contract products.  This allows linked recurring credits with an allocation per seat to be consumed by only one seat's usage.
	SeatGroupKey *string `form:"seat_group_key" json:"seat_group_key"`
}
type V1ContractsCreateContractSubscriptionBillingCycleConfigParams struct {
	// The date to anchor the billing cycle to. If omitted, defaults to the contract's usage invoice billing cycle anchor date.
	AnchorDate *string `form:"anchor_date" json:"anchor_date,omitempty"`
	// Controls whether this subscription consolidates onto usage invoices or gets its own scheduled invoice. Defaults to ON_USAGE_INVOICE if omitted.
	InvoicePlacement *string `form:"invoice_placement" json:"invoice_placement,omitempty"`
}
type V1ContractsCreateContractSubscriptionPaymentGateConfigParams struct {
	// Gate access to the subscription based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe.
	PaymentGateType *string `form:"payment_gate_type" json:"payment_gate_type"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType *string `form:"tax_type" json:"tax_type,omitempty"`
}
type V1ContractsCreateContractSubscriptionParams struct {
	BillingCycleConfig *V1ContractsCreateContractSubscriptionBillingCycleConfigParams `form:"billing_cycle_config" json:"billing_cycle_config,omitempty"`
	CollectionSchedule *string                                                        `form:"collection_schedule" json:"collection_schedule"`
	CustomFields       map[string]string                                              `form:"custom_fields" json:"custom_fields,omitempty"`
	Description        *string                                                        `form:"description" json:"description,omitempty"`
	// Exclusive end time for the subscription. If not provided, subscription inherits contract end date.
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// The initial quantity for the subscription. It must be non-negative value. Required if quantity_management_mode is QUANTITY_ONLY.
	InitialQuantity   *Number                                                       `form:"initial_quantity" json:"initial_quantity,omitempty"`
	Name              *string                                                       `form:"name" json:"name,omitempty"`
	PaymentGateConfig *V1ContractsCreateContractSubscriptionPaymentGateConfigParams `form:"payment_gate_config" json:"payment_gate_config,omitempty"`
	Proration         *V1ContractsCreateContractSubscriptionProrationParams         `form:"proration" json:"proration"`
	// Determines how the subscription's quantity is controlled. Defaults to QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified directly on the subscription. `initial_quantity` must be provided with this option. Compatible with recurring commits/credits that use POOLED allocation. **SEAT_BASED**: Use when you want to pass specific seat identifiers (e.g. add user_123) to increment and decrement a subscription quantity, rather than directly providing the quantity. You must use a **SEAT_BASED** subscription to use a linked recurring credit with an allocation per seat. `seat_config` must be provided with this option.
	QuantityManagementMode *string                                                `form:"quantity_management_mode" json:"quantity_management_mode,omitempty"`
	SeatConfig             *V1ContractsCreateContractSubscriptionSeatConfigParams `form:"seat_config" json:"seat_config,omitempty"`
	// Inclusive start time for the subscription. If not provided, defaults to contract start date
	StartingAt       *string                                                      `form:"starting_at" json:"starting_at,omitempty"`
	SubscriptionRate *V1ContractsCreateContractSubscriptionSubscriptionRateParams `form:"subscription_rate" json:"subscription_rate"`
	// A temporary ID used to reference the subscription in recurring commit/credit subscription configs created within the same payload.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
}
type V1ContractsCreateContractHierarchyConfigurationParentParams struct {
	ContractID *string `form:"contract_id" json:"contract_id"`
	CustomerID *string `form:"customer_id" json:"customer_id"`
}
type V1ContractsCreateContractHierarchyConfigurationParentBehaviorParams struct {
	// Indicates the desired behavior of consolidated invoices generated by the parent in a customer hierarchy
	//
	// **CONCATENATE**: Statements on the invoices of child customers will be appended to the consolidated invoice
	//
	// **NONE**: Do not generate consolidated invoices
	InvoiceConsolidationType *string `form:"invoice_consolidation_type" json:"invoice_consolidation_type,omitempty"`
}
type V1ContractsCreateContractHierarchyConfigurationParams struct {
	Parent         *V1ContractsCreateContractHierarchyConfigurationParentParams         `form:"parent" json:"parent,omitempty"`
	ParentBehavior *V1ContractsCreateContractHierarchyConfigurationParentBehaviorParams `form:"parent_behavior" json:"parent_behavior,omitempty"`
	// Indicates which customer should pay for the child's invoice charges
	//
	// **SELF**: The child pays for its own invoice charges
	//
	// **PARENT**: The parent pays for the child's invoice charges
	Payer *string `form:"payer" json:"payer,omitempty"`
	// Indicates the behavior of the child's invoice statements on the parent's invoices.
	//
	// **CONSOLIDATE**: Child's invoice statements will be added to parent's consolidated invoices
	//
	// **SEPARATE**: Child's invoice statements will appear not appear on parent's consolidated invoices
	UsageStatementBehavior *string `form:"usage_statement_behavior" json:"usage_statement_behavior,omitempty"`
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
type V1ContractsCreateContractParams struct {
	Params                       `form:"*"`
	BillingProviderConfiguration *V1ContractsCreateContractBillingProviderConfigurationParams `form:"billing_provider_configuration" json:"billing_provider_configuration,omitempty"`
	Commits                      []*V1ContractsCreateContractCommitParams                     `form:"commits" json:"commits,omitempty"`
	Credits                      []*V1ContractsCreateContractCreditParams                     `form:"credits" json:"credits,omitempty"`
	CustomerID                   *string                                                      `form:"customer_id" json:"customer_id"`
	CustomFields                 map[string]string                                            `form:"custom_fields" json:"custom_fields,omitempty"`
	// This field's availability is dependent on your client's configuration.
	Discounts []*V1ContractsCreateContractDiscountParams `form:"discounts" json:"discounts,omitempty"`
	// exclusive contract end time
	EndingBefore           *string                                                `form:"ending_before" json:"ending_before,omitempty"`
	HierarchyConfiguration *V1ContractsCreateContractHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list prices automatically. EXPLICIT prioritization requires specifying priorities for each multiplier; the one with the lowest priority value will be prioritized first. If tiered overrides are used, prioritization must be explicit.
	MultiplierOverridePrioritization *string `form:"multiplier_override_prioritization" json:"multiplier_override_prioritization,omitempty"`
	Name                             *string `form:"name" json:"name,omitempty"`
	NetPaymentTermsDays              *Number `form:"net_payment_terms_days" json:"net_payment_terms_days,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string                                    `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	Overrides            []*V1ContractsCreateContractOverrideParams `form:"overrides" json:"overrides,omitempty"`
	// Selects the package linked to the specified alias as of the contract's start date. Mutually exclusive with package_id.
	PackageAlias          *string                                               `form:"package_alias" json:"package_alias,omitempty"`
	PackageCustomizations *V1ContractsCreateContractPackageCustomizationsParams `form:"package_customizations" json:"package_customizations,omitempty"`
	// If provided, provisions a customer on a package instead of creating a traditional contract. When specified, only customer_id, starting_at, package_id, uniqueness_key, transition, and custom_fields are allowed.
	PackageID                            *string                                                              `form:"package_id" json:"package_id,omitempty"`
	PrepaidBalanceThresholdConfiguration *V1ContractsCreateContractPrepaidBalanceThresholdConfigurationParams `form:"prepaid_balance_threshold_configuration" json:"prepaid_balance_threshold_configuration,omitempty"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []*V1ContractsCreateContractProfessionalServiceParams `form:"professional_services" json:"professional_services,omitempty"`
	// Selects the rate card linked to the specified alias as of the contract's start date.
	RateCardAlias    *string                                           `form:"rate_card_alias" json:"rate_card_alias,omitempty"`
	RateCardID       *string                                           `form:"rate_card_id" json:"rate_card_id,omitempty"`
	RecurringCommits []*V1ContractsCreateContractRecurringCommitParams `form:"recurring_commits" json:"recurring_commits,omitempty"`
	RecurringCredits []*V1ContractsCreateContractRecurringCreditParams `form:"recurring_credits" json:"recurring_credits,omitempty"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties          []*V1ContractsCreateContractResellerRoyaltyParams          `form:"reseller_royalties" json:"reseller_royalties,omitempty"`
	RevenueSystemConfiguration *V1ContractsCreateContractRevenueSystemConfigurationParams `form:"revenue_system_configuration" json:"revenue_system_configuration,omitempty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID         *string                                                     `form:"salesforce_opportunity_id" json:"salesforce_opportunity_id,omitempty"`
	ScheduledCharges                []*V1ContractsCreateContractScheduledChargeParams           `form:"scheduled_charges" json:"scheduled_charges,omitempty"`
	ScheduledChargesOnUsageInvoices *string                                                     `form:"scheduled_charges_on_usage_invoices" json:"scheduled_charges_on_usage_invoices,omitempty"`
	SpendThresholdConfiguration     *V1ContractsCreateContractSpendThresholdConfigurationParams `form:"spend_threshold_configuration" json:"spend_threshold_configuration,omitempty"`
	// Spend trackers to attach to this contract. Aliases must be unique within a contract.
	SpendTrackers []*V1ContractsCreateContractSpendTrackerParams `form:"spend_trackers" json:"spend_trackers,omitempty"`
	// inclusive contract start time
	StartingAt    *string                                        `form:"starting_at" json:"starting_at"`
	Subscriptions []*V1ContractsCreateContractSubscriptionParams `form:"subscriptions" json:"subscriptions,omitempty"`
	// This field's availability is dependent on your client's configuration.
	TotalContractValue     *Number                                                `form:"total_contract_value" json:"total_contract_value,omitempty"`
	Transition             *V1ContractsCreateContractTransitionParams             `form:"transition" json:"transition,omitempty"`
	UniquenessKey          *string                                                `form:"uniqueness_key" json:"uniqueness_key,omitempty"`
	UsageFilter            *V1ContractsCreateContractUsageFilterParams            `form:"usage_filter" json:"usage_filter,omitempty"`
	UsageStatementSchedule *V1ContractsCreateContractUsageStatementScheduleParams `form:"usage_statement_schedule" json:"usage_statement_schedule,omitempty"`
}
type V1ContractsCreateCustomerWithContractCustomerCustomerBillingProviderConfigurationParams struct {
	// The billing provider set for this configuration.
	BillingProvider *string `form:"billing_provider" json:"billing_provider"`
	// Configuration for the billing provider. The structure of this object is specific to the billing provider and delivery provider combination. Defaults to an empty object, however, for most billing provider + delivery method combinations, it will not be a valid configuration.
	Configuration map[string]JSON `form:"configuration" json:"configuration,omitempty"`
	// The method to use for delivering invoices to this customer. If not provided, the `delivery_method_id` must be provided.
	DeliveryMethod *string `form:"delivery_method" json:"delivery_method,omitempty"`
	// ID of the delivery method to use for this customer. If not provided, the `delivery_method` must be provided.
	DeliveryMethodID *string `form:"delivery_method_id" json:"delivery_method_id,omitempty"`
	// Specifies which tax provider Metronome should use for tax calculation when billing through Stripe. This is only supported for Stripe billing provider configurations.
	TaxProvider *string `form:"tax_provider" json:"tax_provider,omitempty"`
}
type V1ContractsCreateCustomerWithContractCustomerCustomerRevenueSystemConfigurationParams struct {
	// Configuration for the revenue system provider. The structure of this object is specific to the revenue system provider. For NetSuite, this should contain `netsuite_customer_id`.
	Configuration map[string]JSON `form:"configuration" json:"configuration,omitempty"`
	// The method to use for delivering invoices to this customer. If not provided, the `delivery_method_id` must be provided.
	DeliveryMethod *string `form:"delivery_method" json:"delivery_method,omitempty"`
	// ID of the delivery method to use for this customer. If not provided, the `delivery_method` must be provided.
	DeliveryMethodID *string `form:"delivery_method_id" json:"delivery_method_id,omitempty"`
	// The revenue system provider set for this configuration.
	Provider *string `form:"provider" json:"provider"`
}
type V1ContractsCreateCustomerWithContractCustomerParams struct {
	CustomerBillingProviderConfigurations []*V1ContractsCreateCustomerWithContractCustomerCustomerBillingProviderConfigurationParams `form:"customer_billing_provider_configurations" json:"customer_billing_provider_configurations,omitempty"`
	CustomerRevenueSystemConfigurations   []*V1ContractsCreateCustomerWithContractCustomerCustomerRevenueSystemConfigurationParams   `form:"customer_revenue_system_configurations" json:"customer_revenue_system_configurations,omitempty"`
	CustomFields                          map[string]string                                                                          `form:"custom_fields" json:"custom_fields,omitempty"`
	// Aliases that can be used to refer to this customer in usage events
	IngestAliases []*string `form:"ingest_aliases" json:"ingest_aliases,omitempty"`
	// This will be truncated to 160 characters if the provided name is longer.
	Name *string `form:"name" json:"name"`
}
type V1ContractsCreateCustomerWithContractContractCommitAccessScheduleScheduleItemParams struct {
	Amount *Number `form:"amount" json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	// RFC 3339 timestamp (inclusive)
	StartingAt *string `form:"starting_at" json:"starting_at"`
}

// Required: Schedule for distributing the commit to the customer. For "POSTPAID" commits only one schedule item is allowed and amount must match invoice_schedule total.
type V1ContractsCreateCustomerWithContractContractCommitAccessScheduleParams struct {
	// Determines how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used. Defaults to `SPEND` if omitted.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID  *string                                                                                `form:"credit_type_id" json:"credit_type_id,omitempty"`
	ScheduleItems []*V1ContractsCreateCustomerWithContractContractCommitAccessScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items"`
}

// Either provide amount or provide both unit_price and quantity.
type V1ContractsCreateCustomerWithContractContractCommitInvoiceScheduleScheduleItemParams struct {
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
type V1ContractsCreateCustomerWithContractContractCommitInvoiceScheduleRecurringScheduleParams struct {
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

// Required for "POSTPAID" commits: the true up invoice will be generated at this time and only one schedule item is allowed; the total must match access_schedule amount. Optional for "PREPAID" commits: if not provided, this will be a "complimentary" commit with no invoice.
type V1ContractsCreateCustomerWithContractContractCommitInvoiceScheduleParams struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is only applicable to commit invoice schedules. If true, this schedule will not generate an invoice.
	DoNotInvoice *bool `form:"do_not_invoice" json:"do_not_invoice,omitempty"`
	// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
	RecurringSchedule *V1ContractsCreateCustomerWithContractContractCommitInvoiceScheduleRecurringScheduleParams `form:"recurring_schedule" json:"recurring_schedule,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []*V1ContractsCreateCustomerWithContractContractCommitInvoiceScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items,omitempty"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreateCustomerWithContractContractCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreateCustomerWithContractContractCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreateCustomerWithContractContractCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Optional configuration for commit hierarchy access control
type V1ContractsCreateCustomerWithContractContractCommitHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}

// Optional attributes for spend tracker integration. Immutable after creation.
type V1ContractsCreateCustomerWithContractContractCommitSpendTrackerAttributesParams struct {
	// If true, this commit will be included in spend trackers with discounted set to DISCOUNTED_ONLY
	CountsAsDiscounted *bool `form:"counts_as_discounted" json:"counts_as_discounted"`
}
type V1ContractsCreateCustomerWithContractContractCommitParams struct {
	// Required: Schedule for distributing the commit to the customer. For "POSTPAID" commits only one schedule item is allowed and amount must match invoice_schedule total.
	AccessSchedule *V1ContractsCreateCustomerWithContractContractCommitAccessScheduleParams `form:"access_schedule" json:"access_schedule,omitempty"`
	// (DEPRECATED) Use access_schedule and invoice_schedule instead.
	Amount *Number `form:"amount" json:"amount,omitempty"`
	// Which products the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductTags []*string         `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	CustomFields          map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// Used only in UI/API. It is not exposed to end customers.
	Description *string `form:"description" json:"description,omitempty"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration *V1ContractsCreateCustomerWithContractContractCommitHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// Required for "POSTPAID" commits: the true up invoice will be generated at this time and only one schedule item is allowed; the total must match access_schedule amount. Optional for "PREPAID" commits: if not provided, this will be a "complimentary" commit with no invoice.
	InvoiceSchedule *V1ContractsCreateCustomerWithContractContractCommitInvoiceScheduleParams `form:"invoice_schedule" json:"invoice_schedule,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// If multiple commits are applicable, the one with the lower priority will apply first.
	Priority  *Number `form:"priority" json:"priority,omitempty"`
	ProductID *string `form:"product_id" json:"product_id"`
	RateType  *string `form:"rate_type" json:"rate_type,omitempty"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreateCustomerWithContractContractCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
	// Optional attributes for spend tracker integration. Immutable after creation.
	SpendTrackerAttributes *V1ContractsCreateCustomerWithContractContractCommitSpendTrackerAttributesParams `form:"spend_tracker_attributes" json:"spend_tracker_attributes,omitempty"`
	// A temporary ID for the commit that can be used to reference the commit for commit specific overrides.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
	Type        *string `form:"type" json:"type"`
}
type V1ContractsCreateCustomerWithContractContractCreditAccessScheduleScheduleItemParams struct {
	Amount *Number `form:"amount" json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	// RFC 3339 timestamp (inclusive)
	StartingAt *string `form:"starting_at" json:"starting_at"`
}

// Schedule for distributing the credit to the customer.
type V1ContractsCreateCustomerWithContractContractCreditAccessScheduleParams struct {
	// Determines how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used. Defaults to `SPEND` if omitted.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID  *string                                                                                `form:"credit_type_id" json:"credit_type_id,omitempty"`
	ScheduleItems []*V1ContractsCreateCustomerWithContractContractCreditAccessScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreateCustomerWithContractContractCreditSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreateCustomerWithContractContractCreditSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreateCustomerWithContractContractCreditSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Optional configuration for credit hierarchy access control
type V1ContractsCreateCustomerWithContractContractCreditHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}
type V1ContractsCreateCustomerWithContractContractCreditParams struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule *V1ContractsCreateCustomerWithContractContractCreditAccessScheduleParams `form:"access_schedule" json:"access_schedule"`
	// Which products the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductTags []*string         `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	CustomFields          map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// Used only in UI/API. It is not exposed to end customers.
	Description *string `form:"description" json:"description,omitempty"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration *V1ContractsCreateCustomerWithContractContractCreditHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// If multiple credits are applicable, the one with the lower priority will apply first.
	Priority  *Number `form:"priority" json:"priority,omitempty"`
	ProductID *string `form:"product_id" json:"product_id"`
	RateType  *string `form:"rate_type" json:"rate_type,omitempty"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreateCustomerWithContractContractCreditSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
}

// The amount of commit to grant.
type V1ContractsCreateCustomerWithContractContractRecurringCommitAccessAmountParams struct {
	// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is required unless a subscription is attached via `subscription_config`.
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	UnitPrice *Number `form:"unit_price" json:"unit_price"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreateCustomerWithContractContractRecurringCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreateCustomerWithContractContractRecurringCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreateCustomerWithContractContractRecurringCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.
type V1ContractsCreateCustomerWithContractContractRecurringCommitCommitDurationParams struct {
	Unit  *string `form:"unit" json:"unit,omitempty"`
	Value *Number `form:"value" json:"value"`
}
type V1ContractsCreateCustomerWithContractContractRecurringCommitSubscriptionConfigApplySeatIncreaseConfigParams struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated *bool `form:"is_prorated" json:"is_prorated"`
}

// Attach a subscription to the recurring commit/credit.
type V1ContractsCreateCustomerWithContractContractRecurringCommitSubscriptionConfigParams struct {
	// Controls when a customer can access recurring commit or credit child balances when the attached subscription is payment-gated. `BILLING_PERIOD_PAID` releases each balance after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases all balances after the first payment.
	AccessPolicy            *string                                                                                                      `form:"access_policy" json:"access_policy,omitempty"`
	Allocation              *string                                                                                                      `form:"allocation" json:"allocation,omitempty"`
	ApplySeatIncreaseConfig *V1ContractsCreateCustomerWithContractContractRecurringCommitSubscriptionConfigApplySeatIncreaseConfigParams `form:"apply_seat_increase_config" json:"apply_seat_increase_config"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID *string `form:"subscription_id" json:"subscription_id"`
}

// Optional configuration for recurring commit/credit hierarchy access control
type V1ContractsCreateCustomerWithContractContractRecurringCommitHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}

// The amount the customer should be billed for the commit. Not required.
type V1ContractsCreateCustomerWithContractContractRecurringCommitInvoiceAmountParams struct {
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id"`
	Quantity     *Number `form:"quantity" json:"quantity"`
	UnitPrice    *Number `form:"unit_price" json:"unit_price"`
}
type V1ContractsCreateCustomerWithContractContractRecurringCommitProrationRoundingAccessParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}
type V1ContractsCreateCustomerWithContractContractRecurringCommitProrationRoundingInvoiceParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}

// Optional rounding configuration for prorated recurring commit amounts.
type V1ContractsCreateCustomerWithContractContractRecurringCommitProrationRoundingParams struct {
	Access  *V1ContractsCreateCustomerWithContractContractRecurringCommitProrationRoundingAccessParams  `form:"access" json:"access,omitempty"`
	Invoice *V1ContractsCreateCustomerWithContractContractRecurringCommitProrationRoundingInvoiceParams `form:"invoice" json:"invoice,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractRecurringCommitParams struct {
	// The amount of commit to grant.
	AccessAmount *V1ContractsCreateCustomerWithContractContractRecurringCommitAccessAmountParams `form:"access_amount" json:"access_amount"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Will be passed down to the individual commits
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	// Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.
	CommitDuration *V1ContractsCreateCustomerWithContractContractRecurringCommitCommitDurationParams `form:"commit_duration" json:"commit_duration"`
	// Will be passed down to the individual commits
	Description *string `form:"description" json:"description,omitempty"`
	// Determines when the contract will stop creating recurring commits. optional
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// Optional configuration for recurring commit/credit hierarchy access control
	HierarchyConfiguration *V1ContractsCreateCustomerWithContractContractRecurringCommitHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount *V1ContractsCreateCustomerWithContractContractRecurringCommitInvoiceAmountParams `form:"invoice_amount" json:"invoice_amount,omitempty"`
	// displayed on invoices. will be passed through to the individual commits
	Name *string `form:"name" json:"name,omitempty"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// Will be passed down to the individual commits
	Priority  *Number `form:"priority" json:"priority"`
	ProductID *string `form:"product_id" json:"product_id"`
	// Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration *string `form:"proration" json:"proration,omitempty"`
	// Optional rounding configuration for prorated recurring commit amounts.
	ProrationRounding *V1ContractsCreateCustomerWithContractContractRecurringCommitProrationRoundingParams `form:"proration_rounding" json:"proration_rounding,omitempty"`
	// Whether the created commits will use the commit rate or list rate
	RateType *string `form:"rate_type" json:"rate_type,omitempty"`
	// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
	RecurrenceFrequency *string `form:"recurrence_frequency" json:"recurrence_frequency,omitempty"`
	// Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreateCustomerWithContractContractRecurringCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
	// determines the start time for the first commit
	StartingAt *string `form:"starting_at" json:"starting_at"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig *V1ContractsCreateCustomerWithContractContractRecurringCommitSubscriptionConfigParams `form:"subscription_config" json:"subscription_config,omitempty"`
	// A temporary ID that can be used to reference the recurring commit for commit specific overrides.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
}

// The amount of commit to grant.
type V1ContractsCreateCustomerWithContractContractRecurringCreditAccessAmountParams struct {
	// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is required unless a subscription is attached via `subscription_config`.
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	UnitPrice *Number `form:"unit_price" json:"unit_price"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreateCustomerWithContractContractRecurringCreditSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreateCustomerWithContractContractRecurringCreditSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreateCustomerWithContractContractRecurringCreditSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.
type V1ContractsCreateCustomerWithContractContractRecurringCreditCommitDurationParams struct {
	Unit  *string `form:"unit" json:"unit,omitempty"`
	Value *Number `form:"value" json:"value"`
}
type V1ContractsCreateCustomerWithContractContractRecurringCreditSubscriptionConfigApplySeatIncreaseConfigParams struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated *bool `form:"is_prorated" json:"is_prorated"`
}

// Attach a subscription to the recurring commit/credit.
type V1ContractsCreateCustomerWithContractContractRecurringCreditSubscriptionConfigParams struct {
	// Controls when a customer can access recurring commit or credit child balances when the attached subscription is payment-gated. `BILLING_PERIOD_PAID` releases each balance after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases all balances after the first payment.
	AccessPolicy            *string                                                                                                      `form:"access_policy" json:"access_policy,omitempty"`
	Allocation              *string                                                                                                      `form:"allocation" json:"allocation,omitempty"`
	ApplySeatIncreaseConfig *V1ContractsCreateCustomerWithContractContractRecurringCreditSubscriptionConfigApplySeatIncreaseConfigParams `form:"apply_seat_increase_config" json:"apply_seat_increase_config"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID *string `form:"subscription_id" json:"subscription_id"`
}

// Optional configuration for recurring commit/credit hierarchy access control
type V1ContractsCreateCustomerWithContractContractRecurringCreditHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}
type V1ContractsCreateCustomerWithContractContractRecurringCreditProrationRoundingAccessParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}

// Optional rounding configuration for prorated recurring credit amounts.
type V1ContractsCreateCustomerWithContractContractRecurringCreditProrationRoundingParams struct {
	Access *V1ContractsCreateCustomerWithContractContractRecurringCreditProrationRoundingAccessParams `form:"access" json:"access,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractRecurringCreditParams struct {
	// The amount of commit to grant.
	AccessAmount *V1ContractsCreateCustomerWithContractContractRecurringCreditAccessAmountParams `form:"access_amount" json:"access_amount"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Will be passed down to the individual commits
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	// Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.
	CommitDuration *V1ContractsCreateCustomerWithContractContractRecurringCreditCommitDurationParams `form:"commit_duration" json:"commit_duration"`
	// Will be passed down to the individual commits
	Description *string `form:"description" json:"description,omitempty"`
	// Determines when the contract will stop creating recurring commits. optional
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// Optional configuration for recurring commit/credit hierarchy access control
	HierarchyConfiguration *V1ContractsCreateCustomerWithContractContractRecurringCreditHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// displayed on invoices. will be passed through to the individual commits
	Name *string `form:"name" json:"name,omitempty"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// Will be passed down to the individual commits
	Priority  *Number `form:"priority" json:"priority"`
	ProductID *string `form:"product_id" json:"product_id"`
	// Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration *string `form:"proration" json:"proration,omitempty"`
	// Optional rounding configuration for prorated recurring credit amounts.
	ProrationRounding *V1ContractsCreateCustomerWithContractContractRecurringCreditProrationRoundingParams `form:"proration_rounding" json:"proration_rounding,omitempty"`
	// Whether the created commits will use the commit rate or list rate
	RateType *string `form:"rate_type" json:"rate_type,omitempty"`
	// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
	RecurrenceFrequency *string `form:"recurrence_frequency" json:"recurrence_frequency,omitempty"`
	// Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreateCustomerWithContractContractRecurringCreditSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
	// determines the start time for the first commit
	StartingAt *string `form:"starting_at" json:"starting_at"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig *V1ContractsCreateCustomerWithContractContractRecurringCreditSubscriptionConfigParams `form:"subscription_config" json:"subscription_config,omitempty"`
	// A temporary ID that can be used to reference the recurring commit for commit specific overrides.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
}

// Only set for TIERED rate_type.
type V1ContractsCreateCustomerWithContractContractOverrideOverwriteRateTierParams struct {
	Price *Number `form:"price" json:"price"`
	Size  *Number `form:"size" json:"size,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractOverrideOverwriteRateMinimumConfigParams struct {
	Minimum *Number `form:"minimum" json:"minimum"`
}

// Required for OVERWRITE type.
type V1ContractsCreateCustomerWithContractContractOverrideOverwriteRateParams struct {
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate processors.
	CustomRate map[string]JSON `form:"custom_rate" json:"custom_rate,omitempty"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be set to true.
	IsProrated    *bool                                                                                  `form:"is_prorated" json:"is_prorated,omitempty"`
	MinimumConfig *V1ContractsCreateCustomerWithContractContractOverrideOverwriteRateMinimumConfigParams `form:"minimum_config" json:"minimum_config,omitempty"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price *Number `form:"price" json:"price,omitempty"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity *Number `form:"quantity" json:"quantity,omitempty"`
	RateType *string `form:"rate_type" json:"rate_type"`
	// Only set for TIERED rate_type.
	Tiers []*V1ContractsCreateCustomerWithContractContractOverrideOverwriteRateTierParams `form:"tiers" json:"tiers,omitempty"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreateCustomerWithContractContractOverrideOverrideSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Cannot be used in conjunction with product_id or applicable_product_tags. If provided, the override will apply to all products with the specified specifiers.
type V1ContractsCreateCustomerWithContractContractOverrideOverrideSpecifierParams struct {
	// Can only be used for commit specific overrides. Must be used in conjunction with one of `product_id`, `product_tags`, `pricing_group_values`, or `presentation_group_values`. Must be used instead of both `commit_ids` and `recurring_commit_ids` If provided, the override will apply to any specified commit, credit, recurring commit or recurring credit IDs.
	AnyCommitOrCreditIDs []*string `form:"any_commit_or_credit_ids" json:"any_commit_or_credit_ids,omitempty"`
	BillingFrequency     *string   `form:"billing_frequency" json:"billing_frequency,omitempty"`
	// Can only be used for commit specific overrides. Must be used in conjunction with one of `product_id`, `product_tags`, `pricing_group_values`, or `presentation_group_values`. If provided, the override will only apply to the specified commits. If not provided, the override will apply to all commits.
	CommitIDs []*string `form:"commit_ids" json:"commit_ids,omitempty"`
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreateCustomerWithContractContractOverrideOverrideSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// A map of group names to values. The override will only apply to line items with the specified presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// A map of pricing group names to values. The override will only apply to products with the specified pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the override will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the override will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
	// Can only be used for commit specific overrides. Must be used in conjunction with one of `product_id`, `product_tags`, `pricing_group_values`, or `presentation_group_values`. If provided, the override will only apply to commits created by the specified recurring commit ids.
	RecurringCommitIDs []*string `form:"recurring_commit_ids" json:"recurring_commit_ids,omitempty"`
}

// Required for TIERED type. Must have at least one tier.
type V1ContractsCreateCustomerWithContractContractOverrideTierParams struct {
	Multiplier *Number `form:"multiplier" json:"multiplier"`
	Size       *Number `form:"size" json:"size,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractOverrideParams struct {
	// tags identifying products whose rates are being overridden. Cannot be used in conjunction with override_specifiers.
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	Entitled     *bool   `form:"entitled" json:"entitled,omitempty"`
	// Indicates whether the override should only apply to commits. Defaults to `false`. If `true` you can specify relevant commits in `override_specifiers` by passing `commit_ids`, `recurring_commit_ids`, or `any_commit_or_credit_ids`.  If you do not specify any of these fields, the override will apply when consuming any prepaid commit, postpaid commit, or credit
	IsCommitSpecific *bool `form:"is_commit_specific" json:"is_commit_specific,omitempty"`
	// Required for MULTIPLIER type. Must be >=0.
	Multiplier *Number `form:"multiplier" json:"multiplier,omitempty"`
	// Cannot be used in conjunction with product_id or applicable_product_tags. If provided, the override will apply to all products with the specified specifiers.
	OverrideSpecifiers []*V1ContractsCreateCustomerWithContractContractOverrideOverrideSpecifierParams `form:"override_specifiers" json:"override_specifiers,omitempty"`
	// Required for OVERWRITE type.
	OverwriteRate *V1ContractsCreateCustomerWithContractContractOverrideOverwriteRateParams `form:"overwrite_rate" json:"overwrite_rate,omitempty"`
	// Required for EXPLICIT multiplier prioritization scheme and all TIERED overrides. Under EXPLICIT prioritization, overwrites are prioritized first, and then tiered and multiplier overrides are prioritized by their priority value (lowest first). Must be > 0.
	Priority *Number `form:"priority" json:"priority,omitempty"`
	// ID of the product whose rate is being overridden. Cannot be used in conjunction with override_specifiers.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// RFC 3339 timestamp indicating when the override will start applying (inclusive)
	StartingAt *string `form:"starting_at" json:"starting_at"`
	// Indicates whether the override applies to commit rates or list rates. Can only be used for overrides that have `is_commit_specific` set to `true`. Defaults to `"LIST_RATE"`.
	Target *string `form:"target" json:"target,omitempty"`
	// Required for TIERED type. Must have at least one tier.
	Tiers []*V1ContractsCreateCustomerWithContractContractOverrideTierParams `form:"tiers" json:"tiers,omitempty"`
	// Overwrites are prioritized over multipliers and tiered overrides.
	Type *string `form:"type" json:"type,omitempty"`
}

// Either provide amount or provide both unit_price and quantity.
type V1ContractsCreateCustomerWithContractContractDiscountScheduleScheduleItemParams struct {
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
type V1ContractsCreateCustomerWithContractContractDiscountScheduleRecurringScheduleParams struct {
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
type V1ContractsCreateCustomerWithContractContractDiscountScheduleParams struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is only applicable to commit invoice schedules. If true, this schedule will not generate an invoice.
	DoNotInvoice *bool `form:"do_not_invoice" json:"do_not_invoice,omitempty"`
	// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
	RecurringSchedule *V1ContractsCreateCustomerWithContractContractDiscountScheduleRecurringScheduleParams `form:"recurring_schedule" json:"recurring_schedule,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []*V1ContractsCreateCustomerWithContractContractDiscountScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items,omitempty"`
}

// This field's availability is dependent on your client's configuration.
type V1ContractsCreateCustomerWithContractContractDiscountParams struct {
	CustomFields map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string                                                              `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	ProductID            *string                                                              `form:"product_id" json:"product_id"`
	Schedule             *V1ContractsCreateCustomerWithContractContractDiscountScheduleParams `form:"schedule" json:"schedule"`
}
type V1ContractsCreateCustomerWithContractContractResellerRoyaltyAwsOptionsParams struct {
	AwsAccountNumber    *string `form:"aws_account_number" json:"aws_account_number,omitempty"`
	AwsOfferID          *string `form:"aws_offer_id" json:"aws_offer_id,omitempty"`
	AwsPayerReferenceID *string `form:"aws_payer_reference_id" json:"aws_payer_reference_id,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractResellerRoyaltyGcpOptionsParams struct {
	GcpAccountID *string `form:"gcp_account_id" json:"gcp_account_id,omitempty"`
	GcpOfferID   *string `form:"gcp_offer_id" json:"gcp_offer_id,omitempty"`
}

// This field's availability is dependent on your client's configuration.
type V1ContractsCreateCustomerWithContractContractResellerRoyaltyParams struct {
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductTags []*string                                                                     `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	AwsOptions            *V1ContractsCreateCustomerWithContractContractResellerRoyaltyAwsOptionsParams `form:"aws_options" json:"aws_options,omitempty"`
	EndingBefore          *string                                                                       `form:"ending_before" json:"ending_before,omitempty"`
	Fraction              *Number                                                                       `form:"fraction" json:"fraction"`
	GcpOptions            *V1ContractsCreateCustomerWithContractContractResellerRoyaltyGcpOptionsParams `form:"gcp_options" json:"gcp_options,omitempty"`
	NetsuiteResellerID    *string                                                                       `form:"netsuite_reseller_id" json:"netsuite_reseller_id"`
	ResellerContractValue *Number                                                                       `form:"reseller_contract_value" json:"reseller_contract_value,omitempty"`
	ResellerType          *string                                                                       `form:"reseller_type" json:"reseller_type"`
	StartingAt            *string                                                                       `form:"starting_at" json:"starting_at"`
}

// This field's availability is dependent on your client's configuration.
type V1ContractsCreateCustomerWithContractContractProfessionalServiceParams struct {
	CustomFields map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	Description  *string           `form:"description" json:"description,omitempty"`
	// Maximum amount for the term.
	MaxAmount *Number `form:"max_amount" json:"max_amount"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	ProductID            *string `form:"product_id" json:"product_id"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the amount.
	Quantity *Number `form:"quantity" json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified.
	UnitPrice *Number `form:"unit_price" json:"unit_price"`
}

// Either provide amount or provide both unit_price and quantity.
type V1ContractsCreateCustomerWithContractContractScheduledChargeScheduleScheduleItemParams struct {
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
type V1ContractsCreateCustomerWithContractContractScheduledChargeScheduleRecurringScheduleParams struct {
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
type V1ContractsCreateCustomerWithContractContractScheduledChargeScheduleParams struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is only applicable to commit invoice schedules. If true, this schedule will not generate an invoice.
	DoNotInvoice *bool `form:"do_not_invoice" json:"do_not_invoice,omitempty"`
	// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
	RecurringSchedule *V1ContractsCreateCustomerWithContractContractScheduledChargeScheduleRecurringScheduleParams `form:"recurring_schedule" json:"recurring_schedule,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []*V1ContractsCreateCustomerWithContractContractScheduledChargeScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractScheduledChargeParams struct {
	CustomFields map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string                                                                     `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	ProductID            *string                                                                     `form:"product_id" json:"product_id"`
	Schedule             *V1ContractsCreateCustomerWithContractContractScheduledChargeScheduleParams `form:"schedule" json:"schedule"`
}
type V1ContractsCreateCustomerWithContractContractUsageFilterParams struct {
	GroupKey    *string   `form:"group_key" json:"group_key"`
	GroupValues []*string `form:"group_values" json:"group_values"`
	StartingAt  *string   `form:"starting_at" json:"starting_at,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractUsageStatementScheduleParams struct {
	// Required when using CUSTOM_DATE. This option lets you set a historical billing anchor date, aligning future billing cycles with a chosen cadence. For example, if a contract starts on 2024-09-15 and you set the anchor date to 2024-09-10 with a MONTHLY frequency, the first usage statement will cover 09-15 to 10-10. Subsequent statements will follow the 10th of each month.
	BillingAnchorDate *string `form:"billing_anchor_date" json:"billing_anchor_date,omitempty"`
	// If not provided, defaults to the first day of the month.
	Day       *string `form:"day" json:"day,omitempty"`
	Frequency *string `form:"frequency" json:"frequency"`
	// The date Metronome should start generating usage invoices. If unspecified, contract start date will be used. This is useful to set if you want to import historical invoices via our 'Create Historical Invoices' API rather than having Metronome automatically generate them.
	InvoiceGenerationStartingAt *string `form:"invoice_generation_starting_at" json:"invoice_generation_starting_at,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractBillingProviderConfigurationParams struct {
	// Do not specify if using billing_provider_configuration_id.
	BillingProvider *string `form:"billing_provider" json:"billing_provider,omitempty"`
	// The Metronome ID of the billing provider configuration. Use when a customer has multiple configurations with the same billing provider and delivery method. Otherwise, specify the billing_provider and delivery_method.
	BillingProviderConfigurationID *string `form:"billing_provider_configuration_id" json:"billing_provider_configuration_id,omitempty"`
	// Do not specify if using billing_provider_configuration_id.
	DeliveryMethod *string `form:"delivery_method" json:"delivery_method,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractRevenueSystemConfigurationParams struct {
	// How revenue recognition records should be delivered to the revenue system. Do not specify if using revenue_system_configuration_id.
	DeliveryMethod *string `form:"delivery_method" json:"delivery_method,omitempty"`
	// The system that is providing services for revenue recognition. Do not specify if using revenue_system_configuration_id.
	Provider *string `form:"provider" json:"provider,omitempty"`
	// The Metronome ID of the revenue system configuration. Use when a customer has multiple configurations with the same provider and delivery method. Otherwise, specify the provider and delivery_method.
	RevenueSystemConfigurationID *string `form:"revenue_system_configuration_id" json:"revenue_system_configuration_id,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractSpendThresholdConfigurationCommitParams struct {
	Description *string `form:"description" json:"description,omitempty"`
	// Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.
	Name *string `form:"name" json:"name,omitempty"`
	// The priority of the commit, used to determine drawdown order. Lower priority commits are consumed first. Defaults to 100 if not specified.
	Priority *Number `form:"priority" json:"priority,omitempty"`
	// The commit product that will be used to generate the line item for commit payment.
	ProductID *string `form:"product_id" json:"product_id"`
}

// Only applicable if using STRIPE as your payment gate type.
type V1ContractsCreateCustomerWithContractContractSpendThresholdConfigurationPaymentGateConfigStripeConfigParams struct {
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as your payment type.
	InvoiceMetadata map[string]string `form:"invoice_metadata" json:"invoice_metadata,omitempty"`
	// If left blank, will default to INVOICE
	PaymentType *string `form:"payment_type" json:"payment_type"`
}

// Only applicable if using PRECALCULATED as your tax type.
type V1ContractsCreateCustomerWithContractContractSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams struct {
	// Amount of tax to be applied. This should be in the same currency and denomination  as the commit's invoice schedule
	TaxAmount *Number `form:"tax_amount" json:"tax_amount"`
	// Name of the tax to be applied. This may be used in an invoice line item description.
	TaxName *string `form:"tax_name" json:"tax_name,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractSpendThresholdConfigurationPaymentGateConfigParams struct {
	// Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.
	PaymentGateType *string `form:"payment_gate_type" json:"payment_gate_type"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig *V1ContractsCreateCustomerWithContractContractSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams `form:"precalculated_tax_config" json:"precalculated_tax_config,omitempty"`
	// Only applicable if using STRIPE as your payment gate type.
	StripeConfig *V1ContractsCreateCustomerWithContractContractSpendThresholdConfigurationPaymentGateConfigStripeConfigParams `form:"stripe_config" json:"stripe_config,omitempty"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType *string `form:"tax_type" json:"tax_type,omitempty"`
}

// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
type V1ContractsCreateCustomerWithContractContractSpendThresholdConfigurationDiscountConfigurationCapParams struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount *Number `form:"amount" json:"amount"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias *string `form:"spend_tracker_alias" json:"spend_tracker_alias"`
}
type V1ContractsCreateCustomerWithContractContractSpendThresholdConfigurationDiscountConfigurationParams struct {
	// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
	Cap *V1ContractsCreateCustomerWithContractContractSpendThresholdConfigurationDiscountConfigurationCapParams `form:"cap" json:"cap,omitempty"`
	// The fraction of the original amount that the customer pays after applying the discount. For example, 0.85 means the customer pays 85% of the original amount (a 15% discount).
	PaymentFraction *Number `form:"payment_fraction" json:"payment_fraction"`
}
type V1ContractsCreateCustomerWithContractContractSpendThresholdConfigurationParams struct {
	Commit                *V1ContractsCreateCustomerWithContractContractSpendThresholdConfigurationCommitParams                `form:"commit" json:"commit"`
	DiscountConfiguration *V1ContractsCreateCustomerWithContractContractSpendThresholdConfigurationDiscountConfigurationParams `form:"discount_configuration" json:"discount_configuration,omitempty"`
	// When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.
	IsEnabled         *bool                                                                                            `form:"is_enabled" json:"is_enabled"`
	PaymentGateConfig *V1ContractsCreateCustomerWithContractContractSpendThresholdConfigurationPaymentGateConfigParams `form:"payment_gate_config" json:"payment_gate_config"`
	// Specify the threshold amount for the contract. Each time the contract's usage hits this amount, a threshold charge will be initiated.
	ThresholdAmount *Number `form:"threshold_amount" json:"threshold_amount"`
}

// The length of time the created commit will be valid, starting from the end of the invoice's service period. If not provided, defaults to one year.
type V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationCommitDurationParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationCommitParams struct {
	// Which products the threshold commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the threshold commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	Description           *string   `form:"description" json:"description,omitempty"`
	// The length of time the created commit will be valid, starting from the end of the invoice's service period. If not provided, defaults to one year.
	Duration *V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationCommitDurationParams `form:"duration" json:"duration,omitempty"`
	// Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.
	Name *string `form:"name" json:"name,omitempty"`
	// The priority of the commit, used to determine drawdown order. Lower priority commits are consumed first. Defaults to 100 if not specified.
	Priority *Number `form:"priority" json:"priority,omitempty"`
	// The commit product that will be used to generate the line item for commit payment.
	ProductID *string `form:"product_id" json:"product_id"`
	// Whether the created commits will be charged at commit rate or list rate.
	RateType *string `form:"rate_type" json:"rate_type,omitempty"`
	// Fraction of the created commit's unused balance that will roll over. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
}

// Only applicable if using STRIPE as your payment gate type.
type V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigParams struct {
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as your payment type.
	InvoiceMetadata map[string]string `form:"invoice_metadata" json:"invoice_metadata,omitempty"`
	// If left blank, will default to INVOICE
	PaymentType *string `form:"payment_type" json:"payment_type"`
}

// Only applicable if using PRECALCULATED as your tax type.
type V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams struct {
	// Amount of tax to be applied. This should be in the same currency and denomination  as the commit's invoice schedule
	TaxAmount *Number `form:"tax_amount" json:"tax_amount"`
	// Name of the tax to be applied. This may be used in an invoice line item description.
	TaxName *string `form:"tax_name" json:"tax_name,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationPaymentGateConfigParams struct {
	// Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.
	PaymentGateType *string `form:"payment_gate_type" json:"payment_gate_type"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig *V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams `form:"precalculated_tax_config" json:"precalculated_tax_config,omitempty"`
	// Only applicable if using STRIPE as your payment gate type.
	StripeConfig *V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigParams `form:"stripe_config" json:"stripe_config,omitempty"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType *string `form:"tax_type" json:"tax_type,omitempty"`
}

// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
type V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationDiscountConfigurationCapParams struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount *Number `form:"amount" json:"amount"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias *string `form:"spend_tracker_alias" json:"spend_tracker_alias"`
}
type V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationDiscountConfigurationParams struct {
	// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
	Cap *V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationDiscountConfigurationCapParams `form:"cap" json:"cap,omitempty"`
	// The fraction of the original amount that the customer pays after applying the discount. For example, 0.85 means the customer pays 85% of the original amount (a 15% discount).
	PaymentFraction *Number `form:"payment_fraction" json:"payment_fraction"`
}

// If provided, balances with all the custom fields will not be considered when evaluating threshold billing
type V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilterParams struct {
	Entity *string `form:"entity" json:"entity"`
	Key    *string `form:"key" json:"key"`
	Value  *string `form:"value" json:"value"`
}

// If any of the exclude specifier is met, the balance is not considered when evaluating threshold billing
type V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeParams struct {
	// If provided, balances with all the custom fields will not be considered when evaluating threshold billing
	CustomFieldFilters []*V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilterParams `form:"custom_field_filters" json:"custom_field_filters"`
}

// Determines which balances are excluded from remaining balance calculation for threshold billing.
type V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierParams struct {
	// If any of the exclude specifier is met, the balance is not considered when evaluating threshold billing
	Exclude []*V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeParams `form:"exclude" json:"exclude"`
}
type V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationParams struct {
	Commit *V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationCommitParams `form:"commit" json:"commit"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID    *string                                                                                                       `form:"custom_credit_type_id" json:"custom_credit_type_id,omitempty"`
	DiscountConfiguration *V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationDiscountConfigurationParams `form:"discount_configuration" json:"discount_configuration,omitempty"`
	// When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.
	IsEnabled         *bool                                                                                                     `form:"is_enabled" json:"is_enabled"`
	PaymentGateConfig *V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationPaymentGateConfigParams `form:"payment_gate_config" json:"payment_gate_config"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount *Number `form:"recharge_to_amount" json:"recharge_to_amount"`
	// Specify the threshold amount for the contract. Each time the contract's prepaid balance lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount *Number `form:"threshold_amount" json:"threshold_amount"`
	// Determines which balances are excluded from remaining balance calculation for threshold billing.
	ThresholdBalanceSpecifiers []*V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierParams `form:"threshold_balance_specifiers" json:"threshold_balance_specifiers,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractSpendTrackerApplicableSpendSpecifierParams struct {
	// Filter by whether the spend was discounted. Defaults to ANY if omitted.
	Discounted *string   `form:"discounted" json:"discounted,omitempty"`
	Sources    []*string `form:"sources" json:"sources"`
	SpendType  *string   `form:"spend_type" json:"spend_type"`
}

// Spend trackers to attach to this contract. Aliases must be unique within a contract.
type V1ContractsCreateCustomerWithContractContractSpendTrackerParams struct {
	// Human-readable identifier, unique per contract.
	Alias                     *string                                                                                    `form:"alias" json:"alias"`
	ApplicableSpendSpecifiers []*V1ContractsCreateCustomerWithContractContractSpendTrackerApplicableSpendSpecifierParams `form:"applicable_spend_specifiers" json:"applicable_spend_specifiers"`
	CreditTypeID              *string                                                                                    `form:"credit_type_id" json:"credit_type_id"`
	ResetFrequency            *string                                                                                    `form:"reset_frequency" json:"reset_frequency"`
}
type V1ContractsCreateCustomerWithContractContractSubscriptionSubscriptionRateParams struct {
	// Frequency to bill subscription with. Together with product_id, must match existing rate on the rate card.
	BillingFrequency *string `form:"billing_frequency" json:"billing_frequency"`
	// Must be subscription type product
	ProductID *string `form:"product_id" json:"product_id"`
}
type V1ContractsCreateCustomerWithContractContractSubscriptionProrationRoundingParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}
type V1ContractsCreateCustomerWithContractContractSubscriptionProrationParams struct {
	// Indicates how mid-period quantity adjustments are invoiced.
	// **BILL_IMMEDIATELY**: Only available when collection schedule is `ADVANCE`. The quantity increase will be billed immediately on the scheduled date.
	// **BILL_ON_NEXT_COLLECTION_DATE**: The quantity increase will be billed for in-arrears at the end of the period.
	//
	InvoiceBehavior *string `form:"invoice_behavior" json:"invoice_behavior,omitempty"`
	// Indicates if the partial period will be prorated or charged a full amount.
	IsProrated *bool                                                                             `form:"is_prorated" json:"is_prorated,omitempty"`
	Rounding   *V1ContractsCreateCustomerWithContractContractSubscriptionProrationRoundingParams `form:"rounding" json:"rounding,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractSubscriptionSeatConfigParams struct {
	// The initial assigned seats on this subscription.
	InitialSeatIDs []*string `form:"initial_seat_ids" json:"initial_seat_ids"`
	// The initial amount of unassigned seats on this subscription.
	InitialUnassignedSeats *Number `form:"initial_unassigned_seats" json:"initial_unassigned_seats,omitempty"`
	// The property name, sent on usage events, that identifies the seat ID associated with the usage event.  For example, the property name might be seat_id or user_id. The property must be set as a group key on billable metrics and a presentation/pricing group key on contract products.  This allows linked recurring credits with an allocation per seat to be consumed by only one seat's usage.
	SeatGroupKey *string `form:"seat_group_key" json:"seat_group_key"`
}
type V1ContractsCreateCustomerWithContractContractSubscriptionBillingCycleConfigParams struct {
	// The date to anchor the billing cycle to. If omitted, defaults to the contract's usage invoice billing cycle anchor date.
	AnchorDate *string `form:"anchor_date" json:"anchor_date,omitempty"`
	// Controls whether this subscription consolidates onto usage invoices or gets its own scheduled invoice. Defaults to ON_USAGE_INVOICE if omitted.
	InvoicePlacement *string `form:"invoice_placement" json:"invoice_placement,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractSubscriptionPaymentGateConfigParams struct {
	// Gate access to the subscription based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe.
	PaymentGateType *string `form:"payment_gate_type" json:"payment_gate_type"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType *string `form:"tax_type" json:"tax_type,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractSubscriptionParams struct {
	BillingCycleConfig *V1ContractsCreateCustomerWithContractContractSubscriptionBillingCycleConfigParams `form:"billing_cycle_config" json:"billing_cycle_config,omitempty"`
	CollectionSchedule *string                                                                            `form:"collection_schedule" json:"collection_schedule"`
	CustomFields       map[string]string                                                                  `form:"custom_fields" json:"custom_fields,omitempty"`
	Description        *string                                                                            `form:"description" json:"description,omitempty"`
	// Exclusive end time for the subscription. If not provided, subscription inherits contract end date.
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// The initial quantity for the subscription. It must be non-negative value. Required if quantity_management_mode is QUANTITY_ONLY.
	InitialQuantity   *Number                                                                           `form:"initial_quantity" json:"initial_quantity,omitempty"`
	Name              *string                                                                           `form:"name" json:"name,omitempty"`
	PaymentGateConfig *V1ContractsCreateCustomerWithContractContractSubscriptionPaymentGateConfigParams `form:"payment_gate_config" json:"payment_gate_config,omitempty"`
	Proration         *V1ContractsCreateCustomerWithContractContractSubscriptionProrationParams         `form:"proration" json:"proration"`
	// Determines how the subscription's quantity is controlled. Defaults to QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified directly on the subscription. `initial_quantity` must be provided with this option. Compatible with recurring commits/credits that use POOLED allocation. **SEAT_BASED**: Use when you want to pass specific seat identifiers (e.g. add user_123) to increment and decrement a subscription quantity, rather than directly providing the quantity. You must use a **SEAT_BASED** subscription to use a linked recurring credit with an allocation per seat. `seat_config` must be provided with this option.
	QuantityManagementMode *string                                                                    `form:"quantity_management_mode" json:"quantity_management_mode,omitempty"`
	SeatConfig             *V1ContractsCreateCustomerWithContractContractSubscriptionSeatConfigParams `form:"seat_config" json:"seat_config,omitempty"`
	// Inclusive start time for the subscription. If not provided, defaults to contract start date
	StartingAt       *string                                                                          `form:"starting_at" json:"starting_at,omitempty"`
	SubscriptionRate *V1ContractsCreateCustomerWithContractContractSubscriptionSubscriptionRateParams `form:"subscription_rate" json:"subscription_rate"`
	// A temporary ID used to reference the subscription in recurring commit/credit subscription configs created within the same payload.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractHierarchyConfigurationParentParams struct {
	ContractID *string `form:"contract_id" json:"contract_id"`
	CustomerID *string `form:"customer_id" json:"customer_id"`
}
type V1ContractsCreateCustomerWithContractContractHierarchyConfigurationParentBehaviorParams struct {
	// Indicates the desired behavior of consolidated invoices generated by the parent in a customer hierarchy
	//
	// **CONCATENATE**: Statements on the invoices of child customers will be appended to the consolidated invoice
	//
	// **NONE**: Do not generate consolidated invoices
	InvoiceConsolidationType *string `form:"invoice_consolidation_type" json:"invoice_consolidation_type,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractHierarchyConfigurationParams struct {
	Parent         *V1ContractsCreateCustomerWithContractContractHierarchyConfigurationParentParams         `form:"parent" json:"parent,omitempty"`
	ParentBehavior *V1ContractsCreateCustomerWithContractContractHierarchyConfigurationParentBehaviorParams `form:"parent_behavior" json:"parent_behavior,omitempty"`
	// Indicates which customer should pay for the child's invoice charges
	//
	// **SELF**: The child pays for its own invoice charges
	//
	// **PARENT**: The parent pays for the child's invoice charges
	Payer *string `form:"payer" json:"payer,omitempty"`
	// Indicates the behavior of the child's invoice statements on the parent's invoices.
	//
	// **CONSOLIDATE**: Child's invoice statements will be added to parent's consolidated invoices
	//
	// **SEPARATE**: Child's invoice statements will appear not appear on parent's consolidated invoices
	UsageStatementBehavior *string `form:"usage_statement_behavior" json:"usage_statement_behavior,omitempty"`
}
type V1ContractsCreateCustomerWithContractContractParams struct {
	BillingProviderConfiguration *V1ContractsCreateCustomerWithContractContractBillingProviderConfigurationParams `form:"billing_provider_configuration" json:"billing_provider_configuration,omitempty"`
	Commits                      []*V1ContractsCreateCustomerWithContractContractCommitParams                     `form:"commits" json:"commits,omitempty"`
	Credits                      []*V1ContractsCreateCustomerWithContractContractCreditParams                     `form:"credits" json:"credits,omitempty"`
	CustomFields                 map[string]string                                                                `form:"custom_fields" json:"custom_fields,omitempty"`
	// This field's availability is dependent on your client's configuration.
	Discounts []*V1ContractsCreateCustomerWithContractContractDiscountParams `form:"discounts" json:"discounts,omitempty"`
	// exclusive contract end time
	EndingBefore           *string                                                                    `form:"ending_before" json:"ending_before,omitempty"`
	HierarchyConfiguration *V1ContractsCreateCustomerWithContractContractHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list prices automatically. EXPLICIT prioritization requires specifying priorities for each multiplier; the one with the lowest priority value will be prioritized first.
	MultiplierOverridePrioritization *string `form:"multiplier_override_prioritization" json:"multiplier_override_prioritization,omitempty"`
	Name                             *string `form:"name" json:"name,omitempty"`
	NetPaymentTermsDays              *Number `form:"net_payment_terms_days" json:"net_payment_terms_days,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID                 *string                                                                                  `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	Overrides                            []*V1ContractsCreateCustomerWithContractContractOverrideParams                           `form:"overrides" json:"overrides,omitempty"`
	PrepaidBalanceThresholdConfiguration *V1ContractsCreateCustomerWithContractContractPrepaidBalanceThresholdConfigurationParams `form:"prepaid_balance_threshold_configuration" json:"prepaid_balance_threshold_configuration,omitempty"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []*V1ContractsCreateCustomerWithContractContractProfessionalServiceParams `form:"professional_services" json:"professional_services,omitempty"`
	// Selects the rate card linked to the specified alias as of the contract's start date.
	RateCardAlias    *string                                                               `form:"rate_card_alias" json:"rate_card_alias,omitempty"`
	RateCardID       *string                                                               `form:"rate_card_id" json:"rate_card_id,omitempty"`
	RecurringCommits []*V1ContractsCreateCustomerWithContractContractRecurringCommitParams `form:"recurring_commits" json:"recurring_commits,omitempty"`
	RecurringCredits []*V1ContractsCreateCustomerWithContractContractRecurringCreditParams `form:"recurring_credits" json:"recurring_credits,omitempty"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties          []*V1ContractsCreateCustomerWithContractContractResellerRoyaltyParams          `form:"reseller_royalties" json:"reseller_royalties,omitempty"`
	RevenueSystemConfiguration *V1ContractsCreateCustomerWithContractContractRevenueSystemConfigurationParams `form:"revenue_system_configuration" json:"revenue_system_configuration,omitempty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID         *string                                                                         `form:"salesforce_opportunity_id" json:"salesforce_opportunity_id,omitempty"`
	ScheduledCharges                []*V1ContractsCreateCustomerWithContractContractScheduledChargeParams           `form:"scheduled_charges" json:"scheduled_charges,omitempty"`
	ScheduledChargesOnUsageInvoices *string                                                                         `form:"scheduled_charges_on_usage_invoices" json:"scheduled_charges_on_usage_invoices,omitempty"`
	SpendThresholdConfiguration     *V1ContractsCreateCustomerWithContractContractSpendThresholdConfigurationParams `form:"spend_threshold_configuration" json:"spend_threshold_configuration,omitempty"`
	// Spend trackers to attach to this contract. Aliases must be unique within a contract.
	SpendTrackers []*V1ContractsCreateCustomerWithContractContractSpendTrackerParams `form:"spend_trackers" json:"spend_trackers,omitempty"`
	// inclusive contract start time
	StartingAt    *string                                                            `form:"starting_at" json:"starting_at"`
	Subscriptions []*V1ContractsCreateCustomerWithContractContractSubscriptionParams `form:"subscriptions" json:"subscriptions,omitempty"`
	// This field's availability is dependent on your client's configuration.
	TotalContractValue     *Number                                                                    `form:"total_contract_value" json:"total_contract_value,omitempty"`
	UniquenessKey          *string                                                                    `form:"uniqueness_key" json:"uniqueness_key,omitempty"`
	UsageFilter            *V1ContractsCreateCustomerWithContractContractUsageFilterParams            `form:"usage_filter" json:"usage_filter,omitempty"`
	UsageStatementSchedule *V1ContractsCreateCustomerWithContractContractUsageStatementScheduleParams `form:"usage_statement_schedule" json:"usage_statement_schedule,omitempty"`
}

// Create a new customer and provision a contract. This endpoint's availability is dependent on your client's configuration.
type V1ContractsCreateCustomerWithContractParams struct {
	Params   `form:"*"`
	Contract *V1ContractsCreateCustomerWithContractContractParams `form:"contract" json:"contract"`
	Customer *V1ContractsCreateCustomerWithContractCustomerParams `form:"customer" json:"customer"`
}
type V1ContractsCreateHistoricalContractUsageInvoicesInvoiceUsageLineItemSubtotalsWithQuantityParams struct {
	ExclusiveEndDate   *string `form:"exclusive_end_date" json:"exclusive_end_date"`
	InclusiveStartDate *string `form:"inclusive_start_date" json:"inclusive_start_date"`
	Quantity           *Number `form:"quantity" json:"quantity"`
}
type V1ContractsCreateHistoricalContractUsageInvoicesInvoiceUsageLineItemParams struct {
	ExclusiveEndDate        *string                                                                                            `form:"exclusive_end_date" json:"exclusive_end_date"`
	InclusiveStartDate      *string                                                                                            `form:"inclusive_start_date" json:"inclusive_start_date"`
	PresentationGroupValues map[string]string                                                                                  `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	PricingGroupValues      map[string]string                                                                                  `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	ProductID               *string                                                                                            `form:"product_id" json:"product_id"`
	Quantity                *Number                                                                                            `form:"quantity" json:"quantity,omitempty"`
	SubtotalsWithQuantity   []*V1ContractsCreateHistoricalContractUsageInvoicesInvoiceUsageLineItemSubtotalsWithQuantityParams `form:"subtotals_with_quantity" json:"subtotals_with_quantity,omitempty"`
}
type V1ContractsCreateHistoricalContractUsageInvoicesInvoiceParams struct {
	// This field's availability is dependent on your client's configuration.
	BillableStatus       *string                                                                       `form:"billable_status" json:"billable_status,omitempty"`
	BreakdownGranularity *string                                                                       `form:"breakdown_granularity" json:"breakdown_granularity,omitempty"`
	ContractID           *string                                                                       `form:"contract_id" json:"contract_id"`
	CreditTypeID         *string                                                                       `form:"credit_type_id" json:"credit_type_id"`
	CustomerID           *string                                                                       `form:"customer_id" json:"customer_id"`
	CustomFields         map[string]string                                                             `form:"custom_fields" json:"custom_fields,omitempty"`
	ExclusiveEndDate     *string                                                                       `form:"exclusive_end_date" json:"exclusive_end_date"`
	InclusiveStartDate   *string                                                                       `form:"inclusive_start_date" json:"inclusive_start_date"`
	IssueDate            *string                                                                       `form:"issue_date" json:"issue_date"`
	UsageLineItems       []*V1ContractsCreateHistoricalContractUsageInvoicesInvoiceUsageLineItemParams `form:"usage_line_items" json:"usage_line_items"`
}

// Create historical usage invoices for past billing periods on specific contracts. Use this endpoint to generate retroactive invoices with custom usage line items, quantities, and date ranges. Supports preview mode to validate invoice data before creation. Ideal for billing migrations or correcting past billing periods.
type V1ContractsCreateHistoricalContractUsageInvoicesParams struct {
	Params   `form:"*"`
	Invoices []*V1ContractsCreateHistoricalContractUsageInvoicesInvoiceParams `form:"invoices" json:"invoices"`
	Preview  *bool                                                            `form:"preview" json:"preview"`
}

// Reference this alias when creating a contract. If the same alias is assigned to multiple packages, it will reference the package to which it was most recently assigned. It is not exposed to end customers.
type V1ContractsCreatePackageAliasParams struct {
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	Name         *string `form:"name" json:"name"`
	StartingAt   *string `form:"starting_at" json:"starting_at,omitempty"`
}
type V1ContractsCreatePackageDurationParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}

// Date relative to the contract start date indicating the start of this schedule segment.
type V1ContractsCreatePackageCommitAccessScheduleScheduleItemStartingAtOffsetParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}

// Offset relative to the start of this segment indicating when it should end.
type V1ContractsCreatePackageCommitAccessScheduleScheduleItemDurationParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}
type V1ContractsCreatePackageCommitAccessScheduleScheduleItemParams struct {
	Amount *Number `form:"amount" json:"amount"`
	// Offset relative to the start of this segment indicating when it should end.
	Duration *V1ContractsCreatePackageCommitAccessScheduleScheduleItemDurationParams `form:"duration" json:"duration"`
	// Date relative to the contract start date indicating the start of this schedule segment.
	StartingAtOffset *V1ContractsCreatePackageCommitAccessScheduleScheduleItemStartingAtOffsetParams `form:"starting_at_offset" json:"starting_at_offset"`
}

// Required: Schedule for distributing the commit to the customer. For "POSTPAID" commits only one schedule item is allowed and amount must match invoice_schedule total.
type V1ContractsCreatePackageCommitAccessScheduleParams struct {
	// Determines how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used. Defaults to `SPEND` if omitted.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID  *string                                                           `form:"credit_type_id" json:"credit_type_id,omitempty"`
	ScheduleItems []*V1ContractsCreatePackageCommitAccessScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items"`
}

// Date relative to the contract start date.
type V1ContractsCreatePackageCommitInvoiceScheduleScheduleItemDateOffsetParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}

// Either provide amount or provide both unit_price and quantity.
type V1ContractsCreatePackageCommitInvoiceScheduleScheduleItemParams struct {
	// Date relative to the contract start date.
	DateOffset *V1ContractsCreatePackageCommitInvoiceScheduleScheduleItemDateOffsetParams `form:"date_offset" json:"date_offset"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the amount.
	Quantity *Number `form:"quantity" json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the amount.
	UnitPrice *Number `form:"unit_price" json:"unit_price"`
}

// Required for "POSTPAID" commits: the true up invoice will be generated at this time and only one schedule item is allowed; the total must match access_schedule amount. Optional for "PREPAID" commits: if not provided, this will be a "complimentary" commit with no invoice.
type V1ContractsCreatePackageCommitInvoiceScheduleParams struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// If true, this schedule will not generate an invoice.
	DoNotInvoice *bool `form:"do_not_invoice" json:"do_not_invoice,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []*V1ContractsCreatePackageCommitInvoiceScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreatePackageCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreatePackageCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreatePackageCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}
type V1ContractsCreatePackageCommitParams struct {
	// Required: Schedule for distributing the commit to the customer. For "POSTPAID" commits only one schedule item is allowed and amount must match invoice_schedule total.
	AccessSchedule *V1ContractsCreatePackageCommitAccessScheduleParams `form:"access_schedule" json:"access_schedule"`
	// Which products the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductTags []*string         `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	CustomFields          map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// Used only in UI/API. It is not exposed to end customers.
	Description *string `form:"description" json:"description,omitempty"`
	// Required for "POSTPAID" commits: the true up invoice will be generated at this time and only one schedule item is allowed; the total must match access_schedule amount. Optional for "PREPAID" commits: if not provided, this will be a "complimentary" commit with no invoice.
	InvoiceSchedule *V1ContractsCreatePackageCommitInvoiceScheduleParams `form:"invoice_schedule" json:"invoice_schedule,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// If multiple commits are applicable, the one with the lower priority will apply first.
	Priority  *Number `form:"priority" json:"priority,omitempty"`
	ProductID *string `form:"product_id" json:"product_id"`
	RateType  *string `form:"rate_type" json:"rate_type,omitempty"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreatePackageCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
	// A temporary ID for the commit that can be used to reference the commit for commit specific overrides.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
	Type        *string `form:"type" json:"type"`
}

// Date relative to the contract start date indicating the start of this schedule segment.
type V1ContractsCreatePackageCreditAccessScheduleScheduleItemStartingAtOffsetParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}

// Offset relative to the start of this segment indicating when it should end.
type V1ContractsCreatePackageCreditAccessScheduleScheduleItemDurationParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}
type V1ContractsCreatePackageCreditAccessScheduleScheduleItemParams struct {
	Amount *Number `form:"amount" json:"amount"`
	// Offset relative to the start of this segment indicating when it should end.
	Duration *V1ContractsCreatePackageCreditAccessScheduleScheduleItemDurationParams `form:"duration" json:"duration"`
	// Date relative to the contract start date indicating the start of this schedule segment.
	StartingAtOffset *V1ContractsCreatePackageCreditAccessScheduleScheduleItemStartingAtOffsetParams `form:"starting_at_offset" json:"starting_at_offset"`
}

// Schedule for distributing the credit to the customer.
type V1ContractsCreatePackageCreditAccessScheduleParams struct {
	// Determines how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used. Defaults to `SPEND` if omitted.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID  *string                                                           `form:"credit_type_id" json:"credit_type_id,omitempty"`
	ScheduleItems []*V1ContractsCreatePackageCreditAccessScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreatePackageCreditSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreatePackageCreditSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreatePackageCreditSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}
type V1ContractsCreatePackageCreditParams struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule *V1ContractsCreatePackageCreditAccessScheduleParams `form:"access_schedule" json:"access_schedule"`
	// Which products the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductTags []*string         `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	CustomFields          map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// Used only in UI/API. It is not exposed to end customers.
	Description *string `form:"description" json:"description,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// If multiple credits are applicable, the one with the lower priority will apply first.
	Priority  *Number `form:"priority" json:"priority,omitempty"`
	ProductID *string `form:"product_id" json:"product_id"`
	RateType  *string `form:"rate_type" json:"rate_type,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreatePackageCreditSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
}

// The amount of commit to grant.
type V1ContractsCreatePackageRecurringCommitAccessAmountParams struct {
	// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is required unless a subscription is attached via `subscription_config`.
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	UnitPrice *Number `form:"unit_price" json:"unit_price"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreatePackageRecurringCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreatePackageRecurringCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreatePackageRecurringCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Offset relative to the contract start date that determines the start time for the first commit
type V1ContractsCreatePackageRecurringCommitStartingAtOffsetParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}

// Offset relative to the recurring credit start that determines when the contract will stop creating recurring commits. optional
type V1ContractsCreatePackageRecurringCommitDurationParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}

// Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.
type V1ContractsCreatePackageRecurringCommitCommitDurationParams struct {
	Unit  *string `form:"unit" json:"unit,omitempty"`
	Value *Number `form:"value" json:"value"`
}
type V1ContractsCreatePackageRecurringCommitSubscriptionConfigApplySeatIncreaseConfigParams struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated *bool `form:"is_prorated" json:"is_prorated"`
}

// Attach a subscription to the recurring commit/credit.
type V1ContractsCreatePackageRecurringCommitSubscriptionConfigParams struct {
	// Controls when a customer can access recurring commit or credit child balances when the attached subscription is payment-gated. `BILLING_PERIOD_PAID` releases each balance after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases all balances after the first payment.
	AccessPolicy            *string                                                                                 `form:"access_policy" json:"access_policy,omitempty"`
	Allocation              *string                                                                                 `form:"allocation" json:"allocation,omitempty"`
	ApplySeatIncreaseConfig *V1ContractsCreatePackageRecurringCommitSubscriptionConfigApplySeatIncreaseConfigParams `form:"apply_seat_increase_config" json:"apply_seat_increase_config"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID *string `form:"subscription_id" json:"subscription_id"`
}

// The amount the customer should be billed for the commit. Not required.
type V1ContractsCreatePackageRecurringCommitInvoiceAmountParams struct {
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id"`
	Quantity     *Number `form:"quantity" json:"quantity"`
	UnitPrice    *Number `form:"unit_price" json:"unit_price"`
}
type V1ContractsCreatePackageRecurringCommitProrationRoundingAccessParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}
type V1ContractsCreatePackageRecurringCommitProrationRoundingInvoiceParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}

// Optional rounding configuration for prorated recurring commit amounts.
type V1ContractsCreatePackageRecurringCommitProrationRoundingParams struct {
	Access  *V1ContractsCreatePackageRecurringCommitProrationRoundingAccessParams  `form:"access" json:"access,omitempty"`
	Invoice *V1ContractsCreatePackageRecurringCommitProrationRoundingInvoiceParams `form:"invoice" json:"invoice,omitempty"`
}
type V1ContractsCreatePackageRecurringCommitParams struct {
	// The amount of commit to grant.
	AccessAmount *V1ContractsCreatePackageRecurringCommitAccessAmountParams `form:"access_amount" json:"access_amount"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Will be passed down to the individual commits
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	// Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.
	CommitDuration *V1ContractsCreatePackageRecurringCommitCommitDurationParams `form:"commit_duration" json:"commit_duration"`
	// Will be passed down to the individual commits
	Description *string `form:"description" json:"description,omitempty"`
	// Offset relative to the recurring credit start that determines when the contract will stop creating recurring commits. optional
	Duration *V1ContractsCreatePackageRecurringCommitDurationParams `form:"duration" json:"duration,omitempty"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount *V1ContractsCreatePackageRecurringCommitInvoiceAmountParams `form:"invoice_amount" json:"invoice_amount,omitempty"`
	// displayed on invoices. will be passed through to the individual commits
	Name *string `form:"name" json:"name,omitempty"`
	// Will be passed down to the individual commits
	Priority  *Number `form:"priority" json:"priority"`
	ProductID *string `form:"product_id" json:"product_id"`
	// Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration *string `form:"proration" json:"proration,omitempty"`
	// Optional rounding configuration for prorated recurring commit amounts.
	ProrationRounding *V1ContractsCreatePackageRecurringCommitProrationRoundingParams `form:"proration_rounding" json:"proration_rounding,omitempty"`
	// Whether the created commits will use the commit rate or list rate
	RateType *string `form:"rate_type" json:"rate_type,omitempty"`
	// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
	RecurrenceFrequency *string `form:"recurrence_frequency" json:"recurrence_frequency,omitempty"`
	// Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreatePackageRecurringCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
	// Offset relative to the contract start date that determines the start time for the first commit
	StartingAtOffset *V1ContractsCreatePackageRecurringCommitStartingAtOffsetParams `form:"starting_at_offset" json:"starting_at_offset"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig *V1ContractsCreatePackageRecurringCommitSubscriptionConfigParams `form:"subscription_config" json:"subscription_config,omitempty"`
	// A temporary ID that can be used to reference the recurring commit for commit specific overrides.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
}

// The amount of commit to grant.
type V1ContractsCreatePackageRecurringCreditAccessAmountParams struct {
	// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is required unless a subscription is attached via `subscription_config`.
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	UnitPrice *Number `form:"unit_price" json:"unit_price"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreatePackageRecurringCreditSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreatePackageRecurringCreditSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreatePackageRecurringCreditSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Offset relative to the contract start date that determines the start time for the first commit
type V1ContractsCreatePackageRecurringCreditStartingAtOffsetParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}

// Offset relative to the recurring credit start that determines when the contract will stop creating recurring commits. optional
type V1ContractsCreatePackageRecurringCreditDurationParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}

// Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.
type V1ContractsCreatePackageRecurringCreditCommitDurationParams struct {
	Unit  *string `form:"unit" json:"unit,omitempty"`
	Value *Number `form:"value" json:"value"`
}
type V1ContractsCreatePackageRecurringCreditSubscriptionConfigApplySeatIncreaseConfigParams struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated *bool `form:"is_prorated" json:"is_prorated"`
}

// Attach a subscription to the recurring commit/credit.
type V1ContractsCreatePackageRecurringCreditSubscriptionConfigParams struct {
	// Controls when a customer can access recurring commit or credit child balances when the attached subscription is payment-gated. `BILLING_PERIOD_PAID` releases each balance after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases all balances after the first payment.
	AccessPolicy            *string                                                                                 `form:"access_policy" json:"access_policy,omitempty"`
	Allocation              *string                                                                                 `form:"allocation" json:"allocation,omitempty"`
	ApplySeatIncreaseConfig *V1ContractsCreatePackageRecurringCreditSubscriptionConfigApplySeatIncreaseConfigParams `form:"apply_seat_increase_config" json:"apply_seat_increase_config"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID *string `form:"subscription_id" json:"subscription_id"`
}
type V1ContractsCreatePackageRecurringCreditProrationRoundingAccessParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}

// Optional rounding configuration for prorated recurring credit amounts.
type V1ContractsCreatePackageRecurringCreditProrationRoundingParams struct {
	Access *V1ContractsCreatePackageRecurringCreditProrationRoundingAccessParams `form:"access" json:"access,omitempty"`
}
type V1ContractsCreatePackageRecurringCreditParams struct {
	// The amount of commit to grant.
	AccessAmount *V1ContractsCreatePackageRecurringCreditAccessAmountParams `form:"access_amount" json:"access_amount"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Will be passed down to the individual commits
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	// Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.
	CommitDuration *V1ContractsCreatePackageRecurringCreditCommitDurationParams `form:"commit_duration" json:"commit_duration"`
	// Will be passed down to the individual commits
	Description *string `form:"description" json:"description,omitempty"`
	// Offset relative to the recurring credit start that determines when the contract will stop creating recurring commits. optional
	Duration *V1ContractsCreatePackageRecurringCreditDurationParams `form:"duration" json:"duration,omitempty"`
	// displayed on invoices. will be passed through to the individual commits
	Name *string `form:"name" json:"name,omitempty"`
	// Will be passed down to the individual commits
	Priority  *Number `form:"priority" json:"priority"`
	ProductID *string `form:"product_id" json:"product_id"`
	// Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration *string `form:"proration" json:"proration,omitempty"`
	// Optional rounding configuration for prorated recurring credit amounts.
	ProrationRounding *V1ContractsCreatePackageRecurringCreditProrationRoundingParams `form:"proration_rounding" json:"proration_rounding,omitempty"`
	// Whether the created commits will use the commit rate or list rate
	RateType *string `form:"rate_type" json:"rate_type,omitempty"`
	// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
	RecurrenceFrequency *string `form:"recurrence_frequency" json:"recurrence_frequency,omitempty"`
	// Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreatePackageRecurringCreditSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
	// Offset relative to the contract start date that determines the start time for the first commit
	StartingAtOffset *V1ContractsCreatePackageRecurringCreditStartingAtOffsetParams `form:"starting_at_offset" json:"starting_at_offset"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig *V1ContractsCreatePackageRecurringCreditSubscriptionConfigParams `form:"subscription_config" json:"subscription_config,omitempty"`
	// A temporary ID that can be used to reference the recurring commit for commit specific overrides.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
}

// Offset relative to contract start date indicating when the override will start applying (inclusive)
type V1ContractsCreatePackageOverrideStartingAtOffsetParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}

// Offset relative to override start indicating when the override will stop applying (exclusive)
type V1ContractsCreatePackageOverrideDurationParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}

// Only set for TIERED rate_type.
type V1ContractsCreatePackageOverrideOverwriteRateTierParams struct {
	Price *Number `form:"price" json:"price"`
	Size  *Number `form:"size" json:"size,omitempty"`
}
type V1ContractsCreatePackageOverrideOverwriteRateMinimumConfigParams struct {
	Minimum *Number `form:"minimum" json:"minimum"`
}

// Required for OVERWRITE type.
type V1ContractsCreatePackageOverrideOverwriteRateParams struct {
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate processors.
	CustomRate map[string]JSON `form:"custom_rate" json:"custom_rate,omitempty"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be set to true.
	IsProrated    *bool                                                             `form:"is_prorated" json:"is_prorated,omitempty"`
	MinimumConfig *V1ContractsCreatePackageOverrideOverwriteRateMinimumConfigParams `form:"minimum_config" json:"minimum_config,omitempty"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price *Number `form:"price" json:"price,omitempty"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity *Number `form:"quantity" json:"quantity,omitempty"`
	RateType *string `form:"rate_type" json:"rate_type"`
	// Only set for TIERED rate_type.
	Tiers []*V1ContractsCreatePackageOverrideOverwriteRateTierParams `form:"tiers" json:"tiers,omitempty"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreatePackageOverrideOverrideSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Specifies which products the override will apply to.
type V1ContractsCreatePackageOverrideOverrideSpecifierParams struct {
	// Can only be used for commit specific overrides. Must be used in conjunction with one of `product_id`, `product_tags`, `pricing_group_values`, or `presentation_group_values`. Must be used instead of both `commit_ids` and `recurring_commit_ids` If provided, the override will apply to any specified commit, credit, recurring commit or recurring credit IDs.
	AnyCommitOrCreditIDs []*string `form:"any_commit_or_credit_ids" json:"any_commit_or_credit_ids,omitempty"`
	BillingFrequency     *string   `form:"billing_frequency" json:"billing_frequency,omitempty"`
	// Can only be used for commit specific overrides. Must be used in conjunction with one of `product_id`, `product_tags`, `pricing_group_values`, or `presentation_group_values`. If provided, the override will only apply to the specified commits. If not provided, the override will apply to all commits.
	CommitIDs []*string `form:"commit_ids" json:"commit_ids,omitempty"`
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreatePackageOverrideOverrideSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// A map of group names to values. The override will only apply to line items with the specified presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// A map of pricing group names to values. The override will only apply to products with the specified pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the override will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the override will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
	// Can only be used for commit specific overrides. Must be used in conjunction with one of `product_id`, `product_tags`, `pricing_group_values`, or `presentation_group_values`. If provided, the override will only apply to commits created by the specified recurring commit ids.
	RecurringCommitIDs []*string `form:"recurring_commit_ids" json:"recurring_commit_ids,omitempty"`
}

// Required for TIERED type. Must have at least one tier.
type V1ContractsCreatePackageOverrideTierParams struct {
	Multiplier *Number `form:"multiplier" json:"multiplier"`
	Size       *Number `form:"size" json:"size,omitempty"`
}
type V1ContractsCreatePackageOverrideParams struct {
	// Offset relative to override start indicating when the override will stop applying (exclusive)
	Duration *V1ContractsCreatePackageOverrideDurationParams `form:"duration" json:"duration,omitempty"`
	Entitled *bool                                           `form:"entitled" json:"entitled,omitempty"`
	// Indicates whether the override should only apply to commits. Defaults to `false`. If `true` you can specify relevant commits in `override_specifiers` by passing `commit_ids`, `recurring_commit_ids`, or `any_commit_or_credit_ids`.  If you do not specify any of these fields, the override will apply when consuming any prepaid commit, postpaid commit, or credit
	IsCommitSpecific *bool `form:"is_commit_specific" json:"is_commit_specific,omitempty"`
	// Required for MULTIPLIER type. Must be >=0.
	Multiplier *Number `form:"multiplier" json:"multiplier,omitempty"`
	// Specifies which products the override will apply to.
	OverrideSpecifiers []*V1ContractsCreatePackageOverrideOverrideSpecifierParams `form:"override_specifiers" json:"override_specifiers"`
	// Required for OVERWRITE type.
	OverwriteRate *V1ContractsCreatePackageOverrideOverwriteRateParams `form:"overwrite_rate" json:"overwrite_rate,omitempty"`
	// Required for EXPLICIT multiplier prioritization scheme and all TIERED overrides. Under EXPLICIT prioritization, overwrites are prioritized first, and then tiered and multiplier overrides are prioritized by their priority value (lowest first). Must be > 0.
	Priority *Number `form:"priority" json:"priority,omitempty"`
	// Offset relative to contract start date indicating when the override will start applying (inclusive)
	StartingAtOffset *V1ContractsCreatePackageOverrideStartingAtOffsetParams `form:"starting_at_offset" json:"starting_at_offset"`
	// Indicates whether the override applies to commit rates or list rates. Can only be used for overrides that have `is_commit_specific` set to `true`. Defaults to `"LIST_RATE"`.
	Target *string `form:"target" json:"target,omitempty"`
	// Required for TIERED type. Must have at least one tier.
	Tiers []*V1ContractsCreatePackageOverrideTierParams `form:"tiers" json:"tiers,omitempty"`
	// Overwrites are prioritized over multipliers and tiered overrides.
	Type *string `form:"type" json:"type,omitempty"`
}

// Date relative to the contract start date.
type V1ContractsCreatePackageScheduledChargeScheduleScheduleItemDateOffsetParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}

// Either provide amount or provide both unit_price and quantity.
type V1ContractsCreatePackageScheduledChargeScheduleScheduleItemParams struct {
	// Date relative to the contract start date.
	DateOffset *V1ContractsCreatePackageScheduledChargeScheduleScheduleItemDateOffsetParams `form:"date_offset" json:"date_offset"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the amount.
	Quantity *Number `form:"quantity" json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the amount.
	UnitPrice *Number `form:"unit_price" json:"unit_price"`
}
type V1ContractsCreatePackageScheduledChargeScheduleParams struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []*V1ContractsCreatePackageScheduledChargeScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items"`
}
type V1ContractsCreatePackageScheduledChargeParams struct {
	CustomFields map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// displayed on invoices
	Name      *string                                                `form:"name" json:"name,omitempty"`
	ProductID *string                                                `form:"product_id" json:"product_id"`
	Schedule  *V1ContractsCreatePackageScheduledChargeScheduleParams `form:"schedule" json:"schedule"`
}

// The offset at which Metronome should start generating usage invoices, relative to the contract start date.  If unspecified, contract start date will be used. This is useful to set if you want to import historical invoices via our 'Create Historical Invoices' API rather than having Metronome automatically generate them.
type V1ContractsCreatePackageUsageStatementScheduleInvoiceGenerationStartingAtOffsetParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}
type V1ContractsCreatePackageUsageStatementScheduleParams struct {
	// If not provided, defaults to the first day of the month.
	Day       *string `form:"day" json:"day,omitempty"`
	Frequency *string `form:"frequency" json:"frequency"`
	// The offset at which Metronome should start generating usage invoices, relative to the contract start date.  If unspecified, contract start date will be used. This is useful to set if you want to import historical invoices via our 'Create Historical Invoices' API rather than having Metronome automatically generate them.
	InvoiceGenerationStartingAtOffset *V1ContractsCreatePackageUsageStatementScheduleInvoiceGenerationStartingAtOffsetParams `form:"invoice_generation_starting_at_offset" json:"invoice_generation_starting_at_offset,omitempty"`
}
type V1ContractsCreatePackageSpendThresholdConfigurationCommitParams struct {
	Description *string `form:"description" json:"description,omitempty"`
	// Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.
	Name *string `form:"name" json:"name,omitempty"`
	// The priority of the commit, used to determine drawdown order. Lower priority commits are consumed first. Defaults to 100 if not specified.
	Priority *Number `form:"priority" json:"priority,omitempty"`
	// The commit product that will be used to generate the line item for commit payment.
	ProductID *string `form:"product_id" json:"product_id"`
}

// Only applicable if using STRIPE as your payment gate type.
type V1ContractsCreatePackageSpendThresholdConfigurationPaymentGateConfigStripeConfigParams struct {
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as your payment type.
	InvoiceMetadata map[string]string `form:"invoice_metadata" json:"invoice_metadata,omitempty"`
	// If left blank, will default to INVOICE
	PaymentType *string `form:"payment_type" json:"payment_type"`
}

// Only applicable if using PRECALCULATED as your tax type.
type V1ContractsCreatePackageSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams struct {
	// Amount of tax to be applied. This should be in the same currency and denomination  as the commit's invoice schedule
	TaxAmount *Number `form:"tax_amount" json:"tax_amount"`
	// Name of the tax to be applied. This may be used in an invoice line item description.
	TaxName *string `form:"tax_name" json:"tax_name,omitempty"`
}
type V1ContractsCreatePackageSpendThresholdConfigurationPaymentGateConfigParams struct {
	// Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.
	PaymentGateType *string `form:"payment_gate_type" json:"payment_gate_type"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig *V1ContractsCreatePackageSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams `form:"precalculated_tax_config" json:"precalculated_tax_config,omitempty"`
	// Only applicable if using STRIPE as your payment gate type.
	StripeConfig *V1ContractsCreatePackageSpendThresholdConfigurationPaymentGateConfigStripeConfigParams `form:"stripe_config" json:"stripe_config,omitempty"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType *string `form:"tax_type" json:"tax_type,omitempty"`
}

// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
type V1ContractsCreatePackageSpendThresholdConfigurationDiscountConfigurationCapParams struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount *Number `form:"amount" json:"amount"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias *string `form:"spend_tracker_alias" json:"spend_tracker_alias"`
}
type V1ContractsCreatePackageSpendThresholdConfigurationDiscountConfigurationParams struct {
	// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
	Cap *V1ContractsCreatePackageSpendThresholdConfigurationDiscountConfigurationCapParams `form:"cap" json:"cap,omitempty"`
	// The fraction of the original amount that the customer pays after applying the discount. For example, 0.85 means the customer pays 85% of the original amount (a 15% discount).
	PaymentFraction *Number `form:"payment_fraction" json:"payment_fraction"`
}
type V1ContractsCreatePackageSpendThresholdConfigurationParams struct {
	Commit                *V1ContractsCreatePackageSpendThresholdConfigurationCommitParams                `form:"commit" json:"commit"`
	DiscountConfiguration *V1ContractsCreatePackageSpendThresholdConfigurationDiscountConfigurationParams `form:"discount_configuration" json:"discount_configuration,omitempty"`
	// When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.
	IsEnabled         *bool                                                                       `form:"is_enabled" json:"is_enabled"`
	PaymentGateConfig *V1ContractsCreatePackageSpendThresholdConfigurationPaymentGateConfigParams `form:"payment_gate_config" json:"payment_gate_config"`
	// Specify the threshold amount for the contract. Each time the contract's usage hits this amount, a threshold charge will be initiated.
	ThresholdAmount *Number `form:"threshold_amount" json:"threshold_amount"`
}

// The length of time the created commit will be valid, starting from the end of the invoice's service period. If not provided, defaults to one year.
type V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationCommitDurationParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
type V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}
type V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationCommitParams struct {
	// Which products the threshold commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the threshold commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	Description           *string   `form:"description" json:"description,omitempty"`
	// The length of time the created commit will be valid, starting from the end of the invoice's service period. If not provided, defaults to one year.
	Duration *V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationCommitDurationParams `form:"duration" json:"duration,omitempty"`
	// Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.
	Name *string `form:"name" json:"name,omitempty"`
	// The priority of the commit, used to determine drawdown order. Lower priority commits are consumed first. Defaults to 100 if not specified.
	Priority *Number `form:"priority" json:"priority,omitempty"`
	// The commit product that will be used to generate the line item for commit payment.
	ProductID *string `form:"product_id" json:"product_id"`
	// Whether the created commits will be charged at commit rate or list rate.
	RateType *string `form:"rate_type" json:"rate_type,omitempty"`
	// Fraction of the created commit's unused balance that will roll over. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.
	Specifiers []*V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
}

// Only applicable if using STRIPE as your payment gate type.
type V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigParams struct {
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as your payment type.
	InvoiceMetadata map[string]string `form:"invoice_metadata" json:"invoice_metadata,omitempty"`
	// If left blank, will default to INVOICE
	PaymentType *string `form:"payment_type" json:"payment_type"`
}

// Only applicable if using PRECALCULATED as your tax type.
type V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams struct {
	// Amount of tax to be applied. This should be in the same currency and denomination  as the commit's invoice schedule
	TaxAmount *Number `form:"tax_amount" json:"tax_amount"`
	// Name of the tax to be applied. This may be used in an invoice line item description.
	TaxName *string `form:"tax_name" json:"tax_name,omitempty"`
}
type V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationPaymentGateConfigParams struct {
	// Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.
	PaymentGateType *string `form:"payment_gate_type" json:"payment_gate_type"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig *V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams `form:"precalculated_tax_config" json:"precalculated_tax_config,omitempty"`
	// Only applicable if using STRIPE as your payment gate type.
	StripeConfig *V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigParams `form:"stripe_config" json:"stripe_config,omitempty"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType *string `form:"tax_type" json:"tax_type,omitempty"`
}

// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
type V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationDiscountConfigurationCapParams struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount *Number `form:"amount" json:"amount"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias *string `form:"spend_tracker_alias" json:"spend_tracker_alias"`
}
type V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationDiscountConfigurationParams struct {
	// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
	Cap *V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationDiscountConfigurationCapParams `form:"cap" json:"cap,omitempty"`
	// The fraction of the original amount that the customer pays after applying the discount. For example, 0.85 means the customer pays 85% of the original amount (a 15% discount).
	PaymentFraction *Number `form:"payment_fraction" json:"payment_fraction"`
}

// If provided, balances with all the custom fields will not be considered when evaluating threshold billing
type V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilterParams struct {
	Entity *string `form:"entity" json:"entity"`
	Key    *string `form:"key" json:"key"`
	Value  *string `form:"value" json:"value"`
}

// If any of the exclude specifier is met, the balance is not considered when evaluating threshold billing
type V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeParams struct {
	// If provided, balances with all the custom fields will not be considered when evaluating threshold billing
	CustomFieldFilters []*V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilterParams `form:"custom_field_filters" json:"custom_field_filters"`
}

// Determines which balances are excluded from remaining balance calculation for threshold billing.
type V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierParams struct {
	// If any of the exclude specifier is met, the balance is not considered when evaluating threshold billing
	Exclude []*V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeParams `form:"exclude" json:"exclude"`
}
type V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationParams struct {
	Commit *V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationCommitParams `form:"commit" json:"commit"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID    *string                                                                                  `form:"custom_credit_type_id" json:"custom_credit_type_id,omitempty"`
	DiscountConfiguration *V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationDiscountConfigurationParams `form:"discount_configuration" json:"discount_configuration,omitempty"`
	// When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.
	IsEnabled         *bool                                                                                `form:"is_enabled" json:"is_enabled"`
	PaymentGateConfig *V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationPaymentGateConfigParams `form:"payment_gate_config" json:"payment_gate_config"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount *Number `form:"recharge_to_amount" json:"recharge_to_amount"`
	// Specify the threshold amount for the contract. Each time the contract's prepaid balance lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount *Number `form:"threshold_amount" json:"threshold_amount"`
	// Determines which balances are excluded from remaining balance calculation for threshold billing.
	ThresholdBalanceSpecifiers []*V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierParams `form:"threshold_balance_specifiers" json:"threshold_balance_specifiers,omitempty"`
}
type V1ContractsCreatePackageSpendTrackerApplicableSpendSpecifierParams struct {
	// Filter by whether the spend was discounted. Defaults to ANY if omitted.
	Discounted *string   `form:"discounted" json:"discounted,omitempty"`
	Sources    []*string `form:"sources" json:"sources"`
	SpendType  *string   `form:"spend_type" json:"spend_type"`
}
type V1ContractsCreatePackageSpendTrackerParams struct {
	// Human-readable identifier, unique per contract.
	Alias                     *string                                                               `form:"alias" json:"alias"`
	ApplicableSpendSpecifiers []*V1ContractsCreatePackageSpendTrackerApplicableSpendSpecifierParams `form:"applicable_spend_specifiers" json:"applicable_spend_specifiers"`
	CreditTypeID              *string                                                               `form:"credit_type_id" json:"credit_type_id"`
	ResetFrequency            *string                                                               `form:"reset_frequency" json:"reset_frequency"`
}
type V1ContractsCreatePackageSubscriptionSubscriptionRateParams struct {
	// Frequency to bill subscription with. Together with product_id, must match existing rate on the rate card.
	BillingFrequency *string `form:"billing_frequency" json:"billing_frequency"`
	// Must be subscription type product
	ProductID *string `form:"product_id" json:"product_id"`
}
type V1ContractsCreatePackageSubscriptionProrationRoundingParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}
type V1ContractsCreatePackageSubscriptionProrationParams struct {
	// Indicates how mid-period quantity adjustments are invoiced.
	// **BILL_IMMEDIATELY**: Only available when collection schedule is `ADVANCE`. The quantity increase will be billed immediately on the scheduled date.
	// **BILL_ON_NEXT_COLLECTION_DATE**: The quantity increase will be billed for in-arrears at the end of the period.
	//
	InvoiceBehavior *string `form:"invoice_behavior" json:"invoice_behavior,omitempty"`
	// Indicates if the partial period will be prorated or charged a full amount.
	IsProrated *bool                                                        `form:"is_prorated" json:"is_prorated,omitempty"`
	Rounding   *V1ContractsCreatePackageSubscriptionProrationRoundingParams `form:"rounding" json:"rounding,omitempty"`
}

// Relative date from contract start date corresponding to the inclusive start time for the subscription. If not provided, defaults to contract start date
type V1ContractsCreatePackageSubscriptionStartingAtOffsetParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}

// Lifetime of the subscription from its start. If not provided, subscription inherits contract end date.
type V1ContractsCreatePackageSubscriptionDurationParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}
type V1ContractsCreatePackageSubscriptionSeatConfigParams struct {
	// The initial amount of unassigned seats on this subscription.
	InitialUnassignedSeats *Number `form:"initial_unassigned_seats" json:"initial_unassigned_seats,omitempty"`
	// The property name, sent on usage events, that identifies the seat ID associated with the usage event. For example, the property name might be seat_id or user_id. The property must be set as a group key on billable metrics and a presentation/pricing group key on contract products. This allows linked recurring credits with an allocation per seat to be consumed by only one seat's usage.
	SeatGroupKey *string `form:"seat_group_key" json:"seat_group_key"`
}
type V1ContractsCreatePackageSubscriptionBillingCycleConfigParams struct {
	// Controls whether subscriptions consolidate onto usage invoices. Defaults to ON_USAGE_INVOICE if omitted.
	InvoicePlacement *string `form:"invoice_placement" json:"invoice_placement,omitempty"`
}
type V1ContractsCreatePackageSubscriptionPaymentGateConfigParams struct {
	// Gate access to the subscription based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe.
	PaymentGateType *string `form:"payment_gate_type" json:"payment_gate_type"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType *string `form:"tax_type" json:"tax_type,omitempty"`
}
type V1ContractsCreatePackageSubscriptionParams struct {
	BillingCycleConfig *V1ContractsCreatePackageSubscriptionBillingCycleConfigParams `form:"billing_cycle_config" json:"billing_cycle_config,omitempty"`
	CollectionSchedule *string                                                       `form:"collection_schedule" json:"collection_schedule"`
	CustomFields       map[string]string                                             `form:"custom_fields" json:"custom_fields,omitempty"`
	Description        *string                                                       `form:"description" json:"description,omitempty"`
	// Lifetime of the subscription from its start. If not provided, subscription inherits contract end date.
	Duration *V1ContractsCreatePackageSubscriptionDurationParams `form:"duration" json:"duration,omitempty"`
	// The initial quantity for the subscription. It must be non-negative value. Required if quantity_management_mode is QUANTITY_ONLY.
	InitialQuantity   *Number                                                      `form:"initial_quantity" json:"initial_quantity,omitempty"`
	Name              *string                                                      `form:"name" json:"name,omitempty"`
	PaymentGateConfig *V1ContractsCreatePackageSubscriptionPaymentGateConfigParams `form:"payment_gate_config" json:"payment_gate_config,omitempty"`
	Proration         *V1ContractsCreatePackageSubscriptionProrationParams         `form:"proration" json:"proration"`
	// Determines how the subscription's quantity is controlled. Defaults to QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified directly on the subscription. `initial_quantity` must be provided with this option. Compatible with recurring commits/credits that use POOLED allocation. **SEAT_BASED**: Use when you want to pass specific seat identifiers (e.g. add user_123) to increment and decrement a subscription quantity, rather than directly providing the quantity. You must use a **SEAT_BASED** subscription to use a linked recurring credit with an allocation per seat. `seat_config` must be provided with this option.
	QuantityManagementMode *string                                               `form:"quantity_management_mode" json:"quantity_management_mode,omitempty"`
	SeatConfig             *V1ContractsCreatePackageSubscriptionSeatConfigParams `form:"seat_config" json:"seat_config,omitempty"`
	// Relative date from contract start date corresponding to the inclusive start time for the subscription. If not provided, defaults to contract start date
	StartingAtOffset *V1ContractsCreatePackageSubscriptionStartingAtOffsetParams `form:"starting_at_offset" json:"starting_at_offset,omitempty"`
	// Optional human-readable, per-package-unique alias for this subscription template. Lets a contract customization reference the subscription by alias instead of its template id.
	SubscriptionAlias *string                                                     `form:"subscription_alias" json:"subscription_alias,omitempty"`
	SubscriptionRate  *V1ContractsCreatePackageSubscriptionSubscriptionRateParams `form:"subscription_rate" json:"subscription_rate"`
	// A temporary ID used to reference the subscription in recurring commit/credit subscription configs created within the same payload.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
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
type V1ContractsCreatePackageParams struct {
	Params `form:"*"`
	// Reference this alias when creating a contract. If the same alias is assigned to multiple packages, it will reference the package to which it was most recently assigned. It is not exposed to end customers.
	Aliases         []*V1ContractsCreatePackageAliasParams  `form:"aliases" json:"aliases,omitempty"`
	BillingProvider *string                                 `form:"billing_provider" json:"billing_provider,omitempty"`
	Commits         []*V1ContractsCreatePackageCommitParams `form:"commits" json:"commits,omitempty"`
	ContractName    *string                                 `form:"contract_name" json:"contract_name,omitempty"`
	Credits         []*V1ContractsCreatePackageCreditParams `form:"credits" json:"credits,omitempty"`
	DeliveryMethod  *string                                 `form:"delivery_method" json:"delivery_method,omitempty"`
	Duration        *V1ContractsCreatePackageDurationParams `form:"duration" json:"duration,omitempty"`
	// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list prices automatically. EXPLICIT prioritization requires specifying priorities for each multiplier; the one with the lowest priority value will be prioritized first. If tiered overrides are used, prioritization must be explicit.
	MultiplierOverridePrioritization     *string                                                             `form:"multiplier_override_prioritization" json:"multiplier_override_prioritization,omitempty"`
	Name                                 *string                                                             `form:"name" json:"name"`
	NetPaymentTermsDays                  *Number                                                             `form:"net_payment_terms_days" json:"net_payment_terms_days,omitempty"`
	Overrides                            []*V1ContractsCreatePackageOverrideParams                           `form:"overrides" json:"overrides,omitempty"`
	PrepaidBalanceThresholdConfiguration *V1ContractsCreatePackagePrepaidBalanceThresholdConfigurationParams `form:"prepaid_balance_threshold_configuration" json:"prepaid_balance_threshold_configuration,omitempty"`
	// Selects the rate card linked to the specified alias as of the contract's start date.
	RateCardAlias                   *string                                                    `form:"rate_card_alias" json:"rate_card_alias,omitempty"`
	RateCardID                      *string                                                    `form:"rate_card_id" json:"rate_card_id,omitempty"`
	RecurringCommits                []*V1ContractsCreatePackageRecurringCommitParams           `form:"recurring_commits" json:"recurring_commits,omitempty"`
	RecurringCredits                []*V1ContractsCreatePackageRecurringCreditParams           `form:"recurring_credits" json:"recurring_credits,omitempty"`
	ScheduledCharges                []*V1ContractsCreatePackageScheduledChargeParams           `form:"scheduled_charges" json:"scheduled_charges,omitempty"`
	ScheduledChargesOnUsageInvoices *string                                                    `form:"scheduled_charges_on_usage_invoices" json:"scheduled_charges_on_usage_invoices,omitempty"`
	SpendThresholdConfiguration     *V1ContractsCreatePackageSpendThresholdConfigurationParams `form:"spend_threshold_configuration" json:"spend_threshold_configuration,omitempty"`
	SpendTrackers                   []*V1ContractsCreatePackageSpendTrackerParams              `form:"spend_trackers" json:"spend_trackers,omitempty"`
	Subscriptions                   []*V1ContractsCreatePackageSubscriptionParams              `form:"subscriptions" json:"subscriptions,omitempty"`
	UniquenessKey                   *string                                                    `form:"uniqueness_key" json:"uniqueness_key,omitempty"`
	UsageStatementSchedule          *V1ContractsCreatePackageUsageStatementScheduleParams      `form:"usage_statement_schedule" json:"usage_statement_schedule,omitempty"`
}

// This is the v1 endpoint to get a contract. New clients should implement using the v2 endpoint.
type V1ContractsGetContractParams struct {
	Params     `form:"*"`
	ContractID *string `form:"contract_id" json:"contract_id"`
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// Include the balance of credits and commits in the response. Setting this flag may cause the query to be slower.
	IncludeBalance *bool `form:"include_balance" json:"include_balance,omitempty"`
	// Include commit ledgers in the response. Setting this flag may cause the query to be slower.
	IncludeLedgers *bool `form:"include_ledgers" json:"include_ledgers,omitempty"`
}

// List of rate selectors, rates matching ANY of the selectors will be included in the response. Passing no selectors will result in all rates being returned.
type V1ContractsGetContractRateScheduleSelectorParams struct {
	// Subscription rates matching the billing frequency will be included in the response.
	BillingFrequency *string `form:"billing_frequency" json:"billing_frequency,omitempty"`
	// List of pricing group key value pairs, rates containing the matching key / value pairs will be included in the response.
	PartialPricingGroupValues map[string]string `form:"partial_pricing_group_values" json:"partial_pricing_group_values,omitempty"`
	// List of pricing group key value pairs, rates matching all of the key / value pairs will be included in the response.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// Rates matching the product id will be included in the response.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// List of product tags, rates matching any of the tags will be included in the response.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// For a specific customer and contract, get the rates at a specific point in time. This endpoint takes the contract's rate card into consideration, including scheduled changes. It also takes into account overrides on the contract.
//
// For example, if you want to show your customer a summary of the prices they are paying, inclusive of any negotiated discounts or promotions, use this endpoint. This endpoint only returns rates that are entitled.
type V1ContractsGetContractRateScheduleParams struct {
	Params `form:"*"`
	// optional timestamp which overlaps with the returned rate schedule segments. When not specified, the current timestamp will be used.
	At *string `form:"at" json:"at,omitempty"`
	// ID of the contract to get the rate schedule for.
	ContractID *string `form:"contract_id" json:"contract_id"`
	// ID of the customer for whose contract to get the rate schedule for.
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// Max number of results that should be returned
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
	// List of rate selectors, rates matching ANY of the selectors will be included in the response. Passing no selectors will result in all rates being returned.
	Selectors []*V1ContractsGetContractRateScheduleSelectorParams `form:"selectors" json:"selectors,omitempty"`
}

// Get the history of subscription quantities and prices over time for a given `subscription_id`. This endpoint can be used to power an in-product experience where you show a customer their historical changes to seat count. Future changes are not included in this endpoint - use the `getContract` endpoint to view the future scheduled changes to a subscription's quantity.
//
// Subscriptions are used to model fixed recurring fees as well as seat-based recurring fees. To model changes to the number of seats in Metronome, you can increment or decrement the quantity on a subscription at any point in the past or future.
type V1ContractsGetSubscriptionQuantityHistoryParams struct {
	Params         `form:"*"`
	ContractID     *string `form:"contract_id" json:"contract_id"`
	CustomerID     *string `form:"customer_id" json:"customer_id"`
	SubscriptionID *string `form:"subscription_id" json:"subscription_id"`
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
type V1ContractsGetSubscriptionSeatsHistoryParams struct {
	Params     `form:"*"`
	ContractID *string `form:"contract_id" json:"contract_id"`
	// Get the seats history segment for the covering date. Cannot be used with `starting_at` or `ending_before`.
	CoveringDate *string `form:"covering_date" json:"covering_date,omitempty"`
	// Cursor for pagination. Use the value from the `next_page` field of the previous response to retrieve the next page of results.
	Cursor     *string `form:"cursor" json:"cursor,omitempty"`
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// Include seats history segments that are active at or before this timestamp. Use with `starting_at` to get a specific time range. If not set, there's no upper bound.
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// Maximum number of seat schedule entries to return. Defaults to 10. Required range: 1 <= x <= 10.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Include seats history segments that are active at or after this timestamp. Use with `ending_before` to get a specific time range. If not set, there's no lower bound.
	StartingAt     *string `form:"starting_at" json:"starting_at,omitempty"`
	SubscriptionID *string `form:"subscription_id" json:"subscription_id"`
}

// Get the current payment and invoice status for a payment-gated subscription.
type V1ContractsGetSubscriptionStatusParams struct {
	Params         `form:"*"`
	ContractID     *string `form:"contract_id" json:"contract_id"`
	CustomerID     *string `form:"customer_id" json:"customer_id"`
	SubscriptionID *string `form:"subscription_id" json:"subscription_id"`
}

// Retrieves a page of contracts for a specific customer, including pricing, terms, credits, and commitments. Use this to view a customer's contract history and current agreements for billing management. Returns contract details with optional ledgers and balance information.
//
// ### Usage guidelines:
// - Pagination: Results are limited to 20 contracts per page; use 'cursor' for more
//
// ⚠️ Note: This is the legacy v1 endpoint - new integrations should use the v2 endpoint for enhanced features.
type V1ContractsListContractsParams struct {
	Params `form:"*"`
	// Optional RFC 3339 timestamp. If provided, the response will include only contracts effective on the provided date.  This cannot be provided if the starting_at filter is provided.
	CoveringDate *string `form:"covering_date" json:"covering_date,omitempty"`
	// Cursor from a previous response to fetch the next page of contracts.
	Cursor     *string `form:"cursor" json:"cursor,omitempty"`
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// Include archived contracts in the response
	IncludeArchived *bool `form:"include_archived" json:"include_archived,omitempty"`
	// Include the balance of credits and commits in the response. Setting this flag may cause the query to be slower.
	IncludeBalance *bool `form:"include_balance" json:"include_balance,omitempty"`
	// Include commit ledgers in the response. Setting this flag may cause the query to be slower.
	IncludeLedgers *bool `form:"include_ledgers" json:"include_ledgers,omitempty"`
	// Max number of contracts to return per page. Range: 1-20. Default: 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Optional RFC 3339 timestamp. If provided, the response will include only contracts where starting_at is on or after the provided date.  This cannot be provided if the covering_date filter is provided.
	StartingAt *string `form:"starting_at" json:"starting_at,omitempty"`
}

// For a given package, returns all contract IDs and customer IDs associated with the package over a specific time period.
//
// ### Use this endpoint to:
// - Understand which customers are provisioned on a package at any given time for internal cohort management
// - Manage customer migrations to a new package. For example, to migrate all active customers to a new package, call this endpoint, end contracts, and provision customers on a new package.
//
// ### **Usage guidelines:**
// Use the **`starting_at`**, **`covering_date`**, and **`include_archived`** parameters to filter the list of returned contracts. For example, to list only currently active contracts, pass **`covering_date`** equal to the current time.
type V1ContractsListContractsOnPackageParams struct {
	Params `form:"*"`
	// Optional RFC 3339 timestamp. Only include contracts active on the provided date. This cannot be provided if starting_at filter is provided.
	CoveringDate *string `form:"covering_date" json:"covering_date,omitempty"`
	// Default false. Determines whether to include archived contracts in the results
	IncludeArchived *bool `form:"include_archived" json:"include_archived,omitempty"`
	// Max number of results that should be returned
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor that indicates where the next page of results should start.
	NextPage  *string `form:"next_page" json:"next_page,omitempty"`
	PackageID *string `form:"package_id" json:"package_id"`
	// Optional RFC 3339 timestamp. Only include contracts that started on or after this date. This cannot be provided if covering_date filter is provided.
	StartingAt *string `form:"starting_at" json:"starting_at,omitempty"`
}

// Each line requires an amount or both unit_price and quantity.
type V1ContractsScheduleProServicesInvoiceLineItemParams struct {
	// If the professional_service_id was added on an amendment, this is required.
	AmendmentID *string `form:"amendment_id" json:"amendment_id,omitempty"`
	// Amount for the term on the new invoice.
	Amount *Number `form:"amount" json:"amount,omitempty"`
	// For client use.
	Metadata *string `form:"metadata" json:"metadata,omitempty"`
	// The end date for the billing period on the invoice.
	NetsuiteInvoiceBillingEnd *string `form:"netsuite_invoice_billing_end" json:"netsuite_invoice_billing_end,omitempty"`
	// The start date for the billing period on the invoice.
	NetsuiteInvoiceBillingStart *string `form:"netsuite_invoice_billing_start" json:"netsuite_invoice_billing_start,omitempty"`
	ProfessionalServiceID       *string `form:"professional_service_id" json:"professional_service_id"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the amount.
	Quantity *Number `form:"quantity" json:"quantity,omitempty"`
	// If specified, this overrides the unit price on the pro service term. Must also provide quantity (but not amount) if providing unit_price.
	UnitPrice *Number `form:"unit_price" json:"unit_price,omitempty"`
}

// Create a new scheduled invoice for Professional Services terms on a contract. This endpoint's availability is dependent on your client's configuration.
type V1ContractsScheduleProServicesInvoiceParams struct {
	Params     `form:"*"`
	ContractID *string `form:"contract_id" json:"contract_id"`
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// The date the invoice is issued
	IssuedAt *string `form:"issued_at" json:"issued_at"`
	// Each line requires an amount or both unit_price and quantity.
	LineItems []*V1ContractsScheduleProServicesInvoiceLineItemParams `form:"line_items" json:"line_items"`
	// The end date of the invoice header in Netsuite
	NetsuiteInvoiceHeaderEnd *string `form:"netsuite_invoice_header_end" json:"netsuite_invoice_header_end,omitempty"`
	// The start date of the invoice header in Netsuite
	NetsuiteInvoiceHeaderStart *string `form:"netsuite_invoice_header_start" json:"netsuite_invoice_header_start,omitempty"`
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
type V1ContractsSetUsageFilterParams struct {
	Params      `form:"*"`
	ContractID  *string   `form:"contract_id" json:"contract_id"`
	CustomerID  *string   `form:"customer_id" json:"customer_id"`
	GroupKey    *string   `form:"group_key" json:"group_key"`
	GroupValues []*string `form:"group_values" json:"group_values"`
	StartingAt  *string   `form:"starting_at" json:"starting_at"`
}

// Update or add an end date to a contract. Ending a contract early will impact draft usage statements, truncate any terms, and remove upcoming scheduled invoices. Moving the date into the future will only extend the contract length. Terms and scheduled invoices are not extended. In-advance subscriptions will not be extended. Use this if a contract's end date has changed or if a perpetual contract ends.
type V1ContractsUpdateContractEndDateParams struct {
	Params `form:"*"`
	// If true, allows setting the contract end date earlier than the end_timestamp of existing finalized invoices. Finalized invoices will be unchanged; if you want to incorporate the new end date, you can void and regenerate finalized usage invoices. Defaults to true.
	AllowEndingBeforeFinalizedInvoice *bool `form:"allow_ending_before_finalized_invoice" json:"allow_ending_before_finalized_invoice,omitempty"`
	// ID of the contract to update
	ContractID *string `form:"contract_id" json:"contract_id"`
	// ID of the customer whose contract is to be updated
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// RFC 3339 timestamp indicating when the contract will end (exclusive). If not provided, the contract will be updated to be open-ended.
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
}

// Updates the issue date of a specific DRAFT invoice within a contract. Use this endpoint to reschedule when an invoice should be issued without affecting future billing cycles or the underlying contract terms. Only works with invoices still in DRAFT status, and the new issue date cannot be later than the contract's end date.
//
// ### Usage guidelines:
// This only changes the individual invoice's issue date - it does not modify the recurring invoice schedule of associated charges or commits. To update both the issue date and future billing schedule, use the 'edit contract' or 'edit commit' endpoints instead.
type V1ContractsUpdateInvoiceIssueDateParams struct {
	Params `form:"*"`
	// ID of the invoice to update. The invoice must still be in DRAFT status.
	InvoiceID *string `form:"invoice_id" json:"invoice_id"`
	// RFC 3339 timestamp. This will be the new issue date of the invoice. It must not be after the end date of the contract.
	IssueDate *string `form:"issue_date" json:"issue_date"`
}
