func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return min(nums[0], nums[1])
	}

	left, right := 0, len(nums) - 1

	for left < right {
		midpoint := (right + left + 1) / 2

		if nums[midpoint] > nums[right] {
			left = midpoint + 1
		} else {
			right = midpoint 
		}

		if right-left == 1 {
			break
		}
	}

	return min(nums[left], nums[right])
}
