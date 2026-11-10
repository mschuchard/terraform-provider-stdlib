# Return the list repeated twice:
provider::stdlib::repeat(tolist(["zero", "one", "two"]), 2)
# result => ["zero", "one", "two", "zero", "one", "two"]

# Return the string repeated twice:
provider::stdlib::repeat("pizza", 2)
# result => "pizzapizza"

# Return the tuple repeated twice:
provider::stdlib::repeat(["zero", true], 2)
# result => ["zero", true, "zero", true]

# Return the list repeated zero times (empty):
provider::stdlib::repeat(tolist(["zero", "one", "two"]), 0)
# result => []

# Return the string repeated zero times (empty):
provider::stdlib::repeat("pizza", 0)
# result => ""

# Return the tuple repeated zero times (empty):
provider::stdlib::repeat(["zero", true], 0)
# result => [])