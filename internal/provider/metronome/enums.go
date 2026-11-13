//
// File generated from our OpenAPI spec
//

package metronome

var managedEntityValues = map[string]struct{}{"alert": {},
	"billable_metric":          {},
	"charge":                   {},
	"commit":                   {},
	"contract":                 {},
	"contract_credit":          {},
	"contract_product":         {},
	"customer":                 {},
	"discount":                 {},
	"invoice":                  {},
	"package_commit":           {},
	"package_credit":           {},
	"package_scheduled_charge": {},
	"package_subscription":     {},
	"product":                  {},
	"professional_service":     {},
	"rate_card":                {},
	"scheduled_charge":         {},
	"subscription":             {}}
var customerAlertTypeValues = map[string]struct{}{"invoice_total_reached": {},
	"low_remaining_commit_balance_reached":                        {},
	"low_remaining_commit_percentage_reached":                     {},
	"low_remaining_contract_credit_and_commit_balance_reached":    {},
	"low_remaining_contract_credit_and_commit_percentage_reached": {},
	"low_remaining_contract_credit_balance_reached":               {},
	"low_remaining_contract_credit_percentage_reached":            {},
	"low_remaining_days_for_commit_segment_reached":               {},
	"low_remaining_days_for_contract_credit_segment_reached":      {},
	"low_remaining_seat_balance_reached":                          {},
	"monthly_invoice_total_spend_threshold_reached":               {},
	"spend_threshold_reached":                                     {},
	"usage_threshold_reached":                                     {}}
var billingProviderValues = map[string]struct{}{"aws_marketplace": {},
	"azure_marketplace": {},
	"custom":            {},
	"gcp_marketplace":   {},
	"metronome":         {},
	"netsuite":          {},
	"quickbooks_online": {},
	"stripe":            {},
	"workday":           {}}
var revenueProviderValues = map[string]struct{}{"netsuite": {}}
var deliveryMethodValues = map[string]struct{}{"aws_sns": {},
	"aws_sqs":                    {},
	"direct_to_billing_provider": {},
	"tackle":                     {}}
var taxProviderValues = map[string]struct{}{"anrok": {},
	"avalara": {},
	"stripe":  {}}
