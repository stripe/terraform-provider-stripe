package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stripe/stripe-go/v84"
)

// ResourceProductFeature returns the schema for the stripe_product_feature resource
func ResourceProductFeature() *schema.Resource {
	return &schema.Resource{
		Description: "A product feature represents an attachment between a Stripe product and an entitlements feature. When the product is purchased, Stripe creates an active entitlement to the feature for the customer.",

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"product": {
				Type:        schema.TypeString,
				Description: "The ID of the product to attach the feature to.",
				Required:    true,
				ForceNew:    true,
			},
			"entitlement_feature": {
				Type:        schema.TypeString,
				Description: "The ID of the entitlements feature to attach to the product.",
				Required:    true,
				ForceNew:    true,
			},
		},

		CreateContext: resourceProductFeatureCreate,
		ReadContext:   resourceProductFeatureRead,
		DeleteContext: resourceProductFeatureDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceProductFeatureImportState,
		},
	}
}

func resourceProductFeatureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "Creating stripe_product_feature resource")
	c := meta.(*stripe.Client)

	params := &stripe.ProductFeatureCreateParams{}

	if v, ok := d.Get("product").(string); ok && v != "" {
		params.Product = stripe.String(v)
	}
	if v, ok := d.Get("entitlement_feature").(string); ok && v != "" {
		params.EntitlementFeature = stripe.String(v)
	}

	product_feature, err := c.V1ProductFeatures.Create(ctx, params)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create product_feature: %w", err))
	}

	d.SetId(product_feature.ID)
	return resourceProductFeatureRead(ctx, d, meta)
}

func resourceProductFeatureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	importing := ctx.Value("importing") != nil
	_ = importing
	productID, _ := d.Get("product").(string)
	tflog.Debug(ctx, "Reading stripe_product_feature resource", map[string]interface{}{"id": d.Id(), "product": productID})
	c := meta.(*stripe.Client)

	params := &stripe.ProductFeatureRetrieveParams{}
	if productID != "" {
		params.Product = stripe.String(productID)
	}

	product_feature, err := c.V1ProductFeatures.Retrieve(ctx, d.Id(), params)
	if err != nil {
		if stripeErr, ok := err.(*stripe.Error); ok && stripeErr.HTTPStatusCode == 404 {
			d.SetId("")
			return diags
		}
		return diag.FromErr(fmt.Errorf("failed to read product_feature: %w", err))
	}

	if product_feature.EntitlementFeature != nil {
		if err := d.Set("entitlement_feature", product_feature.EntitlementFeature.ID); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	return diags
}

func resourceProductFeatureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	productID, _ := d.Get("product").(string)
	tflog.Debug(ctx, "Deleting stripe_product_feature resource", map[string]interface{}{"id": d.Id(), "product": productID})
	c := meta.(*stripe.Client)

	params := &stripe.ProductFeatureDeleteParams{}
	if productID != "" {
		params.Product = stripe.String(productID)
	}

	_, err := c.V1ProductFeatures.Delete(ctx, d.Id(), params)
	if err != nil {
		if stripeErr, ok := err.(*stripe.Error); ok && stripeErr.HTTPStatusCode == 404 {
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("failed to delete product_feature: %w", err))
	}

	d.SetId("")
	return nil
}

func resourceProductFeatureImportState(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.SplitN(d.Id(), "/", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("import ID must be {product_id}/{id}, got: %s", d.Id())
	}

	d.Set("product", parts[0])
	d.SetId(parts[1])

	diags := resourceProductFeatureRead(context.WithValue(ctx, "importing", true), d, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("%s", diags[0].Summary)
	}

	return []*schema.ResourceData{d}, nil
}
