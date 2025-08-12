package mapfunc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	util "github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &compactMapExperimentalFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewCompactMapExperimentalFunction() function.Function {
	return &compactMapExperimentalFunction{}
}

// function implementation
type compactMapExperimentalFunction struct{}

// function metadata
func (*compactMapExperimentalFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "compact_map"
}

// define the provider-level definition for the function
func (*compactMapExperimentalFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Compact a map (experimental and not yet included in provider plugin).",
		MarkdownDescription: "Return a map with all of the key-value pairs removed where the corresponding value is `null` or empty. The types checked for emptiness are String, Set, List, and Map. Other types will error due to lack of definition for emptiness. Note this function is unsupported in the current version of the Terraform Plugin Framework due to explicit schema enforcement, and represents future functionality once it is supported (it currently behaves as expected according to unit test cases). As such this function is not currently included in the provider plugin.",
		Parameters: []function.Parameter{
			function.MapParameter{
				ElementType: types.DynamicType,
				Name:        "map",
				Description: "Input map to compact.",
			},
		},
		Return: function.MapReturn{ElementType: types.DynamicType},
	}
}

func (*compactMapExperimentalFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize map to compact from input parameters
	var inputMap map[string]types.Dynamic

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &inputMap))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "compact_map: map", inputMap)

	// iterate through map
	for key, value := range inputMap {
		// check if value is null
		if value.IsUnderlyingValueNull() {
			// delete kv pair if null
			delete(inputMap, key)
		} else if empty, funcErr := util.IsDynamicEmpty(value, ctx); empty || funcErr != nil { // check if value is empty
			// check on error during emptiness check
			if funcErr != nil {
				resp.Error = function.ConcatFuncErrors(resp.Error, funcErr)
				return
			} else {
				// delete kv pair if empty
				delete(inputMap, key)
			}
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
