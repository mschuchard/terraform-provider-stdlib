package mapfunc

import (
	"context"
	"maps"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &equalMapFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewEqualMapFunction() function.Function {
	return &equalMapFunction{}
}

// function implementation
type equalMapFunction struct{}

// function metadata
func (*equalMapFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "equal_map"
}

// define the provider-level definition for the function
func (*equalMapFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Check equality of two maps",
		MarkdownDescription: "Return whether the two input map parameters contain the same key-value pairs (equality check). The input maps must be single-level.",
		Parameters: []function.Parameter{
			function.MapParameter{
				ElementType: types.StringType,
				Name:        "map_one",
				Description: "First input map parameter to check for equality with the second.",
			},
			function.MapParameter{
				ElementType: types.StringType,
				Name:        "map_two",
				Description: "Second input map parameter to check for equality with the first.",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (*equalMapFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize comparison maps from input parameters
	var mapOne, mapTwo map[string]string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &mapOne, &mapTwo))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "equal_map: map one", mapOne)
	ctx = tflog.SetField(ctx, "equal_map: map two", mapTwo)

	// check equality of maps and assign to model field member
	result := maps.Equal(mapOne, mapTwo)

	// store the result as a bool
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &result))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "equal_map: result", result)
	tflog.Debug(ctx, "equal_map: successful return", map[string]any{"success": true})
}
