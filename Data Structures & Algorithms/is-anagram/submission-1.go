func isAnagram(s string, t string) bool {
	chars := map[byte]int{}

	sS, sT := len(s), len(t)

	if sS != sT {
		return false
	}

	l := sS

	if sT > l {
		l = sT
	}

	put := func(c byte, v int) {
		_, e := chars[c]

		if !e {
			chars[c] = 0
		}

		chars[c] += v
	}

	for i :=  range l {
		if i < sS {
			put(s[i], 1)
		}

		if i < sT {
			put(t[i], -1)
		}
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}
	return true
}
