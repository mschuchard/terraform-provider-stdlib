package slicefunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestRepeatFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.ListUnknown(types.StringType))

	testCases := util.TestCases{
		"double": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two")}),
					types.Int32Value(2),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("zero"), types.StringValue("one"), types.StringValue("two")})),
			},
		},
		"empty": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two")}),
					types.Int32Value(0),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{})),
			},
		},
		"insert-values-length": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{}),
					types.Int32Value(0),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "repeat: list parameter length must be at least 1"),
				Result: resultData,
			},
		},
		"negative-index": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo")}),
					types.Int32Value(-1),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "repeat: count parameter value cannot be negative"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, slicefunc.NewRepeatFunction(), test)
}
