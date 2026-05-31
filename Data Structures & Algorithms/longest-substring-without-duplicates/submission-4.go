func lengthOfLongestSubstring(s string) int {
	longest := 0
	counter := map[byte]int{}
	left := 0

	for right := 0; right < len(s); right++ {
		counter[s[right]]++

		for counter[s[right]] > 1 {
			counter[s[left]]--
			left++
		}

		if right-left+1 > longest {
			longest = right-left+1
		}
	}
	return longest
}
