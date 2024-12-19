# Return the sorted list.
data "stdlib_insert" "numbers" {
  list_param = [0, 4, -10, 8]
}
# result => ["-10", "0", "4", "8"]

# Return the sorted list.
data "stdlib_insert" "strings" {
  list_param = ["gamma", "beta", "alpha", "delta"]
}
# result => ["alpha", "beta", "delta", "gamma"]
