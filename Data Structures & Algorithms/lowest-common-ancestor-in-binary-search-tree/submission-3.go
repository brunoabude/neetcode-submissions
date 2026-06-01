/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	var dfs func(node *TreeNode, depth int)
	lca := root

	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return 
		}

		if p.Val < node.Val && q.Val < node.Val {
			dfs(node.Left, depth+1)
			return
		}
		
		if p.Val > node.Val && q.Val > node.Val {
			dfs(node.Right, depth+1)
			return
		}

		// p and q are on different sides, lca can't get any lower.
		lca = node
	}

	dfs(root, 0)

	return lca
}
