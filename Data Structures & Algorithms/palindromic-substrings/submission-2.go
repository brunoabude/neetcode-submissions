func countSubstrings(s string) int {
    palindromes := 0

	for i := range len(s) {
		for j := 0; i-j >= 0 && i+j < len(s) && s[i-j] == s[i+j]; j++ {
			palindromes++
		}

		for j := 0; i-j >= 0 && i+j+1 < len(s) && s[i-j] == s[i+j+1]; j++ {
			palindromes++
		}
	}

	return palindromes
}
