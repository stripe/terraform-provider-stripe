//
// File generated from our OpenAPI spec
//

package metronome

type SubscriptionProrationModelInvoiceBehavior string

// List of values that SubscriptionProrationModelInvoiceBehavior can take
const (
	SubscriptionProrationModelInvoiceBehaviorBILLIMMEDIATELY          SubscriptionProrationModelInvoiceBehavior = "BILL_IMMEDIATELY"
	SubscriptionProrationModelInvoiceBehaviorBILLONNEXTCOLLECTIONDATE SubscriptionProrationModelInvoiceBehavior = "BILL_ON_NEXT_COLLECTION_DATE"
	SubscriptionProrationModelInvoiceBehaviorBillImmediately          SubscriptionProrationModelInvoiceBehavior = "bill_immediately"
	SubscriptionProrationModelInvoiceBehaviorBillOnNextCollectionDate SubscriptionProrationModelInvoiceBehavior = "bill_on_next_collection_date"
)

type SubscriptionProrationModel struct {
	APIResource
	InvoiceBehavior SubscriptionProrationModelInvoiceBehavior `json:"invoice_behavior"`
	IsProrated      bool                                      `json:"is_prorated"`
	Rounding        *ProrationRoundingConfigModel             `json:"rounding,omitempty"`
}
