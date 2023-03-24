package stdlib

import (
  "fmt"
  "context"

  "golang.org/x/exp/maps"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var (
  _ datasource.DataSource = &flattenMapDataSource{}
)

// helper pseudo-constructor to simplify provider server and testing implementation
func NewFlattenMapDataSource() datasource.DataSource {
  return &flattenMapDataSource{}
}

// data source implementation
type flattenMapDataSource struct{}

// maps the data source schema data
type flattenMapDataSourceModel struct {
  ID     types.String `tfsdk:"id"`
  Param  []paramModel `tfsdk:"param"`
  Result types.Map `tfsdk:"result"`
}

// maps param schema data
type paramModel struct {
  Map types.Map `tfsdk:"map"`
}

// data source metadata
func (tfData *flattenMapDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_flatten_map"
}

// define the provider-level schema for configuration data
func (tfData *flattenMapDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Computed:    true,
        Description: "Aliased to name of first key in map for efficiency.",
      },
      // TODO: also support set
      "param": schema.ListNestedAttribute{
        Description: "Input list of maps to flatten.",
        Required:    true,
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "map": schema.MapAttribute{
              Computed:    true,
              Description: "The map elements of the input list.",
              // TODO: allow non-strings with interface or generics
              ElementType: types.StringType,
            },
          },
        },
      },
      "result": schema.MapAttribute{
        Computed:    true,
        Description: "Function result storing the flattened map.",
        // TODO: allow non-strings with interface or generics
        ElementType: types.StringType,
      },
    },
  }
}

// TODO: need to revisit when plugin framework supports list(map) in the schema
// read executes the actual function
func (tfData *flattenMapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
  // determine input values
  var state flattenMapDataSourceModel
  resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
  if resp.Diagnostics.HasError() {
    return
  }

  // TODO: allow non-strings with interface or generics
  // initialize input list of maps, nested maps, and output map
  var nestedMap map[string]string
  var outputMap map[string]string

  // iterate through list of maps and merge the maps into new map
  for _, nestedMaps := range state.Param {
    nestedMaps.Map.ElementsAs(ctx, &nestedMap, false)
    maps.Copy(outputMap, nestedMap)
  }
  // provide debug logging
  ctx = tflog.SetField(ctx, "stdlib_flatten_map_param", state.Param)
  ctx = tflog.SetField(ctx, "stdlib_flatten_map_result", outputMap)
  tflog.Debug(ctx, fmt.Sprintf("Flattened map is \"%v\"", outputMap))

  // store first key of output map in input list as id
  state.ID = types.StringValue(maps.Keys(outputMap)[0])
  // TODO: allow non-strings with interface or generics
  var mapConvertDiags diag.Diagnostics
  state.Result, mapConvertDiags = types.MapValueFrom(ctx, types.StringType, outputMap)
  resp.Diagnostics.Append(mapConvertDiags...)

  // set state
  diagnostics := resp.State.Set(ctx, &state)
  resp.Diagnostics.Append(diagnostics...)
  if resp.Diagnostics.HasError() {
    return
  }
  tflog.Info(ctx, "Determined flattened map", map[string]any{"success": true})
}
