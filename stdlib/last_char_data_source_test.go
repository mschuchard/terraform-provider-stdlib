package stdlib

import (
  "testing"
  "fmt"

  "github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLastCharDataSource(test *testing.T) {
  // init input param
  param := "hello"

  // invoke test
  resource.Test(test, resource.TestCase{
    ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
    Steps: []resource.TestStep{
      // test basic string slice last char
      {
        Config: providerConfig + fmt.Sprintf(`data "stdlib_last_char" "test" { param = "%s" }`, param),
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
