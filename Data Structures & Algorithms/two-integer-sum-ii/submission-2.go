import "slices"

func twoSum(numbers []int, target int) []int {
	if len(numbers) == 2 {
		return []int{1, 2}
	}

	// bruteforce n log n
	for i := range numbers {
		complement := target - numbers[i]
		value, found := slices.BinarySearch(numbers[i+1:], complement)
		if found {
			return []int{1+i, 1+i + 1 + value}
		}
	}

	return []int{-1, -1}
}