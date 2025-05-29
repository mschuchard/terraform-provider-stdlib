package slicefunc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &lastElementFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewLastElementFunction() function.Function {
	return &lastElementFunction{}
}

// function implementation
type lastElementFunction struct{}

// function metadata
func (*lastElementFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "last_element"
}

// define the provider-level definition for the function
func (*lastElementFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Return the last element of a list",
		MarkdownDescription: "Return one or more terminating element(s) of an input list parameter.",
		Parameters: []function.Parameter{
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "list",
				Description: "Input list parameter for determining the last element(s).",
			},
		},
		VariadicParameter: function.Int32Parameter{
			Name:        "number_of_elements",
			Description: "Optional: The number of terminating elements at the end of the list to return (default: 1). This can be thought of as functionally analogous to a 'reverse slice'.",
		},
		Return: function.ListReturn{ElementType: types.StringType},
	}
}

func (*lastElementFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize list and num elem from input parameters
	var list []string
	var numElementsVar []int
	var numElements int

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &list, &numElementsVar))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "last_element: list", list)
	ctx = tflog.SetField(ctx, "last_element: number of elements variadic", numElementsVar)

	// validation
	if len(numElementsVar) == 0 {
		// assign default numElements value of 1
		numElements = 1
	} else {
		// assign numElements from variadic
		numElements = numElementsVar[0]

		// and then continue validation
		if numElements < 1 {
			resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(1, "last_element: number of elements parameter value must be at least 1"))
		}
	}
	if numElements >= len(list) {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(1, "last_element: the number of terminating elements to return must be fewer than the length of the input list parameter"))
	}
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "last_element: number of elements", numElements)

	// determine last element of slice
	lastElement := list[len(list)-numElements:]

	// store the result as a list of strings
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &lastElement))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "last_element: last element(s)", lastElement)
	tflog.Debug(ctx, "last_element: successful return", map[string]any{"success": true})
}
