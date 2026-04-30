import "slices"

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}

	for _, str := range strs {
		v := []byte(str)
		slices.Sort(v)
		h := string(v)	

		g, ok := groups[h]

		if !ok{
			groups[h] = []string{}
		}

		groups[h] = append(g, str)
	}

	solution := [][]string{}

	for _, v := range groups {
		solution = append(solution, v)
	}

	return solution
}
