# Stripe Terraform


The Stripe Terraform Provider enables you to manage Stripe resources using infrastructure as code. Configure products, prices, billing meters, and complex pricing plans with declarative Terraform syntax. Your Stripe infrastructure becomes version controlled, reproducible, and auditable. For API details, see the [Stripe API reference](https://stripe.com/docs/api).


## Quick start

In this quickstart we will create a product with a recurring price and webhook endpoint to receive events. 

Create a new Terraform file, `main.tf`:
```hcl
terraform {
  required_providers {
    stripe = {
      source = "stripe/stripe"
      version = "0.2.2"
    }
  }
}

provider "stripe" {
  # API key is read from STRIPE_API_KEY environment variable
  # Alternatively, set it explicitly (not recommended for production)
  # api_key = "sk_test_..."
}


# Define a product
resource "stripe_product" "pro_plan" {
  name        = "Pro Plan"
  description = "Professional tier with advanced features"
}


# Create a recurring price for the product
resource "stripe_price" "pro_monthly" {
  product     = stripe_product.pro_plan.id
  currency    = "usd"
  unit_amount = 2900  # $29.00
  
  recurring {
    interval = "month"
  }
}


# Set up a webhook endpoint for payment events
resource "stripe_webhook_endpoint" "payments" {
  url = "https://api.example.com/webhooks/stripe"
  
  enabled_events = [
    "payment_intent.succeeded",
    "payment_intent.payment_failed",
    "customer.subscription.created",
    "customer.subscription.deleted",
  ]
}

output "price_id" {
  value = stripe_price.pro_monthly.id
}
```

Set your API Key:

```bash
export STRIPE_API_KEY="sk_test_..."
```

Preview and apply changes:

```bash
terraform plan    # Preview changes
terraform apply   # Create resources in Stripe
terraform output  # View created resource IDs
```

## Supported resources

### Product Catalog

| Resource | Description |
|----------|-------------|
| `stripe_product` | Product definitions |
| `stripe_price` | Pricing configurations |
| `stripe_coupon` | Discount coupons |
| `stripe_shipping_rate` | Shipping rate configurations |
| `stripe_tax_rate` | Tax rate definitions |
| `stripe_entitlements_feature` | Feature flags for entitlements |

### Core Resources

| Resource | Description |
|----------|-------------|
| `stripe_customer` | Customer records |
| `stripe_webhook_endpoint` | Webhook endpoint configurations |
| `stripe_billing_meter` | Usage tracking meters |

### Advanced Usage-Based Billing (Private Preview)

These resources are part of the V2 Billing API and require access to the private preview.

| Resource | Description |
|----------|-------------|
| `stripe_v2_billing_pricing_plan` | Pricing plan containers |
| `stripe_v2_billing_pricing_plan_component` | Plan components |
| `stripe_v2_billing_licensed_item` | Licensed access items |
| `stripe_v2_billing_license_fee` | Subscription fees |
| `stripe_v2_billing_metered_item` | Usage-based billable items |
| `stripe_v2_billing_rate_card` | Pricing containers |
| `stripe_v2_billing_rate_card_rate` | Individual rates |
| `stripe_v2_billing_service_action` | Credits and adjustments |

### Data sources

| Data Source | Description |
|-------------|-------------|
| `stripe_billing_meter` | Look up existing billing meters |

## Installation

### Prerequisites

- Terraform 1.0 or later
- Your [Stripe API key](https://dashboard.stripe.com/apikeys) (test mode recommended for initial setup)

### Install from the Terraform Registry

The provider is published to the Terraform Registry as [`stripe/stripe`](https://registry.terraform.io/providers/stripe/stripe/latest). Add it to your Terraform configuration:

```hcl
terraform {
  required_providers {
    stripe = {
      source  = "stripe/stripe"
      version = "0.2.2"
    }
  }
}

provider "stripe" {
  # API key is read from STRIPE_API_KEY environment variable
  # Alternatively, set it explicitly (not recommended for production)
  # api_key = "sk_test_..."
}
```

Set your [Stripe API key](https://dashboard.stripe.com/apikeys) as an environment variable:

```bash
export STRIPE_API_KEY="sk_test_..."
```

Initialize Terraform:

```bash
terraform init
```

Terraform will download the provider from the Registry.

### Local development install

To build and test a local checkout of the provider, install Go 1.19 or later, then run:

```bash
git clone https://github.com/stripe/terraform-provider-stripe.git
cd terraform-provider-stripe
make install
```

The `make install` command builds the provider binary and installs it to `~/.terraform.d/plugins/registry.terraform.io/stripe/stripe/0.2.2/` for your operating system and architecture. Use this path only when developing the provider locally.

Configure Terraform to use the locally installed development build:

```hcl
terraform {
  required_providers {
    stripe = {
      source  = "stripe/stripe"
      version = "0.2.2"
    }
  }
}

provider "stripe" {
  # API key is read from STRIPE_API_KEY environment variable
  # Alternatively, set it explicitly (not recommended for production)
  # api_key = "sk_test_..."
}
```

### Project organization

Organize your Terraform configuration alongside the cloned repository. A typical project structure:

```
my-stripe-infrastructure/
├── main.tf              # Your Terraform configuration
├── variables.tf         # Input variables
├── outputs.tf           # Outputs
└── terraform.tfvars     # Variable values (don't commit secrets!)

terraform-provider-stripe/
  └── modules/
      └── token-billing/
```

With this structure, reference modules using relative paths:

```hcl
module "my_plan" {
  source = "../terraform-provider-stripe/modules/token-billing"
  # ...
}
```

Alternatively, use absolute paths or copy modules directly into your project for a self-contained configuration.

## Managing multiple environments with workspaces

Terraform workspaces allow you to manage separate Stripe environments (sandbox vs livemode) with isolated state files. This prevents accidental changes to production resources when working in test mode.

### Setting up workspaces

Create workspaces for sandbox (test mode) and livemode:

```bash
# Create workspaces
terraform workspace new sandbox
terraform workspace new livemode

# List available workspaces
terraform workspace list
```

### Switching between environments

Each workspace maintains its own state file. Switch workspaces and set the corresponding API key:

```bash
# Work in sandbox (test mode)
terraform workspace select sandbox
export STRIPE_API_KEY="sk_test_..."
terraform plan
terraform apply

# Work in livemode (production)
terraform workspace select livemode
export STRIPE_API_KEY="sk_live_..."
terraform plan
terraform apply
```

## Using modules

Modules provide simplified interfaces for common Stripe billing patterns. Instead of managing individual resources and their dependencies, configure a module with high-level parameters.

### Token billing module (Private Preview)

The token-billing module creates AI token-based pricing plans. It handles the coordination of multiple Stripe billing resources including pricing plans, meters, metered items, rate cards, and license fees.

Modules are located in the cloned repository at `terraform-provider-stripe/modules/`. Reference them using a relative or absolute path to the module directory.

#### Basic configuration

Create a pricing plan with a monthly subscription fee and usage-based pricing for AI models:

```hcl
module "ai_plan" {
  source = "../terraform-provider-stripe/modules/token-billing"
  
  plan_name = "AI Pro Plan"
  
  license_fee = {
    amount = 2000  # $20.00/month
  }
  
  models = {
    "openai/gpt-4o" = {
      display_name         = "GPT-4o"
      provider             = "OpenAI"
      input_price          = 2500
      output_price         = 10000
      cached_input_price   = 1250  # 50% discount for cached tokens
      cached_output_price  = 5000
    }
    "anthropic/claude-sonnet-4.5" = {
      display_name         = "Claude Sonnet 4.5"
      provider             = "Anthropic"
      input_price          = 3000
      output_price         = 15000
      cached_input_price   = 1500
      cached_output_price  = 7500
    }
    "google/gemini-2.5-flash" = {
      display_name = "Gemini 2.5 Flash"
      provider     = "Google"
      input_price  = 300
      output_price = 2500
    }
  }
}

output "plan_id" {
  value = module.ai_plan.pricing_plan_id
}

output "meter_event_name" {
  value = module.ai_plan.meter_event_name
}
```

This configuration creates approximately 15 Stripe resources with correct dependencies and relationships.

## Contributing

Thanks for your interest in the Stripe Terraform Provider! This repository is automatically generated from Stripe's internal tooling. If you've found a bug or have a feature request, open an issue and we'll incorporate the changes through our generator.
