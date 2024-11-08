package numberfunc_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	provider "github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccSqrt(test *testing.T) {
	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test error on NaN
			{
				Config:      `data "stdlib_sqrt" "test" { param = -1 }`,
				ExpectError: regexp.MustCompile("The square root of the input parameter must return a valid number"),
			},
			// test square root of four
			{
				Config: `data "stdlib_sqrt" "test" { param = 4 }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_sqrt.test", "param", "4"),
					// verify sqrt result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_sqrt.test", "result", "2"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_sqrt.test", "id", "4"),
				),
			},
			// test square root of zero
			{
				Config: `data "stdlib_sqrt" "test" { param = 0 }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_sqrt.test", "param", "0"),
					// verify sqrt result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_sqrt.test", "result", "0"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_sqrt.test", "id", "0"),
				),
			},
			// test square root of two
			{
				Config: `data "stdlib_sqrt" "test" { param = 2 }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_sqrt.test", "param", "2"),
					// verify sqrt result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_sqrt.test", "result", "1.4142135623730951"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_sqrt.test", "id", "2"),
				),
			},
		},
	})
}
