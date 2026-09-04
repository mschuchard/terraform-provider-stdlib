# Tranpose a map of strings:
provider::stdlib::transpose({"hello" = "world", "foo" = "bar"})
# result => {"world" = "hello", "bar" = "foo"}

# Transpose a map of strings with duplicate values:
provider::stdlib::transpose({"hello" = "world", "foo" = "bar", "baz" = "bar"})
# result => {"world" = "hello", "bar" = "baz"}
