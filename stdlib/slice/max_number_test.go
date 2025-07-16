package slicefunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestMaxNumberFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.Float64Unknown())

	testCases := util.TestCases{
		"fibonacci": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.Float64Type, []attr.Value{types.Float64Value(0), types.Float64Value(1), types.Float64Value(1), types.Float64Value(2), types.Float64Value(3), types.Float64Value(5), types.Float64Value(8), types.Float64Value(13)}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(13)),
			},
		},
		"list-length": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.Float64Type, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "max_number: list parameter length must be at least 1"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, slicefunc.NewMaxNumberFunction(), test)
}
