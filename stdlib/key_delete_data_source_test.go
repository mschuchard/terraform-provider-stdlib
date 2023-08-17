package stdlib

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccKeyDeleteDataSource(test *testing.T) {
	// init input params
	key := "foo"

	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
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
