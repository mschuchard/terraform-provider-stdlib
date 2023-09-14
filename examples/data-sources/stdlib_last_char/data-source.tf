# Return the last character of a string;
# "hello"
# => "o"
data "stdlib_last_char" "hello" {
  param = "hello"
}

# Return the last three characters of a string:
# "hello", 3
# => "llo"
data "stdlib_last_char" "llo" {
  param     = "hello"
  num_chars = 3
}
