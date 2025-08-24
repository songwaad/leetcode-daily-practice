// 134. Gas Station
// Link: https://leetcode.com/problems/gas-station
// Time Complexity: O(n)

package arrayandstring

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost := 0, 0
	curGas := 0
	startingPoint := 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]

		curGas += gas[i] - cost[i]
		if curGas < 0 {
			curGas = 0
			startingPoint = i + 1
		}
	}

	if totalGas >= totalCost {
		return startingPoint
	} else {
		return -1
	}
}
