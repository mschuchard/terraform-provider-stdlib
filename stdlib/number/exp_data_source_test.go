package numberfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	provider "github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccExp(test *testing.T) {
	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test basic 0 exponential
			{
				Config: `data "stdlib_exp" "test" { param = 0 }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_exp.test", "param", "0"),
					// verify exponential result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_exp.test", "result", "1"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_exp.test", "id", "0"),
				),
			},
			// test basic float exponential
			{
				Config: `data "stdlib_exp" "test" { param = 1.0986122	}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_exp.test", "param", "1.0986122"),
					// verify exponential result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_exp.test", "result", "2.9999997339956828"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_exp.test", "id", "1.0986122"),
				),
			},
		},
	})
}
