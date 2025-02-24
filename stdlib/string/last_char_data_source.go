package stringfunc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
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
	ID       types.String `tfsdk:"id"`
	Param    types.String `tfsdk:"param"`
	NumChars types.Int64  `tfsdk:"num_chars"`
	Result   types.String `tfsdk:"result"`
}

// data source metadata
func (*lastCharDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_last_char"
}

// define the provider-level schema for configuration data
func (*lastCharDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			"param": schema.StringAttribute{
				Description: "Input string parameter for determining the last character.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"num_chars": schema.Int64Attribute{
				Description: "The number of terminating characters at the end of the string to return (default: 1).",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"result": schema.StringAttribute{
				Computed:    true,
				Description: "Function result storing the last character of the input string.",
			},
		},
		MarkdownDescription: "Return the last character(s) of an input string parameter.",
	}
}

// read executes the actual function
func (*lastCharDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state lastCharDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// initialize input string and number of terminating chars
	inputString := state.Param.ValueString()
	numChars := 1

	// validate num_chars if input and re-assign from default value
	if !state.NumChars.IsNull() {
		// re-assign
		numChars = int(state.NumChars.ValueInt64())

		// number of terminating chars must be fewer than length of input string
		if numChars >= len(inputString) {
			resp.Diagnostics.AddAttributeError(
				path.Root("num_chars"),
				"Invalid Value",
				"The number of terminating characters to return must be fewer than the length of the input string parameter",
			)
			return
		}
	}
	// determine last char of string
	lastCharacter := inputString[len(inputString)-numChars:]

	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_last_char_result", lastCharacter)
	tflog.Debug(ctx, fmt.Sprintf("Input string parameter \"%s\" last character is \"%s\"", inputString, lastCharacter))

	// store lastChar element of list in state
	state.ID = types.StringValue(inputString)
	state.Result = types.StringValue(lastCharacter)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined last character of string", map[string]any{"success": true})
}
