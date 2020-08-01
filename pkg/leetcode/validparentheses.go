package leetcode

/*
	https://leetcode.com/problems/valid-parentheses/
*/
func isValid(s string) bool {

	arr := make([]rune, len(s))

	if len(s)%2 == 1 {
		return false
	}

	j := 0
	result := true

	for _, val := range s {

		if val == '(' || val == '[' || val == '{' {
			arr[j] = val
			j = j + 1
		} else {

			if j == 0 {
				return false
			}

			previous := arr[j-1]
			if val == ')' && previous == '(' {
				j = j - 1
			} else if val == ']' && previous == '[' {
				j = j - 1
			} else if val == '}' && previous == '{' {
				j = j - 1
			} else {
				return false
			}
		}
	}

	if j > 0 {
		return false
	}

	return result
}

var m = map[rune]rune{
	'[': ']',
	'{': '}',
	'(': ')',
}

func isValidUpdated(s string) bool {
	stack := make([]rune, len(s))
	idx := -1
	for _, v := range s {
		if idx == -1 {
			idx++
			stack[idx] = v
		} else {
			if m[stack[idx]] == v {
				idx--
			} else {
				idx++
				stack[idx] = v
			}
		}
	}
	return idx == -1
}
