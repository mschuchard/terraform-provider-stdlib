package multiple_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	"github.com/mschuchard/terraform-provider-stdlib/stdlib/multiple"
)

func TestRepeatFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.DynamicUnknown())

	testCases := util.TestCases{
		"list-double": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.DynamicValue(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two")})),
					types.Int32Value(2),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.DynamicValue(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("zero"), types.StringValue("one"), types.StringValue("two")}))),
			},
		},
		// string double
		"list-empty": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.DynamicValue(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two")})),
					types.Int32Value(0),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.DynamicValue(types.ListValueMust(types.StringType, []attr.Value{}))),
			},
		},
		// string empty
		"repeater-list-length": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.DynamicValue(types.ListValueMust(types.StringType, []attr.Value{})),
					types.Int32Value(0),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "repeat: repeater parameter length must be at least 1"),
				Result: resultData,
			},
		},
		"repeater-string-length": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.DynamicValue(types.StringValue("")),
					types.Int32Value(0),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "repeat: repeater parameter length must be at least 1"),
				Result: resultData,
			},
		},
		"count-negative": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.DynamicValue(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo")})),
					types.Int32Value(-1),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "repeat: count parameter value cannot be negative"),
				Result: resultData,
			},
		},
		"invalid-type": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.DynamicValue(types.SetValueMust(types.StringType, []attr.Value{types.StringValue("foo")})),
					types.Int32Value(2),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "repeat: invalid input parameter type"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, multiple.NewRepeatFunction(), test)
}
