package mapfunc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &transposeFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewTransposeFunction() function.Function {
	return &transposeFunction{}
}

// function implementation
type transposeFunction struct{}

// function metadata
func (*transposeFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "transpose"
}

// define the provider-level definition for the function
func (*transposeFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Transpose keys and values of a single-level map of strings.",
		MarkdownDescription: "Return a map of key value pairs where the keys and values are swapped. The input map must be single-level and contain string type values. This is analogous to the core function `transpose`, but for `map(string)` instead of `map(list(string))`. If a value is repeated in the input map, then the return will contain only the last occurrence. A warning is emitted in this situation.",
		Parameters: []function.Parameter{
			function.MapParameter{
				ElementType: types.StringType,
				Name:        "map",
				Description: "Input map parameter to transpose.",
			},
		},
		Return: function.MapReturn{ElementType: types.StringType},
	}
}

func (*transposeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize map and key from input parameters
	var inputMap map[string]string

	resp.Error = req.Arguments.Get(ctx, &inputMap)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "transpose: map", inputMap)

	// allocate a new map to hold the transposed key-value pairs
	transposedMap := make(map[string]string, len(inputMap))
	// iterate through the input map and transpose keys and values
	for key, value := range inputMap {
		// validate that the value is not already present in the transposed map
		if _, exists := transposedMap[value]; exists {
			tflog.Warn(ctx, "transpose: duplicate value found in input map; only the last occurrence will exist in the transposed map", map[string]any{"duplicate_value": value})
		}
		// store the transposed key-value pair in the new map
		transposedMap[value] = key
	}

	// store the result as map of strings
	resp.Error = resp.Result.Set(ctx, &transposedMap)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "transpose: transposed map", transposedMap)
	tflog.Debug(ctx, "transpose: successful return", map[string]any{"success": true})
}
