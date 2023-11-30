package multiple

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"golang.org/x/exp/maps"

	"github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &emptyDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewEmptyDataSource() datasource.DataSource {
	return &emptyDataSource{}
}

// data source implementation
type emptyDataSource struct{}

// maps the data source schema data to the model
type emptyDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	ListParam   types.List   `tfsdk:"list_param"`
	MapParam    types.Map    `tfsdk:"map_param"`
	SetParam    types.Set    `tfsdk:"set_param"`
	StringParam types.String `tfsdk:"string_param"`
	Result      types.Bool   `tfsdk:"result"`
}

// data source metadata
func (_ *emptyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_empty"
}

// define the provider-level schema for configuration data
func (_ *emptyDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			"list_param": schema.ListAttribute{
				Description: "List type parameter to check for emptiness. Must be single-level.",
				ElementType: types.StringType,
				Optional:    true,
				// validate that at least one, but no more than one input param is specified
				// this attribute is implicitly included in the following path expressions slice
				Validators: []validator.List{
					listvalidator.ExactlyOneOf(path.Expressions{
						path.MatchRoot("map_param"),
						path.MatchRoot("set_param"),
						path.MatchRoot("string_param"),
					}...),
				},
			},
			"map_param": schema.MapAttribute{
				Description: "Map type parameter to check for emptiness. Must be single-level.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"set_param": schema.SetAttribute{
				Description: "Set type parameter to check for emptiness. Must be single-level.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"string_param": schema.StringAttribute{
				Description: "String type parameter to check for emptiness.",
				Optional:    true,
			},
			"result": schema.BoolAttribute{
				Computed:    true,
				Description: "Function result storing whether input parameter is empty or not.",
			},
		},
		MarkdownDescription: "Return whether the input parameter of one of four possible different types is empty or not.",
	}
}

// read executes the actual function
func (_ *emptyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state emptyDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// initialize default result
	result := false

	// conditionals for which type was input
	if !state.ListParam.IsNull() {
		// convert list param
		var listParam []string
		resp.Diagnostics.Append(state.ListParam.ElementsAs(ctx, &listParam, false)...)
		if resp.Diagnostics.HasError() {
			return
		}

		// check if list param is empty and set state id
		if len(listParam) == 0 {
			result = true
			state.ID = types.StringValue("zero")
		} else {
			state.ID = types.StringValue(listParam[0])
		}
	} else if !state.MapParam.IsNull() {
		// convert map param
		var mapParam map[string]string
		resp.Diagnostics.Append(state.MapParam.ElementsAs(ctx, &mapParam, false)...)
		if resp.Diagnostics.HasError() {
			return
		}

		// check if map is empty and set state id
		if len(mapParam) == 0 {
			result = true
			state.ID = types.StringValue("zero")
		} else {
			state.ID = types.StringValue(maps.Keys(mapParam)[0])
		}
	} else if !state.SetParam.IsNull() {
		// convert set param
		var setParam []string
		resp.Diagnostics.Append(state.SetParam.ElementsAs(ctx, &setParam, false)...)
		if resp.Diagnostics.HasError() {
			return
		}

		// check if set param is empty and set state id
		if len(setParam) == 0 {
			result = true
			state.ID = types.StringValue("zero")
		} else {
			state.ID = types.StringValue(setParam[0])
		}
	} else if !state.StringParam.IsNull() && len(state.StringParam.ValueString()) == 0 {
		// check if string param is empty
		result = true
		// set state id
		state.ID = state.StringParam
	}

	// convert result of emptiness test and assign to model field member
	state.Result = types.BoolValue(result)

	// provide more debug logging
	ctx = tflog.SetField(ctx, "stdlib_empty_result", result)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined whether the input parameter is empty", map[string]any{"success": true})
}
