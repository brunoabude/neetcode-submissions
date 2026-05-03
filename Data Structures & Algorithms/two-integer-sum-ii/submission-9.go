func twoSum(numbers []int, target int) []int {
	if len(numbers) == 2 {
		return []int{1, 2}
	}

	l, r := 0, len(numbers)-1

	for range numbers {
		if numbers[l]+numbers[r] > target {
			r--
			continue
		}

		if numbers[l]+numbers[r] < target {
			l++
			continue
		}

		if numbers[l]+numbers[r] == target {
			return []int{1 + l, 1 + r}
		}
	}

	return []int{-1, -1}
}