package collection_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

func TestAccHasValue(test *testing.T) {
	// initialize test params
	resourceConfig := `data "stdlib_has_value" "test_%s" {
    map = { "hello" = "world", "foo" = "bar" }
    value = "%s"
  }`
	paramsResults := map[string]bool{"foo": false, "bar": true}

	// iterate through tests
	for value, result := range paramsResults {
		// init data source name for this iteration
		dataSourceName := fmt.Sprintf("data.stdlib_has_value.test_%s", value)

		// invoke test
		resource.Test(test, resource.TestCase{
			ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				// test basic value existence in map
				{
					Config: fmt.Sprintf(resourceConfig, value, value),
					Check: resource.ComposeAggregateTestCheckFunc(
						// verify input params are stored correctly
						resource.TestCheckResourceAttr(dataSourceName, "value", value),
						resource.TestCheckResourceAttr(dataSourceName, "map.%", "2"),
						resource.TestCheckResourceAttr(dataSourceName, "map.hello", "world"),
						resource.TestCheckResourceAttr(dataSourceName, "map.foo", "bar"),
						// verify value existence result is stored correctly
						resource.TestCheckResourceAttr(dataSourceName, "result", strconv.FormatBool(result)),
						// verify id stored correctly
						resource.TestCheckResourceAttr(dataSourceName, "id", value),
					),
				},
			},
		})
	}
}
