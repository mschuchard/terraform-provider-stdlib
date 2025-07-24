package slicefunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestReplaceFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.ListUnknown(types.StringType))

	testCases := util.TestCases{
		"begin": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar"), types.StringValue("two"), types.StringValue("three")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one")}),
					types.Int32Value(0),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three")})),
			},
		},
		"middle": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("foo"), types.StringValue("bar"), types.StringValue("baz"), types.StringValue("four"), types.StringValue("five")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("one"), types.StringValue("two"), types.StringValue("three")}),
					types.Int32Value(1),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three"), types.StringValue("four"), types.StringValue("five")})),
			},
		},
		"middle-zeroed": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("foo"), types.StringValue("bar"), types.StringValue("four"), types.StringValue("five")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("one")}),
					types.Int32Value(1),
					types.TupleValueMust([]attr.Type{types.Int32Type}, []attr.Value{types.Int32Value(2)}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("four"), types.StringValue("five")})),
			},
		},
		"terminating": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("foo"), types.StringValue("bar"), types.StringValue("baz")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("one"), types.StringValue("two"), types.StringValue("three")}),
					types.Int32Value(1),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three")})),
			},
		},
		// this returns two errors, and as far as I can tell this cannot be compared with func (*FuncError) Equal as it involves two structs
		// see also "elements_delete"
		/*"list-length": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo")}),
					types.Int32Value(0),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "replace: list parameter must not be empty"),
				Result: resultData,
			},
		},*/
		"replace-values-length": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo")}),
					types.ListValueMust(types.StringType, []attr.Value{}),
					types.Int32Value(0),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "replace: replace values parameter must not be empty"),
				Result: resultData,
			},
		},
		"negative-index": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar")}),
					types.Int32Value(-1),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(2, "replace: index parameter must not be a negative number"),
				Result: resultData,
			},
		},
		"too-few-end-index": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar")}),
					types.Int32Value(0),
					types.TupleValueMust([]attr.Type{types.Int32Type}, []attr.Value{types.Int32Value(-1)}),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(3, "replace: end index parameter must be greater than or equal to the index parameter"),
				Result: resultData,
			},
		},
		"out-of-bounds": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar"), types.StringValue("two"), types.StringValue("three")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one")}),
					types.Int32Value(3),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(3, "replace: The index at which to replace the values added to the length of the replacement values (i.e. 'endIndex') cannot be greater than the length of the list where the values will be replaced as that would be out of range."),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, slicefunc.NewReplaceFunction(), test)
}
