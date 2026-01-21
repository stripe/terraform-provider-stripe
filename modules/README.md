# Terraform Modules

Hand-built, opinionated Terraform modules that provide simplified interfaces for common Stripe billing patterns.

## Available Modules

### token-billing

Simplifies creation of AI token-based billing plans with:
- Subscription base fees (license fees)
- Usage-based pricing per model and token type
- Optional promotional credit grants
- Multi-provider support (OpenAI, Anthropic, Google, etc.)

[View Documentation](./token-billing/README.md)

## Using Modules

Modules provide a simplified interface compared to using raw provider resources directly.

### Example: Token Billing

**With Module** (~20 lines):

```hcl
module "ai_plan" {
  source = "./modules/token-billing"
  
  plan_name = "Premium AI Plan"
  
  license_fee = {
    amount = 5000  # $50.00/month
  }
  
  models = {
    "openai/gpt-4o" = {
      display_name = "GPT-4o"
      provider     = "OpenAI"
      input_price  = 2500
      output_price = 10000
    }
  }
}
```

**Without Module** (~500+ lines):
- Create billing meter
- Create pricing plan
- Create licensed item
- Create license fee
- Create rate card
- Create metered items (for each model/token type)
- Create rate card rates
- Create service actions
- Create pricing plan components
- Coordinate all dependencies

## When to Use Modules

**Use Modules When:**
- Implementing a common billing pattern
- You want sensible defaults
- You need simplified configuration
- Automatic resource coordination is helpful

**Use Raw Resources When:**
- You need fine-grained control
- Your use case is highly custom
- You're learning the Stripe API
- You need parameters not exposed by modules

## Module Structure

Each module follows this standard structure:

```
module-name/
├── main.tf          # Resource definitions
├── variables.tf     # Input variables
├── outputs.tf       # Output values
└── README.md        # Documentation and examples
```

## Creating Your Own Modules

Modules are hand-built Terraform configurations that use the provider resources. To create a new module:

1. Create directory:

```bash
mkdir modules/my-module
```

2. Write Terraform files:

```hcl
# modules/my-module/main.tf
resource "stripe_product" "main" {
  name = var.product_name
}
```

3. Add documentation:

```markdown
# modules/my-module/README.md
```

4. Create example:

```bash
mkdir examples/using-my-module
# Create example usage
```

5. Test:

```bash
cd examples/using-my-module
terraform init && terraform plan
```

## Philosophy

Modules wrap provider resources to:
- Encode best practices
- Simplify common use cases
- Reduce configuration code
- Prevent common mistakes
- Make patterns reusable

Users can always drop down to the resource level when they need fine-grained control.

## Examples

See the `dev-examples/` directory for working examples of module usage.

## Support

For issues or questions about modules, see the module-specific README files or check the dev-examples directory.



