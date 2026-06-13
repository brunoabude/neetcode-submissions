import "slices"

func insertSorted(slice []int, value int) []int {
	i, _ := slices.BinarySearch(slice, value)
	return slices.Insert(slice, i, value)
}

func combinationSum(nums []int, target int) [][]int {
	combinations := map[[31]int]bool{}
	currComb := make([]int, len(nums))

	push := func(idx int) {
		currComb[idx]++
	}

	pop := func(idx int) {
		currComb[idx]--
	}

	collect := func() {
		arr := [31]int{}
		j := 1

		for i := range len(nums) {
			for range currComb[i] {
				arr[j] = nums[i]
				arr[0]++
				j++
			}
		}

		combinations[arr] = true
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

		for j := range nums {
			find(accum + v, j)
		}
	}

	for i := range nums {
		find(0, i)
	}

	res := [][]int{}

	for k := range combinations {
		res = append(res, k[1:k[0]+1])
	}

	return res
}
