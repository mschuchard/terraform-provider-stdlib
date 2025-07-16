package slicefunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestLastElementFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.ListUnknown(types.StringType))

	testCases := util.TestCases{
		"single": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("h"), types.StringValue("e"), types.StringValue("l"), types.StringValue("l"), types.StringValue("o")}),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("o")})),
			},
		},
		"three": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("h"), types.StringValue("e"), types.StringValue("l"), types.StringValue("l"), types.StringValue("o")}),
					types.TupleValueMust([]attr.Type{types.Int32Type}, []attr.Value{types.Int32Value(3)}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("l"), types.StringValue("l"), types.StringValue("o")})),
			},
		},
		"num-elem-too-great": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("h"), types.StringValue("e"), types.StringValue("l"), types.StringValue("l"), types.StringValue("o")}),
					types.TupleValueMust([]attr.Type{types.Int32Type}, []attr.Value{types.Int32Value(10)}),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "last_element: the number of terminating elements to return must be fewer than the length of the input list parameter"),
				Result: resultData,
			},
		},
		"num-elem-minimum": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("h"), types.StringValue("e"), types.StringValue("l"), types.StringValue("l"), types.StringValue("o")}),
					types.TupleValueMust([]attr.Type{types.Int32Type}, []attr.Value{types.Int32Value(0)}),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "last_element: number of elements parameter value must be at least 1"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, slicefunc.NewLastElementFunction(), test)
}
