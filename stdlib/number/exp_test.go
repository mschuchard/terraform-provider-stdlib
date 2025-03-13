package numberfunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	numberfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/number"
)

func TestExpFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"zero-one": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(0)}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(1)),
			},
		},
		"float-one-three": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(1.0986122)}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(2.9999997339956828)),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.Float64Unknown())}

			// execute function and store result
			numberfunc.NewExpFunction().Run(context.Background(), testCase.request, &result)

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
