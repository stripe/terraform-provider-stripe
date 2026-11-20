//
// File generated from our OpenAPI spec
//

package metronome_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	runner "github.com/stripe/terraform-provider-stripe/internal/acctest/runner/metronome"
	metronome "github.com/stripe/terraform-provider-stripe/metronome-go"
)

func TestAccManagedMetronomeCustomerContract(t *testing.T) {
	a := runner.StartAcceptance(t, "metronome_customer_contract", "stripe_metronome_contract", "managed")
	start := time.Now().UTC().Truncate(time.Hour).Add(24 * time.Hour).Format(time.RFC3339)
	config := func(name string) string {
		return runner.ProviderConfig + fmt.Sprintf(`
resource "stripe_metronome_customer" "test" {
  name = %q
  ingest_aliases = [%q]
  deletion_protection = false
}
resource "stripe_metronome_custom_field_key" "test" {
  entity = "customer"
  key = %q
  enforce_uniqueness = false
}
resource "stripe_metronome_rate_card" "test" { name = %q }
resource "stripe_metronome_package" "test" {
  package = {
    name = %q
    uniqueness_key = %q
    rate_card_id = stripe_metronome_rate_card.test.id
  }
}
resource "stripe_metronome_contract" "test" {
  customer_id = stripe_metronome_customer.test.id
  package_id = stripe_metronome_package.test.id
  starting_at = %q
  uniqueness_key = %q
}
`, name, a.Name, a.Name, a.Name, a.Name, a.Name, start, a.Name)
	}
	var customerID, contractID string
	check := func(want string) resource.TestCheckFunc {
		return func(state *terraform.State) error {
			customerID = runner.ResourceID(state, "stripe_metronome_customer.test")
			contractID = runner.ResourceID(state, "stripe_metronome_contract.test")
			if customerID == "" || contractID == "" {
				return fmt.Errorf("customer/contract IDs missing")
			}
			body, err := a.Call(metronome.Operation{ID: "getCustomer-v1", Method: "GET", Path: "/v1/customers/{customer_id}", Read: true}, map[string]string{"customer_id": customerID}, "")
			if err != nil {
				return err
			}
			var response struct {
				Data struct {
					Name string `json:"name"`
				} `json:"data"`
			}
			if err = json.Unmarshal(body, &response); err != nil {
				return err
			}
			if response.Data.Name != want {
				return fmt.Errorf("remote customer name differs from applied configuration")
			}
			return nil
		}
	}
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: runner.ProviderFactories(),
		CheckDestroy: func(_ *terraform.State) error {
			if customerID == "" {
				return nil
			}
			body, err := a.Call(metronome.Operation{ID: "getCustomer-v1", Method: "GET", Path: "/v1/customers/{customer_id}", Read: true}, map[string]string{"customer_id": customerID}, "")
			if metronome.IsNotFound(err) {
				return nil
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
				return fmt.Errorf("customer was not archived")
			}
			return nil
		},
		Steps: []resource.TestStep{
			{Config: config(a.Name), Check: check(a.Name)},
			{Config: config(a.Name + "-updated"), Check: check(a.Name + "-updated")},
			{ResourceName: "stripe_metronome_customer.test", ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"deletion_protection"}},
			{ResourceName: "stripe_metronome_custom_field_key.test", ImportState: true, ImportStateVerify: true},
			{ResourceName: "stripe_metronome_package.test", ImportState: true, ImportStateVerify: true, ImportStateIdFunc: func(state *terraform.State) (string, error) {
				return fmt.Sprintf(`{"id":%q,"package":{"name":%q,"uniqueness_key":%q,"rate_card_id":%q}}`, runner.ResourceID(state, "stripe_metronome_package.test"), a.Name, a.Name, runner.ResourceID(state, "stripe_metronome_rate_card.test")), nil
			}},
			{ResourceName: "stripe_metronome_contract.test", ImportState: true, ImportStateVerify: true, ImportStateIdFunc: func(_ *terraform.State) (string, error) { return customerID + "/" + contractID, nil }, ImportStateVerifyIgnore: []string{"uniqueness_key"}},
		},
	})
}
