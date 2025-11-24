package stringfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	stringfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/string"
)

func TestLastCharFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.StringUnknown())

	testCases := util.TestCases{
		"optional-param-absent": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("hello"), types.TupleValueMust([]attr.Type{}, []attr.Value{})}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.StringValue("o")),
			},
		},
		"three-terminating-chars": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("hello"), types.TupleValueMust([]attr.Type{types.Int32Type}, []attr.Value{types.Int32Value(3)})}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.StringValue("llo")),
			},
		},
		"empty-string": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue(""), types.TupleValueMust([]attr.Type{}, []attr.Value{})}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "last_char: input string parameter must be at least length 1"),
				Result: resultData,
			},
		},
		"zero-num-chars": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("foo"), types.TupleValueMust([]attr.Type{types.Int32Type}, []attr.Value{types.Int32Value(0)})}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "last_char: number_of_characters parameter must be at least 1"),
				Result: resultData,
			},
		},
		"num-chars-too-high": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("hello"), types.TupleValueMust([]attr.Type{types.Int32Type}, []attr.Value{types.Int32Value(10)})}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "last_char: number_of_characters parameter must be less than or equal to the length of the input string parameter"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, stringfunc.NewLastCharFunction(), test)
}
