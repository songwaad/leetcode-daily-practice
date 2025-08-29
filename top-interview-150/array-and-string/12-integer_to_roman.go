// 12. Integer to Roman
// https://leetcode.com/problems/integer-to-roman
// Time Complexity: O(n)

package arrayandstring

func intToRoman(num int) string {
	result := ""
	romanMaps := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	romanNums := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	for i := len(romanNums) - 1; i >= 0; i-- {
		for num >= romanNums[i] {
			num -= romanNums[i]
			result += romanMaps[romanNums[i]]
		}
	}

	return result
}
