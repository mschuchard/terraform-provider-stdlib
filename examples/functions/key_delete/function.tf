# Remove the "foo" key from a map:
provider::stdlib::key_delete({"hello" = "world", "foo" = "bar"}, "foo")
# result => {"hello" = "world"}
