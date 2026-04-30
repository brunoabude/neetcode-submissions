func twoSum(nums []int, target int) []int {
    l := map[int]int{}

	for i, n := range nums {
		c := target - n
		l[c] = i 
	}

	for i, n := range nums {
		j, exists := l[n]

		if exists  && j != i {
			return []int{i, j}
		}
	}

	return []int{}
}
