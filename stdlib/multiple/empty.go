package multiple

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	util "github.com/mschuchard/terraform-provider-stdlib/internal"
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
		MarkdownDescription: "Return whether the input parameter of one of four possible different types (String, Set, List, or Map) is empty or not. Other types will error due to lack of definition for emptiness.",
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

	resp.Error = req.Arguments.Get(ctx, &parameter)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "empty: input", parameter.String())

	// retrieve underlying dynamic value
	paramValue, unknown, null := util.GetDynamicUnderlyingValue(parameter, ctx)
	if unknown || null {
		resp.Error = function.NewArgumentFuncError(0, "empty: input parameter is unknown or null")
		return
	}

	// check if empty
	result, funcErr := util.IsDynamicEmpty(paramValue, ctx)
	if funcErr != nil {
		resp.Error = funcErr
		return
	}

	// store the result as a boolean
	resp.Error = resp.Result.Set(ctx, &result)
	if resp.Error != nil {
		return
	}
	ctx = tflog.SetField(ctx, "empty: result", result)

	tflog.Debug(ctx, "empty: successful return", map[string]any{"success": true})
}
