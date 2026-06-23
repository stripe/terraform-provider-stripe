// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import "testing"

func TestAccManagedBillingAlertBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"billing_alert_basic",
		"stripe_billing_alert",
		"stripe_billing_alert.test",
		"billing_alert_basic_create.tf",
		"",
		true,
		verifyBillingAlert(billingAlertExpectations{
			Address:              "stripe_billing_alert.test",
			ExpectedAlertType:    "usage_threshold",
			ExpectedTitle:        "Codex Billing Alert",
			ExpectedStatus:       "active",
			ExpectedThresholdGTE: 100,
			ExpectedRecurrence:   "one_time",
			ExpectedMeterAddress: "stripe_billing_meter.meter",
		}),
		nil,
		verifyBillingAlertDestroyStateOnly,
	)
}

func TestAccManagedBillingAlertRecurringThreshold(t *testing.T) {
	runBaseManagedCase(
		t,
		"billing_alert_recurring_threshold",
		"stripe_billing_alert",
		"stripe_billing_alert.test",
		"billing_alert_recurring_threshold_create.tf",
		"billing_alert_recurring_threshold_update.tf",
		true,
		verifyBillingAlert(billingAlertExpectations{
			Address:              "stripe_billing_alert.test",
			ExpectedAlertType:    "usage_threshold",
			ExpectedTitle:        "Codex Billing Alert Recurring",
			ExpectedStatus:       "active",
			ExpectedThresholdGTE: 250,
			ExpectedRecurrence:   "one_time",
			ExpectedMeterAddress: "stripe_billing_meter.meter",
		}),
		verifyBillingAlert(billingAlertExpectations{
			Address:              "stripe_billing_alert.test",
			ExpectedAlertType:    "usage_threshold",
			ExpectedTitle:        "Codex Billing Alert Recurring Updated",
			ExpectedStatus:       "active",
			ExpectedThresholdGTE: 350,
			ExpectedRecurrence:   "one_time",
			ExpectedMeterAddress: "stripe_billing_meter.meter",
		}),
		verifyBillingAlertDestroyStateOnly,
	)
}

func TestAccManagedBillingCreditGrantBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"billing_credit_grant_basic",
		"stripe_billing_credit_grant",
		"stripe_billing_credit_grant.test",
		"billing_credit_grant_basic_create.tf",
		"billing_credit_grant_basic_update.tf",
		true,
		verifyBillingCreditGrant(billingCreditGrantExpectations{
			Address:                 "stripe_billing_credit_grant.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedCategory:        "promotional",
			ExpectedName:            "SDK Codegen Credit Grant",
			ExpectedPriority:        50,
			ExpectedExpiresAt:       1893456000,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "billing_credit_grant_basic",
				"phase": "create",
			},
			ExpectedAmountCurrency: "usd",
			ExpectedAmountValue:    500,
			ExpectedPriceType:      "metered",
		}),
		verifyBillingCreditGrant(billingCreditGrantExpectations{
			Address:                 "stripe_billing_credit_grant.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedCategory:        "promotional",
			ExpectedName:            "SDK Codegen Credit Grant",
			ExpectedPriority:        50,
			ExpectedExpiresAt:       1896134400,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "billing_credit_grant_basic",
				"phase": "update",
			},
			ExpectedAmountCurrency: "usd",
			ExpectedAmountValue:    500,
			ExpectedPriceType:      "metered",
		}),
		verifyBillingCreditGrantDestroyStateOnly,
	)
}

func TestAccManagedBillingCreditGrantAmountPriorityUpdate(t *testing.T) {
	runBaseManagedCase(
		t,
		"billing_credit_grant_amount_priority_update",
		"stripe_billing_credit_grant",
		"stripe_billing_credit_grant.test",
		"billing_credit_grant_amount_priority_update_create.tf",
		"billing_credit_grant_amount_priority_update_update.tf",
		true,
		verifyBillingCreditGrant(billingCreditGrantExpectations{
			Address:                 "stripe_billing_credit_grant.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedCategory:        "promotional",
			ExpectedName:            "SDK Codegen Credit Grant Alt",
			ExpectedPriority:        10,
			ExpectedExpiresAt:       1893456000,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "billing_credit_grant_amount_priority_update",
				"phase": "create",
			},
			ExpectedAmountCurrency: "usd",
			ExpectedAmountValue:    300,
			ExpectedPriceType:      "metered",
		}),
		verifyBillingCreditGrant(billingCreditGrantExpectations{
			Address:                 "stripe_billing_credit_grant.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedCategory:        "promotional",
			ExpectedName:            "SDK Codegen Credit Grant Alt Updated",
			ExpectedPriority:        90,
			ExpectedExpiresAt:       1896134400,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "billing_credit_grant_amount_priority_update",
				"phase": "update",
			},
			ExpectedAmountCurrency: "usd",
			ExpectedAmountValue:    900,
			ExpectedPriceType:      "metered",
		}),
		verifyBillingCreditGrantDestroyStateOnly,
	)
}

func TestAccManagedClimateOrderBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"climate_order_basic",
		"stripe_climate_order",
		"stripe_climate_order.test",
		"climate_order_basic_create.tf",
		"climate_order_basic_update.tf",
		true,
		verifyClimateOrder(climateOrderExpectations{
			Address:             "stripe_climate_order.test",
			ExpectedProduct:     "climsku_frontier_offtake_portfolio_2027",
			ExpectedMetricTons:  0.1,
			ExpectedBeneficiary: "SDK Codegen Climate Create",
			ExpectedStatus:      "confirmed",
			ExpectedCurrency:    "usd",
			ExpectedAmountTotal: 5150,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "climate_order_basic",
				"phase": "create",
			},
		}),
		verifyClimateOrder(climateOrderExpectations{
			Address:             "stripe_climate_order.test",
			ExpectedProduct:     "climsku_frontier_offtake_portfolio_2027",
			ExpectedMetricTons:  0.1,
			ExpectedBeneficiary: "SDK Codegen Climate Update",
			ExpectedStatus:      "confirmed",
			ExpectedCurrency:    "usd",
			ExpectedAmountTotal: 5150,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "climate_order_basic",
				"phase": "update",
			},
		}),
		verifyClimateOrderDestroyStateOnly,
	)
}

func TestAccManagedClimateOrderExplicitCurrency(t *testing.T) {
	runBaseManagedCase(
		t,
		"climate_order_explicit_currency",
		"stripe_climate_order",
		"stripe_climate_order.test",
		"climate_order_explicit_currency_create.tf",
		"climate_order_explicit_currency_update.tf",
		true,
		verifyClimateOrder(climateOrderExpectations{
			Address:             "stripe_climate_order.test",
			ExpectedProduct:     "climsku_frontier_offtake_portfolio_2027",
			ExpectedMetricTons:  0.1,
			ExpectedBeneficiary: "SDK Codegen Climate Currency Create",
			ExpectedStatus:      "confirmed",
			ExpectedCurrency:    "usd",
			ExpectedAmountTotal: 5150,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "climate_order_explicit_currency",
				"phase": "create",
			},
		}),
		verifyClimateOrder(climateOrderExpectations{
			Address:             "stripe_climate_order.test",
			ExpectedProduct:     "climsku_frontier_offtake_portfolio_2027",
			ExpectedMetricTons:  0.1,
			ExpectedBeneficiary: "SDK Codegen Climate Currency Update",
			ExpectedStatus:      "confirmed",
			ExpectedCurrency:    "usd",
			ExpectedAmountTotal: 5150,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "climate_order_explicit_currency",
				"phase": "update",
			},
		}),
		verifyClimateOrderDestroyStateOnly,
	)
}

func TestAccManagedPlanBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"plan_basic",
		"stripe_plan",
		"stripe_plan.test",
		"plan_basic_create.tf",
		"plan_basic_update.tf",
		true,
		verifyPlan(planExpectations{
			Address:           "stripe_plan.test",
			CompareStateAttrs: []string{"currency", "interval", "nickname"},
			ExpectedAmount:    1200,
			ExpectedUsageType: "licensed",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "plan_basic",
				"phase": "create",
			},
			ExpectedProductAddress: "stripe_product.product",
		}),
		verifyPlan(planExpectations{
			Address:           "stripe_plan.test",
			CompareStateAttrs: []string{"currency", "interval", "nickname"},
			ExpectedAmount:    1200,
			ExpectedUsageType: "licensed",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "plan_basic",
				"phase": "update",
			},
			ExpectedProductAddress: "stripe_product.product",
		}),
		verifyPlanDestroyDeleted,
	)
}

func TestAccManagedPlanIntervalTrial(t *testing.T) {
	runBaseManagedCase(
		t,
		"plan_interval_trial",
		"stripe_plan",
		"stripe_plan.test",
		"plan_interval_trial_metered_create.tf",
		"",
		false,
		verifyPlan(planExpectations{
			Address:                 "stripe_plan.test",
			CompareStateAttrs:       []string{"currency", "interval", "nickname"},
			ExpectedAmount:          3500,
			ExpectedIntervalCount:   3,
			ExpectedTrialPeriodDays: 14,
			ExpectedUsageType:       "licensed",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "plan_interval_trial",
			},
			ExpectedProductAddress: "stripe_product.product",
		}),
		nil,
		verifyPlanDestroyDeleted,
	)
}

func TestAccManagedRadarValueListBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"radar_value_list_basic",
		"stripe_radar_value_list",
		"stripe_radar_value_list.test",
		"radar_value_list_basic_create.tf",
		"radar_value_list_basic_update.tf",
		true,
		verifyRadarValueList(radarValueListExpectations{
			Address:          "stripe_radar_value_list.test",
			ExpectedName:     "SDK Codegen Radar List Create",
			ExpectedItemType: "string",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "radar_value_list_basic",
				"phase": "create",
			},
		}),
		verifyRadarValueList(radarValueListExpectations{
			Address:          "stripe_radar_value_list.test",
			ExpectedName:     "SDK Codegen Radar List Update",
			ExpectedItemType: "string",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "radar_value_list_basic",
				"phase": "update",
			},
		}),
		verifyRadarValueListDestroyDeleted,
	)
}

func TestAccManagedRadarValueListEmailAliasUpdate(t *testing.T) {
	runBaseManagedCase(
		t,
		"radar_value_list_email_alias_update",
		"stripe_radar_value_list",
		"stripe_radar_value_list.test",
		"radar_value_list_email_alias_update_create.tf",
		"radar_value_list_email_alias_update_update.tf",
		true,
		verifyRadarValueList(radarValueListExpectations{
			Address:          "stripe_radar_value_list.test",
			ExpectedName:     "SDK Codegen Radar Email List Create",
			ExpectedItemType: "email",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "radar_value_list_email_alias_update",
				"phase": "create",
			},
		}),
		verifyRadarValueList(radarValueListExpectations{
			Address:          "stripe_radar_value_list.test",
			ExpectedName:     "SDK Codegen Radar Email List Update",
			ExpectedItemType: "email",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "radar_value_list_email_alias_update",
				"phase": "update",
			},
		}),
		verifyRadarValueListDestroyDeleted,
	)
}

func TestAccManagedRadarValueListItemBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"radar_value_list_item_basic",
		"stripe_radar_value_list_item",
		"stripe_radar_value_list_item.test",
		"radar_value_list_item_basic_create.tf",
		"",
		true,
		verifyRadarValueListItem(radarValueListItemExpectations{
			Address:                  "stripe_radar_value_list_item.test",
			ExpectedValue:            "sdk-codegen@example.com",
			ExpectedValueListAddress: "stripe_radar_value_list.list",
		}),
		nil,
		verifyRadarValueListItemDestroyDeleted,
	)
}
