package collection_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccEqualMap(test *testing.T) {
	// test basic keys existence in map
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `data "stdlib_equal_map" "test" {
				  map_one = { "hello" = "world" }
				  map_two = { "hello" = "world" }
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_equal_map.test", "map_one.%", "1"),
					resource.TestCheckResourceAttr("data.stdlib_equal_map.test", "map_one.hello", "world"),
					resource.TestCheckResourceAttr("data.stdlib_equal_map.test", "map_two.%", "1"),
					resource.TestCheckResourceAttr("data.stdlib_equal_map.test", "map_two.hello", "world"),
					// verify map equality result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_equal_map.test", "id", "hellohello"),
					resource.TestCheckResourceAttr("data.stdlib_equal_map.test", "result", "true"),
				),
			},
			{
				Config: `data "stdlib_equal_map" "test" {
				  map_one = { "hello" = "world" }
				  map_two = { "foo" = "bar" }
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_equal_map.test", "map_one.%", "1"),
					resource.TestCheckResourceAttr("data.stdlib_equal_map.test", "map_one.hello", "world"),
					resource.TestCheckResourceAttr("data.stdlib_equal_map.test", "map_two.%", "1"),
					resource.TestCheckResourceAttr("data.stdlib_equal_map.test", "map_two.foo", "bar"),
					// verify map equality result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_equal_map.test", "id", "hellofoo"),
					resource.TestCheckResourceAttr("data.stdlib_equal_map.test", "result", "false"),
				),
			},
		},
	})
}
