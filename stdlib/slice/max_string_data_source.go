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
	util "github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &maxStringDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewMaxStringDataSource() datasource.DataSource {
	return &maxStringDataSource{}
}

// data source implementation
type maxStringDataSource struct{}

// maps the data source schema data to the model
type maxStringDataSourceModel struct {
	ID     types.String `tfsdk:"id"`
	Param  types.List   `tfsdk:"param"`
	Result types.String `tfsdk:"result"`
}

// data source metadata
func (*maxStringDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_max_string"
}

// define the provider-level schema for configuration data
func (*maxStringDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			"param": schema.ListAttribute{
				Description: "Input list parameter for determining the maximum string.",
				ElementType: types.StringType,
				Required:    true,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"result": schema.StringAttribute{
				Computed:    true,
				Description: "Function result storing the maximum string (last by lexical ordering) from the element(s) of the input list.",
			},
		},
		MarkdownDescription: "Return the maximum string (last by lexical ordering) from the elements of an input list parameter.",
	}
}

// read executes the actual function
func (*maxStringDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state maxStringDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// convert tf list to go slice
	var inputList []string
	resp.Diagnostics.Append(state.Param.ElementsAs(ctx, &inputList, false)...)

	// determine maximum string element of slice
	maxString := slices.Max(inputList)

	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_max_string_result", maxString)
	tflog.Debug(ctx, fmt.Sprintf("Input list parameter \"%v\" max string is \"%s\"", inputList, maxString))

	// store maxString from element(s) of list in state
	state.ID = types.StringValue(inputList[0])
	state.Result = types.StringValue(maxString)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined maximum string from element(s) of list", map[string]any{"success": true})
}
