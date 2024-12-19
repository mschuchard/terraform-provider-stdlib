package stringfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	provider "github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccCut(test *testing.T) {
	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test basic string cut
			{
				Config: `data "stdlib_cut" "test" {
				  param     = "foobarbaz"
				  separator = "bar"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_cut.test", "param", "foobarbaz"),
					resource.TestCheckResourceAttr("data.stdlib_cut.test", "separator", "bar"),
					// verify before, after, and found are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_cut.test", "before", "foo"),
					resource.TestCheckResourceAttr("data.stdlib_cut.test", "after", "baz"),
					resource.TestCheckResourceAttr("data.stdlib_cut.test", "found", "true"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_cut.test", "id", "foobarbaz"),
				),
			},
			// test basic string cut absent separator
			{
				Config: `data "stdlib_cut" "test" {
					param     = "foobarbaz"
					separator = "pizza"
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_cut.test", "param", "foobarbaz"),
					resource.TestCheckResourceAttr("data.stdlib_cut.test", "separator", "pizza"),
					// verify before, after, and found are stored correctly
					resource.TestCheckResourceAttr("data.stdlib_cut.test", "before", "foobarbaz"),
					resource.TestCheckResourceAttr("data.stdlib_cut.test", "after", ""),
					resource.TestCheckResourceAttr("data.stdlib_cut.test", "found", "false"),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_cut.test", "id", "foobarbaz"),
				),
			},
		},
	})
}
