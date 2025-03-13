package numberfunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	numberfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/number"
)

func TestRoundFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"round-down": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(1.2)}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(1)),
			},
		},
		"round-up": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(1.8)}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(2)),
			},
		},
		"round-half": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(1.5)}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(2)),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.Int64Unknown())}

			// execute function and store result
			numberfunc.NewRoundFunction().Run(context.Background(), testCase.request, &result)

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
