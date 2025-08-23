// 169. Majority Element
// Link: https://leetcode.com/problems/majority-element
// Time Complexity: O(n)

package arrayandstring

import "sort"

func majorityElement(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	return nums[n/2]
}
