package mapfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	mapfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/map"
)

func TestEqualMapFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.BoolUnknown())

	testCases := util.TestCases{
		"equal": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world")}),
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world")}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(true)),
			},
		},
		"unequal": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world")}),
					types.MapValueMust(types.StringType, map[string]attr.Value{"foo": types.StringValue("bar")}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(false)),
			},
		},
	}

	util.UnitTests(testCases, resultData, mapfunc.NewEqualMapFunction(), test)
}
