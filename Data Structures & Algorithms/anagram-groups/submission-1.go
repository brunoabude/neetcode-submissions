
func key(st string) [26]int{
	arr := [26]int{}

	for _, c := range st {
		arr[c-'a'] += 1
	}

	return arr
}

func groupAnagrams(strs []string) [][]string {
	groups := map[[26]int][]string{}

	for _, str := range strs {
		k := key(str)
		g, ok := groups[k]

		if !ok{
			groups[k] = []string{}
		}

		groups[k] = append(g, str)
	}

	solution := [][]string{}

	for _, v := range groups {
		solution = append(solution, v)
	}

	return solution
}
