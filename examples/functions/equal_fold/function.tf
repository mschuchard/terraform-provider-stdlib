# Determine case insensitive equality of two strings:
provider::stdlib::equal_fold("aaa", "AaA")
# result => true

# Determine case insensitive equality of two strings (simple case-folding):
provider::stdlib::equal_fold("AB", "ab")
# result => true

# Determine case insensitive equality of two strings (not full case-folding):
provider::stdlib::equal_fold("ß", "ss")
# result => false