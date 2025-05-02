package mapfunc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &hasKeysFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewHasKeysFunction() function.Function {
	return &hasKeysFunction{}
}

// function implementation
type hasKeysFunction struct{}

// function metadata
func (*hasKeysFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "has_keys"
}

// define the provider-level definition for the function
func (*hasKeysFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Check existence of keys in map",
		MarkdownDescription: "Return whether any or all of the input key parameters are present in the input map parameter. The input map must be single-level.",
		Parameters: []function.Parameter{
			function.MapParameter{
				ElementType: types.StringType,
				Name:        "map",
				Description: "Input map parameter from which to check the keys' existence.",
			},
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "keys",
				Description: "Names of the keys to check for existence in the map.",
			},
		},
		VariadicParameter: function.BoolParameter{
			Name:        "all",
			Description: "Optional: Whether to check for all of the keys instead of the default any of the keys.",
		},
		Return: function.BoolReturn{},
	}
}

func (*hasKeysFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize map and keys from input parameters
	var inputMap map[string]string
	var keys []string
	var allVar []bool

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &inputMap, &keys, &allVar))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "has_keys: map", inputMap)
	ctx = tflog.SetField(ctx, "has_keys: keys", keys)
	ctx = tflog.SetField(ctx, "has_keys: all", allVar)

	// validate input parameters
	if len(keys) < 2 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(1, "has_keys: keys parameter must be at least length 2"))
	}
	if resp.Error != nil {
		return
	}

	// declare key existence and all vs. any, and then determine all value
	var keyExists, all bool
	if len(allVar) > 0 {
		// assign all from variadic
		all = allVar[0]

		// assume all or none of the keys exist until single check proves otherwise
		keyExists = all
	}

	// iterate through keys to check
	for _, keyCheck := range keys {
		// check input key's existence
		if _, ok := inputMap[keyCheck]; ok != all {
			// if all is false and single key exists, or all is true and single key does not exist, then flip the existence bool and break
			keyExists = !keyExists
			break
		}
	}

	// store the result as a bool
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &keyExists))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "has_keys: existence", keyExists)
	tflog.Debug(ctx, "has_keys: successful return", map[string]any{"success": true})
}
