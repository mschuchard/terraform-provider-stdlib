package stringfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	stringfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/string"
)

func TestCutFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.TupleUnknown([]attr.Type{types.StringType, types.StringType, types.BoolType}))

	testCases := util.TestCases{
		"normal": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("foobarbaz"), types.StringValue("bar")}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.TupleValueMust(
					[]attr.Type{types.StringType, types.StringType, types.BoolType},
					[]attr.Value{types.StringValue("foo"), types.StringValue("baz"), types.BoolValue(true)},
				)),
			},
		},
		"separator-absent": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("foobarbaz"), types.StringValue("pizza")}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.TupleValueMust(
					[]attr.Type{types.StringType, types.StringType, types.BoolType},
					[]attr.Value{types.StringValue("foobarbaz"), types.StringValue(""), types.BoolValue(false)},
				)),
			},
		},
		"empty-string": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue(""), types.StringValue("foo")}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "cut: input string parameter must be at least length 1"),
				Result: resultData,
			},
		},
		"empty-separator": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("foo"), types.StringValue("")}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "cut: separator parameter must be at least length 1"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, stringfunc.NewCutFunction(), test)
}
