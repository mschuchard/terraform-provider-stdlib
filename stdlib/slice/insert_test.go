package slicefunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestInsertFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"prepend": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("one"), types.StringValue("two"), types.StringValue("three")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero")}),
					types.Int32Value(0),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three")})),
			},
		},
		"insert": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("four"), types.StringValue("five")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("two"), types.StringValue("three")}),
					types.Int32Value(2),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three"), types.StringValue("four"), types.StringValue("five")})),
			},
		},
		"append": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("three")}),
					types.Int32Value(3),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three")})),
			},
		},
		"insert-values-length": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo")}),
					types.ListValueMust(types.StringType, []attr.Value{}),
					types.Int32Value(0),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "insert: insert values parameter must be at least length 1"),
				Result: function.NewResultData(types.ListUnknown(types.StringType)),
			},
		},
		"negative-index": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar")}),
					types.Int32Value(-1),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(2, "insert: index parameter must not be a negative number"),
				Result: function.NewResultData(types.ListUnknown(types.StringType)),
			},
		},
		"out-of-bounds": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar")}),
					types.Int32Value(2),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(2, "insert: index parameter must not be out of range for list"),
				Result: function.NewResultData(types.ListUnknown(types.StringType)),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.ListUnknown(types.StringType))}

			// execute function and store result
			slicefunc.NewInsertFunction().Run(context.Background(), testCase.request, &result)

			// compare results
			if !result.Error.Equal(testCase.expected.Error) {
				test.Errorf("expected value: %s", testCase.expected.Error)
				test.Errorf("actual value: %s", result.Error)
			}
			if !result.Result.Equal(testCase.expected.Result) {
				test.Errorf("expected value: %+q", testCase.expected.Result.Value())
				test.Errorf("actual value: %+q", result.Result.Value())
			}
		})
	}
}
