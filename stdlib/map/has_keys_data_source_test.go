package mapfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	provider "github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccHasKeys(test *testing.T) {
	// test basic keys existence in map
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `data "stdlib_has_keys" "test" {
				  map = { "hello" = "world", "foo" = "bar", "baz" = "bat" }
				  keys = ["bar", "foo"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "keys.#", "2"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "keys.0", "bar"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "keys.1", "foo"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "map.%", "3"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "map.hello", "world"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "map.foo", "bar"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "map.baz", "bat"),
					// verify keys existence result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "result", "true"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "id", "bar"),
				),
			},
			{
				Config: `data "stdlib_has_keys" "test" {
					map = { "hello" = "world", "foo" = "bar", "baz" = "bat" }
					keys = ["bar", "pizza"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "keys.#", "2"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "keys.0", "bar"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "keys.1", "pizza"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "map.%", "3"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "map.hello", "world"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "map.foo", "bar"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "map.baz", "bat"),
					// verify keys existence result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "result", "false"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "id", "bar"),
				),
			},
			{
				Config: `data "stdlib_has_keys" "test" {
					map = { "hello" = "world", "foo" = "bar", "baz" = "bat" }
					keys = ["bar", "foo"]
					all  = true
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "keys.#", "2"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "keys.0", "bar"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "keys.1", "foo"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "map.%", "3"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "map.hello", "world"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "map.foo", "bar"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "map.baz", "bat"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "all", "true"),
					// verify keys existence result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "result", "false"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "id", "bar"),
				),
			},
			{
				Config: `data "stdlib_has_keys" "test" {
					map = { "hello" = "world", "foo" = "bar", "baz" = "bat" }
					keys = ["hello", "foo", "baz"]
					all  = true
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "keys.#", "3"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "keys.0", "hello"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "keys.1", "foo"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "keys.2", "baz"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "map.%", "3"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "map.hello", "world"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "map.foo", "bar"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "map.baz", "bat"),
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "all", "true"),
					// verify keys existence result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "result", "true"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_has_keys.test", "id", "hello"),
				),
			},
		},
	})
}
