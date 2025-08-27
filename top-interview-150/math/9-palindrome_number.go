// 9. Palindrome Number
// Link: https://leetcode.com/problems/palindrome-number
// Time Complexity: O(n)

package math

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	num := x
	rev := 0

	for num != 0 {
		rev = (num % 10) + (rev * 10)
		num /= 10
	}

	return rev == x
}
