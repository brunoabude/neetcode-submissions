import "slices"
func checkInclusion(s1 string, s2 string) bool {
	getFrequencies := func() []int {
		frequencies := make([]int, 26)

		for i := range s1 {
			frequencies[s1[i]-'a']++
		}

		return frequencies
	}

	for i := 0; i < len(s2)-len(s1)+1; i++ {
		frequencies := getFrequencies()

		for j := 0; j < len(s1); j++ {
			frequencies[s2[i+j]-'a']--
		}

		hasNonZero := slices.ContainsFunc(frequencies, func(v int) bool {
			return v != 0
		})

		if !hasNonZero {
			return true
		}
	}
	return false
}
