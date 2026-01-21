# Token Billing Module

A Terraform module that simplifies the creation of AI token-based billing plans in Stripe. This module handles the complexity of coordinating multiple Stripe billing resources (pricing plans, meters, metered items, rate cards, license fees, and service actions) into a single interface.

## Features

- Simple interface: Create complete pricing plans with minimal configuration
- Multi-token support: Handles regular and cached token pricing for input and output tokens
- Flexible billing: Combines license fees, usage-based pricing, and optional credit grants
- Rich metadata: Add custom metadata to plans, license fees, and rate cards
- Meter reuse: Share meters across multiple plans or use existing ones
- Production ready: Handles all resource relationships and dependencies automatically

## Usage

### Basic Example

```hcl
module "ai_plan" {
  source = "./modules/token-billing"

  plan_name = "AI Pro Plan"
  
  license_fee = {
    amount = 2000  # $20.00/month
  }
  
  models = {
    "openai/gpt-4o" = {
      display_name = "GPT-4o"
      provider     = "OpenAI"
      input_price  = 2500   # $2.50 per 1M tokens
      output_price = 10000  # $10.00 per 1M tokens
    }
  }
}
```

### Advanced Example with Cached Pricing

```hcl
module "premium_plan" {
  source = "./modules/token-billing"

  plan_name = "Premium AI Plan"
  
  license_fee = {
    amount   = 5000  # $50.00/month
    metadata = {
      tier     = "premium"
      includes = "all_models"
    }
  }
  
  credit_grant = {
    amount = 2000  # $20.00 monthly credit
  }
  
  rate_card = {
    metadata = {
      version  = "v1"
      category = "ai_tokens"
    }
  }
  
  models = {
    "openai/gpt-4o" = {
      display_name         = "GPT-4o"
      provider             = "OpenAI"
      input_price          = 2500   # $2.50 per 1M tokens
      output_price         = 10000  # $10.00 per 1M tokens
      cached_input_price   = 1250   # $1.25 per 1M cached tokens
      cached_output_price  = 5000   # $5.00 per 1M cached tokens
    }
    "anthropic/claude-sonnet-4.5" = {
      display_name         = "Claude Sonnet 4.5"
      provider             = "Anthropic"
      input_price          = 3000
      output_price         = 15000
      cached_input_price   = 1500
      cached_output_price  = 7500
    }
  }
  
  plan_metadata = {
    tier    = "premium"
    support = "priority"
  }
}
```

### Using an Existing Meter

```hcl
module "basic_plan" {
  source = "./modules/token-billing"

  plan_name = "Basic Plan"
  
  license_fee = {
    amount = 1000  # $10.00/month
  }
  
  existing_meter_event_name = "ai_token_usage"
  
  models = {
    "openai/gpt-4o-mini" = {
      display_name = "GPT-4o mini"
      provider     = "OpenAI"
      input_price  = 150
      output_price = 600
    }
  }
}
```

## Input Variables

### Required Variables

| Name | Type | Description |
|------|------|-------------|
| `plan_name` | `string` | Name of the pricing plan |
| `license_fee` | `object` | License fee configuration with `amount` (in cents) and optional `metadata` |
| `models` | `map(object)` | Map of AI models with pricing configuration |

### Model Configuration

Each model in the `models` map should include:

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `display_name` | `string` | Yes | Human-readable name |
| `provider` | `string` | Yes | Provider name (e.g., "OpenAI") |
| `input_price` | `number` | Yes | Price per 1M input tokens in cents |
| `output_price` | `number` | Yes | Price per 1M output tokens in cents |
| `cached_input_price` | `number` | No | Price per 1M cached input tokens in cents |
| `cached_output_price` | `number` | No | Price per 1M cached output tokens in cents |

### Optional Variables

| Name | Type | Default | Description |
|------|------|---------|-------------|
| `currency` | `string` | `"usd"` | Currency code (ISO 4217) |
| `tax_behavior` | `string` | `"exclusive"` | Tax behavior: "inclusive" or "exclusive" |
| `service_interval` | `string` | `"month"` | Billing interval: "day", "week", "month", or "year" |
| `service_interval_count` | `number` | `1` | Number of intervals per billing period |
| `credit_grant` | `object` | `{ amount = 0 }` | Credit grant configuration with `amount` in cents |
| `rate_card` | `object` | `{ metadata = {} }` | Rate card configuration with optional `metadata` |
| `plan_metadata` | `map(string)` | `{}` | Additional metadata for the pricing plan |
| `license_metadata` | `map(string)` | `{}` | Additional metadata for the licensed item |
| `meter_event_name` | `string` | `"ai_token_usage"` | Event name for new meter |
| `existing_meter_event_name` | `string` | `null` | Event name of existing meter to use |
| `create_meter` | `bool` | `true` | Whether to create a new meter |
| `credit_grant_category` | `string` | `"promotional"` | Credit category: "promotional" or "operational" |
| `credit_grant_priority` | `number` | `50` | Credit priority (0-100, higher = applied first) |
| `credit_expiry_type` | `string` | `"end_of_service_period"` | Credit expiry: "end_of_service_period" or "never" |

## Outputs

| Name | Description |
|------|-------------|
| `pricing_plan_id` | ID of the created pricing plan |
| `meter_id` | ID of the meter (created or existing) |
| `meter_event_name` | Event name to send token usage events to |
| `rate_card_id` | ID of the rate card |
| `license_fee_id` | ID of the license fee |
| `licensed_item_id` | ID of the licensed item |
| `service_action_id` | ID of the credit grant (if enabled) |
| `component_ids` | Map of pricing plan component IDs |
| `metered_items` | Map of metered item IDs and display names |
| `model_count` | Number of models configured |
| `metered_item_count` | Total number of metered items created |
| `summary` | Summary of the pricing plan configuration |

## Token Types

The module automatically creates metered items for each configured token type:

- `input`: Regular input tokens (always created)
- `output`: Regular output tokens (always created)
- `cached_input`: Cached input tokens (created if `cached_input_price` is set)
- `cached_output`: Cached output tokens (created if `cached_output_price` is set)

Each metered item uses meter segment conditions to track usage by:
- `model`: The model ID
- `token_type`: The token type (input, output, cached_input, or cached_output)

## Metadata Support

The module supports custom metadata at multiple levels:

### License Fee Metadata

```hcl
license_fee = {
  amount   = 5000
  metadata = {
    tier       = "premium"
    department = "sales"
  }
}
```

### Rate Card Metadata

```hcl
rate_card = {
  metadata = {
    version = "2.0"
    region  = "us-east"
  }
}
```

### Plan Metadata

```hcl
plan_metadata = {
  tier      = "enterprise"
  sales_rep = "john.doe"
}
```

Note: Service actions (credit grants) do not support custom metadata due to Stripe API limitations.

## Architecture

This module creates and coordinates the following Stripe resources:

1. Billing Meter: Tracks token usage events with dimensions (model, token_type)
2. Licensed Item: Represents plan access
3. License Fee: Recurring subscription fee
4. Rate Card: Container for usage-based pricing
5. Metered Items: Individual billable items per model and token type
6. Rate Card Rates: Pricing for each metered item
7. Service Action: Optional credit grant
8. Pricing Plan: Main container with up to three components:
   - License component (subscription fee)
   - Usage component (token usage)
   - Credit component (optional)

## Resource Dependencies

The module handles all resource dependencies automatically:

```
Meter → Metered Items → Rate Card Rates → Rate Card
                                              ↓
License Fee ← Licensed Item              Pricing Plan
                                              ↑
Service Action (optional) ────────────────────┘
```

## Migration from Earlier Versions

If you're upgrading from an earlier version, update your configuration:

**Previous:**

```hcl
license_fee_amount  = 5000
credit_grant_amount = 2000
```

**Current:**

```hcl
license_fee = {
  amount = 5000
}

credit_grant = {
  amount = 2000
}
```

## Requirements

- Terraform >= 1.0
- Stripe Terraform Provider >= 0.1.0

## License

This module is part of the Stripe Terraform Provider project.

