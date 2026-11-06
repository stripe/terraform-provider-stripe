// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	stripe "github.com/stripe/stripe-go/v86"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

func verifyTokenEphemeralState(
	_ runner.TestEnv,
	_ *stripe.Client,
	state *terraform.State,
) error {
	if err := runner.ExpectResourceAbsent(state, "ephemeral.stripe_token.test"); err != nil {
		return err
	}
	for _, outputName := range []string{"token_id", "token_type", "token_reference", "token_bank_country"} {
		if err := runner.ExpectOutputAbsent(state, outputName); err != nil {
			return err
		}
	}

	return nil
}
