func missingNumber(nums []int) int {
	expectedSum := 0
	realSum := 0

	for i, v := range nums {
		if v <= len(nums) {
			realSum += v
		}
		expectedSum += i+1
	}

	return expectedSum - realSum
}
