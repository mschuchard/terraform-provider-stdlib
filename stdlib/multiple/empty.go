package multiple

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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
	var result bool

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &parameter))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "empty: input", parameter.String())

	// determine input parameter type and check for emptiness
	// ascertain parameter was not refined to a specific value type
	if parameter.IsUnderlyingValueNull() || parameter.IsUnderlyingValueUnknown() {
		tflog.Error(ctx, fmt.Sprintf("empty: input parameter '%s' was refined by terraform to a specific underlying value type, and this prevents evaluation of the value's emptiness", parameter.String()))
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "empty: underlying value type refined"))
		return
	}
	// access underlying value of dynamic type parameter
	value := parameter.UnderlyingValue()

	// convert to one of four acceptable types
	// string
	if stringType, ok := value.(types.String); ok {
		// emptiness check
		result = len(stringType.ValueString()) == 0
	} else if set, ok := value.Type(ctx).(types.SetType); ok { // set
		// emptiness check
		result = value.Equal(types.SetValueMust(set.ElementType(), []attr.Value{}))
	} else if list, ok := value.Type(ctx).(types.ListType); ok { // list
		// emptiness check
		result = value.Equal(types.ListValueMust(list.ElementType(), []attr.Value{}))
	} else if mapType, ok := value.Type(ctx).(types.MapType); ok { // map
		// emptiness check
		result = value.Equal(types.MapValueMust(mapType.ElementType(), map[string]attr.Value{}))
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
