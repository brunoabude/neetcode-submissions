func climbStairs(n int) int {
	if n == 1 || n == 2 || n == 3 {
		return n
	}

	dp := make([]int, n)
	dp[n-1] = 0
	dp[n-2] = 1

	for i := n-3; i >= 0; i-- {
		dp[i] =  1 + dp[i+1] + dp[i+2]
	}

	return 1 + dp[0]
}
