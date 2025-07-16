package slicefunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestInsertFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.ListUnknown(types.StringType))

	testCases := util.TestCases{
		"prepend": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("one"), types.StringValue("two"), types.StringValue("three")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero")}),
					types.Int32Value(0),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three")})),
			},
		},
		"insert": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("four"), types.StringValue("five")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("two"), types.StringValue("three")}),
					types.Int32Value(2),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three"), types.StringValue("four"), types.StringValue("five")})),
			},
		},
		"append": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("three")}),
					types.Int32Value(3),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three")})),
			},
		},
		"insert-values-length": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo")}),
					types.ListValueMust(types.StringType, []attr.Value{}),
					types.Int32Value(0),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "insert: insert values parameter must be at least length 1"),
				Result: resultData,
			},
		},
		"negative-index": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar")}),
					types.Int32Value(-1),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(2, "insert: index parameter must not be a negative number"),
				Result: resultData,
			},
		},
		"out-of-bounds": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar")}),
					types.Int32Value(2),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(2, "insert: index parameter must not be out of range for list"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, slicefunc.NewInsertFunction(), test)
}
