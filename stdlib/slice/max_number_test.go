package slicefunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestMaxNumberFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"fibonacci": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.Float64Type, []attr.Value{types.Float64Value(0), types.Float64Value(1), types.Float64Value(1), types.Float64Value(2), types.Float64Value(3), types.Float64Value(5), types.Float64Value(8), types.Float64Value(13)}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(13)),
			},
		},
		"list-length": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.Float64Type, []attr.Value{}),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "max_number: list parameter length must be at least 1"),
				Result: function.NewResultData(types.Float64Unknown()),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.Float64Unknown())}

			// execute function and store result
			slicefunc.NewMaxNumberFunction().Run(context.Background(), testCase.request, &result)

			// compare results
			if !result.Error.Equal(testCase.expected.Error) {
				test.Errorf("expected error: %s", testCase.expected.Error)
				test.Errorf("actual error: %s", result.Error)
			}
			if !result.Result.Equal(testCase.expected.Result) {
				test.Errorf("expected value: %f", testCase.expected.Result.Value())
				test.Errorf("actual value: %f", result.Result.Value())
			}
		})
	}
}
