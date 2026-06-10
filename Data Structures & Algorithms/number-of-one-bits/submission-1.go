func hammingWeight(n int) int {
	res := 0

	for i := range 32 {
		if (n & (0x1<<i)) != 0 {
			res++
		}
	}

	return res
}
