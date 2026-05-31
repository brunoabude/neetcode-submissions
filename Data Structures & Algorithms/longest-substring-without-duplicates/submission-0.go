func lengthOfLongestSubstring(s string) int {
	longest := 0

	//Bruteforce
	for i := range s {
		characters := map[byte]int{}
		length := 0

		for j := range s[i:] {
			digit := s[i+j]
			characters[digit]++

			if  characters[digit] == 1 {
				length++
			} else {
				break
			}
		}

		if length > longest {
			longest = length
		}
	}

	return longest
}
