package slicefunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestSortListNumberFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.ListUnknown(types.Float64Type))

	testCases := util.TestCases{
		"integers": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.Float64Type, []attr.Value{types.Float64Value(0), types.Float64Value(4), types.Float64Value(-10), types.Float64Value(8)}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.Float64Type, []attr.Value{types.Float64Value(-10), types.Float64Value(0), types.Float64Value(4), types.Float64Value(8)})),
			},
		},
		"floats": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.Float64Type, []attr.Value{types.Float64Value(9.0), types.Float64Value(45.5), types.Float64Value(123.4), types.Float64Value(-0.5)}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.Float64Type, []attr.Value{types.Float64Value(-0.5), types.Float64Value(9.0), types.Float64Value(45.5), types.Float64Value(123.4)})),
			},
		},
		"minimum-length": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.Float64Type, []attr.Value{types.Float64Value(0)}),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "sort_list_number: list parameter length must be at least 2"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, slicefunc.NewSortListNumberFunction(), test)
}
