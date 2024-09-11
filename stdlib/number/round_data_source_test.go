package numberfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccRound(test *testing.T) {
	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test 1.2 round down
			{
				Config: `data "stdlib_round" "test" { param = 1.2 }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_round.test", "param", "1.2"),
					// verify rounding result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_round.test", "result", "1"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_round.test", "id", "1.2"),
				),
			},
			// test 1.8 round up
			{
				Config: `data "stdlib_round" "test" { param = 1.8 }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_round.test", "param", "1.8"),
					// verify rounding result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_round.test", "result", "2"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_round.test", "id", "1.8"),
				),
			},
			// test 1.5 round up
			{
				Config: `data "stdlib_round" "test" { param = 1.5 }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_round.test", "param", "1.5"),
					// verify rounding result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_round.test", "result", "2"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_round.test", "id", "1.5"),
				),
			},
		},
	})
}
