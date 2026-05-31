func min(a,b int) int {
	if a < b {
		return a
	}

	return b
}

func trap(height []int) int {
	trapped_water := 0

	for i := range height {
		h := height[i]
		left_max, right_max := h, h

		for l := i-1; l >= 0; l-- {
			if height[l] > left_max {
				left_max = height[l]
			}
		}

		for r := i+1; r < len(height); r++ {
			if height[r] > right_max {
				right_max = height[r]
			}
		}

		trapped_water += min(left_max, right_max) - h
	}

	return trapped_water
}
