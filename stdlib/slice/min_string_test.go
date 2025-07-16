package slicefunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestMinStringFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.StringUnknown())

	testCases := util.TestCases{
		"zero-to-seven": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three"), types.StringValue("four"), types.StringValue("five"), types.StringValue("six"), types.StringValue("seven")}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.StringValue("five")),
			},
		},
		"greek": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("alpha"), types.StringValue("beta"), types.StringValue("gamma"), types.StringValue("delta"), types.StringValue("epsilon")}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.StringValue("alpha")),
			},
		},
		"list-length": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.Float64Type, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "min_string: list parameter length must be at least 1"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, slicefunc.NewMinStringFunction(), test)
}
