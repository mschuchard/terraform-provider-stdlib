package collection_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccMinNumber(test *testing.T) {
	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test minimum number
			{
				Config: `data "stdlib_min_number" "test" { param = [0, 1, 1, 2, 3, 5, 8, 13] }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_min_number.test", "param.#", "8"),
					// verify minimum number result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_min_number.test", "result", "0"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_min_number.test", "id", "0.000000"),
				),
			},
		},
	})
}
