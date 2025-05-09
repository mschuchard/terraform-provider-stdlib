package mapfunc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &keysDeleteFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewKeysDeleteFunction() function.Function {
	return &keysDeleteFunction{}
}

// function implementation
type keysDeleteFunction struct{}

// function metadata
func (*keysDeleteFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "keys_delete"
}

// define the provider-level definition for the function
func (*keysDeleteFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Delete keys from a map",
		MarkdownDescription: "Return the input map parameter with the key-value pairs corresponding to the keys parameter deleted from the map.",
		Parameters: []function.Parameter{
			function.MapParameter{
				ElementType: types.StringType,
				Name:        "map",
				Description: "Input map parameter from which to delete the keys.",
			},
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "keys",
				Description: "Names of the keys to delete from the map.",
			},
		},
		Return: function.MapReturn{ElementType: types.StringType},
	}
}

func (*keysDeleteFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize map and keys from input parameters
	var inputMap map[string]string
	var keys []string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &inputMap, &keys))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "keys_delete: input map", inputMap)
	ctx = tflog.SetField(ctx, "keys_delete: keys", keys)

	// validate input parameters
	if len(keys) < 2 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(1, "delete_keys: keys parameter must be at least length 2"))
	}
	if resp.Error != nil {
		return
	}

	// iterate through keys to delete
	for _, key := range keys {
		// verify key exists in map
		if _, ok := inputMap[key]; ok {
			// delete key from map
			delete(inputMap, key)
		} else {
			// key did not exist in map
			resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(1, fmt.Sprintf("keys_delete: the key to be deleted '%s' does not exist in the input map", key)))
			return
		}
	}

	// store the result as a map
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &inputMap))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "keys_delete: result map", inputMap)
	tflog.Debug(ctx, "keys_delete: successful return", map[string]any{"success": true})
}
