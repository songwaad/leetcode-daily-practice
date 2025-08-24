// 55. Jump Game
// Link: https://leetcode.com/problems/jump-game
// Time Complexity: O(n)

package arrayandstring

func canJump(nums []int) bool {
	energy := 0

	for i := 0; i < len(nums); i++ {
		if energy < 0 {
			return false
		}

		if nums[i] > energy {
			energy = nums[i]
		}

		energy--
	}

	return true
}
