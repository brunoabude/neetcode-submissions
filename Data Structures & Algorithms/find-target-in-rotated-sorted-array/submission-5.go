func binarySearch(nums []int, target int) int{
	left, right := 0, len(nums) - 1 

	for left <= right {
		mid := (left+right)/2
		m := nums[mid]

		if m == target {
			return mid
		} 

		if m > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}

		return -1
	}

	left, right := 0, len(nums) - 1 

	for left < right {
		r := nums[right]
		mid := (left+right)/2
		m := nums[mid]

		if m > r {
			left = mid + 1
		} else {
			right = mid
		}
	}

	res := binarySearch(nums[:left], target)

	if res != -1 {
		return res
	}

	res = binarySearch(nums[right:], target)

	if res != -1 {
		return res+right
	}

	return -1
}
