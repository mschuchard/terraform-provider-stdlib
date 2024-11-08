package mapfunc_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	provider "github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccHasKey(test *testing.T) {
	// initialize test params
	resourceConfig := `data "stdlib_has_key" "test_%s" {
    map = { "hello" = "world", "foo" = "bar" }
    key = "%s"
  }`
	paramsResults := map[string]bool{"foo": true, "bar": false}

	// iterate through tests
	for key, result := range paramsResults {
		// init data source name for this iteration
		dataSourceName := fmt.Sprintf("data.stdlib_has_key.test_%s", key)

		// invoke test
		resource.Test(test, resource.TestCase{
			ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				// test basic key existence in map
				{
					Config: fmt.Sprintf(resourceConfig, key, key),
					Check: resource.ComposeAggregateTestCheckFunc(
						// verify input params are stored correctly
						resource.TestCheckResourceAttr(dataSourceName, "key", key),
						resource.TestCheckResourceAttr(dataSourceName, "map.%", "2"),
						resource.TestCheckResourceAttr(dataSourceName, "map.hello", "world"),
						resource.TestCheckResourceAttr(dataSourceName, "map.foo", "bar"),
						// verify key existence result is stored correctly
						resource.TestCheckResourceAttr(dataSourceName, "result", strconv.FormatBool(result)),
						// verify id stored correctly
						resource.TestCheckResourceAttr(dataSourceName, "id", key),
					),
				},
			},
		})
	}
}
