//
// File generated from our OpenAPI spec
//

package metronome

type V2ContractsEditContractAddCommitAccessScheduleScheduleItemParams struct {
	Amount *Number `form:"amount" json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	// RFC 3339 timestamp (inclusive)
	StartingAt *string `form:"starting_at" json:"starting_at"`
}

// Required: Schedule for distributing the commit to the customer. For "POSTPAID" commits only one schedule item is allowed and amount must match invoice_schedule total.
type V2ContractsEditContractAddCommitAccessScheduleParams struct {
	// Determines how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used. Defaults to `SPEND` if omitted.
	AccessType    *string                                                             `form:"access_type" json:"access_type,omitempty"`
	CreditTypeID  *string                                                             `form:"credit_type_id" json:"credit_type_id,omitempty"`
	ScheduleItems []*V2ContractsEditContractAddCommitAccessScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items"`
}

// Either provide amount or provide both unit_price and quantity.
type V2ContractsEditContractAddCommitInvoiceScheduleScheduleItemParams struct {
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
type V2ContractsEditContractAddCommitInvoiceScheduleRecurringScheduleParams struct {
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
type V2ContractsEditContractAddCommitInvoiceScheduleParams struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is only applicable to commit invoice schedules. If true, this schedule will not generate an invoice.
	DoNotInvoice *bool `form:"do_not_invoice" json:"do_not_invoice,omitempty"`
	// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
	RecurringSchedule *V2ContractsEditContractAddCommitInvoiceScheduleRecurringScheduleParams `form:"recurring_schedule" json:"recurring_schedule,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []*V2ContractsEditContractAddCommitInvoiceScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items,omitempty"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V2ContractsEditContractAddCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
type V2ContractsEditContractAddCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V2ContractsEditContractAddCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Only applicable if using STRIPE as your payment gateway type.
type V2ContractsEditContractAddCommitPaymentGateConfigStripeConfigParams struct {
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as your payment type.
	InvoiceMetadata map[string]string `form:"invoice_metadata" json:"invoice_metadata,omitempty"`
	// If true, the payment will be made assuming the customer is present (i.e. on session).
	//
	// If false, the payment will be made assuming the customer is not present (i.e. off session).
	// For cardholders from a country with an e-mandate requirement (e.g. India), the payment may be declined.
	//
	// If left blank, will default to false.
	OnSessionPayment *bool `form:"on_session_payment" json:"on_session_payment,omitempty"`
	// If left blank, will default to INVOICE
	PaymentType *string `form:"payment_type" json:"payment_type"`
}

// Only applicable if using PRECALCULATED as your tax type.
type V2ContractsEditContractAddCommitPaymentGateConfigPrecalculatedTaxConfigParams struct {
	// Amount of tax to be applied. This should be in the same currency and denomination as the commit's invoice schedule
	TaxAmount *Number `form:"tax_amount" json:"tax_amount"`
	// Name of the tax to be applied. This may be used in an invoice line item description.
	TaxName *string `form:"tax_name" json:"tax_name,omitempty"`
}

// optionally payment gate this commit
type V2ContractsEditContractAddCommitPaymentGateConfigParams struct {
	// Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.
	PaymentGateType *string `form:"payment_gate_type" json:"payment_gate_type"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig *V2ContractsEditContractAddCommitPaymentGateConfigPrecalculatedTaxConfigParams `form:"precalculated_tax_config" json:"precalculated_tax_config,omitempty"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig *V2ContractsEditContractAddCommitPaymentGateConfigStripeConfigParams `form:"stripe_config" json:"stripe_config,omitempty"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType *string `form:"tax_type" json:"tax_type,omitempty"`
}

// Optional configuration for commit hierarchy access control
type V2ContractsEditContractAddCommitHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}

// Optional attributes for spend tracker integration. Immutable after creation.
type V2ContractsEditContractAddCommitSpendTrackerAttributesParams struct {
	// If true, this commit will be included in spend trackers with discounted set to DISCOUNTED_ONLY
	CountsAsDiscounted *bool `form:"counts_as_discounted" json:"counts_as_discounted"`
}
type V2ContractsEditContractAddCommitParams struct {
	// Required: Schedule for distributing the commit to the customer. For "POSTPAID" commits only one schedule item is allowed and amount must match invoice_schedule total.
	AccessSchedule *V2ContractsEditContractAddCommitAccessScheduleParams `form:"access_schedule" json:"access_schedule,omitempty"`
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
	HierarchyConfiguration *V2ContractsEditContractAddCommitHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// Required for "POSTPAID" commits: the true up invoice will be generated at this time and only one schedule item is allowed; the total must match access_schedule amount. Optional for "PREPAID" commits: if not provided, this will be a "complimentary" commit with no invoice.
	InvoiceSchedule *V2ContractsEditContractAddCommitInvoiceScheduleParams `form:"invoice_schedule" json:"invoice_schedule,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// optionally payment gate this commit
	PaymentGateConfig *V2ContractsEditContractAddCommitPaymentGateConfigParams `form:"payment_gate_config" json:"payment_gate_config,omitempty"`
	// If multiple commits are applicable, the one with the lower priority will apply first.
	Priority  *Number `form:"priority" json:"priority,omitempty"`
	ProductID *string `form:"product_id" json:"product_id"`
	RateType  *string `form:"rate_type" json:"rate_type,omitempty"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
	Specifiers []*V2ContractsEditContractAddCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
	// Optional attributes for spend tracker integration. Immutable after creation.
	SpendTrackerAttributes *V2ContractsEditContractAddCommitSpendTrackerAttributesParams `form:"spend_tracker_attributes" json:"spend_tracker_attributes,omitempty"`
	// A temporary ID for the commit that can be used to reference the commit for commit specific overrides.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
	Type        *string `form:"type" json:"type"`
}
type V2ContractsEditContractAddCreditAccessScheduleScheduleItemParams struct {
	Amount *Number `form:"amount" json:"amount"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	// RFC 3339 timestamp (inclusive)
	StartingAt *string `form:"starting_at" json:"starting_at"`
}

// Schedule for distributing the credit to the customer.
type V2ContractsEditContractAddCreditAccessScheduleParams struct {
	// Determines how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used. Defaults to `SPEND` if omitted.
	AccessType    *string                                                             `form:"access_type" json:"access_type,omitempty"`
	CreditTypeID  *string                                                             `form:"credit_type_id" json:"credit_type_id,omitempty"`
	ScheduleItems []*V2ContractsEditContractAddCreditAccessScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V2ContractsEditContractAddCreditSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
type V2ContractsEditContractAddCreditSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V2ContractsEditContractAddCreditSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
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
type V2ContractsEditContractAddCreditHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}
type V2ContractsEditContractAddCreditParams struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule *V2ContractsEditContractAddCreditAccessScheduleParams `form:"access_schedule" json:"access_schedule"`
	// Which products the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductTags []*string         `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	CustomFields          map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// Used only in UI/API. It is not exposed to end customers.
	Description *string `form:"description" json:"description,omitempty"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration *V2ContractsEditContractAddCreditHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
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
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
	Specifiers []*V2ContractsEditContractAddCreditSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
}

// The amount of commit to grant.
type V2ContractsEditContractAddRecurringCommitAccessAmountParams struct {
	// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is required unless a subscription is attached via `subscription_config`.
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	UnitPrice *Number `form:"unit_price" json:"unit_price"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V2ContractsEditContractAddRecurringCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
type V2ContractsEditContractAddRecurringCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V2ContractsEditContractAddRecurringCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
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
type V2ContractsEditContractAddRecurringCommitCommitDurationParams struct {
	Unit  *string `form:"unit" json:"unit,omitempty"`
	Value *Number `form:"value" json:"value"`
}

// Optional configuration for recurring credit hierarchy access control
type V2ContractsEditContractAddRecurringCommitHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}
type V2ContractsEditContractAddRecurringCommitSubscriptionConfigApplySeatIncreaseConfigParams struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated *bool `form:"is_prorated" json:"is_prorated"`
}

// Attach a subscription to the recurring commit/credit.
type V2ContractsEditContractAddRecurringCommitSubscriptionConfigParams struct {
	// Controls when a customer can access recurring commit or credit child balances when the attached subscription is payment-gated. `BILLING_PERIOD_PAID` releases each balance after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases all balances after the first payment.
	AccessPolicy            *string                                                                                   `form:"access_policy" json:"access_policy,omitempty"`
	Allocation              *string                                                                                   `form:"allocation" json:"allocation,omitempty"`
	ApplySeatIncreaseConfig *V2ContractsEditContractAddRecurringCommitSubscriptionConfigApplySeatIncreaseConfigParams `form:"apply_seat_increase_config" json:"apply_seat_increase_config"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID *string `form:"subscription_id" json:"subscription_id"`
}

// The amount the customer should be billed for the commit. Not required.
type V2ContractsEditContractAddRecurringCommitInvoiceAmountParams struct {
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id"`
	Quantity     *Number `form:"quantity" json:"quantity"`
	UnitPrice    *Number `form:"unit_price" json:"unit_price"`
}
type V2ContractsEditContractAddRecurringCommitProrationRoundingAccessParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}
type V2ContractsEditContractAddRecurringCommitProrationRoundingInvoiceParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}

// Optional rounding configuration for prorated recurring commit amounts.
type V2ContractsEditContractAddRecurringCommitProrationRoundingParams struct {
	Access  *V2ContractsEditContractAddRecurringCommitProrationRoundingAccessParams  `form:"access" json:"access,omitempty"`
	Invoice *V2ContractsEditContractAddRecurringCommitProrationRoundingInvoiceParams `form:"invoice" json:"invoice,omitempty"`
}
type V2ContractsEditContractAddRecurringCommitParams struct {
	// The amount of commit to grant.
	AccessAmount *V2ContractsEditContractAddRecurringCommitAccessAmountParams `form:"access_amount" json:"access_amount"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Will be passed down to the individual commits
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	// Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.
	CommitDuration *V2ContractsEditContractAddRecurringCommitCommitDurationParams `form:"commit_duration" json:"commit_duration"`
	// Will be passed down to the individual commits
	Description *string `form:"description" json:"description,omitempty"`
	// Determines when the contract will stop creating recurring commits. optional
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration *V2ContractsEditContractAddRecurringCommitHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount *V2ContractsEditContractAddRecurringCommitInvoiceAmountParams `form:"invoice_amount" json:"invoice_amount,omitempty"`
	// displayed on invoices. will be passed through to the individual commits
	Name *string `form:"name" json:"name,omitempty"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// Will be passed down to the individual commits
	Priority  *Number `form:"priority" json:"priority"`
	ProductID *string `form:"product_id" json:"product_id"`
	// Determines whether the first and last commit will be prorated. If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration *string `form:"proration" json:"proration,omitempty"`
	// Optional rounding configuration for prorated recurring commit amounts.
	ProrationRounding *V2ContractsEditContractAddRecurringCommitProrationRoundingParams `form:"proration_rounding" json:"proration_rounding,omitempty"`
	// Whether the created commits will use the commit rate or list rate
	RateType *string `form:"rate_type" json:"rate_type,omitempty"`
	// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
	RecurrenceFrequency *string `form:"recurrence_frequency" json:"recurrence_frequency,omitempty"`
	// Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
	Specifiers []*V2ContractsEditContractAddRecurringCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
	// determines the start time for the first commit
	StartingAt *string `form:"starting_at" json:"starting_at"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig *V2ContractsEditContractAddRecurringCommitSubscriptionConfigParams `form:"subscription_config" json:"subscription_config,omitempty"`
	// A temporary ID that can be used to reference the recurring commit for commit specific overrides.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
}

// The amount of commit to grant.
type V2ContractsEditContractAddRecurringCreditAccessAmountParams struct {
	// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Defaults to USD (cents) if not passed
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is required unless a subscription is attached via `subscription_config`.
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	UnitPrice *Number `form:"unit_price" json:"unit_price"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V2ContractsEditContractAddRecurringCreditSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
type V2ContractsEditContractAddRecurringCreditSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V2ContractsEditContractAddRecurringCreditSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
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
type V2ContractsEditContractAddRecurringCreditCommitDurationParams struct {
	Unit  *string `form:"unit" json:"unit,omitempty"`
	Value *Number `form:"value" json:"value"`
}

// Optional configuration for recurring credit hierarchy access control
type V2ContractsEditContractAddRecurringCreditHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}
type V2ContractsEditContractAddRecurringCreditSubscriptionConfigApplySeatIncreaseConfigParams struct {
	// Indicates whether a mid-period seat increase should be prorated.
	IsProrated *bool `form:"is_prorated" json:"is_prorated"`
}

// Attach a subscription to the recurring commit/credit.
type V2ContractsEditContractAddRecurringCreditSubscriptionConfigParams struct {
	// Controls when a customer can access recurring commit or credit child balances when the attached subscription is payment-gated. `BILLING_PERIOD_PAID` releases each balance after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases all balances after the first payment.
	AccessPolicy            *string                                                                                   `form:"access_policy" json:"access_policy,omitempty"`
	Allocation              *string                                                                                   `form:"allocation" json:"allocation,omitempty"`
	ApplySeatIncreaseConfig *V2ContractsEditContractAddRecurringCreditSubscriptionConfigApplySeatIncreaseConfigParams `form:"apply_seat_increase_config" json:"apply_seat_increase_config"`
	// ID of the subscription to configure on the recurring commit/credit.
	SubscriptionID *string `form:"subscription_id" json:"subscription_id"`
}
type V2ContractsEditContractAddRecurringCreditProrationRoundingAccessParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}

// Optional rounding configuration for prorated recurring credit amounts.
type V2ContractsEditContractAddRecurringCreditProrationRoundingParams struct {
	Access *V2ContractsEditContractAddRecurringCreditProrationRoundingAccessParams `form:"access" json:"access,omitempty"`
}
type V2ContractsEditContractAddRecurringCreditParams struct {
	// The amount of commit to grant.
	AccessAmount *V2ContractsEditContractAddRecurringCreditAccessAmountParams `form:"access_amount" json:"access_amount"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Will be passed down to the individual commits
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	// Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to "PERIODS", where the length of a period is determined by the recurrence_frequency.
	CommitDuration *V2ContractsEditContractAddRecurringCreditCommitDurationParams `form:"commit_duration" json:"commit_duration"`
	// Will be passed down to the individual commits
	Description *string `form:"description" json:"description,omitempty"`
	// Determines when the contract will stop creating recurring commits. optional
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration *V2ContractsEditContractAddRecurringCreditHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// displayed on invoices. will be passed through to the individual commits
	Name *string `form:"name" json:"name,omitempty"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID *string `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	// Will be passed down to the individual commits
	Priority  *Number `form:"priority" json:"priority"`
	ProductID *string `form:"product_id" json:"product_id"`
	// Determines whether the first and last commit will be prorated. If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration *string `form:"proration" json:"proration,omitempty"`
	// Optional rounding configuration for prorated recurring credit amounts.
	ProrationRounding *V2ContractsEditContractAddRecurringCreditProrationRoundingParams `form:"proration_rounding" json:"proration_rounding,omitempty"`
	// Whether the created commits will use the commit rate or list rate
	RateType *string `form:"rate_type" json:"rate_type,omitempty"`
	// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
	RecurrenceFrequency *string `form:"recurrence_frequency" json:"recurrence_frequency,omitempty"`
	// Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
	Specifiers []*V2ContractsEditContractAddRecurringCreditSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
	// determines the start time for the first commit
	StartingAt *string `form:"starting_at" json:"starting_at"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig *V2ContractsEditContractAddRecurringCreditSubscriptionConfigParams `form:"subscription_config" json:"subscription_config,omitempty"`
	// A temporary ID that can be used to reference the recurring commit for commit specific overrides.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
}

// Either provide amount or provide both unit_price and quantity.
type V2ContractsEditContractAddDiscountScheduleScheduleItemParams struct {
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
type V2ContractsEditContractAddDiscountScheduleRecurringScheduleParams struct {
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
type V2ContractsEditContractAddDiscountScheduleParams struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is only applicable to commit invoice schedules. If true, this schedule will not generate an invoice.
	DoNotInvoice *bool `form:"do_not_invoice" json:"do_not_invoice,omitempty"`
	// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
	RecurringSchedule *V2ContractsEditContractAddDiscountScheduleRecurringScheduleParams `form:"recurring_schedule" json:"recurring_schedule,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []*V2ContractsEditContractAddDiscountScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items,omitempty"`
}
type V2ContractsEditContractAddDiscountParams struct {
	CustomFields map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string                                           `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	ProductID            *string                                           `form:"product_id" json:"product_id"`
	Schedule             *V2ContractsEditContractAddDiscountScheduleParams `form:"schedule" json:"schedule"`
}

// Only set for TIERED rate_type.
type V2ContractsEditContractAddOverrideOverwriteRateTierParams struct {
	Price *Number `form:"price" json:"price"`
	Size  *Number `form:"size" json:"size,omitempty"`
}
type V2ContractsEditContractAddOverrideOverwriteRateMinimumConfigParams struct {
	Minimum *Number `form:"minimum" json:"minimum"`
}

// Required for OVERWRITE type.
type V2ContractsEditContractAddOverrideOverwriteRateParams struct {
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate processors.
	CustomRate map[string]JSON `form:"custom_rate" json:"custom_rate,omitempty"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be set to true.
	IsProrated    *bool                                                               `form:"is_prorated" json:"is_prorated,omitempty"`
	MinimumConfig *V2ContractsEditContractAddOverrideOverwriteRateMinimumConfigParams `form:"minimum_config" json:"minimum_config,omitempty"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price *Number `form:"price" json:"price,omitempty"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity *Number `form:"quantity" json:"quantity,omitempty"`
	RateType *string `form:"rate_type" json:"rate_type"`
	// Only set for TIERED rate_type.
	Tiers []*V2ContractsEditContractAddOverrideOverwriteRateTierParams `form:"tiers" json:"tiers,omitempty"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V2ContractsEditContractAddOverrideOverrideSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// Cannot be used in conjunction with product_id or applicable_product_tags. If provided, the override will apply to all products with the specified specifiers.
type V2ContractsEditContractAddOverrideOverrideSpecifierParams struct {
	// Can only be used for commit specific overrides. Must be used in conjunction with one of `product_id`, `product_tags`, `pricing_group_values`, or `presentation_group_values`. Must be used instead of both `commit_ids` and `recurring_commit_ids` If provided, the override will apply to any specified commit, credit, recurring commit or recurring credit IDs.
	AnyCommitOrCreditIDs []*string `form:"any_commit_or_credit_ids" json:"any_commit_or_credit_ids,omitempty"`
	BillingFrequency     *string   `form:"billing_frequency" json:"billing_frequency,omitempty"`
	// If provided, the override will only apply to the specified commits. Can only be used for commit specific overrides. If not provided, the override will apply to all commits.
	CommitIDs []*string `form:"commit_ids" json:"commit_ids,omitempty"`
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V2ContractsEditContractAddOverrideOverrideSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// A map of group names to values. The override will only apply to line items with the specified presentation group values. Can only be used for multiplier overrides.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// A map of pricing group names to values. The override will only apply to products with the specified pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the override will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the override will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
	// Can only be used for commit specific overrides. Must be used in conjunction with one of product_id, product_tags, pricing_group_values, or presentation_group_values. If provided, the override will only apply to commits created by the specified recurring commit ids.
	RecurringCommitIDs []*string `form:"recurring_commit_ids" json:"recurring_commit_ids,omitempty"`
}

// Required for TIERED type. Must have at least one tier.
type V2ContractsEditContractAddOverrideTierParams struct {
	Multiplier *Number `form:"multiplier" json:"multiplier"`
	Size       *Number `form:"size" json:"size,omitempty"`
}
type V2ContractsEditContractAddOverrideParams struct {
	// tags identifying products whose rates are being overridden
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	Entitled     *bool   `form:"entitled" json:"entitled,omitempty"`
	// Indicates whether the override should only apply to commits. Defaults to `false`. If `true`, you can specify relevant commits in `override_specifiers` by passing `commit_ids`.
	IsCommitSpecific *bool `form:"is_commit_specific" json:"is_commit_specific,omitempty"`
	// Required for MULTIPLIER type. Must be >=0.
	Multiplier *Number `form:"multiplier" json:"multiplier,omitempty"`
	// Cannot be used in conjunction with product_id or applicable_product_tags. If provided, the override will apply to all products with the specified specifiers.
	OverrideSpecifiers []*V2ContractsEditContractAddOverrideOverrideSpecifierParams `form:"override_specifiers" json:"override_specifiers,omitempty"`
	// Required for OVERWRITE type.
	OverwriteRate *V2ContractsEditContractAddOverrideOverwriteRateParams `form:"overwrite_rate" json:"overwrite_rate,omitempty"`
	// Required for EXPLICIT multiplier prioritization scheme and all TIERED overrides. Under EXPLICIT prioritization, overwrites are prioritized first, and then tiered and multiplier overrides are prioritized by their priority value (lowest first). Must be > 0.
	Priority *Number `form:"priority" json:"priority,omitempty"`
	// ID of the product whose rate is being overridden
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// RFC 3339 timestamp indicating when the override will start applying (inclusive)
	StartingAt *string `form:"starting_at" json:"starting_at"`
	// Indicates whether the override applies to commit rates or list rates. Can only be used for overrides that have `is_commit_specific` set to `true`. Defaults to `"LIST_RATE"`.
	Target *string `form:"target" json:"target,omitempty"`
	// Required for TIERED type. Must have at least one tier.
	Tiers []*V2ContractsEditContractAddOverrideTierParams `form:"tiers" json:"tiers,omitempty"`
	// Overwrites are prioritized over multipliers and tiered overrides.
	Type *string `form:"type" json:"type,omitempty"`
}

// Either provide amount or provide both unit_price and quantity.
type V2ContractsEditContractAddScheduledChargeScheduleScheduleItemParams struct {
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
type V2ContractsEditContractAddScheduledChargeScheduleRecurringScheduleParams struct {
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
type V2ContractsEditContractAddScheduledChargeScheduleParams struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// This field is only applicable to commit invoice schedules. If true, this schedule will not generate an invoice.
	DoNotInvoice *bool `form:"do_not_invoice" json:"do_not_invoice,omitempty"`
	// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
	RecurringSchedule *V2ContractsEditContractAddScheduledChargeScheduleRecurringScheduleParams `form:"recurring_schedule" json:"recurring_schedule,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []*V2ContractsEditContractAddScheduledChargeScheduleScheduleItemParams `form:"schedule_items" json:"schedule_items,omitempty"`
}
type V2ContractsEditContractAddScheduledChargeParams struct {
	CustomFields map[string]string `form:"custom_fields" json:"custom_fields,omitempty"`
	// displayed on invoices
	Name *string `form:"name" json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string                                                  `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	ProductID            *string                                                  `form:"product_id" json:"product_id"`
	Schedule             *V2ContractsEditContractAddScheduledChargeScheduleParams `form:"schedule" json:"schedule"`
}

// This field's availability is dependent on your client's configuration.
type V2ContractsEditContractAddProfessionalServiceParams struct {
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
type V2ContractsEditContractAddResellerRoyaltyAwsOptionsParams struct {
	AwsAccountNumber    *string `form:"aws_account_number" json:"aws_account_number,omitempty"`
	AwsOfferID          *string `form:"aws_offer_id" json:"aws_offer_id,omitempty"`
	AwsPayerReferenceID *string `form:"aws_payer_reference_id" json:"aws_payer_reference_id,omitempty"`
}
type V2ContractsEditContractAddResellerRoyaltyGcpOptionsParams struct {
	GcpAccountID *string `form:"gcp_account_id" json:"gcp_account_id,omitempty"`
	GcpOfferID   *string `form:"gcp_offer_id" json:"gcp_offer_id,omitempty"`
}
type V2ContractsEditContractAddResellerRoyaltyParams struct {
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductTags []*string                                                  `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	AwsOptions            *V2ContractsEditContractAddResellerRoyaltyAwsOptionsParams `form:"aws_options" json:"aws_options,omitempty"`
	// Use null to indicate that the existing end timestamp should be removed.
	EndingBefore          *string                                                    `form:"ending_before" json:"ending_before,omitempty"`
	Fraction              *Number                                                    `form:"fraction" json:"fraction,omitempty"`
	GcpOptions            *V2ContractsEditContractAddResellerRoyaltyGcpOptionsParams `form:"gcp_options" json:"gcp_options,omitempty"`
	NetsuiteResellerID    *string                                                    `form:"netsuite_reseller_id" json:"netsuite_reseller_id,omitempty"`
	ResellerContractValue *Number                                                    `form:"reseller_contract_value" json:"reseller_contract_value,omitempty"`
	ResellerType          *string                                                    `form:"reseller_type" json:"reseller_type"`
	StartingAt            *string                                                    `form:"starting_at" json:"starting_at,omitempty"`
}
type V2ContractsEditContractAddSubscriptionSubscriptionRateParams struct {
	// Frequency to bill subscription with. Together with product_id, must match existing rate on the rate card.
	BillingFrequency *string `form:"billing_frequency" json:"billing_frequency"`
	// Must be subscription type product
	ProductID *string `form:"product_id" json:"product_id"`
}
type V2ContractsEditContractAddSubscriptionProrationRoundingParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}
type V2ContractsEditContractAddSubscriptionProrationParams struct {
	// Indicates how mid-period quantity adjustments are invoiced.
	// **BILL_IMMEDIATELY**: Only available when collection schedule is `ADVANCE`. The quantity increase will be billed immediately on the scheduled date.
	// **BILL_ON_NEXT_COLLECTION_DATE**: The quantity increase will be billed for in-arrears at the end of the period.
	//
	InvoiceBehavior *string `form:"invoice_behavior" json:"invoice_behavior,omitempty"`
	// Indicates if the partial period will be prorated or charged a full amount.
	IsProrated *bool                                                          `form:"is_prorated" json:"is_prorated,omitempty"`
	Rounding   *V2ContractsEditContractAddSubscriptionProrationRoundingParams `form:"rounding" json:"rounding,omitempty"`
}
type V2ContractsEditContractAddSubscriptionSeatConfigParams struct {
	// The initial assigned seats on this subscription.
	InitialSeatIDs []*string `form:"initial_seat_ids" json:"initial_seat_ids"`
	// The initial amount of unassigned seats on this subscription.
	InitialUnassignedSeats *Number `form:"initial_unassigned_seats" json:"initial_unassigned_seats,omitempty"`
	// The property name, sent on usage events, that identifies the seat ID associated with the usage event.  For example, the property name might be seat_id or user_id. The property must be set as a group key on billable metrics and a presentation/pricing group key on contract products.  This allows linked recurring credits with an allocation per seat to be consumed by only one seat's usage.
	SeatGroupKey *string `form:"seat_group_key" json:"seat_group_key"`
}
type V2ContractsEditContractAddSubscriptionBillingCycleConfigParams struct {
	// The date to anchor the billing cycle to. If omitted, defaults to the contract's usage invoice billing cycle anchor date.
	AnchorDate *string `form:"anchor_date" json:"anchor_date,omitempty"`
	// Controls whether this subscription consolidates onto usage invoices or gets its own scheduled invoice. Defaults to ON_USAGE_INVOICE if omitted.
	InvoicePlacement *string `form:"invoice_placement" json:"invoice_placement,omitempty"`
}
type V2ContractsEditContractAddSubscriptionPaymentGateConfigParams struct {
	// Gate access to the subscription based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe.
	PaymentGateType *string `form:"payment_gate_type" json:"payment_gate_type"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType *string `form:"tax_type" json:"tax_type,omitempty"`
}
type V2ContractsEditContractAddSubscriptionParams struct {
	BillingCycleConfig *V2ContractsEditContractAddSubscriptionBillingCycleConfigParams `form:"billing_cycle_config" json:"billing_cycle_config,omitempty"`
	CollectionSchedule *string                                                         `form:"collection_schedule" json:"collection_schedule"`
	CustomFields       map[string]string                                               `form:"custom_fields" json:"custom_fields,omitempty"`
	Description        *string                                                         `form:"description" json:"description,omitempty"`
	// Exclusive end time for the subscription. If not provided, subscription inherits contract end date.
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// The initial quantity for the subscription. It must be non-negative value. Required if quantity_management_mode is QUANTITY_ONLY.
	InitialQuantity   *Number                                                        `form:"initial_quantity" json:"initial_quantity,omitempty"`
	Name              *string                                                        `form:"name" json:"name,omitempty"`
	PaymentGateConfig *V2ContractsEditContractAddSubscriptionPaymentGateConfigParams `form:"payment_gate_config" json:"payment_gate_config,omitempty"`
	Proration         *V2ContractsEditContractAddSubscriptionProrationParams         `form:"proration" json:"proration"`
	// Determines how the subscription's quantity is controlled. Defaults to QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified directly on the subscription. `initial_quantity` must be provided with this option. Compatible with recurring commits/credits that use POOLED allocation. **SEAT_BASED**: Use when you want to pass specific seat identifiers (e.g. add user_123) to increment and decrement a subscription quantity, rather than directly providing the quantity. You must use a **SEAT_BASED** subscription to use a linked recurring credit with an allocation per seat. `seat_config` must be provided with this option.
	QuantityManagementMode *string                                                 `form:"quantity_management_mode" json:"quantity_management_mode,omitempty"`
	SeatConfig             *V2ContractsEditContractAddSubscriptionSeatConfigParams `form:"seat_config" json:"seat_config,omitempty"`
	// Inclusive start time for the subscription. If not provided, defaults to contract start date
	StartingAt       *string                                                       `form:"starting_at" json:"starting_at,omitempty"`
	SubscriptionRate *V2ContractsEditContractAddSubscriptionSubscriptionRateParams `form:"subscription_rate" json:"subscription_rate"`
	// A temporary ID used to reference the subscription in recurring commit/credit subscription configs created within the same payload.
	TemporaryID *string `form:"temporary_id" json:"temporary_id,omitempty"`
}
type V2ContractsEditContractAddSpendThresholdConfigurationCommitParams struct {
	Description *string `form:"description" json:"description,omitempty"`
	// Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.
	Name *string `form:"name" json:"name,omitempty"`
	// The priority of the commit, used to determine drawdown order. Lower priority commits are consumed first. Defaults to 100 if not specified.
	Priority *Number `form:"priority" json:"priority,omitempty"`
	// The commit product that will be used to generate the line item for commit payment.
	ProductID *string `form:"product_id" json:"product_id"`
}

// Only applicable if using STRIPE as your payment gateway type.
type V2ContractsEditContractAddSpendThresholdConfigurationPaymentGateConfigStripeConfigParams struct {
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as your payment type.
	InvoiceMetadata map[string]string `form:"invoice_metadata" json:"invoice_metadata,omitempty"`
	// If left blank, will default to INVOICE
	PaymentType *string `form:"payment_type" json:"payment_type"`
}

// Only applicable if using PRECALCULATED as your tax type.
type V2ContractsEditContractAddSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams struct {
	// Amount of tax to be applied. This should be in the same currency and denomination  as the commit's invoice schedule
	TaxAmount *Number `form:"tax_amount" json:"tax_amount"`
	// Name of the tax to be applied. This may be used in an invoice line item description.
	TaxName *string `form:"tax_name" json:"tax_name,omitempty"`
}
type V2ContractsEditContractAddSpendThresholdConfigurationPaymentGateConfigParams struct {
	// Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.
	PaymentGateType *string `form:"payment_gate_type" json:"payment_gate_type"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig *V2ContractsEditContractAddSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams `form:"precalculated_tax_config" json:"precalculated_tax_config,omitempty"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig *V2ContractsEditContractAddSpendThresholdConfigurationPaymentGateConfigStripeConfigParams `form:"stripe_config" json:"stripe_config,omitempty"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish  Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType *string `form:"tax_type" json:"tax_type,omitempty"`
}

// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
type V2ContractsEditContractAddSpendThresholdConfigurationDiscountConfigurationCapParams struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount *Number `form:"amount" json:"amount"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias *string `form:"spend_tracker_alias" json:"spend_tracker_alias"`
}
type V2ContractsEditContractAddSpendThresholdConfigurationDiscountConfigurationParams struct {
	// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
	Cap *V2ContractsEditContractAddSpendThresholdConfigurationDiscountConfigurationCapParams `form:"cap" json:"cap,omitempty"`
	// The fraction of the original amount that the customer pays after applying the discount. For example, 0.85 means the customer pays 85% of the original amount (a 15% discount).
	PaymentFraction *Number `form:"payment_fraction" json:"payment_fraction"`
}
type V2ContractsEditContractAddSpendThresholdConfigurationParams struct {
	Commit                *V2ContractsEditContractAddSpendThresholdConfigurationCommitParams                `form:"commit" json:"commit"`
	DiscountConfiguration *V2ContractsEditContractAddSpendThresholdConfigurationDiscountConfigurationParams `form:"discount_configuration" json:"discount_configuration,omitempty"`
	// When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.
	IsEnabled         *bool                                                                         `form:"is_enabled" json:"is_enabled"`
	PaymentGateConfig *V2ContractsEditContractAddSpendThresholdConfigurationPaymentGateConfigParams `form:"payment_gate_config" json:"payment_gate_config"`
	// Specify the threshold amount for the contract. Each time the contract's usage hits this amount, a threshold charge will be initiated.
	ThresholdAmount *Number `form:"threshold_amount" json:"threshold_amount"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
type V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// The length of time the created commit will be valid, starting from the end of the invoice's service period. If not provided, defaults to one year.
type V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationCommitDurationParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}
type V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationCommitParams struct {
	// Which products the threshold commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the threshold commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	Description           *string   `form:"description" json:"description,omitempty"`
	// The length of time the created commit will be valid, starting from the end of the invoice's service period. If not provided, defaults to one year.
	Duration *V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationCommitDurationParams `form:"duration" json:"duration,omitempty"`
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
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
	Specifiers []*V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
}

// Only applicable if using STRIPE as your payment gateway type.
type V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigParams struct {
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as your payment type.
	InvoiceMetadata map[string]string `form:"invoice_metadata" json:"invoice_metadata,omitempty"`
	// If left blank, will default to INVOICE
	PaymentType *string `form:"payment_type" json:"payment_type"`
}

// Only applicable if using PRECALCULATED as your tax type.
type V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams struct {
	// Amount of tax to be applied. This should be in the same currency and denomination  as the commit's invoice schedule
	TaxAmount *Number `form:"tax_amount" json:"tax_amount"`
	// Name of the tax to be applied. This may be used in an invoice line item description.
	TaxName *string `form:"tax_name" json:"tax_name,omitempty"`
}
type V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationPaymentGateConfigParams struct {
	// Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.
	PaymentGateType *string `form:"payment_gate_type" json:"payment_gate_type"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig *V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams `form:"precalculated_tax_config" json:"precalculated_tax_config,omitempty"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig *V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigParams `form:"stripe_config" json:"stripe_config,omitempty"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish  Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType *string `form:"tax_type" json:"tax_type,omitempty"`
}

// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
type V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationDiscountConfigurationCapParams struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount *Number `form:"amount" json:"amount"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias *string `form:"spend_tracker_alias" json:"spend_tracker_alias"`
}
type V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationDiscountConfigurationParams struct {
	// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
	Cap *V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationDiscountConfigurationCapParams `form:"cap" json:"cap,omitempty"`
	// The fraction of the original amount that the customer pays after applying the discount. For example, 0.85 means the customer pays 85% of the original amount (a 15% discount).
	PaymentFraction *Number `form:"payment_fraction" json:"payment_fraction"`
}
type V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilterParams struct {
	Entity *string `form:"entity" json:"entity"`
	Key    *string `form:"key" json:"key"`
	Value  *string `form:"value" json:"value"`
}
type V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeParams struct {
	CustomFieldFilters []*V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilterParams `form:"custom_field_filters" json:"custom_field_filters"`
}

// Determines which balances are excluded from remaining balance calculation for threshold billing.
type V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierParams struct {
	Exclude []*V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeParams `form:"exclude" json:"exclude"`
}
type V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationParams struct {
	Commit *V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationCommitParams `form:"commit" json:"commit"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID    *string                                                                                    `form:"custom_credit_type_id" json:"custom_credit_type_id,omitempty"`
	DiscountConfiguration *V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationDiscountConfigurationParams `form:"discount_configuration" json:"discount_configuration,omitempty"`
	// When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.
	IsEnabled         *bool                                                                                  `form:"is_enabled" json:"is_enabled"`
	PaymentGateConfig *V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationPaymentGateConfigParams `form:"payment_gate_config" json:"payment_gate_config"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount *Number `form:"recharge_to_amount" json:"recharge_to_amount"`
	// Specify the threshold amount for the contract. Each time the contract's balance lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount *Number `form:"threshold_amount" json:"threshold_amount"`
	// Determines which balances are excluded from remaining balance calculation for threshold billing.
	ThresholdBalanceSpecifiers []*V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierParams `form:"threshold_balance_specifiers" json:"threshold_balance_specifiers,omitempty"`
}
type V2ContractsEditContractAddBillingProviderConfigurationUpdateBillingProviderConfigurationParams struct {
	BillingProvider                *string `form:"billing_provider" json:"billing_provider,omitempty"`
	BillingProviderConfigurationID *string `form:"billing_provider_configuration_id" json:"billing_provider_configuration_id,omitempty"`
	DeliveryMethod                 *string `form:"delivery_method" json:"delivery_method,omitempty"`
}

// Indicates when the billing provider will be active on the contract. Any charges accrued during the schedule will be billed to the indicated billing provider.
type V2ContractsEditContractAddBillingProviderConfigurationUpdateScheduleParams struct {
	// When the billing provider update will take effect.
	EffectiveAt *string `form:"effective_at" json:"effective_at"`
}

// Update the billing provider configuration on the contract. Currently only supports adding a billing provider configuration to a contract that does not already have one.
type V2ContractsEditContractAddBillingProviderConfigurationUpdateParams struct {
	BillingProviderConfiguration *V2ContractsEditContractAddBillingProviderConfigurationUpdateBillingProviderConfigurationParams `form:"billing_provider_configuration" json:"billing_provider_configuration"`
	// Indicates when the billing provider will be active on the contract. Any charges accrued during the schedule will be billed to the indicated billing provider.
	Schedule *V2ContractsEditContractAddBillingProviderConfigurationUpdateScheduleParams `form:"schedule" json:"schedule"`
}
type V2ContractsEditContractAddRevenueSystemConfigurationUpdateRevenueSystemConfigurationParams struct {
	DeliveryMethod               *string `form:"delivery_method" json:"delivery_method,omitempty"`
	Provider                     *string `form:"provider" json:"provider,omitempty"`
	RevenueSystemConfigurationID *string `form:"revenue_system_configuration_id" json:"revenue_system_configuration_id,omitempty"`
}
type V2ContractsEditContractAddRevenueSystemConfigurationUpdateScheduleParams struct {
	EffectiveAt *string `form:"effective_at" json:"effective_at"`
}

// Update the revenue system configuration on the contract. Currently only supports adding a revenue system configuration to a contract that does not already have one.
type V2ContractsEditContractAddRevenueSystemConfigurationUpdateParams struct {
	RevenueSystemConfiguration *V2ContractsEditContractAddRevenueSystemConfigurationUpdateRevenueSystemConfigurationParams `form:"revenue_system_configuration" json:"revenue_system_configuration"`
	Schedule                   *V2ContractsEditContractAddRevenueSystemConfigurationUpdateScheduleParams                   `form:"schedule" json:"schedule"`
}
type V2ContractsEditContractAddSpendTrackerApplicableSpendSpecifierParams struct {
	// Filter by whether the spend was discounted. Defaults to ANY if omitted.
	Discounted *string   `form:"discounted" json:"discounted,omitempty"`
	Sources    []*string `form:"sources" json:"sources"`
	SpendType  *string   `form:"spend_type" json:"spend_type"`
}

// Spend trackers to add to this contract. Aliases must be unique within a contract.
type V2ContractsEditContractAddSpendTrackerParams struct {
	// Human-readable identifier, unique per contract.
	Alias                     *string                                                                 `form:"alias" json:"alias"`
	ApplicableSpendSpecifiers []*V2ContractsEditContractAddSpendTrackerApplicableSpendSpecifierParams `form:"applicable_spend_specifiers" json:"applicable_spend_specifiers"`
	CreditTypeID              *string                                                                 `form:"credit_type_id" json:"credit_type_id"`
	ResetFrequency            *string                                                                 `form:"reset_frequency" json:"reset_frequency"`
}
type V2ContractsEditContractUpdateScheduledChargeInvoiceScheduleAddScheduleItemParams struct {
	Amount    *Number `form:"amount" json:"amount,omitempty"`
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	Timestamp *string `form:"timestamp" json:"timestamp"`
	UnitPrice *Number `form:"unit_price" json:"unit_price,omitempty"`
}
type V2ContractsEditContractUpdateScheduledChargeInvoiceScheduleUpdateScheduleItemParams struct {
	Amount    *Number `form:"amount" json:"amount,omitempty"`
	ID        *string `form:"id" json:"id"`
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	Timestamp *string `form:"timestamp" json:"timestamp,omitempty"`
	UnitPrice *Number `form:"unit_price" json:"unit_price,omitempty"`
}
type V2ContractsEditContractUpdateScheduledChargeInvoiceScheduleRemoveScheduleItemParams struct {
	ID *string `form:"id" json:"id"`
}
type V2ContractsEditContractUpdateScheduledChargeInvoiceScheduleParams struct {
	AddScheduleItems    []*V2ContractsEditContractUpdateScheduledChargeInvoiceScheduleAddScheduleItemParams    `form:"add_schedule_items" json:"add_schedule_items,omitempty"`
	RemoveScheduleItems []*V2ContractsEditContractUpdateScheduledChargeInvoiceScheduleRemoveScheduleItemParams `form:"remove_schedule_items" json:"remove_schedule_items,omitempty"`
	UpdateScheduleItems []*V2ContractsEditContractUpdateScheduledChargeInvoiceScheduleUpdateScheduleItemParams `form:"update_schedule_items" json:"update_schedule_items,omitempty"`
}
type V2ContractsEditContractUpdateScheduledChargeParams struct {
	InvoiceSchedule      *V2ContractsEditContractUpdateScheduledChargeInvoiceScheduleParams `form:"invoice_schedule" json:"invoice_schedule,omitempty"`
	NetsuiteSalesOrderID *string                                                            `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	ScheduledChargeID    *string                                                            `form:"scheduled_charge_id" json:"scheduled_charge_id"`
}
type V2ContractsEditContractUpdateCommitAccessScheduleAddScheduleItemParams struct {
	Amount       *Number `form:"amount" json:"amount"`
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	StartingAt   *string `form:"starting_at" json:"starting_at"`
}
type V2ContractsEditContractUpdateCommitAccessScheduleUpdateScheduleItemParams struct {
	Amount       *Number `form:"amount" json:"amount,omitempty"`
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	ID           *string `form:"id" json:"id"`
	StartingAt   *string `form:"starting_at" json:"starting_at,omitempty"`
}
type V2ContractsEditContractUpdateCommitAccessScheduleRemoveScheduleItemParams struct {
	ID *string `form:"id" json:"id"`
}
type V2ContractsEditContractUpdateCommitAccessScheduleParams struct {
	AddScheduleItems    []*V2ContractsEditContractUpdateCommitAccessScheduleAddScheduleItemParams    `form:"add_schedule_items" json:"add_schedule_items,omitempty"`
	RemoveScheduleItems []*V2ContractsEditContractUpdateCommitAccessScheduleRemoveScheduleItemParams `form:"remove_schedule_items" json:"remove_schedule_items,omitempty"`
	UpdateScheduleItems []*V2ContractsEditContractUpdateCommitAccessScheduleUpdateScheduleItemParams `form:"update_schedule_items" json:"update_schedule_items,omitempty"`
}
type V2ContractsEditContractUpdateCommitInvoiceScheduleAddScheduleItemParams struct {
	Amount    *Number `form:"amount" json:"amount,omitempty"`
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	Timestamp *string `form:"timestamp" json:"timestamp"`
	UnitPrice *Number `form:"unit_price" json:"unit_price,omitempty"`
}
type V2ContractsEditContractUpdateCommitInvoiceScheduleUpdateScheduleItemParams struct {
	Amount    *Number `form:"amount" json:"amount,omitempty"`
	ID        *string `form:"id" json:"id"`
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	Timestamp *string `form:"timestamp" json:"timestamp,omitempty"`
	UnitPrice *Number `form:"unit_price" json:"unit_price,omitempty"`
}
type V2ContractsEditContractUpdateCommitInvoiceScheduleRemoveScheduleItemParams struct {
	ID *string `form:"id" json:"id"`
}
type V2ContractsEditContractUpdateCommitInvoiceScheduleParams struct {
	AddScheduleItems    []*V2ContractsEditContractUpdateCommitInvoiceScheduleAddScheduleItemParams    `form:"add_schedule_items" json:"add_schedule_items,omitempty"`
	RemoveScheduleItems []*V2ContractsEditContractUpdateCommitInvoiceScheduleRemoveScheduleItemParams `form:"remove_schedule_items" json:"remove_schedule_items,omitempty"`
	UpdateScheduleItems []*V2ContractsEditContractUpdateCommitInvoiceScheduleUpdateScheduleItemParams `form:"update_schedule_items" json:"update_schedule_items,omitempty"`
}

// Optional configuration for commit hierarchy access control
type V2ContractsEditContractUpdateCommitHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}
type V2ContractsEditContractUpdateCommitParams struct {
	AccessSchedule *V2ContractsEditContractUpdateCommitAccessScheduleParams `form:"access_schedule" json:"access_schedule,omitempty"`
	// Which products the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	CommitID              *string   `form:"commit_id" json:"commit_id"`
	Description           *string   `form:"description" json:"description,omitempty"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration *V2ContractsEditContractUpdateCommitHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	InvoiceSchedule        *V2ContractsEditContractUpdateCommitInvoiceScheduleParams        `form:"invoice_schedule" json:"invoice_schedule,omitempty"`
	Name                   *string                                                          `form:"name" json:"name,omitempty"`
	NetsuiteSalesOrderID   *string                                                          `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	Priority               *Number                                                          `form:"priority" json:"priority,omitempty"`
	ProductID              *string                                                          `form:"product_id" json:"product_id,omitempty"`
	// If provided, updates the commit to use the specified rate type for current and future invoices. Previously finalized invoices will need to be voided and regenerated to reflect the rate type change.
	RateType         *string `form:"rate_type" json:"rate_type,omitempty"`
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
}
type V2ContractsEditContractUpdateCreditAccessScheduleAddScheduleItemParams struct {
	Amount       *Number `form:"amount" json:"amount"`
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	StartingAt   *string `form:"starting_at" json:"starting_at"`
}
type V2ContractsEditContractUpdateCreditAccessScheduleUpdateScheduleItemParams struct {
	Amount       *Number `form:"amount" json:"amount,omitempty"`
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	ID           *string `form:"id" json:"id"`
	StartingAt   *string `form:"starting_at" json:"starting_at,omitempty"`
}
type V2ContractsEditContractUpdateCreditAccessScheduleRemoveScheduleItemParams struct {
	ID *string `form:"id" json:"id"`
}
type V2ContractsEditContractUpdateCreditAccessScheduleParams struct {
	AddScheduleItems    []*V2ContractsEditContractUpdateCreditAccessScheduleAddScheduleItemParams    `form:"add_schedule_items" json:"add_schedule_items,omitempty"`
	RemoveScheduleItems []*V2ContractsEditContractUpdateCreditAccessScheduleRemoveScheduleItemParams `form:"remove_schedule_items" json:"remove_schedule_items,omitempty"`
	UpdateScheduleItems []*V2ContractsEditContractUpdateCreditAccessScheduleUpdateScheduleItemParams `form:"update_schedule_items" json:"update_schedule_items,omitempty"`
}

// Optional configuration for commit hierarchy access control
type V2ContractsEditContractUpdateCreditHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}
type V2ContractsEditContractUpdateCreditParams struct {
	AccessSchedule *V2ContractsEditContractUpdateCreditAccessScheduleParams `form:"access_schedule" json:"access_schedule,omitempty"`
	// Which products the credit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the credit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the credit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the credit applies to all products.
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	CreditID              *string   `form:"credit_id" json:"credit_id"`
	Description           *string   `form:"description" json:"description,omitempty"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration *V2ContractsEditContractUpdateCreditHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	Name                   *string                                                          `form:"name" json:"name,omitempty"`
	NetsuiteSalesOrderID   *string                                                          `form:"netsuite_sales_order_id" json:"netsuite_sales_order_id,omitempty"`
	Priority               *Number                                                          `form:"priority" json:"priority,omitempty"`
	ProductID              *string                                                          `form:"product_id" json:"product_id,omitempty"`
	// If provided, updates the credit to use the specified rate type for current and future invoices. Previously finalized invoices will need to be voided and regenerated to reflect the rate type change.
	RateType         *string `form:"rate_type" json:"rate_type,omitempty"`
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
}
type V2ContractsEditContractUpdateRecurringCommitAccessAmountParams struct {
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	UnitPrice *Number `form:"unit_price" json:"unit_price,omitempty"`
}
type V2ContractsEditContractUpdateRecurringCommitInvoiceAmountParams struct {
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	UnitPrice *Number `form:"unit_price" json:"unit_price,omitempty"`
}
type V2ContractsEditContractUpdateRecurringCommitProrationRoundingAccessParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}
type V2ContractsEditContractUpdateRecurringCommitProrationRoundingInvoiceParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}

// If provided, updates the rounding config on the recurring commit. Set to null to clear rounding. Omit to leave unchanged.
type V2ContractsEditContractUpdateRecurringCommitProrationRoundingParams struct {
	Access  *V2ContractsEditContractUpdateRecurringCommitProrationRoundingAccessParams  `form:"access" json:"access,omitempty"`
	Invoice *V2ContractsEditContractUpdateRecurringCommitProrationRoundingInvoiceParams `form:"invoice" json:"invoice,omitempty"`
}

// Edits to these recurring commits will only affect commits whose access schedules has not started. Expired commits, and commits with an active access schedule will remain unchanged.
type V2ContractsEditContractUpdateRecurringCommitParams struct {
	AccessAmount  *V2ContractsEditContractUpdateRecurringCommitAccessAmountParams  `form:"access_amount" json:"access_amount,omitempty"`
	EndingBefore  *string                                                          `form:"ending_before" json:"ending_before,omitempty"`
	InvoiceAmount *V2ContractsEditContractUpdateRecurringCommitInvoiceAmountParams `form:"invoice_amount" json:"invoice_amount,omitempty"`
	// If provided, updates the rounding config on the recurring commit. Set to null to clear rounding. Omit to leave unchanged.
	ProrationRounding *V2ContractsEditContractUpdateRecurringCommitProrationRoundingParams `form:"proration_rounding" json:"proration_rounding,omitempty"`
	// If provided, updates the recurring commit to use the specified rate type when generating future commits.
	RateType          *string `form:"rate_type" json:"rate_type,omitempty"`
	RecurringCommitID *string `form:"recurring_commit_id" json:"recurring_commit_id"`
}
type V2ContractsEditContractUpdateRecurringCreditAccessAmountParams struct {
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	UnitPrice *Number `form:"unit_price" json:"unit_price,omitempty"`
}
type V2ContractsEditContractUpdateRecurringCreditProrationRoundingAccessParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}

// If provided, updates the rounding config on the recurring credit. Set to null to clear rounding. Omit to leave unchanged.
type V2ContractsEditContractUpdateRecurringCreditProrationRoundingParams struct {
	Access *V2ContractsEditContractUpdateRecurringCreditProrationRoundingAccessParams `form:"access" json:"access,omitempty"`
}

// Edits to these recurring credits will only affect credits whose access schedules has not started. Expired credits, and credits with an active access schedule will remain unchanged.
type V2ContractsEditContractUpdateRecurringCreditParams struct {
	AccessAmount *V2ContractsEditContractUpdateRecurringCreditAccessAmountParams `form:"access_amount" json:"access_amount,omitempty"`
	EndingBefore *string                                                         `form:"ending_before" json:"ending_before,omitempty"`
	// If provided, updates the rounding config on the recurring credit. Set to null to clear rounding. Omit to leave unchanged.
	ProrationRounding *V2ContractsEditContractUpdateRecurringCreditProrationRoundingParams `form:"proration_rounding" json:"proration_rounding,omitempty"`
	// If provided, updates the recurring credit to use the specified rate type when generating future credits.
	RateType          *string `form:"rate_type" json:"rate_type,omitempty"`
	RecurringCreditID *string `form:"recurring_credit_id" json:"recurring_credit_id"`
}
type V2ContractsEditContractUpdateSubscriptionQuantityManagementModeUpdateSeatConfigParams struct {
	SeatGroupKey *string `form:"seat_group_key" json:"seat_group_key"`
}

// Update the subscription's quantity management mode from QUANTITY_ONLY to SEAT_BASED with the provided seat_group_key.
type V2ContractsEditContractUpdateSubscriptionQuantityManagementModeUpdateParams struct {
	QuantityManagementMode *string                                                                                `form:"quantity_management_mode" json:"quantity_management_mode"`
	SeatConfig             *V2ContractsEditContractUpdateSubscriptionQuantityManagementModeUpdateSeatConfigParams `form:"seat_config" json:"seat_config"`
}

// Quantity changes are applied on the effective date based on the order which they are sent. For example, if I scheduled the quantity to be 12 on May 21 and then scheduled a quantity delta change of -1, the result from that day would be 11.
type V2ContractsEditContractUpdateSubscriptionQuantityUpdateParams struct {
	// The new quantity for the subscription. Must be provided if quantity_delta is not provided. Must be non-negative.
	Quantity *Number `form:"quantity" json:"quantity,omitempty"`
	// The delta to add to the subscription's quantity. Must be provided if quantity is not provided. Can't be zero. It also can't result in a negative quantity on the subscription.
	QuantityDelta *Number `form:"quantity_delta" json:"quantity_delta,omitempty"`
	StartingAt    *string `form:"starting_at" json:"starting_at"`
}

// Adds seat IDs to the subscription.  If there are unassigned seats, the new seat IDs will fill these unassigned seats and not increase the total subscription quantity. Otherwise, if there are more new seat IDs than unassigned seats, the total subscription quantity will increase.
type V2ContractsEditContractUpdateSubscriptionSeatUpdatesAddSeatIDParams struct {
	SeatIDs []*string `form:"seat_ids" json:"seat_ids"`
	// Assigned seats will be added/removed starting at this date.
	StartingAt *string `form:"starting_at" json:"starting_at"`
}

// Removes seat IDs from the subscription, if possible.  If a seat ID is removed, the total subscription quantity will decrease. Otherwise, if the seat ID is not found on the subscription, this is a no-op.
type V2ContractsEditContractUpdateSubscriptionSeatUpdatesRemoveSeatIDParams struct {
	SeatIDs []*string `form:"seat_ids" json:"seat_ids"`
	// Assigned seats will be added/removed starting at this date.
	StartingAt *string `form:"starting_at" json:"starting_at"`
}

// Adds unassigned seats to the subscription. This will increase the total subscription quantity.
type V2ContractsEditContractUpdateSubscriptionSeatUpdatesAddUnassignedSeatParams struct {
	// The number of unassigned seats on the subscription will increase/decrease by this delta. Must be greater than 0.
	Quantity *Number `form:"quantity" json:"quantity"`
	// Unassigned seats will be updated starting at this date.
	StartingAt *string `form:"starting_at" json:"starting_at"`
}

// Removes unassigned seats from the subscription. This will decrease the total subscription quantity if there are are unassigned seats.
type V2ContractsEditContractUpdateSubscriptionSeatUpdatesRemoveUnassignedSeatParams struct {
	// The number of unassigned seats on the subscription will increase/decrease by this delta. Must be greater than 0.
	Quantity *Number `form:"quantity" json:"quantity"`
	// Unassigned seats will be updated starting at this date.
	StartingAt *string `form:"starting_at" json:"starting_at"`
}
type V2ContractsEditContractUpdateSubscriptionSeatUpdatesParams struct {
	// Adds seat IDs to the subscription.  If there are unassigned seats, the new seat IDs will fill these unassigned seats and not increase the total subscription quantity. Otherwise, if there are more new seat IDs than unassigned seats, the total subscription quantity will increase.
	AddSeatIDs []*V2ContractsEditContractUpdateSubscriptionSeatUpdatesAddSeatIDParams `form:"add_seat_ids" json:"add_seat_ids,omitempty"`
	// Adds unassigned seats to the subscription. This will increase the total subscription quantity.
	AddUnassignedSeats []*V2ContractsEditContractUpdateSubscriptionSeatUpdatesAddUnassignedSeatParams `form:"add_unassigned_seats" json:"add_unassigned_seats,omitempty"`
	// Removes seat IDs from the subscription, if possible.  If a seat ID is removed, the total subscription quantity will decrease. Otherwise, if the seat ID is not found on the subscription, this is a no-op.
	RemoveSeatIDs []*V2ContractsEditContractUpdateSubscriptionSeatUpdatesRemoveSeatIDParams `form:"remove_seat_ids" json:"remove_seat_ids,omitempty"`
	// Removes unassigned seats from the subscription. This will decrease the total subscription quantity if there are are unassigned seats.
	RemoveUnassignedSeats []*V2ContractsEditContractUpdateSubscriptionSeatUpdatesRemoveUnassignedSeatParams `form:"remove_unassigned_seats" json:"remove_unassigned_seats,omitempty"`
}
type V2ContractsEditContractUpdateSubscriptionProrationRoundingParams struct {
	// Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit).
	DecimalPlaces  *Number `form:"decimal_places" json:"decimal_places"`
	RoundingMethod *string `form:"rounding_method" json:"rounding_method"`
}

// Optional list of subscriptions to update.
type V2ContractsEditContractUpdateSubscriptionParams struct {
	EndingBefore      *string                                                           `form:"ending_before" json:"ending_before,omitempty"`
	Name              *string                                                           `form:"name" json:"name,omitempty"`
	ProrationRounding *V2ContractsEditContractUpdateSubscriptionProrationRoundingParams `form:"proration_rounding" json:"proration_rounding,omitempty"`
	// Update the subscription's quantity management mode from QUANTITY_ONLY to SEAT_BASED with the provided seat_group_key.
	QuantityManagementModeUpdate *V2ContractsEditContractUpdateSubscriptionQuantityManagementModeUpdateParams `form:"quantity_management_mode_update" json:"quantity_management_mode_update,omitempty"`
	// Quantity changes are applied on the effective date based on the order which they are sent. For example, if I scheduled the quantity to be 12 on May 21 and then scheduled a quantity delta change of -1, the result from that day would be 11.
	QuantityUpdates []*V2ContractsEditContractUpdateSubscriptionQuantityUpdateParams `form:"quantity_updates" json:"quantity_updates,omitempty"`
	SeatUpdates     *V2ContractsEditContractUpdateSubscriptionSeatUpdatesParams      `form:"seat_updates" json:"seat_updates,omitempty"`
	SubscriptionID  *string                                                          `form:"subscription_id" json:"subscription_id"`
}
type V2ContractsEditContractUpdateSpendThresholdConfigurationCommitParams struct {
	Description *string `form:"description" json:"description,omitempty"`
	// Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.
	Name *string `form:"name" json:"name,omitempty"`
	// The priority of the commit, used to determine drawdown order. Lower priority commits are consumed first. Defaults to 100 if not specified. On updates, set to null to clear a previously configured priority.
	Priority *Number `form:"priority" json:"priority,omitempty"`
	// The commit product that will be used to generate the line item for commit payment.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
}

// Only applicable if using STRIPE as your payment gateway type.
type V2ContractsEditContractUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigParams struct {
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as your payment type.
	InvoiceMetadata map[string]string `form:"invoice_metadata" json:"invoice_metadata,omitempty"`
	// If left blank, will default to INVOICE
	PaymentType *string `form:"payment_type" json:"payment_type"`
}

// Only applicable if using PRECALCULATED as your tax type.
type V2ContractsEditContractUpdateSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams struct {
	// Amount of tax to be applied. This should be in the same currency and denomination  as the commit's invoice schedule
	TaxAmount *Number `form:"tax_amount" json:"tax_amount"`
	// Name of the tax to be applied. This may be used in an invoice line item description.
	TaxName *string `form:"tax_name" json:"tax_name,omitempty"`
}
type V2ContractsEditContractUpdateSpendThresholdConfigurationPaymentGateConfigParams struct {
	// Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.
	PaymentGateType *string `form:"payment_gate_type" json:"payment_gate_type"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig *V2ContractsEditContractUpdateSpendThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams `form:"precalculated_tax_config" json:"precalculated_tax_config,omitempty"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig *V2ContractsEditContractUpdateSpendThresholdConfigurationPaymentGateConfigStripeConfigParams `form:"stripe_config" json:"stripe_config,omitempty"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish  Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType *string `form:"tax_type" json:"tax_type,omitempty"`
}

// Update the discount cap. Set to null to remove an existing cap.
type V2ContractsEditContractUpdateSpendThresholdConfigurationDiscountConfigurationCapParams struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount *Number `form:"amount" json:"amount"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias *string `form:"spend_tracker_alias" json:"spend_tracker_alias"`
}
type V2ContractsEditContractUpdateSpendThresholdConfigurationDiscountConfigurationParams struct {
	// Update the discount cap. Set to null to remove an existing cap.
	Cap *V2ContractsEditContractUpdateSpendThresholdConfigurationDiscountConfigurationCapParams `form:"cap" json:"cap,omitempty"`
	// The fraction of the original amount that the customer pays after applying the discount. Set to null to remove the discount fraction. For example, 0.85 means the customer pays 85% of the original amount (a 15% discount).
	//
	PaymentFraction *Number `form:"payment_fraction" json:"payment_fraction,omitempty"`
}
type V2ContractsEditContractUpdateSpendThresholdConfigurationParams struct {
	Commit                *V2ContractsEditContractUpdateSpendThresholdConfigurationCommitParams                `form:"commit" json:"commit,omitempty"`
	DiscountConfiguration *V2ContractsEditContractUpdateSpendThresholdConfigurationDiscountConfigurationParams `form:"discount_configuration" json:"discount_configuration,omitempty"`
	// When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.
	IsEnabled         *bool                                                                            `form:"is_enabled" json:"is_enabled,omitempty"`
	PaymentGateConfig *V2ContractsEditContractUpdateSpendThresholdConfigurationPaymentGateConfigParams `form:"payment_gate_config" json:"payment_gate_config,omitempty"`
	// Specify the threshold amount for the contract. Each time the contract's usage hits this amount, a threshold charge will be initiated.
	ThresholdAmount *Number `form:"threshold_amount" json:"threshold_amount,omitempty"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
type V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `form:"presentation_group_values" json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// The length of time the created commit will be valid, starting from the end of the invoice's service period. Set to null to clear a previously configured duration.
type V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationCommitDurationParams struct {
	Unit  *string `form:"unit" json:"unit"`
	Value *int64  `form:"value" json:"value"`
}
type V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationCommitParams struct {
	// Which products the threshold commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the threshold commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	Description           *string   `form:"description" json:"description,omitempty"`
	// The length of time the created commit will be valid, starting from the end of the invoice's service period. Set to null to clear a previously configured duration.
	Duration *V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationCommitDurationParams `form:"duration" json:"duration,omitempty"`
	// Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.
	Name *string `form:"name" json:"name,omitempty"`
	// The priority of the commit, used to determine drawdown order. Lower priority commits are consumed first. Defaults to 100 if not specified. On updates, set to null to clear a previously configured priority.
	Priority *Number `form:"priority" json:"priority,omitempty"`
	// The commit product that will be used to generate the line item for commit payment.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// Whether the created commits will be charged at commit rate or list rate. Set to null to clear a previously configured rate type.
	RateType *string `form:"rate_type" json:"rate_type,omitempty"`
	// Fraction of the created commit's unused balance that will roll over. Must be between 0 and 1. Set to null to clear a previously configured rollover fraction.
	RolloverFraction *Number `form:"rollover_fraction" json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
	Specifiers []*V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
}

// Only applicable if using STRIPE as your payment gateway type.
type V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigParams struct {
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as your payment type.
	InvoiceMetadata map[string]string `form:"invoice_metadata" json:"invoice_metadata,omitempty"`
	// If left blank, will default to INVOICE
	PaymentType *string `form:"payment_type" json:"payment_type"`
}

// Only applicable if using PRECALCULATED as your tax type.
type V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams struct {
	// Amount of tax to be applied. This should be in the same currency and denomination  as the commit's invoice schedule
	TaxAmount *Number `form:"tax_amount" json:"tax_amount"`
	// Name of the tax to be applied. This may be used in an invoice line item description.
	TaxName *string `form:"tax_name" json:"tax_name,omitempty"`
}
type V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigParams struct {
	// Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.
	PaymentGateType *string `form:"payment_gate_type" json:"payment_gate_type"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig *V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigPrecalculatedTaxConfigParams `form:"precalculated_tax_config" json:"precalculated_tax_config,omitempty"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig *V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigStripeConfigParams `form:"stripe_config" json:"stripe_config,omitempty"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish  Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType *string `form:"tax_type" json:"tax_type,omitempty"`
}

// Update the discount cap. Set to null to remove an existing cap.
type V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationDiscountConfigurationCapParams struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount *Number `form:"amount" json:"amount"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias *string `form:"spend_tracker_alias" json:"spend_tracker_alias"`
}
type V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationDiscountConfigurationParams struct {
	// Update the discount cap. Set to null to remove an existing cap.
	Cap *V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationDiscountConfigurationCapParams `form:"cap" json:"cap,omitempty"`
	// The fraction of the original amount that the customer pays after applying the discount. Set to null to remove the discount fraction. For example, 0.85 means the customer pays 85% of the original amount (a 15% discount).
	//
	PaymentFraction *Number `form:"payment_fraction" json:"payment_fraction,omitempty"`
}
type V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilterParams struct {
	Entity *string `form:"entity" json:"entity"`
	Key    *string `form:"key" json:"key"`
	Value  *string `form:"value" json:"value"`
}
type V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeParams struct {
	CustomFieldFilters []*V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeCustomFieldFilterParams `form:"custom_field_filters" json:"custom_field_filters"`
}

// Determines which balances are excluded from remaining balance calculation for threshold billing.
type V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierParams struct {
	Exclude []*V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierExcludeParams `form:"exclude" json:"exclude"`
}
type V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationParams struct {
	Commit *V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationCommitParams `form:"commit" json:"commit,omitempty"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID    *string                                                                                       `form:"custom_credit_type_id" json:"custom_credit_type_id,omitempty"`
	DiscountConfiguration *V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationDiscountConfigurationParams `form:"discount_configuration" json:"discount_configuration,omitempty"`
	// When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.
	IsEnabled         *bool                                                                                     `form:"is_enabled" json:"is_enabled,omitempty"`
	PaymentGateConfig *V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationPaymentGateConfigParams `form:"payment_gate_config" json:"payment_gate_config,omitempty"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount *Number `form:"recharge_to_amount" json:"recharge_to_amount,omitempty"`
	// Specify the threshold amount for the contract. Each time the contract's balance lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount *Number `form:"threshold_amount" json:"threshold_amount,omitempty"`
	// Determines which balances are excluded from remaining balance calculation for threshold billing.
	ThresholdBalanceSpecifiers []*V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationThresholdBalanceSpecifierParams `form:"threshold_balance_specifiers" json:"threshold_balance_specifiers,omitempty"`
}

// IDs of commits to archive
type V2ContractsEditContractArchiveCommitParams struct {
	ID *string `form:"id" json:"id"`
}

// IDs of credits to archive
type V2ContractsEditContractArchiveCreditParams struct {
	ID *string `form:"id" json:"id"`
}

// IDs of scheduled charges to archive
type V2ContractsEditContractArchiveScheduledChargeParams struct {
	ID *string `form:"id" json:"id"`
}

// IDs of overrides to remove
type V2ContractsEditContractRemoveOverrideParams struct {
	ID *string `form:"id" json:"id"`
}

// The ability to edit a contract helps you react quickly to the needs of your customers and your business.
//
// ### Use this endpoint to:
// - Encode mid-term commitment and discount changes
// - Fix configuration mistakes and easily roll back packaging changes
//
// ### Key response fields:
// - The `id` of the edit
// - Complete edit details. For example, if you edited the contract to add new overrides and credits, you will receive the IDs of those overrides and credits in the response.
//
// ### Usage guidelines:
// - When you edit a contract, any draft invoices update immediately to reflect that edit. Finalized invoices remain unchanged - you must void and regenerate them in the UI or API to reflect the edit.
// - Contract editing must be enabled to use this endpoint. Contact us via the [Metronome support portal](https://support.metronome.com/) to learn more.
type V2ContractsEditContractParams struct {
	Params `form:"*"`
	// Update the billing provider configuration on the contract. Currently only supports adding a billing provider configuration to a contract that does not already have one.
	AddBillingProviderConfigurationUpdate   *V2ContractsEditContractAddBillingProviderConfigurationUpdateParams   `form:"add_billing_provider_configuration_update" json:"add_billing_provider_configuration_update,omitempty"`
	AddCommits                              []*V2ContractsEditContractAddCommitParams                             `form:"add_commits" json:"add_commits,omitempty"`
	AddCredits                              []*V2ContractsEditContractAddCreditParams                             `form:"add_credits" json:"add_credits,omitempty"`
	AddDiscounts                            []*V2ContractsEditContractAddDiscountParams                           `form:"add_discounts" json:"add_discounts,omitempty"`
	AddOverrides                            []*V2ContractsEditContractAddOverrideParams                           `form:"add_overrides" json:"add_overrides,omitempty"`
	AddPrepaidBalanceThresholdConfiguration *V2ContractsEditContractAddPrepaidBalanceThresholdConfigurationParams `form:"add_prepaid_balance_threshold_configuration" json:"add_prepaid_balance_threshold_configuration,omitempty"`
	// This field's availability is dependent on your client's configuration.
	AddProfessionalServices []*V2ContractsEditContractAddProfessionalServiceParams `form:"add_professional_services" json:"add_professional_services,omitempty"`
	AddRecurringCommits     []*V2ContractsEditContractAddRecurringCommitParams     `form:"add_recurring_commits" json:"add_recurring_commits,omitempty"`
	AddRecurringCredits     []*V2ContractsEditContractAddRecurringCreditParams     `form:"add_recurring_credits" json:"add_recurring_credits,omitempty"`
	AddResellerRoyalties    []*V2ContractsEditContractAddResellerRoyaltyParams     `form:"add_reseller_royalties" json:"add_reseller_royalties,omitempty"`
	// Update the revenue system configuration on the contract. Currently only supports adding a revenue system configuration to a contract that does not already have one.
	AddRevenueSystemConfigurationUpdate *V2ContractsEditContractAddRevenueSystemConfigurationUpdateParams `form:"add_revenue_system_configuration_update" json:"add_revenue_system_configuration_update,omitempty"`
	AddScheduledCharges                 []*V2ContractsEditContractAddScheduledChargeParams                `form:"add_scheduled_charges" json:"add_scheduled_charges,omitempty"`
	AddSpendThresholdConfiguration      *V2ContractsEditContractAddSpendThresholdConfigurationParams      `form:"add_spend_threshold_configuration" json:"add_spend_threshold_configuration,omitempty"`
	// Spend trackers to add to this contract. Aliases must be unique within a contract.
	AddSpendTrackers []*V2ContractsEditContractAddSpendTrackerParams `form:"add_spend_trackers" json:"add_spend_trackers,omitempty"`
	AddSubscriptions []*V2ContractsEditContractAddSubscriptionParams `form:"add_subscriptions" json:"add_subscriptions,omitempty"`
	// If true, allows setting the contract end date earlier than the end_timestamp of existing finalized invoices. Finalized invoices will be unchanged; if you want to incorporate the new end date, you can void and regenerate finalized usage invoices. Defaults to true.
	AllowContractEndingBeforeFinalizedInvoice *bool `form:"allow_contract_ending_before_finalized_invoice" json:"allow_contract_ending_before_finalized_invoice,omitempty"`
	// IDs of commits to archive
	ArchiveCommits []*V2ContractsEditContractArchiveCommitParams `form:"archive_commits" json:"archive_commits,omitempty"`
	// IDs of credits to archive
	ArchiveCredits []*V2ContractsEditContractArchiveCreditParams `form:"archive_credits" json:"archive_credits,omitempty"`
	// IDs of scheduled charges to archive
	ArchiveScheduledCharges []*V2ContractsEditContractArchiveScheduledChargeParams `form:"archive_scheduled_charges" json:"archive_scheduled_charges,omitempty"`
	// Aliases of spend trackers to archive.
	ArchiveSpendTrackers []*string `form:"archive_spend_trackers" json:"archive_spend_trackers,omitempty"`
	// ID of the contract being edited
	ContractID *string `form:"contract_id" json:"contract_id"`
	// ID of the customer whose contract is being edited
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// IDs of overrides to remove
	RemoveOverrides []*V2ContractsEditContractRemoveOverrideParams `form:"remove_overrides" json:"remove_overrides,omitempty"`
	// Optional uniqueness key to prevent duplicate contract edits.
	UniquenessKey *string                                      `form:"uniqueness_key" json:"uniqueness_key,omitempty"`
	UpdateCommits []*V2ContractsEditContractUpdateCommitParams `form:"update_commits" json:"update_commits,omitempty"`
	// RFC 3339 timestamp indicating when the contract will end (exclusive).
	UpdateContractEndDate *string `form:"update_contract_end_date" json:"update_contract_end_date,omitempty"`
	// Value to update the contract name to. If not provided, the contract name will remain unchanged.
	UpdateContractName *string                                      `form:"update_contract_name" json:"update_contract_name,omitempty"`
	UpdateCredits      []*V2ContractsEditContractUpdateCreditParams `form:"update_credits" json:"update_credits,omitempty"`
	// Number of days after issuance of invoice after which the invoice is due (e.g. Net 30).
	UpdateNetPaymentTermsDays                  *Number                                                                  `form:"update_net_payment_terms_days" json:"update_net_payment_terms_days,omitempty"`
	UpdatePrepaidBalanceThresholdConfiguration *V2ContractsEditContractUpdatePrepaidBalanceThresholdConfigurationParams `form:"update_prepaid_balance_threshold_configuration" json:"update_prepaid_balance_threshold_configuration,omitempty"`
	// Edits to these recurring commits will only affect commits whose access schedules has not started. Expired commits, and commits with an active access schedule will remain unchanged.
	UpdateRecurringCommits []*V2ContractsEditContractUpdateRecurringCommitParams `form:"update_recurring_commits" json:"update_recurring_commits,omitempty"`
	// Edits to these recurring credits will only affect credits whose access schedules has not started. Expired credits, and credits with an active access schedule will remain unchanged.
	UpdateRecurringCredits            []*V2ContractsEditContractUpdateRecurringCreditParams           `form:"update_recurring_credits" json:"update_recurring_credits,omitempty"`
	UpdateScheduledCharges            []*V2ContractsEditContractUpdateScheduledChargeParams           `form:"update_scheduled_charges" json:"update_scheduled_charges,omitempty"`
	UpdateSpendThresholdConfiguration *V2ContractsEditContractUpdateSpendThresholdConfigurationParams `form:"update_spend_threshold_configuration" json:"update_spend_threshold_configuration,omitempty"`
	// Optional list of subscriptions to update.
	UpdateSubscriptions []*V2ContractsEditContractUpdateSubscriptionParams `form:"update_subscriptions" json:"update_subscriptions,omitempty"`
}

// Gets the details for a specific contract, including contract term, rate card information, credits and commits, and more.
//
// ### Use this endpoint to:
// - Check the duration of a customer's current contract
// - Get details on contract terms, including access schedule amounts for commitments and credits
// - Understand the state of a contract at a past time. As you can evolve the terms of a contract over time through editing, use the `as_of_date` parameter to view the full contract configuration as of that point in time.
//
// ### Usage guidelines:
// - Optionally, use the `include_balance` and `include_ledger` fields to include balances and ledgers in the credit and commit responses. Using these fields will cause the query to be slower.
type V2ContractsGetContractParams struct {
	Params `form:"*"`
	// Optional RFC 3339 timestamp. Return the contract as of this date. Cannot be used with include_ledgers parameter.
	AsOfDate   *string `form:"as_of_date" json:"as_of_date,omitempty"`
	ContractID *string `form:"contract_id" json:"contract_id"`
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// Include the balance of credits and commits in the response. Setting this flag may cause the query to be slower.
	IncludeBalance *bool `form:"include_balance" json:"include_balance,omitempty"`
	// Include commit/credit ledgers in the response. Setting this flag may cause the query to be slower. Cannot be used with as_of_date parameter.
	IncludeLedgers *bool `form:"include_ledgers" json:"include_ledgers,omitempty"`
	// Indicates that this API request was triggered by a webhook notification with the provided ID.
	WebhookNotificationID *string `form:"webhook_notification_id" json:"webhook_notification_id,omitempty"`
}

// List all the edits made to a contract over time. In Metronome, you can edit a contract at any point after it's created to fix mistakes or reflect changes in terms. Metronome stores a full history of all edits that were ever made to a contract, whether through the UI, `editContract` endpoint, or other endpoints like `updateContractEndDate`.
//
// ### Use this endpoint to:
// - Understand what changes were made to a contract, when, and by who
//
// ### Key response fields:
// - An array of every edit ever made to the contract
// - Details on each individual edit - for example showing that in one edit, a user added two discounts and incremented a subscription quantity.
type V2ContractsGetContractEditHistoryParams struct {
	Params     `form:"*"`
	ContractID *string `form:"contract_id" json:"contract_id"`
	CustomerID *string `form:"customer_id" json:"customer_id"`
}

// For a given customer, lists a page of their contracts in chronological order.
//
// ### Use this endpoint to:
// - Check if a customer is provisioned with any contract, and at which tier
// - Check the duration and terms of a customer's current contract
// - Power a page in your end customer experience that shows the customer's history of tiers (e.g. this customer started out on the Pro Plan, then downgraded to the Starter plan).
//
// ### Usage guidelines:
// Use the `starting_at`, `covering_date`, and `include_archived` parameters to filter the list of returned contracts. For example, to list only currently active contracts, pass `covering_date` equal to the current time.
//
// Results are limited to 20 contracts per page. When the response includes a non-null `cursor`, pass it back as the `cursor` parameter to fetch the next page.
type V2ContractsListContractsParams struct {
	Params `form:"*"`
	// Optional RFC 3339 timestamp. Only include contracts active on the provided date. This cannot be provided if starting_at filter is provided.
	CoveringDate *string `form:"covering_date" json:"covering_date,omitempty"`
	// Cursor from a previous response to fetch the next page of contracts.
	Cursor     *string `form:"cursor" json:"cursor,omitempty"`
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// Include archived contracts in the response.
	IncludeArchived *bool `form:"include_archived" json:"include_archived,omitempty"`
	// Include the balance of credits and commits in the response. Setting this flag may cause the response to be slower.
	IncludeBalance *bool `form:"include_balance" json:"include_balance,omitempty"`
	// Include commit/credit ledgers in the response. Setting this flag may cause the response to be slower.
	IncludeLedgers *bool `form:"include_ledgers" json:"include_ledgers,omitempty"`
	// Max number of contracts to return per page. Range: 1-20. Default: 20.
	Limit *Number `form:"limit" json:"limit,omitempty"`
	// Optional RFC 3339 timestamp. Only include contracts that started on or after this date. This cannot be provided if covering_date filter is provided.
	StartingAt *string `form:"starting_at" json:"starting_at,omitempty"`
}
