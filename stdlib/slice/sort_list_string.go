package slicefunc

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &sortListStringFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewSortListStringFunction() function.Function {
	return &sortListStringFunction{}
}

// function implementation
type sortListStringFunction struct{}

// function metadata
func (*sortListStringFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "sort_list_string"
}

// define the provider-level definition for the function
func (*sortListStringFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Sort the elements in a list",
		MarkdownDescription: "Return the list where values are sorted in ascending order. Note that the Terraform 'types' package has issues converting some numbers for comparisons such that e.g. 49 will be sorted before 5 due to 4 < 5, but 45 would be correctly sorted before 49. Therefore, it would be preferred to use `sort_list_number` for lists of numbers not implicitly cast or recast to strings.",
		Parameters: []function.Parameter{
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "list",
				Description: "Input list parameter for sorting. This must be at least size 2.",
			},
		},
		Return: function.ListReturn{ElementType: types.StringType},
	}
}

func (*sortListStringFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize list
	var list []string

	resp.Error = req.Arguments.Get(ctx, &list)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "sort_list_string: list", list)

	// validation
	if len(list) < 2 {
		resp.Error = function.NewArgumentFuncError(0, "sort_list_string: list parameter length must be at least 2")
		return
	}

	// sort the list
	slices.Sort(list)

	// store the result as a string
	resp.Error = resp.Result.Set(ctx, &list)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "sort_list_string: sorted list", list)
	tflog.Debug(ctx, "sort_list_string: successful return", map[string]any{"success": true})
}
