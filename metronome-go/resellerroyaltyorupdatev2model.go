//
// File generated from our OpenAPI spec
//

package metronome

type ResellerRoyaltyOrUpdateV2ModelResellerType string

// List of values that ResellerRoyaltyOrUpdateV2ModelResellerType can take
const (
	ResellerRoyaltyOrUpdateV2ModelResellerTypeAWS           ResellerRoyaltyOrUpdateV2ModelResellerType = "AWS"
	ResellerRoyaltyOrUpdateV2ModelResellerTypeAWSPROSERVICE ResellerRoyaltyOrUpdateV2ModelResellerType = "AWS_PRO_SERVICE"
	ResellerRoyaltyOrUpdateV2ModelResellerTypeGCP           ResellerRoyaltyOrUpdateV2ModelResellerType = "GCP"
	ResellerRoyaltyOrUpdateV2ModelResellerTypeGCPPROSERVICE ResellerRoyaltyOrUpdateV2ModelResellerType = "GCP_PRO_SERVICE"
)

type ResellerRoyaltyOrUpdateV2Model struct {
	APIResource
	ApplicableProductIDs  []string                                   `json:"applicable_product_ids,omitempty"`
	ApplicableProductTags []string                                   `json:"applicable_product_tags,omitempty"`
	AwsAccountNumber      string                                     `json:"aws_account_number,omitempty"`
	AwsOfferID            string                                     `json:"aws_offer_id,omitempty"`
	AwsPayerReferenceID   string                                     `json:"aws_payer_reference_id,omitempty"`
	EndingBefore          string                                     `json:"ending_before,omitempty"`
	Fraction              Number                                     `json:"fraction,omitempty"`
	GcpAccountID          string                                     `json:"gcp_account_id,omitempty"`
	GcpOfferID            string                                     `json:"gcp_offer_id,omitempty"`
	NetsuiteResellerID    string                                     `json:"netsuite_reseller_id,omitempty"`
	ResellerContractValue Number                                     `json:"reseller_contract_value,omitempty"`
	ResellerType          ResellerRoyaltyOrUpdateV2ModelResellerType `json:"reseller_type"`
	StartingAt            string                                     `json:"starting_at,omitempty"`
}
