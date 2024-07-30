package slicefunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccMaxNumber(test *testing.T) {
	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test maximum number
			{
				Config: `data "stdlib_max_number" "test" { param = [0, 1, 1, 2, 3, 5, 8, 13] }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_max_number.test", "param.#", "8"),
					// verify maximum number result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_max_number.test", "result", "13"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_max_number.test", "id", "0"),
				),
			},
		},
	})
}
