

// Serializes a tree by levels (BFS)
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var st strings.Builder

	var dfs func(*TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			st.WriteString("nil,")
			return
		}

		st.WriteString(fmt.Sprintf("%d,", node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	return st.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	i, j := 0, 0

	nextToken := func() (string, bool) {
		for j = i; j < len(data) && data[j] != ','; j++ {
		}

		if j > i && i < len(data)-1 {
			tmp := data[i:j]
			i = j + 1
			return tmp, false
		}

		return "", true
	}

	var dfs func() *TreeNode

	dfs = func() *TreeNode {
		token, eof := nextToken()

		if eof {
			return nil
		}

		if token == "nil" {
			return nil
		}

		v, _ := strconv.Atoi(token)

		node := &TreeNode{
			Val: v,
		}

		node.Left = dfs()
		node.Right = dfs()

		return node
	}

	return dfs()
}