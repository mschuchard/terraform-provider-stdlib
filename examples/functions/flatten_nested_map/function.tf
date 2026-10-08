# Example is related to common design pattern algorithm described at https://developer.hashicorp.com/terraform/language/functions/flatten#flattening-nested-structures-for-for_each
# The below example demonstrates how to perform the above transformation in a very compact and clean function invocation as opposed to the obfuscated and messy standard approach described in the above documentation link.

# Flatten a nested map into a flat map:
resource "aws_subnet" "example" {
  for_each = provider::stdlib::flatten_nested_map(var.networks, "subnets")

  vpc_id            = aws_vpc.example[each.value.parent_key].id
  availability_zone = each.value.child_key
  cidr_block        = each.value.child.cidr_block
}
# result => {
#   "<parent_key>.<child_key>" = {
#     parent_key = "<parent_key>",
#     child_key  = "<child_key>",
#     child      = { cidr_block = "<cidr_block>" }
#   }
# }
