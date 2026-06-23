//
// File generated from our OpenAPI spec
//

package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &CustomerBalanceTransactionResource{}

var _ resource.ResourceWithConfigure = &CustomerBalanceTransactionResource{}

var _ resource.ResourceWithImportState = &CustomerBalanceTransactionResource{}

func NewCustomerBalanceTransactionResource() resource.Resource {
	return &CustomerBalanceTransactionResource{}
}

type CustomerBalanceTransactionResource struct {
	client *stripe.Client
}

type CustomerBalanceTransactionResourceModel struct {
	Object          types.String `tfsdk:"object"`
	Amount          types.Int64  `tfsdk:"amount"`
	CheckoutSession types.String `tfsdk:"checkout_session"`
	Created         types.Int64  `tfsdk:"created"`
	CreditNote      types.String `tfsdk:"credit_note"`
	Currency        types.String `tfsdk:"currency"`
	Customer        types.String `tfsdk:"customer"`
	Description     types.String `tfsdk:"description"`
	EndingBalance   types.Int64  `tfsdk:"ending_balance"`
	ID              types.String `tfsdk:"id"`
	Invoice         types.String `tfsdk:"invoice"`
	Livemode        types.Bool   `tfsdk:"livemode"`
	Metadata        types.Map    `tfsdk:"metadata"`
	Type            types.String `tfsdk:"type"`
}

func (r *CustomerBalanceTransactionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CustomerBalanceTransactionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_customer_balance_transaction"
}

func (r *CustomerBalanceTransactionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Each customer has a [Balance](https://docs.stripe.com/api/customers/object#customer_object-balance) value,\nwhich denotes a debit or credit that's automatically applied to their next invoice upon finalization.\nYou may modify the value directly by using the [update customer API](https://docs.stripe.com/api/customers/update),\nor by creating a Customer Balance Transaction, which increments or decrements the customer's `balance` by the specified `amount`.\n\nRelated guide: [Customer balance](https://docs.stripe.com/billing/customer/balance)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("customer_balance_transaction")},
			},
			"amount": schema.Int64Attribute{
				Required:      true,
				Description:   "The amount of the transaction. A negative value is a credit for the customer's balance, and a positive value is a debit to the customer's `balance`.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"checkout_session": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of the checkout session (if any) that created the transaction.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"credit_note": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of the credit note (if any) related to the transaction.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				Required:      true,
				Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"customer": schema.StringAttribute{
				Required:      true,
				Description:   "The ID of the customer the transaction belongs to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "An arbitrary string attached to the object. Often useful for displaying to users.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ending_balance": schema.Int64Attribute{
				Computed:      true,
				Description:   "The customer's `balance` after the transaction was applied. A negative value decreases the amount due on the customer's next invoice. A positive value increases the amount due on the customer's next invoice.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"invoice": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of the invoice (if any) related to the transaction.",
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
			"type": schema.StringAttribute{
				Computed:      true,
				Description:   "Transaction type: `adjustment`, `applied_to_invoice`, `credit_note`, `initial`, `invoice_overpaid`, `invoice_too_large`, `invoice_too_small`, `unspent_receiver_credit`, `unapplied_from_invoice`, `checkout_session_subscription_payment`, or `checkout_session_subscription_payment_canceled`. See the [Customer Balance page](https://docs.stripe.com/billing/customer/balance#types) to learn more about transaction types.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("adjustment", "applied_to_invoice", "checkout_session_subscription_payment", "checkout_session_subscription_payment_canceled", "credit_note", "initial", "invoice_overpaid", "invoice_too_large", "invoice_too_small", "migration", "unapplied_from_invoice", "unspent_receiver_credit")},
			},
		},
	}
}

func (r *CustomerBalanceTransactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CustomerBalanceTransactionResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandCustomerBalanceTransactionCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building CustomerBalanceTransaction create params", err.Error())
		return
	}

	if !plan.Customer.IsNull() && !plan.Customer.IsUnknown() {
		if !assignStringToNamedField(params, "Customer", "ID", plan.Customer.ValueString()) {
			resp.Diagnostics.AddError("Error building CustomerBalanceTransaction create path params", fmt.Sprintf("Failed to assign path parameter %q on %T", "customer", params))
			return
		}
	}
	obj, err := r.client.V1CustomerBalanceTransactions.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating CustomerBalanceTransaction", err.Error())
		return
	}

	rawReadParams := &stripe.CustomerBalanceTransactionRetrieveParams{}
	if !plan.Customer.IsNull() && !plan.Customer.IsUnknown() {
		if !assignStringToNamedField(rawReadParams, "Customer", "ID", plan.Customer.ValueString()) {
			resp.Diagnostics.AddError("Error building CustomerBalanceTransaction read params for raw hydration after create", fmt.Sprintf("Failed to assign path parameter %q on %T", "customer", rawReadParams))
			return
		}
	}

	if err := ensureRawResponse(obj, r.client.V1CustomerBalanceTransactions.B, r.client.V1CustomerBalanceTransactions.Key, stripe.FormatURLPath("/v1/customers/%s/balance_transactions/%s", plan.Customer.ValueString(), obj.ID), rawReadParams); err != nil {
		resp.Diagnostics.AddError("Error hydrating CustomerBalanceTransaction create raw response", err.Error())
		return
	}

	if err := flattenCustomerBalanceTransaction(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening CustomerBalanceTransaction create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CustomerBalanceTransactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState CustomerBalanceTransactionResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state CustomerBalanceTransactionResourceModel
	state = priorState

	params := &stripe.CustomerBalanceTransactionRetrieveParams{}
	if !state.Customer.IsNull() && !state.Customer.IsUnknown() {
		if !assignStringToNamedField(params, "Customer", "ID", state.Customer.ValueString()) {
			resp.Diagnostics.AddError("Error building CustomerBalanceTransaction read params", fmt.Sprintf("Failed to assign path parameter %q on %T", "customer", params))
			return
		}
	}

	obj, err := r.client.V1CustomerBalanceTransactions.Retrieve(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error reading CustomerBalanceTransaction", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1CustomerBalanceTransactions.B, r.client.V1CustomerBalanceTransactions.Key, stripe.FormatURLPath("/v1/customers/%s/balance_transactions/%s", state.Customer.ValueString(), state.ID.ValueString()), params); err != nil {
		resp.Diagnostics.AddError("Error hydrating CustomerBalanceTransaction raw response", err.Error())
		return
	}

	if err := flattenCustomerBalanceTransaction(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening CustomerBalanceTransaction read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *CustomerBalanceTransactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan CustomerBalanceTransactionResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state CustomerBalanceTransactionResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandCustomerBalanceTransactionUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building CustomerBalanceTransaction update params", err.Error())
		return
	}

	if !state.Customer.IsNull() && !state.Customer.IsUnknown() {
		if !assignStringToNamedField(params, "Customer", "ID", state.Customer.ValueString()) {
			resp.Diagnostics.AddError("Error building CustomerBalanceTransaction update path params", fmt.Sprintf("Failed to assign path parameter %q on %T", "customer", params))
			return
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building CustomerBalanceTransaction update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1CustomerBalanceTransactions.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating CustomerBalanceTransaction", err.Error())
		return
	}

	rawReadParams := &stripe.CustomerBalanceTransactionRetrieveParams{}
	if !state.Customer.IsNull() && !state.Customer.IsUnknown() {
		if !assignStringToNamedField(rawReadParams, "Customer", "ID", state.Customer.ValueString()) {
			resp.Diagnostics.AddError("Error building CustomerBalanceTransaction read params for raw hydration after update", fmt.Sprintf("Failed to assign path parameter %q on %T", "customer", rawReadParams))
			return
		}
	}

	if err := ensureRawResponse(obj, r.client.V1CustomerBalanceTransactions.B, r.client.V1CustomerBalanceTransactions.Key, stripe.FormatURLPath("/v1/customers/%s/balance_transactions/%s", state.Customer.ValueString(), obj.ID), rawReadParams); err != nil {
		resp.Diagnostics.AddError("Error hydrating CustomerBalanceTransaction update raw response", err.Error())
		return
	}

	if err := flattenCustomerBalanceTransaction(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening CustomerBalanceTransaction update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CustomerBalanceTransactionResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *CustomerBalanceTransactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Unexpected import identifier", fmt.Sprintf("Expected import identifier in the form \"customer/id\", got %q", req.ID))
		return
	}

	diags := resp.State.SetAttribute(ctx, path.Root("customer"), types.StringValue(parts[0]))
	resp.Diagnostics.Append(diags...)
	diags = resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(parts[1]))
	resp.Diagnostics.Append(diags...)
}

func expandCustomerBalanceTransactionCreate(plan CustomerBalanceTransactionResourceModel) (*stripe.CustomerBalanceTransactionCreateParams, error) {
	params := &stripe.CustomerBalanceTransactionCreateParams{}

	if !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		params.Amount = stripe.Int64(plan.Amount.ValueInt64())
	}
	if !plan.Currency.IsNull() && !plan.Currency.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Currency", "Currency", plan.Currency.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "currency", params)
		}
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "description", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}

	return params, nil
}

func expandCustomerBalanceTransactionUpdate(plan CustomerBalanceTransactionResourceModel, state CustomerBalanceTransactionResourceModel) (*stripe.CustomerBalanceTransactionUpdateParams, error) {
	params := &stripe.CustomerBalanceTransactionUpdateParams{}

	if !plan.Description.Equal(state.Description) && !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			if !plan.Description.Equal(state.Description) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "description", params)
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

	return params, nil
}

func flattenCustomerBalanceTransaction(obj *stripe.CustomerBalanceTransaction, state *CustomerBalanceTransactionResourceModel) error {
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
		if rawValueAmount, rawOk := plainValueAtPath(raw, "amount"); rawOk {
			if valueAmount, err := flattenPlainValue(rawValueAmount, types.Int64Type, "amount", "raw response"); err != nil {
				return err
			} else {
				if typedAmount, ok := valueAmount.(types.Int64); ok {
					state.Amount = typedAmount
				}
			}
		} else if !hasRaw {
			if responseValueAmount, ok := plainFromResponseField(obj, "Amount"); ok {
				if valueAmount, err := flattenPlainValue(responseValueAmount, types.Int64Type, "amount", "response struct"); err != nil {
					return err
				} else {
					if typedAmount, ok := valueAmount.(types.Int64); ok {
						state.Amount = typedAmount
					}
				}
			}
		}
	}
	{
		if true {
			if rawValueCheckoutSession, rawOk := plainValueAtPath(raw, "checkout_session"); rawOk {
				if typedCheckoutSession, ok := plainToStringIDValue(rawValueCheckoutSession); ok {
					state.CheckoutSession = typedCheckoutSession
				}
			} else if !hasRaw {
				if responseValueCheckoutSession, ok := plainFromResponseField(obj, "CheckoutSession"); ok {
					if typedCheckoutSession, ok := plainToStringIDValue(responseValueCheckoutSession); ok {
						state.CheckoutSession = typedCheckoutSession
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
		if true {
			if rawValueCreditNote, rawOk := plainValueAtPath(raw, "credit_note"); rawOk {
				if typedCreditNote, ok := plainToStringIDValue(rawValueCreditNote); ok {
					state.CreditNote = typedCreditNote
				}
			} else if !hasRaw {
				if responseValueCreditNote, ok := plainFromResponseField(obj, "CreditNote"); ok {
					if typedCreditNote, ok := plainToStringIDValue(responseValueCreditNote); ok {
						state.CreditNote = typedCreditNote
					}
				}
			}
		}
	}
	{
		if rawValueCurrency, rawOk := plainValueAtPath(raw, "currency"); rawOk {
			if valueCurrency, err := flattenPlainValue(rawValueCurrency, types.StringType, "currency", "raw response"); err != nil {
				return err
			} else {
				if typedCurrency, ok := valueCurrency.(types.String); ok {
					state.Currency = typedCurrency
				}
			}
		} else if !hasRaw {
			if responseValueCurrency, ok := plainFromResponseField(obj, "Currency"); ok {
				if valueCurrency, err := flattenPlainValue(responseValueCurrency, types.StringType, "currency", "response struct"); err != nil {
					return err
				} else {
					if typedCurrency, ok := valueCurrency.(types.String); ok {
						state.Currency = typedCurrency
					}
				}
			}
		}
	}
	{
		if state.Customer.IsNull() || state.Customer.IsUnknown() {
			if rawValueCustomer, rawOk := plainValueAtPath(raw, "customer"); rawOk {
				if typedCustomer, ok := plainToStringIDValue(rawValueCustomer); ok {
					state.Customer = typedCustomer
				}
			} else if !hasRaw {
				if responseValueCustomer, ok := plainFromResponseField(obj, "Customer"); ok {
					if typedCustomer, ok := plainToStringIDValue(responseValueCustomer); ok {
						state.Customer = typedCustomer
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
		if rawValueEndingBalance, rawOk := plainValueAtPath(raw, "ending_balance"); rawOk {
			if valueEndingBalance, err := flattenPlainValue(rawValueEndingBalance, types.Int64Type, "ending_balance", "raw response"); err != nil {
				return err
			} else {
				if typedEndingBalance, ok := valueEndingBalance.(types.Int64); ok {
					state.EndingBalance = typedEndingBalance
				}
			}
		} else if !hasRaw {
			if responseValueEndingBalance, ok := plainFromResponseField(obj, "EndingBalance"); ok {
				if valueEndingBalance, err := flattenPlainValue(responseValueEndingBalance, types.Int64Type, "ending_balance", "response struct"); err != nil {
					return err
				} else {
					if typedEndingBalance, ok := valueEndingBalance.(types.Int64); ok {
						state.EndingBalance = typedEndingBalance
					}
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
		if true {
			if rawValueInvoice, rawOk := plainValueAtPath(raw, "invoice"); rawOk {
				if typedInvoice, ok := plainToStringIDValue(rawValueInvoice); ok {
					state.Invoice = typedInvoice
				}
			} else if !hasRaw {
				if responseValueInvoice, ok := plainFromResponseField(obj, "Invoice"); ok {
					if typedInvoice, ok := plainToStringIDValue(responseValueInvoice); ok {
						state.Invoice = typedInvoice
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
