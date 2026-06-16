const inf int = 2147483647

func islandsAndTreasure(grid [][]int) {
    var dfs func(int, int, int)
	m, n := len(grid), len(grid[0])

	dfs = func(i, j, treasureDist int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}

		if grid[i][j] == -1 {
			return
		}

		if grid[i][j] != inf && grid[i][j] < treasureDist {
			return
		}

		grid[i][j] = treasureDist

		dfs(i+1, j, treasureDist+1)
		dfs(i-1, j, treasureDist+1)
		dfs(i, j+1, treasureDist+1)
		dfs(i, j-1, treasureDist+1)
	}

	for i := range m {
		for j := range n {
			if grid[i][j] == 0 {
				dfs(i, j, 0)
			}
		}
	}
}
