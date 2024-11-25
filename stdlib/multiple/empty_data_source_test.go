package multiple_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	provider "github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccEmpty(test *testing.T) {
	// test basic keys existence in map
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test list param
			{
				Config: `data "stdlib_empty" "test" { list_param = [] }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_empty.test", "list_param.#", "0"),
					// verify emptiness result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_empty.test", "id", "zero"),
					resource.TestCheckResourceAttr("data.stdlib_empty.test", "result", "true"),
				),
			},
			// test map param
			{
				Config: `data "stdlib_empty" "test" {
                  map_param = { "foo" = "bar" }
                }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_empty.test", "map_param.%", "1"),
					resource.TestCheckResourceAttr("data.stdlib_empty.test", "map_param.foo", "bar"),
					// verify emptiness result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_empty.test", "id", "foo"),
					resource.TestCheckResourceAttr("data.stdlib_empty.test", "result", "false"),
				),
			},
			// test set param
			{
				Config: `data "stdlib_empty" "test" { set_param = ["no"] }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_empty.test", "set_param.#", "1"),
					// verify emptiness result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_empty.test", "id", "no"),
					resource.TestCheckResourceAttr("data.stdlib_empty.test", "result", "false"),
				),
			},
			// test string param
			{
				Config: `data "stdlib_empty" "test" { string_param = "" }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_empty.test", "string_param", ""),
					// verify emptiness result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_empty.test", "id", ""),
					resource.TestCheckResourceAttr("data.stdlib_empty.test", "result", "true"),
				),
			},
		},
	})
}
