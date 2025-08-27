// 209. Minimum Size Subarray Sum
// Link: https://leetcode.com/problems/minimum-size-subarray-sum
// Time Complexity: O(2n)

package slidingwindow

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	n := len(nums)
	result := n + 1
	sum := 0

	for left < n && right < n && left <= right {
		sum += nums[right]
		if sum >= target {
			sum = sum - nums[left] - nums[right]
			result = min(result, right-left+1)
			left++
		} else {
			right++
		}
	}

	if result < n+1 {
		return result
	} else {
		return 0
	}
}
