func min(a, b int) int {
	if a < b { return a }
	return b
}

func minCostClimbingStairs(cost []int) int {
    n := len(cost)
	var dfs func(int) int
	memo := map[int]int{}

	dfs = func(step int) int {
		if step > n {
			return 99999999
		}

		if step == n {
			return 0
		}

		if v, e := memo[step]; e {
			return v
		}
		
		res := cost[step] + min(dfs(step+1), dfs(step+2))
		memo[step] = res
		return res
	}

	return min(dfs(0), dfs(1))
}
