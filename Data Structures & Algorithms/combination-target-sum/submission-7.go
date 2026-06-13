
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

	var find func(accum, start int) 

	find = func(accum, start int) {
		if accum == target {
			collect()
			return
		}

		for j := start; j < len(nums); j++ {
			v := nums[j]

			if (accum + v) > target {
				continue
			}
			push(v)
			find(accum + v, j)
			pop()
		}
	}

	find(0, 0)

	return res
}
