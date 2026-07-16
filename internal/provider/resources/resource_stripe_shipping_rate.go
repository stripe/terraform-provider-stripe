//
// File generated from our OpenAPI spec
//

package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"reflect"
	"strconv"
	"strings"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &ShippingRateResource{}

var _ resource.ResourceWithConfigure = &ShippingRateResource{}

var _ resource.ResourceWithImportState = &ShippingRateResource{}

func NewShippingRateResource() resource.Resource {
	return &ShippingRateResource{}
}

type ShippingRateResource struct {
	client *stripe.Client
}

type ShippingRateResourceModel struct {
	Object           types.String `tfsdk:"object"`
	Active           types.Bool   `tfsdk:"active"`
	Created          types.Int64  `tfsdk:"created"`
	DeliveryEstimate types.List   `tfsdk:"delivery_estimate"`
	DisplayName      types.String `tfsdk:"display_name"`
	FixedAmount      types.List   `tfsdk:"fixed_amount"`
	ID               types.String `tfsdk:"id"`
	Livemode         types.Bool   `tfsdk:"livemode"`
	Metadata         types.Map    `tfsdk:"metadata"`
	TaxBehavior      types.String `tfsdk:"tax_behavior"`
	TaxCode          types.String `tfsdk:"tax_code"`
	Type             types.String `tfsdk:"type"`
}

func (r *ShippingRateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*stripe.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *stripe.Client, got: %T", req.ProviderData))
		return
	}

	r.client = client
}

func (r *ShippingRateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_shipping_rate"
}

var _ resource.ResourceWithUpgradeState = &ShippingRateResource{}

func (r *ShippingRateResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: shippingRateResourceV0Schema(),
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior ShippingRateResourceModel
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgraded, diags := upgradeShippingRateStateV1(ctx, prior)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, &upgraded)...)
			},
		},
		1: {
			PriorSchema: shippingRateResourceV1Schema(),
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior ShippingRateResourceV1Model
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgraded, diags := upgradeShippingRateStateV1(ctx, prior)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, &upgraded)...)
			},
		},
	}
}

func shippingRateResourceV1Schema() *schema.Schema {
	return &schema.Schema{
		Description: "Shipping rates describe the price of shipping presented to your customers and\napplied to a purchase. For more information, see [Charge for shipping](https://docs.stripe.com/payments/during-payment/charge-shipping).",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("shipping_rate")},
			},
			"active": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether the shipping rate can be used for new purchases. Defaults to `true`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"delivery_estimate": schema.SingleNestedAttribute{
				Optional:      true,
				Description:   "The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"maximum": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"unit": schema.StringAttribute{
								Required:      true,
								Description:   "A unit of time.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("business_day", "day", "hour", "month", "week")},
							},
							"value": schema.Int64Attribute{
								Required:      true,
								Description:   "Must be greater than 0.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
							},
						},
					},
					"minimum": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The lower bound of the estimated range. If empty, represents no lower bound.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"unit": schema.StringAttribute{
								Required:      true,
								Description:   "A unit of time.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("business_day", "day", "hour", "month", "week")},
							},
							"value": schema.Int64Attribute{
								Required:      true,
								Description:   "Must be greater than 0.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
							},
						},
					},
				},
			},
			"display_name": schema.StringAttribute{
				Required:      true,
				Description:   "The name of the shipping rate, meant to be displayable to the customer. This will appear on CheckoutSessions.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"fixed_amount": schema.SingleNestedAttribute{
				Optional: true,

				Attributes: map[string]schema.Attribute{
					"amount": schema.Int64Attribute{
						Required:      true,
						Description:   "A non-negative integer in cents representing how much to charge.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
					},
					"currency": schema.StringAttribute{
						Required:      true,
						Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"currency_options": schema.ListNestedAttribute{
						Optional:    true,
						Description: "Shipping rates defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Required:    true,
									Description: "Key for this entry.",
								},
								"amount": schema.Int64Attribute{
									Required:    true,
									Description: "A non-negative integer in cents representing how much to charge.",
								},
								"tax_behavior": schema.StringAttribute{
									Optional:      true,
									Computed:      true,
									Description:   "Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									Validators:    []validator.String{stringvalidator.OneOf("exclusive", "inclusive", "unspecified")},
								},
							},
						},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"tax_behavior": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("exclusive", "inclusive", "unspecified")},
			},
			"tax_code": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A [tax code](https://docs.stripe.com/tax/tax-categories) ID. The Shipping tax code is `txcd_92010001`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The type of calculation to use on the shipping rate.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("fixed_amount")},
			},
		},
	}
}

func shippingRateResourceV0Schema() *schema.Schema {
	return &schema.Schema{
		Description: "Shipping rates describe the price of shipping presented to your customers and\napplied to a purchase. For more information, see [Charge for shipping](https://docs.stripe.com/payments/during-payment/charge-shipping).",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("shipping_rate")},
			},
			"active": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether the shipping rate can be used for new purchases. Defaults to `true`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"display_name": schema.StringAttribute{
				Required:      true,
				Description:   "The name of the shipping rate, meant to be displayable to the customer. This will appear on CheckoutSessions.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"tax_behavior": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("exclusive", "inclusive", "unspecified")},
			},
			"tax_code": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A [tax code](https://docs.stripe.com/tax/tax-categories) ID. The Shipping tax code is `txcd_92010001`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The type of calculation to use on the shipping rate.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("fixed_amount")},
			},
		},
		Blocks: map[string]schema.Block{
			"delivery_estimate": schema.ListNestedBlock{
				Description:   "The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{},
					Blocks: map[string]schema.Block{
						"maximum": schema.ListNestedBlock{
							Description:   "The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.",
							PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown(), listplanmodifier.RequiresReplace()},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"unit": schema.StringAttribute{
										Required:      true,
										Description:   "A unit of time.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("business_day", "day", "hour", "month", "week")},
									},
									"value": schema.Int64Attribute{
										Required:      true,
										Description:   "Must be greater than 0.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
									},
								},
							},
						},
						"minimum": schema.ListNestedBlock{
							Description:   "The lower bound of the estimated range. If empty, represents no lower bound.",
							PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown(), listplanmodifier.RequiresReplace()},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"unit": schema.StringAttribute{
										Required:      true,
										Description:   "A unit of time.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("business_day", "day", "hour", "month", "week")},
									},
									"value": schema.Int64Attribute{
										Required:      true,
										Description:   "Must be greater than 0.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
				},
			},
			"fixed_amount": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"amount": schema.Int64Attribute{
							Required:      true,
							Description:   "A non-negative integer in cents representing how much to charge.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
						},
						"currency": schema.StringAttribute{
							Required:      true,
							Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
					},
					Blocks: map[string]schema.Block{
						"currency_options": schema.ListNestedBlock{
							Description: "Shipping rates defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Required:    true,
										Description: "Key for this entry.",
									},
									"amount": schema.Int64Attribute{
										Required:    true,
										Description: "A non-negative integer in cents representing how much to charge.",
									},
									"tax_behavior": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("exclusive", "inclusive", "unspecified")},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

type ShippingRateResourceV1Model struct {
	Object           types.String `tfsdk:"object"`
	Active           types.Bool   `tfsdk:"active"`
	Created          types.Int64  `tfsdk:"created"`
	DeliveryEstimate types.Object `tfsdk:"delivery_estimate"`
	DisplayName      types.String `tfsdk:"display_name"`
	FixedAmount      types.Object `tfsdk:"fixed_amount"`
	ID               types.String `tfsdk:"id"`
	Livemode         types.Bool   `tfsdk:"livemode"`
	Metadata         types.Map    `tfsdk:"metadata"`
	TaxBehavior      types.String `tfsdk:"tax_behavior"`
	TaxCode          types.String `tfsdk:"tax_code"`
	Type             types.String `tfsdk:"type"`
}

type shippingrateStateUpgradeAttrMeta struct {
	AttrType                attr.Type
	Behavior                string
	LegacyBehavior          string
	PreserveConfiguredValue bool
	Nested                  map[string]shippingrateStateUpgradeAttrMeta
}

var shippingrateStateUpgradeRootMeta = map[string]shippingrateStateUpgradeAttrMeta{"object": shippingrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "active": shippingrateStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "created": shippingrateStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "computed", LegacyBehavior: "computed"}, "delivery_estimate": shippingrateStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"maximum": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}}, "minimum": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}}}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]shippingrateStateUpgradeAttrMeta{"maximum": shippingrateStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed", Nested: map[string]shippingrateStateUpgradeAttrMeta{"unit": shippingrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "value": shippingrateStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "required", LegacyBehavior: "required"}}}, "minimum": shippingrateStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed", Nested: map[string]shippingrateStateUpgradeAttrMeta{"unit": shippingrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "value": shippingrateStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "required", LegacyBehavior: "required"}}}}}, "display_name": shippingrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "fixed_amount": shippingrateStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType, "currency_options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "amount": types.Int64Type, "tax_behavior": types.StringType}}}}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]shippingrateStateUpgradeAttrMeta{"amount": shippingrateStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "required", LegacyBehavior: "required"}, "currency": shippingrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "currency_options": shippingrateStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "amount": types.Int64Type, "tax_behavior": types.StringType}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]shippingrateStateUpgradeAttrMeta{"key": shippingrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "amount": shippingrateStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "required", LegacyBehavior: "required"}, "tax_behavior": shippingrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}}}}}, "id": shippingrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "livemode": shippingrateStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "computed", LegacyBehavior: "computed"}, "metadata": shippingrateStateUpgradeAttrMeta{AttrType: types.MapType{ElemType: types.StringType}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "tax_behavior": shippingrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "tax_code": shippingrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "type": shippingrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}}

var shippingrateStateUpgradeSingletonPaths = map[string]struct{}{}

var shippingrateStateUpgradeLegacyObjectPaths = map[string]struct{}{"delivery_estimate": struct{}{}, "delivery_estimate.maximum": struct{}{}, "delivery_estimate.minimum": struct{}{}, "fixed_amount": struct{}{}}

func shippingrateAttrMapFromModel(model interface{}) map[string]attr.Value {
	value := reflect.ValueOf(model)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if !value.IsValid() || value.Kind() != reflect.Struct {
		return map[string]attr.Value{}
	}

	result := make(map[string]attr.Value, value.NumField())
	typeInfo := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := typeInfo.Field(i)
		tag := field.Tag.Get("tfsdk")
		if tag == "" || tag == "-" {
			continue
		}

		attrValue, ok := value.Field(i).Interface().(attr.Value)
		if !ok {
			continue
		}
		result[tag] = attrValue
	}
	return result
}

func shippingrateSetModelFromAttrMap(target interface{}, values map[string]attr.Value) {
	value := reflect.ValueOf(target)
	if value.Kind() != reflect.Ptr || value.IsNil() {
		return
	}

	elem := value.Elem()
	if !elem.IsValid() || elem.Kind() != reflect.Struct {
		return
	}

	typeInfo := elem.Type()
	for i := 0; i < elem.NumField(); i++ {
		field := typeInfo.Field(i)
		tag := field.Tag.Get("tfsdk")
		if tag == "" || tag == "-" {
			continue
		}

		attrValue, ok := values[tag]
		if !ok {
			continue
		}

		fieldValue := elem.Field(i)
		renderedValue := reflect.ValueOf(attrValue)
		if renderedValue.Type().AssignableTo(fieldValue.Type()) {
			fieldValue.Set(renderedValue)
			continue
		}
		if renderedValue.Type().ConvertibleTo(fieldValue.Type()) {
			fieldValue.Set(renderedValue.Convert(fieldValue.Type()))
		}
	}
}

func shippingrateIsComputedBehavior(behavior string) bool {
	return behavior == "computed" || behavior == "optional_and_computed"
}

func shippingrateShouldPreserveChild(parent shippingrateStateUpgradeAttrMeta, child shippingrateStateUpgradeAttrMeta) bool {
	if parent.PreserveConfiguredValue {
		return true
	}
	if !shippingrateIsComputedBehavior(child.LegacyBehavior) {
		return true
	}
	return !shippingrateIsComputedBehavior(child.Behavior)
}

func shippingrateNullValueForType(attributeType attr.Type) attr.Value {
	switch t := attributeType.(type) {
	case basetypes.StringType:
		return types.StringNull()
	case basetypes.Int64Type:
		return types.Int64Null()
	case basetypes.Float64Type:
		return types.Float64Null()
	case basetypes.BoolType:
		return types.BoolNull()
	case basetypes.MapType:
		return types.MapNull(t.ElemType)
	case basetypes.ListType:
		return types.ListNull(t.ElemType)
	case basetypes.SetType:
		return types.SetNull(t.ElemType)
	case basetypes.ObjectType:
		return types.ObjectNull(t.AttrTypes)
	default:
		return types.StringNull()
	}
}

func shippingrateLegacyUpgradeIsEmptyValue(value attr.Value) bool {
	if value == nil || value.IsNull() || value.IsUnknown() {
		return true
	}
	switch typed := value.(type) {
	case types.String:
		return typed.ValueString() == ""
	case types.Int64:
		return typed.ValueInt64() == 0
	case types.Float64:
		return typed.ValueFloat64() == 0
	case types.Bool:
		return !typed.ValueBool()
	case types.Map:
		return len(typed.Elements()) == 0
	case types.List:
		return len(typed.Elements()) == 0
	case types.Set:
		return len(typed.Elements()) == 0
	case types.Object:
		return len(typed.Attributes()) == 0
	default:
		return false
	}
}

func shippingrateLegacyUpgradeInt64Value(value attr.Value) (int64, bool) {
	switch typed := value.(type) {
	case types.Int64:
		if typed.IsNull() || typed.IsUnknown() {
			return 0, false
		}
		return typed.ValueInt64(), true
	case types.Float64:
		if typed.IsNull() || typed.IsUnknown() {
			return 0, false
		}
		return int64(typed.ValueFloat64()), true
	default:
		return 0, false
	}
}

func shippingrateLegacyUpgradeDecimalMirrorsBase(name string, value attr.Value, siblings map[string]attr.Value) bool {
	typedValue, ok := value.(types.String)
	if !ok || typedValue.IsNull() || typedValue.IsUnknown() || !strings.HasSuffix(name, "_decimal") {
		return false
	}
	baseValue, ok := siblings[strings.TrimSuffix(name, "_decimal")]
	if !ok {
		return false
	}
	baseInt, ok := shippingrateLegacyUpgradeInt64Value(baseValue)
	if !ok {
		return false
	}
	return typedValue.ValueString() == strconv.FormatInt(baseInt, 10)
}

func shippingrateLegacyUpgradeZeroAmountMirrorsDecimal(name string, value attr.Value, siblings map[string]attr.Value) bool {
	typedValue, ok := value.(types.Int64)
	if !ok || typedValue.IsNull() || typedValue.IsUnknown() || typedValue.ValueInt64() != 0 {
		return false
	}
	decimalValue, ok := siblings[name+"_decimal"]
	if !ok {
		return false
	}
	typedDecimal, ok := decimalValue.(types.String)
	if !ok || typedDecimal.IsNull() || typedDecimal.IsUnknown() {
		return false
	}
	return typedDecimal.ValueString() == "0"
}

func shippingrateLegacyUpgradeNormalizeChild(parent shippingrateStateUpgradeAttrMeta, name string, child shippingrateStateUpgradeAttrMeta, value attr.Value, siblings map[string]attr.Value) attr.Value {
	if !parent.PreserveConfiguredValue || child.Behavior != "optional" {
		return value
	}
	if shippingrateLegacyUpgradeIsEmptyValue(value) {
		return shippingrateNullValueForType(child.AttrType)
	}
	if shippingrateLegacyUpgradeDecimalMirrorsBase(name, value, siblings) {
		return shippingrateNullValueForType(child.AttrType)
	}
	if shippingrateLegacyUpgradeZeroAmountMirrorsDecimal(name, value, siblings) {
		return shippingrateNullValueForType(child.AttrType)
	}
	return value
}

func shippingrateLegacyUpgradeChildAttr(path []string, parent shippingrateStateUpgradeAttrMeta, name string, child shippingrateStateUpgradeAttrMeta, source map[string]attr.Value) attr.Value {
	if !shippingrateShouldPreserveChild(parent, child) {
		return shippingrateNullValueForType(child.AttrType)
	}

	childValue, ok := source[name]
	if !ok {
		return shippingrateNullValueForType(child.AttrType)
	}

	nextPath := append(append([]string{}, path...), name)
	upgradedChild := shippingrateUpgradeValue(nextPath, child, childValue)
	return shippingrateLegacyUpgradeNormalizeChild(parent, name, child, upgradedChild, source)
}

func shippingrateUpgradeAttrs(path []string, meta map[string]shippingrateStateUpgradeAttrMeta, prior map[string]attr.Value) map[string]attr.Value {
	upgraded := make(map[string]attr.Value, len(meta))
	for name, fieldMeta := range meta {
		nextPath := append(append([]string{}, path...), name)
		priorValue, ok := prior[name]
		if !ok {
			upgraded[name] = shippingrateNullValueForType(fieldMeta.AttrType)
			continue
		}
		upgradedValue := shippingrateUpgradeValue(nextPath, fieldMeta, priorValue)
		if fieldMeta.PreserveConfiguredValue && fieldMeta.Behavior == "optional" {
			upgradedValue = shippingrateLegacyUpgradeNormalizeChild(
				shippingrateStateUpgradeAttrMeta{PreserveConfiguredValue: true},
				name,
				fieldMeta,
				upgradedValue,
				prior,
			)
		}
		upgraded[name] = upgradedValue
	}
	return upgraded
}

func shippingrateUpgradeObjectValue(path []string, meta shippingrateStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue types.Object) attr.Value {
	if priorValue.IsNull() {
		return types.ObjectNull(objectType.AttrTypes)
	}
	if priorValue.IsUnknown() {
		return types.ObjectUnknown(objectType.AttrTypes)
	}
	sourceAttrs := priorValue.Attributes()
	upgradedAttrs := make(map[string]attr.Value, len(meta.Nested))
	for name, childMeta := range meta.Nested {
		upgradedAttrs[name] = shippingrateLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func shippingrateUpgradeSingletonListToObject(path []string, meta shippingrateStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return shippingrateNullValueForType(meta.AttrType)
	}
	if listValue.IsNull() {
		return types.ObjectNull(objectType.AttrTypes)
	}
	if listValue.IsUnknown() {
		return types.ObjectUnknown(objectType.AttrTypes)
	}

	elements := listValue.Elements()
	if len(elements) == 0 {
		return types.ObjectNull(objectType.AttrTypes)
	}

	firstObject, ok := elements[0].(types.Object)
	if !ok {
		if baseObject, baseOk := elements[0].(basetypes.ObjectValue); baseOk {
			firstObject = types.Object(baseObject)
		} else {
			return types.ObjectUnknown(objectType.AttrTypes)
		}
	}

	sourceAttrs := firstObject.Attributes()
	upgradedAttrs := make(map[string]attr.Value, len(meta.Nested))
	for name, childMeta := range meta.Nested {
		upgradedAttrs[name] = shippingrateLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func shippingrateUpgradeObjectValueToSingletonList(path []string, meta shippingrateStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	if listValue, ok := priorValue.(types.List); ok {
		return shippingrateUpgradeListValue(path, meta, listType, listValue)
	}
	if baseList, ok := priorValue.(basetypes.ListValue); ok {
		return shippingrateUpgradeListValue(path, meta, listType, types.List(baseList))
	}

	objectValue, ok := priorValue.(types.Object)
	if !ok {
		if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
			objectValue = types.Object(baseObject)
		} else {
			return shippingrateNullValueForType(meta.AttrType)
		}
	}
	if objectValue.IsNull() {
		return types.ListNull(listType.ElemType)
	}
	if objectValue.IsUnknown() {
		return types.ListUnknown(listType.ElemType)
	}

	elementObjectType, ok := listType.ElemType.(basetypes.ObjectType)
	if !ok {
		return shippingrateNullValueForType(meta.AttrType)
	}

	upgradedObject := shippingrateUpgradeObjectValue(path, meta, elementObjectType, objectValue)
	return types.ListValueMust(listType.ElemType, []attr.Value{upgradedObject})
}

func shippingrateUpgradeListValue(path []string, meta shippingrateStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return shippingrateNullValueForType(meta.AttrType)
	}
	if listValue.IsNull() {
		return types.ListNull(listType.ElemType)
	}
	if listValue.IsUnknown() {
		return types.ListUnknown(listType.ElemType)
	}
	if len(meta.Nested) == 0 {
		return listValue
	}

	elementObjectType, ok := listType.ElemType.(basetypes.ObjectType)
	if !ok {
		return listValue
	}

	elements := listValue.Elements()
	upgradedElements := make([]attr.Value, 0, len(elements))
	for _, element := range elements {
		objectValue, ok := element.(types.Object)
		if !ok {
			if baseObject, baseOk := element.(basetypes.ObjectValue); baseOk {
				objectValue = types.Object(baseObject)
			} else {
				upgradedElements = append(upgradedElements, shippingrateNullValueForType(elementObjectType))
				continue
			}
		}
		upgradedElements = append(
			upgradedElements,
			shippingrateUpgradeObjectValue(path, meta, elementObjectType, objectValue),
		)
	}

	return types.ListValueMust(listType.ElemType, upgradedElements)
}

func shippingrateUpgradeValue(path []string, meta shippingrateStateUpgradeAttrMeta, priorValue attr.Value) attr.Value {
	pathKey := strings.Join(path, ".")

	switch attrType := meta.AttrType.(type) {
	case basetypes.ObjectType:
		if _, ok := shippingrateStateUpgradeSingletonPaths[pathKey]; ok {
			return shippingrateUpgradeSingletonListToObject(path, meta, attrType, priorValue)
		}

		objectValue, ok := priorValue.(types.Object)
		if !ok {
			if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
				objectValue = types.Object(baseObject)
			} else {
				return shippingrateNullValueForType(meta.AttrType)
			}
		}
		return shippingrateUpgradeObjectValue(path, meta, attrType, objectValue)
	case basetypes.ListType:
		if _, ok := shippingrateStateUpgradeLegacyObjectPaths[pathKey]; ok {
			return shippingrateUpgradeObjectValueToSingletonList(path, meta, attrType, priorValue)
		}
		return shippingrateUpgradeListValue(path, meta, attrType, priorValue)
	default:
		return priorValue
	}
}

func upgradeShippingRateStateV1(ctx context.Context, prior interface{}) (ShippingRateResourceModel, diag.Diagnostics) {
	_ = ctx
	upgradedAttrs := shippingrateUpgradeAttrs(nil, shippingrateStateUpgradeRootMeta, shippingrateAttrMapFromModel(prior))
	var upgraded ShippingRateResourceModel
	shippingrateSetModelFromAttrMap(&upgraded, upgradedAttrs)
	return upgraded, nil
}

func (r *ShippingRateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version:     2,
		Description: "Shipping rates describe the price of shipping presented to your customers and\napplied to a purchase. For more information, see [Charge for shipping](https://docs.stripe.com/payments/during-payment/charge-shipping).",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("shipping_rate")},
			},
			"active": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether the shipping rate can be used for new purchases. Defaults to `true`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"display_name": schema.StringAttribute{
				Required:      true,
				Description:   "The name of the shipping rate, meant to be displayable to the customer. This will appear on CheckoutSessions.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"tax_behavior": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("exclusive", "inclusive", "unspecified")},
			},
			"tax_code": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A [tax code](https://docs.stripe.com/tax/tax-categories) ID. The Shipping tax code is `txcd_92010001`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The type of calculation to use on the shipping rate.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("fixed_amount")},
			},
		},
		Blocks: map[string]schema.Block{
			"delivery_estimate": schema.ListNestedBlock{
				Description:   "The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{},
					Blocks: map[string]schema.Block{
						"maximum": schema.ListNestedBlock{
							Description:   "The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.",
							PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown(), listplanmodifier.RequiresReplace()},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"unit": schema.StringAttribute{
										Required:      true,
										Description:   "A unit of time.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("business_day", "day", "hour", "month", "week")},
									},
									"value": schema.Int64Attribute{
										Required:      true,
										Description:   "Must be greater than 0.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
									},
								},
							},
						},
						"minimum": schema.ListNestedBlock{
							Description:   "The lower bound of the estimated range. If empty, represents no lower bound.",
							PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown(), listplanmodifier.RequiresReplace()},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"unit": schema.StringAttribute{
										Required:      true,
										Description:   "A unit of time.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("business_day", "day", "hour", "month", "week")},
									},
									"value": schema.Int64Attribute{
										Required:      true,
										Description:   "Must be greater than 0.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
				},
			},
			"fixed_amount": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"amount": schema.Int64Attribute{
							Required:      true,
							Description:   "A non-negative integer in cents representing how much to charge.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
						},
						"currency": schema.StringAttribute{
							Required:      true,
							Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
					},
					Blocks: map[string]schema.Block{
						"currency_options": schema.ListNestedBlock{
							Description: "Shipping rates defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Required:    true,
										Description: "Key for this entry.",
									},
									"amount": schema.Int64Attribute{
										Required:    true,
										Description: "A non-negative integer in cents representing how much to charge.",
									},
									"tax_behavior": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("exclusive", "inclusive", "unspecified")},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *ShippingRateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ShippingRateResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandShippingRateCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building ShippingRate create params", err.Error())
		return
	}

	obj, err := r.client.V1ShippingRates.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating ShippingRate", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1ShippingRates.B, r.client.V1ShippingRates.Key, stripe.FormatURLPath("/v1/shipping_rates/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating ShippingRate create raw response", err.Error())
		return
	}

	var createdState ShippingRateResourceModel
	if err := flattenShippingRate(obj, &createdState); err != nil {
		resp.Diagnostics.AddError("Error flattening ShippingRate create response", err.Error())
		return
	}
	normalizeUnknownValues(&createdState)

	diffPlan := plan
	diffCreatedState := createdState

	postCreateParams, err := expandShippingRatePostCreateUpdate(diffPlan, diffCreatedState)
	if err != nil {
		resp.Diagnostics.AddError("Error building ShippingRate post-create update params", err.Error())
		return
	}

	if paramsHaveValues(postCreateParams) {
		if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
			!createdState.Metadata.IsNull() && !createdState.Metadata.IsUnknown() {
			if !assignMetadataDiffToNamedField(postCreateParams, "Metadata", plan.Metadata, createdState.Metadata) {
				resp.Diagnostics.AddError("Error building ShippingRate update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", postCreateParams))
				return
			}
		}
		obj, err = r.client.V1ShippingRates.Update(ctx, createdState.ID.ValueString(), postCreateParams)
		if err != nil {
			resp.Diagnostics.AddError("Error finalizing ShippingRate after create", err.Error())
			return
		}
		if err := ensureRawResponse(obj, r.client.V1ShippingRates.B, r.client.V1ShippingRates.Key, stripe.FormatURLPath("/v1/shipping_rates/%s", obj.ID), nil); err != nil {
			resp.Diagnostics.AddError("Error hydrating ShippingRate post-create update raw response", err.Error())
			return
		}
	}

	if err := flattenShippingRate(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening ShippingRate create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ShippingRateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState ShippingRateResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ShippingRateResourceModel
	state = priorState

	obj, err := r.client.V1ShippingRates.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading ShippingRate", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1ShippingRates.B, r.client.V1ShippingRates.Key, stripe.FormatURLPath("/v1/shipping_rates/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating ShippingRate raw response", err.Error())
		return
	}

	if err := flattenShippingRate(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening ShippingRate read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *ShippingRateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ShippingRateResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ShippingRateResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandShippingRateUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building ShippingRate update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building ShippingRate update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1ShippingRates.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating ShippingRate", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1ShippingRates.B, r.client.V1ShippingRates.Key, stripe.FormatURLPath("/v1/shipping_rates/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating ShippingRate update raw response", err.Error())
		return
	}

	if err := flattenShippingRate(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening ShippingRate update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ShippingRateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ShippingRateResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !state.Active.IsNull() && !state.Active.IsUnknown() && !state.Active.ValueBool() {
		return
	}

	params := &stripe.ShippingRateUpdateParams{}
	activeField := reflect.ValueOf(params).Elem().FieldByName("Active")
	if activeField.IsValid() && activeField.CanSet() {
		if activeField.Kind() == reflect.Pointer && activeField.Type().Elem().Kind() == reflect.Bool {
			activeField.Set(reflect.ValueOf(stripe.Bool(false)))
		} else if activeField.Kind() == reflect.Bool {
			activeField.SetBool(false)
		}
	}

	_, err := r.client.V1ShippingRates.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error deactivating ShippingRate", err.Error())
		return
	}
}

func (r *ShippingRateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandShippingRateCreate(plan ShippingRateResourceModel) (*stripe.ShippingRateCreateParams, error) {
	params := &stripe.ShippingRateCreateParams{}

	if !plan.DeliveryEstimate.IsNull() && !plan.DeliveryEstimate.IsUnknown() {
		if !assignAttrValueToNamedField(params, "DeliveryEstimate", plan.DeliveryEstimate) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "delivery_estimate", params)
		}
	}
	if !plan.DisplayName.IsNull() && !plan.DisplayName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DisplayName", "DisplayName", plan.DisplayName.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "display_name", params)
		}
	}
	if !plan.FixedAmount.IsNull() && !plan.FixedAmount.IsUnknown() {
		if !assignAttrValueToNamedField(params, "FixedAmount", plan.FixedAmount) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "fixed_amount", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.TaxBehavior.IsNull() && !plan.TaxBehavior.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TaxBehavior", "TaxBehavior", plan.TaxBehavior.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tax_behavior", params)
		}
	}
	if !plan.TaxCode.IsNull() && !plan.TaxCode.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TaxCodeID", "TaxCode", plan.TaxCode.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tax_code", params)
		}
	}
	if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Type", "Type", plan.Type.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "type", params)
		}
	}

	return params, nil
}

func expandShippingRateUpdate(plan ShippingRateResourceModel, state ShippingRateResourceModel) (*stripe.ShippingRateUpdateParams, error) {
	params := &stripe.ShippingRateUpdateParams{}

	if !plan.Active.Equal(state.Active) && !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}
	if !plan.FixedAmount.Equal(state.FixedAmount) && !plan.FixedAmount.IsNull() && !plan.FixedAmount.IsUnknown() {
		if !assignAttrValueToNamedField(params, "FixedAmount", plan.FixedAmount) {
			if !plan.FixedAmount.Equal(state.FixedAmount) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "fixed_amount", params)
			}
		}
	}
	if !plan.Metadata.Equal(state.Metadata) && !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			if !plan.Metadata.Equal(state.Metadata) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "metadata", params)
			}
		}
	}
	if !plan.TaxBehavior.Equal(state.TaxBehavior) && !plan.TaxBehavior.IsNull() && !plan.TaxBehavior.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TaxBehavior", "TaxBehavior", plan.TaxBehavior.ValueString()) {
			if !plan.TaxBehavior.Equal(state.TaxBehavior) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "tax_behavior", params)
			}
		}
	}

	return params, nil
}

func expandShippingRatePostCreateUpdate(plan ShippingRateResourceModel, state ShippingRateResourceModel) (*stripe.ShippingRateUpdateParams, error) {
	params := &stripe.ShippingRateUpdateParams{}

	if !plan.Active.Equal(state.Active) && !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}

	return params, nil
}

func flattenShippingRate(obj *stripe.ShippingRate, state *ShippingRateResourceModel) error {
	raw, hasRaw := rawResponseToPlain(obj)
	{
		if rawValueObject, rawOk := plainValueAtPath(raw, "object"); rawOk {
			if valueObject, err := flattenPlainValue(rawValueObject, types.StringType, "object", "raw response"); err != nil {
				return err
			} else {
				if typedObject, ok := valueObject.(types.String); ok {
					state.Object = typedObject
				}
			}
		} else if !hasRaw {
			if responseValueObject, ok := plainFromResponseField(obj, "Object"); ok {
				if valueObject, err := flattenPlainValue(responseValueObject, types.StringType, "object", "response struct"); err != nil {
					return err
				} else {
					if typedObject, ok := valueObject.(types.String); ok {
						state.Object = typedObject
					}
				}
			}
		}
	}
	{
		if rawValueActive, rawOk := plainValueAtPath(raw, "active"); rawOk {
			if valueActive, err := flattenPlainValue(rawValueActive, types.BoolType, "active", "raw response"); err != nil {
				return err
			} else {
				if typedActive, ok := valueActive.(types.Bool); ok {
					state.Active = typedActive
				}
			}
		} else if !hasRaw {
			if responseValueActive, ok := plainFromResponseField(obj, "Active"); ok {
				if valueActive, err := flattenPlainValue(responseValueActive, types.BoolType, "active", "response struct"); err != nil {
					return err
				} else {
					if typedActive, ok := valueActive.(types.Bool); ok {
						state.Active = typedActive
					}
				}
			}
		}
	}
	{
		if rawValueCreated, rawOk := plainValueAtPath(raw, "created"); rawOk {
			if valueCreated, err := flattenPlainValue(rawValueCreated, types.Int64Type, "created", "raw response"); err != nil {
				return err
			} else {
				if typedCreated, ok := valueCreated.(types.Int64); ok {
					state.Created = typedCreated
				}
			}
		} else if !hasRaw {
			if responseValueCreated, ok := plainFromResponseField(obj, "Created"); ok {
				if valueCreated, err := flattenPlainValue(responseValueCreated, types.Int64Type, "created", "response struct"); err != nil {
					return err
				} else {
					if typedCreated, ok := valueCreated.(types.Int64); ok {
						state.Created = typedCreated
					}
				}
			}
		}
	}
	{
		assignedDeliveryEstimate := false
		hadRawDeliveryEstimate := false
		if rawValueDeliveryEstimate, rawOk := plainValueAtPath(raw, "delivery_estimate"); rawOk {
			hadRawDeliveryEstimate = true
			if rawValueDeliveryEstimate != nil {
				sourceDeliveryEstimate := applyConfiguredKeyedListShapes(rawValueDeliveryEstimate, unwrapPlainSingletonList(attrValueToPlain(state.DeliveryEstimate)))
				if valueDeliveryEstimate, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourceDeliveryEstimate, [][]string{[]string{}, []string{"maximum"}, []string{"minimum"}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"maximum": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}}, "minimum": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}}}}}, "delivery_estimate", "raw response"); err != nil {
					return err
				} else {
					if typedDeliveryEstimate, ok := valueDeliveryEstimate.(types.List); ok {
						state.DeliveryEstimate = typedDeliveryEstimate
						assignedDeliveryEstimate = true
					}
				}
			}
		}
		if !assignedDeliveryEstimate {
			if !hasRaw {
				if responseValueDeliveryEstimate, ok := plainFromResponseField(obj, "DeliveryEstimate"); ok {
					sourceDeliveryEstimate := applyConfiguredKeyedListShapes(responseValueDeliveryEstimate, unwrapPlainSingletonList(attrValueToPlain(state.DeliveryEstimate)))
					if valueDeliveryEstimate, err := flattenPlainValue(
						applyPlainSingletonListShapePaths(sourceDeliveryEstimate, [][]string{[]string{}, []string{"maximum"}, []string{"minimum"}}),
						types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"maximum": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}}, "minimum": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}}}}},
						"delivery_estimate",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedDeliveryEstimate, ok := valueDeliveryEstimate.(types.List); ok {
							state.DeliveryEstimate = typedDeliveryEstimate
							assignedDeliveryEstimate = true
						}
					}
				}
			}
		}
		if !assignedDeliveryEstimate && hadRawDeliveryEstimate {
			if nullDeliveryEstimate, ok := nullTerraformValue(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"maximum": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}}, "minimum": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}}}}}); ok {
				if typedDeliveryEstimate, ok := nullDeliveryEstimate.(types.List); ok {
					state.DeliveryEstimate = typedDeliveryEstimate
				}
			}
		}
	}
	{
		if rawValueDisplayName, rawOk := plainValueAtPath(raw, "display_name"); rawOk {
			if valueDisplayName, err := flattenPlainValue(rawValueDisplayName, types.StringType, "display_name", "raw response"); err != nil {
				return err
			} else {
				if typedDisplayName, ok := valueDisplayName.(types.String); ok {
					state.DisplayName = typedDisplayName
				}
			}
		} else if !hasRaw {
			if responseValueDisplayName, ok := plainFromResponseField(obj, "DisplayName"); ok {
				if valueDisplayName, err := flattenPlainValue(responseValueDisplayName, types.StringType, "display_name", "response struct"); err != nil {
					return err
				} else {
					if typedDisplayName, ok := valueDisplayName.(types.String); ok {
						state.DisplayName = typedDisplayName
					}
				}
			}
		}
	}
	{
		assignedFixedAmount := false
		hadRawFixedAmount := false
		if rawValueFixedAmount, rawOk := plainValueAtPath(raw, "fixed_amount"); rawOk {
			hadRawFixedAmount = true
			if rawValueFixedAmount != nil {
				sourceFixedAmount := applyConfiguredKeyedListShapes(rawValueFixedAmount, unwrapPlainSingletonList(attrValueToPlain(state.FixedAmount)))
				if valueFixedAmount, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourceFixedAmount, [][]string{[]string{}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType, "currency_options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "amount": types.Int64Type, "tax_behavior": types.StringType}}}}}}, "fixed_amount", "raw response"); err != nil {
					return err
				} else {
					if typedFixedAmount, ok := valueFixedAmount.(types.List); ok {
						state.FixedAmount = typedFixedAmount
						assignedFixedAmount = true
					}
				}
			}
		}
		if !assignedFixedAmount {
			if !hasRaw {
				if responseValueFixedAmount, ok := plainFromResponseField(obj, "FixedAmount"); ok {
					sourceFixedAmount := applyConfiguredKeyedListShapes(responseValueFixedAmount, unwrapPlainSingletonList(attrValueToPlain(state.FixedAmount)))
					if valueFixedAmount, err := flattenPlainValue(
						applyPlainSingletonListShapePaths(sourceFixedAmount, [][]string{[]string{}}),
						types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType, "currency_options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "amount": types.Int64Type, "tax_behavior": types.StringType}}}}}},
						"fixed_amount",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedFixedAmount, ok := valueFixedAmount.(types.List); ok {
							state.FixedAmount = typedFixedAmount
							assignedFixedAmount = true
						}
					}
				}
			}
		}
		if !assignedFixedAmount && hadRawFixedAmount {
			if nullFixedAmount, ok := nullTerraformValue(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType, "currency_options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "amount": types.Int64Type, "tax_behavior": types.StringType}}}}}}); ok {
				if typedFixedAmount, ok := nullFixedAmount.(types.List); ok {
					state.FixedAmount = typedFixedAmount
				}
			}
		}
	}
	{
		if rawValueID, rawOk := plainValueAtPath(raw, "id"); rawOk {
			if valueID, err := flattenPlainValue(rawValueID, types.StringType, "id", "raw response"); err != nil {
				return err
			} else {
				if typedID, ok := valueID.(types.String); ok {
					state.ID = typedID
				}
			}
		} else if !hasRaw {
			if responseValueID, ok := plainFromResponseField(obj, "ID"); ok {
				if valueID, err := flattenPlainValue(responseValueID, types.StringType, "id", "response struct"); err != nil {
					return err
				} else {
					if typedID, ok := valueID.(types.String); ok {
						state.ID = typedID
					}
				}
			}
		}
	}
	{
		if rawValueLivemode, rawOk := plainValueAtPath(raw, "livemode"); rawOk {
			if valueLivemode, err := flattenPlainValue(rawValueLivemode, types.BoolType, "livemode", "raw response"); err != nil {
				return err
			} else {
				if typedLivemode, ok := valueLivemode.(types.Bool); ok {
					state.Livemode = typedLivemode
				}
			}
		} else if !hasRaw {
			if responseValueLivemode, ok := plainFromResponseField(obj, "Livemode"); ok {
				if valueLivemode, err := flattenPlainValue(responseValueLivemode, types.BoolType, "livemode", "response struct"); err != nil {
					return err
				} else {
					if typedLivemode, ok := valueLivemode.(types.Bool); ok {
						state.Livemode = typedLivemode
					}
				}
			}
		}
	}
	{
		if rawValueMetadata, rawOk := plainValueAtPath(raw, "metadata"); rawOk {
			if valueMetadata, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueMetadata, attrValueToPlain(state.Metadata)), types.MapType{ElemType: types.StringType}, "metadata", "raw response"); err != nil {
				return err
			} else {
				if typedMetadata, ok := valueMetadata.(types.Map); ok {
					state.Metadata = typedMetadata
				}
			}
		} else if !hasRaw {
			if responseValueMetadata, ok := plainFromResponseField(obj, "Metadata"); ok {
				if valueMetadata, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueMetadata, attrValueToPlain(state.Metadata)),
					types.MapType{ElemType: types.StringType},
					"metadata",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedMetadata, ok := valueMetadata.(types.Map); ok {
						state.Metadata = typedMetadata
					}
				}
			}
		}
	}
	{
		if rawValueTaxBehavior, rawOk := plainValueAtPath(raw, "tax_behavior"); rawOk {
			if valueTaxBehavior, err := flattenPlainValue(rawValueTaxBehavior, types.StringType, "tax_behavior", "raw response"); err != nil {
				return err
			} else {
				if typedTaxBehavior, ok := valueTaxBehavior.(types.String); ok {
					state.TaxBehavior = typedTaxBehavior
				}
			}
		} else if !hasRaw {
			if responseValueTaxBehavior, ok := plainFromResponseField(obj, "TaxBehavior"); ok {
				if valueTaxBehavior, err := flattenPlainValue(responseValueTaxBehavior, types.StringType, "tax_behavior", "response struct"); err != nil {
					return err
				} else {
					if typedTaxBehavior, ok := valueTaxBehavior.(types.String); ok {
						state.TaxBehavior = typedTaxBehavior
					}
				}
			}
		}
	}
	{
		if state.TaxCode.IsNull() || state.TaxCode.IsUnknown() {
			if rawValueTaxCode, rawOk := plainValueAtPath(raw, "tax_code"); rawOk {
				if typedTaxCode, ok := plainToStringIDValue(rawValueTaxCode); ok {
					state.TaxCode = typedTaxCode
				}
			} else if !hasRaw {
				if responseValueTaxCode, ok := plainFromResponseField(obj, "TaxCode"); ok {
					if typedTaxCode, ok := plainToStringIDValue(responseValueTaxCode); ok {
						state.TaxCode = typedTaxCode
					}
				}
			}
		}
	}
	{
		if rawValueType, rawOk := plainValueAtPath(raw, "type"); rawOk {
			if valueType, err := flattenPlainValue(rawValueType, types.StringType, "type", "raw response"); err != nil {
				return err
			} else {
				if typedType, ok := valueType.(types.String); ok {
					state.Type = typedType
				}
			}
		} else if !hasRaw {
			if responseValueType, ok := plainFromResponseField(obj, "Type"); ok {
				if valueType, err := flattenPlainValue(responseValueType, types.StringType, "type", "response struct"); err != nil {
					return err
				} else {
					if typedType, ok := valueType.(types.String); ok {
						state.Type = typedType
					}
				}
			}
		}
	}
	return nil
}
