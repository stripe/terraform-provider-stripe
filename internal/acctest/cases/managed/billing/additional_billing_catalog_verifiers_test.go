// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	stripe "github.com/stripe/stripe-go/v86"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

type billingAlertExpectations struct {
	Address              string
	ExpectedAlertType    string
	ExpectedTitle        string
	ExpectedStatus       string
	ExpectedThresholdGTE int64
	ExpectedRecurrence   string
	ExpectedMeterAddress string
}

type billingCreditGrantExpectations struct {
	Address                 string
	ExpectedCustomerAddress string
	ExpectedCategory        string
	ExpectedName            string
	ExpectedPriority        int64
	ExpectedExpiresAt       int64
	ExpectedMetadata        map[string]string
	ExpectedAmountCurrency  string
	ExpectedAmountValue     int64
	ExpectedPriceType       string
}

type climateOrderExpectations struct {
	Address             string
	ExpectedProduct     string
	ExpectedMetricTons  float64
	ExpectedBeneficiary string
	ExpectedStatus      string
	ExpectedCurrency    string
	ExpectedAmountTotal int64
	ExpectedMetadata    map[string]string
}

type planExpectations struct {
	Address                 string
	CompareStateAttrs       []string
	ExpectedAmount          int64
	ExpectedIntervalCount   int64
	ExpectedTrialPeriodDays int64
	ExpectedUsageType       string
	ExpectedMetadata        map[string]string
	ExpectedProductAddress  string
}

type radarValueListExpectations struct {
	Address          string
	ExpectedName     string
	ExpectedItemType string
	ExpectedMetadata map[string]string
}

type radarValueListItemExpectations struct {
	Address                  string
	ExpectedValue            string
	ExpectedValueListAddress string
}

func verifyBillingAlert(expect billingAlertExpectations) runner.StateVerifier {
	return func(_ runner.TestEnv, client *stripe.Client, state *terraform.State) error {
		alert, err := retrieveBillingAlert(client, state, expect.Address)
		if err != nil {
			return err
		}
		if string(alert.AlertType) != expect.ExpectedAlertType {
			return fmt.Errorf("remote %s.alert_type mismatch: expected %q, got %q", expect.Address, expect.ExpectedAlertType, string(alert.AlertType))
		}
		if alert.Title != expect.ExpectedTitle {
			return fmt.Errorf("remote %s.title mismatch: expected %q, got %q", expect.Address, expect.ExpectedTitle, alert.Title)
		}
		if string(alert.Status) != expect.ExpectedStatus {
			return fmt.Errorf("remote %s.status mismatch: expected %q, got %q", expect.Address, expect.ExpectedStatus, string(alert.Status))
		}
		if alert.UsageThreshold == nil {
			return fmt.Errorf("remote %s.usage_threshold missing", expect.Address)
		}
		if alert.UsageThreshold.GTE != expect.ExpectedThresholdGTE {
			return fmt.Errorf("remote %s.usage_threshold.gte mismatch: expected %d, got %d", expect.Address, expect.ExpectedThresholdGTE, alert.UsageThreshold.GTE)
		}
		if string(alert.UsageThreshold.Recurrence) != expect.ExpectedRecurrence {
			return fmt.Errorf("remote %s.usage_threshold.recurrence mismatch: expected %q, got %q", expect.Address, expect.ExpectedRecurrence, string(alert.UsageThreshold.Recurrence))
		}
		expectedMeterID, err := runner.ResourcePrimaryID(state, expect.ExpectedMeterAddress)
		if err != nil {
			return err
		}
		if alert.UsageThreshold.Meter == nil || alert.UsageThreshold.Meter.ID != expectedMeterID {
			actual := ""
			if alert.UsageThreshold.Meter != nil {
				actual = alert.UsageThreshold.Meter.ID
			}
			return fmt.Errorf("remote %s.usage_threshold.meter mismatch: expected %q, got %q", expect.Address, expectedMeterID, actual)
		}
		return nil
	}
}

func verifyBillingAlertDestroyStateOnly(_ runner.TestEnv, client *stripe.Client, state *terraform.State) error {
	alert, err := retrieveBillingAlert(client, state, "stripe_billing_alert.test")
	if err != nil {
		return err
	}
	if string(alert.Status) == "" {
		return fmt.Errorf("expected stripe_billing_alert.test to remain retrievable after destroy")
	}
	return nil
}

func verifyBillingCreditGrant(expect billingCreditGrantExpectations) runner.StateVerifier {
	return func(_ runner.TestEnv, client *stripe.Client, state *terraform.State) error {
		grant, err := retrieveBillingCreditGrant(client, state, expect.Address)
		if err != nil {
			return err
		}
		expectedCustomerID, err := runner.ResourcePrimaryID(state, expect.ExpectedCustomerAddress)
		if err != nil {
			return err
		}
		if grant.Customer == nil || grant.Customer.ID != expectedCustomerID {
			actual := ""
			if grant.Customer != nil {
				actual = grant.Customer.ID
			}
			return fmt.Errorf("remote %s.customer mismatch: expected %q, got %q", expect.Address, expectedCustomerID, actual)
		}
		if string(grant.Category) != expect.ExpectedCategory {
			return fmt.Errorf("remote %s.category mismatch: expected %q, got %q", expect.Address, expect.ExpectedCategory, string(grant.Category))
		}
		if grant.Name != expect.ExpectedName {
			return fmt.Errorf("remote %s.name mismatch: expected %q, got %q", expect.Address, expect.ExpectedName, grant.Name)
		}
		if grant.Priority != expect.ExpectedPriority {
			return fmt.Errorf("remote %s.priority mismatch: expected %d, got %d", expect.Address, expect.ExpectedPriority, grant.Priority)
		}
		if grant.ExpiresAt != expect.ExpectedExpiresAt {
			return fmt.Errorf("remote %s.expires_at mismatch: expected %d, got %d", expect.Address, expect.ExpectedExpiresAt, grant.ExpiresAt)
		}
		if err := expectMetadataSubset(expect.Address+".metadata", grant.Metadata, expect.ExpectedMetadata); err != nil {
			return err
		}
		if grant.Amount == nil || grant.Amount.Monetary == nil {
			return fmt.Errorf("remote %s.amount.monetary missing", expect.Address)
		}
		if string(grant.Amount.Type) != "monetary" {
			return fmt.Errorf("remote %s.amount.type mismatch: expected %q, got %q", expect.Address, "monetary", string(grant.Amount.Type))
		}
		if string(grant.Amount.Monetary.Currency) != expect.ExpectedAmountCurrency {
			return fmt.Errorf("remote %s.amount.monetary.currency mismatch: expected %q, got %q", expect.Address, expect.ExpectedAmountCurrency, string(grant.Amount.Monetary.Currency))
		}
		if grant.Amount.Monetary.Value != expect.ExpectedAmountValue {
			return fmt.Errorf("remote %s.amount.monetary.value mismatch: expected %d, got %d", expect.Address, expect.ExpectedAmountValue, grant.Amount.Monetary.Value)
		}
		if grant.ApplicabilityConfig == nil || grant.ApplicabilityConfig.Scope == nil {
			return fmt.Errorf("remote %s.applicability_config.scope missing", expect.Address)
		}
		if string(grant.ApplicabilityConfig.Scope.PriceType) != expect.ExpectedPriceType {
			return fmt.Errorf("remote %s.applicability_config.scope.price_type mismatch: expected %q, got %q", expect.Address, expect.ExpectedPriceType, string(grant.ApplicabilityConfig.Scope.PriceType))
		}
		return nil
	}
}

func verifyBillingCreditGrantDestroyStateOnly(_ runner.TestEnv, client *stripe.Client, state *terraform.State) error {
	grant, err := retrieveBillingCreditGrant(client, state, "stripe_billing_credit_grant.test")
	if err != nil {
		return err
	}
	if grant.ID == "" {
		return fmt.Errorf("expected stripe_billing_credit_grant.test to remain retrievable after destroy")
	}
	return nil
}

func verifyClimateOrder(expect climateOrderExpectations) runner.StateVerifier {
	return func(_ runner.TestEnv, client *stripe.Client, state *terraform.State) error {
		order, err := retrieveClimateOrder(client, state, expect.Address)
		if err != nil {
			return err
		}
		if order.Product == nil || order.Product.ID != expect.ExpectedProduct {
			actual := ""
			if order.Product != nil {
				actual = order.Product.ID
			}
			return fmt.Errorf("remote %s.product mismatch: expected %q, got %q", expect.Address, expect.ExpectedProduct, actual)
		}
		if err := expectFloat(expect.Address+".metric_tons", order.MetricTons, expect.ExpectedMetricTons); err != nil {
			return err
		}
		if order.Beneficiary == nil || order.Beneficiary.PublicName != expect.ExpectedBeneficiary {
			actual := ""
			if order.Beneficiary != nil {
				actual = order.Beneficiary.PublicName
			}
			return fmt.Errorf("remote %s.beneficiary.public_name mismatch: expected %q, got %q", expect.Address, expect.ExpectedBeneficiary, actual)
		}
		if string(order.Status) != expect.ExpectedStatus {
			return fmt.Errorf("remote %s.status mismatch: expected %q, got %q", expect.Address, expect.ExpectedStatus, string(order.Status))
		}
		if string(order.Currency) != expect.ExpectedCurrency {
			return fmt.Errorf("remote %s.currency mismatch: expected %q, got %q", expect.Address, expect.ExpectedCurrency, string(order.Currency))
		}
		if order.AmountTotal != expect.ExpectedAmountTotal {
			return fmt.Errorf("remote %s.amount_total mismatch: expected %d, got %d", expect.Address, expect.ExpectedAmountTotal, order.AmountTotal)
		}
		return expectMetadataSubset(expect.Address+".metadata", order.Metadata, expect.ExpectedMetadata)
	}
}

func verifyClimateOrderDestroyStateOnly(_ runner.TestEnv, client *stripe.Client, state *terraform.State) error {
	order, err := retrieveClimateOrder(client, state, "stripe_climate_order.test")
	if err != nil {
		return err
	}
	if order.ID == "" {
		return fmt.Errorf("expected stripe_climate_order.test to remain retrievable after destroy")
	}
	return nil
}

func verifyPlan(expect planExpectations) runner.StateVerifier {
	return func(_ runner.TestEnv, client *stripe.Client, state *terraform.State) error {
		plan, err := retrievePlan(client, state, expect.Address)
		if err != nil {
			return err
		}
		for _, attr := range expect.CompareStateAttrs {
			actual, err := remotePlanString(plan, attr)
			if err != nil {
				return err
			}
			if err := expectRemoteStateString(state, expect.Address, attr, actual); err != nil {
				return err
			}
		}
		if plan.Amount != expect.ExpectedAmount {
			return fmt.Errorf("remote %s.amount mismatch: expected %d, got %d", expect.Address, expect.ExpectedAmount, plan.Amount)
		}
		if expect.ExpectedIntervalCount != 0 && plan.IntervalCount != expect.ExpectedIntervalCount {
			return fmt.Errorf("remote %s.interval_count mismatch: expected %d, got %d", expect.Address, expect.ExpectedIntervalCount, plan.IntervalCount)
		}
		if expect.ExpectedTrialPeriodDays != 0 && plan.TrialPeriodDays != expect.ExpectedTrialPeriodDays {
			return fmt.Errorf("remote %s.trial_period_days mismatch: expected %d, got %d", expect.Address, expect.ExpectedTrialPeriodDays, plan.TrialPeriodDays)
		}
		if string(plan.UsageType) != expect.ExpectedUsageType {
			return fmt.Errorf("remote %s.usage_type mismatch: expected %q, got %q", expect.Address, expect.ExpectedUsageType, string(plan.UsageType))
		}
		if err := expectMetadataSubset(expect.Address+".metadata", plan.Metadata, expect.ExpectedMetadata); err != nil {
			return err
		}
		expectedProductID, err := runner.ResourcePrimaryID(state, expect.ExpectedProductAddress)
		if err != nil {
			return err
		}
		if plan.Product == nil || plan.Product.ID != expectedProductID {
			actual := ""
			if plan.Product != nil {
				actual = plan.Product.ID
			}
			return fmt.Errorf("remote %s.product mismatch: expected %q, got %q", expect.Address, expectedProductID, actual)
		}
		return nil
	}
}

func verifyPlanDestroyDeleted(_ runner.TestEnv, client *stripe.Client, state *terraform.State) error {
	id, err := runner.ResourcePrimaryID(state, "stripe_plan.test")
	if err != nil {
		return err
	}
	_, err = client.V1Plans.Retrieve(context.Background(), id, nil)
	return expectRemoteMissing("stripe_plan.test", id, err)
}

func verifyRadarValueList(expect radarValueListExpectations) runner.StateVerifier {
	return func(_ runner.TestEnv, client *stripe.Client, state *terraform.State) error {
		list, err := retrieveRadarValueList(client, state, expect.Address)
		if err != nil {
			return err
		}
		if err := expectRemoteStateString(state, expect.Address, "alias", list.Alias); err != nil {
			return err
		}
		if list.Name != expect.ExpectedName {
			return fmt.Errorf("remote %s.name mismatch: expected %q, got %q", expect.Address, expect.ExpectedName, list.Name)
		}
		if string(list.ItemType) != expect.ExpectedItemType {
			return fmt.Errorf("remote %s.item_type mismatch: expected %q, got %q", expect.Address, expect.ExpectedItemType, string(list.ItemType))
		}
		return expectMetadataSubset(expect.Address+".metadata", list.Metadata, expect.ExpectedMetadata)
	}
}

func verifyRadarValueListDestroyDeleted(_ runner.TestEnv, client *stripe.Client, state *terraform.State) error {
	id, err := runner.ResourcePrimaryID(state, "stripe_radar_value_list.test")
	if err != nil {
		return err
	}
	_, err = client.V1RadarValueLists.Retrieve(context.Background(), id, nil)
	return expectRemoteMissing("stripe_radar_value_list.test", id, err)
}

func verifyRadarValueListItem(expect radarValueListItemExpectations) runner.StateVerifier {
	return func(_ runner.TestEnv, client *stripe.Client, state *terraform.State) error {
		item, err := retrieveRadarValueListItem(client, state, expect.Address)
		if err != nil {
			return err
		}
		if item.Value != expect.ExpectedValue {
			return fmt.Errorf("remote %s.value mismatch: expected %q, got %q", expect.Address, expect.ExpectedValue, item.Value)
		}
		expectedValueListID, err := runner.ResourcePrimaryID(state, expect.ExpectedValueListAddress)
		if err != nil {
			return err
		}
		if item.ValueList != expectedValueListID {
			return fmt.Errorf("remote %s.value_list mismatch: expected %q, got %q", expect.Address, expectedValueListID, item.ValueList)
		}
		return nil
	}
}

func verifyRadarValueListItemDestroyDeleted(_ runner.TestEnv, client *stripe.Client, state *terraform.State) error {
	id, err := runner.ResourcePrimaryID(state, "stripe_radar_value_list_item.test")
	if err != nil {
		return err
	}
	_, err = client.V1RadarValueListItems.Retrieve(context.Background(), id, nil)
	return expectRemoteMissing("stripe_radar_value_list_item.test", id, err)
}

func retrieveBillingAlert(client *stripe.Client, state *terraform.State, address string) (*stripe.BillingAlert, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}
	obj, err := client.V1BillingAlerts.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}
	return obj, nil
}

func retrieveBillingCreditGrant(client *stripe.Client, state *terraform.State, address string) (*stripe.BillingCreditGrant, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}
	obj, err := client.V1BillingCreditGrants.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}
	return obj, nil
}

func retrieveClimateOrder(client *stripe.Client, state *terraform.State, address string) (*stripe.ClimateOrder, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}
	obj, err := client.V1ClimateOrders.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}
	return obj, nil
}

func retrievePlan(client *stripe.Client, state *terraform.State, address string) (*stripe.Plan, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}
	obj, err := client.V1Plans.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}
	return obj, nil
}

func retrieveRadarValueList(client *stripe.Client, state *terraform.State, address string) (*stripe.RadarValueList, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}
	obj, err := client.V1RadarValueLists.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}
	return obj, nil
}

func retrieveRadarValueListItem(client *stripe.Client, state *terraform.State, address string) (*stripe.RadarValueListItem, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}
	obj, err := client.V1RadarValueListItems.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}
	return obj, nil
}

func remotePlanString(plan *stripe.Plan, attribute string) (string, error) {
	switch attribute {
	case "currency":
		return string(plan.Currency), nil
	case "interval":
		return string(plan.Interval), nil
	case "nickname":
		return plan.Nickname, nil
	default:
		return "", fmt.Errorf("unsupported plan attribute %q", attribute)
	}
}
