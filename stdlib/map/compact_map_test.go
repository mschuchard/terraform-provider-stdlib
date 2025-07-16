package mapfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	mapfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/map"
)

func TestCompactMapFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.MapUnknown(types.StringType))

	testCases := util.TestCases{
		"standard": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue(""), "bar": types.StringNull()}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world")})),
			},
		},
	}

	util.UnitTests(testCases, resultData, mapfunc.NewCompactMapFunction(), test)
}
