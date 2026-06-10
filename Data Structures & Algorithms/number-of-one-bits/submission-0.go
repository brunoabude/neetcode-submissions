func hammingWeight(n int) int {
	res := 0

	for range 32 {
		res += n & 0x1
		n = n >> 1
	}

	return res
}
