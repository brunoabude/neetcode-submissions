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
	s1 := make([]byte, len(s))
	s2 := make([]byte, len(s))
	l := 0

	for i := range s {
		if !isAlphanumeric(s[i]) {
			continue
		}

		s1[l] = s[i]
		s2[len(s)-1-l] = s[i]

		l++
	}

	for i := range l {
		if !equal(s1[i], s2[len(s)-l+i]) {
			return false
		}
	}

	return true
}