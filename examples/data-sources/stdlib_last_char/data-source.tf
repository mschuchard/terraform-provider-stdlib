# Return the last character of a string:
data "stdlib_last_char" "hello" {
  param = "hello"
}
# result => "o"

# Return the last three characters of a string:
data "stdlib_last_char" "llo" {
  param     = "hello"
  num_chars = 3
}
# result => "llo"
