# Return the separated strings:
data "stdlib_cut" "foobarbaz" {
  param     = "foobarbaz"
  separator = "bar"
}
# before => "foo", after => "baz", found = true

# Return the separated strings with absent separator:
data "stdlib_cut" "pizza" {
  param     = "foobarbaz"
  separator = "pizza"
}
# before => "foobarbaz", after => "", found = false