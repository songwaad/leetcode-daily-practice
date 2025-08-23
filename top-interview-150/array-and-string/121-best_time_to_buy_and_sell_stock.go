// 121. Best Time to Buy and Sell Stock
// Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock
// Time Complexity: O(n)

package arrayandstring

func maxProfit(prices []int) int {
	maxProfit := 0
	maxCur := 0
	for i := 1; i < len(prices); i++ {
		maxCur = max(0, maxCur+prices[i]-prices[i-1])
		maxProfit = max(maxCur, maxProfit)
	}

	return maxProfit
}
