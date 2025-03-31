package slicefunc_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	provider "github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccReplace(test *testing.T) {
	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test error on out of range index
			{
				Config: `data "stdlib_replace" "test" {
				  list_param     = ["foo", "bar", "two", "three"]
				  replace_values = ["zero", "one"]
				  index          = 3
				}`,
				ExpectError: regexp.MustCompile("The index at which to replace the values added to the length"),
			},
			// test list values replace
			{
				Config: `data "stdlib_replace" "test" {
				  list_param     = ["foo", "bar", "two", "three"]
				  replace_values = ["zero", "one"]
				  index          = 0
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "list_param.#", "4"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "replace_values.#", "2"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "index", "0"),
					// verify end_index automatically deduced correctly
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "end_index", "1"),
					// verify replaced values list result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "result.#", "4"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "result.0", "zero"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "result.3", "three"),
					// verify id stored correctly
					//resource.TestCheckResourceAttr("data.stdlib_replace.test", "id", "foo"),
				),
			},
			{
				Config: `data "stdlib_replace" "test" {
				  list_param     = ["zero", "foo", "bar", "baz", "four", "five"]
				  replace_values = ["one", "two", "three"]
				  index          = 1
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "list_param.#", "6"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "replace_values.#", "3"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "index", "1"),
					// verify end_index automatically deduced correctly
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "end_index", "3"),
					// verify replaced values list result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "result.#", "6"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "result.0", "zero"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "result.2", "two"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "result.4", "four"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "id", "zero"),
				),
			},
			{
				Config: `data "stdlib_replace" "test" {
				  list_param     = ["zero", "foo", "bar", "four", "five"]
				  replace_values = ["one"]
				  index          = 1
				  end_index      = 2
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "list_param.#", "5"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "replace_values.#", "1"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "index", "1"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "end_index", "2"),
					// verify replaced values list result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "result.#", "4"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "result.0", "zero"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "result.1", "one"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "result.2", "four"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "result.3", "five"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "id", "zero"),
				),
			},
			{
				Config: `data "stdlib_replace" "test" {
				  list_param     = ["zero", "foo", "bar", "baz"]
				  replace_values = ["one", "two", "three"]
				  index          = length(["zero", "foo", "bar", "baz"]) - length(["one", "two", "three"])
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "list_param.#", "4"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "replace_values.#", "3"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "index", "1"),
					// verify end_index automatically deduced correctly
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "end_index", "3"),
					// verify replaced values list result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "result.#", "4"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "result.0", "zero"),
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "result.3", "three"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_replace.test", "id", "zero"),
				),
			},
		},
	})
}
