package stdlib

import (
  "fmt"
  "context"

  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/hashicorp/terraform-plugin-log/tflog"
  "github.com/hashicorp/terraform-plugin-framework/path"

  "github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var (
  _ datasource.DataSource = &lastCharDataSource{}
)

// helper pseudo-constructor to simplify provider server and testing implementation
func NewLastCharDataSource() datasource.DataSource {
  return &lastCharDataSource{}
}

// data source implementation
type lastCharDataSource struct{}

// maps the data source schema data
type lastCharDataSourceModel struct {
  ID     types.String `tfsdk:"id"`
  Param  types.String `tfsdk:"param"`
  Result types.String `tfsdk:"result"`
}

// data source metadata
func (tfData *lastCharDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_last_char"
}

// define the provider-level schema for configuration data
func (tfData *lastCharDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": utils.IDStringAttribute(),
      "param": schema.StringAttribute{
        Description: "Input string parameter for determining the last character.",
        Required:    true,
      },
      "result": schema.StringAttribute{
        Computed:    true,
        Description: "Function result storing the last character of the input string.",
      },
    },
    MarkdownDescription: "Return the last character of an input string parameter.",
  }
}

// read executes the actual function
func (tfData *lastCharDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
  // determine input param string value
  var state lastCharDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
  inputString := state.Param.ValueString()

  // determine last char of string
  var lastCharacter string
  if len(inputString) > 0 {
    lastCharacter = inputString[len(inputString)-1:]
  } else {
    resp.Diagnostics.AddAttributeError(
      path.Root("param"),
      "Empty Value",
      "Expected param value to be non-empty",
    )
  }

  // provide debug logging
  ctx = tflog.SetField(ctx, "stdlib_last_char_param", inputString)
  ctx = tflog.SetField(ctx, "stdlib_last_char_result", lastCharacter)
  tflog.Debug(ctx, fmt.Sprintf("Input string parameter \"%s\" last character is \"%s\"", inputString, lastCharacter))

  // store lastChar element of list in state
  state.ID = types.StringValue(inputString)
  state.Result = types.StringValue(lastCharacter)

  // set state
  diagnostics := resp.State.Set(ctx, &state)
  resp.Diagnostics.Append(diagnostics...)
  if resp.Diagnostics.HasError() {
    return
  }
  tflog.Info(ctx, "Determined last character of string", map[string]any{"success": true})
}
