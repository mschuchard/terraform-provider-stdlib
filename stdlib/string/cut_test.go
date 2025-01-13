package stringfunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stringfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/string"
)

func TestCutFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"normal": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("foobarbaz"), types.StringValue("bar")}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.TupleValueMust(
					[]attr.Type{types.StringType, types.StringType, types.BoolType},
					[]attr.Value{types.StringValue("foo"), types.StringValue("baz"), types.BoolValue(true)})),
			},
		},
		"separator-absent": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("foobarbaz"), types.StringValue("pizza")}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.TupleValueMust(
					[]attr.Type{types.StringType, types.StringType, types.BoolType},
					[]attr.Value{types.StringValue("foobarbaz"), types.StringValue(""), types.BoolValue(false)})),
			},
		},
		"empty-string": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue(""), types.StringValue("foo")}),
			},
			expected: function.RunResponse{
				Error: function.NewArgumentFuncError(0, "cut: input string parameter must be at least length 1"),
				Result: function.NewResultData(types.TupleUnknown(
					[]attr.Type{types.StringType, types.StringType, types.BoolType})),
			},
		},
		"empty-separator": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("foo"), types.StringValue("")}),
			},
			expected: function.RunResponse{
				Error: function.NewArgumentFuncError(1, "cut: separator parameter must be at least length 1"),
				Result: function.NewResultData(types.TupleUnknown(
					[]attr.Type{types.StringType, types.StringType, types.BoolType})),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			result := function.RunResponse{Result: function.NewResultData(types.TupleUnknown([]attr.Type{types.StringType, types.StringType, types.BoolType}))}

			stringfunc.NewCutFunction().Run(context.Background(), testCase.request, &result)

			//if result.Error.Equal(testCase.expected.Error) {
			if testCase.expected.Error != nil && result.Error.Error() != testCase.expected.Error.Text {
				test.Errorf("expected value: %s", testCase.expected.Error)
				test.Errorf("actual value: %s", result.Error)
			}
			//if result.Result.Equal(testCase.expected.Result) {
			/* if result.Result.Value().Equal(testCase.expected.Result.Value()) {
				test.Errorf("expected value: %+q", testCase.expected.Result.Value())
				test.Errorf("actual value: %+q", result.Result.Value())
			}*/
		})
	}
}
