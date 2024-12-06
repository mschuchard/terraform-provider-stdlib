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
var _ datasource.DataSource = &modDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewModDataSource() datasource.DataSource {
	return &modDataSource{}
}

// data source implementation
type modDataSource struct{}

// maps the data source schema data to the model
type modDataSourceModel struct {
	ID       types.Float64 `tfsdk:"id"`
	Dividend types.Float64 `tfsdk:"dividend"`
	Divisor  types.Float64 `tfsdk:"divisor"`
	Result   types.Float64 `tfsdk:"result"`
}

// data source metadata
func (*modDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mod"
}

// define the provider-level schema for configuration data
func (*modDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDFloat64Attribute(),
			"dividend": schema.Float64Attribute{
				Description: "The dividend number from which to divide.",
				Required:    true,
			},
			"divisor": schema.Float64Attribute{
				Description: "The divisor number by which to divide.",
				Required:    true,
			},
			"result": schema.Float64Attribute{
				Computed:    true,
				Description: "Function result storing the remainder of the dividend divided by the divisor.",
			},
		},
		MarkdownDescription: "Return the remainder of the dividend divided by the divisor.",
	}
}

// read executes the actual function
func (*modDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state modDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// initialize input numbers
	dividend := state.Dividend.ValueFloat64()
	divisor := state.Divisor.ValueFloat64()

	// determine the remainder
	remainder := math.Mod(dividend, divisor)

	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_mod_result", remainder)
	tflog.Debug(ctx, fmt.Sprintf("Input number dividend \"%f\" divided by input number divisor \"%f\" remainder is \"%f\"", dividend, divisor, remainder))

	// store remainder result in state
	state.ID = types.Float64Value(dividend)
	state.Result = types.Float64Value(remainder)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined remainder of input numbers", map[string]any{"success": true})
}
