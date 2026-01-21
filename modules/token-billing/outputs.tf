# Token Billing Module Outputs

output "pricing_plan_id" {
  description = "ID of the created pricing plan"
  value       = stripe_v2_billing_pricing_plan.plan.id
}

output "meter_id" {
  description = "ID of the token usage meter (created or existing)"
  value       = local.meter_id
}

output "meter_event_name" {
  description = "Event name for the meter"
  value       = local.meter_event_name
}

output "rate_card_id" {
  description = "ID of the rate card"
  value       = stripe_v2_billing_rate_card.plan_rate_card.id
}

output "license_fee_id" {
  description = "ID of the license fee"
  value       = stripe_v2_billing_license_fee.plan_license_fee.id
}

output "licensed_item_id" {
  description = "ID of the licensed item"
  value       = stripe_v2_billing_licensed_item.plan_access.id
}

output "service_action_id" {
  description = "ID of the credit grant service action (if enabled)"
  value       = var.credit_grant.amount > 0 ? stripe_v2_billing_service_action.credit_grant[0].id : null
}

output "component_ids" {
  description = "IDs of all pricing plan components"
  value = {
    license = stripe_v2_billing_pricing_plan_component.license_component.id
    usage   = stripe_v2_billing_pricing_plan_component.usage_component.id
    credit  = var.credit_grant.amount > 0 ? stripe_v2_billing_pricing_plan_component.credit_component[0].id : null
  }
}

output "metered_items" {
  description = "Map of metered item IDs by model/token type key"
  value = {
    for key, item in stripe_v2_billing_metered_item.tokens :
    key => {
      id           = item.id
      display_name = item.display_name
    }
  }
}

output "model_count" {
  description = "Number of AI models configured"
  value       = length(var.models)
}

output "metered_item_count" {
  description = "Total number of metered items (models × token types)"
  value       = length(local.model_token_map)
}

output "summary" {
  description = "Summary of the pricing plan configuration"
  value = {
    plan_name          = var.plan_name
    plan_id            = stripe_v2_billing_pricing_plan.plan.id
    currency           = var.currency
    license_fee        = "$${var.license_fee.amount / 100}"
    credit_grant       = var.credit_grant.amount > 0 ? "$${var.credit_grant.amount / 100}" : "disabled"
    models_configured  = length(var.models)
    metered_items      = length(local.model_token_map)
    billing_interval   = "${var.service_interval_count} ${var.service_interval}${var.service_interval_count > 1 ? "s" : ""}"
  }
}

