type Matrix struct {
	data []bool
	m    int
	n    int
}

func newMatrix(m, n int) *Matrix {
	return &Matrix{
		data: make([]bool, m*n),
		m:    m,
		n:    n,
	}
}

func (mt *Matrix) g(i, j int) bool {
	return mt.data[i*mt.n+j]
}

func (mt *Matrix) s(i, j int) {
	mt.data[i*mt.n+j] = true
}

func bfs(visited *Matrix, heights *[][]int, m, n, i, j, currHeight int) {
	if i < 0 || i >= m {
		return
	}

	if j < 0 || j >= n {
		return
	}

	if visited.g(i, j) {
		return
	}

	h := (*heights)[i][j]

	if h < currHeight {
		return
	}

	visited.s(i, j)

	bfs(visited, heights, m, n, i+1, j, h)
	bfs(visited, heights, m, n, i, j+1, h)
	bfs(visited, heights, m, n, i-1, j, h)
	bfs(visited, heights, m, n, i, j-1, h)
}

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])

	pacific := newMatrix(m, n)

	for i := range m {
		bfs(pacific, &heights, m, n, i, 0, heights[i][0])
	}

	for j := range n {
		bfs(pacific, &heights, m, n, 0, j, heights[0][j])
	}

	atlantic := newMatrix(m, n)

	for i := range m {
		bfs(atlantic, &heights, m, n, i, n-1, heights[i][n-1])
	}

	for j := range n {
		bfs(atlantic, &heights, m, n, m-1, j, heights[m-1][j])
	}

	solution := newMatrix(m, n)

	for i := range m {
		for j := range n {
			if (i == 0 || j == 0) && atlantic.g(i, j) {
				solution.s(i, j)
				continue
			}

			if (i == m-1 || j == n-1) && pacific.g(i, j) {
				solution.s(i, j)
				continue
			}

			if pacific.g(i, j) && atlantic.g(i, j) {
				solution.s(i, j)
			}
		}
	}

	result := [][]int{}

	for i := range m {
		for j := range n {
			if solution.g(i, j) {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}