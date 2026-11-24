//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v1CustomFieldsService is used to invoke v1customfields related APIs.
type v1CustomFieldsService struct {
	B   Backend
	Key string
}

// Creates a new custom field key for a given entity (e.g. billable metric, contract, alert).
//
// Custom fields are properties that you can add to Metronome objects to store metadata like foreign keys or other descriptors. This metadata can get transferred to or accessed by other systems to contextualize Metronome data and power business processes. For example, to service workflows like revenue recognition, reconciliation, and invoicing, custom fields help Metronome know the relationship between entities in the platform and third-party systems.
//
// ### Use this endpoint to:
// - Create a new custom field key for Customer objects in Metronome. You can then use the Set Custom Field Values endpoint to set the value of this key for a specific customer.
// - Specify whether the key should enforce uniqueness. If the key is set to enforce uniqueness and you attempt to set a custom field value for the key that already exists, it will fail.
//
// ### Usage guidelines:
// - Custom fields set on commits, credits, and contracts can be used to scope alert evaluation. For example, you can create a spend threshold alert that only considers spend associated with contracts with custom field key `contract_type` and value `paygo`
// - Custom fields set on products can be used in the Stripe integration to set metadata on invoices.
// - Custom fields for customers, contracts, invoices, products, commits, scheduled charges, and subscriptions are passed down to the invoice.
func (c v1CustomFieldsService) AddCustomFieldKey(ctx context.Context, params *V1CustomFieldsAddCustomFieldKeyParams) (*V1AddCustomFieldKeyResponse, error) {
	if params == nil {
		params = &V1CustomFieldsAddCustomFieldKeyParams{}
	}
	params.Context = ctx
	v1addcustomfieldkeyresponse := &V1AddCustomFieldKeyResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/customFields/addKey", c.Key, params, v1addcustomfieldkeyresponse)
	return v1addcustomfieldkeyresponse, err
}

// Remove specific custom field values from a Metronome entity instance by specifying the field keys to delete. Use this endpoint to clean up unwanted custom field data while preserving other fields on the same entity. Requires the entity type, entity ID, and array of keys to remove.
func (c v1CustomFieldsService) DeleteCustomFields(ctx context.Context, params *V1CustomFieldsDeleteCustomFieldsParams) (*V1DeleteCustomFieldsResponse, error) {
	if params == nil {
		params = &V1CustomFieldsDeleteCustomFieldsParams{}
	}
	params.Context = ctx
	v1deletecustomfieldsresponse := &V1DeleteCustomFieldsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/customFields/deleteValues", c.Key, params, v1deletecustomfieldsresponse)
	return v1deletecustomfieldsresponse, err
}

// Removes a custom field key from the allowlist for a specific entity type, preventing future use of that key across all instances of the entity. Existing values for this key on entity instances will no longer be accessible once the key is removed.
func (c v1CustomFieldsService) DisableCustomFieldKey(ctx context.Context, params *V1CustomFieldsDisableCustomFieldKeyParams) (*V1DisableCustomFieldKeyResponse, error) {
	if params == nil {
		params = &V1CustomFieldsDisableCustomFieldKeyParams{}
	}
	params.Context = ctx
	v1disablecustomfieldkeyresponse := &V1DisableCustomFieldKeyResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/customFields/removeKey", c.Key, params, v1disablecustomfieldkeyresponse)
	return v1disablecustomfieldkeyresponse, err
}

// Retrieve all your active custom field keys, with optional filtering by entity type (customer, contract, product, etc.). Use this endpoint to discover what custom field keys are available before setting values on entities or to audit your custom field configuration across different entity types.
func (c v1CustomFieldsService) ListCustomFieldKeys(ctx context.Context, params *V1CustomFieldsListCustomFieldKeysParams) (*V1ListCustomFieldKeysResponse, error) {
	if params == nil {
		params = &V1CustomFieldsListCustomFieldKeysParams{}
	}
	params.Context = ctx
	v1listcustomfieldkeysresponse := &V1ListCustomFieldKeysResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/customFields/listKeys", c.Key, params, v1listcustomfieldkeysresponse)
	return v1listcustomfieldkeysresponse, err
}

// Sets custom field values on a specific Metronome entity instance. Overwrites existing values for matching keys while preserving other fields. All updates are transactional—either all values are set or none are. Custom field values are limited to 200 characters each.
func (c v1CustomFieldsService) SetCustomFields(ctx context.Context, params *V1CustomFieldsSetCustomFieldsParams) (*V1SetCustomFieldsResponse, error) {
	if params == nil {
		params = &V1CustomFieldsSetCustomFieldsParams{}
	}
	params.Context = ctx
	v1setcustomfieldsresponse := &V1SetCustomFieldsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/customFields/setValues", c.Key, params, v1setcustomfieldsresponse)
	return v1setcustomfieldsresponse, err
}
