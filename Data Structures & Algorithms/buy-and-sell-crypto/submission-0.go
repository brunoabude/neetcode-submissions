func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	profit, l, r := 0, 0, 0

	for i := range prices {
		if prices[i] < prices[l] {
			l = i
		}

		if prices[i] > prices[l] {
			r = i
		}

		if r > l && prices[r]-prices[l] > profit {
			profit = prices[r] - prices[l]
		}

	}

	return max(profit, 0)
}