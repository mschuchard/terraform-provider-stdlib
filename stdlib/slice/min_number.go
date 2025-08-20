package slicefunc

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &minNumberFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewMinNumberFunction() function.Function {
	return &minNumberFunction{}
}

// function implementation
type minNumberFunction struct{}

// function metadata
func (*minNumberFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "min_number"
}

// define the provider-level definition for the function
func (*minNumberFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Return minimum number in list",
		MarkdownDescription: "Return the minimum number among the elements of an input list parameter.",
		Parameters: []function.Parameter{
			function.ListParameter{
				ElementType: types.Float64Type,
				Name:        "list",
				Description: "Input list parameter for determining the minimum number.",
			},
		},
		Return: function.Float64Return{},
	}
}

func (*minNumberFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize list
	var list []float64

	resp.Error = req.Arguments.Get(ctx, &list)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "min_number: list", list)

	// validation
	if len(list) == 0 {
		resp.Error = function.NewArgumentFuncError(0, "min_number: list parameter length must be at least 1")
		return
	}

	// determine maximum number element of slice
	minNumber := slices.Min(list)

	// store the result as a float
	resp.Error = resp.Result.Set(ctx, &minNumber)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "min_number: minimum number", minNumber)
	tflog.Debug(ctx, "min_number: successful return", map[string]any{"success": true})
}
