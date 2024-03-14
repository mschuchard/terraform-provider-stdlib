package slicefunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccMaxString(test *testing.T) {
	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test maximum string
			{
				Config: `data "stdlib_max_string" "test" { param = ["zero", "one", "two", "three", "four", "five", "six", "seven"] }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_max_string.test", "param.#", "8"),
					// verify maximum string result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_max_string.test", "result", "zero"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_max_string.test", "id", "zero"),
				),
			},
			// test maximum string
			{
				Config: `data "stdlib_max_string" "test" { param = ["alpha", "beta", "gamma", "delta", "epsilon"] }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_max_string.test", "param.#", "5"),
					// verify maximum string result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_max_string.test", "result", "gamma"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_max_string.test", "id", "alpha"),
				),
			},
		},
	})
}
