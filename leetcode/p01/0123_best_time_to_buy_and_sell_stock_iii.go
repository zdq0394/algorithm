package p01

/*
给定一个数组，它的第 i 个元素是一支给定的股票在第i天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成"两笔"交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii
*/

func maxProfit(prices []int) int {
	count := 2
	// 交易次数len为count + 1, 因为包括0
	// 不能把count写在2的位置，提示Go error: non-constant array bound, 数组的len不能是常量
	dp := make([][3][2]int, len(prices))

	for i := 0; i < len(prices); i++ {
		for j := count; j >= 1; j-- {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}

			dp[i][j][0] = maxNum(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = maxNum(dp[i-1][j-1][0]-prices[i], dp[i-1][j][1])
		}
	}

	return dp[len(prices)-1][count][0]
}

func maxNum(a1, a2 int) int {
	if a1 >= a2 {
		return a1
	}
	return a2
}
