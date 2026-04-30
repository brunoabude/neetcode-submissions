func dfs(visited map[int]bool, graph map[int][]int, i int) int {
	if visited[i] {
		return -1
	}

	visited[i] = true

	for iE := len(graph[i])-1; iE >= 0; iE-- {

		if dfs(visited, graph, graph[i][iE]) == -1 {
			return -1
		}

		graph[i] = graph[i][:iE]
	}

	visited[i] = false
	return 0
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := map[int][]int{}

	for i := range len(prerequisites) {
		reqs := prerequisites[i]

		for j := len(reqs) - 1; j >= 1; j-- {
			_, exists := graph[reqs[j-1]]

			if !exists {
				graph[reqs[j-1]] = []int{}
			}

			graph[reqs[j-1]] = append(graph[reqs[j-1]], reqs[j])
		}
	}

	for i := range numCourses {
		visited := map[int]bool{}
		if dfs(visited, graph, i) == -1 {
			return false
		}
	}

	return true
}