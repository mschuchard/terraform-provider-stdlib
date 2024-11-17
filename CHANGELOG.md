### 1.5.1
- Refine attribute error messages.
- Coerce result of `round` function to integer.
- Do not coerce `id` of `min` and `max` functions to integer.

### 1.5.0
- Do not coerce `number` type ID to `string` type ID (may cause superficial plan changes to existing states).
- Add `exp`, `mod`, `round`, and `sqrt` functions.

### 1.4.1
- Add `sorted` parameter to `list_index` function.
- Add `end_index` parameter to `replace` function.
- Fix `end_index` auto-deduction in `replace` function.
- Add config validation to `insert` and `replace` functions.

### 1.4.0
- Add `insert`, `replace`, and `sort_list` functions.

### 1.3.0
- Add `list_index`, `max_number`, `min_number`, `max_string`, and `min_string` functions.

### 1.2.0
- Add `empty`, `compare_list`, and `last_element` functions.

### 1.1.1
- Add `all` parameter to `has_keys` and `has_values`.

### 1.1.0
- Add `num_chars` parameter to `last_char`.
- Add `equal_map`, `keys_delete`, `has_keys`, and `has_values` functions.
- Miscellaneous fixes and improvements to collection functions.

### 1.0.0
- Initial Release
