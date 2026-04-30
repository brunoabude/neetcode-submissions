func createMatrix(m int, n int)[]int {
    matrix := make([]int, m*n)
    return matrix
}

func uniquePaths(m int, n int) int {
	matrix := createMatrix(m, n)

	matrix[0] = 1

	for i := range m {
		for j := range n {
			moveTo(&matrix, m, n, i, j, i, j+1)
			moveTo(&matrix, m, n, i, j, i+1, j)
		}
	}

	return matrix[idx(m,n,m-1,n-1)]
}

func idx(m,n,i,j int) int {
    return i * n + j
}

func moveTo(matrix *[]int, m, n, i, j, x, y int) {
	if x >= 0 && x < m && y >= 0 && y < n {
		(*matrix)[idx(m,n,x,y)] += (*matrix)[idx(m,n,i,j)]
	}
}