package stdlib

import (
  //"fmt"
  "context"

  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/types"
  //"github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-log/tflog"
  //"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
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
  Param  types.Map `tfsdk:"map"`
  Result types.Map `tfsdk:"result"`
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
      "param": schema.ListAttribute{
        Description: "Input map parameter from which to delete a key.",
        // TODO: allow non-strings with interface or generics
        // https://github.com/hashicorp/terraform-plugin-framework/issues/700
        ElementType: types.StringType,
        //ElementType: basetypes.MapType{},
        Required:    true,
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

// read executes the actual function
func (tfData *flattenMapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
  // determine input values
  var state flattenMapDataSourceModel
  resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
  if resp.Diagnostics.HasError() {
    return
  }

  // set state
  diagnostics := resp.State.Set(ctx, &state)
  resp.Diagnostics.Append(diagnostics...)
  if resp.Diagnostics.HasError() {
    return
  }
  tflog.Info(ctx, "Determined flattened map", map[string]any{"success": true})
}
