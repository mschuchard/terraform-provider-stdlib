package slicefunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	provider "github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccProduct(test *testing.T) {
	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test product
			{
				Config: `data "stdlib_product" "test" { param = [0, 1, 2] }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_product.test", "param.#", "3"),
					// verify product result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_product.test", "result", "0"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_product.test", "id", "0"),
				),
			},
			{
				Config: `data "stdlib_product" "test" { param = [5] }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_product.test", "param.#", "1"),
					// verify product result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_product.test", "result", "5"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_product.test", "id", "5"),
				),
			},
			{
				Config: `data "stdlib_product" "test" { param = [1, 2, 3, 4, 5] }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_product.test", "param.#", "5"),
					// verify product result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_product.test", "result", "120"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_product.test", "id", "1"),
				),
			},
		},
	})
}
