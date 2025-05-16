package multiple_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mschuchard/terraform-provider-stdlib/stdlib/multiple"
)

func TestEmptyFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"string": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.DynamicValue(types.StringValue(""))}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(true)),
			},
		},
		"set": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.DynamicValue(types.SetValueMust(types.StringType, []attr.Value{types.StringValue("no")}))}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(false)),
			},
		},
		"list": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.DynamicValue(types.ListValueMust(types.StringType, []attr.Value{}))}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(true)),
			},
		},
		"map": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.DynamicValue(types.MapValueMust(types.StringType, map[string]attr.Value{"foo": types.StringValue("bar")}))}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(false)),
			},
		},
		"invalid type": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.DynamicValue(types.BoolValue(false))}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "empty: invalid input parameter type"),
				Result: function.NewResultData(types.BoolUnknown()),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.BoolUnknown())}

			// execute function and store result
			multiple.NewEmptyFunction().Run(context.Background(), testCase.request, &result)

			// compare results
			if !result.Error.Equal(testCase.expected.Error) {
				test.Errorf("%s failed", name)
				test.Errorf("expected error: %s", testCase.expected.Error)
				test.Errorf("actual error: %s", result.Error)
			}
			if !result.Result.Equal(testCase.expected.Result) {
				test.Errorf("%s failed", name)
				test.Errorf("expected value: %+v", testCase.expected.Result.Value())
				test.Errorf("actual value: %+v", result.Result.Value())
			}
		})
	}
}
