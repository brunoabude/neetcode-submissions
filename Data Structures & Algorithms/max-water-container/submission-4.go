func maxArea(heights []int) int {
	if len(heights) == 2 {
		return min(heights[0], heights[1])
	}

	n := len(heights) - 1

	l, r := 0, n
	maxVolume := min(heights[0], heights[n]) * n

	for range heights {
		volume := (r - l) * min(heights[l], heights[r])

		if volume > maxVolume  {
			maxVolume = volume
		}

		if heights[l] <= heights[r] {
			l++
		} else {
			r--
		}


		if r <= l {
			break
		}
	}

	return maxVolume
}
