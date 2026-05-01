package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stripe/stripe-go/v84"
)

// DataSourceEntitlementsFeature returns the data source for reading Stripe stripe_entitlements_feature
func DataSourceEntitlementsFeature() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve information about an existing Stripe Entitlements feature.",
		ReadContext: dataSourceEntitlementsFeatureRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Unique identifier for the feature.",
			},
			"lookup_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The lookup key for the feature.",
			},
			"metadata": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Set of key-value pairs attached to the feature.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The feature's internal name.",
			},
			"active": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the feature is active.",
			},
		},
	}
}

func dataSourceEntitlementsFeatureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*stripe.Client)

	if id, ok := d.GetOk("id"); ok {
		obj, err := client.V1EntitlementsFeatures.Retrieve(ctx, id.(string), nil)
		if err != nil {
			return diag.FromErr(err)
		}
		return dataSourceEntitlementsFeatureSetData(d, obj)
	}

	if lookupKey, ok := d.GetOk("lookup_key"); ok {
		obj, err := findEntitlementsFeatureByLookupKey(ctx, client, lookupKey.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		return dataSourceEntitlementsFeatureSetData(d, obj)
	}

	return diag.Errorf("must provide either id or lookup_key")
}

func dataSourceEntitlementsFeatureSetData(d *schema.ResourceData, obj *stripe.EntitlementsFeature) diag.Diagnostics {
	d.SetId(obj.ID)
	d.Set("lookup_key", obj.LookupKey)
	d.Set("metadata", obj.Metadata)
	d.Set("name", obj.Name)
	d.Set("active", obj.Active)
	return nil
}

func findEntitlementsFeatureByLookupKey(ctx context.Context, client *stripe.Client, lookupKey string) (*stripe.EntitlementsFeature, error) {
	params := &stripe.EntitlementsFeatureListParams{}
	params.LookupKey = stripe.String(lookupKey)

	for obj, err := range client.V1EntitlementsFeatures.List(ctx, params) {
		if err != nil {
			return nil, err
		}
		if obj.LookupKey == lookupKey {
			return obj, nil
		}
	}

	archived := true
	archivedParams := &stripe.EntitlementsFeatureListParams{}
	archivedParams.LookupKey = stripe.String(lookupKey)
	archivedParams.Archived = stripe.Bool(archived)

	for obj, err := range client.V1EntitlementsFeatures.List(ctx, archivedParams) {
		if err != nil {
			return nil, err
		}
		if obj.LookupKey == lookupKey {
			return obj, nil
		}
	}

	return nil, fmt.Errorf("no entitlements feature found with lookup_key: %s", lookupKey)
}
