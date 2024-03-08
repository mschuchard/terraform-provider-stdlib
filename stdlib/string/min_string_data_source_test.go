package stringfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccMinString(test *testing.T) {
	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test minimum string
			{
				Config: `data "stdlib_min_string" "test" { param = ["zero", "one", "two", "three", "four", "five", "six", "seven"] }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_min_string.test", "param.#", "8"),
					// verify minimum string result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_min_string.test", "result", "five"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_min_string.test", "id", "zero"),
				),
			},
			// test minimum string
			{
				Config: `data "stdlib_min_string" "test" { param = ["alpha", "beta", "gamma", "delta", "epsilon"] }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_min_string.test", "param.#", "5"),
					// verify minimum string result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_min_string.test", "result", "alpha"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_min_string.test", "id", "alpha"),
				),
			},
		},
	})
}
