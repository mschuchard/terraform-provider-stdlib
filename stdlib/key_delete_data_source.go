package stdlib

import (
  "fmt"
  "context"

  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/hashicorp/terraform-plugin-log/tflog"
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
      "id": schema.StringAttribute{
        Computed:    true,
        Description: "Aliased to name of key to delete for efficiency.",
      },
      "key": schema.StringAttribute{
        Description: "Name of the key to delete from the map.",
        Required:    true,
      },
      "map": schema.MapAttribute{
        Description: "Input map parameter from which to delete a key.",
        Required:    true,
      },
      "result": schema.MapAttribute{
        Computed:    true,
        Description: "Function result storing the map with the key removed.",
      },
    },
  }
}

// read executes the actual function
func (tfData *keyDeleteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

}
