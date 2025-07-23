package util

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

type TestCases map[string]struct {
	Request  function.RunRequest
	Expected function.RunResponse
}

func UnitTests(testCases TestCases, resultData function.ResultData, tfFunction function.Function, test *testing.T) {
	test.Parallel()

	// iterate through test cases
	for name, testCase := range testCases {
		// execute unit test
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: resultData}

			// execute function and store result
			tfFunction.Run(context.Background(), testCase.Request, &result)

			// initialize expected for efficiency
			expected := testCase.Expected

			// compare result versus expected errors
			if !result.Error.Equal(expected.Error) {
				// check for error with function argument
				if expected.Error.FunctionArgument != nil {
					test.Errorf("expected error func arg: %d", *expected.Error.FunctionArgument)
					test.Errorf("actual error func arg: %d", *result.Error.FunctionArgument)
				}
				// display information for error text
				test.Errorf("expected error text: %s", expected.Error.Text)
				test.Errorf("actual error text: %s", result.Error.Text)
			}
			// compare result versus expected values
			if !result.Result.Equal(expected.Result) {
				test.Errorf("expected value: %v", expected.Result.Value())
				test.Errorf("actual value: %v", result.Result.Value())
			}
		})
	}
}
