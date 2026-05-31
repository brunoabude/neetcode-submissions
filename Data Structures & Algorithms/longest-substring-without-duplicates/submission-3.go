func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	l, r := 0, 1
	longest := 1

	counter := map[byte]int{}
	counter[s[0]]++

	for r < len(s) {
		// increase window until we find a repeated char
		length := 0

		for r < len(s) {
			d := s[r]
			counter[d]++
			r++

			if counter[d] == 2 {
				length--
				break
			}
		}

		length += r - l

		if length > longest {
			longest = length
		}

	
		for l < r {
			d := s[l]
			counter[d]--
			l++
			if counter[d] == 1 {
				break
			}
		}
	}

	return longest
}
