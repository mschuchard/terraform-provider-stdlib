package mapfunc

import (
	"context"
	"maps"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &flattenMapFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewFlattenMapFunction() function.Function {
	return &flattenMapFunction{}
}

// function implementation
type flattenMapFunction struct{}

// function metadata
func (*flattenMapFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "flatten_map"
}

// define the provider-level definition for the function
func (*flattenMapFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Flatten a list of maps",
		MarkdownDescription: "Return the flattened map of an input list of maps. Note that if a key is repeated between distinct element maps, then the last entry will overwrite any previous entries in the result (maps cannot contain repeated keys).",
		Parameters: []function.Parameter{
			function.ListParameter{
				ElementType: types.MapType{
					ElemType: types.StringType,
				},
				Name:        "list_of_maps",
				Description: "Input list of maps to flatten.",
			},
		},
		Return: function.MapReturn{ElementType: types.StringType},
	}
}

func (*flattenMapFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize comparison maps from input parameters
	var listMaps []map[string]string

	resp.Error = req.Arguments.Get(ctx, &listMaps)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "flatten_map: list of maps", listMaps)

	// validation
	if len(listMaps) < 1 {
		resp.Error = function.NewArgumentFuncError(0, "flatten_map: list of maps parameter must be at least length 1")
		return
	}

	// iterate through list of maps, and merge each map into new map
	outputMap := map[string]string{}
	for _, nestedMap := range listMaps {
		maps.Copy(outputMap, nestedMap)
	}

	// store the result as a map
	resp.Error = resp.Result.Set(ctx, &outputMap)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "flatten_map: result map", outputMap)
	tflog.Debug(ctx, "flatten_map: successful return", map[string]any{"success": true})
}
