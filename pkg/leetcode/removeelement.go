package leetcode

/*
Remove Element
https://leetcode.com/problems/remove-element/
*/

func removeElement(nums []int, val int) int {

	j := 0
	for _, num := range nums {
		if num != val {
			nums[j] = num
			j++
		}
	}
	return j
}
