package mapfunc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	util "github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &keyDeleteDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewKeyDeleteDataSource() datasource.DataSource {
	return &keyDeleteDataSource{}
}

// data source implementation
type keyDeleteDataSource struct{}

// maps the data source schema data to the model
type keyDeleteDataSourceModel struct {
	ID     types.String `tfsdk:"id"`
	Key    types.String `tfsdk:"key"`
	Map    types.Map    `tfsdk:"map"`
	Result types.Map    `tfsdk:"result"`
}

// data source metadata
func (*keyDeleteDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_key_delete"
}

// define the provider-level schema for configuration data
func (*keyDeleteDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			"key": schema.StringAttribute{
				Description: "Name of the key to delete from the map.",
				Required:    true,
			},
			"map": schema.MapAttribute{
				Description: "Input map parameter from which to delete a key.",
				ElementType: types.StringType,
				Required:    true,
			},
			"result": schema.MapAttribute{
				Computed:    true,
				Description: "Function result storing the map with the key removed.",
				ElementType: types.StringType,
			},
		},
		MarkdownDescription: "Return the input map parameter with the key parameter deleted from the map.",
	}
}

// validate data source config
func (*keyDeleteDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
	// determine input values
	var state keyDeleteDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// return if map or key is unknown
	if state.Key.IsUnknown() || state.Map.IsUnknown() {
		return
	}

	// initialize map and key
	deleteKey := state.Key.ValueString()
	var inputMap map[string]string
	resp.Diagnostics.Append(state.Map.ElementsAs(ctx, &inputMap, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// verify key exists in map
	if _, ok := inputMap[deleteKey]; !ok {
		resp.Diagnostics.AddAttributeError(
			path.Root("key"),
			"Improper Attribute Value",
			fmt.Sprintf("The key to be deleted '%s' does not exist in the input map", deleteKey),
		)
	}
}

// read executes the actual function
func (*keyDeleteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state keyDeleteDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// initialize map and key
	deleteKey := state.Key.ValueString()
	var inputMap map[string]string
	resp.Diagnostics.Append(state.Map.ElementsAs(ctx, &inputMap, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// verify key exists in map
	if _, ok := inputMap[deleteKey]; ok {
		// delete key from map
		delete(inputMap, deleteKey)
	} else {
		// key did not exist in map
		resp.Diagnostics.AddAttributeError(
			path.Root("key"),
			"Improper Attribute Value",
			fmt.Sprintf("The key to be deleted '%s' does not exist in the input map", deleteKey),
		)
		return
	}

	// provide more debug logging
	ctx = tflog.SetField(ctx, "stdlib_key_delete_result", inputMap)
	tflog.Debug(ctx, fmt.Sprintf("Map with key removed is \"%v\"", inputMap))

	// store resultant map in state
	state.ID = types.StringValue(deleteKey)
	var mapConvertDiags diag.Diagnostics
	state.Result, mapConvertDiags = types.MapValueFrom(ctx, types.StringType, inputMap)
	resp.Diagnostics.Append(mapConvertDiags...)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined map with key removed", map[string]any{"success": true})
}
