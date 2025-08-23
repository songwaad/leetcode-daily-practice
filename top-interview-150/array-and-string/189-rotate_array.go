// 189. Rotate Array
// Link: https://leetcode.com/problems/rotate-array
// Time Complexity: O(2n)

package arrayandstring

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	if k == 0 {
		return
	}

	arr := make([]int, n)
	j := 0

	for i := n - k; i < n; i++ {
		arr[j] = nums[i]
		j++
	}

	for i := 0; i < n-k; i++ {
		arr[j] = nums[i]
		j++
	}

	for i := 0; i < n; i++ {
		nums[i] = arr[i]
	}
}
