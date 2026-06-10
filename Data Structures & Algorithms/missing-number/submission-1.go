func missingNumber(nums []int) int {
	complement := 0

	for i, v := range nums {
		complement += i+1 - v
	}
	
	return complement
}
