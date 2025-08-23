// 27. Remove Element
// Link: https://leetcode.com/problems/remove-element
// Time Complexity: O(n)

package arrayandstring

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
