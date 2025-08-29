// 219. Contains Duplicate II
// Link: https://leetcode.com/problems/contains-duplicate-ii
// Time Complexity: O(n)

package hashmap

func containsNearbyDuplicate(nums []int, k int) bool {
	maps := make(map[int]int, 0)

	for j, num := range nums {
		if index, ok := maps[num]; ok {
			if j-index <= k {
				return true
			} else {
				maps[num] = j
			}
		}
		maps[num] = j
	}

	return false
}
