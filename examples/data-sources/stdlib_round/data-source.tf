# Return the rounding of of 1.2
data "stdlib_round" "down" {
  param = 1.2
}
# result => 1

# Return the rounding of 1.8
data "stdlib_round" "up" {
  param = 1.8
}
# result => 2

# Return the rounding of 1.5
data "stdlib_round" "half" {
  param = 1.5
}
# result => 2
