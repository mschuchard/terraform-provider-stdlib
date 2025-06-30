package slicefunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestRepeatFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"double": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two")}),
					types.Int32Value(2),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("zero"), types.StringValue("one"), types.StringValue("two")})),
			},
		},
		"empty": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two")}),
					types.Int32Value(0),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{})),
			},
		},
		"insert-values-length": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{}),
					types.Int32Value(0),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "repeat: list parameter length must be at least 1"),
				Result: function.NewResultData(types.ListUnknown(types.StringType)),
			},
		},
		"negative-index": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo")}),
					types.Int32Value(-1),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "repeat: count parameter value cannot be negative"),
				Result: function.NewResultData(types.ListUnknown(types.StringType)),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.ListUnknown(types.StringType))}

			// execute function and store result
			slicefunc.NewRepeatFunction().Run(context.Background(), testCase.request, &result)

			// compare results
			if !result.Error.Equal(testCase.expected.Error) {
				test.Errorf("expected error: %s", testCase.expected.Error)
				test.Errorf("actual error: %s", result.Error)
			}
			if !result.Result.Equal(testCase.expected.Result) {
				test.Errorf("expected value: %+q", testCase.expected.Result.Value())
				test.Errorf("actual value: %+q", result.Result.Value())
			}
		})
	}
}
