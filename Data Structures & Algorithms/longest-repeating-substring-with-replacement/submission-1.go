func MAX(arr map[byte]int) int {
	res := arr[0]

	for _, v := range arr {
		if v > res {
			res = v
		}
	}

	return res
}

func characterReplacement(s string, k int) int {
	if len(s) == 1 {
		return 1
	}

	freq := map[byte]int{}

	res := 0
	l := 0

	for r := range s {
		if c, e := freq[s[r]]; e {
			freq[s[r]] = c + 1
		} else {
			freq[s[r]] = 1
		}

		for r - l + 1 - MAX(freq) > k {
			freq[s[l]] -= 1
			l++
		}

		res = max(res, r - l + 1)
	}

	return res
}