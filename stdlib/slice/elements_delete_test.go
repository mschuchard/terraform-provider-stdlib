package slicefunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestElementsDeleteFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.ListUnknown(types.StringType))

	testCases := util.TestCases{
		"begin": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three")}),
					types.Int32Value(0),
					types.Int32Value(1),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("two"), types.StringValue("three")})),
			},
		},
		"end": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three")}),
					types.Int32Value(2),
					types.Int32Value(3),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one")})),
			},
		},
		// this returns two errors, and as far as I can tell this cannot be compared with func (*FuncError) Equal as it involves two structs
		// see also "replace"
		/*"empty-list": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{}),
					types.Int32Value(0),
					types.Int32Value(0),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "elements_delete: list parameter must not be empty"),
				Result: resultData,
			},
		},*/
		"negative-index": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three")}),
					types.Int32Value(-1),
					types.Int32Value(2),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "elements_delete: index parameter must not be a negative number"),
				Result: resultData,
			},
		},
		"end-index-fewer-than-index": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three")}),
					types.Int32Value(2),
					types.Int32Value(1),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(2, "elements_delete: end_index parameter must be greater than or equal to the index parameter"),
				Result: resultData,
			},
		},
		"end-index-out-of-bounds": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three")}),
					types.Int32Value(2),
					types.Int32Value(4),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(2, "elements_delete: end_index parameter must not be greater than or equal to the length of the list"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, slicefunc.NewElementsDeleteFunction(), test)
}
