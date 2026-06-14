func climbStairs(n int) int {
	var dfs func(int) int
	memo := map[int]int{}

	dfs = func(step int) int {
		if step == n {
			return 1
		}

		if step > n {
			return 0
		}

		if v, e := memo[step]; e {
			return v
		}

		memo[step] = dfs(step+1) + dfs(step+2)

		return memo[step]
	}

	return dfs(0)
}
