import "slices"

func insertSorted(slice []int, value int) []int {
	i, _ := slices.BinarySearch(slice, value)
	return slices.Insert(slice, i, value)
}

func combinationSum(nums []int, target int) [][]int {
	combinations := map[[31]int]bool{}
	currComb := []int{}

	push := func(v int) {
		currComb = append(currComb, v)
	}

	pop := func() {
		currComb = currComb[:len(currComb)-1]
	}

	collect := func() {
		arr := [31]int{}
		slc := slices.Sorted(slices.Values(currComb))
		copy(arr[1:], slc)
		arr[0] = len(currComb)

		combinations[arr] = true
	}

	var find func(accum, i int) 

	find = func(accum, i int) {
		v := nums[i]
		push(v)
		defer pop()

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
