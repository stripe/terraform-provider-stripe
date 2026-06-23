// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package runner

import stripe "github.com/stripe/stripe-go/v86"

func NewStripeClient(env TestEnv) *stripe.Client {
	if env.StripeAccount != "" {
		backends := stripe.NewBackendsWithConfig(&stripe.BackendConfig{
			StripeContext: stripe.String(env.StripeAccount),
		})
		return stripe.NewClient(env.APIKey, stripe.WithBackends(backends))
	}

	return stripe.NewClient(env.APIKey)
}
