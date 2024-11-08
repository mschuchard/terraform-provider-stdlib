package numberfunc

import (
	"context"
	"fmt"
	"math"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	util "github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &roundDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewRoundDataSource() datasource.DataSource {
	return &roundDataSource{}
}

// data source implementation
type roundDataSource struct{}

// maps the data source schema data to the model
type roundDataSourceModel struct {
	ID     types.Float64 `tfsdk:"id"`
	Param  types.Float64 `tfsdk:"param"`
	Result types.Int64   `tfsdk:"result"`
}

// data source metadata
func (_ *roundDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_round"
}

// define the provider-level schema for configuration data
func (_ *roundDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDFloat64Attribute(),
			"param": schema.Float64Attribute{
				Description: "Input number parameter for determining the rounding.",
				Required:    true,
			},
			"result": schema.Int64Attribute{
				Computed:    true,
				Description: "Function result storing the rounding of the input parameter.",
			},
		},
		MarkdownDescription: "Return the nearest integer of an input parameter; rounding half away from zero.",
	}
}

// read executes the actual function
func (_ *roundDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state roundDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// initialize input number
	inputNum := state.Param.ValueFloat64()

	// determine the rounded integer
	round := int64(math.Round(inputNum))

	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_round_result", round)
	tflog.Debug(ctx, fmt.Sprintf("Input number parameter \"%f\" rounded is \"%d\"", inputNum, round))

	// store rounded result in state
	state.ID = types.Float64Value(inputNum)
	state.Result = types.Int64Value(round)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined rounded integer of input number", map[string]any{"success": true})
}
