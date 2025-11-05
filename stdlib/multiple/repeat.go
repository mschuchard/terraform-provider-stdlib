package multiple

import (
	"context"
	"slices"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	util "github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &repeatFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewRepeatFunction() function.Function {
	return &repeatFunction{}
}

// function implementation
type repeatFunction struct{}

// function metadata
func (*repeatFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "repeat"
}

// define the provider-level definition for the function
func (*repeatFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Return a repeated list or string.",
		MarkdownDescription: "Return a list or string that repeats the input list or string the specified number of times.",
		Parameters: []function.Parameter{
			function.DynamicParameter{
				Name:        "repeater",
				Description: "Input list or string parameter for repeating. The list or string must contain at least one element.",
			},
			function.Int32Parameter{
				Name:        "count",
				Description: "Number of times to repeat the input list or string. This cannot be a negative number.",
			},
		},
		Return: function.DynamicReturn{},
	}
}

func (*repeatFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize repeater and count
	var repeater types.Dynamic
	var count int

	resp.Error = req.Arguments.Get(ctx, &repeater, &count)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "repeat: repeater", repeater.String())
	ctx = tflog.SetField(ctx, "repeat: count", count)

	// retrieve underlying dynamic value
	repeaterValue, unknown, null := util.GetDynamicUnderlyingValue(repeater, ctx)
	if unknown || null {
		resp.Error = function.NewArgumentFuncError(0, "repeat: input parameter is unknown or null")
		return
	}

	// validation
	if empty, err := util.IsDynamicEmpty(repeaterValue, 0, ctx); err != nil || empty {
		// invalid type probably
		if err != nil {
			resp.Error = function.ConcatFuncErrors(resp.Error, err)
		} else { // empty value
			resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "repeat: repeater parameter length must be at least 1"))
		}
	}
	if count < 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(1, "repeat: count parameter value cannot be negative"))
	}
	if resp.Error != nil {
		return
	}

	// determine the repeated list or string
	var repeated types.Dynamic
	// check if string type
	if stringValue, ok := repeaterValue.(types.String); ok {
		// repeat the string value
		repeatedString := strings.Repeat(stringValue.ValueString(), count)

		// convert to explicit string type
		repeatedTFString := types.StringValue(repeatedString)

		// convert to dynamic type
		repeated = types.DynamicValue(repeatedTFString)
	} else if list, ok := repeaterValue.(types.List); ok { // check if list type
		// convert list to slice of string
		var listValue []string
		if diags := list.ElementsAs(ctx, &listValue, false); diags.HasError() {
			resp.Error = function.FuncErrorFromDiags(ctx, diags)
			return
		}

		// repeat the slice value
		repeatedList := slices.Repeat(listValue, count)

		// convert to explicit list type
		repeatedTFList, diags := types.ListValueFrom(ctx, types.StringType, repeatedList)
		if diags.HasError() {
			resp.Error = function.FuncErrorFromDiags(ctx, diags)
			return
		}

		// convert to dynamic type
		repeated = types.DynamicValue(repeatedTFList)
	} else {
		resp.Error = function.NewArgumentFuncError(0, "repeat: invalid input parameter type")
		return
	}

	// store the result as a dynamic value
	resp.Error = resp.Result.Set(ctx, &repeated)
	//resp.Error = resp.Result.Set(ctx, &repeated)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "repeat: repeated", repeated)
	tflog.Debug(ctx, "repeat: successful return", map[string]any{"success": true})
}
