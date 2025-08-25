// 125. Valide Palindrome
// Link: https://leetcode.com/problems/valid-palindrome
// Time Complexity: O(n)

package twopointers

import "unicode"

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if !IsAlNum(rune(s[left])) {
			left++
			continue
		}

		if !IsAlNum(rune(s[right])) {
			right--
			continue
		}

		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}

func IsAlNum(char rune) bool {
	if unicode.IsLetter(char) || unicode.IsDigit(char) {
		return true
	} else {
		return false
	}
}
