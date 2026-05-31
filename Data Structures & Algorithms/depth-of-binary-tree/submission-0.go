/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

    maxDepth := 1

	var dfs func(node *TreeNode, depth int)

	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return 
		}

		if depth > maxDepth {
			maxDepth = depth
		}

		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}


	dfs(root.Left, 2)
	dfs(root.Right, 2)

	return maxDepth
}
