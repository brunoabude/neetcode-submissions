func hasDuplicate(nums []int) bool {
    m := map[int]bool{}

	for _, n := range nums {
		_, exists := m[n]

		if exists {
			return true
		}

		m[n] = true
	}

	return false
}
