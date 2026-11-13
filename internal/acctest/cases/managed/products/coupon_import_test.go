// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
package cases

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	stripe "github.com/stripe/stripe-go/v87"
	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

func TestAccManagedCouponAppliesToImport(t *testing.T) {
	if !runner.AcceptanceEnabled() {
		t.Skip("set TF_ACC=1 to run Terraform acceptance tests")
	}
	if !runner.ShouldRunCase(runner.CaseDefinition{
		Name: "coupon_applies_to_import", Surface: "stripe_coupon", Group: "base", Kind: "resource",
	}) {
		t.Skip("case filtered out by selectors")
	}
	env := runner.ResolveTestEnv(t, "base", nil)
	runner.ApplyResolvedAccountEnv(t, env)
	t.Setenv(resource.EnvTfAccProviderNamespace, "stripe")
	client := runner.NewStripeClient(env)
	product, err := client.V1Products.Create(context.Background(), &stripe.ProductCreateParams{
		Name: stripe.String("acctest-coupon-import"),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := client.V1Products.Delete(context.Background(), product.ID, nil); err != nil {
			t.Errorf("cleaning up import product: %v", err)
		}
	})
	coupon, err := client.V1Coupons.Create(context.Background(), &stripe.CouponCreateParams{
		Duration:   stripe.String("forever"),
		PercentOff: stripe.Float64(10),
		AppliesTo:  &stripe.CouponCreateAppliesToParams{Products: []*string{stripe.String(product.ID)}},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_, err := client.V1Coupons.Delete(context.Background(), coupon.ID, nil)
		if e, ok := err.(*stripe.Error); ok && e.HTTPStatusCode == 404 {
			return
		}
		if err != nil {
			t.Errorf("cleaning up import coupon: %v", err)
		}
	})
	config := fmt.Sprintf(`
terraform {
  required_providers {
    stripe = { source = "stripe/stripe" }
  }
}
provider "stripe" {}
resource "stripe_coupon" "test" {
  duration = "forever"
  percent_off = 10
  applies_to { products = [%q] }
}
`, product.ID)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: runner.NewProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{Config: config, ResourceName: "stripe_coupon.test", ImportState: true, ImportStateId: coupon.ID, ImportStatePersist: true},
			{Config: config, PlanOnly: true, Check: resource.TestCheckResourceAttr("stripe_coupon.test", "applies_to.0.products.0", product.ID)},
			{RefreshState: true},
			{Config: config, PlanOnly: true},
		},
	})
}
