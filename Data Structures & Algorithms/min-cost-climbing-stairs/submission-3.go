func min(a, b int) int {
	if a < b { return a }
	return b
}

func minCostClimbingStairs(cost []int) int {
    n := len(cost)
	res := [2]int{}
	dp := [2]int{}
	var aux int

	// Starting from 0
	dp[0] = 0
	dp[1] = cost[0]

	for i := 2; i <= n; i++ {
		aux = min(dp[0] + cost[i-2], dp[1] + cost[i-1])
		dp[0] = dp[1]
		dp[1] = aux
	}

	res[0] = aux

	// Starting from 1
	dp[0] = 0
	dp[1] = min(cost[0], cost[1])

	for i := 3; i <= n; i++ {
		aux = min(dp[0] + cost[i-2], dp[1] + cost[i-1])
		dp[0] = dp[1]
		dp[1] = aux
	}

	res[1] = aux

	return min(res[0], res[1])
}
