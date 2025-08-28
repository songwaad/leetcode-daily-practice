// 58. Length of Last Word
// Link: https://leetcode.com/problems/length-of-last-word
// Time Complexity: O(n)

package arrayandstring

import "unicode"

func lengthOfLastWord(s string) int {
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		if !isAlNum(rune(s[i])) {
			continue
		}

		for i >= 0 && isAlNum(rune(s[i])) {
			result++
			i--
		}

		break
	}

	return result
}

func isAlNum(char rune) bool {
	if unicode.IsLetter(char) || unicode.IsDigit(char) {
		return true
	} else {
		return false
	}
}
