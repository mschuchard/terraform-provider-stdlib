package slicefunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	provider "github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccCompareList(test *testing.T) {
	// test basic keys existence in map
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `data "stdlib_compare_list" "test" {
				  list_one = ["foo", "bar", "b"]
				  list_two = ["foo", "bar", "baz"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "list_one.#", "3"),
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "list_one.2", "b"),
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "list_two.#", "3"),
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "list_two.2", "baz"),
					// verify list comparison result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "id", "foofoo"),
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "result", "-1"),
				),
			},
			{
				Config: `data "stdlib_compare_list" "test" {
				  list_one = ["pizza", "cake"]
				  list_two = ["pizza", "cake"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "list_one.#", "2"),
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "list_one.0", "pizza"),
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "list_two.#", "2"),
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "list_two.1", "cake"),
					// verify list comparison result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "id", "pizzapizza"),
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "result", "0"),
				),
			},
			{
				Config: `data "stdlib_compare_list" "test" {
				  list_one = ["super", "hyper", "turbo"]
				  list_two = ["pizza", "cake", "punch"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "list_one.#", "3"),
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "list_one.2", "turbo"),
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "list_two.#", "3"),
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "list_two.1", "cake"),
					// verify list comparison result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "id", "superpizza"),
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "result", "1"),
				),
			},
			{
				Config: `data "stdlib_compare_list" "test" {
				  list_one = ["pizza", "cake", "punch"]
				  list_two = ["pizza", "cake"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input params are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "list_one.#", "3"),
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "list_one.0", "pizza"),
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "list_two.#", "2"),
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "list_two.1", "cake"),
					// verify list comparison result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "id", "pizzapizza"),
					resource.TestCheckResourceAttr("data.stdlib_compare_list.test", "result", "1"),
				),
			},
		},
	})
}
