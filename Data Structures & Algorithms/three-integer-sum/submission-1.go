
func threeSum(nums []int) [][]int {
	m := map[int]int{}

	for i := range len(nums) {
		m[nums[i]] = i
	}
	solution := map[[3]int]bool{}

	for i := range len(nums) {
		for j := range len(nums) {
			sum := nums[i] + nums[j]

			v, e := m[-sum]

			if e && i != j && i != v && j != v {
				s := []int{nums[i], nums[j], nums[v]}

				if s[2] < s[1] {
					s[1], s[2] = s[2],s[1]
				}

				if s[1] < s[0] {
					s[1], s[0] = s[0], s[1]
				}

				if s[2] < s[1] {
					s[1], s[2] = s[2],s[1]
				}

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
