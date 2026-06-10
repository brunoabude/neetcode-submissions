/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxPathSum(root *TreeNode) int {
	maxPath := root.Val

	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftMax := max(0, dfs(node.Left))
		rightMax := max(0, dfs(node.Right))
	
		maxPath = max(maxPath, leftMax + rightMax + node.Val)

		return node.Val + max(leftMax, rightMax)
	}

	dfs(root)

	return maxPath
}
