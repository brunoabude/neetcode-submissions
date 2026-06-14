func min(a, b int) int {
	if a < b { return a}
	return b
}

func longestPalindrome(s string) string {
	res := [2]int{0, 1}

	for i := range(len(s)-1) {
		if i+1 < len(s) && s[i] == s[i+1] { 
			j := 0
			for i-j >= 0 && i+j+1 < len(s) {
				if s[i-j] == s[i+j+1] {
					j++
				} else {
					break
				}
			}
			j--

			if 2 + j*2 > res[1] - res[0] {
				res[0] = i-j
				res[1] = i+j+2
			}
		} 
		if i >= 1 && s[i-1] == s[i+1] {
			j := 0
			for i-j >= 0 && i+j < len(s) {
				if s[i-j] == s[i+j] {
					j++
				} else {
					break
				}
			}
			j--

			if 1 + j*2 > res[1] - res[0] {
				res[0] = i-j
				res[1] = i+j+1
			}
		}
	}

	return s[res[0]:res[1]]
}
