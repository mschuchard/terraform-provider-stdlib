package stringfunc

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &strcontainsMultipleFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewStrcontainsMultipleFunction() function.Function {
	return &strcontainsMultipleFunction{}
}

// function implementation
type strcontainsMultipleFunction struct{}

// function metadata
func (*strcontainsMultipleFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "strcontains_multiple"
}

// define the provider-level definition for the function
func (*strcontainsMultipleFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Checks whether substrings are within another string.",
		MarkdownDescription: "Determine if one or more strings are contained within another string. The optional parameter `all` can be set to true to require that all substrings are contained within the string. If `all` is `false` or not provided, the function will return `true` if any of the substrings are contained within the string. This function is analogous to the core function `strcontains`, but allows for multiple substrings to be checked at once.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "the_string",
				Description: "Input string parameter for checking substrings' existence.",
			},
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "substrings",
				Description: "List of substrings to check for existence within the input string.",
			},
		},
		VariadicParameter: function.BoolParameter{
			Name:                "all",
			MarkdownDescription: "Optional: Whether or not to check if all substrings are contained within the string. The default behavior is any.",
		},
		Return: function.BoolReturn{},
	}
}

func (*strcontainsMultipleFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input parameters
	var theString string
	var substrings []string
	var all bool
	var allVar []bool

	resp.Error = req.Arguments.Get(ctx, &theString, &substrings, &allVar)
	if resp.Error != nil {
		return
	}

	// optional all parameter
	if len(allVar) == 1 {
		// assign variadic parameter value to all
		all = allVar[0]
	}

	ctx = tflog.SetField(ctx, "strcontains_multiple: the_string", theString)
	ctx = tflog.SetField(ctx, "strcontains_multiple: substrings", substrings)
	ctx = tflog.SetField(ctx, "strcontains_multiple: all", all)

	// determine if any or all substrings are contained within the string
	// initialize to false in case the substrings list is empty
	result := false
	// iterate through substrings and check if they are contained within the string
	for _, substring := range substrings {
		// check if current substring is contained in the string
		if strings.Contains(theString, substring) {
			// if "any" then we should break as soon as one substring is valid
			if !all {
				result = true
				break
			}
		} else {
			// if "all" then we should break as soon as one substring is invalid
			if all {
				result = false
				break
			}
		}
		// if we are here then either it is "any" with no valid subtrings, or "all" with all valid substrings
		// therefore we can use the all boolean value to determine the result
		result = all
	}
	ctx = tflog.SetField(ctx, "strcontains_multiple: result", result)

	// store the result as a boolean
	resp.Error = resp.Result.Set(ctx, &result)
	if resp.Error != nil {
		return
	}

	tflog.Debug(ctx, "strcontains_multiple: successful return", map[string]any{"success": true})
}
