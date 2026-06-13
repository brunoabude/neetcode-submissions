
func combinationSum(nums []int, target int) [][]int {
	res := [][]int{}

	currComb := make([]int, len(nums))

	push := func(idx int) {
		currComb[idx]++
	}

	pop := func(idx int) {
		currComb[idx]--
	}

	collect := func() {
		arr := [30]int{}
		j := 0

		for i := range len(nums) {
			for range currComb[i] {
				arr[j] = nums[i]
				j++
			}
		}

		res = append(res, arr[:j])
	}

	var find func(accum, i int) 

	find = func(accum, i int) {
		if i >= len(nums) {
			return
		}
		v := nums[i]
		push(i)
		defer pop(i)

		if (accum + v) == target {
			collect()
			return
		}

		if (accum + v) > target {
			return
		}

		for j := i; j < len(nums); j++ {
			find(accum + v, j)
		}
	}

	for i := range nums {
		find(0, i)
	}

	return res
}
