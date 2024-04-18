package slicefunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccInsert(test *testing.T) {
	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test list values insert
			{
				Config: `data "stdlib_insert" "test" {
				  list_param    = ["one", "two", "three"]
				  insert_values = ["zero"]
				  index         = 0
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "list_param.#", "3"),
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "insert_values.#", "1"),
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "index", "0"),
					// verify inserted values list result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "result.#", "4"),
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "result.0", "zero"),
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "result.3", "three"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "id", "one"),
				),
			},
			{
				Config: `data "stdlib_insert" "test" {
				  list_param    = ["zero", "one", "four", "five"]
				  insert_values = ["two", "three"]
				  index         = 2
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "list_param.#", "4"),
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "insert_values.#", "2"),
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "index", "2"),
					// verify inserted values list result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "result.#", "6"),
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "result.0", "zero"),
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "result.5", "five"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "id", "zero"),
				),
			},
			{
				Config: `data "stdlib_insert" "test" {
				  list_param    = ["zero", "one", "two"]
				  insert_values = ["three"]
				  index         = length(["zero", "one", "two"])
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "list_param.#", "3"),
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "insert_values.#", "1"),
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "index", "3"),
					// verify inserted values list result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "result.#", "4"),
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "result.0", "zero"),
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "result.3", "three"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_insert.test", "id", "zero"),
				),
			},
		},
	})
}
