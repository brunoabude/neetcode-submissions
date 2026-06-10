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
		if node.Left == nil && node.Right == nil {
			return node.Val
		}

		var left int
		var right int

		if node.Left != nil {
			left = dfs(node.Left)
			maxPath = max(maxPath, left)
			maxPath = max(maxPath, left+node.Val)
		}  
		
		if node.Right != nil {
			right = dfs(node.Right)
			maxPath = max(maxPath, right)
			maxPath = max(maxPath, right+node.Val)
		}

		if node.Left != nil && node.Right != nil {
			maxPath = max(maxPath, left + right + node.Val)
			return max(node.Val, max(left, right) + node.Val)
		}  
		
		if node.Left != nil {
			return max(node.Val, left + node.Val)
		}

		return max(node.Val, right + node.Val)
	}

	dfs(root)

	return maxPath
}
