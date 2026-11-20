//
// File generated from our OpenAPI spec
//

package metronome

// Only set for TIERED rate_type.
type V1RateCardsAddRateTierParams struct {
	Price *Number `form:"price" json:"price"`
	Size  *Number `form:"size" json:"size,omitempty"`
}
type V1RateCardsAddRateMinimumConfigParams struct {
	Minimum *Number `form:"minimum" json:"minimum"`
}

// Only set for TIERED rate_type.
type V1RateCardsAddRateCommitRateTierParams struct {
	Price *Number `form:"price" json:"price"`
	Size  *Number `form:"size" json:"size,omitempty"`
}
type V1RateCardsAddRateCommitRateMinimumConfigParams struct {
	Minimum *Number `form:"minimum" json:"minimum"`
}
type V1RateCardsAddRateCommitRateParams struct {
	MinimumConfig *V1RateCardsAddRateCommitRateMinimumConfigParams `form:"minimum_config" json:"minimum_config,omitempty"`
	// Commit rate price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    *Number `form:"price" json:"price,omitempty"`
	RateType *string `form:"rate_type" json:"rate_type"`
	// Only set for TIERED rate_type.
	Tiers []*V1RateCardsAddRateCommitRateTierParams `form:"tiers" json:"tiers,omitempty"`
}

// Add a new rate
//
// This endpoint is heavily rate limited. For adding multiple rates, using the [addRates](https://docs.metronome.com/api-reference/rate-cards/add-rates) endpoint is strongly encouraged.
type V1RateCardsAddRateParams struct {
	Params `form:"*"`
	// Optional. Frequency to bill subscriptions with. Required for subscription type products with Flat rate.
	BillingFrequency *string                             `form:"billing_frequency" json:"billing_frequency,omitempty"`
	CommitRate       *V1RateCardsAddRateCommitRateParams `form:"commit_rate" json:"commit_rate,omitempty"`
	// The Metronome ID of the credit type to associate with price, defaults to USD (cents) if not passed. Used by all rate_types except type PERCENTAGE. PERCENTAGE rates use the credit type of associated rates.
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate processors.
	CustomRate map[string]JSON `form:"custom_rate" json:"custom_rate,omitempty"`
	// exclusive end date
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	Entitled     *bool   `form:"entitled" json:"entitled"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be set to true.
	IsProrated    *bool                                  `form:"is_prorated" json:"is_prorated,omitempty"`
	MinimumConfig *V1RateCardsAddRateMinimumConfigParams `form:"minimum_config" json:"minimum_config,omitempty"`
	// Default price. For FLAT and SUBSCRIPTION rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price *Number `form:"price" json:"price,omitempty"`
	// Optional. List of pricing group key value pairs which will be used to calculate the price.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// ID of the product to add a rate for
	ProductID *string `form:"product_id" json:"product_id"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity *Number `form:"quantity" json:"quantity,omitempty"`
	// ID of the rate card to update
	RateCardID *string `form:"rate_card_id" json:"rate_card_id"`
	RateType   *string `form:"rate_type" json:"rate_type"`
	// inclusive effective date
	StartingAt *string `form:"starting_at" json:"starting_at"`
	// Only set for TIERED rate_type.
	Tiers []*V1RateCardsAddRateTierParams `form:"tiers" json:"tiers,omitempty"`
}

// Only set for TIERED rate_type.
type V1RateCardsAddRatesRateTierParams struct {
	Price *Number `form:"price" json:"price"`
	Size  *Number `form:"size" json:"size,omitempty"`
}
type V1RateCardsAddRatesRateMinimumConfigParams struct {
	Minimum *Number `form:"minimum" json:"minimum"`
}

// Only set for TIERED rate_type.
type V1RateCardsAddRatesRateCommitRateTierParams struct {
	Price *Number `form:"price" json:"price"`
	Size  *Number `form:"size" json:"size,omitempty"`
}
type V1RateCardsAddRatesRateCommitRateMinimumConfigParams struct {
	Minimum *Number `form:"minimum" json:"minimum"`
}
type V1RateCardsAddRatesRateCommitRateParams struct {
	MinimumConfig *V1RateCardsAddRatesRateCommitRateMinimumConfigParams `form:"minimum_config" json:"minimum_config,omitempty"`
	// Commit rate price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    *Number `form:"price" json:"price,omitempty"`
	RateType *string `form:"rate_type" json:"rate_type"`
	// Only set for TIERED rate_type.
	Tiers []*V1RateCardsAddRatesRateCommitRateTierParams `form:"tiers" json:"tiers,omitempty"`
}
type V1RateCardsAddRatesRateParams struct {
	// Optional. Frequency to bill subscriptions with. Required for subscription type products with Flat rate.
	BillingFrequency *string                                  `form:"billing_frequency" json:"billing_frequency,omitempty"`
	CommitRate       *V1RateCardsAddRatesRateCommitRateParams `form:"commit_rate" json:"commit_rate,omitempty"`
	// "The Metronome ID of the credit type to associate with price, defaults to USD (cents) if not passed. Used by all rate_types except type PERCENTAGE. PERCENTAGE rates use the credit type of associated rates."
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate processors.
	CustomRate map[string]JSON `form:"custom_rate" json:"custom_rate,omitempty"`
	// exclusive end date
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	Entitled     *bool   `form:"entitled" json:"entitled"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be set to true.
	IsProrated    *bool                                       `form:"is_prorated" json:"is_prorated,omitempty"`
	MinimumConfig *V1RateCardsAddRatesRateMinimumConfigParams `form:"minimum_config" json:"minimum_config,omitempty"`
	// Default price. For FLAT and SUBSCRIPTION rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price *Number `form:"price" json:"price,omitempty"`
	// Optional. List of pricing group key value pairs which will be used to calculate the price.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// ID of the product to add a rate for
	ProductID *string `form:"product_id" json:"product_id"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity *Number `form:"quantity" json:"quantity,omitempty"`
	RateType *string `form:"rate_type" json:"rate_type"`
	// inclusive effective date
	StartingAt *string `form:"starting_at" json:"starting_at"`
	// Only set for TIERED rate_type.
	Tiers []*V1RateCardsAddRatesRateTierParams `form:"tiers" json:"tiers,omitempty"`
}

// Add new rates
type V1RateCardsAddRatesParams struct {
	Params     `form:"*"`
	RateCardID *string                          `form:"rate_card_id" json:"rate_card_id"`
	Rates      []*V1RateCardsAddRatesRateParams `form:"rates" json:"rates"`
}

// Permanently disable a rate card by archiving it, preventing use in new contracts while preserving existing contract pricing. Use this when retiring old pricing models, consolidating rate cards, or removing outdated pricing structures. Returns the archived rate card ID and stops the rate card from appearing in contract creation workflows.
type V1RateCardsArchiveRateCardParams struct {
	Params `form:"*"`
	ID     *string `form:"id" json:"id"`
}

// Required when using custom pricing units in rates.
type V1RateCardsCreateRateCardCreditTypeConversionParams struct {
	CustomCreditTypeID  *string `form:"custom_credit_type_id" json:"custom_credit_type_id"`
	FiatPerCustomCredit *Number `form:"fiat_per_custom_credit" json:"fiat_per_custom_credit"`
}

// Reference this alias when creating a contract. If the same alias is assigned to multiple rate cards, it will reference the rate card to which it was most recently assigned. It is not exposed to end customers.
type V1RateCardsCreateRateCardAliasParams struct {
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	Name         *string `form:"name" json:"name"`
	StartingAt   *string `form:"starting_at" json:"starting_at,omitempty"`
}

// In Metronome, the rate card is the central location for your pricing. Rate cards were built with new product launches and pricing changes in mind - you can update your products and pricing in one place, and that change will be automatically propagated across your customer cohorts. Most clients need only maintain one or a few rate cards within Metronome.
//
// ### Use this endpoint to:
// - Create a rate card with a name and description
// - Define the rate card's single underlying fiat currency, and any number of conversion rates between that fiat currency and custom pricing units. You can then add products and associated rates in the fiat currency or custom pricing unit for which you have defined a conversion rate.
// - Set aliases for the rate card. Aliases are human-readable names that you can use in the place of the id of the rate card when provisioning a customer's contract. By using an alias, you can easily create a contract and provision a customer by choosing the paygo rate card, without storing the rate card id in your internal systems. This is helpful when launching a new rate card for paygo customers, you can update the alias for paygo to be scheduled to be assigned to the new rate card without updating your code.
//
// ### Key response fields:
// - The ID of the rate card you just created
//
// ### Usage guidelines:
// - After creating a rate card, you can now use the addRate or addRates endpoints to add products and their prices to it
// - A rate card alias can only be used by one rate card at a time. If you create a contract with a rate card alias that is already in use by another rate card, the original rate card's alias schedule will be updated. The alias will reference the rate card to which it was most recently assigned.
type V1RateCardsCreateRateCardParams struct {
	Params `form:"*"`
	// Reference this alias when creating a contract. If the same alias is assigned to multiple rate cards, it will reference the rate card to which it was most recently assigned. It is not exposed to end customers.
	Aliases []*V1RateCardsCreateRateCardAliasParams `form:"aliases" json:"aliases,omitempty"`
	// Required when using custom pricing units in rates.
	CreditTypeConversions []*V1RateCardsCreateRateCardCreditTypeConversionParams `form:"credit_type_conversions" json:"credit_type_conversions,omitempty"`
	CustomFields          map[string]string                                      `form:"custom_fields" json:"custom_fields,omitempty"`
	Description           *string                                                `form:"description" json:"description,omitempty"`
	// The Metronome ID of the credit type to associate with the rate card, defaults to USD (cents) if not passed.
	FiatCreditTypeID *string `form:"fiat_credit_type_id" json:"fiat_credit_type_id,omitempty"`
	// Used only in UI/API. It is not exposed to end customers.
	Name *string `form:"name" json:"name"`
}

// Return details for a specific rate card including name, description, and aliases. This endpoint does not return rates - use the dedicated getRates or getRateSchedule endpoints to understand the rates on a rate card.
type V1RateCardsGetRateCardParams struct {
	Params `form:"*"`
	ID     *string `form:"id" json:"id"`
}

// List of rate selectors, rates matching ANY of the selector will be included in the response Passing no selectors will result in all rates being returned.
type V1RateCardsGetRatesSelectorParams struct {
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

// Understand the rate schedule at a given timestamp, optionally filtering the list of rates returned based on properties such as `product_id` and `pricing_group_values`. For example, you may want to display the current price for a given product in your product experience - use this endpoint to fetch that information from its source of truth in Metronome.
//
// If you want to understand the rates for a specific customer's contract, inclusive of contract-level overrides, use the `getContractRateSchedule` endpoint.
type V1RateCardsGetRatesParams struct {
	Params `form:"*"`
	// inclusive starting point for the rates schedule
	At *string `form:"at" json:"at"`
	// Max number of results that should be returned
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
	// ID of the rate card to get the schedule for
	RateCardID *string `form:"rate_card_id" json:"rate_card_id"`
	// List of rate selectors, rates matching ANY of the selector will be included in the response Passing no selectors will result in all rates being returned.
	Selectors []*V1RateCardsGetRatesSelectorParams `form:"selectors" json:"selectors,omitempty"`
}

// List of rate selectors, rates matching ANY of the selector will be included in the response Passing no selectors will result in all rates being returned.
type V1RateCardsGetRateScheduleSelectorParams struct {
	// Subscription rates matching the billing frequency will be included in the response.
	BillingFrequency *string `form:"billing_frequency" json:"billing_frequency,omitempty"`
	// List of pricing group key value pairs, rates containing the matching key / value pairs will be included in the response.
	PartialPricingGroupValues map[string]string `form:"partial_pricing_group_values" json:"partial_pricing_group_values,omitempty"`
	// List of pricing group key value pairs, rates matching all of the key / value pairs will be included in the response.
	PricingGroupValues map[string]string `form:"pricing_group_values" json:"pricing_group_values,omitempty"`
	// Rates matching the product id will be included in the response.
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
}

// A rate card defines the prices that you charge for your products. Rate cards support scheduled changes over time, to allow you to easily roll out pricing changes and new product launches across your customer base. Use this endpoint to understand the rate schedule `starting_at` a given date, optionally filtering the list of rates returned based on product id or pricing group values. For example, you may want to display a schedule of upcoming price changes for a given product in your product experience - use this endpoint to fetch that information from its source of truth in Metronome.
//
// If you want to understand the rates for a specific customer's contract, inclusive of contract-level overrides, use the `getContractRateSchedule` endpoint.
type V1RateCardsGetRateScheduleParams struct {
	Params `form:"*"`
	// optional exclusive end date for the rates schedule. When not specified rates will show all future schedule segments.
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// Max number of results that should be returned
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
	// ID of the rate card to get the schedule for
	RateCardID *string `form:"rate_card_id" json:"rate_card_id"`
	// List of rate selectors, rates matching ANY of the selector will be included in the response Passing no selectors will result in all rates being returned.
	Selectors []*V1RateCardsGetRateScheduleSelectorParams `form:"selectors" json:"selectors,omitempty"`
	// inclusive starting point for the rates schedule
	StartingAt *string `form:"starting_at" json:"starting_at"`
}

// List all rate cards. Returns rate card IDs, names, descriptions, aliases, and other details. To view the rates associated with a given rate card, use the getRates or getRateSchedule endpoints.
type V1RateCardsListRateCardsParams struct {
	Params `form:"*"`
	// Max number of results that should be returned
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
}
type V1RateCardsMoveRateCardProductsProductMoveParams struct {
	// 0-based index of the new position of the product
	Position *Number `form:"position" json:"position"`
	// ID of the product to move
	ProductID *string `form:"product_id" json:"product_id"`
}

// The ordering of products on a rate card determines the order in which the products will appear on customers' invoices. Use this endpoint to set the order of specific products on the rate card by moving them relative to their current location.
type V1RateCardsMoveRateCardProductsParams struct {
	Params       `form:"*"`
	ProductMoves []*V1RateCardsMoveRateCardProductsProductMoveParams `form:"product_moves" json:"product_moves"`
	// ID of the rate card to update
	RateCardID *string `form:"rate_card_id" json:"rate_card_id"`
}

// The ordering of products on a rate card determines the order in which the products will appear on customers' invoices. Use this endpoint to set the order of products on the rate card.
type V1RateCardsSetRateCardProductsOrderParams struct {
	Params       `form:"*"`
	ProductOrder []*string `form:"product_order" json:"product_order"`
	// ID of the rate card to update
	RateCardID *string `form:"rate_card_id" json:"rate_card_id"`
}

// Reference this alias when creating a contract. If the same alias is assigned to multiple rate cards, it will reference the rate card to which it was most recently assigned. It is not exposed to end customers.
type V1RateCardsUpdateRateCardAliasParams struct {
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	Name         *string `form:"name" json:"name"`
	StartingAt   *string `form:"starting_at" json:"starting_at,omitempty"`
}

// Add credit type conversions for using custom pricing units in rates. Existing conversions cannot be modified.
type V1RateCardsUpdateRateCardAddCreditTypeConversionParams struct {
	CustomCreditTypeID  *string `form:"custom_credit_type_id" json:"custom_credit_type_id"`
	FiatPerCustomCredit *Number `form:"fiat_per_custom_credit" json:"fiat_per_custom_credit"`
}

// Update a rate card's name, description, aliases, and credit type conversion rates. This endpoint does not affect underlying pricing rates or schedules.
//
// ### Use this endpoint to:
// - Rename rate cards: Update display names or descriptions for organizational clarity
// - Manage aliases: Add, modify, or schedule alias transitions for seamless and code-free rate card migrations
// - Update documentation: Keep rate card descriptions current with business context
// - Configure custom pricing units: Add credit type conversions to enable rates with different pricing units
//
// #### Active contract impact:
// - Alias changes: Already-created contracts continue using their originally assigned rate cards.
// - Other changes made using this endpoint will only impact the Metronome UI.
//
// #### Grandfathering existing PLG customer pricing:
// - Rate card aliases support scheduled transitions, enabling seamless rate card migrations for new customers, allowing existing customers to be grandfathered into their existing prices without code. Note that there are multiple mechanisms to support grandfathering in Metronome.
//
// #### How scheduled aliases work for PLG grandfathering:
// Initial setup:
// - Add alias to current rate card: Assign a stable alias (e.g., "standard-pricing") to your active rate card
// - Reference alias during contract creation: Configure your self-serve workflow to create contracts using `rate_card_alias` instead of direct `rate_card_id`
// - Automatic resolution: New contracts referencing the alias automatically resolve to the rate card associated with the alias at the point in time of provisioning
//
// #### Grandfathering process:
// - Create new rate card: Build your new rate card with updated pricing structure
// - Schedule alias transition: Add the same alias to the new rate card with a `starting_at` timestamp
// - Automatic cutover: Starting at the scheduled time, new contracts created in your PLG workflow using that alias will automatically reference the new rate card
type V1RateCardsUpdateRateCardParams struct {
	Params `form:"*"`
	// Add credit type conversions for using custom pricing units in rates. Existing conversions cannot be modified.
	AddCreditTypeConversions []*V1RateCardsUpdateRateCardAddCreditTypeConversionParams `form:"add_credit_type_conversions" json:"add_credit_type_conversions,omitempty"`
	// Reference this alias when creating a contract. If the same alias is assigned to multiple rate cards, it will reference the rate card to which it was most recently assigned. It is not exposed to end customers.
	Aliases     []*V1RateCardsUpdateRateCardAliasParams `form:"aliases" json:"aliases,omitempty"`
	Description *string                                 `form:"description" json:"description,omitempty"`
	// Used only in UI/API. It is not exposed to end customers.
	Name *string `form:"name" json:"name,omitempty"`
	// ID of the rate card to update
	RateCardID *string `form:"rate_card_id" json:"rate_card_id"`
}
