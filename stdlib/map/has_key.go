package mapfunc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &hasKeyFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewHasKeyFunction() function.Function {
	return &hasKeyFunction{}
}

// function implementation
type hasKeyFunction struct{}

// function metadata
func (*hasKeyFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "has_key"
}

// define the provider-level definition for the function
func (*hasKeyFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Check existence of key in map",
		MarkdownDescription: "Return whether the input key parameter is present in the input map parameter. The input map must be single-level.",
		Parameters: []function.Parameter{
			function.MapParameter{
				ElementType: types.StringType,
				Name:        "map",
				Description: "Input map parameter from which to check a key's existence.",
			},
			function.StringParameter{
				Name:        "key",
				Description: "Name of the key to check for existence in the map.",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (*hasKeyFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize comparison maps from input parameters
	var inputMap map[string]string
	var key string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &inputMap, &key))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "has_key: map", inputMap)
	ctx = tflog.SetField(ctx, "has_key: key", key)

	// check key's existence
	keyExists := false
	if _, ok := inputMap[key]; ok {
		keyExists = true
	}

	// store the result as a bool
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &keyExists))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "has_key: existence", keyExists)
	tflog.Debug(ctx, "has_key: successful return", map[string]any{"success": true})
}
