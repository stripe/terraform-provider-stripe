# Stripe Terraform Provider

Terraform provider for Stripe.

This provider lets you manage Stripe resources with Terraform, including common
objects like products, prices, customers, invoices, subscriptions, payment
intents, webhooks, issuing, terminal, and treasury resources. It also includes
ephemeral resources and action-based operations for workflows such as refunds,
payouts, transfers, and token creation.


Published releases: <https://registry.terraform.io/providers/stripe/stripe/latest>

## Install

```hcl
terraform {
  required_providers {
    stripe = {
      source  = "stripe/stripe"
      version = "= 0.2.2"
    }
  }
}

provider "stripe" {}
```

Set your API key with either the provider argument or the environment variable.

```bash
export STRIPE_API_KEY="sk_test_..."
```

For stable releases, use your normal version constraint. For prereleases,
Terraform only installs an exact prerelease version.

## Example

```hcl
terraform {
  required_providers {
    stripe = {
      source  = "stripe/stripe"
      version = "= 0.3.0-beta1"
    }
  }
}

provider "stripe" {}

resource "stripe_product" "starter" {
  name = "Starter"
}

resource "stripe_price" "starter_monthly" {
  product     = stripe_product.starter.id
  currency    = "usd"
  unit_amount = 1500

  recurring {
    interval = "month"
  }
}
```

## Development

Run the Go test suite:

```bash
go test ./...
```

Run acceptance tests against real Stripe test accounts:

```bash
bash ./scripts/testacc.sh --require-env
```

The acceptance runner loads `.env` if present. `STRIPE_API_KEY` is always
required. Some suites also require `STRIPE_ACCOUNT`,
`STRIPE_ISSUING_ACCOUNT`, or `STRIPE_TREASURY_ACCOUNT`.

If you use `just`, the repo also exposes helper commands in [justfile](./justfile).

## Contributing

Thanks for your interest in the Stripe Terraform Provider! This repository is automatically generated from Stripe's internal tooling. If you've found a bug or have a feature request, open an issue and we'll incorporate the changes through our generator.
