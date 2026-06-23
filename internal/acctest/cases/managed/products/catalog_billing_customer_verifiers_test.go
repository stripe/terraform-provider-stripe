// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	stripe "github.com/stripe/stripe-go/v86"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

type productPackageDimensionsExpectation struct {
	Height float64
	Length float64
	Weight float64
	Width  float64
}

type productDefaultPriceExpectation struct {
	Currency      string
	UnitAmount    int64
	Recurring     bool
	Interval      string
	IntervalCount int64
}

type stateStringExpectation struct {
	Attribute string
	Expected  string
}

type decimalStringExpectation struct {
	Attribute string
	Expected  string
}

type productExpectations struct {
	Address                   string
	CompareStateAttrs         []string
	CompareMetadata           bool
	CheckActive               bool
	ExpectedActive            bool
	CheckType                 bool
	ExpectedType              string
	CheckImages               bool
	ExpectedImages            []string
	CheckMarketingFeatures    bool
	ExpectedMarketingFeatures []string
	CheckShippable            bool
	ExpectedShippable         bool
	CheckPackageDimensions    bool
	ExpectedPackageDimensions productPackageDimensionsExpectation
	CheckDefaultPrice         bool
	ExpectedDefaultPrice      productDefaultPriceExpectation
	StateStrings              []stateStringExpectation
	DefaultPriceDecimals      []decimalStringExpectation
}

type priceRecurringExpectation struct {
	Interval  string
	UsageType string
}

type priceTierExpectation struct {
	UnitAmount int64
	UpTo       int64
}

type priceExpectations struct {
	Address                   string
	CompareStateAttrs         []string
	CompareMetadata           bool
	MissingStateAttrs         []string
	CheckActive               bool
	ExpectedActive            bool
	CheckType                 bool
	ExpectedType              string
	CheckBillingScheme        bool
	ExpectedBillingScheme     string
	CheckUnitAmount           bool
	ExpectedUnitAmount        int64
	CheckRecurring            bool
	ExpectedRecurring         priceRecurringExpectation
	CheckTiersMode            bool
	ExpectedTiersMode         string
	CheckTiers                bool
	ExpectedTiers             []priceTierExpectation
	CheckTransformQuantity    bool
	ExpectedTransformDivideBy int64
	ExpectedTransformRound    string
	StateStrings              []stateStringExpectation
	DecimalStrings            []decimalStringExpectation
}

type webhookEndpointExpectations struct {
	Address               string
	CompareStateAttrs     []string
	CompareMetadata       bool
	CheckEnabledEvents    bool
	ExpectedEnabledEvents []string
}

type eventDestinationExpectations struct {
	Address                 string
	CompareStateAttrs       []string
	CompareMetadata         bool
	CheckEnabledEvents      bool
	ExpectedEnabledEvents   []string
	CheckAmazonEventbridge  bool
	ExpectedAWSAccountID    string
	CheckWebhookEndpointURL bool
	ExpectedWebhookURLParts []string
}

type billingMeterExpectations struct {
	Address           string
	CompareStateAttrs []string
	CheckStatus       bool
	ExpectedStatus    string
}

type entitlementsFeatureExpectations struct {
	Address           string
	CompareStateAttrs []string
	CompareMetadata   bool
	CheckActive       bool
	ExpectedActive    bool
}

type couponExpectations struct {
	Address                  string
	CompareStateStringAttrs  []string
	CompareMetadata          bool
	CheckPercentOff          bool
	ExpectedPercentOff       float64
	CheckAmountOff           bool
	ExpectedAmountOff        int64
	CheckDurationInMonths    bool
	ExpectedDurationInMonths int64
	CheckAppliesToProduct    bool
}

type customerExpectations struct {
	Address           string
	CompareStateAttrs []string
	CompareMetadata   bool
}

type promotionCodeRestrictionsExpectation struct {
	MinimumAmount         int64
	MinimumAmountCurrency string
	FirstTimeTransaction  bool
}

type promotionCodeExpectations struct {
	Address              string
	CompareStateAttrs    []string
	CompareMetadata      bool
	CheckActive          bool
	ExpectedActive       bool
	CheckRestrictions    bool
	ExpectedRestrictions promotionCodeRestrictionsExpectation
}

func verifyCustomer(expect customerExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		customer, err := retrieveCustomer(client, state, expect.Address)
		if err != nil {
			return err
		}

		for _, attr := range expect.CompareStateAttrs {
			actual, err := remoteCustomerString(customer, attr)
			if err != nil {
				return err
			}
			if err := expectRemoteStateString(state, expect.Address, attr, actual); err != nil {
				return err
			}
		}
		if expect.CompareMetadata {
			if err := expectRemoteMetadataFromState(
				state,
				expect.Address,
				customer.Metadata,
			); err != nil {
				return err
			}
		}

		return nil
	}
}

func verifyCustomerDestroyDeleted(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	customer, err := retrieveCustomer(client, state, "stripe_customer.test")
	if err != nil {
		return err
	}
	if !customer.Deleted {
		return fmt.Errorf(
			"expected stripe_customer.test (%s) to be deleted remotely",
			customer.ID,
		)
	}

	return nil
}

func verifyProduct(expect productExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		product, err := retrieveProduct(client, state, expect.Address)
		if err != nil {
			return err
		}

		for _, attr := range expect.CompareStateAttrs {
			actual, err := remoteProductString(product, attr)
			if err != nil {
				return err
			}
			if err := expectRemoteStateString(state, expect.Address, attr, actual); err != nil {
				return err
			}
		}
		if expect.CompareMetadata {
			if err := expectRemoteMetadataFromState(
				state,
				expect.Address,
				product.Metadata,
			); err != nil {
				return err
			}
		}
		for _, stateString := range expect.StateStrings {
			if err := expectStateString(
				state,
				expect.Address,
				stateString.Attribute,
				stateString.Expected,
			); err != nil {
				return err
			}
		}
		if expect.CheckActive && product.Active != expect.ExpectedActive {
			return fmt.Errorf(
				"remote %s.active mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedActive,
				product.Active,
			)
		}
		if expect.CheckType && string(product.Type) != expect.ExpectedType {
			return fmt.Errorf(
				"remote %s.type mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedType,
				string(product.Type),
			)
		}
		if expect.CheckImages {
			if err := expectRemoteStringList(
				expect.Address+".images",
				product.Images,
				expect.ExpectedImages,
			); err != nil {
				return err
			}
		}
		if expect.CheckMarketingFeatures {
			actual := []string{}
			for _, feature := range product.MarketingFeatures {
				if feature == nil {
					continue
				}
				actual = append(actual, feature.Name)
			}
			if err := expectRemoteStringList(
				expect.Address+".marketing_features",
				actual,
				expect.ExpectedMarketingFeatures,
			); err != nil {
				return err
			}
		}
		if expect.CheckShippable && product.Shippable != expect.ExpectedShippable {
			return fmt.Errorf(
				"remote %s.shippable mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedShippable,
				product.Shippable,
			)
		}
		if expect.CheckPackageDimensions {
			if product.PackageDimensions == nil {
				return fmt.Errorf("remote %s.package_dimensions missing", expect.Address)
			}
			if err := expectFloat(
				expect.Address+".package_dimensions.height",
				product.PackageDimensions.Height,
				expect.ExpectedPackageDimensions.Height,
			); err != nil {
				return err
			}
			if err := expectFloat(
				expect.Address+".package_dimensions.length",
				product.PackageDimensions.Length,
				expect.ExpectedPackageDimensions.Length,
			); err != nil {
				return err
			}
			if err := expectFloat(
				expect.Address+".package_dimensions.weight",
				product.PackageDimensions.Weight,
				expect.ExpectedPackageDimensions.Weight,
			); err != nil {
				return err
			}
			if err := expectFloat(
				expect.Address+".package_dimensions.width",
				product.PackageDimensions.Width,
				expect.ExpectedPackageDimensions.Width,
			); err != nil {
				return err
			}
		}
		if expect.CheckDefaultPrice {
			defaultPrice, err := retrieveProductDefaultPrice(client, product, expect.Address)
			if err != nil {
				return err
			}
			if err := expectRemoteStateNonEmpty(state, expect.Address, "default_price"); err != nil {
				return err
			}
			if string(defaultPrice.Currency) != expect.ExpectedDefaultPrice.Currency {
				return fmt.Errorf(
					"remote %s.default_price.currency mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedDefaultPrice.Currency,
					string(defaultPrice.Currency),
				)
			}
			if defaultPrice.UnitAmount != expect.ExpectedDefaultPrice.UnitAmount {
				return fmt.Errorf(
					"remote %s.default_price.unit_amount mismatch: expected %d, got %d",
					expect.Address,
					expect.ExpectedDefaultPrice.UnitAmount,
					defaultPrice.UnitAmount,
				)
			}
			if expect.ExpectedDefaultPrice.Recurring {
				if defaultPrice.Recurring == nil {
					return fmt.Errorf("remote %s.default_price.recurring missing", expect.Address)
				}
				if string(defaultPrice.Recurring.Interval) != expect.ExpectedDefaultPrice.Interval {
					return fmt.Errorf(
						"remote %s.default_price.recurring.interval mismatch: expected %q, got %q",
						expect.Address,
						expect.ExpectedDefaultPrice.Interval,
						string(defaultPrice.Recurring.Interval),
					)
				}
				if defaultPrice.Recurring.IntervalCount != expect.ExpectedDefaultPrice.IntervalCount {
					return fmt.Errorf(
						"remote %s.default_price.recurring.interval_count mismatch: expected %d, got %d",
						expect.Address,
						expect.ExpectedDefaultPrice.IntervalCount,
						defaultPrice.Recurring.IntervalCount,
					)
				}
			}
		}
		if len(expect.DefaultPriceDecimals) > 0 {
			defaultPrice, err := retrieveProductDefaultPrice(client, product, expect.Address)
			if err != nil {
				return err
			}
			for _, decimalExpectation := range expect.DefaultPriceDecimals {
				actual, err := remotePriceDecimalString(defaultPrice, decimalExpectation.Attribute)
				if err != nil {
					return err
				}
				if err := expectEquivalentDecimalString(
					expect.Address+".default_price."+decimalExpectation.Attribute,
					actual,
					decimalExpectation.Expected,
				); err != nil {
					return err
				}
			}
		}

		return nil
	}
}

func verifyProductDestroyInactive(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	product, err := retrieveProduct(client, state, "stripe_product.test")
	if err != nil {
		return err
	}
	if product.Active {
		return fmt.Errorf("expected stripe_product.test to be inactive after destroy")
	}

	return nil
}

func verifyPrice(expect priceExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		price, err := retrievePrice(client, state, expect.Address)
		if err != nil {
			return err
		}

		for _, attr := range expect.CompareStateAttrs {
			actual, err := remotePriceString(price, attr)
			if err != nil {
				return err
			}
			if err := expectRemoteStateString(state, expect.Address, attr, actual); err != nil {
				return err
			}
		}
		if expect.CompareMetadata {
			if err := expectRemoteMetadataFromState(
				state,
				expect.Address,
				price.Metadata,
			); err != nil {
				return err
			}
		}
		for _, stateString := range expect.StateStrings {
			if err := expectStateString(
				state,
				expect.Address,
				stateString.Attribute,
				stateString.Expected,
			); err != nil {
				return err
			}
		}
		for _, missingAttr := range expect.MissingStateAttrs {
			if err := expectStateMissing(state, expect.Address, missingAttr); err != nil {
				return err
			}
		}
		if expect.CheckActive && price.Active != expect.ExpectedActive {
			return fmt.Errorf(
				"remote %s.active mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedActive,
				price.Active,
			)
		}
		if expect.CheckType && string(price.Type) != expect.ExpectedType {
			return fmt.Errorf(
				"remote %s.type mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedType,
				string(price.Type),
			)
		}
		if expect.CheckBillingScheme && string(price.BillingScheme) != expect.ExpectedBillingScheme {
			return fmt.Errorf(
				"remote %s.billing_scheme mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedBillingScheme,
				string(price.BillingScheme),
			)
		}
		if expect.CheckUnitAmount && price.UnitAmount != expect.ExpectedUnitAmount {
			return fmt.Errorf(
				"remote %s.unit_amount mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedUnitAmount,
				price.UnitAmount,
			)
		}
		if expect.CheckRecurring {
			if price.Recurring == nil {
				return fmt.Errorf("remote %s.recurring missing", expect.Address)
			}
			if string(price.Recurring.Interval) != expect.ExpectedRecurring.Interval {
				return fmt.Errorf(
					"remote %s.recurring.interval mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedRecurring.Interval,
					string(price.Recurring.Interval),
				)
			}
			if string(price.Recurring.UsageType) != expect.ExpectedRecurring.UsageType {
				return fmt.Errorf(
					"remote %s.recurring.usage_type mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedRecurring.UsageType,
					string(price.Recurring.UsageType),
				)
			}
		}
		if expect.CheckTiersMode && string(price.TiersMode) != expect.ExpectedTiersMode {
			return fmt.Errorf(
				"remote %s.tiers_mode mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedTiersMode,
				string(price.TiersMode),
			)
		}
		if expect.CheckTiers {
			if len(price.Tiers) != len(expect.ExpectedTiers) {
				return fmt.Errorf(
					"remote %s.tiers length mismatch: expected %d, got %d",
					expect.Address,
					len(expect.ExpectedTiers),
					len(price.Tiers),
				)
			}
			for index, expectedTier := range expect.ExpectedTiers {
				tier := price.Tiers[index]
				if tier == nil {
					return fmt.Errorf("remote %s.tiers[%d] is nil", expect.Address, index)
				}
				if tier.UnitAmount != expectedTier.UnitAmount {
					return fmt.Errorf(
						"remote %s.tiers[%d].unit_amount mismatch: expected %d, got %d",
						expect.Address,
						index,
						expectedTier.UnitAmount,
						tier.UnitAmount,
					)
				}
				if tier.UpTo != expectedTier.UpTo {
					return fmt.Errorf(
						"remote %s.tiers[%d].up_to mismatch: expected %d, got %d",
						expect.Address,
						index,
						expectedTier.UpTo,
						tier.UpTo,
					)
				}
			}
		}
		if expect.CheckTransformQuantity {
			if price.TransformQuantity == nil {
				return fmt.Errorf("remote %s.transform_quantity missing", expect.Address)
			}
			if price.TransformQuantity.DivideBy != expect.ExpectedTransformDivideBy {
				return fmt.Errorf(
					"remote %s.transform_quantity.divide_by mismatch: expected %d, got %d",
					expect.Address,
					expect.ExpectedTransformDivideBy,
					price.TransformQuantity.DivideBy,
				)
			}
			if string(price.TransformQuantity.Round) != expect.ExpectedTransformRound {
				return fmt.Errorf(
					"remote %s.transform_quantity.round mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedTransformRound,
					string(price.TransformQuantity.Round),
				)
			}
		}
		for _, decimalExpectation := range expect.DecimalStrings {
			actual, err := remotePriceDecimalString(price, decimalExpectation.Attribute)
			if err != nil {
				return err
			}
			if err := expectEquivalentDecimalString(
				expect.Address+"."+decimalExpectation.Attribute,
				actual,
				decimalExpectation.Expected,
			); err != nil {
				return err
			}
		}

		return nil
	}
}

func verifyPriceDestroyInactive(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	price, err := retrievePrice(client, state, "stripe_price.test")
	if err != nil {
		return err
	}
	if price.Active {
		return fmt.Errorf("expected stripe_price.test to be inactive after destroy")
	}

	return nil
}

func verifyWebhookEndpoint(
	expect webhookEndpointExpectations,
) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		webhookEndpoint, err := retrieveWebhookEndpoint(client, state, expect.Address)
		if err != nil {
			return err
		}

		for _, attr := range expect.CompareStateAttrs {
			actual, err := remoteWebhookEndpointString(webhookEndpoint, attr)
			if err != nil {
				return err
			}
			if err := expectRemoteStateString(state, expect.Address, attr, actual); err != nil {
				return err
			}
		}
		if expect.CompareMetadata {
			if err := expectRemoteMetadataFromState(
				state,
				expect.Address,
				webhookEndpoint.Metadata,
			); err != nil {
				return err
			}
		}
		if expect.CheckEnabledEvents {
			if err := expectRemoteStringList(
				expect.Address+".enabled_events",
				webhookEndpoint.EnabledEvents,
				expect.ExpectedEnabledEvents,
			); err != nil {
				return err
			}
		}

		return nil
	}
}

func verifyWebhookEndpointDestroyMissing(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	id, err := runner.ResourcePrimaryID(state, "stripe_webhook_endpoint.test")
	if err != nil {
		return err
	}

	_, err = client.V1WebhookEndpoints.Retrieve(context.Background(), id, nil)
	return expectRemoteMissing("stripe_webhook_endpoint.test", id, err)
}

func verifyEventDestination(
	expect eventDestinationExpectations,
) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		eventDestination, err := retrieveEventDestination(client, state, expect.Address)
		if err != nil {
			return err
		}

		for _, attr := range expect.CompareStateAttrs {
			actual, err := remoteEventDestinationString(eventDestination, attr)
			if err != nil {
				return err
			}
			if err := expectRemoteStateString(state, expect.Address, attr, actual); err != nil {
				return err
			}
		}
		if expect.CompareMetadata {
			if err := expectRemoteMetadataFromState(
				state,
				expect.Address,
				eventDestination.Metadata,
			); err != nil {
				return err
			}
		}
		if expect.CheckEnabledEvents {
			if err := expectRemoteStringList(
				expect.Address+".enabled_events",
				eventDestination.EnabledEvents,
				expect.ExpectedEnabledEvents,
			); err != nil {
				return err
			}
		}
		if expect.CheckAmazonEventbridge {
			if eventDestination.AmazonEventbridge == nil {
				return fmt.Errorf("remote %s.amazon_eventbridge missing", expect.Address)
			}
			if eventDestination.AmazonEventbridge.AwsAccountID != expect.ExpectedAWSAccountID {
				return fmt.Errorf(
					"remote %s.amazon_eventbridge.aws_account_id mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedAWSAccountID,
					eventDestination.AmazonEventbridge.AwsAccountID,
				)
			}
			if eventDestination.AmazonEventbridge.AwsEventSourceArn == "" {
				return fmt.Errorf(
					"remote %s.amazon_eventbridge.aws_event_source_arn unexpectedly empty",
					expect.Address,
				)
			}
			if string(eventDestination.AmazonEventbridge.AwsEventSourceStatus) == "" {
				return fmt.Errorf(
					"remote %s.amazon_eventbridge.aws_event_source_status unexpectedly empty",
					expect.Address,
				)
			}
		}
		if expect.CheckWebhookEndpointURL {
			if eventDestination.WebhookEndpoint == nil {
				return fmt.Errorf("remote %s.webhook_endpoint missing", expect.Address)
			}
			for _, part := range expect.ExpectedWebhookURLParts {
				if !strings.Contains(eventDestination.WebhookEndpoint.URL, part) {
					return fmt.Errorf(
						"remote %s.webhook_endpoint.url mismatch: expected %q to contain %q",
						expect.Address,
						eventDestination.WebhookEndpoint.URL,
						part,
					)
				}
			}
			if eventDestination.WebhookEndpoint.URL == "" {
				return fmt.Errorf(
					"remote %s.webhook_endpoint.url unexpectedly empty",
					expect.Address,
				)
			}
		}

		return nil
	}
}

func verifyEventDestinationDestroyMissing(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	for _, address := range []string{
		"stripe_v2_core_event_destination.eventbridge",
		"stripe_v2_core_event_destination.webhook",
	} {
		id, err := runner.ResourcePrimaryID(state, address)
		if err != nil {
			return err
		}

		_, err = client.V2CoreEventDestinations.Retrieve(context.Background(), id, nil)
		if err == nil {
			return fmt.Errorf("expected %s (%s) to be missing remotely", address, id)
		}
		if !strings.Contains(err.Error(), "\"status\":404") &&
			!strings.Contains(err.Error(), "\"code\":\"not_found\"") {
			return fmt.Errorf(
				"expected %s (%s) missing error, got: %w",
				address,
				id,
				err,
			)
		}
	}

	return nil
}

func verifyBillingMeter(
	expect billingMeterExpectations,
) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		billingMeter, err := retrieveBillingMeter(client, state, expect.Address)
		if err != nil {
			return err
		}

		for _, attr := range expect.CompareStateAttrs {
			actual, err := remoteBillingMeterString(billingMeter, attr)
			if err != nil {
				return err
			}
			if err := expectRemoteStateString(state, expect.Address, attr, actual); err != nil {
				return err
			}
		}
		if expect.CheckStatus && string(billingMeter.Status) != expect.ExpectedStatus {
			return fmt.Errorf(
				"remote %s.status mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedStatus,
				string(billingMeter.Status),
			)
		}

		return nil
	}
}

func verifyBillingMeterDestroyInactive(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	billingMeter, err := retrieveBillingMeter(client, state, "stripe_billing_meter.test")
	if err != nil {
		return err
	}
	if string(billingMeter.Status) != "inactive" {
		return fmt.Errorf(
			"expected stripe_billing_meter.test to be inactive after destroy, got %q",
			string(billingMeter.Status),
		)
	}
	if billingMeter.StatusTransitions == nil || billingMeter.StatusTransitions.DeactivatedAt == 0 {
		return fmt.Errorf(
			"expected stripe_billing_meter.test to have deactivated_at after destroy",
		)
	}

	return nil
}

func verifyEntitlementsFeature(
	expect entitlementsFeatureExpectations,
) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		feature, err := retrieveEntitlementsFeature(client, state, expect.Address)
		if err != nil {
			return err
		}

		for _, attr := range expect.CompareStateAttrs {
			actual, err := remoteEntitlementsFeatureString(feature, attr)
			if err != nil {
				return err
			}
			if err := expectRemoteStateString(state, expect.Address, attr, actual); err != nil {
				return err
			}
		}
		if expect.CompareMetadata {
			if err := expectRemoteMetadataFromState(
				state,
				expect.Address,
				feature.Metadata,
			); err != nil {
				return err
			}
		}
		if expect.CheckActive && feature.Active != expect.ExpectedActive {
			return fmt.Errorf(
				"remote %s.active mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedActive,
				feature.Active,
			)
		}

		return nil
	}
}

func verifyEntitlementsFeatureDestroyInactive(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	feature, err := retrieveEntitlementsFeature(
		client,
		state,
		"stripe_entitlements_feature.test",
	)
	if err != nil {
		return err
	}
	if feature.Active {
		return fmt.Errorf(
			"expected stripe_entitlements_feature.test to be inactive after destroy",
		)
	}

	return nil
}

func verifyCoupon(expect couponExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		coupon, err := retrieveCoupon(client, state, expect.Address)
		if err != nil {
			return err
		}

		for _, attr := range expect.CompareStateStringAttrs {
			actual, err := remoteCouponString(coupon, attr)
			if err != nil {
				return err
			}
			if err := expectRemoteStateString(state, expect.Address, attr, actual); err != nil {
				return err
			}
		}
		if expect.CompareMetadata {
			if err := expectRemoteMetadataFromState(
				state,
				expect.Address,
				coupon.Metadata,
			); err != nil {
				return err
			}
		}
		if expect.CheckPercentOff {
			if err := expectFloat(
				expect.Address+".percent_off",
				coupon.PercentOff,
				expect.ExpectedPercentOff,
			); err != nil {
				return err
			}
		}
		if expect.CheckAmountOff && coupon.AmountOff != expect.ExpectedAmountOff {
			return fmt.Errorf(
				"remote %s.amount_off mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedAmountOff,
				coupon.AmountOff,
			)
		}
		if expect.CheckDurationInMonths &&
			coupon.DurationInMonths != expect.ExpectedDurationInMonths {
			return fmt.Errorf(
				"remote %s.duration_in_months mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedDurationInMonths,
				coupon.DurationInMonths,
			)
		}
		if expect.CheckAppliesToProduct {
			expectedProductID, err := runner.ResourceAttribute(
				state,
				expect.Address,
				"applies_to.products.0",
			)
			if err != nil {
				return err
			}
			if err := expectRemoteStateString(
				state,
				expect.Address,
				"applies_to.products.0",
				expectedProductID,
			); err != nil {
				return err
			}
		}

		return nil
	}
}

func verifyCouponDestroyMissing(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	id, err := runner.ResourcePrimaryID(state, "stripe_coupon.test")
	if err != nil {
		return err
	}

	_, err = client.V1Coupons.Retrieve(context.Background(), id, nil)
	return expectRemoteMissing("stripe_coupon.test", id, err)
}

func verifyPromotionCode(
	expect promotionCodeExpectations,
) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		promotionCode, err := retrievePromotionCode(client, state, expect.Address)
		if err != nil {
			return err
		}

		for _, attr := range expect.CompareStateAttrs {
			actual, err := remotePromotionCodeString(promotionCode, attr)
			if err != nil {
				return err
			}
			if err := expectRemoteStateString(state, expect.Address, attr, actual); err != nil {
				return err
			}
		}
		if expect.CompareMetadata {
			if err := expectRemoteMetadataFromState(
				state,
				expect.Address,
				promotionCode.Metadata,
			); err != nil {
				return err
			}
		}
		if expect.CheckActive && promotionCode.Active != expect.ExpectedActive {
			return fmt.Errorf(
				"remote %s.active mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedActive,
				promotionCode.Active,
			)
		}
		if expect.CheckRestrictions {
			if promotionCode.Restrictions == nil {
				return fmt.Errorf("remote %s.restrictions missing", expect.Address)
			}
			if promotionCode.Restrictions.MinimumAmount != expect.ExpectedRestrictions.MinimumAmount {
				return fmt.Errorf(
					"remote %s.restrictions.minimum_amount mismatch: expected %d, got %d",
					expect.Address,
					expect.ExpectedRestrictions.MinimumAmount,
					promotionCode.Restrictions.MinimumAmount,
				)
			}
			if string(promotionCode.Restrictions.MinimumAmountCurrency) !=
				expect.ExpectedRestrictions.MinimumAmountCurrency {
				return fmt.Errorf(
					"remote %s.restrictions.minimum_amount_currency mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedRestrictions.MinimumAmountCurrency,
					string(promotionCode.Restrictions.MinimumAmountCurrency),
				)
			}
			if promotionCode.Restrictions.FirstTimeTransaction !=
				expect.ExpectedRestrictions.FirstTimeTransaction {
				return fmt.Errorf(
					"remote %s.restrictions.first_time_transaction mismatch: expected %t, got %t",
					expect.Address,
					expect.ExpectedRestrictions.FirstTimeTransaction,
					promotionCode.Restrictions.FirstTimeTransaction,
				)
			}
		}

		return nil
	}
}

func verifyPromotionCodeDestroyInactive(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	promotionCode, err := retrievePromotionCode(
		client,
		state,
		"stripe_promotion_code.test",
	)
	if err != nil {
		return err
	}
	if promotionCode.Active {
		return fmt.Errorf(
			"expected stripe_promotion_code.test to be inactive after destroy",
		)
	}

	return nil
}

func retrieveCustomer(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.Customer, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	customer, err := client.V1Customers.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return customer, nil
}

func retrievePrice(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.Price, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	params := &stripe.PriceRetrieveParams{}
	params.AddExpand("tiers")
	price, err := client.V1Prices.Retrieve(context.Background(), id, params)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return price, nil
}

func retrieveProduct(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.Product, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	product, err := client.V1Products.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return product, nil
}

func retrieveProductDefaultPrice(
	client *stripe.Client,
	product *stripe.Product,
	address string,
) (*stripe.Price, error) {
	if product.DefaultPrice == nil || product.DefaultPrice.ID == "" {
		return nil, fmt.Errorf("remote %s.default_price missing", address)
	}

	price, err := client.V1Prices.Retrieve(
		context.Background(),
		product.DefaultPrice.ID,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"retrieve %s.default_price (%s): %w",
			address,
			product.DefaultPrice.ID,
			err,
		)
	}

	return price, nil
}

func retrieveWebhookEndpoint(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.WebhookEndpoint, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	webhookEndpoint, err := client.V1WebhookEndpoints.Retrieve(
		context.Background(),
		id,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return webhookEndpoint, nil
}

func retrieveEventDestination(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.V2CoreEventDestination, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	params := &stripe.V2CoreEventDestinationRetrieveParams{
		Include: stripe.StringSlice([]string{"webhook_endpoint.url"}),
	}
	eventDestination, err := client.V2CoreEventDestinations.Retrieve(
		context.Background(),
		id,
		params,
	)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return eventDestination, nil
}

func retrieveBillingMeter(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.BillingMeter, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	billingMeter, err := client.V1BillingMeters.Retrieve(
		context.Background(),
		id,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return billingMeter, nil
}

func retrieveEntitlementsFeature(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.EntitlementsFeature, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	feature, err := client.V1EntitlementsFeatures.Retrieve(
		context.Background(),
		id,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return feature, nil
}

func retrieveCoupon(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.Coupon, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	coupon, err := client.V1Coupons.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return coupon, nil
}

func retrievePromotionCode(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.PromotionCode, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	promotionCode, err := client.V1PromotionCodes.Retrieve(
		context.Background(),
		id,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return promotionCode, nil
}

func expectRemoteStateString(
	state *terraform.State,
	address string,
	attribute string,
	actual string,
) error {
	expected, err := runner.ResourceAttribute(state, address, attribute)
	if err != nil {
		return err
	}
	if expected != actual {
		return fmt.Errorf(
			"remote %s.%s mismatch: expected %q, got %q",
			address,
			attribute,
			expected,
			actual,
		)
	}

	return nil
}

func expectStateString(
	state *terraform.State,
	address string,
	attribute string,
	expected string,
) error {
	actual, err := runner.ResourceAttribute(state, address, attribute)
	if err != nil {
		return err
	}
	if actual != expected {
		return fmt.Errorf(
			"state %s.%s mismatch: expected %q, got %q",
			address,
			attribute,
			expected,
			actual,
		)
	}

	return nil
}

func expectStateMissing(
	state *terraform.State,
	address string,
	attribute string,
) error {
	resourceState, err := runner.ResourceState(state, address)
	if err != nil {
		return err
	}
	if resourceState.Primary == nil {
		return fmt.Errorf("resource %s has no primary state", address)
	}
	if _, ok := resourceState.Primary.Attributes[attribute]; ok {
		return fmt.Errorf("state %s.%s unexpectedly present", address, attribute)
	}

	return nil
}

func expectRemoteStateNonEmpty(
	state *terraform.State,
	address string,
	attribute string,
) error {
	resourceState, err := runner.ResourceState(state, address)
	if err != nil {
		return err
	}
	if resourceState.Primary == nil {
		return fmt.Errorf("resource %s has no primary state", address)
	}

	actual, ok := resourceState.Primary.Attributes[attribute]
	if !ok {
		return fmt.Errorf("state %s.%s missing", address, attribute)
	}
	if actual == "" {
		return fmt.Errorf("state %s.%s unexpectedly empty", address, attribute)
	}

	return nil
}

func expectRemoteMetadataFromState(
	state *terraform.State,
	address string,
	actual map[string]string,
) error {
	expected, err := stateStringMap(state, address, "metadata.")
	if err != nil {
		return err
	}

	for key, expectedValue := range expected {
		actualValue, ok := actual[key]
		if !ok {
			return fmt.Errorf("remote %s.metadata missing key %q", address, key)
		}
		if actualValue != expectedValue {
			return fmt.Errorf(
				"remote %s.metadata[%q] mismatch: expected %q, got %q",
				address,
				key,
				expectedValue,
				actualValue,
			)
		}
	}

	return nil
}

func expectRemoteStringList(name string, actual []string, expected []string) error {
	if !reflect.DeepEqual(actual, expected) {
		return fmt.Errorf(
			"remote %s mismatch: expected %v, got %v",
			name,
			expected,
			actual,
		)
	}

	return nil
}

func expectFloat(name string, actual float64, expected float64) error {
	if math.Abs(actual-expected) > 0.000001 {
		return fmt.Errorf(
			"remote %s mismatch: expected %f, got %f",
			name,
			expected,
			actual,
		)
	}

	return nil
}

func equivalentDecimalString(left string, right string) bool {
	leftValue, ok := new(big.Rat).SetString(left)
	if !ok {
		return false
	}
	rightValue, ok := new(big.Rat).SetString(right)
	if !ok {
		return false
	}

	return leftValue.Cmp(rightValue) == 0
}

func expectEquivalentDecimalString(name string, actual string, expected string) error {
	if equivalentDecimalString(actual, expected) {
		return nil
	}
	return fmt.Errorf(
		"remote %s mismatch: expected decimal %q, got %q",
		name,
		expected,
		actual,
	)
}

func stateStringMap(
	state *terraform.State,
	address string,
	prefix string,
) (map[string]string, error) {
	resourceState, err := runner.ResourceState(state, address)
	if err != nil {
		return nil, err
	}
	if resourceState.Primary == nil {
		return nil, fmt.Errorf("resource %s has no primary state", address)
	}

	expected := map[string]string{}
	keys := []string{}
	for key := range resourceState.Primary.Attributes {
		if strings.HasPrefix(key, prefix) &&
			key != prefix+"%" &&
			!strings.HasSuffix(key, ".%") {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	for _, key := range keys {
		expected[strings.TrimPrefix(key, prefix)] = resourceState.Primary.Attributes[key]
	}

	return expected, nil
}

func remoteCustomerString(customer *stripe.Customer, attribute string) (string, error) {
	switch attribute {
	case "name":
		return customer.Name, nil
	case "email":
		return customer.Email, nil
	case "description":
		return customer.Description, nil
	case "phone":
		return customer.Phone, nil
	case "invoice_prefix":
		return customer.InvoicePrefix, nil
	case "tax_exempt":
		return string(customer.TaxExempt), nil
	default:
		return "", fmt.Errorf("unsupported customer attribute %q", attribute)
	}
}

func remoteProductString(product *stripe.Product, attribute string) (string, error) {
	switch attribute {
	case "name":
		return product.Name, nil
	case "description":
		return product.Description, nil
	case "tax_code":
		if product.TaxCode == nil {
			return "", nil
		}
		return product.TaxCode.ID, nil
	case "url":
		return product.URL, nil
	case "unit_label":
		return product.UnitLabel, nil
	case "statement_descriptor":
		return product.StatementDescriptor, nil
	default:
		return "", fmt.Errorf("unsupported product attribute %q", attribute)
	}
}

func remoteEventDestinationString(
	eventDestination *stripe.V2CoreEventDestination,
	attribute string,
) (string, error) {
	switch attribute {
	case "name":
		return eventDestination.Name, nil
	case "description":
		return eventDestination.Description, nil
	case "type":
		return string(eventDestination.Type), nil
	case "event_payload":
		return string(eventDestination.EventPayload), nil
	case "snapshot_api_version":
		return eventDestination.SnapshotAPIVersion, nil
	case "amazon_eventbridge.0.aws_account_id":
		if eventDestination.AmazonEventbridge == nil {
			return "", nil
		}
		return eventDestination.AmazonEventbridge.AwsAccountID, nil
	case "amazon_eventbridge.0.aws_event_source_arn":
		if eventDestination.AmazonEventbridge == nil {
			return "", nil
		}
		return eventDestination.AmazonEventbridge.AwsEventSourceArn, nil
	case "amazon_eventbridge.0.aws_event_source_status":
		if eventDestination.AmazonEventbridge == nil {
			return "", nil
		}
		return string(eventDestination.AmazonEventbridge.AwsEventSourceStatus), nil
	case "webhook_endpoint.0.url":
		if eventDestination.WebhookEndpoint == nil {
			return "", nil
		}
		return eventDestination.WebhookEndpoint.URL, nil
	default:
		return "", fmt.Errorf("unsupported event destination attribute %q", attribute)
	}
}

func remotePriceString(price *stripe.Price, attribute string) (string, error) {
	switch attribute {
	case "currency":
		return string(price.Currency), nil
	case "product":
		if price.Product == nil {
			return "", nil
		}
		return price.Product.ID, nil
	case "lookup_key":
		return price.LookupKey, nil
	case "nickname":
		return price.Nickname, nil
	case "tax_behavior":
		return string(price.TaxBehavior), nil
	default:
		return "", fmt.Errorf("unsupported price attribute %q", attribute)
	}
}

func formatDecimalFloat(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

func remotePriceDecimalString(price *stripe.Price, attribute string) (string, error) {
	parts := strings.Split(attribute, ".")
	if len(parts) == 0 {
		return "", fmt.Errorf("unsupported price decimal attribute %q", attribute)
	}

	switch parts[0] {
	case "unit_amount_decimal":
		if len(parts) != 1 {
			return "", fmt.Errorf("unsupported price decimal attribute %q", attribute)
		}
		return formatDecimalFloat(price.UnitAmountDecimal), nil
	case "tiers":
		if len(parts) != 3 {
			return "", fmt.Errorf("unsupported price decimal attribute %q", attribute)
		}
		index, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", fmt.Errorf("parse tier index for %q: %w", attribute, err)
		}
		if index < 0 || index >= len(price.Tiers) {
			return "", fmt.Errorf("remote price tier index %d out of range for %q", index, attribute)
		}
		tier := price.Tiers[index]
		if tier == nil {
			return "", fmt.Errorf("remote price tier %d missing for %q", index, attribute)
		}
		switch parts[2] {
		case "unit_amount_decimal":
			return formatDecimalFloat(tier.UnitAmountDecimal), nil
		case "flat_amount_decimal":
			return formatDecimalFloat(tier.FlatAmountDecimal), nil
		default:
			return "", fmt.Errorf("unsupported price decimal attribute %q", attribute)
		}
	case "currency_options":
		if len(parts) < 3 {
			return "", fmt.Errorf("unsupported price decimal attribute %q", attribute)
		}
		option, ok := price.CurrencyOptions[parts[1]]
		if !ok || option == nil {
			return "", fmt.Errorf("remote price currency option %q missing for %q", parts[1], attribute)
		}
		if len(parts) == 3 && parts[2] == "unit_amount_decimal" {
			return formatDecimalFloat(option.UnitAmountDecimal), nil
		}
		if len(parts) == 5 && parts[2] == "tiers" {
			index, err := strconv.Atoi(parts[3])
			if err != nil {
				return "", fmt.Errorf("parse currency tier index for %q: %w", attribute, err)
			}
			if index < 0 || index >= len(option.Tiers) {
				return "", fmt.Errorf(
					"remote price currency option tier index %d out of range for %q",
					index,
					attribute,
				)
			}
			tier := option.Tiers[index]
			if tier == nil {
				return "", fmt.Errorf("remote price currency option tier %d missing for %q", index, attribute)
			}
			switch parts[4] {
			case "unit_amount_decimal":
				return formatDecimalFloat(tier.UnitAmountDecimal), nil
			case "flat_amount_decimal":
				return formatDecimalFloat(tier.FlatAmountDecimal), nil
			default:
				return "", fmt.Errorf("unsupported price decimal attribute %q", attribute)
			}
		}
		return "", fmt.Errorf("unsupported price decimal attribute %q", attribute)
	default:
		return "", fmt.Errorf("unsupported price decimal attribute %q", attribute)
	}
}

func remoteWebhookEndpointString(
	webhookEndpoint *stripe.WebhookEndpoint,
	attribute string,
) (string, error) {
	switch attribute {
	case "api_version":
		return webhookEndpoint.APIVersion, nil
	case "url":
		return webhookEndpoint.URL, nil
	case "description":
		return webhookEndpoint.Description, nil
	default:
		return "", fmt.Errorf("unsupported webhook endpoint attribute %q", attribute)
	}
}

func remoteBillingMeterString(
	billingMeter *stripe.BillingMeter,
	attribute string,
) (string, error) {
	switch attribute {
	case "display_name":
		return billingMeter.DisplayName, nil
	case "event_name":
		return billingMeter.EventName, nil
	case "customer_mapping.event_payload_key":
		if billingMeter.CustomerMapping == nil {
			return "", nil
		}
		return billingMeter.CustomerMapping.EventPayloadKey, nil
	case "customer_mapping.type":
		if billingMeter.CustomerMapping == nil {
			return "", nil
		}
		return string(billingMeter.CustomerMapping.Type), nil
	case "default_aggregation.formula":
		if billingMeter.DefaultAggregation == nil {
			return "", nil
		}
		return string(billingMeter.DefaultAggregation.Formula), nil
	case "value_settings.event_payload_key":
		if billingMeter.ValueSettings == nil {
			return "", nil
		}
		return billingMeter.ValueSettings.EventPayloadKey, nil
	default:
		return "", fmt.Errorf("unsupported billing meter attribute %q", attribute)
	}
}

func remoteEntitlementsFeatureString(
	feature *stripe.EntitlementsFeature,
	attribute string,
) (string, error) {
	switch attribute {
	case "lookup_key":
		return feature.LookupKey, nil
	case "name":
		return feature.Name, nil
	default:
		return "", fmt.Errorf("unsupported entitlements feature attribute %q", attribute)
	}
}

func remoteCouponString(coupon *stripe.Coupon, attribute string) (string, error) {
	switch attribute {
	case "name":
		return coupon.Name, nil
	case "currency":
		return string(coupon.Currency), nil
	case "duration":
		return string(coupon.Duration), nil
	default:
		return "", fmt.Errorf("unsupported coupon attribute %q", attribute)
	}
}

func remotePromotionCodeString(
	promotionCode *stripe.PromotionCode,
	attribute string,
) (string, error) {
	switch attribute {
	case "code":
		return promotionCode.Code, nil
	case "customer":
		if promotionCode.Customer == nil {
			return "", nil
		}
		return promotionCode.Customer.ID, nil
	case "promotion.coupon":
		if promotionCode.Promotion == nil || promotionCode.Promotion.Coupon == nil {
			return "", nil
		}
		return promotionCode.Promotion.Coupon.ID, nil
	default:
		return "", fmt.Errorf("unsupported promotion code attribute %q", attribute)
	}
}

func expectRemoteMissing(address string, id string, err error) error {
	if err == nil {
		return fmt.Errorf("expected %s (%s) to be missing remotely", address, id)
	}

	var stripeErr *stripe.Error
	if errors.As(err, &stripeErr) && stripeErr.HTTPStatusCode == http.StatusNotFound {
		return nil
	}

	return fmt.Errorf("expected %s (%s) missing error, got: %w", address, id, err)
}
