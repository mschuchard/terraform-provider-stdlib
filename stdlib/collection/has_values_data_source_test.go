package collection_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccHasValuesDataSource(test *testing.T) {
	// test basic values existence in map
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `data "stdlib_has_values" "test" {
					map = { "hello" = "world", "foo" = "bar", "baz" = "bat" }
					values = ["foo", "bar"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_has_values.test", "values.#", "2"),
					resource.TestCheckResourceAttr("data.stdlib_has_values.test", "values.0", "foo"),
					resource.TestCheckResourceAttr("data.stdlib_has_values.test", "values.1", "bar"),
					resource.TestCheckResourceAttr("data.stdlib_has_values.test", "map.%", "3"),
					resource.TestCheckResourceAttr("data.stdlib_has_values.test", "map.hello", "world"),
					resource.TestCheckResourceAttr("data.stdlib_has_values.test", "map.foo", "bar"),
					resource.TestCheckResourceAttr("data.stdlib_has_values.test", "map.baz", "bat"),
					// verify values existence result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_has_values.test", "result", "true"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_has_values.test", "id", "foo"),
				),
			},
			{
				Config: `data "stdlib_has_values" "test" {
					map = { "hello" = "world", "foo" = "bar", "baz" = "bat" }
					values = ["foo", "pizza"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_has_values.test", "values.#", "2"),
					resource.TestCheckResourceAttr("data.stdlib_has_values.test", "values.0", "foo"),
					resource.TestCheckResourceAttr("data.stdlib_has_values.test", "values.1", "pizza"),
					resource.TestCheckResourceAttr("data.stdlib_has_values.test", "map.%", "3"),
					resource.TestCheckResourceAttr("data.stdlib_has_values.test", "map.hello", "world"),
					resource.TestCheckResourceAttr("data.stdlib_has_values.test", "map.foo", "bar"),
					resource.TestCheckResourceAttr("data.stdlib_has_values.test", "map.baz", "bat"),
					// verify values existence result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_has_values.test", "result", "false"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_has_values.test", "id", "foo"),
				),
			},
		},
	})
}
