func climbStairs(n int) int {
	dp := [2]int{0, 1}
	aux := 0
	
	for range n {
		aux = dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = aux
	}

	return aux
}
