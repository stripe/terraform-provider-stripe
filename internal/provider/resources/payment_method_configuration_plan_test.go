package resources_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	stripe "github.com/stripe/stripe-go/v86"
	stripeprovider "github.com/stripe/terraform-provider-stripe/internal/provider"
)

// Override Configure so tests cannot read application credentials or reach a real
// Stripe account. Every SDK request goes through the in-memory transport below.
type paymentConfigurationTestProvider struct {
	provider.Provider
	client *stripe.Client
}

func (p *paymentConfigurationTestProvider) Configure(_ context.Context, _ provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	resp.ResourceData = p.client
}

type paymentConfigurationTransport struct {
	mu         sync.Mutex
	method     string
	preference string
	updates    int
}

func (m *paymentConfigurationTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if req.URL.Path != "/v1/payment_method_configurations/pmc_imported" {
		return nil, fmt.Errorf("unexpected Stripe request: %s %s", req.Method, req.URL.Path)
	}
	switch req.Method {
	case http.MethodGet:
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			return nil, err
		}
		if preference := req.PostForm.Get(m.method + "[display_preference][preference]"); preference != "" {
			m.preference = preference
			m.updates++
		} else if req.PostForm.Get("active") != "false" {
			return nil, fmt.Errorf("unexpected update: %v", req.PostForm)
		}
	default:
		return nil, fmt.Errorf("unexpected method: %s", req.Method)
	}
	method := func(preference string) map[string]any {
		return map[string]any{
			"available":          preference == "on",
			"display_preference": map[string]any{"preference": preference, "value": preference, "overridable": true},
		}
	}
	body, err := json.Marshal(map[string]any{
		"id": "pmc_imported", "object": "payment_method_configuration", "active": true,
		"is_default": true, "livemode": false, "name": "Default",
		m.method: method(m.preference), "link": method("on"),
	})
	if err != nil {
		return nil, err
	}
	return &http.Response{StatusCode: 200, Header: http.Header{"Content-Type": []string{"application/json"}}, Body: io.NopCloser(strings.NewReader(string(body))), Request: req}, nil
}

func TestPaymentMethodConfigurationImportedPreferenceUpdate(t *testing.T) {
	for _, method := range []string{"apple_pay", "google_pay", "sepa_debit"} {
		t.Run(method, func(t *testing.T) {
			transport := &paymentConfigurationTransport{method: method, preference: "off"}
			client := stripe.NewClient("sk_test_dummy", stripe.WithBackends(stripe.NewBackends(&http.Client{Transport: transport})))
			factories := map[string]func() (tfprotov6.ProviderServer, error){"stripe": providerserver.NewProtocol6WithError(&paymentConfigurationTestProvider{Provider: stripeprovider.New("test")(), client: client})}
			config := func(preference string) string {
				return fmt.Sprintf(`
provider "stripe" {}
resource "stripe_payment_method_configuration" "test" {
 %s = { display_preference = { preference = %q } }
}
`, method, preference)
			}
			const address = "stripe_payment_method_configuration.test"
			update := func(preference string) resource.TestStep {
				return resource.TestStep{
					Config: config(preference),
					ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(address, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(address, tfjsonpath.New("link").AtMapKey("available"), knownvalue.Bool(true)),
						plancheck.ExpectKnownValue(address, tfjsonpath.New("link").AtMapKey("display_preference").AtMapKey("value"), knownvalue.StringExact("on")),
						plancheck.ExpectUnknownValue(address, tfjsonpath.New(method).AtMapKey("available")),
						plancheck.ExpectUnknownValue(address, tfjsonpath.New(method).AtMapKey("display_preference").AtMapKey("value")),
					}},
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(address, method+".available", fmt.Sprint(preference == "on")),
						resource.TestCheckResourceAttr(address, method+".display_preference.value", preference),
						resource.TestCheckResourceAttr(address, "link.display_preference.value", "on"),
					),
				}
			}
			unchanged := func(preference string) resource.TestStep {
				return resource.TestStep{Config: config(preference), ConfigPlanChecks: resource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()}}}
			}
			resource.UnitTest(t, resource.TestCase{
				ProtoV6ProviderFactories: factories,
				Steps: []resource.TestStep{
					{Config: config("off"), ResourceName: address, ImportState: true, ImportStateId: "pmc_imported", ImportStatePersist: true},
					unchanged("off"), update("on"), unchanged("on"), update("off"), unchanged("off"),
				},
			})
			if transport.updates != 2 {
				t.Errorf("got %d preference updates, want 2", transport.updates)
			}
		})
	}
}
