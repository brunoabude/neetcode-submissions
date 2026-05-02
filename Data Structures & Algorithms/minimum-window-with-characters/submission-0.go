func checkAll(m map[byte]int) bool {
	for _, v := range m {
		if v < 0 {
			return false
		}
	}

	return true
}

func minWindow(s string, t string) string {
	chars := map[byte]int{}

	for i := range t {
		if c, e := chars[t[i]]; e {
			chars[t[i]] = c - 1
		} else {
			chars[t[i]] = -1
		}
	}

	resI, resJ := 0, 1001

	l, r := 0, 0

	for l = range s {
		for !checkAll(chars) && r < len(s) {
			if _, e := chars[s[r]]; e {
				chars[s[r]] += 1
			}
			r++
		}

		if checkAll(chars) && (resJ-resI) > (r-l) {
			resI, resJ = l, r
		}

		if _, e := chars[s[l]]; e {
			chars[s[l]] -= 1
		}

	}

	if resJ-resI == 1001 {
		return ""
	}

	return s[resI:resJ]
}