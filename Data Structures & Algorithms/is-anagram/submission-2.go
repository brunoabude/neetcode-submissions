import "slices"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sb, tb := []byte(s),[]byte(t)

	slices.Sort(sb)
	slices.Sort(tb)

	return slices.Equal(sb, tb)
}
