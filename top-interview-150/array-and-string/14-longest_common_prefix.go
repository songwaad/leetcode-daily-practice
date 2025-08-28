// 14. Longest Common Prefix
// Link: https://leetcode.com/problems/longest-common-prefix
// Time Complexity: O(nm)

package arrayandstring

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 1 {
		return strs[0]
	}

	minWord := []int{0, len(strs[0])}
	for index, word := range strs {
		if len(word) < minWord[1] {
			minWord[0] = index
		}
	}

	i := 0
	for i < len(strs[minWord[0]]) {
		for j := 0; j < len(strs); j++ {
			if j == minWord[0] {
				continue
			}

			if strs[minWord[0]][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
		i++
	}

	return strs[0][:i]
}
