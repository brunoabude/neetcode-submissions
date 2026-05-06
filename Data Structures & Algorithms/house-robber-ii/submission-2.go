func dfs (memo map[int]int, nums []int, i int) int {
	if i >= len(nums) {
		return 0
	}

	if v, e := memo[i]; e {
		return v
	}

	// Skip this house and rob the next one
	rob1 := dfs(memo, nums, i+1)
	// Rob this house
	rob2 := nums[i] + dfs(memo, nums, i+2)

	solution := max(rob1, rob2)
	memo[i] = solution
	return solution
}

func rob(nums []int) int {
	if len(nums) == 0 {
	return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}
	return max(dfs(map[int]int{}, nums[1:], 0), dfs(map[int]int{}, nums[:len(nums)-1], 0))
}
