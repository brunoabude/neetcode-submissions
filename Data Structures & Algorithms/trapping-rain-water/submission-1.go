func min(a,b int) int {
	if a < b {
		return a
	}

	return b
}

func trap(height []int) int {
	trapped_water := 0

	prefix := make([]int, len(height))
	suffix := make([]int, len(height))
	
	prefix[0] = height[0]

	for i := 1; i < len(height); i++ {
		if height[i] > prefix[i-1] {
			prefix[i] = height[i]
		} else {
			prefix[i] = prefix[i-1]
		}
	}

	suffix[len(height)-1] = height[len(height)-1]

	for i := len(height)-2; i >= 0; i-- {
		if height[i] > suffix[i+1] {
			suffix[i] = height[i]
		} else {
			suffix[i] = suffix[i+1]
		}
	}


	for i := range height {
		t := min(prefix[i], suffix[i]) - height[i]
		if t > 0 {
			trapped_water += t
		}
	}
	return trapped_water
}
