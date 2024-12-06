package slicefunc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &productDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewProductDataSource() datasource.DataSource {
	return &productDataSource{}
}

// data source implementation
type productDataSource struct{}

// maps the data source schema data to the model
type productDataSourceModel struct {
	ID       types.Float64 `tfsdk:"id"`
	SetParam types.Set     `tfsdk:"set_param"`
	Result   types.Float64 `tfsdk:"result"`
}

// data source metadata
func (*productDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_product"
}

// define the provider-level schema for configuration data
func (*productDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDFloat64Attribute(),
			"set_param": schema.SetAttribute{
				Description: "Input set parameter for determining the product. The set must contain at least one element.",
				ElementType: types.Float64Type,
				Required:    true,
				Validators: []validator.Set{
					setvalidator.SizeAtLeast(1),
				},
			},
			"result": schema.Float64Attribute{
				Computed:    true,
				Description: "The resulting list with the values sorted.",
			},
		},
		MarkdownDescription: "Return the product of the elements within a set.",
	}
}

// read executes the actual function
func (*productDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state productDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// convert tf set to go slice
	var setParam []float64
	resp.Diagnostics.Append(state.SetParam.ElementsAs(ctx, &setParam, false)...)

	// determine the product
	result := 1.0
	for _, elem := range setParam {
		result *= elem
	}

	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_product_result", result)
	tflog.Debug(ctx, fmt.Sprintf("Input set \"%f\" product is \"%f\"", setParam, result))

	// store product of set in state
	state.ID = types.Float64Value(setParam[0])
	state.Result = types.Float64Value(result)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined product of set", map[string]any{"success": true})
}
