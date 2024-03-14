package slicefunc

import (
	"context"
	"fmt"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &minStringDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewMinStringDataSource() datasource.DataSource {
	return &minStringDataSource{}
}

// data source implementation
type minStringDataSource struct{}

// maps the data source schema data to the model
type minStringDataSourceModel struct {
	ID     types.String `tfsdk:"id"`
	Param  types.List   `tfsdk:"param"`
	Result types.String `tfsdk:"result"`
}

// data source metadata
func (_ *minStringDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_min_string"
}

// define the provider-level schema for configuration data
func (_ *minStringDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			"param": schema.ListAttribute{
				Description: "Input list parameter for determining the minimum string.",
				ElementType: types.StringType,
				Required:    true,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"result": schema.StringAttribute{
				Computed:    true,
				Description: "Function result storing the minimum string (first by lexical ordering) from the element(s) of the input list.",
			},
		},
		MarkdownDescription: "Return the minimum string (first by lexical ordering) from the elements of an input list parameter.",
	}
}

// read executes the actual function
func (_ *minStringDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state minStringDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// convert tf list to go slice
	var inputList []string
	resp.Diagnostics.Append(state.Param.ElementsAs(ctx, &inputList, false)...)

	// determine minimum string element of slice
	minString := slices.Min(inputList)

	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_min_string_result", minString)
	tflog.Debug(ctx, fmt.Sprintf("Input list parameter \"%v\" min string is \"%s\"", inputList, minString))

	// store minString from element(s) of list in state
	state.ID = types.StringValue(inputList[0])
	state.Result = types.StringValue(minString)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined minimum string from element(s) of list", map[string]any{"success": true})
}
