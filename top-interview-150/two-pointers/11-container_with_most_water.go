// 11. Container With Most Water
// Link: https://leetcode.com/problems/container-with-most-water
// Time Complexity: O(n)

package twopointers

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0
	curArea := 0

	for left < right {
		curArea = min(height[left], height[right]) * (right - left)
		maxArea = max(maxArea, curArea)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
