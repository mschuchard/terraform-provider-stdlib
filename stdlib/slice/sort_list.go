package slicefunc

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &sortListFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewSortListFunction() function.Function {
	return &sortListFunction{}
}

// function implementation
type sortListFunction struct{}

// function metadata
func (*sortListFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "sort_list"
}

// define the provider-level definition for the function
func (*sortListFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Sort the elements in a list",
		MarkdownDescription: "Return the list where values are sorted in ascending order. Note that the Terraform 'types' package has issues converting some numbers for comparisons such that e.g. 49 will be sorted before 5 due to 4 < 5, but 45 would be correctly sorted before 49.",
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

func (*sortListFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize list
	var list []string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &list))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "sort_list: list", list)

	// validation
	if len(list) < 2 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "sort_list: list parameter length must be at least 2"))
	}
	if resp.Error != nil {
		return
	}

	// shallow clone the param to a result
	sortedList := slices.Clone(list)
	// sort the list
	slices.Sort(sortedList)

	// store the result as a string
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &sortedList))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "sort_list: sorted list", sortedList)
	tflog.Debug(ctx, "sort_list: successful return", map[string]any{"success": true})
}
