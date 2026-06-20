func countSubstrings(s string) int {
    palindromes := 0

	for i := range len(s) {
		//odd
		for j := 0; i-j >= 0 && i+j < len(s); j++ {
			if s[i-j] == s[i+j] {
				palindromes++
			} else {
				break
			}
		}

		for j := 0; i-j >= 0 && i+j+1 < len(s); j++ {
			if s[i-j] == s[i+j+1] {
				palindromes++
			}else {
				break
			}
		}
	}

	return palindromes
}
