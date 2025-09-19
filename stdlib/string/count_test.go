package stringfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	stringfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/string"
)

func TestCountFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.Int32Unknown())

	testCases := util.TestCases{
		"optional-param-absent": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("pizza"), types.StringValue("z")}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(2)),
			},
		},
		"empty-string": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("hello"), types.StringValue("")}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "count: substring parameter must be at least length 1"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, stringfunc.NewCountFunction(), test)
}
