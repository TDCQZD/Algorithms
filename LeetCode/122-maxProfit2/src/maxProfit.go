package src

// 一次遍历
func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}

// 峰谷法
func maxProfit2(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	i, valley, peak, maxprofit := 0, prices[0], prices[0], 0
	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
			valley = prices[i]
		}
		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
			peak = prices[i]
		}
		maxprofit += peak - valley
	}
	return maxprofit

}

// 暴力法
func maxProfit3(prices []int) int {
	return calculate(prices, 0)
}
func calculate(prices []int, s int) int {
	if len(prices) <= s {
		return 0
	}
	max := 0
	for i := s; i < len(prices); i++ {
		maxProfit := 0
		for j := i + 1; j < len(prices); j++ {
			if prices[i] < prices[j] {
				profit := calculate(prices, j+1) + prices[j] - prices[i]
				if profit > maxProfit {
					maxProfit = profit
				}
			}
		}

		if maxProfit > max {
			max = maxProfit
		}

	}
	return max
}
