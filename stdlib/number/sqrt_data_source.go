package numberfunc

import (
	"context"
	"fmt"
	"math"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &sqrtDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewSqrtDataSource() datasource.DataSource {
	return &sqrtDataSource{}
}

// data source implementation
type sqrtDataSource struct{}

// maps the data source schema data to the model
type sqrtDataSourceModel struct {
	ID     types.Float64 `tfsdk:"id"`
	Param  types.Float64 `tfsdk:"param"`
	Result types.Float64 `tfsdk:"result"`
}

// data source metadata
func (_ *sqrtDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sqrt"
}

// define the provider-level schema for configuration data
func (_ *sqrtDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDFloat64Attribute(),
			"param": schema.Float64Attribute{
				Description: "Input number parameter for determining the square root.",
				Required:    true,
			},
			"result": schema.Float64Attribute{
				Computed:    true,
				Description: "Function result storing the square root of the input parameter.",
			},
		},
		MarkdownDescription: "Return the square root of an input parameter;.",
	}
}

// read executes the actual function
func (_ *sqrtDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state sqrtDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// initialize input number
	inputNum := state.Param.ValueFloat64()

	// determine the square root
	sqrt := math.Sqrt(inputNum)
	if math.IsNaN(sqrt) {
		resp.Diagnostics.AddAttributeError(
			path.Root("param"),
			"Invalid Value",
			"The square root of the input parameter must return a valid number, but instead returned 'NaN'.",
		)
		return
	}

	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_sqrt_result", sqrt)
	tflog.Debug(ctx, fmt.Sprintf("Input number parameter \"%f\" square root is \"%f\"", inputNum, sqrt))

	// store sqrted result in state
	state.ID = types.Float64Value(inputNum)
	state.Result = types.Float64Value(sqrt)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined square root of input number", map[string]any{"success": true})
}
