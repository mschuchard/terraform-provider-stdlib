package mapfunc

import (
	"context"
	"maps"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &hasValueFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewHasValueFunction() function.Function {
	return &hasValueFunction{}
}

// function implementation
type hasValueFunction struct{}

// function metadata
func (*hasValueFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "has_value"
}

// define the provider-level definition for the function
func (*hasValueFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Check existence of value in map",
		MarkdownDescription: "Return whether the input value parameter is present in the input map parameter. The input map must be single-level.",
		Parameters: []function.Parameter{
			function.MapParameter{
				ElementType: types.StringType,
				Name:        "map",
				Description: "Input map parameter from which to check a value's existence.",
			},
			function.StringParameter{
				Name:        "value",
				Description: "Name of the value to check for existence in the map.",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (*hasValueFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize map and value from input parameters
	var inputMap map[string]string
	var value string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &inputMap, &value))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "has_value: map", inputMap)
	ctx = tflog.SetField(ctx, "has_value: value", value)

	// validate input parameters
	if len(value) == 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(1, "has_value: input value parameter must not be empty"))
		return
	}

	// assign values of map and check input value's existence
	mapValues := slices.Collect(maps.Values(inputMap))
	valueExists := slices.Contains(mapValues, value)

	// store the result as a bool
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &valueExists))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "has_value: existence", valueExists)
	tflog.Debug(ctx, "has_value: successful return", map[string]any{"success": true})
}
