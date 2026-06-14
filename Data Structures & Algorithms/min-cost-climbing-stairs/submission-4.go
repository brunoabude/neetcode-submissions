func min(a, b int) int {
	if a < b { return a }
	return b
}

func minCostClimbingStairs(cost []int) int {
    n := len(cost)
	dp := [2]int{}
	var aux int

	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= n; i++ {
		aux = min(dp[0] + cost[i-2], dp[1] + cost[i-1])
		dp[0] = dp[1]
		dp[1] = aux
	}

	return aux
}
