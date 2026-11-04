// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
package cases

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

func TestAccManagedProductCustomId(t *testing.T) {
	testAccManagedCustomID(t, "product", `name = "acctest-custom-id"`)
}

func TestAccManagedCouponCustomId(t *testing.T) {
	testAccManagedCustomID(t, "coupon", `duration = "forever"
  percent_off = 10`)
}

func TestAccManagedPlanCustomId(t *testing.T) {
	testAccManagedCustomID(t, "plan", `amount = 1000
  currency = "usd"
  interval = "month"
  product = stripe_product.fixture.id`)
}

func testAccManagedCustomID(t *testing.T, kind, attributes string) {
	t.Helper()
	if !runner.AcceptanceEnabled() {
		t.Skip("set TF_ACC=1 to run Terraform acceptance tests")
	}
	surface := "stripe_" + kind
	if !runner.ShouldRunCase(runner.CaseDefinition{
		Name: kind + "_custom_id", Surface: surface, Group: "base", Kind: "resource",
	}) {
		t.Skip("case filtered out by selectors")
	}
	env := runner.ResolveTestEnv(t, "base", nil)
	runner.ApplyResolvedAccountEnv(t, env)
	t.Setenv(resource.EnvTfAccProviderNamespace, "stripe")
	id := fmt.Sprintf("acctest_%s_%d", kind, time.Now().UnixNano())
	replacementID := id + "_replacement"
	address := surface + ".test"
	config := func(configuredID string) string {
		fixture := ""
		if kind == "plan" {
			fixture = `resource "stripe_product" "fixture" { name = "acctest-custom-plan-id" }`
		}
		idAttribute := ""
		if configuredID != "" {
			idAttribute = fmt.Sprintf("id = %q", configuredID)
		}
		return fmt.Sprintf(`
terraform {
  required_providers {
    stripe = { source = "stripe/stripe" }
  }
}
provider "stripe" {}
%s
resource %q "test" {
  %s
  %s
}
`, fixture, surface, idAttribute, attributes)
	}
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: runner.NewProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{Config: config(id), Check: resource.TestCheckResourceAttr(address, "id", id)},
			{RefreshState: true, Check: resource.TestCheckResourceAttr(address, "id", id)},
			{Config: config(id), ResourceName: address, ImportState: true, ImportStateId: id, ImportStateVerify: true},
			{Config: config(id), PlanOnly: true},
			{
				Config: config(replacementID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectResourceAction(address, plancheck.ResourceActionDestroyBeforeCreate)},
				},
				Check: resource.TestCheckResourceAttr(address, "id", replacementID),
			},
			{RefreshState: true, Check: resource.TestCheckResourceAttr(address, "id", replacementID)},
			{Config: config(replacementID), PlanOnly: true},
		},
	})
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: runner.NewProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{Config: config(""), Check: resource.TestCheckResourceAttrSet(address, "id")},
			{Config: config(""), PlanOnly: true},
		},
	})
}
