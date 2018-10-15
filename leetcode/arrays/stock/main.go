package main

import "fmt"

func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	maxProfit := 0
	maxProfitInIndex := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] > prices[maxProfitInIndex] {
			continue
		}
		iMaxProfit := 0
		for j := i + 1; j < len(prices); j++ {
			tProfit := prices[j] - prices[i]
			if tProfit > iMaxProfit {
				iMaxProfit = tProfit
			}
		}
		if iMaxProfit > maxProfit {
			maxProfit = iMaxProfit
			maxProfitInIndex = i
		}
	}
	return maxProfit
}
func maxProfit2(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	minBuy := prices[0]
	maxVal := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minBuy {
			minBuy = prices[i]
		}
		if prices[i]-minBuy > maxVal {
			maxVal = prices[i] - minBuy
		}
	}
	return maxVal
}

func maxProfitManyTransactions(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	profit := 0
	buy := prices[0]
	t := 0
	for i := 1; i < len(prices); i++ {
		fmt.Printf("%d ", i)
		tt := prices[i] - buy
		if tt > t {
			t = prices[i] - buy
			fmt.Printf("get profit: %d\n", tt)
		} else {
			profit += t
			fmt.Printf("get profit total: %d\n", t)
			buy = prices[i]
			t = 0
		}
	}

	return profit + t
}

func main() {
	fmt.Println(maxProfitManyTransactions([]int{1, 2, 3}))
}
