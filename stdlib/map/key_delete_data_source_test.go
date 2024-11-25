package mapfunc_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	provider "github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccKeyDelete(test *testing.T) {
	// init input params
	key := "foo"

	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test error when key does not exist
			{
				Config: `data "stdlib_key_delete" "test" {
                  map = { "hello" = "world", "foo" = "bar" }
                  key = "bar"
                }`,
				ExpectError: regexp.MustCompile("The key to be deleted 'bar' does not exist in the input map"),
			},
			// test basic key removal from map
			{
				Config: fmt.Sprintf(`data "stdlib_key_delete" "test" {
                  map = { "hello" = "world", "foo" = "bar" }
                  key = "%s"
                }`, key),
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_key_delete.test", "key", key),
					resource.TestCheckResourceAttr("data.stdlib_key_delete.test", "map.%", "2"),
					resource.TestCheckResourceAttr("data.stdlib_key_delete.test", "map.hello", "world"),
					resource.TestCheckResourceAttr("data.stdlib_key_delete.test", "map.foo", "bar"),
					// verify map removal result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_key_delete.test", "result.%", "1"),
					resource.TestCheckResourceAttr("data.stdlib_key_delete.test", "result.hello", "world"),
					resource.TestCheckNoResourceAttr("data.stdlib_key_delete.test", fmt.Sprintf("result.%s", key)),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_key_delete.test", "id", key),
				),
			},
		},
	})
}
