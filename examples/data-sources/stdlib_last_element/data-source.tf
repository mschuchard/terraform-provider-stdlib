# Return the last element of a list:
data "stdlib_last_element" "hello" {
  param = ["h", "e", "l", "l", "o"]
}
# => ["o"]

# Return the last three elements of a list (reverse slice)
data "stdlib_last_element" "llo" {
  param        = ["h", "e", "l", "l", "o"]
  num_elements = 3
}
# => ["l", "l", "o"]
