//
// File generated from our OpenAPI spec
//

package metronome_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	runner "github.com/stripe/terraform-provider-stripe/internal/acctest/runner/metronome"
	metronome "github.com/stripe/terraform-provider-stripe/metronome-go"
)

func TestAccManagedMetronomeCatalog(t *testing.T) {
	a := runner.StartAcceptance(t, "metronome_catalog", "stripe_metronome_product", "managed")
	config := func(name string) string {
		return runner.ProviderConfig + fmt.Sprintf(`
resource "stripe_metronome_billable_metric" "test" {
  name = %q
  event_type = %q
  aggregation_type = "COUNT"
}
resource "stripe_metronome_product" "test" {
  name = %q
  type = "USAGE"
  billable_metric_id = stripe_metronome_billable_metric.test.id
}
resource "stripe_metronome_credit_type" "test" { name = %q }
resource "stripe_metronome_rate_card" "test" {
  name = %q
  rates = [{
    product_id = stripe_metronome_product.test.id
    rate_type = "FLAT"
    entitled = true
    price = 12.125
    starting_at = "2026-01-01T00:00:00Z"
  }]
}
data "stripe_metronome_get_product_v1" "test" {
  request = { id = stripe_metronome_product.test.id }
}
data "stripe_metronome_list_credit_types_v1" "test" {
  query_params = { limit = "100" }
  depends_on = [stripe_metronome_credit_type.test]
}
`, name, a.Name, name, a.Name, name)
	}
	var productID, metricID, cardID, creditID string
	check := func(want string) resource.TestCheckFunc {
		return func(state *terraform.State) error {
			productID = runner.ResourceID(state, "stripe_metronome_product.test")
			metricID = runner.ResourceID(state, "stripe_metronome_billable_metric.test")
			cardID = runner.ResourceID(state, "stripe_metronome_rate_card.test")
			creditID = runner.ResourceID(state, "stripe_metronome_credit_type.test")
			if productID == "" || metricID == "" || cardID == "" || creditID == "" {
				return fmt.Errorf("catalog IDs missing from state")
			}
			body, err := a.Call(metronome.Operation{ID: "getProduct-v1", Method: "POST", Path: "/v1/contract-pricing/products/get", Read: true}, nil, fmt.Sprintf(`{"id":%q}`, productID))
			if err != nil {
				return err
			}
			var response struct {
				Data struct {
					Current struct {
						Name string `json:"name"`
					} `json:"current"`
				} `json:"data"`
			}
			if err = json.Unmarshal(body, &response); err != nil {
				return err
			}
			if response.Data.Current.Name != want {
				return fmt.Errorf("remote product name differs from applied configuration")
			}
			return nil
		}
	}
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: runner.ProviderFactories(),
		CheckDestroy: func(_ *terraform.State) error {
			for _, target := range []struct{ id, operation, path, field string }{
				{productID, "getProduct-v1", "/v1/contract-pricing/products/get", "id"},
				{cardID, "getRateCard-v1", "/v1/contract-pricing/rate-cards/get", "id"},
			} {
				if target.id == "" {
					continue
				}
				body, err := a.Call(metronome.Operation{ID: target.operation, Method: "POST", Path: target.path, Read: true}, nil, fmt.Sprintf(`{%q:%q}`, target.field, target.id))
				if metronome.IsNotFound(err) {
					continue
				}
				if err != nil {
					return err
				}
				var response struct {
					Data struct {
						ArchivedAt *string `json:"archived_at"`
					} `json:"data"`
				}
				if err = json.Unmarshal(body, &response); err != nil {
					return err
				}
				if response.Data.ArchivedAt == nil {
					return fmt.Errorf("%s was not archived", target.operation)
				}
			}
			return nil
		},
		Steps: []resource.TestStep{
			{Config: config(a.Name), Check: resource.ComposeTestCheckFunc(check(a.Name), resource.TestCheckResourceAttrSet("data.stripe_metronome_get_product_v1.test", "response_json"))},
			{Config: config(a.Name + "-updated"), Check: check(a.Name + "-updated")},
			{ResourceName: "stripe_metronome_product.test", ImportState: true, ImportStateVerify: true},
			{ResourceName: "stripe_metronome_billable_metric.test", ImportState: true, ImportStateVerify: true},
			{ResourceName: "stripe_metronome_credit_type.test", ImportState: true, ImportStateVerify: true},
			{ResourceName: "stripe_metronome_rate_card.test", ImportState: true, ImportStateVerify: true, ImportStateIdFunc: func(_ *terraform.State) (string, error) {
				return fmt.Sprintf(`{"id":%q,"rates":[{"product_id":%q,"rate_type":"FLAT","entitled":true,"price":12.125,"starting_at":"2026-01-01T00:00:00Z"}]}`, cardID, productID), nil
			}},
		},
	})
}
