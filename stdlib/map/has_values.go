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
var _ function.Function = &hasValuesFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewHasValuesFunction() function.Function {
	return &hasValuesFunction{}
}

// function implementation
type hasValuesFunction struct{}

// function metadata
func (*hasValuesFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "has_values"
}

// define the provider-level definition for the function
func (*hasValuesFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Check existence of values in map",
		MarkdownDescription: "Return whether any or all of the input value parameters are present in the input map parameter. The input map must be single-level.",
		Parameters: []function.Parameter{
			function.MapParameter{
				ElementType: types.StringType,
				Name:        "map",
				Description: "Input map parameter from which to check the values' existence.",
			},
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "values",
				Description: "Name of the values to check for existence in the map.",
			},
		},
		VariadicParameter: function.BoolParameter{
			Name:        "all",
			Description: "Optional: Whether to check for all of the values instead of the default any of the values.",
		},
		Return: function.BoolReturn{},
	}
}

func (*hasValuesFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize map and value from input parameters
	var inputMap map[string]string
	var values []string
	var allVar []bool

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &inputMap, &values, &allVar))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "has_values: map", inputMap)
	ctx = tflog.SetField(ctx, "has_values: values", values)
	ctx = tflog.SetField(ctx, "has_values: all", allVar)

	// validate input parameters
	if len(values) < 2 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(1, "has_values: values parameter must be at least length 2"))
	}
	if resp.Error != nil {
		return
	}

	// declare value existence and all vs. any, and then determine all value and expectation
	var valueExists, all bool
	if len(allVar) > 0 {
		// assign all from variadic
		all = allVar[0]

		// assume all or none of the values exist until single check proves otherwise
		valueExists = all
	}

	// assign values of map
	mapValues := slices.Collect(maps.Values(inputMap))
	// iterate through values to check
	for _, value := range values {
		// check input values' existence
		if slices.Contains(mapValues, value) != all {
			// if all is false and single value exists, or all is true and single value does not exist, then flip the existence bool and break
			valueExists = !valueExists
			break
		}
	}

	// store the result as a bool
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &valueExists))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "has_values: existence", valueExists)
	tflog.Debug(ctx, "has_values: successful return", map[string]any{"success": true})
}
