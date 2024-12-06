package slicefunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	provider "github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccListIndex(test *testing.T) {
	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test list element index
			{
				Config: `data "stdlib_list_index" "test" {
				  list_param = ["zero", "one", "two"]
				  elem_param = "one"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_list_index.test", "list_param.#", "3"),
					resource.TestCheckResourceAttr("data.stdlib_list_index.test", "elem_param", "one"),
					// verify list index result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_list_index.test", "result", "1"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_list_index.test", "id", "zero"),
				),
			},
			{
				Config: `data "stdlib_list_index" "test" {
				  list_param = ["a", "b", "c", "d"]
				  elem_param = "c"
				  sorted     = true
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_list_index.test", "list_param.#", "4"),
					resource.TestCheckResourceAttr("data.stdlib_list_index.test", "elem_param", "c"),
					resource.TestCheckResourceAttr("data.stdlib_list_index.test", "sorted", "true"),
					// verify list index result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_list_index.test", "result", "2"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_list_index.test", "id", "a"),
				),
			},
			{
				Config: `data "stdlib_list_index" "test" {
		  		  list_param = ["zero", "one", "two", "three", "two", "one", "zero"]
		  		  elem_param = "two"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_list_index.test", "list_param.#", "7"),
					resource.TestCheckResourceAttr("data.stdlib_list_index.test", "elem_param", "two"),
					// verify list index result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_list_index.test", "result", "2"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_list_index.test", "id", "zero"),
				),
			},
			{
				Config: `data "stdlib_list_index" "test" {
		  		  list_param = ["hundred", "thousand", "million", "billion"]
	  			  elem_param = "infinity"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_list_index.test", "list_param.#", "4"),
					resource.TestCheckResourceAttr("data.stdlib_list_index.test", "elem_param", "infinity"),
					// verify list index result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_list_index.test", "result", "-1"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_list_index.test", "id", "hundred"),
				),
			},
		},
	})
}
