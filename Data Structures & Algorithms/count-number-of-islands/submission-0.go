func numIslands(grid [][]byte) int {
	var dfs func(int,int) bool

	dfs = func(i, j int) bool {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
			return false
		}

		if grid[i][j] == '0' || grid[i][j] == '#' {
			return false
		}

		grid[i][j] = '#'

		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)

		return true
	}

	total := 0

	for i := range len(grid) {
		for j := range len(grid[0]) {
			if dfs(i, j) {
				total++
			}
		}
	}
	return total
}
