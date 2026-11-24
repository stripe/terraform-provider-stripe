//
// File generated from our OpenAPI spec
//

package runner

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	shared "github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
	"github.com/stripe/terraform-provider-stripe/internal/combinedprovider"
	metronome "github.com/stripe/terraform-provider-stripe/metronome-go"
)

// Acceptance shares the existing runner's selectors and credential requirements.
// Tests use a unique namespace and own every object they create in the sandbox.
type Acceptance struct {
	Client *metronome.Client
	Name   string
}

func StartAcceptance(t *testing.T, name, surface, kind string) Acceptance {
	t.Helper()
	if !shared.AcceptanceEnabled() {
		t.Skip("set TF_ACC=1 to run Terraform acceptance tests")
	}
	definition := shared.CaseDefinition{Name: name, Surface: surface, Kind: kind, Group: "metronome"}
	if !shared.ShouldRunCase(definition) {
		t.Skip("case filtered out by selectors")
	}
	env := shared.ResolveTestEnv(t, "metronome", nil)
	t.Setenv(resource.EnvTfAccProviderNamespace, "stripe")
	if !strings.Contains(env.APIKey, "_test_") {
		t.Fatal("Metronome acceptance tests require a Stripe sandbox API key")
	}
	// Metronome targets the key's account, even when the nightly job also has
	// connected-account fixtures for Stripe tests.
	env.StripeAccount = ""
	shared.ApplyResolvedAccountEnv(t, env)
	client, err := metronome.New(env.APIKey, "", "", "testacc")
	if err != nil {
		t.Fatal(err)
	}
	var nonce [8]byte
	if _, err = rand.Read(nonce[:]); err != nil {
		t.Fatal(err)
	}
	return Acceptance{Client: client, Name: "tfacc-" + hex.EncodeToString(nonce[:])}
}

func ProviderFactories() map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"stripe": providerserver.NewProtocol6WithError(combinedprovider.New("testacc")()),
	}
}

const ProviderConfig = `terraform {
  required_providers {
    stripe = { source = "stripe/stripe" }
  }
}
provider "stripe" {}
`

func ResourceID(state *terraform.State, address string) string {
	if state == nil || state.RootModule() == nil {
		return ""
	}
	r := state.RootModule().Resources[address]
	if r == nil || r.Primary == nil {
		return ""
	}
	return r.Primary.ID
}

// Call bounds verification and cleanup requests independently of Terraform.
func (a Acceptance) Call(operation metronome.Operation, paths map[string]string, body string) (json.RawMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	var payload json.RawMessage
	if body != "" {
		payload = json.RawMessage(body)
	}
	return a.Client.Execute(ctx, operation, paths, nil, payload)
}
