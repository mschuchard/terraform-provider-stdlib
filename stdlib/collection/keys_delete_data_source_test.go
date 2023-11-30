package collection_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccKeysDelete(test *testing.T) {
	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test basic keys removal from map
			{
				Config: `data "stdlib_keys_delete" "test" {
          map = { "hello" = "world", "foo" = "bar", "baz" = "bat" }
          keys = ["foo", "baz"]
        }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_keys_delete.test", "keys.#", "2"),
					resource.TestCheckResourceAttr("data.stdlib_keys_delete.test", "keys.0", "foo"),
					resource.TestCheckResourceAttr("data.stdlib_keys_delete.test", "keys.1", "baz"),
					resource.TestCheckResourceAttr("data.stdlib_keys_delete.test", "map.%", "3"),
					resource.TestCheckResourceAttr("data.stdlib_keys_delete.test", "map.hello", "world"),
					resource.TestCheckResourceAttr("data.stdlib_keys_delete.test", "map.foo", "bar"),
					resource.TestCheckResourceAttr("data.stdlib_keys_delete.test", "map.baz", "bat"),
					// verify map removal result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_keys_delete.test", "result.%", "1"),
					resource.TestCheckResourceAttr("data.stdlib_keys_delete.test", "result.hello", "world"),
					resource.TestCheckNoResourceAttr("data.stdlib_keys_delete.test", fmt.Sprintf("result.%s", "foo")),
					resource.TestCheckNoResourceAttr("data.stdlib_keys_delete.test", fmt.Sprintf("result.%s", "baz")),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_keys_delete.test", "id", "foo"),
				),
			},
		},
	})
}
