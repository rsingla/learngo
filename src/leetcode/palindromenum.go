package leetcode

/*
https://leetcode.com/problems/palindrome-number/
*/
func isPalindrome(x int) bool {
	var reverse int
	num := x

	if x < 0 {
		return false
	}

	for num > 0 {
		reverse = reverse*10 + num%10
		num = num / 10
	}

	if reverse == x {
		return true
	}
	return false
}
