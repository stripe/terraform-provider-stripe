//
// File generated from our OpenAPI spec
//

package actions

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/action"
	actionSchema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ action.Action = &ForwardingRequestAction{}
var _ action.ActionWithConfigure = &ForwardingRequestAction{}

func NewForwardingRequestAction() action.Action {
	return &ForwardingRequestAction{}
}

type ForwardingRequestAction struct {
	client *stripe.Client
}

type ForwardingRequestResourceModel struct {
	Metadata      types.Map    `tfsdk:"metadata"`
	PaymentMethod types.String `tfsdk:"payment_method"`
	Replacements  types.List   `tfsdk:"replacements"`
	URL           types.String `tfsdk:"url"`
	Request       types.Object `tfsdk:"request"`
}

func (r *ForwardingRequestAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_forwarding_request"
}

func (r *ForwardingRequestAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = actionSchema.Schema{
		Description: "Instructs Stripe to make a request on your behalf using the destination URL. The destination URL\nis activated by Stripe at the time of onboarding. Stripe verifies requests with your credentials\nprovided during onboarding, and injects card details from the payment_method into the request.\n\nStripe redacts all sensitive fields and headers, including authentication credentials and card numbers,\nbefore storing the request and response data in the forwarding Request object, which are subject to a\n30-day retention period.\n\nYou can provide a Stripe idempotency key to make sure that requests with the same key result in only one\noutbound request. The Stripe idempotency key provided should be unique and different from any idempotency\nkeys provided on the underlying third-party request.\n\nForwarding Requests are synchronous requests that return a response or time out according to\nStripe’s limits.\n\nRelated guide: [Forward card details to third-party API endpoints](https://docs.stripe.com/payments/forwarding).",
		Attributes: map[string]actionSchema.Attribute{
			"metadata": actionSchema.MapAttribute{
				Optional:    true,
				Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				ElementType: types.StringType,
			},
			"payment_method": actionSchema.StringAttribute{
				Required:    true,
				Description: "The PaymentMethod to insert into the forwarded request. Forwarding previously consumed PaymentMethods is allowed.",
			},
			"replacements": actionSchema.ListAttribute{
				Required:    true,
				Description: "The field kinds to be replaced in the forwarded request.",
				ElementType: types.StringType,
			},
			"url": actionSchema.StringAttribute{
				Required:    true,
				Description: "The destination URL for the forwarded request. Must be supported by the config.",
			},
			"request": actionSchema.SingleNestedAttribute{
				Optional:    true,
				Description: "The request body and headers to be sent to the destination endpoint.",
				Attributes: map[string]actionSchema.Attribute{
					"body": actionSchema.StringAttribute{
						Optional:    true,
						Description: "The body payload to send to the destination endpoint.",
					},
					"headers": actionSchema.ListNestedAttribute{
						Optional:    true,
						Description: "The headers to include in the forwarded request. Can be omitted if no additional headers (excluding Stripe-generated ones such as the Content-Type header) should be included.",
						NestedObject: actionSchema.NestedAttributeObject{
							Attributes: map[string]actionSchema.Attribute{
								"name": actionSchema.StringAttribute{
									Required:    true,
									Description: "The header name.",
								},
								"value": actionSchema.StringAttribute{
									Required:    true,
									Description: "The header value.",
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *ForwardingRequestAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*stripe.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Action Configure Type", fmt.Sprintf("Expected *stripe.Client, got: %T", req.ProviderData))
		return
	}

	r.client = client
}

func expandForwardingRequestCreate(plan ForwardingRequestResourceModel) (*stripe.ForwardingRequestCreateParams, error) {
	params := &stripe.ForwardingRequestCreateParams{}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.PaymentMethod.IsNull() && !plan.PaymentMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PaymentMethod", "PaymentMethod", plan.PaymentMethod.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_method", params)
		}
	}
	if !plan.Replacements.IsNull() && !plan.Replacements.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Replacements", plan.Replacements) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "replacements", params)
		}
	}
	if !plan.URL.IsNull() && !plan.URL.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "URL", "URL", plan.URL.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "url", params)
		}
	}
	if !plan.Request.IsNull() && !plan.Request.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Request", plan.Request) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "request", params)
		}
	}

	return params, nil
}

func (r *ForwardingRequestAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config ForwardingRequestResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandForwardingRequestCreate(config)
	if err != nil {
		resp.Diagnostics.AddError("Error building ForwardingRequest action params", err.Error())
		return
	}

	obj, err := r.client.V1ForwardingRequests.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking ForwardingRequest action", err.Error())
		return
	}

	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{Message: fmt.Sprintf("Created stripe_forwarding_request %s", obj.ID)})
	}
}
