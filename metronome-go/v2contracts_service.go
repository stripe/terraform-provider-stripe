//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v2ContractsService is used to invoke v2contracts related APIs.
type v2ContractsService struct {
	B   Backend
	Key string
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
func (c v2ContractsService) EditContract(ctx context.Context, params *V2ContractsEditContractParams) (*V2EditContractResponse, error) {
	if params == nil {
		params = &V2ContractsEditContractParams{}
	}
	params.Context = ctx
	v2editcontractresponse := &V2EditContractResponse{}
	err := c.B.Call(
		http.MethodPost, "/v2/contracts/edit", c.Key, params, v2editcontractresponse)
	return v2editcontractresponse, err
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
func (c v2ContractsService) GetContract(ctx context.Context, params *V2ContractsGetContractParams) (*V2GetContractResponse, error) {
	if params == nil {
		params = &V2ContractsGetContractParams{}
	}
	params.Context = ctx
	v2getcontractresponse := &V2GetContractResponse{}
	err := c.B.Call(
		http.MethodPost, "/v2/contracts/get", c.Key, params, v2getcontractresponse)
	return v2getcontractresponse, err
}

// List all the edits made to a contract over time. In Metronome, you can edit a contract at any point after it's created to fix mistakes or reflect changes in terms. Metronome stores a full history of all edits that were ever made to a contract, whether through the UI, `editContract` endpoint, or other endpoints like `updateContractEndDate`.
//
// ### Use this endpoint to:
// - Understand what changes were made to a contract, when, and by who
//
// ### Key response fields:
// - An array of every edit ever made to the contract
// - Details on each individual edit - for example showing that in one edit, a user added two discounts and incremented a subscription quantity.
func (c v2ContractsService) GetContractEditHistory(ctx context.Context, params *V2ContractsGetContractEditHistoryParams) (*V2GetContractEditHistoryResponse, error) {
	if params == nil {
		params = &V2ContractsGetContractEditHistoryParams{}
	}
	params.Context = ctx
	v2getcontractedithistoryresponse := &V2GetContractEditHistoryResponse{}
	err := c.B.Call(
		http.MethodPost, "/v2/contracts/getEditHistory", c.Key, params, v2getcontractedithistoryresponse)
	return v2getcontractedithistoryresponse, err
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
func (c v2ContractsService) ListContracts(ctx context.Context, params *V2ContractsListContractsParams) (*V2ListContractsResponse, error) {
	if params == nil {
		params = &V2ContractsListContractsParams{}
	}
	params.Context = ctx
	v2listcontractsresponse := &V2ListContractsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v2/contracts/list", c.Key, params, v2listcontractsresponse)
	return v2listcontractsresponse, err
}
