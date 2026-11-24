//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v2CreditsAndCommitsService is used to invoke v2creditsandcommits related APIs.
type v2CreditsAndCommitsService struct {
	B   Backend
	Key string
}

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
func (c v2CreditsAndCommitsService) ArchiveCommit(ctx context.Context, params *V2CreditsAndCommitsArchiveCommitParams) (*V2ArchiveCommitResponse, error) {
	if params == nil {
		params = &V2CreditsAndCommitsArchiveCommitParams{}
	}
	params.Context = ctx
	v2archivecommitresponse := &V2ArchiveCommitResponse{}
	err := c.B.Call(
		http.MethodPost, "/v2/contracts/commits/archive", c.Key, params, v2archivecommitresponse)
	return v2archivecommitresponse, err
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
func (c v2CreditsAndCommitsService) ArchiveCredit(ctx context.Context, params *V2CreditsAndCommitsArchiveCreditParams) (*V2ArchiveCreditResponse, error) {
	if params == nil {
		params = &V2CreditsAndCommitsArchiveCreditParams{}
	}
	params.Context = ctx
	v2archivecreditresponse := &V2ArchiveCreditResponse{}
	err := c.B.Call(
		http.MethodPost, "/v2/contracts/credits/archive", c.Key, params, v2archivecreditresponse)
	return v2archivecreditresponse, err
}

// Edit specific details for a contract-level or customer-level commit. Use this endpoint to modify individual commit access schedules, invoice schedules, applicable products, invoicing contracts, or other fields.
//
// ### Usage guidelines:
// - As with all edits in Metronome, draft invoices will reflect the edit immediately, while finalized invoices are untouched unless voided and regenerated.
// - If a commit's invoice schedule item is associated with a finalized invoice, you cannot remove or update the invoice schedule item.
// - If a commit's invoice schedule item is associated with a voided invoice, you cannot remove the invoice schedule item.
// - You cannot remove an commit access schedule segment that was applied to a finalized invoice. You can void the invoice beforehand and then remove the access schedule segment.
func (c v2CreditsAndCommitsService) EditCommit(ctx context.Context, params *V2CreditsAndCommitsEditCommitParams) (*V2EditCommitResponse, error) {
	if params == nil {
		params = &V2CreditsAndCommitsEditCommitParams{}
	}
	params.Context = ctx
	v2editcommitresponse := &V2EditCommitResponse{}
	err := c.B.Call(
		http.MethodPost, "/v2/contracts/commits/edit", c.Key, params, v2editcommitresponse)
	return v2editcommitresponse, err
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
func (c v2CreditsAndCommitsService) EditCredit(ctx context.Context, params *V2CreditsAndCommitsEditCreditParams) (*V2EditCreditResponse, error) {
	if params == nil {
		params = &V2CreditsAndCommitsEditCreditParams{}
	}
	params.Context = ctx
	v2editcreditresponse := &V2EditCreditResponse{}
	err := c.B.Call(
		http.MethodPost, "/v2/contracts/credits/edit", c.Key, params, v2editcreditresponse)
	return v2editcreditresponse, err
}

// Move one or more fully depleted, immutable commits on a contract to historical storage. Historical commits are removed from active code paths, which makes future computations for this customer faster. They remain retrievable through the historical view. Set `dry_run` to `true` to preview the result without making changes.
func (c v2CreditsAndCommitsService) MoveToHistory(ctx context.Context, params *V2CreditsAndCommitsMoveToHistoryParams) (*V2MoveToHistoryResponse, error) {
	if params == nil {
		params = &V2CreditsAndCommitsMoveToHistoryParams{}
	}
	params.Context = ctx
	v2movetohistoryresponse := &V2MoveToHistoryResponse{}
	err := c.B.Call(
		http.MethodPost, "/v2/contracts/commits/moveToHistory", c.Key, params, v2movetohistoryresponse)
	return v2movetohistoryresponse, err
}
