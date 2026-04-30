func isAnagram(s string, t string) bool {
	chars := map[rune]int{}

	for _, c :=  range s {
		_, e := chars[c]
		if !e {
			chars[c] = 0
		}

		chars[c] += 1
	}

	for _, c :=  range t {
		_, e := chars[c]
		if !e {
			chars[c] = 0
		}

		chars[c] -= 1
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}
