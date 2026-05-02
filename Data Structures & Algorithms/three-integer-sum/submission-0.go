import (
	"slices"
)

func threeSum(nums []int) [][]int {
	matrix := make([]int, len(nums)*len(nums))
	m := map[int]int{}

	idx := func(i, j int) int {
		return i*len(nums) + j
	}

	for i := range len(nums) {
		m[nums[i]] = i
		for j := range len(nums) {
			matrix[idx(i, j)] = nums[i] + nums[j]
		}
	}

	solution := map[[3]int]bool{}

	for i := range len(nums) {
		for j := range len(nums) {
			n := matrix[idx(i, j)]

			v, e := m[-n]

			if e && i != j && i != v && j != v {
				s := []int{nums[i], nums[j], nums[v]}

				slices.Sort(s)

				solution[[3]int(s)] = true
			}
		}
	}

	res := [][]int{}

	for k := range solution {
		res = append(res, []int{k[0], k[1], k[2]})
	}

	return res
}
