package stringfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	stringfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/string"
)

func TestSplitAfterFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.ListUnknown(types.StringType))

	testCases := util.TestCases{
		"normal": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("foo-bar-baz"), types.StringValue("-")}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo-"), types.StringValue("bar-"), types.StringValue("baz")})),
			},
		},
		"absent-separator": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("foo-bar-baz"), types.StringValue("pizza")}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo-bar-baz")})),
			},
		},
		"empty-string": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue(""), types.StringValue("-")}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "split_after: input string parameter must be at least length 1"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, stringfunc.NewSplitAfterFunction(), test)
}
