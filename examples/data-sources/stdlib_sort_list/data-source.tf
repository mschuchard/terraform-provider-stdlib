# Return the sorted list.
data "stdlib_insert" "numbers" {
  list_param = [0, 4, -10, 8]
}
# => ["-10", "0", "4", "8"]

# Return the sorted list.
data "stdlib_insert" "strings" {
  list_param = ["gamma", "beta", "alpha", "delta"]
}
# => ["alpha", "beta", "delta", "gamma"]
