# Return the base e exponential of 0
data "stdlib_exp" "zero" {
  param = 0
}
# => 1

# Return the base e exponential of 1.0986122
data "stdlib_exp" "deciamal" {
  param = 1.0986122
}
# => 2.9999997339956828
