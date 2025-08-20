package slicefunc

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &compareListFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewCompareListFunction() function.Function {
	return &compareListFunction{}
}

// function implementation
type compareListFunction struct{}

// function metadata
func (*compareListFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "compare_list"
}

// define the provider-level definition for the function
func (*compareListFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Compare the elements of two lists similar to the spaceship operator",
		MarkdownDescription: "Returns a comparison between two lists. The elements are compared sequentially, starting at index 0, until one element is not equal to the other. The result of comparing the first non-matching elements is returned. If both lists are equal until one of them ends, then the shorter list is considered less than the longer one. The result is 0 if list_one == list_two, -1 if list_one < list_two, and +1 if list_one > list_two. The input lists must be single-level.",
		Parameters: []function.Parameter{
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "list_one",
				Description: "First input list parameter to compare with the second.",
			},
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "list_two",
				Description: "Second input list parameter to compare with the first.",
			},
		},
		Return: function.Int32Return{},
	}
}

func (*compareListFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize comparison lists from input parameters
	var listOne, listTwo []string

	resp.Error = req.Arguments.Get(ctx, &listOne, &listTwo)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "compare_list: list one", listOne)
	ctx = tflog.SetField(ctx, "compare_list: list two", listTwo)

	// compare lists and assign result to model field member
	result := slices.Compare(listOne, listTwo)

	// store the result as an integer
	resp.Error = resp.Result.Set(ctx, &result)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "compare_list: result", result)
	tflog.Debug(ctx, "compare_list: successful return", map[string]any{"success": true})
}
