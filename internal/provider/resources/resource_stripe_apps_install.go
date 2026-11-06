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

var _ resource.Resource = &AppsInstallResource{}

var _ resource.ResourceWithConfigure = &AppsInstallResource{}

var _ resource.ResourceWithImportState = &AppsInstallResource{}

func NewAppsInstallResource() resource.Resource {
	return &AppsInstallResource{}
}

type AppsInstallResource struct {
	client *stripe.Client
}

type AppsInstallResourceModel struct {
	Object                       types.String `tfsdk:"object"`
	Account                      types.String `tfsdk:"account"`
	App                          types.String `tfsdk:"app"`
	ApprovalRequired             types.Bool   `tfsdk:"approval_required"`
	AuthCode                     types.String `tfsdk:"auth_code"`
	Channel                      types.String `tfsdk:"channel"`
	ContentSecurityPolicyGranted types.Object `tfsdk:"content_security_policy_granted"`
	ContentSecurityPolicyPending types.Object `tfsdk:"content_security_policy_pending"`
	Created                      types.Int64  `tfsdk:"created"`
	CreatedBy                    types.String `tfsdk:"created_by"`
	EndpointsGranted             types.List   `tfsdk:"endpoints_granted"`
	EndpointsPending             types.List   `tfsdk:"endpoints_pending"`
	ID                           types.String `tfsdk:"id"`
	Livemode                     types.Bool   `tfsdk:"livemode"`
	PermissionsGranted           types.List   `tfsdk:"permissions_granted"`
	PermissionsPending           types.List   `tfsdk:"permissions_pending"`
	Status                       types.String `tfsdk:"status"`
	CodeChallenge                types.String `tfsdk:"code_challenge"`
	CodeChallengeMethod          types.String `tfsdk:"code_challenge_method"`
}

func (r *AppsInstallResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AppsInstallResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apps_install"
}

func (r *AppsInstallResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "An app install represents a Stripe App that is installed on an account. It reports the permissions,\ncontent security policy entries, and endpoints that the installing account has authorized, along with any\nthat the app's latest version requests but the account has not authorized yet. Use the Install API to\ninstall, reauthorize, and uninstall apps, and to check the state of existing installs.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("apps.install")},
			},
			"account": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of the account that the app install belongs to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"app": schema.StringAttribute{
				Required:      true,
				Description:   "The ID of the app installed.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"approval_required": schema.BoolAttribute{
				Computed:      true,
				Description:   "Whether the installer must authorize pending permissions, content security policy entries, or endpoints. For private apps, `approval_required` stays `false`; creating or reauthorizing the install through the API installs the newest completed upload and grants its permissions.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"auth_code": schema.StringAttribute{
				Computed:      true,
				Description:   "The authorization code for an oauth app install.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"channel": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The distribution channel associated with the app install.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("private_live", "private_test", "public", "review", "testing")},
			},
			"content_security_policy_granted": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"connect_src": schema.ListAttribute{
						Computed:      true,
						Description:   "The URLs that the app can make network requests to.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"image_src": schema.ListAttribute{
						Computed:      true,
						Description:   "The URLs that the app can load images from.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
				},
			},
			"content_security_policy_pending": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"connect_src": schema.ListAttribute{
						Computed:      true,
						Description:   "The URLs that the app can make network requests to.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"image_src": schema.ListAttribute{
						Computed:      true,
						Description:   "The URLs that the app can load images from.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
				},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"created_by": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of the embedding platform that created the install, if applicable.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"endpoints_granted": schema.ListAttribute{
				Computed:      true,
				Description:   "The endpoint URLs authorized by the installer.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"endpoints_pending": schema.ListAttribute{
				Computed:      true,
				Description:   "The endpoint URLs requested by the latest app version that the installer has not authorized.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
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
			"permissions_granted": schema.ListAttribute{
				Computed:      true,
				Description:   "The permissions authorized by the installer.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"permissions_pending": schema.ListAttribute{
				Computed:      true,
				Description:   "The permissions requested by the latest app version that the installer has not authorized.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "The status of the app install.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("install_failed", "installed", "installing", "uninstall_failed", "uninstalling")},
			},
			"code_challenge": schema.StringAttribute{
				Optional:      true,
				Description:   "For OAuth apps, the PKCE code challenge used to issue the `auth_code` returned on the install. Must be 43 to 128 characters and contain only letters, numbers, `-`, `.`, `_`, and `~`. Only applies to installs made by the app developer or an embedding platform; ignored when an account installs its own private app.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"code_challenge_method": schema.StringAttribute{
				Optional:      true,
				Description:   "The method used to derive `code_challenge`. Required when `code_challenge` is provided, and must be `S256`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
		},
	}
}

func (r *AppsInstallResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AppsInstallResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config AppsInstallResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"CodeChallenge"}, []string{"CodeChallengeMethod"}})

	params, err := expandAppsInstallCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building AppsInstall create params", err.Error())
		return
	}

	obj, err := r.client.V1AppsInstalls.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating AppsInstall", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1AppsInstalls.B, r.client.V1AppsInstalls.Key, stripe.FormatURLPath("/v1/apps/installs/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating AppsInstall create raw response", err.Error())
		return
	}

	if err := flattenAppsInstall(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening AppsInstall create response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"CodeChallenge"}, []string{"CodeChallengeMethod"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *AppsInstallResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState AppsInstallResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state AppsInstallResourceModel
	state = priorState

	obj, err := r.client.V1AppsInstalls.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading AppsInstall", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1AppsInstalls.B, r.client.V1AppsInstalls.Key, stripe.FormatURLPath("/v1/apps/installs/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating AppsInstall raw response", err.Error())
		return
	}

	if err := flattenAppsInstall(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening AppsInstall read response", err.Error())
		return
	}
	clearWriteOnlyPaths(&state, [][]string{[]string{"CodeChallenge"}, []string{"CodeChallengeMethod"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *AppsInstallResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan AppsInstallResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config AppsInstallResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"CodeChallenge"}, []string{"CodeChallengeMethod"}})

	var state AppsInstallResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state
	clearWriteOnlyPaths(&diffPlan, [][]string{[]string{"CodeChallenge"}, []string{"CodeChallengeMethod"}})
	clearWriteOnlyPaths(&diffState, [][]string{[]string{"CodeChallenge"}, []string{"CodeChallengeMethod"}})

	params, err := expandAppsInstallUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building AppsInstall update params", err.Error())
		return
	}

	obj, err := r.client.V1AppsInstalls.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating AppsInstall", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1AppsInstalls.B, r.client.V1AppsInstalls.Key, stripe.FormatURLPath("/v1/apps/installs/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating AppsInstall update raw response", err.Error())
		return
	}

	if err := flattenAppsInstall(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening AppsInstall update response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"CodeChallenge"}, []string{"CodeChallengeMethod"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *AppsInstallResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *AppsInstallResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandAppsInstallCreate(plan AppsInstallResourceModel) (*stripe.AppsInstallCreateParams, error) {
	params := &stripe.AppsInstallCreateParams{}

	if !plan.App.IsNull() && !plan.App.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "App", "App", plan.App.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "app", params)
		}
	}
	if !plan.Channel.IsNull() && !plan.Channel.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Channel", "Channel", plan.Channel.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "channel", params)
		}
	}
	if !plan.CodeChallenge.IsNull() && !plan.CodeChallenge.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CodeChallenge", "CodeChallenge", plan.CodeChallenge.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "code_challenge", params)
		}
	}
	if !plan.CodeChallengeMethod.IsNull() && !plan.CodeChallengeMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CodeChallengeMethod", "CodeChallengeMethod", plan.CodeChallengeMethod.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "code_challenge_method", params)
		}
	}

	return params, nil
}

func expandAppsInstallUpdate(plan AppsInstallResourceModel, state AppsInstallResourceModel) (*stripe.AppsInstallUpdateParams, error) {
	params := &stripe.AppsInstallUpdateParams{}

	return params, nil
}

func flattenAppsInstall(obj *stripe.AppsInstall, state *AppsInstallResourceModel) error {
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
		if rawValueAccount, rawOk := plainValueAtPath(raw, "account"); rawOk {
			if valueAccount, err := flattenPlainValue(rawValueAccount, types.StringType, "account", "raw response"); err != nil {
				return err
			} else {
				if typedAccount, ok := valueAccount.(types.String); ok {
					state.Account = typedAccount
				}
			}
		} else if !hasRaw {
			if responseValueAccount, ok := plainFromResponseField(obj, "Account"); ok {
				if valueAccount, err := flattenPlainValue(responseValueAccount, types.StringType, "account", "response struct"); err != nil {
					return err
				} else {
					if typedAccount, ok := valueAccount.(types.String); ok {
						state.Account = typedAccount
					}
				}
			}
		}
	}
	{
		if rawValueApp, rawOk := plainValueAtPath(raw, "app"); rawOk {
			if valueApp, err := flattenPlainValue(rawValueApp, types.StringType, "app", "raw response"); err != nil {
				return err
			} else {
				if typedApp, ok := valueApp.(types.String); ok {
					state.App = typedApp
				}
			}
		} else if !hasRaw {
			if responseValueApp, ok := plainFromResponseField(obj, "App"); ok {
				if valueApp, err := flattenPlainValue(responseValueApp, types.StringType, "app", "response struct"); err != nil {
					return err
				} else {
					if typedApp, ok := valueApp.(types.String); ok {
						state.App = typedApp
					}
				}
			}
		}
	}
	{
		if rawValueApprovalRequired, rawOk := plainValueAtPath(raw, "approval_required"); rawOk {
			if valueApprovalRequired, err := flattenPlainValue(rawValueApprovalRequired, types.BoolType, "approval_required", "raw response"); err != nil {
				return err
			} else {
				if typedApprovalRequired, ok := valueApprovalRequired.(types.Bool); ok {
					state.ApprovalRequired = typedApprovalRequired
				}
			}
		} else if !hasRaw {
			if responseValueApprovalRequired, ok := plainFromResponseField(obj, "ApprovalRequired"); ok {
				if valueApprovalRequired, err := flattenPlainValue(responseValueApprovalRequired, types.BoolType, "approval_required", "response struct"); err != nil {
					return err
				} else {
					if typedApprovalRequired, ok := valueApprovalRequired.(types.Bool); ok {
						state.ApprovalRequired = typedApprovalRequired
					}
				}
			}
		}
	}
	{
		if rawValueAuthCode, rawOk := plainValueAtPath(raw, "auth_code"); rawOk {
			if valueAuthCode, err := flattenPlainValue(rawValueAuthCode, types.StringType, "auth_code", "raw response"); err != nil {
				return err
			} else {
				if typedAuthCode, ok := valueAuthCode.(types.String); ok {
					state.AuthCode = typedAuthCode
				}
			}
		} else if !hasRaw {
			if responseValueAuthCode, ok := plainFromResponseField(obj, "AuthCode"); ok {
				if valueAuthCode, err := flattenPlainValue(responseValueAuthCode, types.StringType, "auth_code", "response struct"); err != nil {
					return err
				} else {
					if typedAuthCode, ok := valueAuthCode.(types.String); ok {
						state.AuthCode = typedAuthCode
					}
				}
			}
		}
	}
	{
		if rawValueChannel, rawOk := plainValueAtPath(raw, "channel"); rawOk {
			if valueChannel, err := flattenPlainValue(rawValueChannel, types.StringType, "channel", "raw response"); err != nil {
				return err
			} else {
				if typedChannel, ok := valueChannel.(types.String); ok {
					state.Channel = typedChannel
				}
			}
		} else if !hasRaw {
			if responseValueChannel, ok := plainFromResponseField(obj, "Channel"); ok {
				if valueChannel, err := flattenPlainValue(responseValueChannel, types.StringType, "channel", "response struct"); err != nil {
					return err
				} else {
					if typedChannel, ok := valueChannel.(types.String); ok {
						state.Channel = typedChannel
					}
				}
			}
		}
	}
	{
		assignedContentSecurityPolicyGranted := false
		hadRawContentSecurityPolicyGranted := false
		if rawValueContentSecurityPolicyGranted, rawOk := plainValueAtPath(raw, "content_security_policy_granted"); rawOk {
			hadRawContentSecurityPolicyGranted = true
			if rawValueContentSecurityPolicyGranted != nil {
				sourceContentSecurityPolicyGranted := applyConfiguredKeyedListShapes(rawValueContentSecurityPolicyGranted, attrValueToPlain(state.ContentSecurityPolicyGranted))
				if valueContentSecurityPolicyGranted, err := flattenPlainValue(sourceContentSecurityPolicyGranted, types.ObjectType{AttrTypes: map[string]attr.Type{"connect_src": types.ListType{ElemType: types.StringType}, "image_src": types.ListType{ElemType: types.StringType}}}, "content_security_policy_granted", "raw response"); err != nil {
					return err
				} else {
					if typedContentSecurityPolicyGranted, ok := valueContentSecurityPolicyGranted.(types.Object); ok {
						state.ContentSecurityPolicyGranted = typedContentSecurityPolicyGranted
						assignedContentSecurityPolicyGranted = true
					}
				}
			}
		}
		if !assignedContentSecurityPolicyGranted {
			if !hasRaw {
				if responseValueContentSecurityPolicyGranted, ok := plainFromResponseField(obj, "ContentSecurityPolicyGranted"); ok {
					sourceContentSecurityPolicyGranted := applyConfiguredKeyedListShapes(responseValueContentSecurityPolicyGranted, attrValueToPlain(state.ContentSecurityPolicyGranted))
					if valueContentSecurityPolicyGranted, err := flattenPlainValue(
						sourceContentSecurityPolicyGranted,
						types.ObjectType{AttrTypes: map[string]attr.Type{"connect_src": types.ListType{ElemType: types.StringType}, "image_src": types.ListType{ElemType: types.StringType}}},
						"content_security_policy_granted",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedContentSecurityPolicyGranted, ok := valueContentSecurityPolicyGranted.(types.Object); ok {
							state.ContentSecurityPolicyGranted = typedContentSecurityPolicyGranted
							assignedContentSecurityPolicyGranted = true
						}
					}
				}
			}
		}
		if !assignedContentSecurityPolicyGranted && hadRawContentSecurityPolicyGranted {
			if nullContentSecurityPolicyGranted, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"connect_src": types.ListType{ElemType: types.StringType}, "image_src": types.ListType{ElemType: types.StringType}}}); ok {
				if typedContentSecurityPolicyGranted, ok := nullContentSecurityPolicyGranted.(types.Object); ok {
					state.ContentSecurityPolicyGranted = typedContentSecurityPolicyGranted
				}
			}
		}
	}
	{
		assignedContentSecurityPolicyPending := false
		hadRawContentSecurityPolicyPending := false
		if rawValueContentSecurityPolicyPending, rawOk := plainValueAtPath(raw, "content_security_policy_pending"); rawOk {
			hadRawContentSecurityPolicyPending = true
			if rawValueContentSecurityPolicyPending != nil {
				sourceContentSecurityPolicyPending := applyConfiguredKeyedListShapes(rawValueContentSecurityPolicyPending, attrValueToPlain(state.ContentSecurityPolicyPending))
				if valueContentSecurityPolicyPending, err := flattenPlainValue(sourceContentSecurityPolicyPending, types.ObjectType{AttrTypes: map[string]attr.Type{"connect_src": types.ListType{ElemType: types.StringType}, "image_src": types.ListType{ElemType: types.StringType}}}, "content_security_policy_pending", "raw response"); err != nil {
					return err
				} else {
					if typedContentSecurityPolicyPending, ok := valueContentSecurityPolicyPending.(types.Object); ok {
						state.ContentSecurityPolicyPending = typedContentSecurityPolicyPending
						assignedContentSecurityPolicyPending = true
					}
				}
			}
		}
		if !assignedContentSecurityPolicyPending {
			if !hasRaw {
				if responseValueContentSecurityPolicyPending, ok := plainFromResponseField(obj, "ContentSecurityPolicyPending"); ok {
					sourceContentSecurityPolicyPending := applyConfiguredKeyedListShapes(responseValueContentSecurityPolicyPending, attrValueToPlain(state.ContentSecurityPolicyPending))
					if valueContentSecurityPolicyPending, err := flattenPlainValue(
						sourceContentSecurityPolicyPending,
						types.ObjectType{AttrTypes: map[string]attr.Type{"connect_src": types.ListType{ElemType: types.StringType}, "image_src": types.ListType{ElemType: types.StringType}}},
						"content_security_policy_pending",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedContentSecurityPolicyPending, ok := valueContentSecurityPolicyPending.(types.Object); ok {
							state.ContentSecurityPolicyPending = typedContentSecurityPolicyPending
							assignedContentSecurityPolicyPending = true
						}
					}
				}
			}
		}
		if !assignedContentSecurityPolicyPending && hadRawContentSecurityPolicyPending {
			if nullContentSecurityPolicyPending, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"connect_src": types.ListType{ElemType: types.StringType}, "image_src": types.ListType{ElemType: types.StringType}}}); ok {
				if typedContentSecurityPolicyPending, ok := nullContentSecurityPolicyPending.(types.Object); ok {
					state.ContentSecurityPolicyPending = typedContentSecurityPolicyPending
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
		if rawValueCreatedBy, rawOk := plainValueAtPath(raw, "created_by"); rawOk {
			if valueCreatedBy, err := flattenPlainValue(rawValueCreatedBy, types.StringType, "created_by", "raw response"); err != nil {
				return err
			} else {
				if typedCreatedBy, ok := valueCreatedBy.(types.String); ok {
					state.CreatedBy = typedCreatedBy
				}
			}
		} else if !hasRaw {
			if responseValueCreatedBy, ok := plainFromResponseField(obj, "CreatedBy"); ok {
				if valueCreatedBy, err := flattenPlainValue(responseValueCreatedBy, types.StringType, "created_by", "response struct"); err != nil {
					return err
				} else {
					if typedCreatedBy, ok := valueCreatedBy.(types.String); ok {
						state.CreatedBy = typedCreatedBy
					}
				}
			}
		}
	}
	{
		if rawValueEndpointsGranted, rawOk := plainValueAtPath(raw, "endpoints_granted"); rawOk {
			if valueEndpointsGranted, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueEndpointsGranted, attrValueToPlain(state.EndpointsGranted)), types.ListType{ElemType: types.StringType}, "endpoints_granted", "raw response"); err != nil {
				return err
			} else {
				if typedEndpointsGranted, ok := valueEndpointsGranted.(types.List); ok {
					state.EndpointsGranted = typedEndpointsGranted
				}
			}
		} else if !hasRaw {
			if responseValueEndpointsGranted, ok := plainFromResponseField(obj, "EndpointsGranted"); ok {
				if valueEndpointsGranted, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueEndpointsGranted, attrValueToPlain(state.EndpointsGranted)),
					types.ListType{ElemType: types.StringType},
					"endpoints_granted",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedEndpointsGranted, ok := valueEndpointsGranted.(types.List); ok {
						state.EndpointsGranted = typedEndpointsGranted
					}
				}
			}
		}
	}
	{
		if rawValueEndpointsPending, rawOk := plainValueAtPath(raw, "endpoints_pending"); rawOk {
			if valueEndpointsPending, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueEndpointsPending, attrValueToPlain(state.EndpointsPending)), types.ListType{ElemType: types.StringType}, "endpoints_pending", "raw response"); err != nil {
				return err
			} else {
				if typedEndpointsPending, ok := valueEndpointsPending.(types.List); ok {
					state.EndpointsPending = typedEndpointsPending
				}
			}
		} else if !hasRaw {
			if responseValueEndpointsPending, ok := plainFromResponseField(obj, "EndpointsPending"); ok {
				if valueEndpointsPending, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueEndpointsPending, attrValueToPlain(state.EndpointsPending)),
					types.ListType{ElemType: types.StringType},
					"endpoints_pending",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedEndpointsPending, ok := valueEndpointsPending.(types.List); ok {
						state.EndpointsPending = typedEndpointsPending
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
		if rawValuePermissionsGranted, rawOk := plainValueAtPath(raw, "permissions_granted"); rawOk {
			if valuePermissionsGranted, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValuePermissionsGranted, attrValueToPlain(state.PermissionsGranted)), types.ListType{ElemType: types.StringType}, "permissions_granted", "raw response"); err != nil {
				return err
			} else {
				if typedPermissionsGranted, ok := valuePermissionsGranted.(types.List); ok {
					state.PermissionsGranted = typedPermissionsGranted
				}
			}
		} else if !hasRaw {
			if responseValuePermissionsGranted, ok := plainFromResponseField(obj, "PermissionsGranted"); ok {
				if valuePermissionsGranted, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValuePermissionsGranted, attrValueToPlain(state.PermissionsGranted)),
					types.ListType{ElemType: types.StringType},
					"permissions_granted",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedPermissionsGranted, ok := valuePermissionsGranted.(types.List); ok {
						state.PermissionsGranted = typedPermissionsGranted
					}
				}
			}
		}
	}
	{
		if rawValuePermissionsPending, rawOk := plainValueAtPath(raw, "permissions_pending"); rawOk {
			if valuePermissionsPending, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValuePermissionsPending, attrValueToPlain(state.PermissionsPending)), types.ListType{ElemType: types.StringType}, "permissions_pending", "raw response"); err != nil {
				return err
			} else {
				if typedPermissionsPending, ok := valuePermissionsPending.(types.List); ok {
					state.PermissionsPending = typedPermissionsPending
				}
			}
		} else if !hasRaw {
			if responseValuePermissionsPending, ok := plainFromResponseField(obj, "PermissionsPending"); ok {
				if valuePermissionsPending, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValuePermissionsPending, attrValueToPlain(state.PermissionsPending)),
					types.ListType{ElemType: types.StringType},
					"permissions_pending",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedPermissionsPending, ok := valuePermissionsPending.(types.List); ok {
						state.PermissionsPending = typedPermissionsPending
					}
				}
			}
		}
	}
	{
		if rawValueStatus, rawOk := plainValueAtPath(raw, "status"); rawOk {
			if valueStatus, err := flattenPlainValue(rawValueStatus, types.StringType, "status", "raw response"); err != nil {
				return err
			} else {
				if typedStatus, ok := valueStatus.(types.String); ok {
					state.Status = typedStatus
				}
			}
		} else if !hasRaw {
			if responseValueStatus, ok := plainFromResponseField(obj, "Status"); ok {
				if valueStatus, err := flattenPlainValue(responseValueStatus, types.StringType, "status", "response struct"); err != nil {
					return err
				} else {
					if typedStatus, ok := valueStatus.(types.String); ok {
						state.Status = typedStatus
					}
				}
			}
		}
	}
	return nil
}
