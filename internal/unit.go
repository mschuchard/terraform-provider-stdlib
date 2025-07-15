package util

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TestCases map[string]struct {
	Request  function.RunRequest
	Expected function.RunResponse
}

func UnitTests(testCases TestCases, tfFunction function.Function, test *testing.T) {
	test.Parallel()

	// iterate through test cases
	for name, testCase := range testCases {
		// execute unit test
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.MapUnknown(types.StringType))}

			// execute function and store result
			tfFunction.Run(context.Background(), testCase.Request, &result)

			// compare results
			if !result.Error.Equal(testCase.Expected.Error) {
				test.Errorf("expected error: %s", testCase.Expected.Error)
				test.Errorf("actual error: %s", result.Error)
			}
			if !result.Result.Equal(testCase.Expected.Result) {
				test.Errorf("expected value: %v", testCase.Expected.Result.Value())
				test.Errorf("actual value: %v", result.Result.Value())
			}
		})
	}

}
