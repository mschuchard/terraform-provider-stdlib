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
var _ datasource.DataSource = &minNumberDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewMinNumberDataSource() datasource.DataSource {
	return &minNumberDataSource{}
}

// data source implementation
type minNumberDataSource struct{}

// maps the data source schema data to the model
type minNumberDataSourceModel struct {
	ID     types.Float64 `tfsdk:"id"`
	Param  types.List    `tfsdk:"param"`
	Result types.Float64 `tfsdk:"result"`
}

// data source metadata
func (_ *minNumberDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_min_number"
}

// define the provider-level schema for configuration data
func (_ *minNumberDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDFloat64Attribute(),
			"param": schema.ListAttribute{
				Description: "Input list parameter for determining the minimum number.",
				ElementType: types.Float64Type,
				Required:    true,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"result": schema.Float64Attribute{
				Computed:    true,
				Description: "Function result storing the minimum number from the element(s) of the input list.",
			},
		},
		MarkdownDescription: "Return the minimum number from the elements of an input list parameter.",
	}
}

// read executes the actual function
func (_ *minNumberDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state minNumberDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// convert tf list to go slice
	var inputList []float64
	resp.Diagnostics.Append(state.Param.ElementsAs(ctx, &inputList, false)...)

	// determine minimum number element of slice
	minNumber := slices.Min(inputList)

	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_min_number_result", minNumber)
	tflog.Debug(ctx, fmt.Sprintf("Input list parameter \"%f\" min number is \"%f\"", inputList, minNumber))

	// store minNumber from element(s) of list in state
	state.ID = types.Float64Value(inputList[0])
	state.Result = types.Float64Value(minNumber)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined minimum number from element(s) of list", map[string]any{"success": true})
}
