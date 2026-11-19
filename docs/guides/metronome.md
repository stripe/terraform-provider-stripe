---
page_title: "Metronome setup"
subcategory: "Metronome"
description: |-
  Configure Metronome resources using a shared Stripe API key.
---

# Metronome setup

Metronome uses the provider's Stripe API key as a Bearer token. Set
`STRIPE_API_KEY` or the sensitive top-level `api_key` attribute. Metronome
provisions access automatically when API requests use the Stripe key.
Configure the provider and use Metronome resources directly.

Each provider alias has its own clients and key. `METRONOME_API_KEY` and
`METRONOME_API_TOKEN` are not fallbacks. Use Terraform 1.16.5 or later for the
combined integration.

## Use a module

```terraform
provider "stripe" {}

module "billing" {
  source = "./billing"
}
```

The child module declares `stripe/stripe` in its `required_providers` and uses
the inherited provider. For an alias, set
`providers = { stripe = stripe.billing }` on the module.

Metronome uses the account associated with the API key. The `stripe_account`
attribute and `STRIPE_ACCOUNT` environment variable select Stripe Connect
context only and are not supported for Metronome operations. To manage another
Metronome account, use a separate provider alias with that account's key and
leave its Stripe Connect context unset.

## Transport settings and actions

The optional `metronome` object accepts `base_url` and `ca_cert_file`, also
available as `METRONOME_BASE_URL` and `METRONOME_CA_CERT_FILE`. The URL must be
an HTTPS origin. Defaults use `https://api.metronome.com` and system trust.

Actions are explicitly invoked operations, not durable managed objects.
Invoke them with Terraform's action syntax and review their side effects.
Use managed resources for ongoing lifecycle and drift management.
