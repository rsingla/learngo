package leetcode

/*
https://leetcode.com/problems/remove-duplicates-from-sorted-array/
*/

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	prev := nums[0]
	j := 1
	for _, num := range nums {
		if prev != num {
			nums[j] = num
			j++
			prev = num
		}
	}

	return j
}
