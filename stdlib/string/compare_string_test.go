package stringfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	stringfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/string"
)

func TestCompareStringFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.Int32Unknown())

	testCases := util.TestCases{
		"less": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("aaa"), types.StringValue("aab")}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(-1)),
			},
		},
		"less-length": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("aaa"), types.StringValue("aaaa")}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(-1)),
			},
		},
		"equal": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("abc"), types.StringValue("abc")}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(0)),
			},
		},
		"empty": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue(""), types.StringValue("")}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(0)),
			},
		},
		"greater": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("abd"), types.StringValue("abc")}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(1)),
			},
		},
		"greater-length": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("aaaa"), types.StringValue("aaa")}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(1)),
			},
		},
	}

	util.UnitTests(testCases, resultData, stringfunc.NewCompareStringFunction(), test)
}
