package main

func isValid(s string) bool {
	var valid []rune
	runes := []rune(s)
	for _, v := range runes {

		if v == '(' || v == '[' || v == '{' {
			valid = append(valid, v)

		} else {
			if len(valid) == 0 {
				return false
			}
			if v == ')' {
				if valid[len(valid)-1] != '(' {
					return false
				} else {
					valid = valid[:len(valid)-1]
				}
			} else if v == ']' {
				if valid[len(valid)-1] != '[' {
					return false
				} else {
					valid = valid[:len(valid)-1]
				}
			} else if v == '}' {
				if valid[len(valid)-1] != '{' {
					return false
				} else {
					valid = valid[:len(valid)-1]
				}
			}
		}
	}
	if len(valid) != 0 {
		return false
	}
	return true
}
