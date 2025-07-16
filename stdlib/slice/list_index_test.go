package slicefunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestListIndexFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.Int32Unknown())

	testCases := util.TestCases{
		"one": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two")}),
					types.StringValue("one"),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(1)),
			},
		},
		"sorted": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("a"), types.StringValue("b"), types.StringValue("c"), types.StringValue("d")}),
					types.StringValue("c"),
					types.TupleValueMust([]attr.Type{types.BoolType}, []attr.Value{types.BoolValue(true)}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(2)),
			},
		},
		"repeated": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three"), types.StringValue("two"), types.StringValue("one"), types.StringValue("zero")}),
					types.StringValue("two"),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(2)),
			},
		},
		"absent": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("hundred"), types.StringValue("thousand"), types.StringValue("million"), types.StringValue("billion")}),
					types.StringValue("infinity"),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(-1)),
			},
		},
		"list-length": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{}),
					types.StringValue("foo"),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "list_index: list parameter length must be at least 1"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, slicefunc.NewListIndexFunction(), test)
}
