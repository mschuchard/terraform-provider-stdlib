package slicefunc

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &medianFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewMedianFunction() function.Function {
	return &medianFunction{}
}

// function implementation
type medianFunction struct{}

// function metadata
func (*medianFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "median"
}

// define the provider-level definition for the function
func (*medianFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Return median of numbers in a list",
		MarkdownDescription: "Return the median average of the number elements of an input list parameter.",
		Parameters: []function.Parameter{
			function.ListParameter{
				ElementType: types.Float64Type,
				Name:        "list",
				Description: "Input list parameter for determining the median.",
			},
		},
		Return: function.Float64Return{},
	}
}

func (*medianFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize list
	var list []float64

	resp.Error = req.Arguments.Get(ctx, &list)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "median: list", list)

	// validation
	if len(list) == 0 {
		resp.Error = function.NewArgumentFuncError(0, "median: list parameter length must be at least 1")
		return
	}

	// determine median
	var median float64
	// sort input list
	slices.Sort(list)
	// determine middle index
	midIndex := len(list) / 2
	// determine if length is even or odd
	if len(list)%2 == 0 {
		// even length so average the two middle index values of the sorted list
		median = (list[midIndex-1] + list[midIndex]) / 2

	} else {
		// odd length so use the middle index value of the sorted list
		median = list[midIndex]
	}

	// store the result as a float
	resp.Error = resp.Result.Set(ctx, &median)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "median: median", median)
	tflog.Debug(ctx, "median: successful return", map[string]any{"success": true})
}
