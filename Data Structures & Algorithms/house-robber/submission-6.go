var memo map[int]int = map[int]int{}

func naive(nums *[]int, i int) int {
	if v, e := memo[i]; e {
		return v
	}

	solution := 0

	for j := i+2; j < len(*nums); j++ {
		solution = max(solution, naive(nums, j))
	}

	memo[i] = solution + (*nums)[i]
	return memo[i]
}

func rob(nums []int) int {
	clear(memo)
	rsp := 0 

	for i := range nums {
		rsp = max(rsp, naive(&nums, i))
	}

	return rsp
}
