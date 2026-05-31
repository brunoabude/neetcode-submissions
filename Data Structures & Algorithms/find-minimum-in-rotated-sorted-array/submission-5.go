func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	midpoint := 0
	lv := nums[left]
	rv := nums[right]

	for left < right {
		lv = nums[left]
		rv = nums[right]

		if lv < rv {
			return lv
		}

		midpoint = (left+right)>>1

		if nums[midpoint] > rv {
			left = midpoint + 1
		} else {
			right = midpoint 
		}
	}

	return nums[left]
}
