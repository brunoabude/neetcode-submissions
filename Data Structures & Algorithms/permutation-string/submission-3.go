import "slices"
func checkInclusion(s1 string, s2 string) bool {
	frequencies := make([]int, 26)

	for i := range s1 {
		frequencies[s1[i]-'a']++
	}

	left, right := 0, 0

	for right < len(s2) {
		frequencies[s2[right]-'a']--

		for right-left+1 > len(s1) && left < len(s2) {
			frequencies[s2[left]-'a']++
			left++
		}

		hasNonZeroElement := slices.ContainsFunc(frequencies, func(v int) bool { return v != 0 })

		if !hasNonZeroElement {
			return true
		}

		right++
	}

	return false
}