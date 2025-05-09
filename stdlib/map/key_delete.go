package mapfunc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &keyDeleteFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewKeyDeleteFunction() function.Function {
	return &keyDeleteFunction{}
}

// function implementation
type keyDeleteFunction struct{}

// function metadata
func (*keyDeleteFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "key_delete"
}

// define the provider-level definition for the function
func (*keyDeleteFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Delete a key from a map",
		MarkdownDescription: "Return the input map parameter with the key-value pair corresponding to the key parameter deleted from the map.",
		Parameters: []function.Parameter{
			function.MapParameter{
				ElementType: types.StringType,
				Name:        "map",
				Description: "Input map parameter from which to delete a key.",
			},
			function.StringParameter{
				Name:        "key",
				Description: "Name of the key to delete from the map.",
			},
		},
		Return: function.MapReturn{ElementType: types.StringType},
	}
}

func (*keyDeleteFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize map and key from input parameters
	var inputMap map[string]string
	var key string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &inputMap, &key))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "key_delete: input map", inputMap)
	ctx = tflog.SetField(ctx, "key_delete: key", key)

	// verify key exists in map
	if _, ok := inputMap[key]; ok {
		// delete key from map
		delete(inputMap, key)
	} else {
		// key did not exist in map
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(1, fmt.Sprintf("key_delete: the key to be deleted '%s' does not exist in the input map", key)))
		return
	}

	// store the result as a map
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &inputMap))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "key_delete: result map", inputMap)
	tflog.Debug(ctx, "key_delete: successful return", map[string]any{"success": true})
}
