# Token Billing Module Variables
# Simplified interface for creating AI token-based billing plans

variable "plan_name" {
  description = "Name of the pricing plan"
  type        = string
}

variable "currency" {
  description = "Currency for all pricing (ISO 4217 code)"
  type        = string
  default     = "usd"
}

variable "tax_behavior" {
  description = "Tax behavior for pricing (inclusive or exclusive)"
  type        = string
  default     = "exclusive"
  
  validation {
    condition     = contains(["inclusive", "exclusive"], var.tax_behavior)
    error_message = "tax_behavior must be either 'inclusive' or 'exclusive'"
  }
}

variable "service_interval" {
  description = "Service interval (day, week, month, year)"
  type        = string
  default     = "month"
  
  validation {
    condition     = contains(["day", "week", "month", "year"], var.service_interval)
    error_message = "service_interval must be one of: day, week, month, year"
  }
}

variable "service_interval_count" {
  description = "Number of service intervals"
  type        = number
  default     = 1
  
  validation {
    condition     = var.service_interval_count > 0
    error_message = "service_interval_count must be greater than 0"
  }
}

variable "license_fee" {
  description = <<-EOT
    License fee configuration:
    - amount: License fee amount in cents (e.g., 2000 = $20.00)
    - metadata: Optional metadata to attach to the license fee (default: {})
  EOT
  type = object({
    amount   = number
    metadata = optional(map(string), {})
  })
  
  validation {
    condition     = var.license_fee.amount >= 0
    error_message = "license_fee.amount must be non-negative"
  }
}

variable "meter_event_name" {
  description = "Event name for the token usage meter (used when creating a new meter)"
  type        = string
  default     = "ai_token_usage"
}

variable "existing_meter_event_name" {
  description = "Event name of an existing billing meter to look up and use. The data source will find the active meter with this event name. If not provided, a new meter will be created."
  type        = string
  default     = null
}

variable "create_meter" {
  description = "Whether to create a new meter when existing_meter_event_name is not provided. Set to false to prevent meter creation."
  type        = bool
  default     = true
}

variable "models" {
  description = <<-EOT
    Map of AI models with their pricing. Each model should have:
    - display_name: Human-readable name
    - provider: Provider name (e.g., "OpenAI", "Anthropic", "Google")
    - input_price: Price per 1M input tokens in cents
    - output_price: Price per 1M output tokens in cents
    - cached_input_price: (Optional) Price per 1M cached input tokens in cents
    - cached_output_price: (Optional) Price per 1M cached output tokens in cents
  EOT
  type = map(object({
    display_name         = string
    provider             = string
    input_price          = number
    output_price         = number
    cached_input_price   = optional(number)
    cached_output_price  = optional(number)
  }))
}

variable "credit_grant" {
  description = <<-EOT
    Credit grant configuration:
    - amount: Credit grant amount in cents (0 or null to disable credit grants)
    Note: Service actions do not support custom metadata in the Stripe API
  EOT
  type = object({
    amount = number
  })
  default = {
    amount = 0
  }
  
  validation {
    condition     = var.credit_grant.amount >= 0
    error_message = "credit_grant.amount must be non-negative"
  }
}

variable "credit_grant_category" {
  description = "Category for credit grants"
  type        = string
  default     = "promotional"
  
  validation {
    condition     = contains(["promotional", "operational"], var.credit_grant_category)
    error_message = "credit_grant_category must be either 'promotional' or 'operational'"
  }
}

variable "credit_grant_priority" {
  description = "Priority for credit grant application (higher = applied first)"
  type        = number
  default     = 50
  
  validation {
    condition     = var.credit_grant_priority >= 0 && var.credit_grant_priority <= 100
    error_message = "credit_grant_priority must be between 0 and 100"
  }
}

variable "credit_expiry_type" {
  description = "Credit expiry type (end_of_service_period or never)"
  type        = string
  default     = "end_of_service_period"
  
  validation {
    condition     = contains(["end_of_service_period", "never"], var.credit_expiry_type)
    error_message = "credit_expiry_type must be either 'end_of_service_period' or 'never'"
  }
}

variable "plan_metadata" {
  description = "Additional metadata to attach to the pricing plan"
  type        = map(string)
  default     = {}
}

variable "rate_card" {
  description = <<-EOT
    Rate card configuration:
    - metadata: Optional metadata to attach to the rate card (default: {})
  EOT
  type = object({
    metadata = optional(map(string), {})
  })
  default = {
    metadata = {}
  }
}

variable "license_metadata" {
  description = "Additional metadata to attach to the licensed item"
  type        = map(string)
  default     = {}
}

