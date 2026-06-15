import "slices"
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {

	slices.SortFunc(intervals, func (a, b []int) int {
		return a[0] - b[0]
	})

	res := [][]int{}

	i := 0

	for i < len(intervals) {
		j := i
		start := intervals[i][0]
		end := intervals[i][1]

		for j < len(intervals) && intervals[j][0] <= end {
			end = max(end, intervals[j][1])
			j++
		}
		j--

		res = append(res, []int{start, end})
		i = j+1
	}

	return res
}
