package slicefunc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &productFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewProductFunction() function.Function {
	return &productFunction{}
}

// function implementation
type productFunction struct{}

// function metadata
func (*productFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "product"
}

// define the provider-level definition for the function
func (*productFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Return the product of a set's elements.",
		MarkdownDescription: "Return the product of the elements within a set. This is similar to the core `sum` function, but for a mathematical product instead.",
		Parameters: []function.Parameter{
			function.SetParameter{
				ElementType: types.Float64Type,
				Name:        "set",
				Description: "Input set parameter for determining the product. The set must contain at least one element.",
			},
		},
		Return: function.Float64Return{},
	}
}

func (*productFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize set
	var set []float64

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &set))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "product: set", set)

	// validation
	if len(set) == 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "product: set parameter length must be at least 1"))
	}
	if resp.Error != nil {
		return
	}

	// determine the product
	product := 1.0
	for _, elem := range set {
		product *= elem
	}

	// store the result as a float
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &product))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "product: product", product)
	tflog.Debug(ctx, "product: successful return", map[string]any{"success": true})
}
