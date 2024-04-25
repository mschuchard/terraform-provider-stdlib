package slicefunc

import (
	"context"
	"fmt"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &replaceDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewReplaceDataSource() datasource.DataSource {
	return &replaceDataSource{}
}

// data source implementation
type replaceDataSource struct{}

// maps the data source schema data to the model
type replaceDataSourceModel struct {
	ID            types.String `tfsdk:"id"`
	Index         types.Int64  `tfsdk:"index"`
	ReplaceValues types.List   `tfsdk:"replace_values"`
	ListParam     types.List   `tfsdk:"list_param"`
	Result        types.List   `tfsdk:"result"`
}

// data source metadata
func (_ *replaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_replace"
}

// define the provider-level schema for configuration data
func (_ *replaceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			"list_param": schema.ListAttribute{
				Description: "Input list parameter for which the values will be replaced.",
				ElementType: types.StringType,
				Required:    true,
			},
			"replace_values": schema.ListAttribute{
				Description: "Input list of values which will replace values in the list_param.",
				ElementType: types.StringType,
				Required:    true,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"index": schema.Int64Attribute{
				Description: "Index in the list at which to begin replacing the values.",
				Required:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListAttribute{
				Computed:    true,
				Description: "The resulting list with the replaced values.",
				ElementType: types.StringType,
			},
		},
		MarkdownDescription: "Return the list where values are replaced at a specific element index. This function errors if the specified index plus the length of the replace_values list is out of range for the list (greater than length of list_param).",
	}
}

// read executes the actual function
func (_ *replaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state replaceDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// convert tf list to go slice and tf string to go string
	var listParam, replaceValues []string
	resp.Diagnostics.Append(state.ListParam.ElementsAs(ctx, &listParam, false)...)
	resp.Diagnostics.Append(state.ReplaceValues.ElementsAs(ctx, &replaceValues, false)...)
	index := state.Index.ValueInt64()
}
