package stdlib

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccFlattenMapDataSource(test *testing.T) {
	// invoke test
	resource.Test(test, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test basic key removal from map
			{
				Config: providerConfig + `data "stdlib_flatten_map" "test" {
          param = [
            { "hello" = "world" },
            { "foo" = "bar" }
          ]
        }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_flatten_map.test", "param.#", "2"),
					resource.TestCheckResourceAttr("data.stdlib_flatten_map.test", "param.0.%", "1"),
					resource.TestCheckResourceAttr("data.stdlib_flatten_map.test", "param.0.hello", "world"),
					resource.TestCheckResourceAttr("data.stdlib_flatten_map.test", "param.1.%", "1"),
					resource.TestCheckResourceAttr("data.stdlib_flatten_map.test", "param.1.foo", "bar"),
					// verify result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_flatten_map.test", "result.%", "2"),
					resource.TestCheckResourceAttr("data.stdlib_flatten_map.test", "result.hello", "world"),
					resource.TestCheckResourceAttr("data.stdlib_flatten_map.test", "result.foo", "bar"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_flatten_map.test", "id", "hello"),
				),
			},
		},
	})
}
