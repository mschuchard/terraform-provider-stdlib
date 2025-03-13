package slicefunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestMaxStringFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"zero-to-seven": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three"), types.StringValue("four"), types.StringValue("five"), types.StringValue("six"), types.StringValue("seven")}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.StringValue("zero")),
			},
		},
		"greek": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("alpha"), types.StringValue("beta"), types.StringValue("gamma"), types.StringValue("delta"), types.StringValue("epsilon")}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.StringValue("gamma")),
			},
		},
		"list-length": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.Float64Type, []attr.Value{}),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "max_string: list parameter length must be at least 1"),
				Result: function.NewResultData(types.StringUnknown()),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.StringUnknown())}

			// execute function and store result
			slicefunc.NewMaxStringFunction().Run(context.Background(), testCase.request, &result)

			// compare results
			if !result.Error.Equal(testCase.expected.Error) {
				test.Errorf("expected error: %s", testCase.expected.Error)
				test.Errorf("actual error: %s", result.Error)
			}
			if !result.Result.Equal(testCase.expected.Result) {
				test.Errorf("expected value: %s", testCase.expected.Result.Value())
				test.Errorf("actual value: %s", result.Result.Value())
			}
		})
	}
}
