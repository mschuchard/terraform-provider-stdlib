package numberfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	provider "github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccMod(test *testing.T) {
	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test basic 0 remainder
			{
				Config: `data "stdlib_mod" "test" {
				  dividend = 4
				  divisor  = 2
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_mod.test", "dividend", "4"),
					resource.TestCheckResourceAttr("data.stdlib_mod.test", "divisor", "2"),
					// verify remainder result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_mod.test", "result", "0"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_mod.test", "id", "2"),
				),
			},
			// test basic integer remainder
			{
				Config: `data "stdlib_mod" "test" {
				  dividend = 5
				  divisor  = 3
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_mod.test", "dividend", "5"),
					resource.TestCheckResourceAttr("data.stdlib_mod.test", "divisor", "3"),
					// verify remainder result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_mod.test", "result", "2"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_mod.test", "id", "3"),
				),
			},
			// test basic float remainder
			{
				Config: `data "stdlib_mod" "test" {
				  dividend = 10
				  divisor  = 4.75
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_mod.test", "dividend", "10"),
					resource.TestCheckResourceAttr("data.stdlib_mod.test", "divisor", "4.75"),
					// verify remainder result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_mod.test", "result", "0.5"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_mod.test", "id", "4.75"),
				),
			},
		},
	})
}
