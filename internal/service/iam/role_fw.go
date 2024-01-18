// Code generated by tools/tfsdk2fw/main.go. Manual editing is required.

package iam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"

	fwstringplanmodifier "github.com/hashicorp/terraform-provider-aws/internal/framework/stringplanmodifier"
)

// @FrameworkResource
func newResourceRole(context.Context) (resource.ResourceWithConfigure, error) {
	r := &resourceRole{}
	r.SetMigratedFromPluginSDK(true)

	return r, nil
}

type resourceRole struct {
	framework.ResourceWithConfigure
}

// Metadata should return the full name of the resource, such as
// examplecloud_thing.
func (r *resourceRole) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = "aws_iam_role"
}

// Schema returns the schema for this resource.
func (r *resourceRole) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	s := schema.Schema{
		Attributes: map[string]schema.Attribute{
			"arn": schema.StringAttribute{
				Computed: true,
			},
			"assume_role_policy": schema.StringAttribute{
				Required: true,
				// TODO Validate,
			},
			"create_date": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
				// TODO Validate,
			},
			"force_detach_policies": schema.BoolAttribute{
				Optional: true,
				// TODO Default:false,
			},
			"id": // TODO framework.IDAttribute()
			schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"managed_policy_arns": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
			},
			"max_session_duration": schema.Int64Attribute{
				Optional: true,
				// TODO Default:3600,
				// TODO Validate,
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				// TODO Validate,
			},
			"name_prefix": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				// TODO Validate,
			},
			"path": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					fwstringplanmodifier.DefaultValue("/"),
				},
				// TODO Validate,
			},
			"permissions_boundary": schema.StringAttribute{
				Optional: true,
				// TODO Validate,
			},
			"tags": // TODO tftags.TagsAttribute()
			schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"tags_all": // TODO tftags.TagsAttributeComputedOnly()
			schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
			},
			"unique_id": schema.StringAttribute{
				Computed: true,
			},
		},
		Blocks: map[string]schema.Block{
			"inline_policy": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional: true,
							// TODO Validate,
						},
						"policy": schema.StringAttribute{
							Optional: true,
							// TODO Validate,
						},
					},
				},
			},
		},
	}

	response.Schema = s
}

// Create is called when the provider must create a new resource.
// Config and planned state values should be read from the CreateRequest and new state values set on the CreateResponse.
func (r *resourceRole) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var data resourceRoleData

	response.Diagnostics.Append(request.Plan.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	data.ID = types.StringValue("TODO")

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

// Read is called when the provider must read resource values in order to update state.
// Planned state values should be read from the ReadRequest and new state values set on the ReadResponse.
func (r *resourceRole) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var data resourceRoleData

	response.Diagnostics.Append(request.State.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

// Update is called to update the state of the resource.
// Config, planned state, and prior state values should be read from the UpdateRequest and new state values set on the UpdateResponse.
func (r *resourceRole) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var old, new resourceRoleData

	response.Diagnostics.Append(request.State.Get(ctx, &old)...)

	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.Append(request.Plan.Get(ctx, &new)...)

	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &new)...)
}

// Delete is called when the provider must delete the resource.
// Config values may be read from the DeleteRequest.
//
// If execution completes without error, the framework will automatically call DeleteResponse.State.RemoveResource(),
// so it can be omitted from provider logic.
func (r *resourceRole) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var data resourceRoleData

	response.Diagnostics.Append(request.State.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "deleting TODO", map[string]interface{}{
		"id": data.ID.ValueString(),
	})
}

// ImportState is called when the provider must import the state of a resource instance.
// This method must return enough state so the Read method can properly refresh the full resource.
//
// If setting an attribute with the import identifier, it is recommended to use the ImportStatePassthroughID() call in this method.
func (r *resourceRole) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
}

// ModifyPlan is called when the provider has an opportunity to modify
// the plan: once during the plan phase when Terraform is determining
// the diff that should be shown to the user for approval, and once
// during the apply phase with any unknown values from configuration
// filled in with their final values.
//
// The planned new state is represented by
// ModifyPlanResponse.Plan. It must meet the following
// constraints:
// 1. Any non-Computed attribute set in config must preserve the exact
// config value or return the corresponding attribute value from the
// prior state (ModifyPlanRequest.State).
// 2. Any attribute with a known value must not have its value changed
// in subsequent calls to ModifyPlan or Create/Read/Update.
// 3. Any attribute with an unknown value may either remain unknown
// or take on any value of the expected type.
//
// Any errors will prevent further resource-level plan modifications.
func (r *resourceRole) ModifyPlan(ctx context.Context, request resource.ModifyPlanRequest, response *resource.ModifyPlanResponse) {
	r.SetTagsAll(ctx, request, response)
}

type resourceRoleData struct {
	ARN                 types.String `tfsdk:"arn"`
	AssumeRolePolicy    types.String `tfsdk:"assume_role_policy"`
	CreateDate          types.String `tfsdk:"create_date"`
	Description         types.String `tfsdk:"description"`
	ForceDetachPolicies types.Bool   `tfsdk:"force_detach_policies"`
	ID                  types.String `tfsdk:"id"`
	ManagedPolicyArns   types.Set    `tfsdk:"managed_policy_arns"`
	MaxSessionDuration  types.Int64  `tfsdk:"max_session_duration"`
	Name                types.String `tfsdk:"name"`
	NamePrefix          types.String `tfsdk:"name_prefix"`
	Path                types.String `tfsdk:"path"`
	PermissionsBoundary types.String `tfsdk:"permissions_boundary"`
	Tags                types.Map    `tfsdk:"tags"`
	TagsAll             types.Map    `tfsdk:"tags_all"`
	UniqueID            types.String `tfsdk:"unique_id"`
}
