// 13. Roman to Integer
// Link: https://leetcode.com/problems/roman-to-integer
// Time Complexity: O(n)

package arrayandstring

func romanToInt(s string) int {
	roman := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0

	for i, char := range s {
		if i < len(s)-1 && roman[rune(char)] < roman[rune(s[i+1])] {
			result -= roman[rune(char)]
			continue
		}

		result += roman[rune(char)]
	}

	return result
}
