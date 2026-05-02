import "slices"

func threeSum(nums []int) [][]int {
	m := map[int]int{}

	for i := range len(nums) {
		m[nums[i]] = i
	}

	solution := map[[3]int]bool{}

	
	for i := range len(nums) {
		for j := i+1; j < len(nums); j++ {
			if i == j {
				continue
			}

			sum := nums[i] + nums[j]

			v, e := m[-sum]

			if e && i != v && j != v {
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
