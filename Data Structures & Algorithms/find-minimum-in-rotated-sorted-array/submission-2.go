func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		midpoint := left + (right-left) / 2

		if nums[midpoint] > nums[right] {
			left = midpoint + 1
		} else {
			right = midpoint 
		}
	}

	return nums[left]
}
