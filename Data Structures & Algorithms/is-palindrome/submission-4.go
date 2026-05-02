func isAlphanumeric(r byte) bool {
	switch {
	case r >= '0' && r <= '9':
		return true
	case r >= 'a' && r <= 'z':
		return true
	case r >= 'A' && r <= 'Z':
		return true
	default:
		return false
	}
}

func equal(l, r byte) bool {
	if r >= 'A' && r <= 'Z' {
		r = 'a' + r - 'A'
	}

	if l >= 'A' && l <= 'Z' {
		l = 'a' + l - 'A'
	}

	return l == r
}

func isPalindrome(s string) bool {
	l, r := -1, len(s)

	if len(s) <= 1 {
		return true
	}

	for range len(s) {
		// advance l pointer

		l++

		for ; l <= len(s); l++ {
			if l == len(s) {
				return true
			}

			if isAlphanumeric(s[l]) {
				break
			}
		}

		// advance r pointer
		r--

		for ; r >= -1; r-- {
			if r == -1 {
				return true
			}

			if isAlphanumeric(s[r]) {
				break
			}
		}

		if r <= l {
			return true
		}

		if !equal(s[l], s[r]) {
			return false
		}
	}

	return false
}
