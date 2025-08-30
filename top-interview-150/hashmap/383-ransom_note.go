// 383. Ransom Note
// Link: https://leetcode.com/problems/ransom-note
// Time Complexity: O(mn)

package hashmap

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	maps := make(map[rune]int)

	for _, char := range magazine {
		value, _ := maps[char]
		maps[char] = value + 1
	}

	for _, char := range ransomNote {
		if value, ok := maps[char]; !ok {
			return false
		} else {
			maps[char] = value - 1
			if value <= 0 {
				return false
			}
		}
	}

	return true
}
