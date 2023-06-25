package stdlib

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &lastCharDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewLastCharDataSource() datasource.DataSource {
	return &lastCharDataSource{}
}

// data source implementation
type lastCharDataSource struct{}

// maps the data source schema data to the model
type lastCharDataSourceModel struct {
	ID     types.String `tfsdk:"id"`
	Param  types.String `tfsdk:"param"`
	Result types.String `tfsdk:"result"`
}

// data source metadata
func (_ *lastCharDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_last_char"
}

// define the provider-level schema for configuration data
func (_ *lastCharDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": utils.IDStringAttribute(),
			"param": schema.StringAttribute{
				Description: "Input string parameter for determining the last character.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
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
func (_ *lastCharDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input param string value
	var state lastCharDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// initialize input string
	inputString := state.Param.ValueString()

	// re-validate input string
	if len(inputString) == 0 {
		resp.Diagnostics.AddAttributeError(
			path.Root("param"),
			"Empty Value",
			"Expected param value to be non-empty",
		)
		return
	}
	// determine last char of string
	lastCharacter := inputString[len(inputString)-1:]

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
