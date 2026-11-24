//
// File generated from our OpenAPI spec
//

package metronome

// Archive a contract-level or customer-level commit. Use this endpoint to deactivate a commit while preserving historical records. You will not be able to archive a commit until all of the finalized usage invoices the commit has been applied to are voided, and all of the finalized invoices for commit payment have been voided.
//
// Example workflow:
// The customer was provisioned a prepaid commit erroneously. It was applied to their most recent finalized usage invoice.
// - First, void the finalized invoice that the commit was applied to. Also, void the finalized invoice associated with the commit payment.
// - Then, use the archiveCommit endpoint to deactivate the commit.
// - Finally, regenerate the voided invoice. The invoice will be regenerated without the application of the commit, which has now been archived.
//
// ### Usage guidelines:
// - Once a commit has been archived, it will no longer appear by default on the endpoints `listCustomerCommits` or `listCustomerBalances`. Use the `include_archived` parameter to choose to fetch the details.
// - Once a commit has been archived, it has a null ledger and 0 remaining balance.
// - Archiving a commit fully deactivates the entire access schedule. If you want to reduce the amount granted in a commit, consider editing the access schedule using the `editCommit` endpoint, or adding a manual ledger entry using the `addManualBalanceLedgerEntry` endpoint.
type V2CreditsAndCommitsArchiveCommitParams struct {
	Params `form:"*"`
	// ID of the commit to archive
	CommitID *string `form:"commit_id" json:"commit_id"`
	// ID of the customer whose commit is being archived
	CustomerID *string `form:"customer_id" json:"customer_id"`
}

// Archive a contract-level or customer-level credit. Use this endpoint to deactivate a credit while preserving historical records. You will not be able to archive a credit until all of the finalized invoices the credit has been applied to are voided.
//
// Example workflow:
// The customer was granted a free credit erroneously. It was applied to their most recent finalized invoice.
// - First, void the finalized invoice that the credit was applied to.
// - Then, use the archiveCredit endpoint to deactivate the credit.
// - Finally, regenerate the voided invoice. The invoice will be regenerated without the application of the credit, which has now been archived.
//
// ### Usage guidelines:
// - Once a credit has been archived, it will no longer appear by default on the endpoints `listCustomerCredits` or `listCustomerBalances`. Use the `include_archived` parameter to choose to fetch the details.
// - Once a credit has been archived, it has a null ledger and 0 remaining balance.
// - Archiving a credit fully deactivates the entire access schedule. If you want to reduce the amount granted in a credit, consider editing the access schedule using the `editCredit` endpoint, or adding a manual ledger entry using the `addManualBalanceLedgerEntry` endpoint.
type V2CreditsAndCommitsArchiveCreditParams struct {
	Params `form:"*"`
	// ID of the credit to archive
	CreditID *string `form:"credit_id" json:"credit_id"`
	// ID of the customer whose credit is being archived
	CustomerID *string `form:"customer_id" json:"customer_id"`
}
type V2CreditsAndCommitsEditCommitAccessScheduleAddScheduleItemParams struct {
	Amount       *Number `form:"amount" json:"amount"`
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	StartingAt   *string `form:"starting_at" json:"starting_at"`
}
type V2CreditsAndCommitsEditCommitAccessScheduleUpdateScheduleItemParams struct {
	Amount       *Number `form:"amount" json:"amount,omitempty"`
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	ID           *string `form:"id" json:"id"`
	StartingAt   *string `form:"starting_at" json:"starting_at,omitempty"`
}
type V2CreditsAndCommitsEditCommitAccessScheduleRemoveScheduleItemParams struct {
	ID *string `form:"id" json:"id"`
}
type V2CreditsAndCommitsEditCommitAccessScheduleParams struct {
	AddScheduleItems    []*V2CreditsAndCommitsEditCommitAccessScheduleAddScheduleItemParams    `form:"add_schedule_items" json:"add_schedule_items,omitempty"`
	RemoveScheduleItems []*V2CreditsAndCommitsEditCommitAccessScheduleRemoveScheduleItemParams `form:"remove_schedule_items" json:"remove_schedule_items,omitempty"`
	UpdateScheduleItems []*V2CreditsAndCommitsEditCommitAccessScheduleUpdateScheduleItemParams `form:"update_schedule_items" json:"update_schedule_items,omitempty"`
}
type V2CreditsAndCommitsEditCommitInvoiceScheduleAddScheduleItemParams struct {
	Amount    *Number `form:"amount" json:"amount,omitempty"`
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	Timestamp *string `form:"timestamp" json:"timestamp"`
	UnitPrice *Number `form:"unit_price" json:"unit_price,omitempty"`
}
type V2CreditsAndCommitsEditCommitInvoiceScheduleUpdateScheduleItemParams struct {
	Amount    *Number `form:"amount" json:"amount,omitempty"`
	ID        *string `form:"id" json:"id"`
	Quantity  *Number `form:"quantity" json:"quantity,omitempty"`
	Timestamp *string `form:"timestamp" json:"timestamp,omitempty"`
	UnitPrice *Number `form:"unit_price" json:"unit_price,omitempty"`
}
type V2CreditsAndCommitsEditCommitInvoiceScheduleRemoveScheduleItemParams struct {
	ID *string `form:"id" json:"id"`
}
type V2CreditsAndCommitsEditCommitInvoiceScheduleParams struct {
	AddScheduleItems    []*V2CreditsAndCommitsEditCommitInvoiceScheduleAddScheduleItemParams    `form:"add_schedule_items" json:"add_schedule_items,omitempty"`
	RemoveScheduleItems []*V2CreditsAndCommitsEditCommitInvoiceScheduleRemoveScheduleItemParams `form:"remove_schedule_items" json:"remove_schedule_items,omitempty"`
	UpdateScheduleItems []*V2CreditsAndCommitsEditCommitInvoiceScheduleUpdateScheduleItemParams `form:"update_schedule_items" json:"update_schedule_items,omitempty"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V2CreditsAndCommitsEditCommitSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
type V2CreditsAndCommitsEditCommitSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V2CreditsAndCommitsEditCommitSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
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
type V2CreditsAndCommitsEditCommitHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}

// Edit specific details for a contract-level or customer-level commit. Use this endpoint to modify individual commit access schedules, invoice schedules, applicable products, invoicing contracts, or other fields.
//
// ### Usage guidelines:
// - As with all edits in Metronome, draft invoices will reflect the edit immediately, while finalized invoices are untouched unless voided and regenerated.
// - If a commit's invoice schedule item is associated with a finalized invoice, you cannot remove or update the invoice schedule item.
// - If a commit's invoice schedule item is associated with a voided invoice, you cannot remove the invoice schedule item.
// - You cannot remove an commit access schedule segment that was applied to a finalized invoice. You can void the invoice beforehand and then remove the access schedule segment.
type V2CreditsAndCommitsEditCommitParams struct {
	Params         `form:"*"`
	AccessSchedule *V2CreditsAndCommitsEditCommitAccessScheduleParams `form:"access_schedule" json:"access_schedule,omitempty"`
	// Which contracts the customer-level commit applies to. If set to null, the commit applies to all of the customer's contracts. This field cannot be edited for POSTPAID commits or contract-level commits.
	ApplicableContractIDs []*string `form:"applicable_contract_ids" json:"applicable_contract_ids,omitempty"`
	// Which products the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	// ID of the commit to edit
	CommitID *string `form:"commit_id" json:"commit_id"`
	// ID of the customer whose commit is being edited
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// Updated description for the commit
	Description *string `form:"description" json:"description,omitempty"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration *V2CreditsAndCommitsEditCommitHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// ID of contract to use for invoicing
	InvoiceContractID *string                                             `form:"invoice_contract_id" json:"invoice_contract_id,omitempty"`
	InvoiceSchedule   *V2CreditsAndCommitsEditCommitInvoiceScheduleParams `form:"invoice_schedule" json:"invoice_schedule,omitempty"`
	// Updated name for the commit
	Name *string `form:"name" json:"name,omitempty"`
	// If multiple commits are applicable, the one with the lower priority will apply first.
	Priority  *Number `form:"priority" json:"priority,omitempty"`
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, updates the commit to use the specified rate type for current and future invoices. Previously finalized invoices will need to be voided and regenerated to reflect the rate type change.
	RateType *string `form:"rate_type" json:"rate_type,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
	Specifiers []*V2CreditsAndCommitsEditCommitSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
}
type V2CreditsAndCommitsEditCreditAccessScheduleAddScheduleItemParams struct {
	Amount       *Number `form:"amount" json:"amount"`
	EndingBefore *string `form:"ending_before" json:"ending_before"`
	StartingAt   *string `form:"starting_at" json:"starting_at"`
}
type V2CreditsAndCommitsEditCreditAccessScheduleUpdateScheduleItemParams struct {
	Amount       *Number `form:"amount" json:"amount,omitempty"`
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	ID           *string `form:"id" json:"id"`
	StartingAt   *string `form:"starting_at" json:"starting_at,omitempty"`
}
type V2CreditsAndCommitsEditCreditAccessScheduleRemoveScheduleItemParams struct {
	ID *string `form:"id" json:"id"`
}
type V2CreditsAndCommitsEditCreditAccessScheduleParams struct {
	AddScheduleItems    []*V2CreditsAndCommitsEditCreditAccessScheduleAddScheduleItemParams    `form:"add_schedule_items" json:"add_schedule_items,omitempty"`
	RemoveScheduleItems []*V2CreditsAndCommitsEditCreditAccessScheduleRemoveScheduleItemParams `form:"remove_schedule_items" json:"remove_schedule_items,omitempty"`
	UpdateScheduleItems []*V2CreditsAndCommitsEditCreditAccessScheduleUpdateScheduleItemParams `form:"update_schedule_items" json:"update_schedule_items,omitempty"`
}

// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
type V2CreditsAndCommitsEditCreditSpecifierExcludeParams struct {
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []*string `form:"product_tags" json:"product_tags,omitempty"`
}

// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
type V2CreditsAndCommitsEditCreditSpecifierParams struct {
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*V2CreditsAndCommitsEditCreditSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
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
type V2CreditsAndCommitsEditCreditHierarchyConfigurationParams struct {
	ChildAccess *JSON `form:"child_access" json:"child_access"`
}

// Edit details for a contract-level or customer-level credit.
//
// ### Use this endpoint to:
// - Extend the duration or the amount of an existing free credit like a trial
// - Modify individual credit access schedules, applicable products, priority, or other fields.
//
// ### Usage guidelines:
// - As with all edits in Metronome, draft invoices will reflect the edit immediately, while finalized invoices are untouched unless voided and regenerated.
// - You cannot remove an access schedule segment that was applied to a finalized invoice. You can void the invoice beforehand and then remove the access schedule segment.
type V2CreditsAndCommitsEditCreditParams struct {
	Params         `form:"*"`
	AccessSchedule *V2CreditsAndCommitsEditCreditAccessScheduleParams `form:"access_schedule" json:"access_schedule,omitempty"`
	// Which contracts the customer-level credit applies to. If set to null, the credit applies to all of the customer's contracts. This field cannot be set on a contract-level credit.
	ApplicableContractIDs []*string `form:"applicable_contract_ids" json:"applicable_contract_ids,omitempty"`
	// Which products the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductIDs []*string `form:"applicable_product_ids" json:"applicable_product_ids,omitempty"`
	// Which tags the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductTags []*string `form:"applicable_product_tags" json:"applicable_product_tags,omitempty"`
	// ID of the credit to edit
	CreditID *string `form:"credit_id" json:"credit_id"`
	// ID of the customer whose credit is being edited
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// Updated description for the credit
	Description *string `form:"description" json:"description,omitempty"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration *V2CreditsAndCommitsEditCreditHierarchyConfigurationParams `form:"hierarchy_configuration" json:"hierarchy_configuration,omitempty"`
	// Updated name for the credit
	Name *string `form:"name" json:"name,omitempty"`
	// If multiple commits are applicable, the one with the lower priority will apply first.
	Priority  *Number `form:"priority" json:"priority,omitempty"`
	ProductID *string `form:"product_id" json:"product_id,omitempty"`
	// If provided, updates the credit to use the specified rate type for current and future invoices. Previously finalized invoices will need to be voided and regenerated to reflect the rate type change.
	RateType *string `form:"rate_type" json:"rate_type,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
	Specifiers []*V2CreditsAndCommitsEditCreditSpecifierParams `form:"specifiers" json:"specifiers,omitempty"`
}

// Move one or more fully depleted, immutable commits on a contract to historical storage. Historical commits are removed from active code paths, which makes future computations for this customer faster. They remain retrievable through the historical view. Set `dry_run` to `true` to preview the result without making changes.
type V2CreditsAndCommitsMoveToHistoryParams struct {
	Params `form:"*"`
	// IDs of the commits to move to historical
	CommitIDs []*string `form:"commit_ids" json:"commit_ids"`
	// ID of the contract containing the commits
	ContractID *string `form:"contract_id" json:"contract_id"`
	// ID of the customer whose commits are being moved to historical
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// If true, previews which commits would be moved to historical and which would be skipped, without making changes
	DryRun *bool `form:"dry_run" json:"dry_run,omitempty"`
}
