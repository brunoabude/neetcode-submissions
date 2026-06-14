func climbStairs(n int) int {
	if n == 1 || n == 2 || n == 3 {
		return n
	}

	aux, n1, n2 := 0, 1, 0
	
	for i := n-3; i >= 0; i-- {
		aux =  1 + n1 + n2

		n2 = n1
		n1 = aux
	}

	return 1 + aux
}
