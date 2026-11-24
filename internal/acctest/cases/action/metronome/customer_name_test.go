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

func TestAccActionMetronomeSetCustomerName(t *testing.T) {
	a := runner.StartAcceptance(t, "metronome_set_customer_name", "stripe_metronome_set_customer_name_v1", "action")
	body, err := a.Call(metronome.Operation{ID: "createCustomer-v1", Method: "POST", Path: "/v1/customers"}, nil, fmt.Sprintf(`{"name":%q}`, a.Name))
	if err != nil {
		t.Fatal(err)
	}
	var created struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	if err = json.Unmarshal(body, &created); err != nil || created.Data.ID == "" {
		t.Fatal("create customer response missing ID")
	}
	id := created.Data.ID
	t.Cleanup(func() {
		_, err := a.Call(metronome.Operation{ID: "archiveCustomer-v1", Method: "POST", Path: "/v1/customers/archive"}, nil, fmt.Sprintf(`{"id":%q}`, id))
		if err != nil {
			t.Errorf("archive action fixture customer: %v", err)
		}
	})
	config := runner.ProviderConfig + fmt.Sprintf(`
action "stripe_metronome_set_customer_name_v1" "test" {
  config {
    path_params = { customer_id = %q }
    request = { name = %q }
  }
}
`, id, a.Name+"-renamed")
	t.Setenv("TF_CLI_ARGS_plan", "-invoke=action.stripe_metronome_set_customer_name_v1.test")
	t.Setenv("TF_CLI_ARGS_apply", "-invoke=action.stripe_metronome_set_customer_name_v1.test")
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: runner.ProviderFactories(),
		Steps: []resource.TestStep{{Config: config, Check: func(_ *terraform.State) error {
			body, err := a.Call(metronome.Operation{ID: "getCustomer-v1", Method: "GET", Path: "/v1/customers/{customer_id}", Read: true}, map[string]string{"customer_id": id}, "")
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
			if response.Data.Name != a.Name+"-renamed" {
				return fmt.Errorf("action did not update remote customer")
			}
			return nil
		}}},
	})
}
