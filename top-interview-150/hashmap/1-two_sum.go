// 1. Two Sum
// Link: https://leetcode.com/problems/two-sum
// Time Complexity: O(n)

package hashmap

func twoSum(nums []int, target int) []int {
	maps := make(map[int]int)

	for i, num := range nums {
		if index, ok := maps[target-num]; ok {
			return []int{index, i}
		}
		maps[num] = i
	}

	return []int{}
}
