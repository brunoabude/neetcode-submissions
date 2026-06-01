/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	lca, maxDepth := root, 0

	updateLca := func(node *TreeNode, depth int) {
		if depth > maxDepth {
			maxDepth = depth
			lca = node
		}
	}

	var dfs func(node *TreeNode, depth int)

	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return 
		}

		if p.Val < node.Val && q.Val < node.Val {
			updateLca(node, depth)
			dfs(node.Left, depth+1)
			return
		}
		
		if p.Val > node.Val && q.Val > node.Val {
			updateLca(node, depth)
			dfs(node.Right, depth+1)
			return
		}

		// p and q are on different sides, lca can get any lower.
		updateLca(node, depth)
	}

	dfs(root, 0)

	return lca
}
