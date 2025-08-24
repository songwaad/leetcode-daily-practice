// 45. Jump Game II
// Link: https://leetcode.com/problems/jump-game-ii
// Time Complexity: O(n)

package arrayandstring

func jumpTwo(nums []int) int {
	jumps := 0
	farthest := 0
	curEnd := 0

	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])

		if i == curEnd {
			jumps++
			curEnd = farthest
		}
	}

	return jumps
}
