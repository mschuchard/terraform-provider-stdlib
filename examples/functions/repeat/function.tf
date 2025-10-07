# Return the list repeated twice:
provider::stdlib::repeat(["zero", "one", "two"], 2)
# result => ["zero", "one", "two", "zero", "one", "two"]

# Return the string repeated twice:
provider::stdlib::repeat("pizza", 2)
# result => "pizzapizza"

# Return the list repeated zero times (empty):
provider::stdlib::repeat(["zero", "one", "two"], 0)
# result => []

# Return the string repeated zero times (empty):
provider::stdlib::repeat("pizza", 0)
# result => ""