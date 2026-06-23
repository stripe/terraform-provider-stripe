// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	stripe "github.com/stripe/stripe-go/v86"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

const legacyUpgradeProviderVersion = "0.2.2"

func runBaseManagedCase(
	t *testing.T,
	name string,
	surface string,
	resourceAddress string,
	createTemplate string,
	updateTemplate string,
	importStable bool,
	verifyCreate runner.StateVerifier,
	verifyUpdate runner.StateVerifier,
	verifyDestroy runner.StateVerifier,
) {
	t.Helper()

	runner.RunManagedCase(t, runner.ManagedCase{
		Definition: runner.CaseDefinition{
			Name:    name,
			Surface: surface,
			Group:   "base",
			Kind:    "resource",
		},
		ResourceAddress: resourceAddress,
		CreateTemplate:  createTemplate,
		UpdateTemplate:  updateTemplate,
		ImportStable:    importStable,
		VerifyCreate:    verifyCreate,
		VerifyUpdate:    verifyUpdate,
		VerifyDestroy:   verifyDestroy,
	})
}

func runBaseManagedLegacyUpgradeCase(
	t *testing.T,
	name string,
	surface string,
	resourceAddress string,
	createTemplate string,
	verifyCreate runner.StateVerifier,
	verifyDestroy runner.StateVerifier,
) {
	t.Helper()

	runner.RunManagedLegacyUpgradeCase(t, runner.ManagedCase{
		Definition: runner.CaseDefinition{
			Name:    name,
			Surface: surface,
			Group:   "base",
			Kind:    "resource",
		},
		ResourceAddress: resourceAddress,
		CreateTemplate:  createTemplate,
		VerifyCreate:    verifyCreate,
		VerifyDestroy:   verifyDestroy,
	}, legacyUpgradeProviderVersion)
}

func runBaseManagedCaseWithImportIgnore(
	t *testing.T,
	name string,
	surface string,
	resourceAddress string,
	createTemplate string,
	updateTemplate string,
	importStable bool,
	importStateVerifyIgnore []string,
	verifyCreate runner.StateVerifier,
	verifyUpdate runner.StateVerifier,
	verifyDestroy runner.StateVerifier,
) {
	t.Helper()

	runner.RunManagedCase(t, runner.ManagedCase{
		Definition: runner.CaseDefinition{
			Name:    name,
			Surface: surface,
			Group:   "base",
			Kind:    "resource",
		},
		ResourceAddress:         resourceAddress,
		CreateTemplate:          createTemplate,
		UpdateTemplate:          updateTemplate,
		ImportStable:            importStable,
		ImportStateVerifyIgnore: importStateVerifyIgnore,
		VerifyCreate:            verifyCreate,
		VerifyUpdate:            verifyUpdate,
		VerifyDestroy:           verifyDestroy,
	})
}

func TestAccManagedProductBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"product_basic",
		"stripe_product",
		"stripe_product.test",
		"product_create.tf",
		"product_update.tf",
		true,
		verifyProduct(productExpectations{
			Address:           "stripe_product.test",
			CompareStateAttrs: []string{"name", "description", "tax_code"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    true,
			CheckType:         true,
			ExpectedType:      "service",
		}),
		verifyProduct(productExpectations{
			Address:           "stripe_product.test",
			CompareStateAttrs: []string{"name", "description", "tax_code"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    true,
			CheckType:         true,
			ExpectedType:      "service",
		}),
		verifyProductDestroyInactive,
	)
}

func TestAccManagedProductActiveToggle(t *testing.T) {
	runBaseManagedCase(
		t,
		"product_active_toggle",
		"stripe_product",
		"stripe_product.test",
		"product_active_create.tf",
		"product_active_update.tf",
		true,
		verifyProduct(productExpectations{
			Address:           "stripe_product.test",
			CompareStateAttrs: []string{"name"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    true,
		}),
		verifyProduct(productExpectations{
			Address:           "stripe_product.test",
			CompareStateAttrs: []string{"name"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    false,
		}),
		verifyProductDestroyInactive,
	)
}

func TestAccManagedProductImagesURL(t *testing.T) {
	runBaseManagedCase(
		t,
		"product_images_url",
		"stripe_product",
		"stripe_product.test",
		"product_images_create.tf",
		"product_images_update.tf",
		true,
		verifyProduct(productExpectations{
			Address:           "stripe_product.test",
			CompareStateAttrs: []string{"name", "url"},
			CompareMetadata:   true,
			CheckImages:       true,
			ExpectedImages: []string{
				"https://example.com/sdk-codegen/product-a.png",
			},
		}),
		verifyProduct(productExpectations{
			Address:           "stripe_product.test",
			CompareStateAttrs: []string{"name", "url"},
			CompareMetadata:   true,
			CheckImages:       true,
			ExpectedImages: []string{
				"https://example.com/sdk-codegen/product-b.png",
				"https://example.com/sdk-codegen/product-c.png",
			},
		}),
		verifyProductDestroyInactive,
	)
}

func TestAccManagedProductMarketingFeatures(t *testing.T) {
	runBaseManagedCase(
		t,
		"product_marketing_features",
		"stripe_product",
		"stripe_product.test",
		"product_marketing_features_create.tf",
		"product_marketing_features_update.tf",
		true,
		verifyProduct(productExpectations{
			Address:                   "stripe_product.test",
			CompareStateAttrs:         []string{"name", "unit_label"},
			CompareMetadata:           true,
			CheckMarketingFeatures:    true,
			ExpectedMarketingFeatures: []string{"Priority support", "Usage insights"},
		}),
		verifyProduct(productExpectations{
			Address:                   "stripe_product.test",
			CompareStateAttrs:         []string{"name", "unit_label"},
			CompareMetadata:           true,
			CheckMarketingFeatures:    true,
			ExpectedMarketingFeatures: []string{"Dedicated onboarding", "Usage insights"},
		}),
		verifyProductDestroyInactive,
	)
}

func TestAccManagedProductShippable(t *testing.T) {
	runBaseManagedCase(
		t,
		"product_shippable",
		"stripe_product",
		"stripe_product.test",
		"product_shippable_create.tf",
		"product_shippable_update.tf",
		true,
		verifyProduct(productExpectations{
			Address:                "stripe_product.test",
			CompareStateAttrs:      []string{"name"},
			CompareMetadata:        true,
			CheckType:              true,
			ExpectedType:           "good",
			CheckShippable:         true,
			ExpectedShippable:      true,
			CheckPackageDimensions: true,
			ExpectedPackageDimensions: productPackageDimensionsExpectation{
				Height: 4.25,
				Length: 8.50,
				Weight: 16.00,
				Width:  2.75,
			},
		}),
		verifyProduct(productExpectations{
			Address:                "stripe_product.test",
			CompareStateAttrs:      []string{"name"},
			CompareMetadata:        true,
			CheckType:              true,
			ExpectedType:           "good",
			CheckShippable:         true,
			ExpectedShippable:      true,
			CheckPackageDimensions: true,
			ExpectedPackageDimensions: productPackageDimensionsExpectation{
				Height: 4.25,
				Length: 8.50,
				Weight: 16.00,
				Width:  2.75,
			},
		}),
		verifyProductDestroyInactive,
	)
}

func TestAccManagedProductStatementDescriptor(t *testing.T) {
	runBaseManagedCase(
		t,
		"product_statement_descriptor",
		"stripe_product",
		"stripe_product.test",
		"product_statement_descriptor_create.tf",
		"product_statement_descriptor_update.tf",
		true,
		verifyProduct(productExpectations{
			Address: "stripe_product.test",
			CompareStateAttrs: []string{
				"name",
				"description",
				"tax_code",
				"statement_descriptor",
			},
			CompareMetadata: true,
			CheckType:       true,
			ExpectedType:    "service",
		}),
		verifyProduct(productExpectations{
			Address: "stripe_product.test",
			CompareStateAttrs: []string{
				"name",
				"description",
				"tax_code",
				"statement_descriptor",
			},
			CompareMetadata: true,
			CheckType:       true,
			ExpectedType:    "service",
		}),
		verifyProductDestroyInactive,
	)
}

func TestAccManagedProductDefaultPriceData(t *testing.T) {
	runBaseManagedCase(
		t,
		"product_default_price_data",
		"stripe_product",
		"stripe_product.test",
		"product_default_price_data_create.tf",
		"product_default_price_data_update.tf",
		false,
		verifyProduct(productExpectations{
			Address:           "stripe_product.test",
			CompareStateAttrs: []string{"name", "description"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    true,
			CheckDefaultPrice: true,
			ExpectedDefaultPrice: productDefaultPriceExpectation{
				Currency:      "usd",
				UnitAmount:    1500,
				Recurring:     true,
				Interval:      "month",
				IntervalCount: 1,
			},
		}),
		verifyProduct(productExpectations{
			Address:           "stripe_product.test",
			CompareStateAttrs: []string{"name", "description"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    true,
			CheckDefaultPrice: true,
			ExpectedDefaultPrice: productDefaultPriceExpectation{
				Currency:      "usd",
				UnitAmount:    1500,
				Recurring:     true,
				Interval:      "month",
				IntervalCount: 1,
			},
		}),
		verifyProductDestroyInactive,
	)
}

func TestAccManagedProductDefaultPriceDataLegacyUpgrade(t *testing.T) {
	runBaseManagedLegacyUpgradeCase(
		t,
		"product_default_price_data_legacy_upgrade",
		"stripe_product",
		"stripe_product.test",
		"product_default_price_data_create.tf",
		verifyProduct(productExpectations{
			Address:           "stripe_product.test",
			CompareStateAttrs: []string{"name", "description"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    true,
			CheckDefaultPrice: true,
			ExpectedDefaultPrice: productDefaultPriceExpectation{
				Currency:      "usd",
				UnitAmount:    1500,
				Recurring:     true,
				Interval:      "month",
				IntervalCount: 1,
			},
		}),
		verifyProductDestroyInactive,
	)
}

func TestAccManagedProductDefaultPriceDataDecimalAmounts(t *testing.T) {
	runBaseManagedCaseWithImportIgnore(
		t,
		"product_default_price_data_decimal_amounts",
		"stripe_product",
		"stripe_product.test",
		"product_default_price_data_decimal_amounts_create.tf",
		"product_default_price_data_decimal_amounts_update.tf",
		false,
		[]string{"default_price_data"},
		verifyProduct(productExpectations{
			Address:           "stripe_product.test",
			CompareStateAttrs: []string{"name", "description"},
			CompareMetadata:   true,
			DefaultPriceDecimals: []decimalStringExpectation{
				{Attribute: "unit_amount_decimal", Expected: "1500.0"},
			},
		}),
		verifyProduct(productExpectations{
			Address:           "stripe_product.test",
			CompareStateAttrs: []string{"name", "description"},
			CompareMetadata:   true,
			DefaultPriceDecimals: []decimalStringExpectation{
				{Attribute: "unit_amount_decimal", Expected: "1700.00"},
			},
		}),
		verifyProductDestroyInactive,
	)
}

func TestAccManagedPriceBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"price_basic",
		"stripe_price",
		"stripe_price.test",
		"price_basic_create.tf",
		"price_basic_update.tf",
		true,
		verifyPrice(priceExpectations{
			Address:               "stripe_price.test",
			CompareStateAttrs:     []string{"currency", "product"},
			CompareMetadata:       true,
			CheckActive:           true,
			ExpectedActive:        true,
			CheckType:             true,
			ExpectedType:          "recurring",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "per_unit",
			CheckUnitAmount:       true,
			ExpectedUnitAmount:    1500,
			CheckRecurring:        true,
			ExpectedRecurring: priceRecurringExpectation{
				Interval:  "month",
				UsageType: "licensed",
			},
		}),
		verifyPrice(priceExpectations{
			Address:               "stripe_price.test",
			CompareStateAttrs:     []string{"currency", "product"},
			CompareMetadata:       true,
			CheckActive:           true,
			ExpectedActive:        true,
			CheckType:             true,
			ExpectedType:          "recurring",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "per_unit",
			CheckUnitAmount:       true,
			ExpectedUnitAmount:    1500,
			CheckRecurring:        true,
			ExpectedRecurring: priceRecurringExpectation{
				Interval:  "month",
				UsageType: "licensed",
			},
		}),
		verifyPriceDestroyInactive,
	)
}

func TestAccManagedPriceDecimalAmounts(t *testing.T) {
	runBaseManagedCaseWithImportIgnore(
		t,
		"price_decimal_amounts",
		"stripe_price",
		"stripe_price.test",
		"price_decimal_amounts_create.tf",
		"price_decimal_amounts_update.tf",
		true,
		[]string{"unit_amount_decimal", "currency_options"},
		verifyPrice(priceExpectations{
			Address:               "stripe_price.test",
			CompareStateAttrs:     []string{"currency", "product"},
			CompareMetadata:       true,
			CheckActive:           true,
			ExpectedActive:        true,
			CheckType:             true,
			ExpectedType:          "recurring",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "per_unit",
			CheckRecurring:        true,
			ExpectedRecurring: priceRecurringExpectation{
				Interval:  "month",
				UsageType: "licensed",
			},
			DecimalStrings: []decimalStringExpectation{
				{Attribute: "unit_amount_decimal", Expected: "1500.0"},
			},
		}),
		verifyPrice(priceExpectations{
			Address:               "stripe_price.test",
			CompareStateAttrs:     []string{"currency", "product"},
			CompareMetadata:       true,
			CheckActive:           true,
			ExpectedActive:        true,
			CheckType:             true,
			ExpectedType:          "recurring",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "per_unit",
			CheckRecurring:        true,
			ExpectedRecurring: priceRecurringExpectation{
				Interval:  "month",
				UsageType: "licensed",
			},
			StateStrings: []stateStringExpectation{
				{Attribute: "unit_amount_decimal", Expected: "1700.00"},
				{Attribute: "currency_options.0.unit_amount_decimal", Expected: "1400.2500"},
			},
			DecimalStrings: []decimalStringExpectation{
				{Attribute: "unit_amount_decimal", Expected: "1700.00"},
			},
		}),
		verifyPriceDestroyInactive,
	)
}

func TestAccManagedPriceDecimalAmountsLegacyUpgrade(t *testing.T) {
	runBaseManagedLegacyUpgradeCase(
		t,
		"price_decimal_amounts_legacy_upgrade",
		"stripe_price",
		"stripe_price.test",
		"price_decimal_amounts_create.tf",
		verifyPrice(priceExpectations{
			Address:               "stripe_price.test",
			CompareStateAttrs:     []string{"currency", "product"},
			CompareMetadata:       true,
			CheckActive:           true,
			ExpectedActive:        true,
			CheckType:             true,
			ExpectedType:          "recurring",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "per_unit",
			CheckRecurring:        true,
			ExpectedRecurring: priceRecurringExpectation{
				Interval:  "month",
				UsageType: "licensed",
			},
			DecimalStrings: []decimalStringExpectation{
				{Attribute: "unit_amount_decimal", Expected: "1500.0"},
			},
		}),
		verifyPriceDestroyInactive,
	)
}

func TestAccManagedPriceDecimalAmountsSubcentPrecision(t *testing.T) {
	runBaseManagedCaseWithImportIgnore(
		t,
		"price_decimal_amounts_subcent_precision",
		"stripe_price",
		"stripe_price.test",
		"price_decimal_amounts_subcent_precision_create.tf",
		"price_decimal_amounts_subcent_precision_update.tf",
		true,
		[]string{"unit_amount_decimal"},
		verifyPrice(priceExpectations{
			Address:               "stripe_price.test",
			CompareStateAttrs:     []string{"currency", "product"},
			CompareMetadata:       true,
			CheckActive:           true,
			ExpectedActive:        true,
			CheckType:             true,
			ExpectedType:          "recurring",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "per_unit",
			CheckRecurring:        true,
			ExpectedRecurring: priceRecurringExpectation{
				Interval:  "month",
				UsageType: "licensed",
			},
			StateStrings: []stateStringExpectation{
				{Attribute: "unit_amount_decimal", Expected: "0.0015000000"},
			},
			DecimalStrings: []decimalStringExpectation{
				{Attribute: "unit_amount_decimal", Expected: "0.0015000000"},
			},
		}),
		verifyPrice(priceExpectations{
			Address:               "stripe_price.test",
			CompareStateAttrs:     []string{"currency", "product"},
			CompareMetadata:       true,
			CheckActive:           true,
			ExpectedActive:        true,
			CheckType:             true,
			ExpectedType:          "recurring",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "per_unit",
			CheckRecurring:        true,
			ExpectedRecurring: priceRecurringExpectation{
				Interval:  "month",
				UsageType: "licensed",
			},
			StateStrings: []stateStringExpectation{
				{Attribute: "unit_amount_decimal", Expected: "0.001250000000"},
			},
			DecimalStrings: []decimalStringExpectation{
				{Attribute: "unit_amount_decimal", Expected: "0.001250000000"},
			},
		}),
		verifyPriceDestroyInactive,
	)
}

func TestAccManagedPriceActiveToggle(t *testing.T) {
	runBaseManagedCase(
		t,
		"price_active_toggle",
		"stripe_price",
		"stripe_price.test",
		"price_active_create.tf",
		"price_active_update.tf",
		true,
		verifyPrice(priceExpectations{
			Address:               "stripe_price.test",
			CompareStateAttrs:     []string{"currency", "product"},
			CompareMetadata:       true,
			CheckActive:           true,
			ExpectedActive:        true,
			CheckType:             true,
			ExpectedType:          "recurring",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "per_unit",
			CheckUnitAmount:       true,
			ExpectedUnitAmount:    2400,
			CheckRecurring:        true,
			ExpectedRecurring: priceRecurringExpectation{
				Interval:  "month",
				UsageType: "licensed",
			},
		}),
		verifyPrice(priceExpectations{
			Address:               "stripe_price.test",
			CompareStateAttrs:     []string{"currency", "product"},
			CompareMetadata:       true,
			CheckActive:           true,
			ExpectedActive:        false,
			CheckType:             true,
			ExpectedType:          "recurring",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "per_unit",
			CheckUnitAmount:       true,
			ExpectedUnitAmount:    2400,
			CheckRecurring:        true,
			ExpectedRecurring: priceRecurringExpectation{
				Interval:  "month",
				UsageType: "licensed",
			},
		}),
		verifyPriceDestroyInactive,
	)
}

func TestAccManagedPriceLookupKey(t *testing.T) {
	runBaseManagedCase(
		t,
		"price_lookup_key",
		"stripe_price",
		"stripe_price.test",
		"price_lookup_key_create.tf",
		"price_lookup_key_update.tf",
		true,
		verifyPrice(priceExpectations{
			Address: "stripe_price.test",
			CompareStateAttrs: []string{
				"currency",
				"product",
				"lookup_key",
				"nickname",
				"tax_behavior",
			},
			CompareMetadata: true,
			MissingStateAttrs: []string{
				"transfer_lookup_key",
			},
			CheckActive:           true,
			ExpectedActive:        true,
			CheckType:             true,
			ExpectedType:          "recurring",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "per_unit",
			CheckUnitAmount:       true,
			ExpectedUnitAmount:    3600,
			CheckRecurring:        true,
			ExpectedRecurring: priceRecurringExpectation{
				Interval:  "year",
				UsageType: "licensed",
			},
		}),
		verifyPrice(priceExpectations{
			Address: "stripe_price.test",
			CompareStateAttrs: []string{
				"currency",
				"product",
				"lookup_key",
				"nickname",
				"tax_behavior",
			},
			CompareMetadata: true,
			MissingStateAttrs: []string{
				"transfer_lookup_key",
			},
			CheckActive:           true,
			ExpectedActive:        true,
			CheckType:             true,
			ExpectedType:          "recurring",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "per_unit",
			CheckUnitAmount:       true,
			ExpectedUnitAmount:    3700,
			CheckRecurring:        true,
			ExpectedRecurring: priceRecurringExpectation{
				Interval:  "year",
				UsageType: "licensed",
			},
		}),
		verifyPriceDestroyInactive,
	)
}

func TestAccManagedPriceTransformQuantity(t *testing.T) {
	runBaseManagedCase(
		t,
		"price_transform_quantity",
		"stripe_price",
		"stripe_price.test",
		"price_transform_quantity_create.tf",
		"price_transform_quantity_update.tf",
		true,
		verifyPrice(priceExpectations{
			Address:               "stripe_price.test",
			CompareStateAttrs:     []string{"currency", "product"},
			CompareMetadata:       true,
			CheckActive:           true,
			ExpectedActive:        true,
			CheckType:             true,
			ExpectedType:          "recurring",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "per_unit",
			CheckUnitAmount:       true,
			ExpectedUnitAmount:    99,
			CheckRecurring:        true,
			ExpectedRecurring: priceRecurringExpectation{
				Interval:  "month",
				UsageType: "licensed",
			},
			CheckTransformQuantity:    true,
			ExpectedTransformDivideBy: 100,
			ExpectedTransformRound:    "up",
		}),
		verifyPrice(priceExpectations{
			Address:               "stripe_price.test",
			CompareStateAttrs:     []string{"currency", "product"},
			CompareMetadata:       true,
			CheckActive:           true,
			ExpectedActive:        true,
			CheckType:             true,
			ExpectedType:          "recurring",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "per_unit",
			CheckUnitAmount:       true,
			ExpectedUnitAmount:    99,
			CheckRecurring:        true,
			ExpectedRecurring: priceRecurringExpectation{
				Interval:  "month",
				UsageType: "licensed",
			},
			CheckTransformQuantity:    true,
			ExpectedTransformDivideBy: 100,
			ExpectedTransformRound:    "up",
		}),
		verifyPriceDestroyInactive,
	)
}

func TestAccManagedPriceTiered(t *testing.T) {
	runBaseManagedCase(
		t,
		"price_tiered",
		"stripe_price",
		"stripe_price.test",
		"price_tiered_create.tf",
		"",
		true,
		verifyPrice(priceExpectations{
			Address:               "stripe_price.test",
			CompareStateAttrs:     []string{"currency", "product"},
			CompareMetadata:       true,
			CheckActive:           true,
			ExpectedActive:        true,
			CheckType:             true,
			ExpectedType:          "recurring",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "tiered",
			CheckRecurring:        true,
			ExpectedRecurring: priceRecurringExpectation{
				Interval:  "month",
				UsageType: "licensed",
			},
			CheckTiersMode:    true,
			ExpectedTiersMode: "graduated",
			CheckTiers:        true,
			ExpectedTiers: []priceTierExpectation{
				{UnitAmount: 900, UpTo: 10},
				{UnitAmount: 700, UpTo: 0},
			},
		}),
		nil,
		verifyPriceDestroyInactive,
	)
}

func TestAccManagedPriceTieredLegacyUpgrade(t *testing.T) {
	runBaseManagedLegacyUpgradeCase(
		t,
		"price_tiered_legacy_upgrade",
		"stripe_price",
		"stripe_price.test",
		"price_tiered_create.tf",
		verifyPrice(priceExpectations{
			Address:               "stripe_price.test",
			CompareStateAttrs:     []string{"currency", "product"},
			CompareMetadata:       true,
			CheckActive:           true,
			ExpectedActive:        true,
			CheckType:             true,
			ExpectedType:          "recurring",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "tiered",
			CheckRecurring:        true,
			ExpectedRecurring: priceRecurringExpectation{
				Interval:  "month",
				UsageType: "licensed",
			},
			CheckTiersMode:    true,
			ExpectedTiersMode: "graduated",
			CheckTiers:        true,
			ExpectedTiers: []priceTierExpectation{
				{UnitAmount: 900, UpTo: 10},
				{UnitAmount: 700, UpTo: 0},
			},
		}),
		verifyPriceDestroyInactive,
	)
}

func TestAccManagedPriceDecimalTiers(t *testing.T) {
	runBaseManagedCaseWithImportIgnore(
		t,
		"price_decimal_tiers",
		"stripe_price",
		"stripe_price.test",
		"price_decimal_tiers_create.tf",
		"price_decimal_tiers_update.tf",
		true,
		[]string{"tiers", "currency_options"},
		verifyPrice(priceExpectations{
			Address:               "stripe_price.test",
			CompareStateAttrs:     []string{"currency", "product"},
			CompareMetadata:       true,
			CheckActive:           true,
			ExpectedActive:        true,
			CheckType:             true,
			ExpectedType:          "recurring",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "tiered",
			CheckRecurring:        true,
			ExpectedRecurring: priceRecurringExpectation{
				Interval:  "month",
				UsageType: "licensed",
			},
			CheckTiersMode:    true,
			ExpectedTiersMode: "graduated",
			StateStrings: []stateStringExpectation{
				{Attribute: "tiers.0.unit_amount_decimal", Expected: "900.0"},
				{Attribute: "tiers.1.flat_amount_decimal", Expected: "150.0000"},
				{Attribute: "currency_options.0.tiers.0.unit_amount_decimal", Expected: "800.0"},
				{Attribute: "currency_options.0.tiers.1.flat_amount_decimal", Expected: "125.0000"},
			},
			MissingStateAttrs: []string{
				"tiers.0.unit_amount",
				"tiers.1.flat_amount",
				"currency_options.0.tiers.0.unit_amount",
				"currency_options.0.tiers.1.flat_amount",
			},
			DecimalStrings: []decimalStringExpectation{
				{Attribute: "tiers.0.unit_amount_decimal", Expected: "900.0"},
				{Attribute: "tiers.1.flat_amount_decimal", Expected: "150.0000"},
			},
		}),
		verifyPrice(priceExpectations{
			Address:               "stripe_price.test",
			CompareStateAttrs:     []string{"currency", "product"},
			CompareMetadata:       true,
			CheckActive:           true,
			ExpectedActive:        true,
			CheckType:             true,
			ExpectedType:          "recurring",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "tiered",
			CheckRecurring:        true,
			ExpectedRecurring: priceRecurringExpectation{
				Interval:  "month",
				UsageType: "licensed",
			},
			CheckTiersMode:    true,
			ExpectedTiersMode: "graduated",
			StateStrings: []stateStringExpectation{
				{Attribute: "tiers.0.unit_amount_decimal", Expected: "950.00"},
				{Attribute: "tiers.1.flat_amount_decimal", Expected: "175.0000"},
				{Attribute: "currency_options.0.tiers.0.unit_amount_decimal", Expected: "850.00"},
				{Attribute: "currency_options.0.tiers.1.flat_amount_decimal", Expected: "140.0000"},
			},
			MissingStateAttrs: []string{
				"tiers.0.unit_amount",
				"tiers.1.flat_amount",
				"currency_options.0.tiers.0.unit_amount",
				"currency_options.0.tiers.1.flat_amount",
			},
			DecimalStrings: []decimalStringExpectation{
				{Attribute: "tiers.0.unit_amount_decimal", Expected: "950.00"},
				{Attribute: "tiers.1.flat_amount_decimal", Expected: "175.0000"},
			},
		}),
		verifyPriceDestroyInactive,
	)
}

func TestAccManagedPriceOneTime(t *testing.T) {
	runBaseManagedCase(
		t,
		"price_one_time",
		"stripe_price",
		"stripe_price.test",
		"price_one_time_create.tf",
		"price_one_time_update.tf",
		true,
		verifyPrice(priceExpectations{
			Address:               "stripe_price.test",
			CompareStateAttrs:     []string{"currency", "product", "lookup_key", "nickname"},
			CompareMetadata:       true,
			CheckActive:           true,
			ExpectedActive:        true,
			CheckType:             true,
			ExpectedType:          "one_time",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "per_unit",
			CheckUnitAmount:       true,
			ExpectedUnitAmount:    4800,
		}),
		verifyPrice(priceExpectations{
			Address:               "stripe_price.test",
			CompareStateAttrs:     []string{"currency", "product", "lookup_key", "nickname"},
			CompareMetadata:       true,
			CheckActive:           true,
			ExpectedActive:        true,
			CheckType:             true,
			ExpectedType:          "one_time",
			CheckBillingScheme:    true,
			ExpectedBillingScheme: "per_unit",
			CheckUnitAmount:       true,
			ExpectedUnitAmount:    4800,
		}),
		verifyPriceDestroyInactive,
	)
}

func TestAccManagedWebhookEndpointConnect(t *testing.T) {
	runBaseManagedCase(
		t,
		"webhook_endpoint_connect",
		"stripe_webhook_endpoint",
		"stripe_webhook_endpoint.test",
		"webhook_endpoint_connect_create.tf",
		"webhook_endpoint_connect_update.tf",
		false,
		verifyWebhookEndpoint(webhookEndpointExpectations{
			Address:            "stripe_webhook_endpoint.test",
			CompareStateAttrs:  []string{"url", "description"},
			CompareMetadata:    true,
			CheckEnabledEvents: true,
			ExpectedEnabledEvents: []string{
				"checkout.session.completed",
				"price.created",
			},
		}),
		verifyWebhookEndpoint(webhookEndpointExpectations{
			Address:            "stripe_webhook_endpoint.test",
			CompareStateAttrs:  []string{"url", "description"},
			CompareMetadata:    true,
			CheckEnabledEvents: true,
			ExpectedEnabledEvents: []string{
				"checkout.session.completed",
				"price.created",
				"product.updated",
			},
		}),
		verifyWebhookEndpointDestroyMissing,
	)
}

func TestAccManagedWebhookEndpointURLUpdate(t *testing.T) {
	runBaseManagedCase(
		t,
		"webhook_endpoint_url_update",
		"stripe_webhook_endpoint",
		"stripe_webhook_endpoint.test",
		"webhook_endpoint_url_create.tf",
		"webhook_endpoint_url_update.tf",
		false,
		verifyWebhookEndpoint(webhookEndpointExpectations{
			Address:            "stripe_webhook_endpoint.test",
			CompareStateAttrs:  []string{"url", "description"},
			CompareMetadata:    true,
			CheckEnabledEvents: true,
			ExpectedEnabledEvents: []string{
				"customer.created",
			},
		}),
		verifyWebhookEndpoint(webhookEndpointExpectations{
			Address:            "stripe_webhook_endpoint.test",
			CompareStateAttrs:  []string{"url", "description"},
			CompareMetadata:    true,
			CheckEnabledEvents: true,
			ExpectedEnabledEvents: []string{
				"customer.created",
				"customer.updated",
			},
		}),
		verifyWebhookEndpointDestroyMissing,
	)
}

func TestAccManagedWebhookEndpointURLLegacyUpgrade(t *testing.T) {
	runBaseManagedLegacyUpgradeCase(
		t,
		"webhook_endpoint_url_legacy_upgrade",
		"stripe_webhook_endpoint",
		"stripe_webhook_endpoint.test",
		"webhook_endpoint_url_create.tf",
		verifyWebhookEndpoint(webhookEndpointExpectations{
			Address:            "stripe_webhook_endpoint.test",
			CompareStateAttrs:  []string{"url", "description"},
			CompareMetadata:    true,
			CheckEnabledEvents: true,
			ExpectedEnabledEvents: []string{
				"customer.created",
			},
		}),
		verifyWebhookEndpointDestroyMissing,
	)
}

func TestAccManagedWebhookEndpointAPIVersion(t *testing.T) {
	runBaseManagedCase(
		t,
		"webhook_endpoint_api_version",
		"stripe_webhook_endpoint",
		"stripe_webhook_endpoint.test",
		"webhook_endpoint_api_version_create.tf",
		"webhook_endpoint_api_version_update.tf",
		false,
		verifyWebhookEndpoint(webhookEndpointExpectations{
			Address:            "stripe_webhook_endpoint.test",
			CompareStateAttrs:  []string{"api_version", "url", "description"},
			CompareMetadata:    true,
			CheckEnabledEvents: true,
			ExpectedEnabledEvents: []string{
				"customer.created",
				"customer.deleted",
			},
		}),
		verifyWebhookEndpoint(webhookEndpointExpectations{
			Address:            "stripe_webhook_endpoint.test",
			CompareStateAttrs:  []string{"api_version", "url", "description"},
			CompareMetadata:    true,
			CheckEnabledEvents: true,
			ExpectedEnabledEvents: []string{
				"customer.created",
				"customer.deleted",
				"customer.updated",
			},
		}),
		verifyWebhookEndpointDestroyMissing,
	)
}

func TestAccManagedV2CoreEventDestinationBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"v2_core_event_destination_basic",
		"stripe_v2_core_event_destination",
		"stripe_v2_core_event_destination.eventbridge",
		"v2_core_event_destination_create.tf",
		"v2_core_event_destination_update.tf",
		false,
		func(env runner.TestEnv, client *stripe.Client, state *terraform.State) error {
			if err := verifyEventDestination(eventDestinationExpectations{
				Address: "stripe_v2_core_event_destination.eventbridge",
				CompareStateAttrs: []string{
					"name",
					"description",
					"type",
					"event_payload",
				},
				CompareMetadata:    true,
				CheckEnabledEvents: true,
				ExpectedEnabledEvents: []string{
					"v1.billing.meter.error_report_triggered",
				},
				CheckAmazonEventbridge: true,
				ExpectedAWSAccountID:   "111122223333",
			})(env, client, state); err != nil {
				return err
			}
			return verifyEventDestination(eventDestinationExpectations{
				Address: "stripe_v2_core_event_destination.webhook",
				CompareStateAttrs: []string{
					"name",
					"description",
					"type",
					"event_payload",
				},
				CompareMetadata:    true,
				CheckEnabledEvents: true,
				ExpectedEnabledEvents: []string{
					"v1.billing.meter.error_report_triggered",
				},
				CheckWebhookEndpointURL: true,
				ExpectedWebhookURLParts: []string{
					"https://example.com/sdk-codegen/event-destination/",
				},
			})(env, client, state)
		},
		func(env runner.TestEnv, client *stripe.Client, state *terraform.State) error {
			if err := verifyEventDestination(eventDestinationExpectations{
				Address: "stripe_v2_core_event_destination.eventbridge",
				CompareStateAttrs: []string{
					"name",
					"description",
					"type",
					"event_payload",
				},
				CompareMetadata:    true,
				CheckEnabledEvents: true,
				ExpectedEnabledEvents: []string{
					"v1.billing.meter.error_report_triggered",
				},
				CheckAmazonEventbridge: true,
				ExpectedAWSAccountID:   "111122223333",
			})(env, client, state); err != nil {
				return err
			}
			return verifyEventDestination(eventDestinationExpectations{
				Address: "stripe_v2_core_event_destination.webhook",
				CompareStateAttrs: []string{
					"name",
					"description",
					"type",
					"event_payload",
				},
				CompareMetadata:    true,
				CheckEnabledEvents: true,
				ExpectedEnabledEvents: []string{
					"v1.billing.meter.error_report_triggered",
				},
				CheckWebhookEndpointURL: true,
				ExpectedWebhookURLParts: []string{
					"https://example.com/sdk-codegen/event-destination/",
					"/updated",
				},
			})(env, client, state)
		},
		verifyEventDestinationDestroyMissing,
	)
}

func TestAccManagedV2CoreEventDestinationSnapshot(t *testing.T) {
	runBaseManagedCase(
		t,
		"v2_core_event_destination_snapshot",
		"stripe_v2_core_event_destination",
		"stripe_v2_core_event_destination.eventbridge",
		"v2_core_event_destination_snapshot_create.tf",
		"v2_core_event_destination_snapshot_update.tf",
		false,
		func(env runner.TestEnv, client *stripe.Client, state *terraform.State) error {
			if err := verifyEventDestination(eventDestinationExpectations{
				Address: "stripe_v2_core_event_destination.eventbridge",
				CompareStateAttrs: []string{
					"name",
					"description",
					"type",
					"event_payload",
					"snapshot_api_version",
				},
				CompareMetadata:    true,
				CheckEnabledEvents: true,
				ExpectedEnabledEvents: []string{
					"customer.created",
				},
				CheckAmazonEventbridge: true,
				ExpectedAWSAccountID:   "111122223333",
			})(env, client, state); err != nil {
				return err
			}
			return verifyEventDestination(eventDestinationExpectations{
				Address: "stripe_v2_core_event_destination.webhook",
				CompareStateAttrs: []string{
					"name",
					"description",
					"type",
					"event_payload",
					"snapshot_api_version",
				},
				CompareMetadata:    true,
				CheckEnabledEvents: true,
				ExpectedEnabledEvents: []string{
					"customer.created",
				},
				CheckWebhookEndpointURL: true,
				ExpectedWebhookURLParts: []string{
					"https://example.com/sdk-codegen/event-destination-snapshot/",
				},
			})(env, client, state)
		},
		func(env runner.TestEnv, client *stripe.Client, state *terraform.State) error {
			if err := verifyEventDestination(eventDestinationExpectations{
				Address: "stripe_v2_core_event_destination.eventbridge",
				CompareStateAttrs: []string{
					"name",
					"description",
					"type",
					"event_payload",
					"snapshot_api_version",
				},
				CompareMetadata:    true,
				CheckEnabledEvents: true,
				ExpectedEnabledEvents: []string{
					"customer.created",
					"customer.updated",
				},
				CheckAmazonEventbridge: true,
				ExpectedAWSAccountID:   "111122223333",
			})(env, client, state); err != nil {
				return err
			}
			return verifyEventDestination(eventDestinationExpectations{
				Address: "stripe_v2_core_event_destination.webhook",
				CompareStateAttrs: []string{
					"name",
					"description",
					"type",
					"event_payload",
					"snapshot_api_version",
				},
				CompareMetadata:    true,
				CheckEnabledEvents: true,
				ExpectedEnabledEvents: []string{
					"customer.created",
					"customer.updated",
				},
				CheckWebhookEndpointURL: true,
				ExpectedWebhookURLParts: []string{
					"https://example.com/sdk-codegen/event-destination-snapshot/",
					"/updated",
				},
			})(env, client, state)
		},
		verifyEventDestinationDestroyMissing,
	)
}

func TestAccManagedV2CoreEventDestinationLegacyUpgrade(t *testing.T) {
	// Legacy provider v0.2.2 does not accept the current schema shape for this case.
	// Use a legacy-only create fixture so the upgrade path exercises state migration.
	runBaseManagedLegacyUpgradeCase(
		t,
		"v2_core_event_destination_legacy_upgrade",
		"stripe_v2_core_event_destination",
		"stripe_v2_core_event_destination.eventbridge",
		"v2_core_event_destination_legacy_upgrade_create.tf",
		func(env runner.TestEnv, client *stripe.Client, state *terraform.State) error {
			if err := verifyEventDestination(eventDestinationExpectations{
				Address: "stripe_v2_core_event_destination.eventbridge",
				CompareStateAttrs: []string{
					"name",
					"description",
					"type",
					"event_payload",
				},
				CompareMetadata:    true,
				CheckEnabledEvents: true,
				ExpectedEnabledEvents: []string{
					"v1.billing.meter.error_report_triggered",
				},
				CheckAmazonEventbridge: true,
				ExpectedAWSAccountID:   "111122223333",
			})(env, client, state); err != nil {
				return err
			}
			return verifyEventDestination(eventDestinationExpectations{
				Address: "stripe_v2_core_event_destination.webhook",
				CompareStateAttrs: []string{
					"name",
					"description",
					"type",
					"event_payload",
				},
				CompareMetadata:    true,
				CheckEnabledEvents: true,
				ExpectedEnabledEvents: []string{
					"v1.billing.meter.error_report_triggered",
				},
				CheckWebhookEndpointURL: true,
				ExpectedWebhookURLParts: []string{
					"https://example.com/sdk-codegen/event-destination/",
				},
			})(env, client, state)
		},
		verifyEventDestinationDestroyMissing,
	)
}

func TestAccManagedBillingMeterBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"billing_meter_basic",
		"stripe_billing_meter",
		"stripe_billing_meter.test",
		"billing_meter_create.tf",
		"billing_meter_update.tf",
		true,
		verifyBillingMeter(billingMeterExpectations{
			Address: "stripe_billing_meter.test",
			CompareStateAttrs: []string{
				"display_name",
				"event_name",
				"customer_mapping.event_payload_key",
				"customer_mapping.type",
				"default_aggregation.formula",
				"value_settings.event_payload_key",
			},
			CheckStatus:    true,
			ExpectedStatus: "active",
		}),
		verifyBillingMeter(billingMeterExpectations{
			Address: "stripe_billing_meter.test",
			CompareStateAttrs: []string{
				"display_name",
				"event_name",
				"customer_mapping.event_payload_key",
				"customer_mapping.type",
				"default_aggregation.formula",
				"value_settings.event_payload_key",
			},
			CheckStatus:    true,
			ExpectedStatus: "active",
		}),
		verifyBillingMeterDestroyInactive,
	)
}

func TestAccManagedBillingMeterLegacyUpgrade(t *testing.T) {
	runBaseManagedLegacyUpgradeCase(
		t,
		"billing_meter_legacy_upgrade",
		"stripe_billing_meter",
		"stripe_billing_meter.test",
		"billing_meter_create.tf",
		verifyBillingMeter(billingMeterExpectations{
			Address: "stripe_billing_meter.test",
			CompareStateAttrs: []string{
				"display_name",
				"event_name",
				"customer_mapping.event_payload_key",
				"customer_mapping.type",
				"default_aggregation.formula",
				"value_settings.event_payload_key",
			},
			CheckStatus:    true,
			ExpectedStatus: "active",
		}),
		verifyBillingMeterDestroyInactive,
	)
}

func TestAccManagedBillingMeterDeactivate(t *testing.T) {
	runBaseManagedCase(
		t,
		"billing_meter_deactivate",
		"stripe_billing_meter",
		"stripe_billing_meter.test",
		"billing_meter_deactivate_create.tf",
		"",
		true,
		verifyBillingMeter(billingMeterExpectations{
			Address: "stripe_billing_meter.test",
			CompareStateAttrs: []string{
				"display_name",
				"event_name",
				"customer_mapping.event_payload_key",
				"customer_mapping.type",
				"default_aggregation.formula",
				"value_settings.event_payload_key",
			},
			CheckStatus:    true,
			ExpectedStatus: "active",
		}),
		nil,
		verifyBillingMeterDestroyInactive,
	)
}

func TestAccManagedBillingMeterCountAggregation(t *testing.T) {
	runBaseManagedCase(
		t,
		"billing_meter_count_aggregation",
		"stripe_billing_meter",
		"stripe_billing_meter.test",
		"billing_meter_count_aggregation_create.tf",
		"billing_meter_count_aggregation_update.tf",
		true,
		verifyBillingMeter(billingMeterExpectations{
			Address: "stripe_billing_meter.test",
			CompareStateAttrs: []string{
				"display_name",
				"event_name",
				"customer_mapping.event_payload_key",
				"customer_mapping.type",
				"default_aggregation.formula",
				"value_settings.event_payload_key",
			},
			CheckStatus:    true,
			ExpectedStatus: "active",
		}),
		verifyBillingMeter(billingMeterExpectations{
			Address: "stripe_billing_meter.test",
			CompareStateAttrs: []string{
				"display_name",
				"event_name",
				"customer_mapping.event_payload_key",
				"customer_mapping.type",
				"default_aggregation.formula",
				"value_settings.event_payload_key",
			},
			CheckStatus:    true,
			ExpectedStatus: "active",
		}),
		verifyBillingMeterDestroyInactive,
	)
}

func TestAccManagedEntitlementsFeatureBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"entitlements_feature_basic",
		"stripe_entitlements_feature",
		"stripe_entitlements_feature.test",
		"entitlements_feature_create.tf",
		"entitlements_feature_update.tf",
		true,
		verifyEntitlementsFeature(entitlementsFeatureExpectations{
			Address:           "stripe_entitlements_feature.test",
			CompareStateAttrs: []string{"lookup_key", "name"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    true,
		}),
		verifyEntitlementsFeature(entitlementsFeatureExpectations{
			Address:           "stripe_entitlements_feature.test",
			CompareStateAttrs: []string{"lookup_key", "name"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    true,
		}),
		verifyEntitlementsFeatureDestroyInactive,
	)
}

func TestAccManagedEntitlementsFeatureRename(t *testing.T) {
	runBaseManagedCase(
		t,
		"entitlements_feature_rename",
		"stripe_entitlements_feature",
		"stripe_entitlements_feature.test",
		"entitlements_feature_rename_create.tf",
		"entitlements_feature_rename_update.tf",
		true,
		verifyEntitlementsFeature(entitlementsFeatureExpectations{
			Address:           "stripe_entitlements_feature.test",
			CompareStateAttrs: []string{"lookup_key", "name"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    true,
		}),
		verifyEntitlementsFeature(entitlementsFeatureExpectations{
			Address:           "stripe_entitlements_feature.test",
			CompareStateAttrs: []string{"lookup_key", "name"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    true,
		}),
		verifyEntitlementsFeatureDestroyInactive,
	)
}

func TestAccManagedEntitlementsFeatureLegacyUpgrade(t *testing.T) {
	runBaseManagedLegacyUpgradeCase(
		t,
		"entitlements_feature_legacy_upgrade",
		"stripe_entitlements_feature",
		"stripe_entitlements_feature.test",
		"entitlements_feature_create.tf",
		verifyEntitlementsFeature(entitlementsFeatureExpectations{
			Address:           "stripe_entitlements_feature.test",
			CompareStateAttrs: []string{"lookup_key", "name"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    true,
		}),
		verifyEntitlementsFeatureDestroyInactive,
	)
}

func TestAccManagedEntitlementsFeatureMetadataUpdate(t *testing.T) {
	runBaseManagedCase(
		t,
		"entitlements_feature_metadata_update",
		"stripe_entitlements_feature",
		"stripe_entitlements_feature.test",
		"entitlements_feature_metadata_update_create.tf",
		"entitlements_feature_metadata_update_update.tf",
		true,
		verifyEntitlementsFeature(entitlementsFeatureExpectations{
			Address:           "stripe_entitlements_feature.test",
			CompareStateAttrs: []string{"lookup_key", "name"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    true,
		}),
		verifyEntitlementsFeature(entitlementsFeatureExpectations{
			Address:           "stripe_entitlements_feature.test",
			CompareStateAttrs: []string{"lookup_key", "name"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    true,
		}),
		verifyEntitlementsFeatureDestroyInactive,
	)
}

func TestAccManagedCouponPercentOff(t *testing.T) {
	runBaseManagedCase(
		t,
		"coupon_percent_off",
		"stripe_coupon",
		"stripe_coupon.test",
		"coupon_percent_create.tf",
		"coupon_percent_update.tf",
		true,
		verifyCoupon(couponExpectations{
			Address:                 "stripe_coupon.test",
			CompareStateStringAttrs: []string{"name", "duration"},
			CompareMetadata:         true,
			CheckPercentOff:         true,
			ExpectedPercentOff:      25.0,
		}),
		verifyCoupon(couponExpectations{
			Address:                 "stripe_coupon.test",
			CompareStateStringAttrs: []string{"name", "duration"},
			CompareMetadata:         true,
			CheckPercentOff:         true,
			ExpectedPercentOff:      25.0,
		}),
		verifyCouponDestroyMissing,
	)
}

func TestAccManagedCouponAmountOff(t *testing.T) {
	runBaseManagedCase(
		t,
		"coupon_amount_off",
		"stripe_coupon",
		"stripe_coupon.test",
		"coupon_amount_off_create.tf",
		"coupon_amount_off_update.tf",
		true,
		verifyCoupon(couponExpectations{
			Address:                 "stripe_coupon.test",
			CompareStateStringAttrs: []string{"name", "currency", "duration"},
			CompareMetadata:         true,
			CheckAmountOff:          true,
			ExpectedAmountOff:       500,
		}),
		verifyCoupon(couponExpectations{
			Address:                 "stripe_coupon.test",
			CompareStateStringAttrs: []string{"name", "currency", "duration"},
			CompareMetadata:         true,
			CheckAmountOff:          true,
			ExpectedAmountOff:       500,
		}),
		verifyCouponDestroyMissing,
	)
}

func TestAccManagedCouponRepeating(t *testing.T) {
	runBaseManagedCase(
		t,
		"coupon_repeating",
		"stripe_coupon",
		"stripe_coupon.test",
		"coupon_repeating_create.tf",
		"coupon_repeating_update.tf",
		true,
		verifyCoupon(couponExpectations{
			Address:                  "stripe_coupon.test",
			CompareStateStringAttrs:  []string{"name", "duration"},
			CompareMetadata:          true,
			CheckPercentOff:          true,
			ExpectedPercentOff:       15.0,
			CheckDurationInMonths:    true,
			ExpectedDurationInMonths: 3,
		}),
		verifyCoupon(couponExpectations{
			Address:                  "stripe_coupon.test",
			CompareStateStringAttrs:  []string{"name", "duration"},
			CompareMetadata:          true,
			CheckPercentOff:          true,
			ExpectedPercentOff:       15.0,
			CheckDurationInMonths:    true,
			ExpectedDurationInMonths: 3,
		}),
		verifyCouponDestroyMissing,
	)
}

func TestAccManagedCouponAppliesTo(t *testing.T) {
	runBaseManagedCaseWithImportIgnore(
		t,
		"coupon_applies_to",
		"stripe_coupon",
		"stripe_coupon.test",
		"coupon_applies_to_create.tf",
		"coupon_applies_to_update.tf",
		true,
		[]string{
			"applies_to",
			"applies_to.%", // sdkv2 import path retained in failure diff output
			"applies_to.products",
			"applies_to.products.#",
			"applies_to.products.0",
		},
		verifyCoupon(couponExpectations{
			Address:                 "stripe_coupon.test",
			CompareStateStringAttrs: []string{"name", "duration"},
			CompareMetadata:         true,
			CheckPercentOff:         true,
			ExpectedPercentOff:      10.0,
		}),
		verifyCoupon(couponExpectations{
			Address:                 "stripe_coupon.test",
			CompareStateStringAttrs: []string{"name", "duration"},
			CompareMetadata:         true,
			CheckPercentOff:         true,
			ExpectedPercentOff:      10.0,
		}),
		verifyCouponDestroyMissing,
	)
}

func TestAccManagedCouponAppliesToProductUpdate(t *testing.T) {
	runBaseManagedCaseWithImportIgnore(
		t,
		"coupon_applies_to_product_update",
		"stripe_coupon",
		"stripe_coupon.test",
		"coupon_applies_to_product_update_create.tf",
		"coupon_applies_to_product_update_update.tf",
		true,
		[]string{
			"applies_to",
			"applies_to.%",
			"applies_to.products",
			"applies_to.products.#",
			"applies_to.products.0",
		},
		verifyCoupon(couponExpectations{
			Address:                 "stripe_coupon.test",
			CompareStateStringAttrs: []string{"name", "duration"},
			CompareMetadata:         true,
			CheckPercentOff:         true,
			ExpectedPercentOff:      10.0,
			CheckAppliesToProduct:   true,
		}),
		verifyCoupon(couponExpectations{
			Address:                 "stripe_coupon.test",
			CompareStateStringAttrs: []string{"name", "duration"},
			CompareMetadata:         true,
			CheckPercentOff:         true,
			ExpectedPercentOff:      10.0,
			CheckAppliesToProduct:   true,
		}),
		verifyCouponDestroyMissing,
	)
}

func TestAccManagedCouponAppliesToLegacyUpgrade(t *testing.T) {
	runBaseManagedLegacyUpgradeCase(
		t,
		"coupon_applies_to_legacy_upgrade",
		"stripe_coupon",
		"stripe_coupon.test",
		"coupon_applies_to_create.tf",
		verifyCoupon(couponExpectations{
			Address:                 "stripe_coupon.test",
			CompareStateStringAttrs: []string{"name", "duration"},
			CompareMetadata:         true,
			CheckPercentOff:         true,
			ExpectedPercentOff:      10.0,
		}),
		verifyCouponDestroyMissing,
	)
}

func TestAccManagedCustomerBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"customer_basic",
		"stripe_customer",
		"stripe_customer.test",
		"customer_create.tf",
		"customer_update.tf",
		true,
		verifyCustomer(customerExpectations{
			Address:           "stripe_customer.test",
			CompareStateAttrs: []string{"name", "email", "description"},
			CompareMetadata:   true,
		}),
		verifyCustomer(customerExpectations{
			Address:           "stripe_customer.test",
			CompareStateAttrs: []string{"name", "email", "description"},
			CompareMetadata:   true,
		}),
		verifyCustomerDestroyDeleted,
	)
}

func TestAccManagedCustomerProfile(t *testing.T) {
	runBaseManagedCase(
		t,
		"customer_profile",
		"stripe_customer",
		"stripe_customer.test",
		"customer_profile_create.tf",
		"customer_profile_update.tf",
		true,
		verifyCustomer(customerExpectations{
			Address:           "stripe_customer.test",
			CompareStateAttrs: []string{"name", "email", "description", "phone"},
			CompareMetadata:   true,
		}),
		verifyCustomer(customerExpectations{
			Address:           "stripe_customer.test",
			CompareStateAttrs: []string{"name", "email", "description", "phone"},
			CompareMetadata:   true,
		}),
		verifyCustomerDestroyDeleted,
	)
}

func TestAccManagedCustomerInvoiceTax(t *testing.T) {
	runBaseManagedCase(
		t,
		"customer_invoice_tax",
		"stripe_customer",
		"stripe_customer.test",
		"customer_invoice_tax_create.tf",
		"customer_invoice_tax_update.tf",
		true,
		verifyCustomer(customerExpectations{
			Address: "stripe_customer.test",
			CompareStateAttrs: []string{
				"name",
				"email",
				"invoice_prefix",
				"tax_exempt",
			},
			CompareMetadata: true,
		}),
		verifyCustomer(customerExpectations{
			Address: "stripe_customer.test",
			CompareStateAttrs: []string{
				"name",
				"email",
				"invoice_prefix",
				"tax_exempt",
			},
			CompareMetadata: true,
		}),
		verifyCustomerDestroyDeleted,
	)
}

func TestAccManagedCustomerLegacyUpgrade(t *testing.T) {
	runBaseManagedLegacyUpgradeCase(
		t,
		"customer_legacy_upgrade",
		"stripe_customer",
		"stripe_customer.test",
		"customer_create.tf",
		verifyCustomer(customerExpectations{
			Address:           "stripe_customer.test",
			CompareStateAttrs: []string{"name", "email", "description"},
			CompareMetadata:   true,
		}),
		verifyCustomerDestroyDeleted,
	)
}

func TestAccManagedCustomerInvoiceTaxLegacyUpgrade(t *testing.T) {
	runBaseManagedLegacyUpgradeCase(
		t,
		"customer_invoice_tax_legacy_upgrade",
		"stripe_customer",
		"stripe_customer.test",
		"customer_invoice_tax_create.tf",
		verifyCustomer(customerExpectations{
			Address: "stripe_customer.test",
			CompareStateAttrs: []string{
				"name",
				"email",
				"invoice_prefix",
				"tax_exempt",
			},
			CompareMetadata: true,
		}),
		verifyCustomerDestroyDeleted,
	)
}

func TestAccManagedPromotionCodeCustomer(t *testing.T) {
	runBaseManagedCase(
		t,
		"promotion_code_customer",
		"stripe_promotion_code",
		"stripe_promotion_code.test",
		"promotion_code_customer_create.tf",
		"promotion_code_customer_update.tf",
		true,
		verifyPromotionCode(promotionCodeExpectations{
			Address:           "stripe_promotion_code.test",
			CompareStateAttrs: []string{"code", "customer", "promotion.coupon"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    true,
		}),
		verifyPromotionCode(promotionCodeExpectations{
			Address:           "stripe_promotion_code.test",
			CompareStateAttrs: []string{"code", "customer", "promotion.coupon"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    false,
		}),
		verifyPromotionCodeDestroyInactive,
	)
}

func TestAccManagedPromotionCodeLegacyUpgrade(t *testing.T) {
	runBaseManagedLegacyUpgradeCase(
		t,
		"promotion_code_legacy_upgrade",
		"stripe_promotion_code",
		"stripe_promotion_code.test",
		"promotion_code_restrictions_create.tf",
		verifyPromotionCode(promotionCodeExpectations{
			Address:           "stripe_promotion_code.test",
			CompareStateAttrs: []string{"code", "promotion.coupon"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    true,
			CheckRestrictions: true,
			ExpectedRestrictions: promotionCodeRestrictionsExpectation{
				MinimumAmount:         2000,
				MinimumAmountCurrency: "usd",
				FirstTimeTransaction:  true,
			},
		}),
		verifyPromotionCodeDestroyInactive,
	)
}

func TestAccManagedPromotionCodeBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"promotion_code_basic",
		"stripe_promotion_code",
		"stripe_promotion_code.test",
		"promotion_code_basic_create.tf",
		"promotion_code_basic_update.tf",
		true,
		verifyPromotionCode(promotionCodeExpectations{
			Address:           "stripe_promotion_code.test",
			CompareStateAttrs: []string{"code", "promotion.coupon"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    true,
		}),
		verifyPromotionCode(promotionCodeExpectations{
			Address:           "stripe_promotion_code.test",
			CompareStateAttrs: []string{"code", "promotion.coupon"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    false,
		}),
		verifyPromotionCodeDestroyInactive,
	)
}

func TestAccManagedPromotionCodeRestrictions(t *testing.T) {
	runBaseManagedCase(
		t,
		"promotion_code_restrictions",
		"stripe_promotion_code",
		"stripe_promotion_code.test",
		"promotion_code_restrictions_create.tf",
		"",
		true,
		verifyPromotionCode(promotionCodeExpectations{
			Address:           "stripe_promotion_code.test",
			CompareStateAttrs: []string{"code", "promotion.coupon"},
			CompareMetadata:   true,
			CheckActive:       true,
			ExpectedActive:    true,
			CheckRestrictions: true,
			ExpectedRestrictions: promotionCodeRestrictionsExpectation{
				MinimumAmount:         2000,
				MinimumAmountCurrency: "usd",
				FirstTimeTransaction:  true,
			},
		}),
		nil,
		verifyPromotionCodeDestroyInactive,
	)
}
