func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		if nums[left] < nums[right] {
			return nums[left]
		}
		
		midpoint := (left+right)>>1

		if nums[midpoint] > nums[right] {
			left = midpoint + 1
		} else {
			right = midpoint 
		}
	}

	return nums[left]
}
