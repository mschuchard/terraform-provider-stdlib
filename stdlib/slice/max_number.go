package slicefunc

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &maxNumberFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewMaxNumberFunction() function.Function {
	return &maxNumberFunction{}
}

// function implementation
type maxNumberFunction struct{}

// function metadata
func (*maxNumberFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "max_number"
}

// define the provider-level definition for the function
func (*maxNumberFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Return maximum number in list",
		MarkdownDescription: "Return the maximum number among the elements of an input list parameter.",
		Parameters: []function.Parameter{
			function.ListParameter{
				ElementType: types.Float64Type,
				Name:        "list",
				Description: "Input list parameter for determining the maximum number.",
			},
		},
		Return: function.Float64Return{},
	}
}

func (*maxNumberFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize list
	var list []float64

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &list))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "max_number: list", list)

	// validation
	if len(list) == 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "max_number: list parameter length must be at least 1"))
	}
	if resp.Error != nil {
		return
	}

	// determine maximum number element of slice
	maxNumber := slices.Max(list)

	// store the result as a float
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &maxNumber))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "max_number: maximum number", maxNumber)
	tflog.Debug(ctx, "max_number: successful return", map[string]any{"success": true})
}
