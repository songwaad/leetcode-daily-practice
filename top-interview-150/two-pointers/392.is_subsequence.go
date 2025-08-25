// 392. Is Subsequence
// Link: https://leetcode.com/problems/is-subsequence
// Time Complexity: O(n)

package twopointers

func isSubsequence(s string, t string) bool {
	if len(s) <= 0 {
		return true
	}

	count := 0

	for _, char := range t {
		if count >= len(s) {
			break
		}

		if char == rune(s[count]) {
			count++
			continue
		}
	}

	return count == len(s)
}
