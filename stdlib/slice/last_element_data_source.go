package slicefunc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	util "github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &lastElementDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewLastElementDataSource() datasource.DataSource {
	return &lastElementDataSource{}
}

// data source implementation
type lastElementDataSource struct{}

// maps the data source schema data to the model
type lastElementDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Param       types.List   `tfsdk:"param"`
	NumElements types.Int64  `tfsdk:"num_elements"`
	Result      types.List   `tfsdk:"result"`
}

// data source metadata
func (*lastElementDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_last_element"
}

// define the provider-level schema for configuration data
func (*lastElementDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			"param": schema.ListAttribute{
				Description: "Input list parameter for determining the last element(s).",
				ElementType: types.StringType,
				Required:    true,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"num_elements": schema.Int64Attribute{
				Description: "The number of terminating elements at the end of the list to return (default: 1). This can be thought of as functionally analogous to a 'reverse slice'.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"result": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Function result storing the list containing the last element(s) of the input list.",
			},
		},
		MarkdownDescription: "Return the last element(s) of an input list parameter.",
	}
}

// read executes the actual function
func (*lastElementDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state lastElementDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// convert tf list to go slice, and initialize number of terminating elems
	var inputList []string
	resp.Diagnostics.Append(state.Param.ElementsAs(ctx, &inputList, false)...)
	numElements := 1

	// validate num_elements if input, and re-assign from default value
	if !state.NumElements.IsNull() {
		// re-assign
		numElements = int(state.NumElements.ValueInt64())

		// number of terminating elements must be fewer than length of input list
		if numElements >= len(inputList) {
			resp.Diagnostics.AddAttributeError(
				path.Root("num_elements"),
				"Invalid Value",
				"The number of terminating elements to return must be fewer than the length of the input list parameter.",
			)
			return
		}
	}
	// determine last element of slice
	lastElement := inputList[len(inputList)-numElements:]

	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_last_element_result", lastElement)
	tflog.Debug(ctx, fmt.Sprintf("Input list parameter \"%v\" last element(s) is \"%v\"", inputList, lastElement))

	// store lastElement element(s) of list in state
	state.ID = types.StringValue(inputList[0])
	var resultConvertDiags diag.Diagnostics
	state.Result, resultConvertDiags = types.ListValueFrom(ctx, types.StringType, lastElement)

	// append diagnostics
	resp.Diagnostics.Append(resultConvertDiags...)
	if resultConvertDiags.HasError() {
		return
	}

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined last element(s) of list", map[string]any{"success": true})
}
