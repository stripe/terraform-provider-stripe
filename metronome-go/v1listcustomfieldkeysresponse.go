//
// File generated from our OpenAPI spec
//

package metronome

type V1ListCustomFieldKeysResponseDataEntity string

// List of values that V1ListCustomFieldKeysResponseDataEntity can take
const (
	V1ListCustomFieldKeysResponseDataEntityAlert                  V1ListCustomFieldKeysResponseDataEntity = "alert"
	V1ListCustomFieldKeysResponseDataEntityBillableMetric         V1ListCustomFieldKeysResponseDataEntity = "billable_metric"
	V1ListCustomFieldKeysResponseDataEntityCharge                 V1ListCustomFieldKeysResponseDataEntity = "charge"
	V1ListCustomFieldKeysResponseDataEntityCommit                 V1ListCustomFieldKeysResponseDataEntity = "commit"
	V1ListCustomFieldKeysResponseDataEntityContractCredit         V1ListCustomFieldKeysResponseDataEntity = "contract_credit"
	V1ListCustomFieldKeysResponseDataEntityContractProduct        V1ListCustomFieldKeysResponseDataEntity = "contract_product"
	V1ListCustomFieldKeysResponseDataEntityContract               V1ListCustomFieldKeysResponseDataEntity = "contract"
	V1ListCustomFieldKeysResponseDataEntityCustomer               V1ListCustomFieldKeysResponseDataEntity = "customer"
	V1ListCustomFieldKeysResponseDataEntityDiscount               V1ListCustomFieldKeysResponseDataEntity = "discount"
	V1ListCustomFieldKeysResponseDataEntityInvoice                V1ListCustomFieldKeysResponseDataEntity = "invoice"
	V1ListCustomFieldKeysResponseDataEntityProfessionalService    V1ListCustomFieldKeysResponseDataEntity = "professional_service"
	V1ListCustomFieldKeysResponseDataEntityProduct                V1ListCustomFieldKeysResponseDataEntity = "product"
	V1ListCustomFieldKeysResponseDataEntityRateCard               V1ListCustomFieldKeysResponseDataEntity = "rate_card"
	V1ListCustomFieldKeysResponseDataEntityScheduledCharge        V1ListCustomFieldKeysResponseDataEntity = "scheduled_charge"
	V1ListCustomFieldKeysResponseDataEntitySubscription           V1ListCustomFieldKeysResponseDataEntity = "subscription"
	V1ListCustomFieldKeysResponseDataEntityPackageCommit          V1ListCustomFieldKeysResponseDataEntity = "package_commit"
	V1ListCustomFieldKeysResponseDataEntityPackageCredit          V1ListCustomFieldKeysResponseDataEntity = "package_credit"
	V1ListCustomFieldKeysResponseDataEntityPackageSubscription    V1ListCustomFieldKeysResponseDataEntity = "package_subscription"
	V1ListCustomFieldKeysResponseDataEntityPackageScheduledCharge V1ListCustomFieldKeysResponseDataEntity = "package_scheduled_charge"
)

type V1ListCustomFieldKeysResponseData struct {
	EnforceUniqueness bool                                    `json:"enforce_uniqueness"`
	Entity            V1ListCustomFieldKeysResponseDataEntity `json:"entity"`
	Key               string                                  `json:"key"`
}
type V1ListCustomFieldKeysResponse struct {
	APIResource
	Data     []*V1ListCustomFieldKeysResponseData `json:"data"`
	NextPage string                               `json:"next_page"`
}
