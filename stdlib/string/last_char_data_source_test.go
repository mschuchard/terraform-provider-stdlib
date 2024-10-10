package stringfunc_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccLastChar(test *testing.T) {
	// init input param
	param := "hello"
	numChars := 3

	// invoke test
	resource.ParallelTest(test, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// test error on invalid number of terminating chars
			{
				Config: fmt.Sprintf(`data "stdlib_last_char" "test" {
					param = "%s"
					num_chars = 10
				}`, param),
				ExpectError: regexp.MustCompile("The number of terminating characters to return must be fewer than the length"),
			},
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
			// test basic string slice last three characters
			{
				Config: fmt.Sprintf(`data "stdlib_last_char" "test" {
					param = "%s"
					num_chars = %d
				}`, param, numChars),
				Check: resource.ComposeAggregateTestCheckFunc(
					// verify input param is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_last_char.test", "param", param),
					// verify last character result is stored correctly
					resource.TestCheckResourceAttr("data.stdlib_last_char.test", "result", param[len(param)-numChars:]),
					// verify id stored correctly
					resource.TestCheckResourceAttr("data.stdlib_last_char.test", "id", param),
				),
			},
		},
	})
}
