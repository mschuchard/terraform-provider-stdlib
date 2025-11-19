package numberfunc_test

import (
	"math/big"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	numberfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/number"
)

func TestExpFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.Float64Unknown())

	testCases := util.TestCases{
		"zero-one": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(0)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(1)),
			},
		},
		"float-one-three": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(1.0986122)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(2.9999997339956828)),
			},
		},
		"beyond-upper-limit": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.NumberValue(func() *big.Float { f := new(big.Float); f.SetString("1e+310"); return f }())}),
			},
			Expected: function.RunResponse{
				Result: resultData,
				Error:  function.NewArgumentFuncError(0, "exp: input number is beyond the limits of float64"),
			},
		},
	}

	util.UnitTests(testCases, resultData, numberfunc.NewExpFunction(), test)
}
