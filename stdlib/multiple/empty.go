package multiple

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &emptyFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewEmptyFunction() function.Function {
	return &emptyFunction{}
}

// function implementation
type emptyFunction struct{}

// function metadata
func (*emptyFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "empty"
}

// define the provider-level definition for the function
func (*emptyFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Determine if input is empty",
		MarkdownDescription: "Return whether the input parameter of one of four possible different types (String, Set, List, or Map) is empty or not.",
		Parameters: []function.Parameter{
			function.DynamicParameter{
				Name:        "input",
				Description: "Input parameter to check for emptiness.",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (*emptyFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input param, converted go types, and result
	var parameter types.Dynamic
	var stringConvert string
	var result bool

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &parameter))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "empty: input", parameter.String())

	// determine input parameter type and check for emptiness
	// access terraform value of dynamic type parameter
	tfValue, err := parameter.ToTerraformValue(ctx)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("empty: could not convert input parameter '%s' to an acceptable terraform value", parameter.String()))
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}

	// convert to one of four acceptable types
	// string
	if err = tfValue.As(&stringConvert); err == nil {
		// emptiness check
		result = len(stringConvert) == 0
	} else if parameter.String()[:1] == "[" && parameter.String()[len(parameter.String())-1:] == "]" { // janky set or list
		// emptiness check
		result = parameter.String() == "[]"
	} else if parameter.String()[:1] == "{" && parameter.String()[len(parameter.String())-1:] == "}" { // janky map
		// emptiness check
		result = parameter.String() == "{}"
	} else {
		tflog.Error(ctx, fmt.Sprintf("empty: could not convert input parameter '%s' to an acceptable terraform type", parameter.String()))
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "empty: invalid input parameter type"))
		return
	}

	// store the result as a boolean
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &result))
	if resp.Error != nil {
		return
	}
	ctx = tflog.SetField(ctx, "empty: result", result)

	tflog.Debug(ctx, "empty: successful return", map[string]any{"success": true})
}
