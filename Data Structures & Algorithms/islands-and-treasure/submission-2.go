const inf int = 2147483647
var (
	left        [2]int   = [2]int{0, -1}
	right       [2]int   = [2]int{0, 1}
	up          [2]int   = [2]int{-1, 0}
	down        [2]int   = [2]int{1, 0}
	directions  [][2]int = [][2]int{up, down, left, right}
)

func islandsAndTreasure(grid [][]int) {
	m, n := len(grid), len(grid[0])
	queue := [][2]int{}

	for i := range m {
		for j := range n {
			if grid[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		i, j := cell[0], cell[1]
		
		level := grid[i][j] + 1

		for _, dir := range directions {
			if i+dir[0] < 0 || i+dir[0] >= m || j+dir[1] < 0 || j+dir[1] >= n || grid[i+dir[0]][j+dir[1]] != inf {
				continue	
			}

			grid[i+dir[0]][j+dir[1]] = level
			queue = append(queue, [2]int{i+dir[0], j+dir[1]})
		}
	}
}
