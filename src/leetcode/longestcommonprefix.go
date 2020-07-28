package leetcode

/*
https://leetcode.com/problems/longest-common-prefix/
*/
func longestCommonPrefix(strs []string) string {

	common := ""

	if len(strs) == 0 {
		return common
	}

	common = strs[0]

	for i := 1; i < len(strs); i++ {
		length := len(common)
		if len(common) > len(strs[i]) {
			length = len(strs[i])
		}
		currstr := strs[i]
		k := 0
		for j := 0; j < length; j++ {
			if common[j] == currstr[j] {
				k = k + 1
				continue
			} else {
				break
			}
		}
		common = currstr[0:k]
	}
	return common
}
