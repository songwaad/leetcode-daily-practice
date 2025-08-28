// 238. Product of Array Except Self
// Link: https://leetcode.com/problems/product-of-array-except-self
// Time Complexity: O(n^3)

package arrayandstring

func productExceptSelf(nums []int) []int {
	n := len(nums)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = 1
	}

	left := 1
	for i := 0; i < n; i++ {
		arr[i] *= left
		left *= nums[i]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		arr[i] *= right
		right *= nums[i]
	}

	return arr
}
