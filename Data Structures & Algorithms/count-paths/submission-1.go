func createMatrix(m int, n int) [][]int {
    matrix := make([][]int, m)

    for i := range m {
        matrix[i] = make([]int, n)
    }

    return matrix
}

func uniquePaths(m int, n int) int {
	matrix := createMatrix(m, n)

	matrix[0][0] = 1

	for i := range m {
		for j := range n {
			moveTo(&matrix, m, n, i, j, i, j+1)
			moveTo(&matrix, m, n, i, j, i+1, j)
		}
	}

	return matrix[m-1][n-1]
}

func moveTo(matrix *[][]int, m, n, i, j, x, y int) {
	if x >= 0 && x < m && y >= 0 && y < n {
		(*matrix)[x][y] += (*matrix)[i][j] 
	}
}