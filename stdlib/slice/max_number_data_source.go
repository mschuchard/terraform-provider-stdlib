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
var _ datasource.DataSource = &maxNumberDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewMaxNumberDataSource() datasource.DataSource {
	return &maxNumberDataSource{}
}

// data source implementation
type maxNumberDataSource struct{}

// maps the data source schema data to the model
type maxNumberDataSourceModel struct {
	ID     types.Float64 `tfsdk:"id"`
	Param  types.List    `tfsdk:"param"`
	Result types.Float64 `tfsdk:"result"`
}

// data source metadata
func (_ *maxNumberDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_max_number"
}

// define the provider-level schema for configuration data
func (_ *maxNumberDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDFloat64Attribute(),
			"param": schema.ListAttribute{
				Description: "Input list parameter for determining the maximum number.",
				ElementType: types.Float64Type,
				Required:    true,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"result": schema.Float64Attribute{
				Computed:    true,
				Description: "Function result storing the maximum number from the element(s) of the input list.",
			},
		},
		MarkdownDescription: "Return the maximum number from the elements of an input list parameter.",
	}
}

// read executes the actual function
func (_ *maxNumberDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state maxNumberDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// convert tf list to go slice
	var inputList []float64
	resp.Diagnostics.Append(state.Param.ElementsAs(ctx, &inputList, false)...)

	// determine maximum number element of slice
	maxNumber := slices.Max(inputList)

	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_max_number_result", maxNumber)
	tflog.Debug(ctx, fmt.Sprintf("Input list parameter \"%f\" max number is \"%f\"", inputList, maxNumber))

	// store maxNumber from element(s) of list in state
	state.ID = types.Float64Value(inputList[0])
	state.Result = types.Float64Value(maxNumber)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined maximum number from element(s) of list", map[string]any{"success": true})
}
