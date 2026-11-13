---
page_title: "Metronome setup and migration"
subcategory: "Metronome"
description: |-
  Activate shared Stripe-key access and migrate existing Metronome Terraform state.
---

# Metronome setup and migration

Metronome uses the provider's Stripe API key as a Bearer token. Set
`STRIPE_API_KEY` or the sensitive top-level `api_key` attribute. Each provider
alias has its own clients and key. `METRONOME_API_KEY` and `METRONOME_API_TOKEN`
are not fallbacks. Use Terraform 1.16.5 or later for the combined integration.

## Bootstrap a new module

```terraform
provider "stripe" {}

resource "stripe_metronome_account_bridge" "billing" {
  readiness_timeout = "7m"
}

module "billing" {
  source     = "./billing"
  depends_on = [stripe_metronome_account_bridge.billing]
}
```

The child module declares `stripe/stripe` in its `required_providers` and uses
the inherited provider. It must not define an internal provider configuration
if it is used with module-level `depends_on`. For an alias, set
`providers = { stripe = stripe.billing }` on the module and use that same alias
on the bridge.

Create identifies the authenticated Stripe account, calls
`POST https://api.stripe.com/v1/experimental-met-activate` with an empty payload,
and waits for `GET /v1/credit-types/list?limit=1` on the Metronome origin to
succeed. It probes one page of pricing units, accepts an empty list, and does
not retain their values. The default seven-minute wait allows a five-minute
authorization cache to expire; the configurable range is 5m30s through 1h.
Activation requires a supported account key (`sk_test_`, `sk_live_`, or
`rkcs_test_`), activation permissions, and any required Stripe business profile.
Errors such as disabled activation or a missing business name must be resolved
on the account before retrying. Stripe activation errors remain errors.

The bridge identity is `acct_<id>/test` or `acct_<id>/live`. `client_id` may be
null even after successful activation. Plan, refresh, provider configuration,
and import never activate access. A readiness timeout after activation retains
the account identity so subsequent refresh can recover it. Terraform may mark
a failed create as tainted; inspect its next plan before applying again.

Refresh first verifies the Stripe account and mode. Stripe identity errors
retain the existing state. Read treats probe responses 401, 403, and 404 as a missing bridge and plans
repair. Responses such as 429, 5xx, transport failures, other statuses, and
malformed successes retain state with a diagnostic. This rule applies only to
the bridge. Authorization caching can delay detection after revocation. Key
rotation within the same account and mode preserves identity; a different
account or mode reports an error. `stripe_account` does not switch the
Metronome account; use a key for that account instead.

Destroy removes the bridge from Terraform state. It does not revoke access or
call a remote delete endpoint.

## Bootstrap before refreshing existing state

Module dependencies order creation, but an existing module's resources can
need Metronome access during refresh before the bridge can be repaired.
Activate the intended account in a separate bootstrap configuration first, or
perform the explicit targeted step below before a normal plan or migration:

```shell
terraform apply -target=stripe_metronome_account_bridge.billing
terraform plan
terraform apply
```

Do not use targeting for routine full applies. Keep the bridge in state to
monitor readiness. To adopt an already activated account without an upsert:

```shell
terraform import stripe_metronome_account_bridge.billing acct_example/test
```

## Move standalone state

Back up the state and bootstrap shared-key access first. Keep the old provider
available while Terraform processes the move. The reviewed migration contract
accepts schema version 0 of these standalone managed types: billable metric,
product, rate card, package, customer, customer contract, customer commit,
customer credit, alert, custom field, credit type, billing provider
configuration, revenue system configuration, and notification configuration.
Other source schema versions fail explicitly.

Change each resource type from `metronome_*` to `stripe_metronome_*`, select the
appropriate `stripe` provider alias, and add a `moved` block for its full address:

```terraform
moved {
  from = metronome_customer.customers["example"]
  to   = stripe_metronome_customer.customers["example"]
}
```

For a child module, include `module.billing.` in both addresses. Each moved
resource must retain the same configured attributes. Keep dynamic package
representations unchanged during the move; change representations in a later
plan. Review the plan for preserved IDs and no resource creation, update, or
deletion, then apply the move and confirm a subsequent plan has no changes.
This path supports alerts, which cannot be imported. It also covers typed,
dynamic, and JSON package state and rate cards with managed rates. Do not use
`terraform state replace-provider` alone: it does not rename resource types or
invoke the destination's state-migration contract.

## Transport settings and actions

The optional `metronome` object accepts `base_url` and `ca_cert_file`, also
available as `METRONOME_BASE_URL` and `METRONOME_CA_CERT_FILE`. The URL must be
an HTTPS origin. Defaults use `https://api.metronome.com` and system trust.
The Stripe activation origin remains `https://api.stripe.com`.

Actions are explicitly invoked operations, not durable managed objects.
Invoke them with Terraform's action syntax and review their side effects.
Use managed resources for ongoing lifecycle and drift management. Actions do
not replace the account bridge's dependency and readiness guarantees.
