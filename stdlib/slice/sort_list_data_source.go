package slicefunc

import (
	"context"
	"fmt"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &sortListDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewSortListDataSource() datasource.DataSource {
	return &sortListDataSource{}
}

// data source implementation
type sortListDataSource struct{}

// maps the data source schema data to the model
type sortListDataSourceModel struct {
	ID        types.String `tfsdk:"id"`
	ListParam types.List   `tfsdk:"list_param"`
	Result    types.List   `tfsdk:"result"`
}

// data source metadata
func (_ *sortListDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sort_list"
}

// define the provider-level schema for configuration data
func (_ *sortListDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			"list_param": schema.ListAttribute{
				Description: "Input list parameter for sorting. This must be at least size 2.",
				ElementType: types.StringType,
				Required:    true,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(2),
				},
			},
			"result": schema.ListAttribute{
				Computed:    true,
				Description: "The resulting list with the values sorted.",
				ElementType: types.StringType,
			},
		},
		MarkdownDescription: "Return the list where values are sorted in ascending order. Note that the Terraform 'types' package has issues converting some numbers for comparisons such that e.g. 49 will be sorted before 5 due to 4 < 5, but 45 would be correctly sorted before 49.",
	}
}

// read executes the actual function
func (_ *sortListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state sortListDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// convert tf list to go slice
	var listParam []string
	resp.Diagnostics.Append(state.ListParam.ElementsAs(ctx, &listParam, false)...)

	// shallow clone the param to a result
	result := slices.Clone(listParam)
	// sort the list
	slices.Sort(result)

	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_sort_list_result", result)
	tflog.Debug(ctx, fmt.Sprintf("Input list \"%v\" sorted as \"%v\"", listParam, result))

	// store zeroth element of input as id
	state.ID = types.StringValue(listParam[0])
	// store list with values sorted in state
	var listConvertDiags diag.Diagnostics
	state.Result, listConvertDiags = types.ListValueFrom(ctx, types.StringType, result)
	resp.Diagnostics.Append(listConvertDiags...)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined list with values sorted", map[string]any{"success": true})
}
