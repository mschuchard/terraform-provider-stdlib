package multiple_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mschuchard/terraform-provider-stdlib/stdlib/multiple"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
)

func TestEmptyFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.BoolUnknown())

	testCases := util.TestCases{
		"string": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.DynamicValue(types.StringValue(""))}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(true)),
			},
		},
		"set": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.DynamicValue(types.SetValueMust(types.StringType, []attr.Value{types.StringValue("no")}))}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(false)),
			},
		},
		"list": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.DynamicValue(types.ListValueMust(types.StringType, []attr.Value{}))}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(true)),
			},
		},
		"map": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.DynamicValue(types.MapValueMust(types.StringType, map[string]attr.Value{"foo": types.StringValue("bar")}))}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(false)),
			},
		},
		"invalid type": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.DynamicValue(types.BoolValue(false))}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "IsDynamicEntry (helper): invalid input parameter type"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, multiple.NewEmptyFunction(), test)
}
