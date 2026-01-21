# Token Billing Module
# A simplified interface for creating AI token-based billing pricing plans
# This module wraps the complexity of coordinating multiple Stripe resources

terraform {
  required_providers {
    stripe = {
      source  = "stripe/stripe"
      version = ">= 0.1.0"
    }
  }
}

locals {
  # Flatten model configurations into individual token type items
  model_token_types = flatten([
    for model_id, model in var.models : concat(
      [
        {
          key          = "${model_id}-input"
          model_id     = model_id
          display_name = model.display_name
          provider     = model.provider
          token_type   = "input"
          price        = model.input_price / 1000000
        },
        {
          key          = "${model_id}-output"
          model_id     = model_id
          display_name = model.display_name
          provider     = model.provider
          token_type   = "output"
          price        = model.output_price / 1000000
        }
      ],
      model.cached_input_price != null ? [
        {
          key          = "${model_id}-cached_input"
          model_id     = model_id
          display_name = model.display_name
          provider     = model.provider
          token_type   = "cached_input"
          price        = model.cached_input_price / 1000000
        }
      ] : [],
      model.cached_output_price != null ? [
        {
          key          = "${model_id}-cached_output"
          model_id     = model_id
          display_name = model.display_name
          provider     = model.provider
          token_type   = "cached_output"
          price        = model.cached_output_price / 1000000
        }
      ] : []
    )
  ])
  
  model_token_map = {
    for item in local.model_token_types :
    item.key => item
  }
}

# Look up existing meter by event name if specified
data "stripe_billing_meter" "existing" {
  count = var.existing_meter_event_name != null ? 1 : 0
  
  event_name = var.existing_meter_event_name
}

# Create the billing meter for tracking token usage (optional - can use existing)
resource "stripe_billing_meter" "token_usage_meter" {
  count = var.create_meter && var.existing_meter_event_name == null ? 1 : 0
  
  display_name = "${var.plan_name} - Token Usage Meter"
  event_name   = var.meter_event_name
  
  dimension_payload_keys = [
    "model",
    "token_type"
  ]
  
  default_aggregation {
    formula = "sum"
  }
}

# Use meter from: looked up by event_name > created
locals {
  meter_id = (
    var.existing_meter_event_name != null ? data.stripe_billing_meter.existing[0].id :
    stripe_billing_meter.token_usage_meter[0].id
  )
  
  meter_event_name = (
    var.existing_meter_event_name != null ? data.stripe_billing_meter.existing[0].event_name :
    stripe_billing_meter.token_usage_meter[0].event_name
  )
}

# Create the main pricing plan
resource "stripe_v2_billing_pricing_plan" "plan" {
  display_name = var.plan_name
  currency     = var.currency
  tax_behavior = var.tax_behavior
  
  metadata = merge(
    {
      plan_type = "ai_token_billing"
      managed_by = "terraform_module"
    },
    var.plan_metadata
  )
}

# Create the licensed item (represents access to the plan)
resource "stripe_v2_billing_licensed_item" "plan_access" {
  display_name = "${var.plan_name} Access"
  
  metadata = merge(
    {
      plan = var.plan_name
    },
    var.license_metadata
  )
}

# Create the license fee (monthly subscription fee)
resource "stripe_v2_billing_license_fee" "plan_license_fee" {
  display_name           = "${var.plan_name} License Fee"
  licensed_item          = stripe_v2_billing_licensed_item.plan_access.id
  currency               = var.currency
  tax_behavior           = var.tax_behavior
  service_interval       = var.service_interval
  service_interval_count = var.service_interval_count
  
  unit_amount = tostring(var.license_fee.amount)
  
  metadata = merge(
    {
      pricing_model = "token_based"
    },
    var.license_fee.metadata
  )
}

# Create the rate card for usage-based pricing
resource "stripe_v2_billing_rate_card" "plan_rate_card" {
  display_name           = "${var.plan_name} Rate Card"
  currency               = var.currency
  service_interval       = var.service_interval
  service_interval_count = var.service_interval_count
  tax_behavior           = var.tax_behavior
  
  metadata = merge(
    {
      plan_type = "token_billing"
    },
    var.rate_card.metadata
  )
}

# Create metered items for each model/token type combination
resource "stripe_v2_billing_metered_item" "tokens" {
  for_each = local.model_token_map
  
  display_name = "${var.plan_name} - ${each.value.display_name} ${each.value.token_type}"
  meter        = local.meter_id
  
  meter_segment_conditions {
    dimension = "model"
    value     = each.value.model_id
  }
  
  meter_segment_conditions {
    dimension = "token_type"
    value     = each.value.token_type
  }
  
  metadata = {
    model      = each.value.model_id
    token_type = each.value.token_type
    provider   = each.value.provider
  }
}

# Create rate card rates for each metered item
resource "stripe_v2_billing_rate_card_rate" "token_rates" {
  for_each = local.model_token_map
  
  rate_card_id = stripe_v2_billing_rate_card.plan_rate_card.id
  metered_item = stripe_v2_billing_metered_item.tokens[each.key].id
  
  unit_amount = format("%.10f", each.value.price)
  
  metadata = {
    model        = each.value.model_id
    token_type   = each.value.token_type
    provider     = each.value.provider
    price_per_1m = format("%.2f", each.value.price * 1000000)
  }
}

# Optional: Create credit grant service action
resource "stripe_v2_billing_service_action" "credit_grant" {
  count = var.credit_grant.amount > 0 ? 1 : 0
  
  type                   = "credit_grant"
  service_interval       = var.service_interval
  service_interval_count = var.service_interval_count
  
  credit_grant {
    name     = "${var.plan_name} Credit Grant"
    category = var.credit_grant_category
    priority = var.credit_grant_priority
    
    amount {
      type = "monetary"
      
      monetary {
        currency = var.currency
        amount   = var.credit_grant.amount
      }
    }
    
    applicability_config {
      scope {
        price_type = "metered"
      }
    }
    
    expiry_config {
      type = var.credit_expiry_type
    }
  }
}

# Component 1: License Fee
resource "stripe_v2_billing_pricing_plan_component" "license_component" {
  pricing_plan_id = stripe_v2_billing_pricing_plan.plan.id
  type            = "license_fee"
  
  license_fee {
    id      = stripe_v2_billing_license_fee.plan_license_fee.id
    version = "latest"
  }
  
  metadata = {
    component_type = "license"
  }
}

# Component 2: Rate Card (Usage-based pricing)
resource "stripe_v2_billing_pricing_plan_component" "usage_component" {
  pricing_plan_id = stripe_v2_billing_pricing_plan.plan.id
  type            = "rate_card"
  
  rate_card {
    id      = stripe_v2_billing_rate_card.plan_rate_card.id
    version = "latest"
  }
  
  metadata = {
    component_type = "token_usage"
  }
  
  depends_on = [stripe_v2_billing_pricing_plan_component.license_component]
}

# Component 3: Credit Grant (optional)
resource "stripe_v2_billing_pricing_plan_component" "credit_component" {
  count = var.credit_grant.amount > 0 ? 1 : 0
  
  pricing_plan_id = stripe_v2_billing_pricing_plan.plan.id
  type            = "service_action"
  
  service_action {
    id = stripe_v2_billing_service_action.credit_grant[0].id
  }
  
  metadata = {
    component_type = "credit"
  }
  
  depends_on = [stripe_v2_billing_pricing_plan_component.usage_component]
}

