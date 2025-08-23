// 80. Remove Duplicates from Sorted Array
// Link: https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii
// Time Complexity: O(n)

package arrayandstring

func removeDuplicatesTwo(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	k := 2
	for i := 2; i < n; i++ {
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
