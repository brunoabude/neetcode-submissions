/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	maxDepth := 0
	lca  := root

	var dfs func(node *TreeNode, q, p, depth int) 
	var exists func(node *TreeNode, val int) bool

	exists = func(node *TreeNode, val int) bool {
		if node == nil {
			return false
		}

		if node.Val == val {
			return true
		}

		if val < node.Val {
			return exists(node.Left, val)
		}
		
		if val > node.Val {
		  return exists(node.Right, val)
		}

		// should not be reachable.
		return false
	}


	dfs = func(node *TreeNode, q, p, depth int) {
		if node == nil {
			return 
		}

		foundQ := exists(node, q)
		foundP := exists(node, p)

		if foundQ && foundP {
			if depth > maxDepth {
				maxDepth = depth
				lca = node
			}
			dfs(node.Left, q, p, depth+1)
			dfs(node.Right, q, p, depth+1)
		}
	}

	dfs(root, q.Val, p.Val, 0)

	return lca
}
