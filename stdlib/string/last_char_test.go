package stringfunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stringfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/string"
)

func TestLastCharFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"optional-param-absent": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("hello"), types.Int32Null()}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.StringValue("o")),
			},
		},
		/*"three-terminating-chars": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("hello"), types.Int32Value(3)}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.StringValue("llo")),
			},
		},*/
		"empty-string": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue(""), types.Int32Null()}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "lastChar: input string parameter must be at least length 1"),
				Result: function.NewResultData(types.StringUnknown()),
			},
		},
		/*"zero-num-chars": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("foo"), types.Int32Value(0)}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "lastChar: number_of_characters parameter must be at least 1"),
				Result: function.NewResultData(types.StringUnknown()),
			},
		},
		"num-chars-too-high": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("hello"), types.Int32Value(10)}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "lastChar: number_of_characters parameter must be fewer than the length of the input string parameter"),
				Result: function.NewResultData(types.StringUnknown()),
			},
		},*/
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.StringUnknown())}

			// execute function and store result
			stringfunc.NewLastCharFunction().Run(context.Background(), testCase.request, &result)

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
