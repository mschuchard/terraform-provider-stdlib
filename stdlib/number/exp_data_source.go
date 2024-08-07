package numberfunc

import (
	"context"
	"fmt"
	"math"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &expDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewExpDataSource() datasource.DataSource {
	return &expDataSource{}
}

// data source implementation
type expDataSource struct{}

// maps the data source schema data to the model
type expDataSourceModel struct {
	ID     types.Float64 `tfsdk:"id"`
	Param  types.Float64 `tfsdk:"param"`
	Result types.Float64 `tfsdk:"result"`
}

// data source metadata
func (_ *expDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_exp"
}

// define the provider-level schema for configuration data
func (_ *expDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDFloat64Attribute(),
			"param": schema.Float64Attribute{
				Description: "Input number parameter for determining the base-e exponential.",
				Required:    true,
			},
			"result": schema.Float64Attribute{
				Computed:    true,
				Description: "Function result storing the base-e exponential of the input parameter.",
			},
		},
		MarkdownDescription: "Return the base-e exponential of an inpurt parameter.",
	}
}

// read executes the actual function
func (_ *expDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state expDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// initialize input number
	inputNum := state.Param.ValueFloat64()

	// determine the base e exponential
	exponential := math.Exp(inputNum)

	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_exp_result", exponential)
	tflog.Debug(ctx, fmt.Sprintf("Input number parameter \"%f\" base e exponential is \"%f\"", inputNum, exponential))

	// store exponential result in state
	state.ID = types.Float64Value(inputNum)
	state.Result = types.Float64Value(exponential)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined base e exponential of input number", map[string]any{"success": true})
}
