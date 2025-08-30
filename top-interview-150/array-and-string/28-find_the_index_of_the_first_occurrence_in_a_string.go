// 28. Find The Index of The First Occurrence in a String
// Link: https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string
// Time Complexity: O(mn)

package arrayandstring

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	startIndex := make([]int, 0)

	for index, char := range haystack {
		if char == rune(needle[0]) {
			startIndex = append(startIndex, index)
		}
	}

	var count int
	for i := 0; i < len(startIndex); i++ {
		count = 0
		index := startIndex[i]
		for j := 0; j < len(needle); j++ {
			if haystack[index] != needle[j] {
				index++
				break
			}
			count++
			index++
			if count == len(needle) {
				return startIndex[i]
			}
			if index >= len(haystack) {
				return -1
			}
		}
	}

	return -1
}
