//
// File generated from our OpenAPI spec
//

package metronome

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
type V1CustomFieldsAddCustomFieldKeyParams struct {
	Params            `form:"*"`
	EnforceUniqueness *bool   `form:"enforce_uniqueness" json:"enforce_uniqueness"`
	Entity            *string `form:"entity" json:"entity"`
	Key               *string `form:"key" json:"key"`
}

// Remove specific custom field values from a Metronome entity instance by specifying the field keys to delete. Use this endpoint to clean up unwanted custom field data while preserving other fields on the same entity. Requires the entity type, entity ID, and array of keys to remove.
type V1CustomFieldsDeleteCustomFieldsParams struct {
	Params   `form:"*"`
	Entity   *string   `form:"entity" json:"entity"`
	EntityID *string   `form:"entity_id" json:"entity_id"`
	Keys     []*string `form:"keys" json:"keys"`
}

// Removes a custom field key from the allowlist for a specific entity type, preventing future use of that key across all instances of the entity. Existing values for this key on entity instances will no longer be accessible once the key is removed.
type V1CustomFieldsDisableCustomFieldKeyParams struct {
	Params `form:"*"`
	Entity *string `form:"entity" json:"entity"`
	Key    *string `form:"key" json:"key"`
}

// Retrieve all your active custom field keys, with optional filtering by entity type (customer, contract, product, etc.). Use this endpoint to discover what custom field keys are available before setting values on entities or to audit your custom field configuration across different entity types.
type V1CustomFieldsListCustomFieldKeysParams struct {
	Params `form:"*"`
	// Optional list of entity types to return keys for
	Entities []*string `form:"entities" json:"entities,omitempty"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
}

// Sets custom field values on a specific Metronome entity instance. Overwrites existing values for matching keys while preserving other fields. All updates are transactional—either all values are set or none are. Custom field values are limited to 200 characters each.
type V1CustomFieldsSetCustomFieldsParams struct {
	Params       `form:"*"`
	CustomFields map[string]string `form:"custom_fields" json:"custom_fields"`
	Entity       *string           `form:"entity" json:"entity"`
	EntityID     *string           `form:"entity_id" json:"entity_id"`
}
