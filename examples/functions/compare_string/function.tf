# Lexicographically compare two strings of equal length:
provider::stdlib::compare_string("aaa", "aab")
# result => -1

# Lexicographically compare two strings of unequal length:
provider::stdlib::compare_string("aaa", "aaaa")
# result => -1

# Lexicographically compare two strings of equal length:
provider::stdlib::compare_string("abc", "abc")
# result => 0

# Lexicographically compare two empty strings:
provider::stdlib::compare_string("", "")
# result => 0

# Lexicographically compare two strings of equal length:
provider::stdlib::compare_string("abd", "abc")
# result => 1

# Lexicographically compare two strings of unequal length:
provider::stdlib::compare_string("aaaa", "aaa")
# result => 1