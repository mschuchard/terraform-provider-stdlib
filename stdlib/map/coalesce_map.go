package mapfunc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &coalesceMapFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewCoalesceMapFunction() function.Function {
	return &coalesceMapFunction{}
}

// function implementation
type coalesceMapFunction struct{}

// function metadata
func (*coalesceMapFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "coalesce_map"
}

// define the provider-level definition for the function
func (*coalesceMapFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Empty coalesce one or more maps",
		MarkdownDescription: "Takes any number of map arguments and returns the first one that is not empty.",
		VariadicParameter: function.MapParameter{
			ElementType: types.StringType,
			Name:        "map",
			Description: "One or more of the input maps to coalesce.",
		},
		Return: function.MapReturn{ElementType: types.MapType{ElemType: types.StringType}},
	}
}

func (*coalesceMapFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize maps to coalesce from input parameters
	var inputMaps []map[string]string

	resp.Error = req.Arguments.Get(ctx, &inputMaps)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "coalesce_map: maps", inputMaps)

	// validate at least one input argument
	if len(inputMaps) == 0 {
		resp.Error = function.NewFuncError("coalesce_map: at least one argument is required")
		return
	}

	// iterate through map
	var returnMap map[string]string
	for _, inputMap := range inputMaps {
		// check if map is not empty
		if len(inputMap) != 0 {
			// assign the first non-empty map to return
			returnMap = inputMap
			break
		}
	}

	// validate at least one non-empty map in arguments
	if returnMap == nil {
		resp.Error = function.NewFuncError("coalesce_map: all arguments are empty maps")
		return
	}

	// store the result as a map
	resp.Error = resp.Result.Set(ctx, &returnMap)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "coalesce_map: result map", returnMap)
	tflog.Debug(ctx, "coalesce_map: successful return", map[string]any{"success": true})
}
