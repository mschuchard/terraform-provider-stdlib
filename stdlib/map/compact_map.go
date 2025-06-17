package mapfunc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &compactMapFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewCompactMapFunction() function.Function {
	return &compactMapFunction{}
}

// function implementation
type compactMapFunction struct{}

// function metadata
func (*compactMapFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "compact_map"
}

// define the provider-level definition for the function
func (*compactMapFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Compact a map",
		MarkdownDescription: "Return a map with all of the key-value pairs removed where the corresponding value is `null` or empty.",
		Parameters: []function.Parameter{
			function.MapParameter{
				ElementType: types.StringType,
				Name:        "map",
				Description: "Input map to compact.",
			},
		},
		Return: function.MapReturn{ElementType: types.StringType},
	}
}

func (*compactMapFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize comparison maps from input parameters
	var inputMap map[string]string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &inputMap))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "compact_map: map", inputMap)

	// iterate through map
	for key, value := range inputMap {
		// check if value is null or empty
		if len(value) == 0 {
			// delete kv pair if null or empty
			delete(inputMap, key)
		}
	}

	// store the result as a map
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &inputMap))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "compact_map: result map", inputMap)
	tflog.Debug(ctx, "compact_map: successful return", map[string]any{"success": true})
}
