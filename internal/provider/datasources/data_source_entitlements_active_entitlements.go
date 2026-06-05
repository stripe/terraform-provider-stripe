package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stripe/stripe-go/v84"
)

// DataSourceEntitlementsActiveEntitlements returns the data source for reading Stripe active entitlements
func DataSourceEntitlementsActiveEntitlements() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve the active Stripe Entitlements for a customer.",
		ReadContext: dataSourceEntitlementsActiveEntitlementsRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Internal identifier for this data source result.",
			},
			"customer": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the customer whose active entitlements to retrieve.",
			},
			"lookup_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Filter results to active entitlements for a specific feature lookup key.",
			},
			"active_entitlements": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The active entitlements for the customer.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier for the active entitlement.",
						},
						"feature": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the feature granted by this entitlement.",
						},
						"lookup_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The lookup key of the entitled feature.",
						},
					},
				},
			},
		},
	}
}

func dataSourceEntitlementsActiveEntitlementsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*stripe.Client)
	customer := d.Get("customer").(string)
	lookupKey, hasLookupKey := d.GetOk("lookup_key")

	params := &stripe.EntitlementsActiveEntitlementListParams{}
	params.Customer = stripe.String(customer)

	items := make([]interface{}, 0)
	for obj, err := range client.V1EntitlementsActiveEntitlements.List(ctx, params) {
		if err != nil {
			return diag.FromErr(err)
		}
		if hasLookupKey && obj.LookupKey != lookupKey.(string) {
			continue
		}

		item := map[string]interface{}{
			"id":         obj.ID,
			"lookup_key": obj.LookupKey,
		}
		if obj.Feature != nil {
			item["feature"] = obj.Feature.ID
		}
		items = append(items, item)
	}

	if hasLookupKey && len(items) == 0 {
		return diag.FromErr(fmt.Errorf("no active entitlements found for customer %s with lookup_key: %s", customer, lookupKey.(string)))
	}

	if hasLookupKey {
		if err := d.Set("lookup_key", lookupKey.(string)); err != nil {
			return diag.FromErr(err)
		}
	}
	if err := d.Set("active_entitlements", items); err != nil {
		return diag.FromErr(err)
	}

	if hasLookupKey {
		d.SetId(fmt.Sprintf("%s/%s", customer, lookupKey.(string)))
	} else {
		d.SetId(customer)
	}

	return nil
}
