package stringfunc

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &cutDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewCutDataSource() datasource.DataSource {
	return &cutDataSource{}
}

// data source implementation
type cutDataSource struct{}

// maps the data source schema data to the model
type cutDataSourceModel struct {
	ID        types.String `tfsdk:"id"`
	Param     types.String `tfsdk:"param"`
	Separator types.String `tfsdk:"separator"`
	Before    types.String `tfsdk:"before"`
	After     types.String `tfsdk:"after"`
	Found     types.Bool   `tfsdk:"found"`
}

// data source metadata
func (*cutDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cut"
}

// define the provider-level schema for configuration data
func (*cutDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			"param": schema.StringAttribute{
				Description: "Input string parameter for cutting around a separator.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"separator": schema.StringAttribute{
				Description: "The separator for cutting the input string.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"before": schema.StringAttribute{
				Computed:    true,
				Description: "Function result storing the input string before the separator.",
			},
			"after": schema.StringAttribute{
				Computed:    true,
				Description: "Function result storing the input string after the separator.",
			},
			"found": schema.BoolAttribute{
				Computed:    true,
				Description: "Function result storing whether the input string contained the separator.",
			},
		},
		MarkdownDescription: "Returns the strings before and after the first instance of the separator in the input string. Also returns whether or not the separator was found in the input string. If the separator is not found in the input string, then `found` will be false, `before` will be equal to `param`, and `after` will be an empty string.",
	}
}

// read executes the actual function
func (*cutDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state cutDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// initialize input string param and separator
	inputString := state.Param.ValueString()
	separator := state.Separator.ValueString()

	// determine string cut
	before, after, found := strings.Cut(inputString, separator)

	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_cut_before", before)
	ctx = tflog.SetField(ctx, "stdlib_cut_after", after)
	ctx = tflog.SetField(ctx, "stdlib_cut_found", found)
	tflog.Debug(ctx, fmt.Sprintf("Input string parameter \"%s\" with separator \"%s\" has before \"%s\", after \"%s\", and found \"%t\"", inputString, separator, before, after, found))

	// store returned values in state
	state.ID = types.StringValue(inputString)
	state.Before = types.StringValue(before)
	state.After = types.StringValue(after)
	state.Found = types.BoolValue(found)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined string cut", map[string]any{"success": true})
}
