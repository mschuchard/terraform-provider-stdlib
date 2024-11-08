package slicefunc_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	provider "github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccLastElement(test *testing.T) {
	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test error on invalid number of terminating elements
			{
				Config: `data "stdlib_last_element" "test" {
					param = ["h", "e", "l", "l", "o"]
					num_elements = 10
				}`,
				ExpectError: regexp.MustCompile("The number of terminating elements to return must be fewer than"),
			},
			// test basic list slice last element
			{
				Config: `data "stdlib_last_element" "test" { param = ["h", "e", "l", "l", "o"] }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_last_element.test", "param.#", "5"),
					// verify last element result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_last_element.test", "result.#", "1"),
					resource.TestCheckResourceAttr("data.stdlib_last_element.test", "result.0", "o"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_last_element.test", "id", "h"),
				),
			},
			// test basic reverse list slice last three elements
			{
				Config: `data "stdlib_last_element" "test" {
					param = ["h", "e", "l", "l", "o"]
					num_elements = 3
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_last_element.test", "param.#", "5"),
					// verify last element result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_last_element.test", "result.#", "3"),
					resource.TestCheckResourceAttr("data.stdlib_last_element.test", "result.0", "l"),
					resource.TestCheckResourceAttr("data.stdlib_last_element.test", "result.1", "l"),
					resource.TestCheckResourceAttr("data.stdlib_last_element.test", "result.2", "o"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_last_element.test", "id", "h"),
				),
			},
		},
	})
}
