package stringfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	stringfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/string"
)

func TestLastIndexFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.Int32Unknown())

	testCases := util.TestCases{
		"repeating-substring": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("terra terraform"), types.StringValue("terra")}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(6)),
			},
		},
		"absent-substring": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("terra terraform"), types.StringValue("vault")}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(-1)),
			},
		},
		"empty-string": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue(""), types.StringValue("terra")}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "last_index: input string parameter must be at least length 1"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, stringfunc.NewLastIndexFunction(), test)
}
