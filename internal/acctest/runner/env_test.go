// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
package runner

import (
	"os"
	"testing"
)

func TestTransferDestinationPreservesRequestScope(t *testing.T) {
	for _, group := range []string{"base", "forwarding", "connect"} {
		t.Run(group, func(t *testing.T) {
			t.Setenv("STRIPE_API_KEY", "test-key")
			t.Setenv("STRIPE_ACCOUNT", "acct_connected")
			t.Setenv("STRIPE_TRANSFER_DESTINATION", "acct_destination")
			env := ResolveTestEnv(t, group, []string{"STRIPE_TRANSFER_DESTINATION"})
			if env.TransferDestinationAccount != "acct_destination" {
				t.Fatalf("transfer destination = %q", env.TransferDestinationAccount)
			}
			wantAccount := ""
			if group == "connect" {
				wantAccount = "acct_connected"
			}
			if env.StripeAccount != wantAccount {
				t.Fatalf("request account = %q, want %q", env.StripeAccount, wantAccount)
			}
			ApplyResolvedAccountEnv(t, env)
			if os.Getenv("STRIPE_ACCOUNT") != wantAccount || os.Getenv("STRIPE_ACCOUNT_CLASSIC") != wantAccount {
				t.Fatal("resolved request account was not applied")
			}
			if os.Getenv("STRIPE_TRANSFER_DESTINATION") != "acct_destination" {
				t.Fatal("applying request scope changed the transfer destination")
			}
			if envValue(env, "STRIPE_TRANSFER_DESTINATION") != "acct_destination" {
				t.Fatal("transfer destination environment mapping is missing")
			}
		})
	}
}
