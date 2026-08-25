package slicefunc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &meanFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewMeanFunction() function.Function {
	return &meanFunction{}
}

// function implementation
type meanFunction struct{}

// function metadata
func (*meanFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "mean"
}

// define the provider-level definition for the function
func (*meanFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Return mean of numbers in a list",
		MarkdownDescription: "Return the mean average of the number elements of an input list parameter.",
		Parameters: []function.Parameter{
			function.ListParameter{
				ElementType: types.Float64Type,
				Name:        "list",
				Description: "Input list parameter for determining the mean.",
			},
		},
		Return: function.Float64Return{},
	}
}

func (*meanFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize list
	var list []float64

	resp.Error = req.Arguments.Get(ctx, &list)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "mean: list", list)

	// validation
	if len(list) == 0 {
		resp.Error = function.NewArgumentFuncError(0, "mean: list parameter length must be at least 1")
		return
	}

	// determine mean average
	var mean float64
	for _, elem := range list {
		mean += elem
	}
	mean /= float64(len(list))

	// store the result as a float
	resp.Error = resp.Result.Set(ctx, &mean)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "mean: mean", mean)
	tflog.Debug(ctx, "mean: successful return", map[string]any{"success": true})
}
