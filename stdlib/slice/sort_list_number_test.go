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
	resultData := function.NewResultData(types.ListUnknown(types.Float32Type))

	testCases := util.TestCases{
		"integers": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.Float32Type, []attr.Value{types.Float32Value(0), types.Float32Value(4), types.Float32Value(-10), types.Float32Value(8)}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.Float32Type, []attr.Value{types.Float32Value(-10), types.Float32Value(0), types.Float32Value(4), types.Float32Value(8)})),
			},
		},
		"floats": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.Float32Type, []attr.Value{types.Float32Value(9.0), types.Float32Value(45.5), types.Float32Value(123.4), types.Float32Value(-0.5)}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.Float32Type, []attr.Value{types.Float32Value(-0.5), types.Float32Value(9.0), types.Float32Value(45.5), types.Float32Value(123.4)})),
			},
		},
		"minimum-length": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.Float32Type, []attr.Value{types.Float32Value(0)}),
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
