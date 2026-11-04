// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
package cases

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	stripe "github.com/stripe/stripe-go/v87"
	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

func TestAccManagedV2CoreEventDestinationUrlImport(t *testing.T) {
	if !runner.AcceptanceEnabled() {
		t.Skip("set TF_ACC=1 to run Terraform acceptance tests")
	}
	if !runner.ShouldRunCase(runner.CaseDefinition{
		Name: "v2_core_event_destination_url_import", Surface: "stripe_v2_core_event_destination", Group: "base", Kind: "resource",
	}) {
		t.Skip("case filtered out by selectors")
	}
	env := runner.ResolveTestEnv(t, "base", nil)
	runner.ApplyResolvedAccountEnv(t, env)
	t.Setenv(resource.EnvTfAccProviderNamespace, "stripe")
	client := runner.NewStripeClient(env)
	url := "https://example.com/sdk-codegen/event-destination/import"
	destination, err := client.V2CoreEventDestinations.Create(context.Background(), &stripe.V2CoreEventDestinationCreateParams{
		Name:            stripe.String("acctest-webhook-import"),
		Type:            stripe.String("webhook_endpoint"),
		EventPayload:    stripe.String("thin"),
		EnabledEvents:   stripe.StringSlice([]string{"v1.billing.meter.error_report_triggered"}),
		WebhookEndpoint: &stripe.V2CoreEventDestinationCreateWebhookEndpointParams{URL: stripe.String(url)},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_, err := client.V2CoreEventDestinations.Delete(context.Background(), destination.ID, nil)
		if e, ok := err.(*stripe.Error); ok && e.HTTPStatusCode == 404 {
			return
		}
		if e, ok := err.(*stripe.V2RawError); ok && e.HTTPStatusCode == 404 {
			return
		}
		if err != nil {
			t.Errorf("cleaning up imported event destination: %v", err)
		}
	})
	config := fmt.Sprintf(`
terraform {
  required_providers {
    stripe = { source = "stripe/stripe" }
  }
}
provider "stripe" {}
resource "stripe_v2_core_event_destination" "test" {
  name = "acctest-webhook-import"
  type = "webhook_endpoint"
  event_payload = "thin"
  enabled_events = ["v1.billing.meter.error_report_triggered"]
  include = ["webhook_endpoint.url"]
  webhook_endpoint { url = %q }
}
`, url)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: runner.NewProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{Config: config, ResourceName: "stripe_v2_core_event_destination.test", ImportState: true, ImportStateId: destination.ID, ImportStatePersist: true},
			{Config: config, PlanOnly: true, Check: resource.TestCheckResourceAttr("stripe_v2_core_event_destination.test", "webhook_endpoint.0.url", url)},
			{RefreshState: true},
			{Config: config, PlanOnly: true},
			{Config: strings.Replace(config, url, url+"/updated", 1), PlanOnly: true, ExpectNonEmptyPlan: true},
			{Config: strings.Replace(config, url, url+"/updated", 1), Check: resource.TestCheckResourceAttr("stripe_v2_core_event_destination.test", "webhook_endpoint.0.url", url+"/updated")},
			{Config: strings.Replace(config, url, url+"/updated", 1), PlanOnly: true},
		},
	})
}
