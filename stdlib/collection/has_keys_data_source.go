package collection

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &hasKeysDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewHasKeysDataSource() datasource.DataSource {
	return &hasKeysDataSource{}
}

// data source implementation
type hasKeysDataSource struct{}

// maps the data source schema data to the model
type hasKeysDataSourceModel struct {
	ID     types.String `tfsdk:"id"`
	All    types.Bool   `tfsdk:"all"`
	Keys   types.List   `tfsdk:"keys"`
	Map    types.Map    `tfsdk:"map"`
	Result types.Bool   `tfsdk:"result"`
}

// data source metadata
func (_ *hasKeysDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_has_keys"
}

// define the provider-level schema for configuration data
func (_ *hasKeysDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			"all": schema.BoolAttribute{
				Description: "Whether to check for all of the keys instead of the default any of the keys.",
				Optional:    true,
			},
			"keys": schema.ListAttribute{
				Description: "Names of the keys to check for existence in the map.",
				Required:    true,
				ElementType: types.StringType,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(2),
				},
			},
			"map": schema.MapAttribute{
				Description: "Input map parameter from which to check a key's existence.",
				ElementType: types.StringType,
				Required:    true,
			},
			"result": schema.BoolAttribute{
				Computed:    true,
				Description: "Function result storing whether the key exists in the map.",
			},
		},
		MarkdownDescription: "Return whether any or all of the input key parameters are present in the input map parameter. The input map must be single-level.",
	}
}

// read executes the actual function
func (_ *hasKeysDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state hasKeysDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// initialize keys, map, and return
	var keysCheck []string
	resp.Diagnostics.Append(state.Keys.ElementsAs(ctx, &keysCheck, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var inputMap map[string]string
	resp.Diagnostics.Append(state.Map.ElementsAs(ctx, &inputMap, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// declare key existence and all vs. any, and then determine all value
	var keyExists, all bool
	if !state.All.IsNull() {
		all = state.All.ValueBool()
		// assume all or none of the keys exist until single check proves otherwise
		keyExists = all
	}

	// iterate through keys to check
	for _, keyCheck := range keysCheck {
		// check input key's existence
		if _, ok := inputMap[keyCheck]; ok != all {
			keyExists = !keyExists
			break
		}
	}

	// provide more debug logging
	ctx = tflog.SetField(ctx, "stdlib_has_keys_result", keyExists)
	tflog.Debug(ctx, fmt.Sprintf("Result of whether key '%s' is in map is: %t", keysCheck, keyExists))

	// store resultant map in state
	state.ID = types.StringValue(keysCheck[0])
	state.Result = types.BoolValue(keyExists)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined whether keys exist in map", map[string]any{"success": true})
}
