package slicefunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccSortList(test *testing.T) {
	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test sorted list
			{
				Config: `data "stdlib_sort_list" "test" {
				  list_param = [0, 4, -10, 8]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_sort_list.test", "list_param.#", "4"),
					// verify sorted list result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_sort_list.test", "result.#", "4"),
					resource.TestCheckResourceAttr("data.stdlib_sort_list.test", "result.0", "-10"),
					resource.TestCheckResourceAttr("data.stdlib_sort_list.test", "result.1", "0"),
					resource.TestCheckResourceAttr("data.stdlib_sort_list.test", "result.2", "4"),
					resource.TestCheckResourceAttr("data.stdlib_sort_list.test", "result.3", "8"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_sort_list.test", "id", "0"),
				),
			},
			{
				Config: `data "stdlib_sort_list" "test" {
				  list_param = ["gamma", "beta", "alpha", "delta"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_sort_list.test", "list_param.#", "4"),
					// verify sorted list result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_sort_list.test", "result.#", "4"),
					resource.TestCheckResourceAttr("data.stdlib_sort_list.test", "result.0", "alpha"),
					resource.TestCheckResourceAttr("data.stdlib_sort_list.test", "result.1", "beta"),
					resource.TestCheckResourceAttr("data.stdlib_sort_list.test", "result.2", "delta"),
					resource.TestCheckResourceAttr("data.stdlib_sort_list.test", "result.3", "gamma"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_sort_list.test", "id", "gamma"),
				),
			},
		},
	})
}
