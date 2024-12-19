# Return the product of the elements within a set:
data "stdlib_product" "zero" {
  param = [0, 1, 2]
}
# result => 0

data "stdlib_product" "single" {
  params = [5]
}
# result => 5

data "stdlib_product" "normal" {
  params = [1, 2, 3, 4, 5]
}
# result => 120
