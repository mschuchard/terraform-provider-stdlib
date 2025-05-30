package mapfunc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	util "github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &keysDeleteDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewKeysDeleteDataSource() datasource.DataSource {
	return &keysDeleteDataSource{}
}

// data source implementation
type keysDeleteDataSource struct{}

// maps the data source schema data to the model
type keysDeleteDataSourceModel struct {
	ID     types.String `tfsdk:"id"`
	Keys   types.List   `tfsdk:"keys"`
	Map    types.Map    `tfsdk:"map"`
	Result types.Map    `tfsdk:"result"`
}

// data source metadata
func (*keysDeleteDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_keys_delete"
}

// define the provider-level schema for configuration data
func (*keysDeleteDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			"keys": schema.ListAttribute{
				Description: "Names of the keys to delete from the map.",
				Required:    true,
				ElementType: types.StringType,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(2),
				},
			},
			"map": schema.MapAttribute{
				Description: "Input map parameter from which to delete the keys.",
				ElementType: types.StringType,
				Required:    true,
			},
			"result": schema.MapAttribute{
				Computed:    true,
				Description: "Function result storing the map with the keys removed.",
				ElementType: types.StringType,
			},
		},
		MarkdownDescription: "Return the input map parameter with the keys parameter deleted from the map.",
	}
}

// validate data source config
func (*keysDeleteDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
	// determine input values
	var state keysDeleteDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// return if map or key is unknown
	if state.Keys.IsUnknown() || state.Map.IsUnknown() {
		return
	}

	// initialize keys and map
	var deleteKeys []string
	resp.Diagnostics.Append(state.Keys.ElementsAs(ctx, &deleteKeys, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var inputMap map[string]string
	resp.Diagnostics.Append(state.Map.ElementsAs(ctx, &inputMap, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// iterate through keys to delete
	for _, deleteKey := range deleteKeys {
		// verify key exists in map
		_, ok := inputMap[deleteKey]
		if !ok {
			resp.Diagnostics.AddAttributeError(
				path.Root("key"),
				"Improper Attribute Value",
				fmt.Sprintf("The key to be deleted '%s' does not exist in the input map", deleteKey),
			)
		}
	}
}

// read executes the actual function
func (*keysDeleteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state keysDeleteDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// initialize keys and map
	var deleteKeys []string
	resp.Diagnostics.Append(state.Keys.ElementsAs(ctx, &deleteKeys, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var inputMap map[string]string
	resp.Diagnostics.Append(state.Map.ElementsAs(ctx, &inputMap, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// iterate through keys to delete
	for _, deleteKey := range deleteKeys {
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
	}

	// provide more debug logging
	ctx = tflog.SetField(ctx, "stdlib_keys_delete_result", inputMap)
	tflog.Debug(ctx, fmt.Sprintf("Map with key removed is \"%v\"", inputMap))

	// store resultant map in state
	state.ID = types.StringValue(deleteKeys[0])
	var mapConvertDiags diag.Diagnostics
	state.Result, mapConvertDiags = types.MapValueFrom(ctx, types.StringType, inputMap)
	resp.Diagnostics.Append(mapConvertDiags...)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined map with keys removed", map[string]any{"success": true})
}
