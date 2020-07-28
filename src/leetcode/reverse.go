package leetcode

import "math"

/*
https://leetcode.com/problems/reverse-integer/
*/
func reverse(x int) int {

	var reverse int
	num := x
	var negative bool = false
	limit := math.Pow(2, 31) - 1

	if num < 0 {
		negative = true
		num = num * -1
	}

	for num > 0 {
		temp := num % 10
		num = num / 10
		reverse = (reverse * 10) + temp
	}

	if reverse > int(limit) || x > int(limit) {
		return 0
	}

	if negative {
		reverse = -1 * reverse
	}

	return reverse
}
