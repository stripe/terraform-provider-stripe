//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v1RateCardsService is used to invoke v1ratecards related APIs.
type v1RateCardsService struct {
	B   Backend
	Key string
}

// Add a new rate
//
// This endpoint is heavily rate limited. For adding multiple rates, using the [addRates](https://docs.metronome.com/api-reference/rate-cards/add-rates) endpoint is strongly encouraged.
func (c v1RateCardsService) AddRate(ctx context.Context, params *V1RateCardsAddRateParams) (*V1AddRateResponse, error) {
	if params == nil {
		params = &V1RateCardsAddRateParams{}
	}
	params.Context = ctx
	v1addrateresponse := &V1AddRateResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contract-pricing/rate-cards/addRate", c.Key, params, v1addrateresponse)
	return v1addrateresponse, err
}

// Add new rates
func (c v1RateCardsService) AddRates(ctx context.Context, params *V1RateCardsAddRatesParams) (*V1AddRatesResponse, error) {
	if params == nil {
		params = &V1RateCardsAddRatesParams{}
	}
	params.Context = ctx
	v1addratesresponse := &V1AddRatesResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contract-pricing/rate-cards/addRates", c.Key, params, v1addratesresponse)
	return v1addratesresponse, err
}

// Permanently disable a rate card by archiving it, preventing use in new contracts while preserving existing contract pricing. Use this when retiring old pricing models, consolidating rate cards, or removing outdated pricing structures. Returns the archived rate card ID and stops the rate card from appearing in contract creation workflows.
func (c v1RateCardsService) ArchiveRateCard(ctx context.Context, params *V1RateCardsArchiveRateCardParams) (*V1ArchiveRateCardResponse, error) {
	if params == nil {
		params = &V1RateCardsArchiveRateCardParams{}
	}
	params.Context = ctx
	v1archiveratecardresponse := &V1ArchiveRateCardResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contract-pricing/rate-cards/archive", c.Key, params, v1archiveratecardresponse)
	return v1archiveratecardresponse, err
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
func (c v1RateCardsService) CreateRateCard(ctx context.Context, params *V1RateCardsCreateRateCardParams) (*V1CreateRateCardResponse, error) {
	if params == nil {
		params = &V1RateCardsCreateRateCardParams{}
	}
	params.Context = ctx
	v1createratecardresponse := &V1CreateRateCardResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contract-pricing/rate-cards/create", c.Key, params, v1createratecardresponse)
	return v1createratecardresponse, err
}

// Return details for a specific rate card including name, description, and aliases. This endpoint does not return rates - use the dedicated getRates or getRateSchedule endpoints to understand the rates on a rate card.
func (c v1RateCardsService) GetRateCard(ctx context.Context, params *V1RateCardsGetRateCardParams) (*V1GetRateCardResponse, error) {
	if params == nil {
		params = &V1RateCardsGetRateCardParams{}
	}
	params.Context = ctx
	v1getratecardresponse := &V1GetRateCardResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contract-pricing/rate-cards/get", c.Key, params, v1getratecardresponse)
	return v1getratecardresponse, err
}

// Understand the rate schedule at a given timestamp, optionally filtering the list of rates returned based on properties such as `product_id` and `pricing_group_values`. For example, you may want to display the current price for a given product in your product experience - use this endpoint to fetch that information from its source of truth in Metronome.
//
// If you want to understand the rates for a specific customer's contract, inclusive of contract-level overrides, use the `getContractRateSchedule` endpoint.
func (c v1RateCardsService) GetRates(ctx context.Context, params *V1RateCardsGetRatesParams) (*V1GetRatesResponse, error) {
	if params == nil {
		params = &V1RateCardsGetRatesParams{}
	}
	params.Context = ctx
	v1getratesresponse := &V1GetRatesResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contract-pricing/rate-cards/getRates", c.Key, params, v1getratesresponse)
	return v1getratesresponse, err
}

// A rate card defines the prices that you charge for your products. Rate cards support scheduled changes over time, to allow you to easily roll out pricing changes and new product launches across your customer base. Use this endpoint to understand the rate schedule `starting_at` a given date, optionally filtering the list of rates returned based on product id or pricing group values. For example, you may want to display a schedule of upcoming price changes for a given product in your product experience - use this endpoint to fetch that information from its source of truth in Metronome.
//
// If you want to understand the rates for a specific customer's contract, inclusive of contract-level overrides, use the `getContractRateSchedule` endpoint.
func (c v1RateCardsService) GetRateSchedule(ctx context.Context, params *V1RateCardsGetRateScheduleParams) (*V1GetRateScheduleResponse, error) {
	if params == nil {
		params = &V1RateCardsGetRateScheduleParams{}
	}
	params.Context = ctx
	v1getratescheduleresponse := &V1GetRateScheduleResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contract-pricing/rate-cards/getRateSchedule", c.Key, params, v1getratescheduleresponse)
	return v1getratescheduleresponse, err
}

// List all rate cards. Returns rate card IDs, names, descriptions, aliases, and other details. To view the rates associated with a given rate card, use the getRates or getRateSchedule endpoints.
func (c v1RateCardsService) ListRateCards(ctx context.Context, params *V1RateCardsListRateCardsParams) (*V1ListRateCardsResponse, error) {
	if params == nil {
		params = &V1RateCardsListRateCardsParams{}
	}
	params.Context = ctx
	v1listratecardsresponse := &V1ListRateCardsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contract-pricing/rate-cards/list", c.Key, params, v1listratecardsresponse)
	return v1listratecardsresponse, err
}

// The ordering of products on a rate card determines the order in which the products will appear on customers' invoices. Use this endpoint to set the order of specific products on the rate card by moving them relative to their current location.
func (c v1RateCardsService) MoveRateCardProducts(ctx context.Context, params *V1RateCardsMoveRateCardProductsParams) (*V1MoveRateCardProductsResponse, error) {
	if params == nil {
		params = &V1RateCardsMoveRateCardProductsParams{}
	}
	params.Context = ctx
	v1moveratecardproductsresponse := &V1MoveRateCardProductsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contract-pricing/rate-cards/moveRateCardProducts", c.Key, params, v1moveratecardproductsresponse)
	return v1moveratecardproductsresponse, err
}

// The ordering of products on a rate card determines the order in which the products will appear on customers' invoices. Use this endpoint to set the order of products on the rate card.
func (c v1RateCardsService) SetRateCardProductsOrder(ctx context.Context, params *V1RateCardsSetRateCardProductsOrderParams) (*V1SetRateCardProductsOrderResponse, error) {
	if params == nil {
		params = &V1RateCardsSetRateCardProductsOrderParams{}
	}
	params.Context = ctx
	v1setratecardproductsorderresponse := &V1SetRateCardProductsOrderResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contract-pricing/rate-cards/setRateCardProductsOrder", c.Key, params, v1setratecardproductsorderresponse)
	return v1setratecardproductsorderresponse, err
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
func (c v1RateCardsService) UpdateRateCard(ctx context.Context, params *V1RateCardsUpdateRateCardParams) (*V1UpdateRateCardResponse, error) {
	if params == nil {
		params = &V1RateCardsUpdateRateCardParams{}
	}
	params.Context = ctx
	v1updateratecardresponse := &V1UpdateRateCardResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contract-pricing/rate-cards/update", c.Key, params, v1updateratecardresponse)
	return v1updateratecardresponse, err
}
