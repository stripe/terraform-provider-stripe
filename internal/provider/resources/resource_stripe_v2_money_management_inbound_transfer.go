//
// File generated from our OpenAPI spec
//

package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v87"
)

var _ resource.Resource = &V2MoneyManagementInboundTransferResource{}

var _ resource.ResourceWithConfigure = &V2MoneyManagementInboundTransferResource{}

var _ resource.ResourceWithImportState = &V2MoneyManagementInboundTransferResource{}

func NewV2MoneyManagementInboundTransferResource() resource.Resource {
	return &V2MoneyManagementInboundTransferResource{}
}

type V2MoneyManagementInboundTransferResource struct {
	client *stripe.Client
}

type V2MoneyManagementInboundTransferResourceModel struct {
	Object          types.String `tfsdk:"object"`
	Amount          types.Object `tfsdk:"amount"`
	Created         types.String `tfsdk:"created"`
	Description     types.String `tfsdk:"description"`
	From            types.Object `tfsdk:"from"`
	ID              types.String `tfsdk:"id"`
	Livemode        types.Bool   `tfsdk:"livemode"`
	To              types.Object `tfsdk:"to"`
	TransferHistory types.List   `tfsdk:"transfer_history"`
}

func (r *V2MoneyManagementInboundTransferResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *V2MoneyManagementInboundTransferResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_v2_money_management_inbound_transfer"
}

func (r *V2MoneyManagementInboundTransferResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "An InboundTransfer object, representing a money movement from a\nuser owned PaymentMethod to a FinancialAccount belonging to the same user.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value of the object field.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("v2.money_management.inbound_transfer")},
			},
			"amount": schema.SingleNestedAttribute{
				Required:      true,
				Description:   "The amount in specified currency that will land in the FinancialAccount balance.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"currency": schema.StringAttribute{
						Required:      true,
						Description:   "The currency of the monetary amount.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"value": schema.Int64Attribute{
						Required:      true,
						Description:   "The monetary amount in minor units of the currency.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
					},
				},
			},
			"created": schema.StringAttribute{
				Computed:      true,
				Description:   "Creation time of the InboundTransfer. Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A freeform text field provided by user, containing metadata.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"from": schema.SingleNestedAttribute{
				Required:      true,
				Description:   "A nested object containing information about the origin of the InboundTransfer.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"debited": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "The amount in specified currency that was debited from the Payment Method.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"currency": schema.StringAttribute{
								Computed:      true,
								Description:   "The currency of the monetary amount.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"value": schema.Int64Attribute{
								Computed:      true,
								Description:   "The monetary amount in minor units of the currency.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"payment_method": schema.SingleNestedAttribute{
						Required:      true,
						Description:   "The bank account credential used to fund the InboundTransfer (for example, a US bank account), not a Payments API `pm_` PaymentMethod.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Computed:      true,
								Description:   "The type of PaymentMethod used to create the InboundTransfer. Clients should tolerate future values.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"us_bank_account": schema.StringAttribute{
								Computed:      true,
								Description:   "The destination US bank account identifier. eg \"usba_***\".",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"currency": schema.StringAttribute{
						Optional:      true,
						Description:   "An optional currency field used to specify which currency is debited from the Payment Method.\nSince many Payment Methods support only one currency, this field is optional.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						WriteOnly:     true,
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the InboundTransfer.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"to": schema.SingleNestedAttribute{
				Required:      true,
				Description:   "A nested object containing information about the destination of the InboundTransfer.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"credited": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "The amount by which the FinancialAccount balance is credited.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"currency": schema.StringAttribute{
								Computed:      true,
								Description:   "The currency of the monetary amount.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"value": schema.Int64Attribute{
								Computed:      true,
								Description:   "The monetary amount in minor units of the currency.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"financial_account": schema.StringAttribute{
						Required:      true,
						Description:   "The FinancialAccount that funds will land in.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"currency": schema.StringAttribute{
						Required:      true,
						Description:   "The currency to credit to the FinancialAccount.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						WriteOnly:     true,
					},
				},
			},
			"transfer_history": schema.ListNestedAttribute{
				Computed:      true,
				Description:   "A list of history objects, representing changes in the state of the InboundTransfer. The most recent entry's type indicates the current status of the InboundTransfer.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"bank_debit_failed": schema.SingleNestedAttribute{
							Computed:      true,
							Description:   "The InboundTransfer failed. See `failure_reason` for more details.",
							PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
							Attributes: map[string]schema.Attribute{
								"failure_reason": schema.StringAttribute{
									Computed:      true,
									Description:   "Open Enum. The return reason for the failed InboundTransfer.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									Validators:    []validator.String{stringvalidator.OneOf("bank_account_closed", "bank_account_not_found", "bank_debit_could_not_be_processed", "bank_debit_not_authorized", "insufficient_funds")},
								},
							},
						},
						"bank_debit_returned": schema.SingleNestedAttribute{
							Computed:      true,
							Description:   "The InboundTransfer was returned. The original transaction has been reversed.",
							PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
							Attributes: map[string]schema.Attribute{
								"return_reason": schema.StringAttribute{
									Computed:      true,
									Description:   "Open Enum. The return reason for the returned InboundTransfer.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									Validators:    []validator.String{stringvalidator.OneOf("bank_account_closed", "bank_account_not_found", "bank_debit_could_not_be_processed", "bank_debit_not_authorized", "insufficient_funds")},
								},
							},
						},
						"created": schema.StringAttribute{
							Computed:      true,
							Description:   "Creation time of the HistoryEntry in RFC 3339 format and UTC.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"effective_at": schema.StringAttribute{
							Computed:      true,
							Description:   "Effective at time of the HistoryEntry in RFC 3339 format and UTC.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"id": schema.StringAttribute{
							Computed:      true,
							Description:   "A unique ID for the HistoryEntry.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"type": schema.StringAttribute{
							Computed:      true,
							Description:   "Open Enum. The type of the HistoryEntry.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							Validators:    []validator.String{stringvalidator.OneOf("bank_debit_failed", "bank_debit_processing", "bank_debit_queued", "bank_debit_returned", "bank_debit_succeeded")},
						},
					},
				},
			},
		},
	}
}

func (r *V2MoneyManagementInboundTransferResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan V2MoneyManagementInboundTransferResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config V2MoneyManagementInboundTransferResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"From", "currency"}, []string{"To", "currency"}})

	params, err := expandV2MoneyManagementInboundTransferCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building V2MoneyManagementInboundTransfer create params", err.Error())
		return
	}

	obj, err := r.client.V2MoneyManagementInboundTransfers.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating V2MoneyManagementInboundTransfer", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2MoneyManagementInboundTransfers.B, r.client.V2MoneyManagementInboundTransfers.Key, stripe.FormatURLPath("/v2/money_management/inbound_transfers/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2MoneyManagementInboundTransfer create raw response", err.Error())
		return
	}

	if err := flattenV2MoneyManagementInboundTransfer(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening V2MoneyManagementInboundTransfer create response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"From", "currency"}, []string{"To", "currency"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *V2MoneyManagementInboundTransferResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState V2MoneyManagementInboundTransferResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state V2MoneyManagementInboundTransferResourceModel
	state = priorState

	obj, err := r.client.V2MoneyManagementInboundTransfers.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading V2MoneyManagementInboundTransfer", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2MoneyManagementInboundTransfers.B, r.client.V2MoneyManagementInboundTransfers.Key, stripe.FormatURLPath("/v2/money_management/inbound_transfers/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2MoneyManagementInboundTransfer raw response", err.Error())
		return
	}

	if err := flattenV2MoneyManagementInboundTransfer(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening V2MoneyManagementInboundTransfer read response", err.Error())
		return
	}
	clearWriteOnlyPaths(&state, [][]string{[]string{"From", "currency"}, []string{"To", "currency"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *V2MoneyManagementInboundTransferResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "V2MoneyManagementInboundTransfer does not support updates")
}

func (r *V2MoneyManagementInboundTransferResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *V2MoneyManagementInboundTransferResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandV2MoneyManagementInboundTransferCreate(plan V2MoneyManagementInboundTransferResourceModel) (*stripe.V2MoneyManagementInboundTransferCreateParams, error) {
	params := &stripe.V2MoneyManagementInboundTransferCreateParams{}

	if !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Amount", plan.Amount) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "amount", params)
		}
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "description", params)
		}
	}
	if !plan.From.IsNull() && !plan.From.IsUnknown() {
		if !assignAttrValueToNamedField(params, "From", plan.From) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "from", params)
		}
	}
	if !plan.To.IsNull() && !plan.To.IsUnknown() {
		if !assignAttrValueToNamedField(params, "To", plan.To) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "to", params)
		}
	}

	return params, nil
}

func flattenV2MoneyManagementInboundTransfer(obj *stripe.V2MoneyManagementInboundTransfer, state *V2MoneyManagementInboundTransferResourceModel) error {
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
		assignedAmount := false
		hadRawAmount := false
		if rawValueAmount, rawOk := plainValueAtPath(raw, "amount"); rawOk {
			hadRawAmount = true
			if rawValueAmount != nil {
				sourceAmount := applyConfiguredKeyedListShapes(rawValueAmount, attrValueToPlain(state.Amount))
				if valueAmount, err := flattenPlainValue(sourceAmount, types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "amount", "raw response"); err != nil {
					return err
				} else {
					if typedAmount, ok := valueAmount.(types.Object); ok {
						state.Amount = typedAmount
						assignedAmount = true
					}
				}
			}
		}
		if !assignedAmount {
			if !hasRaw {
				if responseValueAmount, ok := plainFromResponseField(obj, "Amount"); ok {
					sourceAmount := applyConfiguredKeyedListShapes(responseValueAmount, attrValueToPlain(state.Amount))
					if valueAmount, err := flattenPlainValue(
						sourceAmount,
						types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}},
						"amount",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAmount, ok := valueAmount.(types.Object); ok {
							state.Amount = typedAmount
							assignedAmount = true
						}
					}
				}
			}
		}
		if !assignedAmount && hadRawAmount {
			if nullAmount, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}); ok {
				if typedAmount, ok := nullAmount.(types.Object); ok {
					state.Amount = typedAmount
				}
			}
		}
	}
	{
		if rawValueCreated, rawOk := plainValueAtPath(raw, "created"); rawOk {
			if valueCreated, err := flattenPlainValue(rawValueCreated, types.StringType, "created", "raw response"); err != nil {
				return err
			} else {
				if typedCreated, ok := valueCreated.(types.String); ok {
					state.Created = typedCreated
				}
			}
		} else if !hasRaw {
			if responseValueCreated, ok := plainFromResponseField(obj, "Created"); ok {
				if valueCreated, err := flattenPlainValue(responseValueCreated, types.StringType, "created", "response struct"); err != nil {
					return err
				} else {
					if typedCreated, ok := valueCreated.(types.String); ok {
						state.Created = typedCreated
					}
				}
			}
		}
	}
	{
		if rawValueDescription, rawOk := plainValueAtPath(raw, "description"); rawOk {
			if valueDescription, err := flattenPlainValue(rawValueDescription, types.StringType, "description", "raw response"); err != nil {
				return err
			} else {
				if typedDescription, ok := valueDescription.(types.String); ok {
					state.Description = typedDescription
				}
			}
		} else if !hasRaw {
			if responseValueDescription, ok := plainFromResponseField(obj, "Description"); ok {
				if valueDescription, err := flattenPlainValue(responseValueDescription, types.StringType, "description", "response struct"); err != nil {
					return err
				} else {
					if typedDescription, ok := valueDescription.(types.String); ok {
						state.Description = typedDescription
					}
				}
			}
		}
	}
	{
		assignedFrom := false
		hadRawFrom := false
		if rawValueFrom, rawOk := plainValueAtPath(raw, "from"); rawOk {
			hadRawFrom = true
			if rawValueFrom != nil {
				sourceFrom := applyConfiguredKeyedListShapes(rawValueFrom, attrValueToPlain(state.From))
				if valueFrom, err := flattenPlainValue(sourceFrom, types.ObjectType{AttrTypes: map[string]attr.Type{"debited": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "payment_method": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "us_bank_account": types.StringType}}, "currency": types.StringType}}, "from", "raw response"); err != nil {
					return err
				} else {
					if typedFrom, ok := valueFrom.(types.Object); ok {
						state.From = typedFrom
						assignedFrom = true
					}
				}
			}
		}
		if !assignedFrom {
			if !hasRaw {
				if responseValueFrom, ok := plainFromResponseField(obj, "From"); ok {
					sourceFrom := applyConfiguredKeyedListShapes(responseValueFrom, attrValueToPlain(state.From))
					if valueFrom, err := flattenPlainValue(
						sourceFrom,
						types.ObjectType{AttrTypes: map[string]attr.Type{"debited": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "payment_method": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "us_bank_account": types.StringType}}, "currency": types.StringType}},
						"from",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedFrom, ok := valueFrom.(types.Object); ok {
							state.From = typedFrom
							assignedFrom = true
						}
					}
				}
			}
		}
		if !assignedFrom && hadRawFrom {
			if nullFrom, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"debited": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "payment_method": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "us_bank_account": types.StringType}}, "currency": types.StringType}}); ok {
				if typedFrom, ok := nullFrom.(types.Object); ok {
					state.From = typedFrom
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
		assignedTo := false
		hadRawTo := false
		if rawValueTo, rawOk := plainValueAtPath(raw, "to"); rawOk {
			hadRawTo = true
			if rawValueTo != nil {
				sourceTo := applyConfiguredKeyedListShapes(rawValueTo, attrValueToPlain(state.To))
				if valueTo, err := flattenPlainValue(sourceTo, types.ObjectType{AttrTypes: map[string]attr.Type{"credited": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "financial_account": types.StringType, "currency": types.StringType}}, "to", "raw response"); err != nil {
					return err
				} else {
					if typedTo, ok := valueTo.(types.Object); ok {
						state.To = typedTo
						assignedTo = true
					}
				}
			}
		}
		if !assignedTo {
			if !hasRaw {
				if responseValueTo, ok := plainFromResponseField(obj, "To"); ok {
					sourceTo := applyConfiguredKeyedListShapes(responseValueTo, attrValueToPlain(state.To))
					if valueTo, err := flattenPlainValue(
						sourceTo,
						types.ObjectType{AttrTypes: map[string]attr.Type{"credited": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "financial_account": types.StringType, "currency": types.StringType}},
						"to",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedTo, ok := valueTo.(types.Object); ok {
							state.To = typedTo
							assignedTo = true
						}
					}
				}
			}
		}
		if !assignedTo && hadRawTo {
			if nullTo, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"credited": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "financial_account": types.StringType, "currency": types.StringType}}); ok {
				if typedTo, ok := nullTo.(types.Object); ok {
					state.To = typedTo
				}
			}
		}
	}
	{
		if rawValueTransferHistory, rawOk := plainValueAtPath(raw, "transfer_history"); rawOk {
			if valueTransferHistory, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueTransferHistory, attrValueToPlain(state.TransferHistory)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"bank_debit_failed": types.ObjectType{AttrTypes: map[string]attr.Type{"failure_reason": types.StringType}}, "bank_debit_returned": types.ObjectType{AttrTypes: map[string]attr.Type{"return_reason": types.StringType}}, "created": types.StringType, "effective_at": types.StringType, "id": types.StringType, "type": types.StringType}}}, "transfer_history", "raw response"); err != nil {
				return err
			} else {
				if typedTransferHistory, ok := valueTransferHistory.(types.List); ok {
					state.TransferHistory = typedTransferHistory
				}
			}
		} else if !hasRaw {
			if responseValueTransferHistory, ok := plainFromResponseField(obj, "TransferHistory"); ok {
				if valueTransferHistory, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueTransferHistory, attrValueToPlain(state.TransferHistory)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"bank_debit_failed": types.ObjectType{AttrTypes: map[string]attr.Type{"failure_reason": types.StringType}}, "bank_debit_returned": types.ObjectType{AttrTypes: map[string]attr.Type{"return_reason": types.StringType}}, "created": types.StringType, "effective_at": types.StringType, "id": types.StringType, "type": types.StringType}}},
					"transfer_history",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedTransferHistory, ok := valueTransferHistory.(types.List); ok {
						state.TransferHistory = typedTransferHistory
					}
				}
			}
		}
	}
	return nil
}
