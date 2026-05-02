func longestConsecutive(nums []int) int {
	set := map[int]bool{}

	for _, n := range nums {
		set[n] = true
	}

	result := 0

	for _, n := range nums {
		if _, e:= set[n-1]; !e {
			length := 1
			for range len(nums) {
				if _, nE := set[n+length]; nE {
					length++
				} else {
					break
				}
			}

			if length > result {
				result = length
			}
		}
	}

	return result
}
