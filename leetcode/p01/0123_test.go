package p01

import (
	"testing"
)

func TestStock3(t *testing.T) {
	a := []int{3, 3, 5, 0, 0, 3, 1, 4}
	if maxProfits(a) != 6 {
		t.FailNow()
	}
}

func maxProfits(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	count := 2
	dp := make([][3][2]int, len(prices))
	for i := 1; i < len(prices); i++ {
		for j := 2; j >= 1; j-- {
			if i == 1 {
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
