// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
package cases

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	stripe "github.com/stripe/stripe-go/v87"
	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

func TestAccManagedWebhookEndpointEventAliases(t *testing.T) {
	if !runner.AcceptanceEnabled() {
		t.Skip("set TF_ACC=1 to run Terraform acceptance tests")
	}
	if !runner.ShouldRunCase(runner.CaseDefinition{
		Name: "webhook_endpoint_event_aliases", Surface: "stripe_webhook_endpoint", Group: "base", Kind: "resource",
	}) {
		t.Skip("webhook alias case excluded by acceptance selectors")
	}
	env := runner.ResolveTestEnv(t, "base", nil)
	key := env.APIKey
	if !strings.HasPrefix(key, "sk_test_") && !strings.HasPrefix(key, "rk_test_") {
		t.Fatal("a test-mode STRIPE_API_KEY is required")
	}
	t.Setenv("STRIPE_ACCOUNT", "")
	t.Setenv(resource.EnvTfAccProviderNamespace, "stripe")
	client := stripe.NewClient(key)
	address := "stripe_webhook_endpoint.aliases"
	for _, tc := range []struct {
		name              string
		events, canonical []string
		apiVersion        string
	}{
		{"transfer_current", []string{"transfer.reversed"}, []string{"transfer.reversed"}, ""},
		{"payment_method_current", []string{"payment_method.automatically_updated"}, []string{"payment_method.automatically_updated"}, ""},
		{"current", []string{"transfer.reversed", "payment_method.automatically_updated"}, []string{"transfer.reversed", "payment_method.automatically_updated"}, ""},
		{"legacy_payment_method", []string{"payment_method.card_automatically_updated"}, []string{"payment_method.card_automatically_updated"}, "2019-12-03"},
		{"legacy_payment_both", []string{"payment_method.automatically_updated", "payment_method.card_automatically_updated"}, []string{"payment_method.automatically_updated", "payment_method.card_automatically_updated"}, "2019-12-03"},
		{"legacy_transfer_distinct", []string{"transfer.reversed", "transfer.canceled", "payout.canceled", "recipient_transfer.canceled"}, []string{"transfer.reversed", "transfer.canceled", "payout.canceled", "recipient_transfer.canceled"}, "2015-01-26"},
		{"transfer_version_boundary", []string{"transfer.reversed"}, []string{"transfer.reversed"}, "2017-04-06"},
		{"payment_version_boundary", []string{"payment_method.automatically_updated"}, []string{"payment_method.automatically_updated"}, "2020-08-27"},
		{"unrelated_events", []string{"customer.updated", "customer.created"}, []string{"customer.updated", "customer.created"}, ""},
		{"duplicate_current", []string{"transfer.reversed", "transfer.reversed"}, []string{"transfer.reversed"}, ""},
		{"wildcard", []string{"*"}, []string{"*"}, ""},
		{"both", []string{"transfer.reversed", "transfer.canceled", "payment_method.automatically_updated", "payment_method.card_automatically_updated"}, []string{"transfer.reversed", "payment_method.automatically_updated"}, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			config := func(events []string) string {
				quoted := make([]string, len(events))
				for i, event := range events {
					quoted[i] = strconv.Quote(event)
				}
				versionConfig := ""
				if tc.apiVersion != "" {
					versionConfig = "api_version = " + strconv.Quote(tc.apiVersion)
				}
				return fmt.Sprintf(`
terraform {
 required_providers {
  stripe = { source = "stripe/stripe" }
 }
}
provider "stripe" {}
resource "stripe_webhook_endpoint" "aliases" {
 url = "https://example.com/terraform-webhook-alias-regression"
 enabled_events = [%s]
 %s
}
`, strings.Join(quoted, ","), versionConfig)
			}
			checkEvents := func(events []string) resource.TestCheckFunc {
				checks := []resource.TestCheckFunc{resource.TestCheckResourceAttr(address, "enabled_events.#", strconv.Itoa(len(events)))}
				for i, event := range events {
					checks = append(checks, resource.TestCheckResourceAttr(address, fmt.Sprintf("enabled_events.%d", i), event))
				}
				return resource.ComposeTestCheckFunc(checks...)
			}
			var endpointID string
			setRemote := func(events []string) func() {
				return func() {
					params := &stripe.WebhookEndpointUpdateParams{}
					for _, event := range events {
						params.EnabledEvents = append(params.EnabledEvents, stripe.String(event))
					}
					if _, err := client.V1WebhookEndpoints.Update(context.Background(), endpointID, params); err != nil {
						t.Fatal(err)
					}
				}
			}
			initial := config(tc.events)
			testCase := resource.TestCase{
				ProtoV6ProviderFactories: runner.NewProtoV6ProviderFactories(),
				CheckDestroy: func(s *terraform.State) error {
					for _, r := range s.RootModule().Resources {
						if r.Type != "stripe_webhook_endpoint" {
							continue
						}
						_, err := client.V1WebhookEndpoints.Retrieve(context.Background(), r.Primary.ID, nil)
						if e, ok := err.(*stripe.Error); ok && e.HTTPStatusCode == 404 {
							continue
						}
						return fmt.Errorf("webhook %s still exists or retrieval failed: %v", r.Primary.ID, err)
					}
					return nil
				},
				Steps: []resource.TestStep{
					{Config: initial, Check: resource.ComposeTestCheckFunc(checkEvents(tc.events), func(s *terraform.State) error {
						endpointID = s.RootModule().Resources[address].Primary.ID
						return nil
					})},
					{RefreshState: true, Check: checkEvents(tc.events)},
					{Config: initial, PlanOnly: true},
				},
			}
			// Exercise the full update/drift lifecycle for each distinct version behavior.
			switch tc.name {
			case "current", "legacy_payment_both", "legacy_transfer_distinct":
				added := append(append([]string{}, tc.events...), "customer.deleted")
				urlOnly := strings.Replace(initial, "terraform-webhook-alias-regression", "terraform-webhook-alias-regression-updated", 1)
				testCase.Steps = append(testCase.Steps,
					resource.TestStep{Config: config(added), PlanOnly: true, ExpectNonEmptyPlan: true},
					resource.TestStep{Config: config(added), Check: checkEvents(added)},
					resource.TestStep{Config: config(added), PlanOnly: true, ExpectNonEmptyPlan: true, PreConfig: setRemote(tc.events)},
					resource.TestStep{Config: config(added), Check: checkEvents(added)},
					resource.TestStep{Config: initial, PlanOnly: true, ExpectNonEmptyPlan: true},
					resource.TestStep{Config: initial, Check: checkEvents(tc.events)},
					resource.TestStep{Config: initial, PlanOnly: true, ExpectNonEmptyPlan: true, PreConfig: setRemote(added)},
					resource.TestStep{Config: initial, Check: checkEvents(tc.events)},
					resource.TestStep{Config: urlOnly, Check: checkEvents(tc.events)},
					resource.TestStep{Config: urlOnly, PlanOnly: true},
					resource.TestStep{Config: initial, Check: checkEvents(tc.events)},
				)
			}
			if !reflect.DeepEqual(tc.events, tc.canonical) {
				testCase.Steps = append(testCase.Steps,
					resource.TestStep{Config: config(tc.canonical), Check: checkEvents(tc.canonical)},
					resource.TestStep{Config: config(tc.canonical), PlanOnly: true},
					resource.TestStep{Config: initial, Check: checkEvents(tc.events)},
				)
			}
			if tc.name == "legacy_payment_both" {
				modernOnly := tc.events[:1]
				testCase.Steps = append(testCase.Steps,
					resource.TestStep{Config: initial, PlanOnly: true, ExpectNonEmptyPlan: true, PreConfig: setRemote(modernOnly)},
					resource.TestStep{Config: initial, Check: checkEvents(tc.events)},
					resource.TestStep{Config: config(modernOnly), Check: checkEvents(modernOnly)},
					resource.TestStep{Config: config(modernOnly), PlanOnly: true, ExpectNonEmptyPlan: true, PreConfig: setRemote(tc.events)},
					resource.TestStep{Config: config(modernOnly), Check: checkEvents(modernOnly)},
				)
			}
			resource.Test(t, testCase)

			// Persist fresh imports for modern aliases and distinct legacy subscriptions.
			switch tc.name {
			case "current", "both", "legacy_payment_method", "legacy_payment_both", "legacy_transfer_distinct":
			default:
				return
			}
			params := &stripe.WebhookEndpointCreateParams{URL: stripe.String("https://example.com/terraform-webhook-alias-import")}
			for _, event := range tc.events {
				params.EnabledEvents = append(params.EnabledEvents, stripe.String(event))
			}
			if tc.apiVersion != "" {
				params.APIVersion = stripe.String(tc.apiVersion)
			}
			endpoint, err := client.V1WebhookEndpoints.Create(context.Background(), params)
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() {
				_, err := client.V1WebhookEndpoints.Delete(context.Background(), endpoint.ID, nil)
				if e, ok := err.(*stripe.Error); ok && e.HTTPStatusCode == 404 {
					return
				}
				if err != nil {
					t.Errorf("cleaning up import endpoint: %v", err)
				}
			})
			importConfig := strings.Replace(config(tc.canonical), "terraform-webhook-alias-regression", "terraform-webhook-alias-import", 1)
			testCase.Steps = []resource.TestStep{
				{Config: importConfig, ResourceName: address, ImportState: true, ImportStateId: endpoint.ID, ImportStatePersist: true},
				{Config: importConfig, PlanOnly: true},
				{RefreshState: true, Check: checkEvents(tc.canonical)},
				{Config: importConfig, PlanOnly: true},
			}
			resource.Test(t, testCase)
		})
	}
}
