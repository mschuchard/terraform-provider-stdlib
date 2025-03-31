package slicefunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestReplaceFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"begin": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar"), types.StringValue("two"), types.StringValue("three")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one")}),
					types.Int32Value(0),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three")})),
			},
		},
		"middle": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("foo"), types.StringValue("bar"), types.StringValue("baz"), types.StringValue("four"), types.StringValue("five")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("one"), types.StringValue("two"), types.StringValue("three")}),
					types.Int32Value(1),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three"), types.StringValue("four"), types.StringValue("five")})),
			},
		},
		"middle-zeroed": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("foo"), types.StringValue("bar"), types.StringValue("four"), types.StringValue("five")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("one")}),
					types.Int32Value(1),
					types.TupleValueMust([]attr.Type{types.Int32Type}, []attr.Value{types.Int32Value(2)}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("four"), types.StringValue("five")})),
			},
		},
		"terminating": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("foo"), types.StringValue("bar"), types.StringValue("baz")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("one"), types.StringValue("two"), types.StringValue("three")}),
					types.Int32Value(1),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three")})),
			},
		},
		"replace-values-length": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo")}),
					types.ListValueMust(types.StringType, []attr.Value{}),
					types.Int32Value(0),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "replace: replace values parameter must be at least length 1"),
				Result: function.NewResultData(types.ListUnknown(types.StringType)),
			},
		},
		"negative-index": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar")}),
					types.Int32Value(-1),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(2, "replace: index parameter must not be a negative number"),
				Result: function.NewResultData(types.ListUnknown(types.StringType)),
			},
		},
		"negative-end-index": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar")}),
					types.Int32Value(0),
					types.TupleValueMust([]attr.Type{types.Int32Type}, []attr.Value{types.Int32Value(-1)}),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(3, "replace: end index parameter must not be a negative number"),
				Result: function.NewResultData(types.ListUnknown(types.StringType)),
			},
		},
		"out-of-bounds": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar"), types.StringValue("two"), types.StringValue("three")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one")}),
					types.Int32Value(3),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(3, "replace: The index at which to replace the values added to the length of the replacement values (i.e. 'endIndex') cannot be greater than the length of the list where the values will be replaced as that would be out of range."),
				Result: function.NewResultData(types.ListUnknown(types.StringType)),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.ListUnknown(types.StringType))}

			// execute function and store result
			slicefunc.NewReplaceFunction().Run(context.Background(), testCase.request, &result)

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
