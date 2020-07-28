package leetcode

/*
	https://leetcode.com/problems/two-sum/
*/
func twoSum(nums []int, target int) []int {

	var indices []int
	complement := make(map[int]int)

	for i, val := range nums {
		complement_target := target - val
		value, ok := complement[val]
		if ok {
			indices = append(indices, value, i)
			return indices
		}
		complement[complement_target] = i
	}
	return indices
}
