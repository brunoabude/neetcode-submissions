
import "slices"

func combinationSum(nums []int, target int) [][]int {
	res := [][]int{}
	currComb := make([]int, 0, len(nums))

	push := func(v int) {
		currComb = append(currComb, v)
	}

	pop := func() {
		currComb = currComb[:len(currComb)-1]
	}

	collect := func() {
		res = append(res, slices.Clone(currComb))
	}

	var find func(accum, i int) 

	find = func(accum, i int) {
		if i >= len(nums) {
			return
		}
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

		for j := i; j < len(nums); j++ {
			find(accum + v, j)
		}
	}

	for i := range nums {
		find(0, i)
	}

	return res
}
