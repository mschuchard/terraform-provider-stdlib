package stringfunc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccLastCharDataSource(test *testing.T) {
	// init input param
	param := "hello"

	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test basic string slice last char
			{
				Config: fmt.Sprintf(`data "stdlib_last_char" "test" { param = "%s" }`, param),
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_last_char.test", "param", param),
					// verify last character result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_last_char.test", "result", param[len(param)-1:]),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_last_char.test", "id", param),
				),
			},
		},
	})
}
