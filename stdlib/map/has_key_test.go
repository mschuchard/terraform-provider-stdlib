package mapfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	mapfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/map"
)

func TestHasKeyFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.BoolUnknown())

	testCases := util.TestCases{
		"present": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar")}),
					types.StringValue("foo"),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(true)),
			},
		},
		"absent": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar")}),
					types.StringValue("bar"),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(false)),
			},
		},
		"empty-key": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar")}),
					types.StringValue(""),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "has_key: input key parameter must not be empty"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, mapfunc.NewHasKeyFunction(), test)
}
