package stdlib

import (
  "fmt"
  "context"

  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-log/tflog"

  "github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var (
  _ datasource.DataSource = &keyDeleteDataSource{}
)

// helper pseudo-constructor to simplify provider server and testing implementation
func NewKeyDeleteDataSource() datasource.DataSource {
  return &keyDeleteDataSource{}
}

// data source implementation
type keyDeleteDataSource struct{}

// maps the data source schema data
type keyDeleteDataSourceModel struct {
  ID     types.String `tfsdk:"id"`
  Key    types.String `tfsdk:"key"`
  Map    types.Map `tfsdk:"map"`
  Result types.Map `tfsdk:"result"`
}

// data source metadata
func (tfData *keyDeleteDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_key_delete"
}

// define the provider-level schema for configuration data
func (tfData *keyDeleteDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": utils.IDStringAttribute(),
      "key": schema.StringAttribute{
        Description: "Name of the key to delete from the map.",
        Required:    true,
      },
      "map": schema.MapAttribute{
        Description: "Input map parameter from which to delete a key.",
        // TODO: allow non-strings with interface or generics
        ElementType: types.StringType,
        Required:    true,
      },
      "result": schema.MapAttribute{
        Computed:    true,
        Description: "Function result storing the map with the key removed.",
        // TODO: allow non-strings with interface or generics
        ElementType: types.StringType,
      },
    },
  }
}

// read executes the actual function
func (tfData *keyDeleteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
  // determine input values
  var state keyDeleteDataSourceModel
  resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
  if resp.Diagnostics.HasError() {
    return
  }

  // initialize map and key
  deleteKey := state.Key.ValueString()
  // TODO: allow non-strings with interface or generics
  var inputMap map[string]string
  resp.Diagnostics.Append(state.Map.ElementsAs(ctx, &inputMap, false)...)
  if resp.Diagnostics.HasError() {
		return
	}

  // provide debug logging
  ctx = tflog.SetField(ctx, "stdlib_key_delete_key", deleteKey)
  ctx = tflog.SetField(ctx, "stdlib_key_delete_map", inputMap)
  tflog.Debug(ctx, fmt.Sprintf("Input map parameter \"%v\" with key parameter \"%s\" removed", inputMap, deleteKey))

  // delete key from map
  _, ok := inputMap[deleteKey]; if ok {
    delete(inputMap, deleteKey)
  } else {
    tflog.Error(ctx, fmt.Sprintf("The key to be deleted '%s' does not exist in the input map", deleteKey))
  }

  // provide more debug logging
  ctx = tflog.SetField(ctx, "stdlib_key_delete_result", inputMap)
  tflog.Debug(ctx, fmt.Sprintf("Map with key removed is \"%v\"", inputMap))

  // store resultant map in state
  state.ID = types.StringValue(deleteKey)
  // TODO: allow non-strings with interface or generics
  var mapConvertDiags diag.Diagnostics
  state.Result, mapConvertDiags = types.MapValueFrom(ctx, types.StringType, inputMap)
  resp.Diagnostics.Append(mapConvertDiags...)

  // set state
  diagnostics := resp.State.Set(ctx, &state)
  resp.Diagnostics.Append(diagnostics...)
  if resp.Diagnostics.HasError() {
    return
  }
  tflog.Info(ctx, "Determined map with key removed", map[string]any{"success": true})
}
