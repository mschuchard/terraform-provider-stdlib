package slicefunc

import (
	"context"
	"fmt"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &compareListDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewCompareListDataSource() datasource.DataSource {
	return &compareListDataSource{}
}

// data source implementation
type compareListDataSource struct{}

// maps the data source schema data to the model
type compareListDataSourceModel struct {
	ID      types.String `tfsdk:"id"`
	ListOne types.List   `tfsdk:"list_one"`
	ListTwo types.List   `tfsdk:"list_two"`
	Result  types.Int64  `tfsdk:"result"`
}

// data source metadata
func (_ *compareListDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_compare_list"
}

// define the provider-level schema for configuration data
func (_ *compareListDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			"list_one": schema.ListAttribute{
				Description: "First input list parameter to compare with the second.",
				ElementType: types.StringType,
				Required:    true,
			},
			"list_two": schema.ListAttribute{
				Description: "Second input list parameter to compare with the first.",
				ElementType: types.StringType,
				Required:    true,
			},
			"result": schema.Int64Attribute{
				Computed:    true,
				Description: "Function result storing whether the two maps are equal.",
			},
		},
		MarkdownDescription: "Returns a comparison between two lists. The elements are compared sequentially, starting at index 0, until one element is not equal to the other. The result of comparing the first non-matching elements is returned. If both lists are equal until one of them ends, then the shorter list is considered less than the longer one. The result is 0 if list_one == list_two, -1 if list_one < list_two, and +1 if list_one > list_two. The input lists must be single-level",
	}
}

// read executes the actual function
func (_ *compareListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state compareListDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// convert tf lists to go slices
	var listOne, listTwo []string
	resp.Diagnostics.Append(state.ListOne.ElementsAs(ctx, &listOne, false)...)
	resp.Diagnostics.Append(state.ListTwo.ElementsAs(ctx, &listTwo, false)...)

	// compare lists and assign result to model field member
	result := slices.Compare(listOne, listTwo)
	state.Result = types.Int64Value(int64(result))
	// assign id as concatentation of first element of each list
	if len(listOne) > 0 && len(listTwo) > 0 {
		state.ID = types.StringValue(listOne[0] + listTwo[0])
	} else {
		state.ID = types.StringValue("empty")
	}

	// provide more debug logging
	ctx = tflog.SetField(ctx, "stdlib_compare_list_result", result)
	tflog.Debug(ctx, fmt.Sprintf("Result of whether list '%v' equals list '%v' is: %d", listOne, listTwo, result))

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined comparison between two lists", map[string]any{"success": true})
}
