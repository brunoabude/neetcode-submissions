/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type NodePair struct {
	Original *Node
	Copy *Node
}

func cloneGraph(node *Node) *Node {
	nodes := make([]*Node, 100)

	var dfs func(n *Node) *Node

	dfs = func(n *Node) *Node {
		if n == nil {
			return nil
		}

		if nodes[n.Val-1] != nil {
			return nodes[n.Val-1]
		}

		nodes[n.Val-1] = &Node{
			Val: n.Val,
			Neighbors: []*Node{},
		}

		for _, neighbor := range n.Neighbors {
			nodes[n.Val-1].Neighbors = append(nodes[n.Val-1].Neighbors, dfs(neighbor))
		}

		return nodes[n.Val-1]
	}

	dfs(node)

	return nodes[0]
}
